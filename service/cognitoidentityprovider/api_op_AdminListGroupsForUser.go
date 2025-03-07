// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the groups that the user belongs to. Amazon Cognito evaluates Identity
// and Access Management (IAM) policies in requests for this API operation. For
// this operation, you must use IAM credentials to authorize requests, and you must
// grant yourself the corresponding IAM permission in a policy. Learn more
//   - Signing Amazon Web Services API Requests (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html)
//   - Using the Amazon Cognito user pools API and user pool endpoints (https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html)
func (c *Client) AdminListGroupsForUser(ctx context.Context, params *AdminListGroupsForUserInput, optFns ...func(*Options)) (*AdminListGroupsForUserOutput, error) {
	if params == nil {
		params = &AdminListGroupsForUserInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AdminListGroupsForUser", params, optFns, c.addOperationAdminListGroupsForUserMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AdminListGroupsForUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AdminListGroupsForUserInput struct {

	// The user pool ID for the user pool.
	//
	// This member is required.
	UserPoolId *string

	// The username for the user.
	//
	// This member is required.
	Username *string

	// The limit of the request to list groups.
	Limit *int32

	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	NextToken *string

	noSmithyDocumentSerde
}

type AdminListGroupsForUserOutput struct {

	// The groups that the user belongs to.
	Groups []types.GroupType

	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAdminListGroupsForUserMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAdminListGroupsForUser{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAdminListGroupsForUser{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AdminListGroupsForUser"); err != nil {
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
	if err = addOpAdminListGroupsForUserValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAdminListGroupsForUser(options.Region), middleware.Before); err != nil {
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

// AdminListGroupsForUserAPIClient is a client that implements the
// AdminListGroupsForUser operation.
type AdminListGroupsForUserAPIClient interface {
	AdminListGroupsForUser(context.Context, *AdminListGroupsForUserInput, ...func(*Options)) (*AdminListGroupsForUserOutput, error)
}

var _ AdminListGroupsForUserAPIClient = (*Client)(nil)

// AdminListGroupsForUserPaginatorOptions is the paginator options for
// AdminListGroupsForUser
type AdminListGroupsForUserPaginatorOptions struct {
	// The limit of the request to list groups.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// AdminListGroupsForUserPaginator is a paginator for AdminListGroupsForUser
type AdminListGroupsForUserPaginator struct {
	options   AdminListGroupsForUserPaginatorOptions
	client    AdminListGroupsForUserAPIClient
	params    *AdminListGroupsForUserInput
	nextToken *string
	firstPage bool
}

// NewAdminListGroupsForUserPaginator returns a new AdminListGroupsForUserPaginator
func NewAdminListGroupsForUserPaginator(client AdminListGroupsForUserAPIClient, params *AdminListGroupsForUserInput, optFns ...func(*AdminListGroupsForUserPaginatorOptions)) *AdminListGroupsForUserPaginator {
	if params == nil {
		params = &AdminListGroupsForUserInput{}
	}

	options := AdminListGroupsForUserPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &AdminListGroupsForUserPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *AdminListGroupsForUserPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next AdminListGroupsForUser page.
func (p *AdminListGroupsForUserPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*AdminListGroupsForUserOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.AdminListGroupsForUser(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opAdminListGroupsForUser(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AdminListGroupsForUser",
	}
}
