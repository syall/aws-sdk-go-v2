// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
)

// A structure that contains information about one CloudWatch Logs account policy.
type AccountPolicy struct {

	// The Amazon Web Services account ID that the policy applies to.
	AccountId *string

	// The date and time that this policy was most recently updated.
	LastUpdatedTime *int64

	// The policy document for this account policy. The JSON specified in
	// policyDocument can be up to 30,720 characters.
	PolicyDocument *string

	// The name of the account policy.
	PolicyName *string

	// The type of policy for this account policy.
	PolicyType PolicyType

	// The scope of the account policy.
	Scope Scope

	noSmithyDocumentSerde
}

// This structure represents one anomaly that has been found by a logs anomaly
// detector. For more information about patterns and anomalies, see
// CreateLogAnomalyDetector (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_CreateLogAnomalyDetector.html)
// .
type Anomaly struct {

	// Specifies whether this anomaly is still ongoing.
	//
	// This member is required.
	Active *bool

	// The ARN of the anomaly detector that identified this anomaly.
	//
	// This member is required.
	AnomalyDetectorArn *string

	// The unique ID that CloudWatch Logs assigned to this anomaly.
	//
	// This member is required.
	AnomalyId *string

	// A human-readable description of the anomaly. This description is generated by
	// CloudWatch Logs.
	//
	// This member is required.
	Description *string

	// The date and time when the anomaly detector first saw this anomaly. It is
	// specified as epoch time, which is the number of seconds since January 1, 1970,
	// 00:00:00 UTC .
	//
	// This member is required.
	FirstSeen int64

	// A map showing times when the anomaly detector ran, and the number of
	// occurrences of this anomaly that were detected at each of those runs. The times
	// are specified in epoch time, which is the number of seconds since January 1,
	// 1970, 00:00:00 UTC .
	//
	// This member is required.
	Histogram map[string]int64

	// The date and time when the anomaly detector most recently saw this anomaly. It
	// is specified as epoch time, which is the number of seconds since January 1,
	// 1970, 00:00:00 UTC .
	//
	// This member is required.
	LastSeen int64

	// An array of ARNS of the log groups that contained log events considered to be
	// part of this anomaly.
	//
	// This member is required.
	LogGroupArnList []string

	// An array of sample log event messages that are considered to be part of this
	// anomaly.
	//
	// This member is required.
	LogSamples []string

	// The ID of the pattern used to help identify this anomaly.
	//
	// This member is required.
	PatternId *string

	// The pattern used to help identify this anomaly, in string format.
	//
	// This member is required.
	PatternString *string

	// An array of structures where each structure contains information about one
	// token that makes up the pattern.
	//
	// This member is required.
	PatternTokens []PatternToken

	// Indicates the current state of this anomaly. If it is still being treated as an
	// anomaly, the value is Active . If you have suppressed this anomaly by using the
	// UpdateAnomaly (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_UpdateAnomaly.html)
	// operation, the value is Suppressed . If this behavior is now considered to be
	// normal, the value is Baseline .
	//
	// This member is required.
	State State

	// If this anomaly is suppressed, this field is true if the suppression is because
	// the pattern is suppressed. If false , then only this particular anomaly is
	// suppressed.
	IsPatternLevelSuppression *bool

	// The pattern used to help identify this anomaly, in regular expression format.
	PatternRegex *string

	// The priority level of this anomaly, as determined by CloudWatch Logs. Priority
	// is computed based on log severity labels such as FATAL and ERROR and the amount
	// of deviation from the baseline. Possible values are HIGH , MEDIUM , and LOW .
	Priority *string

	// Indicates whether this anomaly is currently suppressed. To suppress an anomaly,
	// use UpdateAnomaly (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_UpdateAnomaly.html)
	// .
	Suppressed *bool

	// If the anomaly is suppressed, this indicates when it was suppressed.
	SuppressedDate int64

	// If the anomaly is suppressed, this indicates when the suppression will end. If
	// this value is 0 , the anomaly was suppressed with no expiration, with the
	// INFINITE value.
	SuppressedUntil int64

	noSmithyDocumentSerde
}

// Contains information about one anomaly detector in the account.
type AnomalyDetector struct {

	// The ARN of the anomaly detector.
	AnomalyDetectorArn *string

	// Specifies the current status of the anomaly detector. To pause an anomaly
	// detector, use the enabled parameter in the UpdateLogAnomalyDetector (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_UpdateLogAnomalyDetector.html)
	// operation.
	AnomalyDetectorStatus AnomalyDetectorStatus

	// The number of days used as the life cycle of anomalies. After this time,
	// anomalies are automatically baselined and the anomaly detector model will treat
	// new occurrences of similar event as normal.
	AnomalyVisibilityTime *int64

	// The date and time when this anomaly detector was created.
	CreationTimeStamp int64

	// The name of the anomaly detector.
	DetectorName *string

	// Specifies how often the anomaly detector runs and look for anomalies.
	EvaluationFrequency EvaluationFrequency

	// A symbolic description of how CloudWatch Logs should interpret the data in each
	// log event. For example, a log event can contain timestamps, IP addresses,
	// strings, and so on. You use the filter pattern to specify what to look for in
	// the log event message.
	FilterPattern *string

	// The ID of the KMS key assigned to this anomaly detector, if any.
	KmsKeyId *string

	// The date and time when this anomaly detector was most recently modified.
	LastModifiedTimeStamp int64

	// A list of the ARNs of the log groups that this anomaly detector watches.
	LogGroupArnList []string

	noSmithyDocumentSerde
}

// This structure contains information about one delivery in your account. A
// delivery is a connection between a logical delivery source and a logical
// delivery destination. For more information, see CreateDelivery (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_CreateDelivery.html)
// . You can't update an existing delivery. You can only create and delete
// deliveries.
type Delivery struct {

	// The Amazon Resource Name (ARN) that uniquely identifies this delivery.
	Arn *string

	// The ARN of the delivery destination that is associated with this delivery.
	DeliveryDestinationArn *string

	// Displays whether the delivery destination associated with this delivery is
	// CloudWatch Logs, Amazon S3, or Kinesis Data Firehose.
	DeliveryDestinationType DeliveryDestinationType

	// The name of the delivery source that is associated with this delivery.
	DeliverySourceName *string

	// The unique ID that identifies this delivery in your account.
	Id *string

	// The tags that have been assigned to this delivery.
	Tags map[string]string

	noSmithyDocumentSerde
}

// This structure contains information about one delivery destination in your
// account. A delivery destination is an Amazon Web Services resource that
// represents an Amazon Web Services service that logs can be sent to. CloudWatch
// Logs, Amazon S3, are supported as Kinesis Data Firehose delivery destinations.
// To configure logs delivery between a supported Amazon Web Services service and a
// destination, you must do the following:
//   - Create a delivery source, which is a logical object that represents the
//     resource that is actually sending the logs. For more information, see
//     PutDeliverySource (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDeliverySource.html)
//     .
//   - Create a delivery destination, which is a logical object that represents
//     the actual delivery destination.
//   - If you are delivering logs cross-account, you must use
//     PutDeliveryDestinationPolicy (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDeliveryDestinationolicy.html)
//     in the destination account to assign an IAM policy to the destination. This
//     policy allows delivery to that destination.
//   - Create a delivery by pairing exactly one delivery source and one delivery
//     destination. For more information, see CreateDelivery (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_CreateDelivery.html)
//     .
//
// You can configure a single delivery source to send logs to multiple
// destinations by creating multiple deliveries. You can also create multiple
// deliveries to configure multiple delivery sources to send logs to the same
// delivery destination.
type DeliveryDestination struct {

	// The Amazon Resource Name (ARN) that uniquely identifies this delivery
	// destination.
	Arn *string

	// A structure that contains the ARN of the Amazon Web Services resource that will
	// receive the logs.
	DeliveryDestinationConfiguration *DeliveryDestinationConfiguration

	// Displays whether this delivery destination is CloudWatch Logs, Amazon S3, or
	// Kinesis Data Firehose.
	DeliveryDestinationType DeliveryDestinationType

	// The name of this delivery destination.
	Name *string

	// The format of the logs that are sent to this delivery destination.
	OutputFormat OutputFormat

	// The tags that have been assigned to this delivery destination.
	Tags map[string]string

	noSmithyDocumentSerde
}

// A structure that contains information about one logs delivery destination.
type DeliveryDestinationConfiguration struct {

	// The ARN of the Amazon Web Services destination that this delivery destination
	// represents. That Amazon Web Services destination can be a log group in
	// CloudWatch Logs, an Amazon S3 bucket, or a delivery stream in Kinesis Data
	// Firehose.
	//
	// This member is required.
	DestinationResourceArn *string

	noSmithyDocumentSerde
}

// This structure contains information about one delivery source in your account.
// A delivery source is an Amazon Web Services resource that sends logs to an
// Amazon Web Services destination. The destination can be CloudWatch Logs, Amazon
// S3, or Kinesis Data Firehose. Only some Amazon Web Services services support
// being configured as a delivery source. These services are listed as Supported
// [V2 Permissions] in the table at Enabling logging from Amazon Web Services
// services. (https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-logs-and-resource-policy.html)
// To configure logs delivery between a supported Amazon Web Services service and a
// destination, you must do the following:
//   - Create a delivery source, which is a logical object that represents the
//     resource that is actually sending the logs. For more information, see
//     PutDeliverySource (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDeliverySource.html)
//     .
//   - Create a delivery destination, which is a logical object that represents
//     the actual delivery destination. For more information, see
//     PutDeliveryDestination (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDeliveryDestination.html)
//     .
//   - If you are delivering logs cross-account, you must use
//     PutDeliveryDestinationPolicy (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDeliveryDestinationolicy.html)
//     in the destination account to assign an IAM policy to the destination. This
//     policy allows delivery to that destination.
//   - Create a delivery by pairing exactly one delivery source and one delivery
//     destination. For more information, see CreateDelivery (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_CreateDelivery.html)
//     .
//
// You can configure a single delivery source to send logs to multiple
// destinations by creating multiple deliveries. You can also create multiple
// deliveries to configure multiple delivery sources to send logs to the same
// delivery destination.
type DeliverySource struct {

	// The Amazon Resource Name (ARN) that uniquely identifies this delivery source.
	Arn *string

	// The type of log that the source is sending. For valid values for this
	// parameter, see the documentation for the source service.
	LogType *string

	// The unique name of the delivery source.
	Name *string

	// This array contains the ARN of the Amazon Web Services resource that sends logs
	// and is represented by this delivery source. Currently, only one ARN can be in
	// the array.
	ResourceArns []string

	// The Amazon Web Services service that is sending logs.
	Service *string

	// The tags that have been assigned to this delivery source.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Represents a cross-account destination that receives subscription log events.
type Destination struct {

	// An IAM policy document that governs which Amazon Web Services accounts can
	// create subscription filters against this destination.
	AccessPolicy *string

	// The ARN of this destination.
	Arn *string

	// The creation time of the destination, expressed as the number of milliseconds
	// after Jan 1, 1970 00:00:00 UTC.
	CreationTime *int64

	// The name of the destination.
	DestinationName *string

	// A role for impersonation, used when delivering log events to the target.
	RoleArn *string

	// The Amazon Resource Name (ARN) of the physical target where the log events are
	// delivered (for example, a Kinesis stream).
	TargetArn *string

	noSmithyDocumentSerde
}

// Represents an export task.
type ExportTask struct {

	// The name of the S3 bucket to which the log data was exported.
	Destination *string

	// The prefix that was used as the start of Amazon S3 key for every object
	// exported.
	DestinationPrefix *string

	// Execution information about the export task.
	ExecutionInfo *ExportTaskExecutionInfo

	// The start time, expressed as the number of milliseconds after Jan 1, 1970
	// 00:00:00 UTC . Events with a timestamp before this time are not exported.
	From *int64

	// The name of the log group from which logs data was exported.
	LogGroupName *string

	// The status of the export task.
	Status *ExportTaskStatus

	// The ID of the export task.
	TaskId *string

	// The name of the export task.
	TaskName *string

	// The end time, expressed as the number of milliseconds after Jan 1, 1970
	// 00:00:00 UTC . Events with a timestamp later than this time are not exported.
	To *int64

	noSmithyDocumentSerde
}

// Represents the status of an export task.
type ExportTaskExecutionInfo struct {

	// The completion time of the export task, expressed as the number of milliseconds
	// after Jan 1, 1970 00:00:00 UTC .
	CompletionTime *int64

	// The creation time of the export task, expressed as the number of milliseconds
	// after Jan 1, 1970 00:00:00 UTC .
	CreationTime *int64

	noSmithyDocumentSerde
}

// Represents the status of an export task.
type ExportTaskStatus struct {

	// The status code of the export task.
	Code ExportTaskStatusCode

	// The status message related to the status code.
	Message *string

	noSmithyDocumentSerde
}

// Represents a matched event.
type FilteredLogEvent struct {

	// The ID of the event.
	EventId *string

	// The time the event was ingested, expressed as the number of milliseconds after
	// Jan 1, 1970 00:00:00 UTC .
	IngestionTime *int64

	// The name of the log stream to which this event belongs.
	LogStreamName *string

	// The data contained in the log event.
	Message *string

	// The time the event occurred, expressed as the number of milliseconds after Jan
	// 1, 1970 00:00:00 UTC .
	Timestamp *int64

	noSmithyDocumentSerde
}

// Represents a log event, which is a record of activity that was recorded by the
// application or resource being monitored.
type InputLogEvent struct {

	// The raw event message. Each log event can be no larger than 256 KB.
	//
	// This member is required.
	Message *string

	// The time the event occurred, expressed as the number of milliseconds after Jan
	// 1, 1970 00:00:00 UTC .
	//
	// This member is required.
	Timestamp *int64

	noSmithyDocumentSerde
}

// Represents a log group.
type LogGroup struct {

	// The Amazon Resource Name (ARN) of the log group.
	Arn *string

	// The creation time of the log group, expressed as the number of milliseconds
	// after Jan 1, 1970 00:00:00 UTC.
	CreationTime *int64

	// Displays whether this log group has a protection policy, or whether it had one
	// in the past. For more information, see PutDataProtectionPolicy (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDataProtectionPolicy.html)
	// .
	DataProtectionStatus DataProtectionStatus

	// Displays all the properties that this log group has inherited from
	// account-level settings.
	InheritedProperties []InheritedProperty

	// The Amazon Resource Name (ARN) of the KMS key to use when encrypting log data.
	KmsKeyId *string

	// This specifies the log group class for this log group. There are two classes:
	//   - The Standard log class supports all CloudWatch Logs features.
	//   - The Infrequent Access log class supports a subset of CloudWatch Logs
	//   features and incurs lower costs.
	// For details about the features supported by each class, see Log classes (https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch_Logs_Log_Classes.html)
	LogGroupClass LogGroupClass

	// The name of the log group.
	LogGroupName *string

	// The number of metric filters.
	MetricFilterCount *int32

	// The number of days to retain the log events in the specified log group.
	// Possible values are: 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545,
	// 731, 1096, 1827, 2192, 2557, 2922, 3288, and 3653. To set a log group so that
	// its log events do not expire, use DeleteRetentionPolicy (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_DeleteRetentionPolicy.html)
	// .
	RetentionInDays *int32

	// The number of bytes stored.
	StoredBytes *int64

	noSmithyDocumentSerde
}

// The fields contained in log events found by a GetLogGroupFields operation,
// along with the percentage of queried log events in which each field appears.
type LogGroupField struct {

	// The name of a log field.
	Name *string

	// The percentage of log events queried that contained the field.
	Percent int32

	noSmithyDocumentSerde
}

// Represents a log stream, which is a sequence of log events from a single
// emitter of logs.
type LogStream struct {

	// The Amazon Resource Name (ARN) of the log stream.
	Arn *string

	// The creation time of the stream, expressed as the number of milliseconds after
	// Jan 1, 1970 00:00:00 UTC .
	CreationTime *int64

	// The time of the first event, expressed as the number of milliseconds after Jan
	// 1, 1970 00:00:00 UTC .
	FirstEventTimestamp *int64

	// The time of the most recent log event in the log stream in CloudWatch Logs.
	// This number is expressed as the number of milliseconds after Jan 1, 1970
	// 00:00:00 UTC . The lastEventTime value updates on an eventual consistency
	// basis. It typically updates in less than an hour from ingestion, but in rare
	// situations might take longer.
	LastEventTimestamp *int64

	// The ingestion time, expressed as the number of milliseconds after Jan 1, 1970
	// 00:00:00 UTC The lastIngestionTime value updates on an eventual consistency
	// basis. It typically updates in less than an hour after ingestion, but in rare
	// situations might take longer.
	LastIngestionTime *int64

	// The name of the log stream.
	LogStreamName *string

	// The number of bytes stored. Important: As of June 17, 2019, this parameter is
	// no longer supported for log streams, and is always reported as zero. This change
	// applies only to log streams. The storedBytes parameter for log groups is not
	// affected.
	//
	// Deprecated: Starting on June 17, 2019, this parameter will be deprecated for
	// log streams, and will be reported as zero. This change applies only to log
	// streams. The storedBytes parameter for log groups is not affected.
	StoredBytes *int64

	// The sequence token. The sequence token is now ignored in PutLogEvents actions.
	// PutLogEvents actions are always accepted regardless of receiving an invalid
	// sequence token. You don't need to obtain uploadSequenceToken to use a
	// PutLogEvents action.
	UploadSequenceToken *string

	noSmithyDocumentSerde
}

// Metric filters express how CloudWatch Logs would extract metric observations
// from ingested log events and transform them into metric data in a CloudWatch
// metric.
type MetricFilter struct {

	// The creation time of the metric filter, expressed as the number of milliseconds
	// after Jan 1, 1970 00:00:00 UTC .
	CreationTime *int64

	// The name of the metric filter.
	FilterName *string

	// A symbolic description of how CloudWatch Logs should interpret the data in each
	// log event. For example, a log event can contain timestamps, IP addresses,
	// strings, and so on. You use the filter pattern to specify what to look for in
	// the log event message.
	FilterPattern *string

	// The name of the log group.
	LogGroupName *string

	// The metric transformations.
	MetricTransformations []MetricTransformation

	noSmithyDocumentSerde
}

// Represents a matched event.
type MetricFilterMatchRecord struct {

	// The raw event data.
	EventMessage *string

	// The event number.
	EventNumber int64

	// The values extracted from the event data by the filter.
	ExtractedValues map[string]string

	noSmithyDocumentSerde
}

// Indicates how to transform ingested log events to metric data in a CloudWatch
// metric.
type MetricTransformation struct {

	// The name of the CloudWatch metric.
	//
	// This member is required.
	MetricName *string

	// A custom namespace to contain your metric in CloudWatch. Use namespaces to
	// group together metrics that are similar. For more information, see Namespaces (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cloudwatch_concepts.html#Namespace)
	// .
	//
	// This member is required.
	MetricNamespace *string

	// The value to publish to the CloudWatch metric when a filter pattern matches a
	// log event.
	//
	// This member is required.
	MetricValue *string

	// (Optional) The value to emit when a filter pattern does not match a log event.
	// This value can be null.
	DefaultValue *float64

	// The fields to use as dimensions for the metric. One metric filter can include
	// as many as three dimensions. Metrics extracted from log events are charged as
	// custom metrics. To prevent unexpected high charges, do not specify
	// high-cardinality fields such as IPAddress or requestID as dimensions. Each
	// different value found for a dimension is treated as a separate metric and
	// accrues charges as a separate custom metric. CloudWatch Logs disables a metric
	// filter if it generates 1000 different name/value pairs for your specified
	// dimensions within a certain amount of time. This helps to prevent accidental
	// high charges. You can also set up a billing alarm to alert you if your charges
	// are higher than expected. For more information, see Creating a Billing Alarm to
	// Monitor Your Estimated Amazon Web Services Charges (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/monitor_estimated_charges_with_cloudwatch.html)
	// .
	Dimensions map[string]string

	// The unit to assign to the metric. If you omit this, the unit is set as None .
	Unit StandardUnit

	noSmithyDocumentSerde
}

// Represents a log event.
type OutputLogEvent struct {

	// The time the event was ingested, expressed as the number of milliseconds after
	// Jan 1, 1970 00:00:00 UTC .
	IngestionTime *int64

	// The data contained in the log event.
	Message *string

	// The time the event occurred, expressed as the number of milliseconds after Jan
	// 1, 1970 00:00:00 UTC .
	Timestamp *int64

	noSmithyDocumentSerde
}

// A tructures that contains information about one pattern token related to an
// anomaly. For more information about patterns and tokens, see
// CreateLogAnomalyDetector (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_CreateLogAnomalyDetector.html)
// .
type PatternToken struct {

	// For a dynamic token, this indicates where in the pattern that this token
	// appears, related to other dynamic tokens. The dynamic token that appears first
	// has a value of 1 , the one that appears second is 2 , and so on.
	DynamicTokenPosition int32

	// Contains the values found for a dynamic token, and the number of times each
	// value was found.
	Enumerations map[string]int64

	// Specifies whether this is a dynamic token.
	IsDynamic *bool

	// The string represented by this token. If this is a dynamic token, the value
	// will be <*>
	TokenString *string

	noSmithyDocumentSerde
}

// A structure that contains information about one delivery destination policy.
type Policy struct {

	// The contents of the delivery destination policy.
	DeliveryDestinationPolicy *string

	noSmithyDocumentSerde
}

// Reserved.
type QueryCompileError struct {

	// Reserved.
	Location *QueryCompileErrorLocation

	// Reserved.
	Message *string

	noSmithyDocumentSerde
}

// Reserved.
type QueryCompileErrorLocation struct {

	// Reserved.
	EndCharOffset *int32

	// Reserved.
	StartCharOffset *int32

	noSmithyDocumentSerde
}

// This structure contains details about a saved CloudWatch Logs Insights query
// definition.
type QueryDefinition struct {

	// The date that the query definition was most recently modified.
	LastModified *int64

	// If this query definition contains a list of log groups that it is limited to,
	// that list appears here.
	LogGroupNames []string

	// The name of the query definition.
	Name *string

	// The unique ID of the query definition.
	QueryDefinitionId *string

	// The query string to use for this definition. For more information, see
	// CloudWatch Logs Insights Query Syntax (https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax.html)
	// .
	QueryString *string

	noSmithyDocumentSerde
}

// Information about one CloudWatch Logs Insights query that matches the request
// in a DescribeQueries operation.
type QueryInfo struct {

	// The date and time that this query was created.
	CreateTime *int64

	// The name of the log group scanned by this query.
	LogGroupName *string

	// The unique ID number of this query.
	QueryId *string

	// The query string used in this query.
	QueryString *string

	// The status of this query. Possible values are Cancelled , Complete , Failed ,
	// Running , Scheduled , and Unknown .
	Status QueryStatus

	noSmithyDocumentSerde
}

// Contains the number of log events scanned by the query, the number of log
// events that matched the query criteria, and the total number of bytes in the log
// events that were scanned.
type QueryStatistics struct {

	// The total number of bytes in the log events scanned during the query.
	BytesScanned float64

	// The number of log events that matched the query string.
	RecordsMatched float64

	// The total number of log events scanned during the query.
	RecordsScanned float64

	noSmithyDocumentSerde
}

// Represents the rejected events.
type RejectedLogEventsInfo struct {

	// The expired log events.
	ExpiredLogEventEndIndex *int32

	// The log events that are too new.
	TooNewLogEventStartIndex *int32

	// The log events that are dated too far in the past.
	TooOldLogEventEndIndex *int32

	noSmithyDocumentSerde
}

// A policy enabling one or more entities to put logs to a log group in this
// account.
type ResourcePolicy struct {

	// Timestamp showing when this policy was last updated, expressed as the number of
	// milliseconds after Jan 1, 1970 00:00:00 UTC .
	LastUpdatedTime *int64

	// The details of the policy.
	PolicyDocument *string

	// The name of the resource policy.
	PolicyName *string

	noSmithyDocumentSerde
}

// Contains one field from one log event returned by a CloudWatch Logs Insights
// query, along with the value of that field. For more information about the fields
// that are generated by CloudWatch logs, see Supported Logs and Discovered Fields (https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_AnalyzeLogData-discoverable-fields.html)
// .
type ResultField struct {

	// The log event field.
	Field *string

	// The value of this field.
	Value *string

	noSmithyDocumentSerde
}

// Represents the search status of a log stream.
type SearchedLogStream struct {

	// The name of the log stream.
	LogStreamName *string

	// Indicates whether all the events in this log stream were searched.
	SearchedCompletely *bool

	noSmithyDocumentSerde
}

// Represents a subscription filter.
type SubscriptionFilter struct {

	// The creation time of the subscription filter, expressed as the number of
	// milliseconds after Jan 1, 1970 00:00:00 UTC .
	CreationTime *int64

	// The Amazon Resource Name (ARN) of the destination.
	DestinationArn *string

	// The method used to distribute log data to the destination, which can be either
	// random or grouped by log stream.
	Distribution Distribution

	// The name of the subscription filter.
	FilterName *string

	// A symbolic description of how CloudWatch Logs should interpret the data in each
	// log event. For example, a log event can contain timestamps, IP addresses,
	// strings, and so on. You use the filter pattern to specify what to look for in
	// the log event message.
	FilterPattern *string

	// The name of the log group.
	LogGroupName *string

	//
	RoleArn *string

	noSmithyDocumentSerde
}

// If you are suppressing an anomaly temporariliy, this structure defines how long
// the suppression period is to be.
type SuppressionPeriod struct {

	// Specifies whether the value of value is in seconds, minutes, or hours.
	SuppressionUnit SuppressionUnit

	// Specifies the number of seconds, minutes or hours to suppress this anomaly.
	// There is no maximum.
	Value int32

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
