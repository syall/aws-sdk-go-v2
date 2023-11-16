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

// Discard the current version of the registration.
func (c *Client) DiscardRegistrationVersion(ctx context.Context, params *DiscardRegistrationVersionInput, optFns ...func(*Options)) (*DiscardRegistrationVersionOutput, error) {
	if params == nil {
		params = &DiscardRegistrationVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DiscardRegistrationVersion", params, optFns, c.addOperationDiscardRegistrationVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DiscardRegistrationVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DiscardRegistrationVersionInput struct {

	// The unique identifier for the registration.
	//
	// This member is required.
	RegistrationId *string

	noSmithyDocumentSerde
}

type DiscardRegistrationVersionOutput struct {

	// The Amazon Resource Name (ARN) for the registration.
	//
	// This member is required.
	RegistrationArn *string

	// The unique identifier for the registration.
	//
	// This member is required.
	RegistrationId *string

	// The status of the registration version.
	//   - DRAFT : The initial status of a registration version after it’s created.
	//   - SUBMITTED : Your registration has been submitted.
	//   - REVIEWING : Your registration has been accepted and is being reviewed.
	//   - APPROVED : Your registration has been approved.
	//   - DISCARDED : You've abandon this version of their registration to start over
	//   with a new version.
	//   - DENIED : You must fix your registration and resubmit it.
	//   - REVOKED : Your previously approved registration has been revoked.
	//   - ARCHIVED : Your previously approved registration version moves into this
	//   status when a more recently submitted version is approved.
	//
	// This member is required.
	RegistrationVersionStatus types.RegistrationVersionStatus

	// The RegistrationVersionStatusHistory object contains the time stamps for when
	// the reservations status changes.
	//
	// This member is required.
	RegistrationVersionStatusHistory *types.RegistrationVersionStatusHistory

	// The version number of the registration.
	//
	// This member is required.
	VersionNumber *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDiscardRegistrationVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDiscardRegistrationVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDiscardRegistrationVersion{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DiscardRegistrationVersion"); err != nil {
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
	if err = addOpDiscardRegistrationVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDiscardRegistrationVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDiscardRegistrationVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DiscardRegistrationVersion",
	}
}
