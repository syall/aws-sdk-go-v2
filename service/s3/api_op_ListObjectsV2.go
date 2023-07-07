// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	"errors"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	s3cust "github.com/aws/aws-sdk-go-v2/service/s3/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns some or all (up to 1,000) of the objects in a bucket with each request.
// You can use the request parameters as selection criteria to return a subset of
// the objects in a bucket. A 200 OK response can contain valid or invalid XML.
// Make sure to design your application to parse the contents of the response and
// handle it appropriately. Objects are returned sorted in an ascending order of
// the respective key names in the list. For more information about listing
// objects, see Listing object keys programmatically (https://docs.aws.amazon.com/AmazonS3/latest/userguide/ListingKeysUsingAPIs.html)
// To use this operation, you must have READ access to the bucket. To use this
// action in an Identity and Access Management (IAM) policy, you must have
// permissions to perform the s3:ListBucket action. The bucket owner has this
// permission by default and can grant this permission to others. For more
// information about permissions, see Permissions Related to Bucket Subresource
// Operations (https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources)
// and Managing Access Permissions to Your Amazon S3 Resources (https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html)
// . This section describes the latest revision of this action. We recommend that
// you use this revised API for application development. For backward
// compatibility, Amazon S3 continues to support the prior version of this API,
// ListObjects (https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListObjects.html)
// . To get a list of your buckets, see ListBuckets (https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListBuckets.html)
// . The following operations are related to ListObjectsV2 :
//   - GetObject (https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html)
//   - PutObject (https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObject.html)
//   - CreateBucket (https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucket.html)
func (c *Client) ListObjectsV2(ctx context.Context, params *ListObjectsV2Input, optFns ...func(*Options)) (*ListObjectsV2Output, error) {
	if params == nil {
		params = &ListObjectsV2Input{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListObjectsV2", params, optFns, c.addOperationListObjectsV2Middlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListObjectsV2Output)
	out.ResultMetadata = metadata
	return out, nil
}

type ListObjectsV2Input struct {

	// Bucket name to list. When using this action with an access point, you must
	// direct requests to the access point hostname. The access point hostname takes
	// the form AccessPointName-AccountId.s3-accesspoint.Region.amazonaws.com. When
	// using this action with an access point through the Amazon Web Services SDKs, you
	// provide the access point ARN in place of the bucket name. For more information
	// about access point ARNs, see Using access points (https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-access-points.html)
	// in the Amazon S3 User Guide. When you use this action with Amazon S3 on
	// Outposts, you must direct requests to the S3 on Outposts hostname. The S3 on
	// Outposts hostname takes the form
	// AccessPointName-AccountId.outpostID.s3-outposts.Region.amazonaws.com . When you
	// use this action with S3 on Outposts through the Amazon Web Services SDKs, you
	// provide the Outposts access point ARN in place of the bucket name. For more
	// information about S3 on Outposts ARNs, see What is S3 on Outposts (https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html)
	// in the Amazon S3 User Guide.
	//
	// This member is required.
	Bucket *string

	// ContinuationToken indicates Amazon S3 that the list is being continued on this
	// bucket with a token. ContinuationToken is obfuscated and is not a real key.
	ContinuationToken *string

	// A delimiter is a character you use to group keys.
	Delimiter *string

	// Encoding type used by Amazon S3 to encode object keys in the response.
	EncodingType types.EncodingType

	// The account ID of the expected bucket owner. If the bucket is owned by a
	// different account, the request fails with the HTTP status code 403 Forbidden
	// (access denied).
	ExpectedBucketOwner *string

	// The owner field is not present in listV2 by default, if you want to return
	// owner field with each key in the result then set the fetch owner field to true.
	FetchOwner bool

	// Sets the maximum number of keys returned in the response. By default the action
	// returns up to 1,000 key names. The response might contain fewer keys but will
	// never contain more.
	MaxKeys int32

	// Limits the response to keys that begin with the specified prefix.
	Prefix *string

	// Confirms that the requester knows that she or he will be charged for the list
	// objects request in V2 style. Bucket owners need not specify this parameter in
	// their requests.
	RequestPayer types.RequestPayer

	// StartAfter is where you want Amazon S3 to start listing from. Amazon S3 starts
	// listing after this specified key. StartAfter can be any key in the bucket.
	StartAfter *string

	noSmithyDocumentSerde
}

type ListObjectsV2Output struct {

	// All of the keys (up to 1,000) rolled up into a common prefix count as a single
	// return when calculating the number of returns. A response can contain
	// CommonPrefixes only if you specify a delimiter. CommonPrefixes contains all (if
	// there are any) keys between Prefix and the next occurrence of the string
	// specified by a delimiter. CommonPrefixes lists keys that act like
	// subdirectories in the directory specified by Prefix . For example, if the prefix
	// is notes/ and the delimiter is a slash ( / ) as in notes/summer/july , the
	// common prefix is notes/summer/ . All of the keys that roll up into a common
	// prefix count as a single return when calculating the number of returns.
	CommonPrefixes []types.CommonPrefix

	// Metadata about each object returned.
	Contents []types.Object

	// If ContinuationToken was sent with the request, it is included in the response.
	ContinuationToken *string

	// Causes keys that contain the same string between the prefix and the first
	// occurrence of the delimiter to be rolled up into a single result element in the
	// CommonPrefixes collection. These rolled-up keys are not returned elsewhere in
	// the response. Each rolled-up result counts as only one return against the
	// MaxKeys value.
	Delimiter *string

	// Encoding type used by Amazon S3 to encode object key names in the XML response.
	// If you specify the encoding-type request parameter, Amazon S3 includes this
	// element in the response, and returns encoded key name values in the following
	// response elements: Delimiter, Prefix, Key, and StartAfter .
	EncodingType types.EncodingType

	// Set to false if all of the results were returned. Set to true if more keys are
	// available to return. If the number of results exceeds that specified by MaxKeys,
	// all of the results might not be returned.
	IsTruncated bool

	// KeyCount is the number of keys returned with this request. KeyCount will always
	// be less than or equal to the MaxKeys field. Say you ask for 50 keys, your
	// result will include 50 keys or fewer.
	KeyCount int32

	// Sets the maximum number of keys returned in the response. By default the action
	// returns up to 1,000 key names. The response might contain fewer keys but will
	// never contain more.
	MaxKeys int32

	// The bucket name. When using this action with an access point, you must direct
	// requests to the access point hostname. The access point hostname takes the form
	// AccessPointName-AccountId.s3-accesspoint.Region.amazonaws.com. When using this
	// action with an access point through the Amazon Web Services SDKs, you provide
	// the access point ARN in place of the bucket name. For more information about
	// access point ARNs, see Using access points (https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-access-points.html)
	// in the Amazon S3 User Guide. When you use this action with Amazon S3 on
	// Outposts, you must direct requests to the S3 on Outposts hostname. The S3 on
	// Outposts hostname takes the form
	// AccessPointName-AccountId.outpostID.s3-outposts.Region.amazonaws.com . When you
	// use this action with S3 on Outposts through the Amazon Web Services SDKs, you
	// provide the Outposts access point ARN in place of the bucket name. For more
	// information about S3 on Outposts ARNs, see What is S3 on Outposts (https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html)
	// in the Amazon S3 User Guide.
	Name *string

	// NextContinuationToken is sent when isTruncated is true, which means there are
	// more keys in the bucket that can be listed. The next list requests to Amazon S3
	// can be continued with this NextContinuationToken . NextContinuationToken is
	// obfuscated and is not a real key
	NextContinuationToken *string

	// Keys that begin with the indicated prefix.
	Prefix *string

	// If StartAfter was sent with the request, it is included in the response.
	StartAfter *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListObjectsV2Middlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpListObjectsV2{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListObjectsV2{}, middleware.After)
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
	if err = swapWithCustomHTTPSignerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addListObjectsV2ResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpListObjectsV2ValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListObjectsV2(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = disableAcceptEncodingGzip(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addListObjectsV2EndpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

// ListObjectsV2APIClient is a client that implements the ListObjectsV2 operation.
type ListObjectsV2APIClient interface {
	ListObjectsV2(context.Context, *ListObjectsV2Input, ...func(*Options)) (*ListObjectsV2Output, error)
}

var _ ListObjectsV2APIClient = (*Client)(nil)

// ListObjectsV2PaginatorOptions is the paginator options for ListObjectsV2
type ListObjectsV2PaginatorOptions struct {
	// Sets the maximum number of keys returned in the response. By default the action
	// returns up to 1,000 key names. The response might contain fewer keys but will
	// never contain more.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListObjectsV2Paginator is a paginator for ListObjectsV2
type ListObjectsV2Paginator struct {
	options   ListObjectsV2PaginatorOptions
	client    ListObjectsV2APIClient
	params    *ListObjectsV2Input
	nextToken *string
	firstPage bool
}

// NewListObjectsV2Paginator returns a new ListObjectsV2Paginator
func NewListObjectsV2Paginator(client ListObjectsV2APIClient, params *ListObjectsV2Input, optFns ...func(*ListObjectsV2PaginatorOptions)) *ListObjectsV2Paginator {
	if params == nil {
		params = &ListObjectsV2Input{}
	}

	options := ListObjectsV2PaginatorOptions{}
	if params.MaxKeys != 0 {
		options.Limit = params.MaxKeys
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListObjectsV2Paginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.ContinuationToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListObjectsV2Paginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListObjectsV2 page.
func (p *ListObjectsV2Paginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListObjectsV2Output, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.ContinuationToken = p.nextToken

	params.MaxKeys = p.options.Limit

	result, err := p.client.ListObjectsV2(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = nil
	if result.IsTruncated {
		p.nextToken = result.NextContinuationToken
	}

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListObjectsV2(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "ListObjectsV2",
	}
}

// getListObjectsV2BucketMember returns a pointer to string denoting a provided
// bucket member valueand a boolean indicating if the input has a modeled bucket
// name,
func getListObjectsV2BucketMember(input interface{}) (*string, bool) {
	in := input.(*ListObjectsV2Input)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addListObjectsV2UpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{
			GetBucketFromInput: getListObjectsV2BucketMember,
		},
		UsePathStyle:                   options.UsePathStyle,
		UseAccelerate:                  options.UseAccelerate,
		SupportsAccelerate:             true,
		TargetS3ObjectLambda:           false,
		EndpointResolver:               options.EndpointResolver,
		EndpointResolverOptions:        options.EndpointOptions,
		UseARNRegion:                   options.UseARNRegion,
		DisableMultiRegionAccessPoints: options.DisableMultiRegionAccessPoints,
	})
}

type opListObjectsV2EndpointDisableHTTPSMiddleware struct {
	EndpointDisableHTTPS bool
}

func (*opListObjectsV2EndpointDisableHTTPSMiddleware) ID() string {
	return "opListObjectsV2EndpointDisableHTTPSMiddleware"
}

func (m *opListObjectsV2EndpointDisableHTTPSMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointDisableHTTPS {
		req.URL.Scheme = "http"
	}

	return next.HandleSerialize(ctx, in)

}
func addListObjectsV2EndpointDisableHTTPSMiddleware(stack *middleware.Stack, o Options) error {
	return stack.Serialize.Insert(&opListObjectsV2EndpointDisableHTTPSMiddleware{
		EndpointDisableHTTPS: o.EndpointOptions.DisableHTTPS,
	}, "opListObjectsV2ResolveEndpointMiddleware", middleware.After)
}

type opListObjectsV2ResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  BuiltInParameterResolver
}

func (*opListObjectsV2ResolveEndpointMiddleware) ID() string {
	return "opListObjectsV2ResolveEndpointMiddleware"
}

func (m *opListObjectsV2ResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*ListObjectsV2Input)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	params.Bucket = input.Bucket

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "s3"
			signingRegion := m.BuiltInResolver.(*BuiltInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			ctx = s3cust.SetSignerVersion(ctx, internalauth.SigV4)
		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "s3"
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*BuiltInResolver).Region
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			ctx = s3cust.SetSignerVersion(ctx, v4Scheme.Name)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				signingNameDefault := "s3"
				v4aScheme.SigningName = &signingNameDefault
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			ctx = s3cust.SetSignerVersion(ctx, v4aScheme.Name)
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addListObjectsV2ResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opListObjectsV2ResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &BuiltInResolver{
			Region:                         options.Region,
			UseFIPS:                        options.EndpointOptions.UseFIPSEndpoint,
			UseDualStack:                   options.EndpointOptions.UseDualStackEndpoint,
			Endpoint:                       options.BaseEndpoint,
			ForcePathStyle:                 options.UsePathStyle,
			Accelerate:                     options.UseAccelerate,
			DisableMultiRegionAccessPoints: options.DisableMultiRegionAccessPoints,
			UseArnRegion:                   options.UseARNRegion,
		},
	}, "ResolveEndpoint", middleware.After)
}
