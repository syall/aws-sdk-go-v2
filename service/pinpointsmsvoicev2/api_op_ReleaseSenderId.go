// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoicev2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoicev2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Releases an existing sender ID in your account.
func (c *Client) ReleaseSenderId(ctx context.Context, params *ReleaseSenderIdInput, optFns ...func(*Options)) (*ReleaseSenderIdOutput, error) {
	if params == nil {
		params = &ReleaseSenderIdInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ReleaseSenderId", params, optFns, c.addOperationReleaseSenderIdMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ReleaseSenderIdOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ReleaseSenderIdInput struct {

	// The two-character code, in ISO 3166-1 alpha-2 format, for the country or region.
	//
	// This member is required.
	IsoCountryCode *string

	// The sender ID to release.
	//
	// This member is required.
	SenderId *string

	noSmithyDocumentSerde
}

type ReleaseSenderIdOutput struct {

	// The two-character code, in ISO 3166-1 alpha-2 format, for the country or region.
	//
	// This member is required.
	IsoCountryCode *string

	// The type of message. Valid values are TRANSACTIONAL for messages that are
	// critical or time-sensitive and PROMOTIONAL for messages that aren't critical or
	// time-sensitive.
	//
	// This member is required.
	MessageTypes []types.MessageType

	// The monthly price, in US dollars, to lease the sender ID.
	//
	// This member is required.
	MonthlyLeasingPrice *string

	// True if the sender ID is registered.
	//
	// This member is required.
	Registered bool

	// The sender ID that was released.
	//
	// This member is required.
	SenderId *string

	// The Amazon Resource Name (ARN) associated with the SenderId.
	//
	// This member is required.
	SenderIdArn *string

	// The unique identifier for the registration.
	RegistrationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationReleaseSenderIdMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpReleaseSenderId{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpReleaseSenderId{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ReleaseSenderId"); err != nil {
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
	if err = addOpReleaseSenderIdValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opReleaseSenderId(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opReleaseSenderId(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ReleaseSenderId",
	}
}
