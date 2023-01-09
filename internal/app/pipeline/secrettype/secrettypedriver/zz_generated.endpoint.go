//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by mga tool. DO NOT EDIT.

package secrettypedriver

import (
	"context"
	"errors"
	"github.com/banzaicloud/pipeline/internal/app/pipeline/secrettype"
	"github.com/go-kit/kit/endpoint"
	kitxendpoint "github.com/sagikazarmark/kitx/endpoint"
)

// endpointError identifies an error that should be returned as an endpoint error.
type endpointError interface {
	EndpointError() bool
}

// serviceError identifies an error that should be returned as a service error.
type serviceError interface {
	ServiceError() bool
}

// Endpoints collects all of the endpoints that compose the underlying service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	GetSecretType   endpoint.Endpoint
	ListSecretTypes endpoint.Endpoint
}

// MakeEndpoints returns a(n) Endpoints struct where each endpoint invokes
// the corresponding method on the provided service.
func MakeEndpoints(service secrettype.Service, middleware ...endpoint.Middleware) Endpoints {
	mw := kitxendpoint.Combine(middleware...)

	return Endpoints{
		GetSecretType:   kitxendpoint.OperationNameMiddleware("secrettype.GetSecretType")(mw(MakeGetSecretTypeEndpoint(service))),
		ListSecretTypes: kitxendpoint.OperationNameMiddleware("secrettype.ListSecretTypes")(mw(MakeListSecretTypesEndpoint(service))),
	}
}

// GetSecretTypeRequest is a request struct for GetSecretType endpoint.
type GetSecretTypeRequest struct {
	SecretType string
}

// GetSecretTypeResponse is a response struct for GetSecretType endpoint.
type GetSecretTypeResponse struct {
	SecretTypeDef secrettype.TypeDefinition
	Err           error
}

func (r GetSecretTypeResponse) Failed() error {
	return r.Err
}

// MakeGetSecretTypeEndpoint returns an endpoint for the matching method of the underlying service.
func MakeGetSecretTypeEndpoint(service secrettype.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetSecretTypeRequest)

		secretTypeDef, err := service.GetSecretType(ctx, req.SecretType)

		if err != nil {
			if serviceErr := serviceError(nil); errors.As(err, &serviceErr) && serviceErr.ServiceError() {
				return GetSecretTypeResponse{
					Err:           err,
					SecretTypeDef: secretTypeDef,
				}, nil
			}

			return GetSecretTypeResponse{
				Err:           err,
				SecretTypeDef: secretTypeDef,
			}, err
		}

		return GetSecretTypeResponse{SecretTypeDef: secretTypeDef}, nil
	}
}

// ListSecretTypesRequest is a request struct for ListSecretTypes endpoint.
type ListSecretTypesRequest struct{}

// ListSecretTypesResponse is a response struct for ListSecretTypes endpoint.
type ListSecretTypesResponse struct {
	SecretTypes map[string]secrettype.TypeDefinition
	Err         error
}

func (r ListSecretTypesResponse) Failed() error {
	return r.Err
}

// MakeListSecretTypesEndpoint returns an endpoint for the matching method of the underlying service.
func MakeListSecretTypesEndpoint(service secrettype.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		secretTypes, err := service.ListSecretTypes(ctx)

		if err != nil {
			if serviceErr := serviceError(nil); errors.As(err, &serviceErr) && serviceErr.ServiceError() {
				return ListSecretTypesResponse{
					Err:         err,
					SecretTypes: secretTypes,
				}, nil
			}

			return ListSecretTypesResponse{
				Err:         err,
				SecretTypes: secretTypes,
			}, err
		}

		return ListSecretTypesResponse{SecretTypes: secretTypes}, nil
	}
}
