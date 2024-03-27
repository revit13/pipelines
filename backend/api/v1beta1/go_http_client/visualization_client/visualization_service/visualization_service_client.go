// Code generated by go-swagger; DO NOT EDIT.

package visualization_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new visualization service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for visualization service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
VisualizationServiceCreateVisualizationV1 visualization service create visualization v1 API
*/
func (a *Client) VisualizationServiceCreateVisualizationV1(params *VisualizationServiceCreateVisualizationV1Params, authInfo runtime.ClientAuthInfoWriter) (*VisualizationServiceCreateVisualizationV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVisualizationServiceCreateVisualizationV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "VisualizationService_CreateVisualizationV1",
		Method:             "POST",
		PathPattern:        "/apis/v1beta1/visualizations/{namespace}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &VisualizationServiceCreateVisualizationV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*VisualizationServiceCreateVisualizationV1OK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
