// Code generated by smithy-go-codegen DO NOT EDIT.

package costexplorer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the name, Amazon Resource Name (ARN), NumberOfRules and effective dates
// of all Cost Categories defined in the account. You have the option to use
// EffectiveOn to return a list of Cost Categories that were active on a specific
// date. If there is no EffectiveOn specified, you’ll see Cost Categories that are
// effective on the current date. If Cost Category is still effective, EffectiveEnd
// is omitted in the response. ListCostCategoryDefinitions supports pagination. The
// request can have a MaxResults range up to 100.
func (c *Client) ListCostCategoryDefinitions(ctx context.Context, params *ListCostCategoryDefinitionsInput, optFns ...func(*Options)) (*ListCostCategoryDefinitionsOutput, error) {
	if params == nil {
		params = &ListCostCategoryDefinitionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCostCategoryDefinitions", params, optFns, c.addOperationListCostCategoryDefinitionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCostCategoryDefinitionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCostCategoryDefinitionsInput struct {

	// The date when the Cost Category was effective.
	EffectiveOn *string

	// The number of entries a paginated response contains.
	MaxResults *int32

	// The token to retrieve the next set of results. Amazon Web Services provides the
	// token when the response from a previous call has more results than the maximum
	// page size.
	NextToken *string

	noSmithyDocumentSerde
}

type ListCostCategoryDefinitionsOutput struct {

	// A reference to a Cost Category that contains enough information to identify the
	// Cost Category.
	CostCategoryReferences []types.CostCategoryReference

	// The token to retrieve the next set of results. Amazon Web Services provides the
	// token when the response from a previous call has more results than the maximum
	// page size.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCostCategoryDefinitionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListCostCategoryDefinitions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListCostCategoryDefinitions{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCostCategoryDefinitions(options.Region), middleware.Before); err != nil {
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

// ListCostCategoryDefinitionsAPIClient is a client that implements the
// ListCostCategoryDefinitions operation.
type ListCostCategoryDefinitionsAPIClient interface {
	ListCostCategoryDefinitions(context.Context, *ListCostCategoryDefinitionsInput, ...func(*Options)) (*ListCostCategoryDefinitionsOutput, error)
}

var _ ListCostCategoryDefinitionsAPIClient = (*Client)(nil)

// ListCostCategoryDefinitionsPaginatorOptions is the paginator options for
// ListCostCategoryDefinitions
type ListCostCategoryDefinitionsPaginatorOptions struct {
	// The number of entries a paginated response contains.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCostCategoryDefinitionsPaginator is a paginator for
// ListCostCategoryDefinitions
type ListCostCategoryDefinitionsPaginator struct {
	options   ListCostCategoryDefinitionsPaginatorOptions
	client    ListCostCategoryDefinitionsAPIClient
	params    *ListCostCategoryDefinitionsInput
	nextToken *string
	firstPage bool
}

// NewListCostCategoryDefinitionsPaginator returns a new
// ListCostCategoryDefinitionsPaginator
func NewListCostCategoryDefinitionsPaginator(client ListCostCategoryDefinitionsAPIClient, params *ListCostCategoryDefinitionsInput, optFns ...func(*ListCostCategoryDefinitionsPaginatorOptions)) *ListCostCategoryDefinitionsPaginator {
	if params == nil {
		params = &ListCostCategoryDefinitionsInput{}
	}

	options := ListCostCategoryDefinitionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCostCategoryDefinitionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCostCategoryDefinitionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCostCategoryDefinitions page.
func (p *ListCostCategoryDefinitionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCostCategoryDefinitionsOutput, error) {
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

	result, err := p.client.ListCostCategoryDefinitions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListCostCategoryDefinitions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ce",
		OperationName: "ListCostCategoryDefinitions",
	}
}
