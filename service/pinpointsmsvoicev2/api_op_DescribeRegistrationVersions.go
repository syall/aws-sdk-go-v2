// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoicev2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoicev2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the specified registration version.
func (c *Client) DescribeRegistrationVersions(ctx context.Context, params *DescribeRegistrationVersionsInput, optFns ...func(*Options)) (*DescribeRegistrationVersionsOutput, error) {
	if params == nil {
		params = &DescribeRegistrationVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeRegistrationVersions", params, optFns, c.addOperationDescribeRegistrationVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeRegistrationVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRegistrationVersionsInput struct {

	// The unique identifier for the registration.
	//
	// This member is required.
	RegistrationId *string

	// An array of RegistrationVersionFilter objects to filter the results.
	Filters []types.RegistrationVersionFilter

	// The maximum number of results to return per each request.
	MaxResults *int32

	// The token to be used for the next set of paginated results. You don't need to
	// supply a value for this field in the initial request.
	NextToken *string

	// An array of registration version numbers.
	VersionNumbers []int64

	noSmithyDocumentSerde
}

type DescribeRegistrationVersionsOutput struct {

	// The Amazon Resource Name (ARN) for the registration.
	//
	// This member is required.
	RegistrationArn *string

	// The unique identifier for the registration.
	//
	// This member is required.
	RegistrationId *string

	// An array of RegistrationVersions objects.
	//
	// This member is required.
	RegistrationVersions []types.RegistrationVersionInformation

	// The token to be used for the next set of paginated results. You don't need to
	// supply a value for this field in the initial request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeRegistrationVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeRegistrationVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeRegistrationVersions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeRegistrationVersions"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDescribeRegistrationVersionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRegistrationVersions(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

// DescribeRegistrationVersionsAPIClient is a client that implements the
// DescribeRegistrationVersions operation.
type DescribeRegistrationVersionsAPIClient interface {
	DescribeRegistrationVersions(context.Context, *DescribeRegistrationVersionsInput, ...func(*Options)) (*DescribeRegistrationVersionsOutput, error)
}

var _ DescribeRegistrationVersionsAPIClient = (*Client)(nil)

// DescribeRegistrationVersionsPaginatorOptions is the paginator options for
// DescribeRegistrationVersions
type DescribeRegistrationVersionsPaginatorOptions struct {
	// The maximum number of results to return per each request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeRegistrationVersionsPaginator is a paginator for
// DescribeRegistrationVersions
type DescribeRegistrationVersionsPaginator struct {
	options   DescribeRegistrationVersionsPaginatorOptions
	client    DescribeRegistrationVersionsAPIClient
	params    *DescribeRegistrationVersionsInput
	nextToken *string
	firstPage bool
}

// NewDescribeRegistrationVersionsPaginator returns a new
// DescribeRegistrationVersionsPaginator
func NewDescribeRegistrationVersionsPaginator(client DescribeRegistrationVersionsAPIClient, params *DescribeRegistrationVersionsInput, optFns ...func(*DescribeRegistrationVersionsPaginatorOptions)) *DescribeRegistrationVersionsPaginator {
	if params == nil {
		params = &DescribeRegistrationVersionsInput{}
	}

	options := DescribeRegistrationVersionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeRegistrationVersionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeRegistrationVersionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeRegistrationVersions page.
func (p *DescribeRegistrationVersionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeRegistrationVersionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.DescribeRegistrationVersions(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeRegistrationVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeRegistrationVersions",
	}
}
