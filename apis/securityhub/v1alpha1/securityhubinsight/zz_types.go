/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by terrajet. DO NOT EDIT.

// +kubebuilder:object:generate=true
// +groupName=securityhub.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	apis "github.com/crossplane-contrib/provider-tf-aws/apis"
	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/securityhub/v1alpha1"
)

type AwsAccountIdObservation struct {
}

type AwsAccountIdParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type CompanyNameObservation struct {
}

type CompanyNameParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type ComplianceStatusObservation struct {
}

type ComplianceStatusParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type ConfidenceObservation struct {
}

type ConfidenceParameters struct {
	Eq *string `json:"eq,omitempty" tf:"eq"`

	Gte *string `json:"gte,omitempty" tf:"gte"`

	Lte *string `json:"lte,omitempty" tf:"lte"`
}

type CreatedAtObservation struct {
}

type CreatedAtParameters struct {
	DateRange []DateRangeParameters `json:"dateRange,omitempty" tf:"date_range"`

	End *string `json:"end,omitempty" tf:"end"`

	Start *string `json:"start,omitempty" tf:"start"`
}

type CriticalityObservation struct {
}

type CriticalityParameters struct {
	Eq *string `json:"eq,omitempty" tf:"eq"`

	Gte *string `json:"gte,omitempty" tf:"gte"`

	Lte *string `json:"lte,omitempty" tf:"lte"`
}

type DateRangeObservation struct {
}

type DateRangeParameters struct {
	Unit string `json:"unit" tf:"unit"`

	Value int64 `json:"value" tf:"value"`
}

type DescriptionObservation struct {
}

type DescriptionParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type FiltersObservation struct {
}

type FiltersParameters struct {
	AwsAccountId []AwsAccountIdParameters `json:"awsAccountId,omitempty" tf:"aws_account_id"`

	CompanyName []CompanyNameParameters `json:"companyName,omitempty" tf:"company_name"`

	ComplianceStatus []ComplianceStatusParameters `json:"complianceStatus,omitempty" tf:"compliance_status"`

	Confidence []ConfidenceParameters `json:"confidence,omitempty" tf:"confidence"`

	CreatedAt []CreatedAtParameters `json:"createdAt,omitempty" tf:"created_at"`

	Criticality []CriticalityParameters `json:"criticality,omitempty" tf:"criticality"`

	Description []DescriptionParameters `json:"description,omitempty" tf:"description"`

	FindingProviderFieldsConfidence []FindingProviderFieldsConfidenceParameters `json:"findingProviderFieldsConfidence,omitempty" tf:"finding_provider_fields_confidence"`

	FindingProviderFieldsCriticality []FindingProviderFieldsCriticalityParameters `json:"findingProviderFieldsCriticality,omitempty" tf:"finding_provider_fields_criticality"`

	FindingProviderFieldsRelatedFindingsId []FindingProviderFieldsRelatedFindingsIdParameters `json:"findingProviderFieldsRelatedFindingsId,omitempty" tf:"finding_provider_fields_related_findings_id"`

	FindingProviderFieldsRelatedFindingsProductArn []FindingProviderFieldsRelatedFindingsProductArnParameters `json:"findingProviderFieldsRelatedFindingsProductArn,omitempty" tf:"finding_provider_fields_related_findings_product_arn"`

	FindingProviderFieldsSeverityLabel []FindingProviderFieldsSeverityLabelParameters `json:"findingProviderFieldsSeverityLabel,omitempty" tf:"finding_provider_fields_severity_label"`

	FindingProviderFieldsSeverityOriginal []FindingProviderFieldsSeverityOriginalParameters `json:"findingProviderFieldsSeverityOriginal,omitempty" tf:"finding_provider_fields_severity_original"`

	FindingProviderFieldsTypes []FindingProviderFieldsTypesParameters `json:"findingProviderFieldsTypes,omitempty" tf:"finding_provider_fields_types"`

	FirstObservedAt []FirstObservedAtParameters `json:"firstObservedAt,omitempty" tf:"first_observed_at"`

	GeneratorId []GeneratorIdParameters `json:"generatorId,omitempty" tf:"generator_id"`

	Id []IdParameters `json:"id,omitempty" tf:"id"`

	Keyword []KeywordParameters `json:"keyword,omitempty" tf:"keyword"`

	LastObservedAt []LastObservedAtParameters `json:"lastObservedAt,omitempty" tf:"last_observed_at"`

	MalwareName []MalwareNameParameters `json:"malwareName,omitempty" tf:"malware_name"`

	MalwarePath []MalwarePathParameters `json:"malwarePath,omitempty" tf:"malware_path"`

	MalwareState []MalwareStateParameters `json:"malwareState,omitempty" tf:"malware_state"`

	MalwareType []MalwareTypeParameters `json:"malwareType,omitempty" tf:"malware_type"`

	NetworkDestinationDomain []NetworkDestinationDomainParameters `json:"networkDestinationDomain,omitempty" tf:"network_destination_domain"`

	NetworkDestinationIpv4 []NetworkDestinationIpv4Parameters `json:"networkDestinationIpv4,omitempty" tf:"network_destination_ipv4"`

	NetworkDestinationIpv6 []NetworkDestinationIpv6Parameters `json:"networkDestinationIpv6,omitempty" tf:"network_destination_ipv6"`

	NetworkDestinationPort []NetworkDestinationPortParameters `json:"networkDestinationPort,omitempty" tf:"network_destination_port"`

	NetworkDirection []NetworkDirectionParameters `json:"networkDirection,omitempty" tf:"network_direction"`

	NetworkProtocol []NetworkProtocolParameters `json:"networkProtocol,omitempty" tf:"network_protocol"`

	NetworkSourceDomain []NetworkSourceDomainParameters `json:"networkSourceDomain,omitempty" tf:"network_source_domain"`

	NetworkSourceIpv4 []NetworkSourceIpv4Parameters `json:"networkSourceIpv4,omitempty" tf:"network_source_ipv4"`

	NetworkSourceIpv6 []NetworkSourceIpv6Parameters `json:"networkSourceIpv6,omitempty" tf:"network_source_ipv6"`

	NetworkSourceMac []NetworkSourceMacParameters `json:"networkSourceMac,omitempty" tf:"network_source_mac"`

	NetworkSourcePort []NetworkSourcePortParameters `json:"networkSourcePort,omitempty" tf:"network_source_port"`

	NoteText []NoteTextParameters `json:"noteText,omitempty" tf:"note_text"`

	NoteUpdatedAt []NoteUpdatedAtParameters `json:"noteUpdatedAt,omitempty" tf:"note_updated_at"`

	NoteUpdatedBy []NoteUpdatedByParameters `json:"noteUpdatedBy,omitempty" tf:"note_updated_by"`

	ProcessLaunchedAt []ProcessLaunchedAtParameters `json:"processLaunchedAt,omitempty" tf:"process_launched_at"`

	ProcessName []ProcessNameParameters `json:"processName,omitempty" tf:"process_name"`

	ProcessParentPid []ProcessParentPidParameters `json:"processParentPid,omitempty" tf:"process_parent_pid"`

	ProcessPath []ProcessPathParameters `json:"processPath,omitempty" tf:"process_path"`

	ProcessPid []ProcessPidParameters `json:"processPid,omitempty" tf:"process_pid"`

	ProcessTerminatedAt []ProcessTerminatedAtParameters `json:"processTerminatedAt,omitempty" tf:"process_terminated_at"`

	ProductArn []ProductArnParameters `json:"productArn,omitempty" tf:"product_arn"`

	ProductFields []ProductFieldsParameters `json:"productFields,omitempty" tf:"product_fields"`

	ProductName []ProductNameParameters `json:"productName,omitempty" tf:"product_name"`

	RecommendationText []RecommendationTextParameters `json:"recommendationText,omitempty" tf:"recommendation_text"`

	RecordState []RecordStateParameters `json:"recordState,omitempty" tf:"record_state"`

	RelatedFindingsId []RelatedFindingsIdParameters `json:"relatedFindingsId,omitempty" tf:"related_findings_id"`

	RelatedFindingsProductArn []RelatedFindingsProductArnParameters `json:"relatedFindingsProductArn,omitempty" tf:"related_findings_product_arn"`

	ResourceAwsEc2InstanceIamInstanceProfileArn []ResourceAwsEc2InstanceIamInstanceProfileArnParameters `json:"resourceAwsEc2InstanceIamInstanceProfileArn,omitempty" tf:"resource_aws_ec2_instance_iam_instance_profile_arn"`

	ResourceAwsEc2InstanceImageId []ResourceAwsEc2InstanceImageIdParameters `json:"resourceAwsEc2InstanceImageId,omitempty" tf:"resource_aws_ec2_instance_image_id"`

	ResourceAwsEc2InstanceIpv4Addresses []ResourceAwsEc2InstanceIpv4AddressesParameters `json:"resourceAwsEc2InstanceIpv4Addresses,omitempty" tf:"resource_aws_ec2_instance_ipv4_addresses"`

	ResourceAwsEc2InstanceIpv6Addresses []ResourceAwsEc2InstanceIpv6AddressesParameters `json:"resourceAwsEc2InstanceIpv6Addresses,omitempty" tf:"resource_aws_ec2_instance_ipv6_addresses"`

	ResourceAwsEc2InstanceKeyName []ResourceAwsEc2InstanceKeyNameParameters `json:"resourceAwsEc2InstanceKeyName,omitempty" tf:"resource_aws_ec2_instance_key_name"`

	ResourceAwsEc2InstanceLaunchedAt []ResourceAwsEc2InstanceLaunchedAtParameters `json:"resourceAwsEc2InstanceLaunchedAt,omitempty" tf:"resource_aws_ec2_instance_launched_at"`

	ResourceAwsEc2InstanceSubnetId []ResourceAwsEc2InstanceSubnetIdParameters `json:"resourceAwsEc2InstanceSubnetId,omitempty" tf:"resource_aws_ec2_instance_subnet_id"`

	ResourceAwsEc2InstanceType []ResourceAwsEc2InstanceTypeParameters `json:"resourceAwsEc2InstanceType,omitempty" tf:"resource_aws_ec2_instance_type"`

	ResourceAwsEc2InstanceVpcId []ResourceAwsEc2InstanceVpcIdParameters `json:"resourceAwsEc2InstanceVpcId,omitempty" tf:"resource_aws_ec2_instance_vpc_id"`

	ResourceAwsIamAccessKeyCreatedAt []ResourceAwsIamAccessKeyCreatedAtParameters `json:"resourceAwsIamAccessKeyCreatedAt,omitempty" tf:"resource_aws_iam_access_key_created_at"`

	ResourceAwsIamAccessKeyStatus []ResourceAwsIamAccessKeyStatusParameters `json:"resourceAwsIamAccessKeyStatus,omitempty" tf:"resource_aws_iam_access_key_status"`

	ResourceAwsIamAccessKeyUserName []ResourceAwsIamAccessKeyUserNameParameters `json:"resourceAwsIamAccessKeyUserName,omitempty" tf:"resource_aws_iam_access_key_user_name"`

	ResourceAwsS3BucketOwnerId []ResourceAwsS3BucketOwnerIdParameters `json:"resourceAwsS3BucketOwnerId,omitempty" tf:"resource_aws_s3_bucket_owner_id"`

	ResourceAwsS3BucketOwnerName []ResourceAwsS3BucketOwnerNameParameters `json:"resourceAwsS3BucketOwnerName,omitempty" tf:"resource_aws_s3_bucket_owner_name"`

	ResourceContainerImageId []ResourceContainerImageIdParameters `json:"resourceContainerImageId,omitempty" tf:"resource_container_image_id"`

	ResourceContainerImageName []ResourceContainerImageNameParameters `json:"resourceContainerImageName,omitempty" tf:"resource_container_image_name"`

	ResourceContainerLaunchedAt []ResourceContainerLaunchedAtParameters `json:"resourceContainerLaunchedAt,omitempty" tf:"resource_container_launched_at"`

	ResourceContainerName []ResourceContainerNameParameters `json:"resourceContainerName,omitempty" tf:"resource_container_name"`

	ResourceDetailsOther []ResourceDetailsOtherParameters `json:"resourceDetailsOther,omitempty" tf:"resource_details_other"`

	ResourceId []ResourceIdParameters `json:"resourceId,omitempty" tf:"resource_id"`

	ResourcePartition []ResourcePartitionParameters `json:"resourcePartition,omitempty" tf:"resource_partition"`

	ResourceRegion []ResourceRegionParameters `json:"resourceRegion,omitempty" tf:"resource_region"`

	ResourceTags []ResourceTagsParameters `json:"resourceTags,omitempty" tf:"resource_tags"`

	ResourceType []ResourceTypeParameters `json:"resourceType,omitempty" tf:"resource_type"`

	SeverityLabel []SeverityLabelParameters `json:"severityLabel,omitempty" tf:"severity_label"`

	SourceUrl []SourceUrlParameters `json:"sourceUrl,omitempty" tf:"source_url"`

	ThreatIntelIndicatorCategory []ThreatIntelIndicatorCategoryParameters `json:"threatIntelIndicatorCategory,omitempty" tf:"threat_intel_indicator_category"`

	ThreatIntelIndicatorLastObservedAt []ThreatIntelIndicatorLastObservedAtParameters `json:"threatIntelIndicatorLastObservedAt,omitempty" tf:"threat_intel_indicator_last_observed_at"`

	ThreatIntelIndicatorSource []ThreatIntelIndicatorSourceParameters `json:"threatIntelIndicatorSource,omitempty" tf:"threat_intel_indicator_source"`

	ThreatIntelIndicatorSourceUrl []ThreatIntelIndicatorSourceUrlParameters `json:"threatIntelIndicatorSourceUrl,omitempty" tf:"threat_intel_indicator_source_url"`

	ThreatIntelIndicatorType []ThreatIntelIndicatorTypeParameters `json:"threatIntelIndicatorType,omitempty" tf:"threat_intel_indicator_type"`

	ThreatIntelIndicatorValue []ThreatIntelIndicatorValueParameters `json:"threatIntelIndicatorValue,omitempty" tf:"threat_intel_indicator_value"`

	Title []TitleParameters `json:"title,omitempty" tf:"title"`

	Type []TypeParameters `json:"type,omitempty" tf:"type"`

	UpdatedAt []UpdatedAtParameters `json:"updatedAt,omitempty" tf:"updated_at"`

	UserDefinedValues []UserDefinedValuesParameters `json:"userDefinedValues,omitempty" tf:"user_defined_values"`

	VerificationState []VerificationStateParameters `json:"verificationState,omitempty" tf:"verification_state"`

	WorkflowStatus []WorkflowStatusParameters `json:"workflowStatus,omitempty" tf:"workflow_status"`
}

type FindingProviderFieldsConfidenceObservation struct {
}

type FindingProviderFieldsConfidenceParameters struct {
	Eq *string `json:"eq,omitempty" tf:"eq"`

	Gte *string `json:"gte,omitempty" tf:"gte"`

	Lte *string `json:"lte,omitempty" tf:"lte"`
}

type FindingProviderFieldsCriticalityObservation struct {
}

type FindingProviderFieldsCriticalityParameters struct {
	Eq *string `json:"eq,omitempty" tf:"eq"`

	Gte *string `json:"gte,omitempty" tf:"gte"`

	Lte *string `json:"lte,omitempty" tf:"lte"`
}

type FindingProviderFieldsRelatedFindingsIdObservation struct {
}

type FindingProviderFieldsRelatedFindingsIdParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type FindingProviderFieldsRelatedFindingsProductArnObservation struct {
}

type FindingProviderFieldsRelatedFindingsProductArnParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type FindingProviderFieldsSeverityLabelObservation struct {
}

type FindingProviderFieldsSeverityLabelParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type FindingProviderFieldsSeverityOriginalObservation struct {
}

type FindingProviderFieldsSeverityOriginalParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type FindingProviderFieldsTypesObservation struct {
}

type FindingProviderFieldsTypesParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type FirstObservedAtDateRangeObservation struct {
}

type FirstObservedAtDateRangeParameters struct {
	Unit string `json:"unit" tf:"unit"`

	Value int64 `json:"value" tf:"value"`
}

type FirstObservedAtObservation struct {
}

type FirstObservedAtParameters struct {
	DateRange []FirstObservedAtDateRangeParameters `json:"dateRange,omitempty" tf:"date_range"`

	End *string `json:"end,omitempty" tf:"end"`

	Start *string `json:"start,omitempty" tf:"start"`
}

type GeneratorIdObservation struct {
}

type GeneratorIdParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type IdObservation struct {
}

type IdParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type KeywordObservation struct {
}

type KeywordParameters struct {
	Value string `json:"value" tf:"value"`
}

type LastObservedAtDateRangeObservation struct {
}

type LastObservedAtDateRangeParameters struct {
	Unit string `json:"unit" tf:"unit"`

	Value int64 `json:"value" tf:"value"`
}

type LastObservedAtObservation struct {
}

type LastObservedAtParameters struct {
	DateRange []LastObservedAtDateRangeParameters `json:"dateRange,omitempty" tf:"date_range"`

	End *string `json:"end,omitempty" tf:"end"`

	Start *string `json:"start,omitempty" tf:"start"`
}

type MalwareNameObservation struct {
}

type MalwareNameParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type MalwarePathObservation struct {
}

type MalwarePathParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type MalwareStateObservation struct {
}

type MalwareStateParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type MalwareTypeObservation struct {
}

type MalwareTypeParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type NetworkDestinationDomainObservation struct {
}

type NetworkDestinationDomainParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type NetworkDestinationIpv4Observation struct {
}

type NetworkDestinationIpv4Parameters struct {
	Cidr string `json:"cidr" tf:"cidr"`
}

type NetworkDestinationIpv6Observation struct {
}

type NetworkDestinationIpv6Parameters struct {
	Cidr string `json:"cidr" tf:"cidr"`
}

type NetworkDestinationPortObservation struct {
}

type NetworkDestinationPortParameters struct {
	Eq *string `json:"eq,omitempty" tf:"eq"`

	Gte *string `json:"gte,omitempty" tf:"gte"`

	Lte *string `json:"lte,omitempty" tf:"lte"`
}

type NetworkDirectionObservation struct {
}

type NetworkDirectionParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type NetworkProtocolObservation struct {
}

type NetworkProtocolParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type NetworkSourceDomainObservation struct {
}

type NetworkSourceDomainParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type NetworkSourceIpv4Observation struct {
}

type NetworkSourceIpv4Parameters struct {
	Cidr string `json:"cidr" tf:"cidr"`
}

type NetworkSourceIpv6Observation struct {
}

type NetworkSourceIpv6Parameters struct {
	Cidr string `json:"cidr" tf:"cidr"`
}

type NetworkSourceMacObservation struct {
}

type NetworkSourceMacParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type NetworkSourcePortObservation struct {
}

type NetworkSourcePortParameters struct {
	Eq *string `json:"eq,omitempty" tf:"eq"`

	Gte *string `json:"gte,omitempty" tf:"gte"`

	Lte *string `json:"lte,omitempty" tf:"lte"`
}

type NoteTextObservation struct {
}

type NoteTextParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type NoteUpdatedAtDateRangeObservation struct {
}

type NoteUpdatedAtDateRangeParameters struct {
	Unit string `json:"unit" tf:"unit"`

	Value int64 `json:"value" tf:"value"`
}

type NoteUpdatedAtObservation struct {
}

type NoteUpdatedAtParameters struct {
	DateRange []NoteUpdatedAtDateRangeParameters `json:"dateRange,omitempty" tf:"date_range"`

	End *string `json:"end,omitempty" tf:"end"`

	Start *string `json:"start,omitempty" tf:"start"`
}

type NoteUpdatedByObservation struct {
}

type NoteUpdatedByParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type ProcessLaunchedAtDateRangeObservation struct {
}

type ProcessLaunchedAtDateRangeParameters struct {
	Unit string `json:"unit" tf:"unit"`

	Value int64 `json:"value" tf:"value"`
}

type ProcessLaunchedAtObservation struct {
}

type ProcessLaunchedAtParameters struct {
	DateRange []ProcessLaunchedAtDateRangeParameters `json:"dateRange,omitempty" tf:"date_range"`

	End *string `json:"end,omitempty" tf:"end"`

	Start *string `json:"start,omitempty" tf:"start"`
}

type ProcessNameObservation struct {
}

type ProcessNameParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type ProcessParentPidObservation struct {
}

type ProcessParentPidParameters struct {
	Eq *string `json:"eq,omitempty" tf:"eq"`

	Gte *string `json:"gte,omitempty" tf:"gte"`

	Lte *string `json:"lte,omitempty" tf:"lte"`
}

type ProcessPathObservation struct {
}

type ProcessPathParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type ProcessPidObservation struct {
}

type ProcessPidParameters struct {
	Eq *string `json:"eq,omitempty" tf:"eq"`

	Gte *string `json:"gte,omitempty" tf:"gte"`

	Lte *string `json:"lte,omitempty" tf:"lte"`
}

type ProcessTerminatedAtDateRangeObservation struct {
}

type ProcessTerminatedAtDateRangeParameters struct {
	Unit string `json:"unit" tf:"unit"`

	Value int64 `json:"value" tf:"value"`
}

type ProcessTerminatedAtObservation struct {
}

type ProcessTerminatedAtParameters struct {
	DateRange []ProcessTerminatedAtDateRangeParameters `json:"dateRange,omitempty" tf:"date_range"`

	End *string `json:"end,omitempty" tf:"end"`

	Start *string `json:"start,omitempty" tf:"start"`
}

type ProductArnObservation struct {
}

type ProductArnParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type ProductFieldsObservation struct {
}

type ProductFieldsParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Key string `json:"key" tf:"key"`

	Value string `json:"value" tf:"value"`
}

type ProductNameObservation struct {
}

type ProductNameParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type RecommendationTextObservation struct {
}

type RecommendationTextParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type RecordStateObservation struct {
}

type RecordStateParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type RelatedFindingsIdObservation struct {
}

type RelatedFindingsIdParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type RelatedFindingsProductArnObservation struct {
}

type RelatedFindingsProductArnParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type ResourceAwsEc2InstanceIamInstanceProfileArnObservation struct {
}

type ResourceAwsEc2InstanceIamInstanceProfileArnParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type ResourceAwsEc2InstanceImageIdObservation struct {
}

type ResourceAwsEc2InstanceImageIdParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type ResourceAwsEc2InstanceIpv4AddressesObservation struct {
}

type ResourceAwsEc2InstanceIpv4AddressesParameters struct {
	Cidr string `json:"cidr" tf:"cidr"`
}

type ResourceAwsEc2InstanceIpv6AddressesObservation struct {
}

type ResourceAwsEc2InstanceIpv6AddressesParameters struct {
	Cidr string `json:"cidr" tf:"cidr"`
}

type ResourceAwsEc2InstanceKeyNameObservation struct {
}

type ResourceAwsEc2InstanceKeyNameParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type ResourceAwsEc2InstanceLaunchedAtDateRangeObservation struct {
}

type ResourceAwsEc2InstanceLaunchedAtDateRangeParameters struct {
	Unit string `json:"unit" tf:"unit"`

	Value int64 `json:"value" tf:"value"`
}

type ResourceAwsEc2InstanceLaunchedAtObservation struct {
}

type ResourceAwsEc2InstanceLaunchedAtParameters struct {
	DateRange []ResourceAwsEc2InstanceLaunchedAtDateRangeParameters `json:"dateRange,omitempty" tf:"date_range"`

	End *string `json:"end,omitempty" tf:"end"`

	Start *string `json:"start,omitempty" tf:"start"`
}

type ResourceAwsEc2InstanceSubnetIdObservation struct {
}

type ResourceAwsEc2InstanceSubnetIdParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type ResourceAwsEc2InstanceTypeObservation struct {
}

type ResourceAwsEc2InstanceTypeParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type ResourceAwsEc2InstanceVpcIdObservation struct {
}

type ResourceAwsEc2InstanceVpcIdParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type ResourceAwsIamAccessKeyCreatedAtDateRangeObservation struct {
}

type ResourceAwsIamAccessKeyCreatedAtDateRangeParameters struct {
	Unit string `json:"unit" tf:"unit"`

	Value int64 `json:"value" tf:"value"`
}

type ResourceAwsIamAccessKeyCreatedAtObservation struct {
}

type ResourceAwsIamAccessKeyCreatedAtParameters struct {
	DateRange []ResourceAwsIamAccessKeyCreatedAtDateRangeParameters `json:"dateRange,omitempty" tf:"date_range"`

	End *string `json:"end,omitempty" tf:"end"`

	Start *string `json:"start,omitempty" tf:"start"`
}

type ResourceAwsIamAccessKeyStatusObservation struct {
}

type ResourceAwsIamAccessKeyStatusParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type ResourceAwsIamAccessKeyUserNameObservation struct {
}

type ResourceAwsIamAccessKeyUserNameParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type ResourceAwsS3BucketOwnerIdObservation struct {
}

type ResourceAwsS3BucketOwnerIdParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type ResourceAwsS3BucketOwnerNameObservation struct {
}

type ResourceAwsS3BucketOwnerNameParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type ResourceContainerImageIdObservation struct {
}

type ResourceContainerImageIdParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type ResourceContainerImageNameObservation struct {
}

type ResourceContainerImageNameParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type ResourceContainerLaunchedAtDateRangeObservation struct {
}

type ResourceContainerLaunchedAtDateRangeParameters struct {
	Unit string `json:"unit" tf:"unit"`

	Value int64 `json:"value" tf:"value"`
}

type ResourceContainerLaunchedAtObservation struct {
}

type ResourceContainerLaunchedAtParameters struct {
	DateRange []ResourceContainerLaunchedAtDateRangeParameters `json:"dateRange,omitempty" tf:"date_range"`

	End *string `json:"end,omitempty" tf:"end"`

	Start *string `json:"start,omitempty" tf:"start"`
}

type ResourceContainerNameObservation struct {
}

type ResourceContainerNameParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type ResourceDetailsOtherObservation struct {
}

type ResourceDetailsOtherParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Key string `json:"key" tf:"key"`

	Value string `json:"value" tf:"value"`
}

type ResourceIdObservation struct {
}

type ResourceIdParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type ResourcePartitionObservation struct {
}

type ResourcePartitionParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type ResourceRegionObservation struct {
}

type ResourceRegionParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type ResourceTagsObservation struct {
}

type ResourceTagsParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Key string `json:"key" tf:"key"`

	Value string `json:"value" tf:"value"`
}

type ResourceTypeObservation struct {
}

type ResourceTypeParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type SecurityhubInsightObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type SecurityhubInsightParameters struct {
	Filters []FiltersParameters `json:"filters" tf:"filters"`

	GroupByAttribute string `json:"groupByAttribute" tf:"group_by_attribute"`

	Name string `json:"name" tf:"name"`
}

type SeverityLabelObservation struct {
}

type SeverityLabelParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type SourceUrlObservation struct {
}

type SourceUrlParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type ThreatIntelIndicatorCategoryObservation struct {
}

type ThreatIntelIndicatorCategoryParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type ThreatIntelIndicatorLastObservedAtDateRangeObservation struct {
}

type ThreatIntelIndicatorLastObservedAtDateRangeParameters struct {
	Unit string `json:"unit" tf:"unit"`

	Value int64 `json:"value" tf:"value"`
}

type ThreatIntelIndicatorLastObservedAtObservation struct {
}

type ThreatIntelIndicatorLastObservedAtParameters struct {
	DateRange []ThreatIntelIndicatorLastObservedAtDateRangeParameters `json:"dateRange,omitempty" tf:"date_range"`

	End *string `json:"end,omitempty" tf:"end"`

	Start *string `json:"start,omitempty" tf:"start"`
}

type ThreatIntelIndicatorSourceObservation struct {
}

type ThreatIntelIndicatorSourceParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type ThreatIntelIndicatorSourceUrlObservation struct {
}

type ThreatIntelIndicatorSourceUrlParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type ThreatIntelIndicatorTypeObservation struct {
}

type ThreatIntelIndicatorTypeParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type ThreatIntelIndicatorValueObservation struct {
}

type ThreatIntelIndicatorValueParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type TitleObservation struct {
}

type TitleParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type TypeObservation struct {
}

type TypeParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type UpdatedAtDateRangeObservation struct {
}

type UpdatedAtDateRangeParameters struct {
	Unit string `json:"unit" tf:"unit"`

	Value int64 `json:"value" tf:"value"`
}

type UpdatedAtObservation struct {
}

type UpdatedAtParameters struct {
	DateRange []UpdatedAtDateRangeParameters `json:"dateRange,omitempty" tf:"date_range"`

	End *string `json:"end,omitempty" tf:"end"`

	Start *string `json:"start,omitempty" tf:"start"`
}

type UserDefinedValuesObservation struct {
}

type UserDefinedValuesParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Key string `json:"key" tf:"key"`

	Value string `json:"value" tf:"value"`
}

type VerificationStateObservation struct {
}

type VerificationStateParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

type WorkflowStatusObservation struct {
}

type WorkflowStatusParameters struct {
	Comparison string `json:"comparison" tf:"comparison"`

	Value string `json:"value" tf:"value"`
}

// SecurityhubInsightSpec defines the desired state of SecurityhubInsight
type SecurityhubInsightSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SecurityhubInsightParameters `json:"forProvider"`
}

// SecurityhubInsightStatus defines the observed state of SecurityhubInsight.
type SecurityhubInsightStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SecurityhubInsightObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityhubInsight is the Schema for the SecurityhubInsights API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type SecurityhubInsight struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecurityhubInsightSpec   `json:"spec"`
	Status            SecurityhubInsightStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityhubInsightList contains a list of SecurityhubInsights
type SecurityhubInsightList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityhubInsight `json:"items"`
}

// Repository type metadata.
var (
	SecurityhubInsightKind             = "SecurityhubInsight"
	SecurityhubInsightGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: SecurityhubInsightKind}.String()
	SecurityhubInsightKindAPIVersion   = SecurityhubInsightKind + "." + v1alpha1.GroupVersion.String()
	SecurityhubInsightGroupVersionKind = v1alpha1.GroupVersion.WithKind(SecurityhubInsightKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &SecurityhubInsight{}, &SecurityhubInsightList{})
}
