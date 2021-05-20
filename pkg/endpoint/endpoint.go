package endpoint

import (
	"context"
	model "development-kit/pkg/model"
	service "development-kit/pkg/service"

	endpoint "github.com/go-kit/kit/endpoint"
)

// CreateDemoRequest collects the request parameters for the CreateDemo method.
type CreateDemoRequest struct {
	Req model.DemoCreateRequest `json:"req"`
}

// CreateDemoResponse collects the response parameters for the CreateDemo method.
type CreateDemoResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeCreateDemoEndpoint returns an endpoint that invokes CreateDemo on the service.
func MakeCreateDemoEndpoint(s service.DevelopmentKitService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateDemoRequest)
		rs, err := s.CreateDemo(ctx, req.Req)
		return CreateDemoResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r CreateDemoResponse) Failed() error {
	return r.Err
}

// UpdateDemoRequest collects the request parameters for the UpdateDemo method.
type UpdateDemoRequest struct {
	Req model.DemoUpdateRequest `json:"req"`
}

// UpdateDemoResponse collects the response parameters for the UpdateDemo method.
type UpdateDemoResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeUpdateDemoEndpoint returns an endpoint that invokes UpdateDemo on the service.
func MakeUpdateDemoEndpoint(s service.DevelopmentKitService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateDemoRequest)
		rs, err := s.UpdateDemo(ctx, req.Req)
		return UpdateDemoResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r UpdateDemoResponse) Failed() error {
	return r.Err
}

// GetOneDemoRequest collects the request parameters for the GetOneDemo method.
type GetOneDemoRequest struct {
	Req model.DemoGetRequest `json:"req"`
}

// GetOneDemoResponse collects the response parameters for the GetOneDemo method.
type GetOneDemoResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeGetOneDemoEndpoint returns an endpoint that invokes GetOneDemo on the service.
func MakeGetOneDemoEndpoint(s service.DevelopmentKitService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetOneDemoRequest)
		rs, err := s.GetOneDemo(ctx, req.Req)
		return GetOneDemoResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r GetOneDemoResponse) Failed() error {
	return r.Err
}

// DeleteDemoRequest collects the request parameters for the DeleteDemo method.
type DeleteDemoRequest struct {
	Req model.DemoDeleteRequest `json:"req"`
}

// DeleteDemoResponse collects the response parameters for the DeleteDemo method.
type DeleteDemoResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeDeleteDemoEndpoint returns an endpoint that invokes DeleteDemo on the service.
func MakeDeleteDemoEndpoint(s service.DevelopmentKitService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteDemoRequest)
		rs, err := s.DeleteDemo(ctx, req.Req)
		return DeleteDemoResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r DeleteDemoResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// CreateDemo implements Service. Primarily useful in a client.
func (e Endpoints) CreateDemo(ctx context.Context, req model.DemoCreateRequest) (rs string, err error) {
	request := CreateDemoRequest{Req: req}
	response, err := e.CreateDemoEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateDemoResponse).Rs, response.(CreateDemoResponse).Err
}

// UpdateDemo implements Service. Primarily useful in a client.
func (e Endpoints) UpdateDemo(ctx context.Context, req model.DemoUpdateRequest) (rs string, err error) {
	request := UpdateDemoRequest{Req: req}
	response, err := e.UpdateDemoEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UpdateDemoResponse).Rs, response.(UpdateDemoResponse).Err
}

// GetOneDemo implements Service. Primarily useful in a client.
func (e Endpoints) GetOneDemo(ctx context.Context, req model.DemoGetRequest) (rs string, err error) {
	request := GetOneDemoRequest{Req: req}
	response, err := e.GetOneDemoEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetOneDemoResponse).Rs, response.(GetOneDemoResponse).Err
}

// DeleteDemo implements Service. Primarily useful in a client.
func (e Endpoints) DeleteDemo(ctx context.Context, req model.DemoDeleteRequest) (rs string, err error) {
	request := DeleteDemoRequest{Req: req}
	response, err := e.DeleteDemoEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteDemoResponse).Rs, response.(DeleteDemoResponse).Err
}

// StatusRequest collects the request parameters for the Status method.
type StatusRequest struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// StatusResponse collects the response parameters for the Status method.
type StatusResponse struct {
	Rs  model.HealthCheckResponse `json:"rs"`
	Err error                     `json:"err"`
}

// MakeStatusEndpoint returns an endpoint that invokes Status on the service.
func MakeStatusEndpoint(s service.DevelopmentKitService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(StatusRequest)

		rs, err := s.Status(ctx, req.Name, req.Version)

		return StatusResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r StatusResponse) Failed() error {
	return r.Err
}

// Status implements Service. Primarily useful in a client.
func (e Endpoints) Status(ctx context.Context, name string, version string) (rs model.HealthCheckResponse, err error) {
	request := StatusRequest{
		Name:    name,
		Version: version,
	}
	response, err := e.StatusEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(StatusResponse).Rs, response.(StatusResponse).Err
}
