// Code generated by smithy-go-codegen DO NOT EDIT.

package securitylake

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/securitylake/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds a natively supported Amazon Web Service as an Amazon Security Lake source.
// Enables source types for member accounts in required Amazon Web Services
// Regions, based on the parameters you specify. You can choose any source type in
// any Region for either accounts that are part of a trusted organization or
// standalone accounts. Once you add an Amazon Web Service as a source, Security
// Lake starts collecting logs and events from it, You can use this API only to
// enable natively supported Amazon Web Services as a source. Use
// CreateCustomLogSource to enable data collection from a custom source.
func (c *Client) CreateAwsLogSource(ctx context.Context, params *CreateAwsLogSourceInput, optFns ...func(*Options)) (*CreateAwsLogSourceOutput, error) {
	if params == nil {
		params = &CreateAwsLogSourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAwsLogSource", params, optFns, c.addOperationCreateAwsLogSourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAwsLogSourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAwsLogSourceInput struct {

	// Specify the natively-supported Amazon Web Services service to add as a source
	// in Security Lake.
	//
	// This member is required.
	Sources []types.AwsLogSourceConfiguration

	noSmithyDocumentSerde
}

type CreateAwsLogSourceOutput struct {

	// Lists all accounts in which enabling a natively supported Amazon Web Service as
	// a Security Lake source failed. The failure occurred as these accounts are not
	// part of an organization.
	Failed []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAwsLogSourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateAwsLogSource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateAwsLogSource{}, middleware.After)
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
	if err = addOpCreateAwsLogSourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAwsLogSource(options.Region), middleware.Before); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opCreateAwsLogSource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securitylake",
		OperationName: "CreateAwsLogSource",
	}
}
