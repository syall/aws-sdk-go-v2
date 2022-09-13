// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This API is in preview release for Amazon Connect and is subject to change.
// Returns a paginated list of all the Amazon Lex V1 bots currently associated with
// the instance. To return both Amazon Lex V1 and V2 bots, use the ListBots
// (https://docs.aws.amazon.com/connect/latest/APIReference/API_ListBots.html) API.
func (c *Client) ListLexBots(ctx context.Context, params *ListLexBotsInput, optFns ...func(*Options)) (*ListLexBotsOutput, error) {
	if params == nil {
		params = &ListLexBotsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListLexBots", params, optFns, c.addOperationListLexBotsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListLexBotsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListLexBotsInput struct {

	// The identifier of the Amazon Connect instance. You can find the instanceId in
	// the ARN of the instance.
	//
	// This member is required.
	InstanceId *string

	// The maximum number of results to return per page. If no value is specified, the
	// default is 10.
	MaxResults *int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListLexBotsOutput struct {

	// The names and Regions of the Amazon Lex bots associated with the specified
	// instance.
	LexBots []types.LexBot

	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListLexBotsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListLexBots{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListLexBots{}, middleware.After)
	if err != nil {
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
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListLexBotsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListLexBots(options.Region), middleware.Before); err != nil {
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
	return nil
}

// ListLexBotsAPIClient is a client that implements the ListLexBots operation.
type ListLexBotsAPIClient interface {
	ListLexBots(context.Context, *ListLexBotsInput, ...func(*Options)) (*ListLexBotsOutput, error)
}

var _ ListLexBotsAPIClient = (*Client)(nil)

// ListLexBotsPaginatorOptions is the paginator options for ListLexBots
type ListLexBotsPaginatorOptions struct {
	// The maximum number of results to return per page. If no value is specified, the
	// default is 10.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListLexBotsPaginator is a paginator for ListLexBots
type ListLexBotsPaginator struct {
	options   ListLexBotsPaginatorOptions
	client    ListLexBotsAPIClient
	params    *ListLexBotsInput
	nextToken *string
	firstPage bool
}

// NewListLexBotsPaginator returns a new ListLexBotsPaginator
func NewListLexBotsPaginator(client ListLexBotsAPIClient, params *ListLexBotsInput, optFns ...func(*ListLexBotsPaginatorOptions)) *ListLexBotsPaginator {
	if params == nil {
		params = &ListLexBotsInput{}
	}

	options := ListLexBotsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListLexBotsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListLexBotsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListLexBots page.
func (p *ListLexBotsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListLexBotsOutput, error) {
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

	result, err := p.client.ListLexBots(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListLexBots(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "ListLexBots",
	}
}
