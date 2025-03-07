// Code generated by smithy-go-codegen DO NOT EDIT.

package bedrock

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/bedrock/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Get details for a provisioned throughput. For more information, see Provisioned
// throughput (https://docs.aws.amazon.com/bedrock/latest/userguide/what-is-service.html)
// in the Bedrock User Guide.
func (c *Client) GetProvisionedModelThroughput(ctx context.Context, params *GetProvisionedModelThroughputInput, optFns ...func(*Options)) (*GetProvisionedModelThroughputOutput, error) {
	if params == nil {
		params = &GetProvisionedModelThroughputInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetProvisionedModelThroughput", params, optFns, c.addOperationGetProvisionedModelThroughputMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetProvisionedModelThroughputOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetProvisionedModelThroughputInput struct {

	// The ARN or name of the provisioned throughput.
	//
	// This member is required.
	ProvisionedModelId *string

	noSmithyDocumentSerde
}

type GetProvisionedModelThroughputOutput struct {

	// The timestamp of the creation time for this provisioned throughput.
	//
	// This member is required.
	CreationTime *time.Time

	// The ARN of the new model to asssociate with this provisioned throughput.
	//
	// This member is required.
	DesiredModelArn *string

	// The desired number of model units that was requested to be available for this
	// provisioned throughput.
	//
	// This member is required.
	DesiredModelUnits *int32

	// ARN of the foundation model.
	//
	// This member is required.
	FoundationModelArn *string

	// The timestamp of the last modified time of this provisioned throughput.
	//
	// This member is required.
	LastModifiedTime *time.Time

	// The ARN or name of the model associated with this provisioned throughput.
	//
	// This member is required.
	ModelArn *string

	// The current number of model units requested to be available for this
	// provisioned throughput.
	//
	// This member is required.
	ModelUnits *int32

	// The ARN of the provisioned throughput.
	//
	// This member is required.
	ProvisionedModelArn *string

	// The name of the provisioned throughput.
	//
	// This member is required.
	ProvisionedModelName *string

	// Status of the provisioned throughput.
	//
	// This member is required.
	Status types.ProvisionedModelStatus

	// Commitment duration of the provisioned throughput.
	CommitmentDuration types.CommitmentDuration

	// Commitment expiration time for the provisioned throughput.
	CommitmentExpirationTime *time.Time

	// Failure message for any issues that the create operation encounters.
	FailureMessage *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetProvisionedModelThroughputMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetProvisionedModelThroughput{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetProvisionedModelThroughput{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetProvisionedModelThroughput"); err != nil {
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
	if err = addOpGetProvisionedModelThroughputValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetProvisionedModelThroughput(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetProvisionedModelThroughput(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetProvisionedModelThroughput",
	}
}
