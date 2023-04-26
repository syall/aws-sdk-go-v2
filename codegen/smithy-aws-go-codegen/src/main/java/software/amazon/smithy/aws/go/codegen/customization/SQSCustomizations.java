/*
 * Copyright 2023 Amazon.com, Inc. or its affiliates. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License").
 * You may not use this file except in compliance with the License.
 * A copy of the License is located at
 *
 *  http://aws.amazon.com/apache2.0
 *
 * or in the "license" file accompanying this file. This file is distributed
 * on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
 * express or implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

package software.amazon.smithy.aws.go.codegen.customization;

import java.util.ArrayList;
import java.util.List;
import java.util.Map;
import java.util.logging.Logger;
import software.amazon.smithy.go.codegen.GoSettings;
import software.amazon.smithy.go.codegen.integration.GoIntegration;
import software.amazon.smithy.model.Model;
import software.amazon.smithy.model.node.Node;
import software.amazon.smithy.model.node.NumberNode;
import software.amazon.smithy.model.shapes.MemberShape;
import software.amazon.smithy.model.shapes.Shape;
import software.amazon.smithy.model.shapes.ShapeId;
import software.amazon.smithy.model.traits.DefaultTrait;
import software.amazon.smithy.model.transform.ModelTransformer;
import software.amazon.smithy.utils.MapUtils;

/**
 * AWS SDK for Go V2 Integrations for SQS services
 */
public class SQSCustomizations implements GoIntegration {
    private static final Logger LOGGER = Logger.getLogger(SQSCustomizations.class.getName());
    private static final ShapeId SQS_SERVICE_ID = ShapeId.from("com.amazonaws.sqs#AmazonSQS");

    /**
     * Default traits that need to be backfilled
     */
    private static final Map<ShapeId, Node> DEFAULT_TRAIT_BACKFILL = MapUtils.of(
            ShapeId.from("com.amazonaws.sqs#SendMessageRequest$DelaySeconds"), NumberNode.from(0),
            ShapeId.from("com.amazonaws.sqs#ChangeMessageVisibilityBatchRequestEntry$VisibilityTimeout"), NumberNode.from(0),
            ShapeId.from("com.amazonaws.sqs#SendMessageBatchRequestEntry$DelaySeconds"), NumberNode.from(0),
            ShapeId.from("com.amazonaws.sqs#ChangeMessageVisibilityRequest$VisibilityTimeout"), NumberNode.from(0),
            ShapeId.from("com.amazonaws.sqs#ReceiveMessageRequest$WaitTimeSeconds"), NumberNode.from(0),
            ShapeId.from("com.amazonaws.sqs#ReceiveMessageRequest$VisibilityTimeout"), NumberNode.from(0),
            ShapeId.from("com.amazonaws.sqs#ReceiveMessageRequest$MaxNumberOfMessages"), NumberNode.from(0));

    @Override
    public Model preprocessModel(Model model, GoSettings settings) {
        ShapeId serviceId = settings.getService();
        if (!serviceId.equals(SQS_SERVICE_ID)) {
            return model;
        }

        List<Shape> updates = new ArrayList<>();

        // Patch default traits to members to avoid breaking changes
        for (Map.Entry<ShapeId, Node> entry : DEFAULT_TRAIT_BACKFILL.entrySet()) {
            ShapeId memberShapeId = entry.getKey();
            MemberShape memberShape = model.expectShape(memberShapeId, MemberShape.class);
            if (memberShape.hasTrait(DefaultTrait.class)) {
                DefaultTrait defaultTrait = memberShape.expectTrait(DefaultTrait.class);
                LOGGER.warning("Overwriting default trait for `" + memberShapeId + "` with value: `"
                        + Node.prettyPrintJson(defaultTrait.toNode()) + "`");
            }
            updates.add(memberShape.toBuilder()
                    .addTrait(new DefaultTrait(entry.getValue()))
                    .build());
        }

        return ModelTransformer.create().replaceShapes(model, updates);
    }
}
