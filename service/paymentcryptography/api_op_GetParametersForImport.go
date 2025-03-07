// Code generated by smithy-go-codegen DO NOT EDIT.

package paymentcryptography

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/paymentcryptography/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets the import token and the wrapping key certificate to initiate a TR-34 key
// import into Amazon Web Services Payment Cryptography. The wrapping key
// certificate wraps the key under import within the TR-34 key payload. The import
// token and wrapping key certificate must be in place and operational before
// calling ImportKey . The import token expires in 7 days. The same import token
// can be used to import multiple keys into your service account. Cross-account
// use: This operation can't be used across different Amazon Web Services accounts.
// Related operations:
//   - GetParametersForExport
//   - ImportKey
func (c *Client) GetParametersForImport(ctx context.Context, params *GetParametersForImportInput, optFns ...func(*Options)) (*GetParametersForImportOutput, error) {
	if params == nil {
		params = &GetParametersForImportInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetParametersForImport", params, optFns, c.addOperationGetParametersForImportMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetParametersForImportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetParametersForImportInput struct {

	// The key block format type such as TR-34 or TR-31 to use during key material
	// import. Import token is only required for TR-34 key import TR34_KEY_BLOCK .
	// Import token is not required for TR-31 key import.
	//
	// This member is required.
	KeyMaterialType types.KeyMaterialType

	// The wrapping key algorithm to generate a wrapping key certificate. This
	// certificate wraps the key under import within the TR-34 key block cryptogram.
	// RSA_2048 is the only wrapping key algorithm allowed.
	//
	// This member is required.
	WrappingKeyAlgorithm types.KeyAlgorithm

	noSmithyDocumentSerde
}

type GetParametersForImportOutput struct {

	// The import token to initiate key import into Amazon Web Services Payment
	// Cryptography. The import token expires after 7 days. You can use the same import
	// token to import multiple keys to the same service account.
	//
	// This member is required.
	ImportToken *string

	// The validity period of the import token.
	//
	// This member is required.
	ParametersValidUntilTimestamp *time.Time

	// The algorithm of the wrapping key for use within TR-34 key block. RSA_2048 is
	// the only wrapping key algorithm allowed.
	//
	// This member is required.
	WrappingKeyAlgorithm types.KeyAlgorithm

	// The wrapping key certificate of the wrapping key for use within the TR-34 key
	// block. The certificate expires in 7 days.
	//
	// This member is required.
	WrappingKeyCertificate *string

	// The Amazon Web Services Payment Cryptography certificate chain that signed the
	// wrapping key certificate. This is the root certificate authority (CA) within
	// your service account.
	//
	// This member is required.
	WrappingKeyCertificateChain *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetParametersForImportMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetParametersForImport{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetParametersForImport{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetParametersForImport"); err != nil {
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
	if err = addOpGetParametersForImportValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetParametersForImport(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetParametersForImport(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetParametersForImport",
	}
}
