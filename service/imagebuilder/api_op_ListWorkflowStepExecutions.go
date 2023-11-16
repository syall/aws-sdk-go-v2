// Code generated by smithy-go-codegen DO NOT EDIT.

package imagebuilder

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/imagebuilder/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Shows runtime data for each step in a runtime instance of the workflow that you
// specify in the request.
func (c *Client) ListWorkflowStepExecutions(ctx context.Context, params *ListWorkflowStepExecutionsInput, optFns ...func(*Options)) (*ListWorkflowStepExecutionsOutput, error) {
	if params == nil {
		params = &ListWorkflowStepExecutionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListWorkflowStepExecutions", params, optFns, c.addOperationListWorkflowStepExecutionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListWorkflowStepExecutionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListWorkflowStepExecutionsInput struct {

	// The unique identifier that Image Builder assigned to keep track of runtime
	// details when it ran the workflow.
	//
	// This member is required.
	WorkflowExecutionId *string

	// The maximum items to return in a request.
	MaxResults *int32

	// A token to specify where to start paginating. This is the nextToken from a
	// previously truncated response.
	NextToken *string

	noSmithyDocumentSerde
}

type ListWorkflowStepExecutionsOutput struct {

	// The image build version resource ARN that's associated with the specified
	// runtime instance of the workflow.
	ImageBuildVersionArn *string

	// The output message from the list action, if applicable.
	Message *string

	// The next token used for paginated responses. When this field isn't empty, there
	// are additional elements that the service hasn't included in this request. Use
	// this token with the next request to retrieve additional objects.
	NextToken *string

	// The request ID that uniquely identifies this request.
	RequestId *string

	// Contains an array of runtime details that represents each step in this runtime
	// instance of the workflow.
	Steps []types.WorkflowStepMetadata

	// The build version ARN for the Image Builder workflow resource that defines the
	// steps for this runtime instance of the workflow.
	WorkflowBuildVersionArn *string

	// The unique identifier that Image Builder assigned to keep track of runtime
	// details when it ran the workflow.
	WorkflowExecutionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListWorkflowStepExecutionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListWorkflowStepExecutions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListWorkflowStepExecutions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListWorkflowStepExecutions"); err != nil {
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
	if err = addOpListWorkflowStepExecutionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListWorkflowStepExecutions(options.Region), middleware.Before); err != nil {
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

// ListWorkflowStepExecutionsAPIClient is a client that implements the
// ListWorkflowStepExecutions operation.
type ListWorkflowStepExecutionsAPIClient interface {
	ListWorkflowStepExecutions(context.Context, *ListWorkflowStepExecutionsInput, ...func(*Options)) (*ListWorkflowStepExecutionsOutput, error)
}

var _ ListWorkflowStepExecutionsAPIClient = (*Client)(nil)

// ListWorkflowStepExecutionsPaginatorOptions is the paginator options for
// ListWorkflowStepExecutions
type ListWorkflowStepExecutionsPaginatorOptions struct {
	// The maximum items to return in a request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListWorkflowStepExecutionsPaginator is a paginator for
// ListWorkflowStepExecutions
type ListWorkflowStepExecutionsPaginator struct {
	options   ListWorkflowStepExecutionsPaginatorOptions
	client    ListWorkflowStepExecutionsAPIClient
	params    *ListWorkflowStepExecutionsInput
	nextToken *string
	firstPage bool
}

// NewListWorkflowStepExecutionsPaginator returns a new
// ListWorkflowStepExecutionsPaginator
func NewListWorkflowStepExecutionsPaginator(client ListWorkflowStepExecutionsAPIClient, params *ListWorkflowStepExecutionsInput, optFns ...func(*ListWorkflowStepExecutionsPaginatorOptions)) *ListWorkflowStepExecutionsPaginator {
	if params == nil {
		params = &ListWorkflowStepExecutionsInput{}
	}

	options := ListWorkflowStepExecutionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListWorkflowStepExecutionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListWorkflowStepExecutionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListWorkflowStepExecutions page.
func (p *ListWorkflowStepExecutionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListWorkflowStepExecutionsOutput, error) {
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

	result, err := p.client.ListWorkflowStepExecutions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListWorkflowStepExecutions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListWorkflowStepExecutions",
	}
}
