package http

import (
	"context"
	endpoint "development-kit/pkg/endpoint"
	"encoding/json"
	"errors"
	http1 "net/http"

	http "github.com/go-kit/kit/transport/http"
	handlers "github.com/gorilla/handlers"
	mux "github.com/gorilla/mux"
	schema "github.com/gorilla/schema"
)

// makeCreateDemoHandler creates the handler logic
func makeCreateDemoHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/api/create-demo").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.CreateDemoEndpoint, decodeCreateDemoRequest, encodeCreateDemoResponse, options...)))
}

// decodeCreateDemoRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
// Add your request body/param/scheme here
func decodeCreateDemoRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.CreateDemoRequest{}
	err := decode(r, nil, nil, &req.Req.Body)
	return req, err
}

// encodeCreateDemoResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeCreateDemoResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUpdateDemoHandler creates the handler logic
func makeUpdateDemoHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("PUT").Path("/api/update-demo").Handler(handlers.CORS(handlers.AllowedMethods([]string{"PUT"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.UpdateDemoEndpoint, decodeUpdateDemoRequest, encodeUpdateDemoResponse, options...)))
}

// decodeUpdateDemoRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
// Add your request body/param/scheme here
func decodeUpdateDemoRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.UpdateDemoRequest{}
	err := decode(r, nil, &req.Req.Scheme, &req.Req.Body)
	return req, err
}

// encodeUpdateDemoResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdateDemoResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetOneDemoHandler creates the handler logic
func makeGetOneDemoHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/api/get-one-demo").Handler(handlers.CORS(handlers.AllowedMethods([]string{"GET"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.GetOneDemoEndpoint, decodeGetOneDemoRequest, encodeGetOneDemoResponse, options...)))
}

// decodeGetOneDemoRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
// Add your request body/param/scheme here
func decodeGetOneDemoRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.GetOneDemoRequest{}
	err := decode(r, &req.Req.Param, nil, nil)
	return req, err
}

// encodeGetOneDemoResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetOneDemoResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeleteDemoHandler creates the handler logic
func makeDeleteDemoHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("DELETE").Path("/api/delete-demo").Handler(handlers.CORS(handlers.AllowedMethods([]string{"DELETE"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.DeleteDemoEndpoint, decodeDeleteDemoRequest, encodeDeleteDemoResponse, options...)))
}

// decodeDeleteDemoRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
// Add your request body/param/scheme here
func decodeDeleteDemoRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.DeleteDemoRequest{}
	err := decode(r, nil, &req.Req.Scheme, nil)
	return req, err
}

// encodeDeleteDemoResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteDemoResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
func ErrorEncoder(_ context.Context, err error, w http1.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}
func ErrorDecoder(r *http1.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}
func decode(r *http1.Request, param interface{}, schemas interface{}, body interface{}) (err error) {
	if param != nil {
		err = schema.NewDecoder().Decode(param, r.URL.Query())
		if err != nil {
			return err
		}
	}
	if schemas != nil {
		vars := mux.Vars(r)
		data, err := json.Marshal(vars)
		if err != nil {
			return err
		}
		err = json.Unmarshal(data, schemas)
		if err != nil {
			return err
		}
	}
	if body != nil {
		err = json.NewDecoder(r.Body).Decode(body)
		if err != nil {
			return err
		}
	}
	return err
}

// This is used to set the http status, see an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/pkg/addtransport/http.go#L133
func err2code(err error) int {
	return http1.StatusInternalServerError
}

type errorWrapper struct {
	Error string `json:"error"`
}

// makeStatusHandler creates the handler logic
func makeStatusHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/api/status").Handler(
		handlers.CORS(handlers.AllowedMethods([]string{"GET"}),
			handlers.AllowedOrigins([]string{"*"}))(http.NewServer(
				endpoints.StatusEndpoint,
				decodeStatusRequest,
				encodeStatusResponse,
				options...)))
}

// decodeStatusRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
// Add your request body/param/scheme here
func decodeStatusRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.StatusRequest{}
	err := decode(r, nil, nil, nil)
	return req, err
}

// encodeStatusResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeStatusResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
