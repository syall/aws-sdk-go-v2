// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dynamodb

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/crr"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Client provides the API operation methods for making requests to
// DynamoDB. See this package's package overview docs
// for details on the service.
//
// The client's methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type Client struct {
	*aws.Client

	// Service specific configurations. (codegen: service_specific_config.go)

	// Disables the computation and validation of request and response checksums.
	DisableComputeChecksums bool

	endpointCache *crr.EndpointCache
}

// Used for custom client initialization logic
var initClient func(*Client)

// Used for custom request initialization logic
var initRequest func(*Client, *aws.Request)

const (
	ServiceName = "DynamoDB" // Service's name
	ServiceID   = "DynamoDB" // Service's identifier
	EndpointsID = "dynamodb" // Service's Endpoint identifier
)

// New creates a new instance of the client from the provided Config.
//
// Example:
//     // Create a client from just a config.
//     svc := dynamodb.New(myConfig)
func New(config aws.Config) *Client {
	svc := &Client{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   ServiceName,
				ServiceID:     ServiceID,
				EndpointsID:   EndpointsID,
				SigningName:   "dynamodb",
				SigningRegion: config.Region,
				APIVersion:    "2012-08-10",
				JSONVersion:   "1.0",
				TargetPrefix:  "DynamoDB_20120810",
			},
		),
	}

	svc.endpointCache = crr.NewEndpointCache(10)

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(jsonrpc.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(jsonrpc.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(jsonrpc.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(jsonrpc.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc)
	}

	return svc
}

// newRequest creates a new request for a client operation and runs any
// custom request initialization.
func (c *Client) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(c, req)
	}

	return req
}
