// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the recommender to modify the recommender configuration. If you update
// the recommender to modify the columns used in training, Amazon Personalize
// automatically starts a full retraining of the models backing your recommender.
// While the update completes, you can still get recommendations from the
// recommender. The recommender uses the previous configuration until the update
// completes. To track the status of this update, use the latestRecommenderUpdate
// returned in the DescribeRecommender (https://docs.aws.amazon.com/personalize/latest/dg/API_DescribeRecommender.html)
// operation.
func (c *Client) UpdateRecommender(ctx context.Context, params *UpdateRecommenderInput, optFns ...func(*Options)) (*UpdateRecommenderOutput, error) {
	if params == nil {
		params = &UpdateRecommenderInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateRecommender", params, optFns, c.addOperationUpdateRecommenderMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateRecommenderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateRecommenderInput struct {

	// The Amazon Resource Name (ARN) of the recommender to modify.
	//
	// This member is required.
	RecommenderArn *string

	// The configuration details of the recommender.
	//
	// This member is required.
	RecommenderConfig *types.RecommenderConfig

	noSmithyDocumentSerde
}

type UpdateRecommenderOutput struct {

	// The same recommender Amazon Resource Name (ARN) as given in the request.
	RecommenderArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateRecommenderMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateRecommender{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateRecommender{}, middleware.After)
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
	if err = addOpUpdateRecommenderValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRecommender(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateRecommender(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "UpdateRecommender",
	}
}
