// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type BuildType string

// Enum values for BuildType
const (
	BuildTypeUserInitiated BuildType = "USER_INITIATED"
	BuildTypeScheduled     BuildType = "SCHEDULED"
	BuildTypeImport        BuildType = "IMPORT"
)

// Values returns all known values for BuildType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (BuildType) Values() []BuildType {
	return []BuildType{
		"USER_INITIATED",
		"SCHEDULED",
		"IMPORT",
	}
}

type ComponentFormat string

// Enum values for ComponentFormat
const (
	ComponentFormatShell ComponentFormat = "SHELL"
)

// Values returns all known values for ComponentFormat. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ComponentFormat) Values() []ComponentFormat {
	return []ComponentFormat{
		"SHELL",
	}
}

type ComponentStatus string

// Enum values for ComponentStatus
const (
	ComponentStatusDeprecated ComponentStatus = "DEPRECATED"
)

// Values returns all known values for ComponentStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ComponentStatus) Values() []ComponentStatus {
	return []ComponentStatus{
		"DEPRECATED",
	}
}

type ComponentType string

// Enum values for ComponentType
const (
	ComponentTypeBuild ComponentType = "BUILD"
	ComponentTypeTest  ComponentType = "TEST"
)

// Values returns all known values for ComponentType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ComponentType) Values() []ComponentType {
	return []ComponentType{
		"BUILD",
		"TEST",
	}
}

type ContainerRepositoryService string

// Enum values for ContainerRepositoryService
const (
	ContainerRepositoryServiceEcr ContainerRepositoryService = "ECR"
)

// Values returns all known values for ContainerRepositoryService. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ContainerRepositoryService) Values() []ContainerRepositoryService {
	return []ContainerRepositoryService{
		"ECR",
	}
}

type ContainerType string

// Enum values for ContainerType
const (
	ContainerTypeDocker ContainerType = "DOCKER"
)

// Values returns all known values for ContainerType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ContainerType) Values() []ContainerType {
	return []ContainerType{
		"DOCKER",
	}
}

type DiskImageFormat string

// Enum values for DiskImageFormat
const (
	DiskImageFormatVmdk DiskImageFormat = "VMDK"
	DiskImageFormatRaw  DiskImageFormat = "RAW"
	DiskImageFormatVhd  DiskImageFormat = "VHD"
)

// Values returns all known values for DiskImageFormat. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DiskImageFormat) Values() []DiskImageFormat {
	return []DiskImageFormat{
		"VMDK",
		"RAW",
		"VHD",
	}
}

type EbsVolumeType string

// Enum values for EbsVolumeType
const (
	EbsVolumeTypeStandard EbsVolumeType = "standard"
	EbsVolumeTypeIo1      EbsVolumeType = "io1"
	EbsVolumeTypeIo2      EbsVolumeType = "io2"
	EbsVolumeTypeGp2      EbsVolumeType = "gp2"
	EbsVolumeTypeGp3      EbsVolumeType = "gp3"
	EbsVolumeTypeSc1      EbsVolumeType = "sc1"
	EbsVolumeTypeSt1      EbsVolumeType = "st1"
)

// Values returns all known values for EbsVolumeType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EbsVolumeType) Values() []EbsVolumeType {
	return []EbsVolumeType{
		"standard",
		"io1",
		"io2",
		"gp2",
		"gp3",
		"sc1",
		"st1",
	}
}

type ImageScanStatus string

// Enum values for ImageScanStatus
const (
	ImageScanStatusPending    ImageScanStatus = "PENDING"
	ImageScanStatusScanning   ImageScanStatus = "SCANNING"
	ImageScanStatusCollecting ImageScanStatus = "COLLECTING"
	ImageScanStatusCompleted  ImageScanStatus = "COMPLETED"
	ImageScanStatusAbandoned  ImageScanStatus = "ABANDONED"
	ImageScanStatusFailed     ImageScanStatus = "FAILED"
	ImageScanStatusTimedOut   ImageScanStatus = "TIMED_OUT"
)

// Values returns all known values for ImageScanStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ImageScanStatus) Values() []ImageScanStatus {
	return []ImageScanStatus{
		"PENDING",
		"SCANNING",
		"COLLECTING",
		"COMPLETED",
		"ABANDONED",
		"FAILED",
		"TIMED_OUT",
	}
}

type ImageSource string

// Enum values for ImageSource
const (
	ImageSourceAmazonManaged  ImageSource = "AMAZON_MANAGED"
	ImageSourceAwsMarketplace ImageSource = "AWS_MARKETPLACE"
	ImageSourceImported       ImageSource = "IMPORTED"
	ImageSourceCustom         ImageSource = "CUSTOM"
)

// Values returns all known values for ImageSource. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ImageSource) Values() []ImageSource {
	return []ImageSource{
		"AMAZON_MANAGED",
		"AWS_MARKETPLACE",
		"IMPORTED",
		"CUSTOM",
	}
}

type ImageStatus string

// Enum values for ImageStatus
const (
	ImageStatusPending      ImageStatus = "PENDING"
	ImageStatusCreating     ImageStatus = "CREATING"
	ImageStatusBuilding     ImageStatus = "BUILDING"
	ImageStatusTesting      ImageStatus = "TESTING"
	ImageStatusDistributing ImageStatus = "DISTRIBUTING"
	ImageStatusIntegrating  ImageStatus = "INTEGRATING"
	ImageStatusAvailable    ImageStatus = "AVAILABLE"
	ImageStatusCancelled    ImageStatus = "CANCELLED"
	ImageStatusFailed       ImageStatus = "FAILED"
	ImageStatusDeprecated   ImageStatus = "DEPRECATED"
	ImageStatusDeleted      ImageStatus = "DELETED"
	ImageStatusDisabled     ImageStatus = "DISABLED"
)

// Values returns all known values for ImageStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ImageStatus) Values() []ImageStatus {
	return []ImageStatus{
		"PENDING",
		"CREATING",
		"BUILDING",
		"TESTING",
		"DISTRIBUTING",
		"INTEGRATING",
		"AVAILABLE",
		"CANCELLED",
		"FAILED",
		"DEPRECATED",
		"DELETED",
		"DISABLED",
	}
}

type ImageType string

// Enum values for ImageType
const (
	ImageTypeAmi    ImageType = "AMI"
	ImageTypeDocker ImageType = "DOCKER"
)

// Values returns all known values for ImageType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ImageType) Values() []ImageType {
	return []ImageType{
		"AMI",
		"DOCKER",
	}
}

type LifecycleExecutionResourceActionName string

// Enum values for LifecycleExecutionResourceActionName
const (
	LifecycleExecutionResourceActionNameAvailable LifecycleExecutionResourceActionName = "AVAILABLE"
	LifecycleExecutionResourceActionNameDelete    LifecycleExecutionResourceActionName = "DELETE"
	LifecycleExecutionResourceActionNameDeprecate LifecycleExecutionResourceActionName = "DEPRECATE"
	LifecycleExecutionResourceActionNameDisable   LifecycleExecutionResourceActionName = "DISABLE"
)

// Values returns all known values for LifecycleExecutionResourceActionName. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (LifecycleExecutionResourceActionName) Values() []LifecycleExecutionResourceActionName {
	return []LifecycleExecutionResourceActionName{
		"AVAILABLE",
		"DELETE",
		"DEPRECATE",
		"DISABLE",
	}
}

type LifecycleExecutionResourceStatus string

// Enum values for LifecycleExecutionResourceStatus
const (
	LifecycleExecutionResourceStatusFailed     LifecycleExecutionResourceStatus = "FAILED"
	LifecycleExecutionResourceStatusInProgress LifecycleExecutionResourceStatus = "IN_PROGRESS"
	LifecycleExecutionResourceStatusSkipped    LifecycleExecutionResourceStatus = "SKIPPED"
	LifecycleExecutionResourceStatusSuccess    LifecycleExecutionResourceStatus = "SUCCESS"
)

// Values returns all known values for LifecycleExecutionResourceStatus. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (LifecycleExecutionResourceStatus) Values() []LifecycleExecutionResourceStatus {
	return []LifecycleExecutionResourceStatus{
		"FAILED",
		"IN_PROGRESS",
		"SKIPPED",
		"SUCCESS",
	}
}

type LifecycleExecutionStatus string

// Enum values for LifecycleExecutionStatus
const (
	LifecycleExecutionStatusInProgress LifecycleExecutionStatus = "IN_PROGRESS"
	LifecycleExecutionStatusCancelled  LifecycleExecutionStatus = "CANCELLED"
	LifecycleExecutionStatusCancelling LifecycleExecutionStatus = "CANCELLING"
	LifecycleExecutionStatusFailed     LifecycleExecutionStatus = "FAILED"
	LifecycleExecutionStatusSuccess    LifecycleExecutionStatus = "SUCCESS"
)

// Values returns all known values for LifecycleExecutionStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (LifecycleExecutionStatus) Values() []LifecycleExecutionStatus {
	return []LifecycleExecutionStatus{
		"IN_PROGRESS",
		"CANCELLED",
		"CANCELLING",
		"FAILED",
		"SUCCESS",
	}
}

type LifecyclePolicyDetailActionType string

// Enum values for LifecyclePolicyDetailActionType
const (
	LifecyclePolicyDetailActionTypeDelete    LifecyclePolicyDetailActionType = "DELETE"
	LifecyclePolicyDetailActionTypeDeprecate LifecyclePolicyDetailActionType = "DEPRECATE"
	LifecyclePolicyDetailActionTypeDisable   LifecyclePolicyDetailActionType = "DISABLE"
)

// Values returns all known values for LifecyclePolicyDetailActionType. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (LifecyclePolicyDetailActionType) Values() []LifecyclePolicyDetailActionType {
	return []LifecyclePolicyDetailActionType{
		"DELETE",
		"DEPRECATE",
		"DISABLE",
	}
}

type LifecyclePolicyDetailFilterType string

// Enum values for LifecyclePolicyDetailFilterType
const (
	LifecyclePolicyDetailFilterTypeAge   LifecyclePolicyDetailFilterType = "AGE"
	LifecyclePolicyDetailFilterTypeCount LifecyclePolicyDetailFilterType = "COUNT"
)

// Values returns all known values for LifecyclePolicyDetailFilterType. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (LifecyclePolicyDetailFilterType) Values() []LifecyclePolicyDetailFilterType {
	return []LifecyclePolicyDetailFilterType{
		"AGE",
		"COUNT",
	}
}

type LifecyclePolicyResourceType string

// Enum values for LifecyclePolicyResourceType
const (
	LifecyclePolicyResourceTypeAmiImage       LifecyclePolicyResourceType = "AMI_IMAGE"
	LifecyclePolicyResourceTypeContainerImage LifecyclePolicyResourceType = "CONTAINER_IMAGE"
)

// Values returns all known values for LifecyclePolicyResourceType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (LifecyclePolicyResourceType) Values() []LifecyclePolicyResourceType {
	return []LifecyclePolicyResourceType{
		"AMI_IMAGE",
		"CONTAINER_IMAGE",
	}
}

type LifecyclePolicyStatus string

// Enum values for LifecyclePolicyStatus
const (
	LifecyclePolicyStatusDisabled LifecyclePolicyStatus = "DISABLED"
	LifecyclePolicyStatusEnabled  LifecyclePolicyStatus = "ENABLED"
)

// Values returns all known values for LifecyclePolicyStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (LifecyclePolicyStatus) Values() []LifecyclePolicyStatus {
	return []LifecyclePolicyStatus{
		"DISABLED",
		"ENABLED",
	}
}

type LifecyclePolicyTimeUnit string

// Enum values for LifecyclePolicyTimeUnit
const (
	LifecyclePolicyTimeUnitDays   LifecyclePolicyTimeUnit = "DAYS"
	LifecyclePolicyTimeUnitWeeks  LifecyclePolicyTimeUnit = "WEEKS"
	LifecyclePolicyTimeUnitMonths LifecyclePolicyTimeUnit = "MONTHS"
	LifecyclePolicyTimeUnitYears  LifecyclePolicyTimeUnit = "YEARS"
)

// Values returns all known values for LifecyclePolicyTimeUnit. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (LifecyclePolicyTimeUnit) Values() []LifecyclePolicyTimeUnit {
	return []LifecyclePolicyTimeUnit{
		"DAYS",
		"WEEKS",
		"MONTHS",
		"YEARS",
	}
}

type Ownership string

// Enum values for Ownership
const (
	OwnershipSelf       Ownership = "Self"
	OwnershipShared     Ownership = "Shared"
	OwnershipAmazon     Ownership = "Amazon"
	OwnershipThirdparty Ownership = "ThirdParty"
)

// Values returns all known values for Ownership. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (Ownership) Values() []Ownership {
	return []Ownership{
		"Self",
		"Shared",
		"Amazon",
		"ThirdParty",
	}
}

type PipelineExecutionStartCondition string

// Enum values for PipelineExecutionStartCondition
const (
	PipelineExecutionStartConditionExpressionMatchOnly                          PipelineExecutionStartCondition = "EXPRESSION_MATCH_ONLY"
	PipelineExecutionStartConditionExpressionMatchAndDependencyUpdatesAvailable PipelineExecutionStartCondition = "EXPRESSION_MATCH_AND_DEPENDENCY_UPDATES_AVAILABLE"
)

// Values returns all known values for PipelineExecutionStartCondition. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (PipelineExecutionStartCondition) Values() []PipelineExecutionStartCondition {
	return []PipelineExecutionStartCondition{
		"EXPRESSION_MATCH_ONLY",
		"EXPRESSION_MATCH_AND_DEPENDENCY_UPDATES_AVAILABLE",
	}
}

type PipelineStatus string

// Enum values for PipelineStatus
const (
	PipelineStatusDisabled PipelineStatus = "DISABLED"
	PipelineStatusEnabled  PipelineStatus = "ENABLED"
)

// Values returns all known values for PipelineStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PipelineStatus) Values() []PipelineStatus {
	return []PipelineStatus{
		"DISABLED",
		"ENABLED",
	}
}

type Platform string

// Enum values for Platform
const (
	PlatformWindows Platform = "Windows"
	PlatformLinux   Platform = "Linux"
)

// Values returns all known values for Platform. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Platform) Values() []Platform {
	return []Platform{
		"Windows",
		"Linux",
	}
}

type ResourceStatus string

// Enum values for ResourceStatus
const (
	ResourceStatusAvailable  ResourceStatus = "AVAILABLE"
	ResourceStatusDeleted    ResourceStatus = "DELETED"
	ResourceStatusDeprecated ResourceStatus = "DEPRECATED"
	ResourceStatusDisabled   ResourceStatus = "DISABLED"
)

// Values returns all known values for ResourceStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ResourceStatus) Values() []ResourceStatus {
	return []ResourceStatus{
		"AVAILABLE",
		"DELETED",
		"DEPRECATED",
		"DISABLED",
	}
}

type WorkflowExecutionStatus string

// Enum values for WorkflowExecutionStatus
const (
	WorkflowExecutionStatusPending            WorkflowExecutionStatus = "PENDING"
	WorkflowExecutionStatusSkipped            WorkflowExecutionStatus = "SKIPPED"
	WorkflowExecutionStatusRunning            WorkflowExecutionStatus = "RUNNING"
	WorkflowExecutionStatusCompleted          WorkflowExecutionStatus = "COMPLETED"
	WorkflowExecutionStatusFailed             WorkflowExecutionStatus = "FAILED"
	WorkflowExecutionStatusRollbackInProgress WorkflowExecutionStatus = "ROLLBACK_IN_PROGRESS"
	WorkflowExecutionStatusRollbackCompleted  WorkflowExecutionStatus = "ROLLBACK_COMPLETED"
)

// Values returns all known values for WorkflowExecutionStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (WorkflowExecutionStatus) Values() []WorkflowExecutionStatus {
	return []WorkflowExecutionStatus{
		"PENDING",
		"SKIPPED",
		"RUNNING",
		"COMPLETED",
		"FAILED",
		"ROLLBACK_IN_PROGRESS",
		"ROLLBACK_COMPLETED",
	}
}

type WorkflowStepExecutionRollbackStatus string

// Enum values for WorkflowStepExecutionRollbackStatus
const (
	WorkflowStepExecutionRollbackStatusRunning   WorkflowStepExecutionRollbackStatus = "RUNNING"
	WorkflowStepExecutionRollbackStatusCompleted WorkflowStepExecutionRollbackStatus = "COMPLETED"
	WorkflowStepExecutionRollbackStatusSkipped   WorkflowStepExecutionRollbackStatus = "SKIPPED"
	WorkflowStepExecutionRollbackStatusFailed    WorkflowStepExecutionRollbackStatus = "FAILED"
)

// Values returns all known values for WorkflowStepExecutionRollbackStatus. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (WorkflowStepExecutionRollbackStatus) Values() []WorkflowStepExecutionRollbackStatus {
	return []WorkflowStepExecutionRollbackStatus{
		"RUNNING",
		"COMPLETED",
		"SKIPPED",
		"FAILED",
	}
}

type WorkflowStepExecutionStatus string

// Enum values for WorkflowStepExecutionStatus
const (
	WorkflowStepExecutionStatusPending   WorkflowStepExecutionStatus = "PENDING"
	WorkflowStepExecutionStatusSkipped   WorkflowStepExecutionStatus = "SKIPPED"
	WorkflowStepExecutionStatusRunning   WorkflowStepExecutionStatus = "RUNNING"
	WorkflowStepExecutionStatusCompleted WorkflowStepExecutionStatus = "COMPLETED"
	WorkflowStepExecutionStatusFailed    WorkflowStepExecutionStatus = "FAILED"
)

// Values returns all known values for WorkflowStepExecutionStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (WorkflowStepExecutionStatus) Values() []WorkflowStepExecutionStatus {
	return []WorkflowStepExecutionStatus{
		"PENDING",
		"SKIPPED",
		"RUNNING",
		"COMPLETED",
		"FAILED",
	}
}

type WorkflowType string

// Enum values for WorkflowType
const (
	WorkflowTypeBuild        WorkflowType = "BUILD"
	WorkflowTypeTest         WorkflowType = "TEST"
	WorkflowTypeDistribution WorkflowType = "DISTRIBUTION"
)

// Values returns all known values for WorkflowType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (WorkflowType) Values() []WorkflowType {
	return []WorkflowType{
		"BUILD",
		"TEST",
		"DISTRIBUTION",
	}
}
