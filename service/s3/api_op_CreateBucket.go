// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	s3cust "github.com/aws/aws-sdk-go-v2/service/s3/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new S3 bucket. To create a bucket, you must register with Amazon S3
// and have a valid Amazon Web Services Access Key ID to authenticate requests.
// Anonymous requests are never allowed to create buckets. By creating the bucket,
// you become the bucket owner. Not every string is an acceptable bucket name. For
// information about bucket naming restrictions, see Bucket naming rules (https://docs.aws.amazon.com/AmazonS3/latest/userguide/bucketnamingrules.html)
// . If you want to create an Amazon S3 on Outposts bucket, see Create Bucket (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateBucket.html)
// . By default, the bucket is created in the US East (N. Virginia) Region. You can
// optionally specify a Region in the request body. You might choose a Region to
// optimize latency, minimize costs, or address regulatory requirements. For
// example, if you reside in Europe, you will probably find it advantageous to
// create buckets in the Europe (Ireland) Region. For more information, see
// Accessing a bucket (https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)
// . If you send your create bucket request to the s3.amazonaws.com endpoint, the
// request goes to the us-east-1 Region. Accordingly, the signature calculations in
// Signature Version 4 must use us-east-1 as the Region, even if the location
// constraint in the request specifies another Region where the bucket is to be
// created. If you create a bucket in a Region other than US East (N. Virginia),
// your application must be able to handle 307 redirect. For more information, see
// Virtual hosting of buckets (https://docs.aws.amazon.com/AmazonS3/latest/dev/VirtualHosting.html)
// . Access control lists (ACLs) When creating a bucket using this operation, you
// can optionally configure the bucket ACL to specify the accounts or groups that
// should be granted specific permissions on the bucket. If your CreateBucket
// request sets bucket owner enforced for S3 Object Ownership and specifies a
// bucket ACL that provides access to an external Amazon Web Services account, your
// request fails with a 400 error and returns the
// InvalidBucketAclWithObjectOwnership error code. For more information, see
// Controlling object ownership (https://docs.aws.amazon.com/AmazonS3/latest/userguide/about-object-ownership.html)
// in the Amazon S3 User Guide. There are two ways to grant the appropriate
// permissions using the request headers.
//   - Specify a canned ACL using the x-amz-acl request header. Amazon S3 supports
//     a set of predefined ACLs, known as canned ACLs. Each canned ACL has a predefined
//     set of grantees and permissions. For more information, see Canned ACL (https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#CannedACL)
//     .
//   - Specify access permissions explicitly using the x-amz-grant-read ,
//     x-amz-grant-write , x-amz-grant-read-acp , x-amz-grant-write-acp , and
//     x-amz-grant-full-control headers. These headers map to the set of permissions
//     Amazon S3 supports in an ACL. For more information, see Access control list
//     (ACL) overview (https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl-overview.html)
//     . You specify each grantee as a type=value pair, where the type is one of the
//     following:
//   - id – if the value specified is the canonical user ID of an Amazon Web
//     Services account
//   - uri – if you are granting permissions to a predefined group
//   - emailAddress – if the value specified is the email address of an Amazon Web
//     Services account Using email addresses to specify a grantee is only supported in
//     the following Amazon Web Services Regions:
//   - US East (N. Virginia)
//   - US West (N. California)
//   - US West (Oregon)
//   - Asia Pacific (Singapore)
//   - Asia Pacific (Sydney)
//   - Asia Pacific (Tokyo)
//   - Europe (Ireland)
//   - South America (São Paulo) For a list of all the Amazon S3 supported Regions
//     and endpoints, see Regions and Endpoints (https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region)
//     in the Amazon Web Services General Reference. For example, the following
//     x-amz-grant-read header grants the Amazon Web Services accounts identified by
//     account IDs permissions to read object data and its metadata:
//     x-amz-grant-read: id="11112222333", id="444455556666"
//
// You can use either a canned ACL or specify access permissions explicitly. You
// cannot do both. Permissions In addition to s3:CreateBucket , the following
// permissions are required when your CreateBucket includes specific headers:
//   - ACLs - If your CreateBucket request specifies ACL permissions and the ACL is
//     public-read, public-read-write, authenticated-read, or if you specify access
//     permissions explicitly through any other ACL, both s3:CreateBucket and
//     s3:PutBucketAcl permissions are needed. If the ACL the CreateBucket request is
//     private or doesn't specify any ACLs, only s3:CreateBucket permission is
//     needed.
//   - Object Lock - If ObjectLockEnabledForBucket is set to true in your
//     CreateBucket request, s3:PutBucketObjectLockConfiguration and
//     s3:PutBucketVersioning permissions are required.
//   - S3 Object Ownership - If your CreateBucket request includes the
//     x-amz-object-ownership header, s3:PutBucketOwnershipControls permission is
//     required.
//
// The following operations are related to CreateBucket :
//   - PutObject (https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObject.html)
//   - DeleteBucket (https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucket.html)
func (c *Client) CreateBucket(ctx context.Context, params *CreateBucketInput, optFns ...func(*Options)) (*CreateBucketOutput, error) {
	if params == nil {
		params = &CreateBucketInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateBucket", params, optFns, c.addOperationCreateBucketMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateBucketOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateBucketInput struct {

	// The name of the bucket to create.
	//
	// This member is required.
	Bucket *string

	// The canned ACL to apply to the bucket.
	ACL types.BucketCannedACL

	// The configuration information for the bucket.
	CreateBucketConfiguration *types.CreateBucketConfiguration

	// Allows grantee the read, write, read ACP, and write ACP permissions on the
	// bucket.
	GrantFullControl *string

	// Allows grantee to list the objects in the bucket.
	GrantRead *string

	// Allows grantee to read the bucket ACL.
	GrantReadACP *string

	// Allows grantee to create new objects in the bucket. For the bucket and object
	// owners of existing objects, also allows deletions and overwrites of those
	// objects.
	GrantWrite *string

	// Allows grantee to write the ACL for the applicable bucket.
	GrantWriteACP *string

	// Specifies whether you want S3 Object Lock to be enabled for the new bucket.
	ObjectLockEnabledForBucket bool

	// The container element for object ownership for a bucket's ownership controls.
	// BucketOwnerPreferred - Objects uploaded to the bucket change ownership to the
	// bucket owner if the objects are uploaded with the bucket-owner-full-control
	// canned ACL. ObjectWriter - The uploading account will own the object if the
	// object is uploaded with the bucket-owner-full-control canned ACL.
	// BucketOwnerEnforced - Access control lists (ACLs) are disabled and no longer
	// affect permissions. The bucket owner automatically owns and has full control
	// over every object in the bucket. The bucket only accepts PUT requests that don't
	// specify an ACL or bucket owner full control ACLs, such as the
	// bucket-owner-full-control canned ACL or an equivalent form of this ACL expressed
	// in the XML format.
	ObjectOwnership types.ObjectOwnership

	noSmithyDocumentSerde
}

type CreateBucketOutput struct {

	// A forward slash followed by the name of the bucket.
	Location *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateBucketMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpCreateBucket{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpCreateBucket{}, middleware.After)
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
	if err = addCreateBucketResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateBucketValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateBucket(options.Region), middleware.Before); err != nil {
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
	if err = addCreateBucketEndpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateBucket(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "CreateBucket",
	}
}

// getCreateBucketBucketMember returns a pointer to string denoting a provided
// bucket member valueand a boolean indicating if the input has a modeled bucket
// name,
func getCreateBucketBucketMember(input interface{}) (*string, bool) {
	in := input.(*CreateBucketInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addCreateBucketUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{
			GetBucketFromInput: getCreateBucketBucketMember,
		},
		UsePathStyle:                   options.UsePathStyle,
		UseAccelerate:                  options.UseAccelerate,
		SupportsAccelerate:             false,
		TargetS3ObjectLambda:           false,
		EndpointResolver:               options.EndpointResolver,
		EndpointResolverOptions:        options.EndpointOptions,
		UseARNRegion:                   options.UseARNRegion,
		DisableMultiRegionAccessPoints: options.DisableMultiRegionAccessPoints,
	})
}

type opCreateBucketEndpointDisableHTTPSMiddleware struct {
	EndpointDisableHTTPS bool
}

func (*opCreateBucketEndpointDisableHTTPSMiddleware) ID() string {
	return "opCreateBucketEndpointDisableHTTPSMiddleware"
}

func (m *opCreateBucketEndpointDisableHTTPSMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
func addCreateBucketEndpointDisableHTTPSMiddleware(stack *middleware.Stack, o Options) error {
	return stack.Serialize.Insert(&opCreateBucketEndpointDisableHTTPSMiddleware{
		EndpointDisableHTTPS: o.EndpointOptions.DisableHTTPS,
	}, "opCreateBucketResolveEndpointMiddleware", middleware.After)
}

type opCreateBucketResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  BuiltInParameterResolver
}

func (*opCreateBucketResolveEndpointMiddleware) ID() string {
	return "opCreateBucketResolveEndpointMiddleware"
}

func (m *opCreateBucketResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*CreateBucketInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	params.Bucket = input.Bucket

	params.DisableAccessPoints = ptr.Bool(true)

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
				v4aScheme.SigningName = aws.String("s3")
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

func addCreateBucketResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opCreateBucketResolveEndpointMiddleware{
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
