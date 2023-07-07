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
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves OwnershipControls for an Amazon S3 bucket. To use this operation, you
// must have the s3:GetBucketOwnershipControls permission. For more information
// about Amazon S3 permissions, see Specifying permissions in a policy (https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html)
// . For information about Amazon S3 Object Ownership, see Using Object Ownership (https://docs.aws.amazon.com/AmazonS3/latest/userguide/about-object-ownership.html)
// . The following operations are related to GetBucketOwnershipControls :
//   - PutBucketOwnershipControls
//   - DeleteBucketOwnershipControls
func (c *Client) GetBucketOwnershipControls(ctx context.Context, params *GetBucketOwnershipControlsInput, optFns ...func(*Options)) (*GetBucketOwnershipControlsOutput, error) {
	if params == nil {
		params = &GetBucketOwnershipControlsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBucketOwnershipControls", params, optFns, c.addOperationGetBucketOwnershipControlsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBucketOwnershipControlsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBucketOwnershipControlsInput struct {

	// The name of the Amazon S3 bucket whose OwnershipControls you want to retrieve.
	//
	// This member is required.
	Bucket *string

	// The account ID of the expected bucket owner. If the bucket is owned by a
	// different account, the request fails with the HTTP status code 403 Forbidden
	// (access denied).
	ExpectedBucketOwner *string

	noSmithyDocumentSerde
}

type GetBucketOwnershipControlsOutput struct {

	// The OwnershipControls (BucketOwnerEnforced, BucketOwnerPreferred, or
	// ObjectWriter) currently in effect for this Amazon S3 bucket.
	OwnershipControls *types.OwnershipControls

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetBucketOwnershipControlsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetBucketOwnershipControls{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetBucketOwnershipControls{}, middleware.After)
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
	if err = addGetBucketOwnershipControlsResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpGetBucketOwnershipControlsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetBucketOwnershipControls(options.Region), middleware.Before); err != nil {
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
	if err = addGetBucketOwnershipControlsEndpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetBucketOwnershipControls(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "GetBucketOwnershipControls",
	}
}

// getGetBucketOwnershipControlsBucketMember returns a pointer to string denoting
// a provided bucket member valueand a boolean indicating if the input has a
// modeled bucket name,
func getGetBucketOwnershipControlsBucketMember(input interface{}) (*string, bool) {
	in := input.(*GetBucketOwnershipControlsInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addGetBucketOwnershipControlsUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{
			GetBucketFromInput: getGetBucketOwnershipControlsBucketMember,
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

type opGetBucketOwnershipControlsEndpointDisableHTTPSMiddleware struct {
	EndpointDisableHTTPS bool
}

func (*opGetBucketOwnershipControlsEndpointDisableHTTPSMiddleware) ID() string {
	return "opGetBucketOwnershipControlsEndpointDisableHTTPSMiddleware"
}

func (m *opGetBucketOwnershipControlsEndpointDisableHTTPSMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
func addGetBucketOwnershipControlsEndpointDisableHTTPSMiddleware(stack *middleware.Stack, o Options) error {
	return stack.Serialize.Insert(&opGetBucketOwnershipControlsEndpointDisableHTTPSMiddleware{
		EndpointDisableHTTPS: o.EndpointOptions.DisableHTTPS,
	}, "opGetBucketOwnershipControlsResolveEndpointMiddleware", middleware.After)
}

type opGetBucketOwnershipControlsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  BuiltInParameterResolver
}

func (*opGetBucketOwnershipControlsResolveEndpointMiddleware) ID() string {
	return "opGetBucketOwnershipControlsResolveEndpointMiddleware"
}

func (m *opGetBucketOwnershipControlsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*GetBucketOwnershipControlsInput)
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

func addGetBucketOwnershipControlsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opGetBucketOwnershipControlsResolveEndpointMiddleware{
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
