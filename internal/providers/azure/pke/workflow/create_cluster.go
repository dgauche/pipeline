// Copyright © 2019 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package workflow

import (
	"time"

	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"

	"github.com/banzaicloud/pipeline/cluster"
	"github.com/banzaicloud/pipeline/internal/providers/pke/pkeworkflow"
	pkgCluster "github.com/banzaicloud/pipeline/pkg/cluster"
)

const CreateClusterWorkflowName = "pke-azure-create-cluster"

// CreateClusterWorkflowInput
type CreateClusterWorkflowInput struct {
	ClusterID                       uint
	ClusterName                     string
	OrganizationID                  uint
	ResourceGroupName               string
	SecretID                        string
	VirtualNetworkTemplate          VirtualNetworkTemplate
	LoadBalancerTemplate            LoadBalancerTemplate
	PublicIPAddress                 PublicIPAddress
	RoleAssignmentTemplates         []RoleAssignmentTemplate
	RouteTable                      RouteTable
	SecurityGroups                  []SecurityGroup
	VirtualMachineScaleSetTemplates []VirtualMachineScaleSetTemplate
	PostHooks                       pkgCluster.PostHooks
}

func CreateClusterWorkflow(ctx workflow.Context, input CreateClusterWorkflowInput) error {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: 5 * time.Minute,
		StartToCloseTimeout:    10 * time.Minute,
		ScheduleToCloseTimeout: 15 * time.Minute,
		WaitForCancellation:    true,
	}
	cwo := workflow.ChildWorkflowOptions{
		ExecutionStartToCloseTimeout: 30 * time.Minute,
	}
	ctx = workflow.WithChildOptions(workflow.WithActivityOptions(ctx, ao), cwo)

	// Generate CA certificates
	{
		activityInput := pkeworkflow.GenerateCertificatesActivityInput{ClusterID: input.ClusterID}

		err := workflow.ExecuteActivity(ctx, pkeworkflow.GenerateCertificatesActivityName, activityInput).Get(ctx, nil)
		if err != nil {
			setClusterErrorStatus(ctx, input.ClusterID, err)
			return err
		}
	}

	infraInput := CreateAzureInfrastructureWorkflowInput{
		OrganizationID:    input.OrganizationID,
		ClusterID:         input.ClusterID,
		ClusterName:       input.ClusterName,
		SecretID:          input.SecretID,
		ResourceGroupName: input.ResourceGroupName,
		LoadBalancer: LoadBalancerFactory{
			Template: input.LoadBalancerTemplate,
		},
		PublicIPAddress: input.PublicIPAddress,
		RoleAssignments: RoleAssignmentsFactory{
			Templates: input.RoleAssignmentTemplates,
		},
		RouteTable: input.RouteTable,
		ScaleSets: VirtualMachineScaleSetsFactory{
			Templates: input.VirtualMachineScaleSetTemplates,
		},
		SecurityGroups: input.SecurityGroups,
		VirtualNetwork: VirtualNetworkFactory{
			Template: input.VirtualNetworkTemplate,
		},
	}
	err := workflow.ExecuteChildWorkflow(ctx, CreateInfraWorkflowName, infraInput).Get(ctx, nil)
	if err != nil {
		setClusterErrorStatus(ctx, input.ClusterID, err)
		return err
	}

	setClusterCreatingStatus(ctx, input.ClusterID, "waiting for Kubernetes master")

	signalName := "master-ready"
	signalChan := workflow.GetSignalChannel(ctx, signalName)
	signalSelector := workflow.NewSelector(ctx).AddReceive(signalChan, func(c workflow.Channel, more bool) {
		c.Receive(ctx, nil)
		workflow.GetLogger(ctx).Info("Received signal!", zap.String("signal", signalName))
	})
	signalSelector.Select(ctx) // wait for signal

	postHookWorkflowInput := cluster.RunPostHooksWorkflowInput{
		ClusterID: input.ClusterID,
		PostHooks: cluster.BuildWorkflowPostHookFunctions(input.PostHooks, true),
	}

	err = workflow.ExecuteChildWorkflow(ctx, cluster.RunPostHooksWorkflowName, postHookWorkflowInput).Get(ctx, nil)
	if err != nil {
		setClusterErrorStatus(ctx, input.ClusterID, err)
		return err
	}

	return nil
}