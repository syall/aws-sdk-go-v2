// Code generated by smithy-go-codegen DO NOT EDIT.

package codeartifact

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a PackageVersionDescription (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_PackageVersionDescription.html)
// object that contains information about the requested package version.
func (c *Client) DescribePackageVersion(ctx context.Context, params *DescribePackageVersionInput, optFns ...func(*Options)) (*DescribePackageVersionOutput, error) {
	if params == nil {
		params = &DescribePackageVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribePackageVersion", params, optFns, c.addOperationDescribePackageVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribePackageVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePackageVersionInput struct {

	// The name of the domain that contains the repository that contains the package
	// version.
	//
	// This member is required.
	Domain *string

	// A format that specifies the type of the requested package version.
	//
	// This member is required.
	Format types.PackageFormat

	// The name of the requested package version.
	//
	// This member is required.
	Package *string

	// A string that contains the package version (for example, 3.5.2 ).
	//
	// This member is required.
	PackageVersion *string

	// The name of the repository that contains the package version.
	//
	// This member is required.
	Repository *string

	// The 12-digit account number of the Amazon Web Services account that owns the
	// domain. It does not include dashes or spaces.
	DomainOwner *string

	// The namespace of the requested package version. The package version component
	// that specifies its namespace depends on its type. For example:
	//   - The namespace of a Maven package version is its groupId .
	//   - The namespace of an npm package version is its scope .
	//   - Python and NuGet package versions do not contain a corresponding component,
	//   package versions of those formats do not have a namespace.
	//   - The namespace of a generic package is its namespace .
	Namespace *string

	noSmithyDocumentSerde
}

type DescribePackageVersionOutput struct {

	// A PackageVersionDescription (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_PackageVersionDescription.html)
	// object that contains information about the requested package version.
	//
	// This member is required.
	PackageVersion *types.PackageVersionDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribePackageVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribePackageVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribePackageVersion{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribePackageVersion"); err != nil {
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
	if err = addOpDescribePackageVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePackageVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribePackageVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribePackageVersion",
	}
}
