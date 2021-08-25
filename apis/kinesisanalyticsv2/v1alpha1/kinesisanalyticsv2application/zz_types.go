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
// +groupName=kinesisanalyticsv2.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	apis "github.com/crossplane-contrib/provider-tf-aws/apis"
	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/kinesisanalyticsv2/v1alpha1"
)

type ApplicationCodeConfigurationObservation struct {
}

type ApplicationCodeConfigurationParameters struct {
	CodeContent []CodeContentParameters `json:"codeContent,omitempty" tf:"code_content"`

	CodeContentType string `json:"codeContentType" tf:"code_content_type"`
}

type ApplicationConfigurationObservation struct {
}

type ApplicationConfigurationParameters struct {
	ApplicationCodeConfiguration []ApplicationCodeConfigurationParameters `json:"applicationCodeConfiguration" tf:"application_code_configuration"`

	ApplicationSnapshotConfiguration []ApplicationSnapshotConfigurationParameters `json:"applicationSnapshotConfiguration,omitempty" tf:"application_snapshot_configuration"`

	EnvironmentProperties []EnvironmentPropertiesParameters `json:"environmentProperties,omitempty" tf:"environment_properties"`

	FlinkApplicationConfiguration []FlinkApplicationConfigurationParameters `json:"flinkApplicationConfiguration,omitempty" tf:"flink_application_configuration"`

	RunConfiguration []RunConfigurationParameters `json:"runConfiguration,omitempty" tf:"run_configuration"`

	SqlApplicationConfiguration []SqlApplicationConfigurationParameters `json:"sqlApplicationConfiguration,omitempty" tf:"sql_application_configuration"`

	VpcConfiguration []VpcConfigurationParameters `json:"vpcConfiguration,omitempty" tf:"vpc_configuration"`
}

type ApplicationRestoreConfigurationObservation struct {
}

type ApplicationRestoreConfigurationParameters struct {
	ApplicationRestoreType *string `json:"applicationRestoreType,omitempty" tf:"application_restore_type"`

	SnapshotName *string `json:"snapshotName,omitempty" tf:"snapshot_name"`
}

type ApplicationSnapshotConfigurationObservation struct {
}

type ApplicationSnapshotConfigurationParameters struct {
	SnapshotsEnabled bool `json:"snapshotsEnabled" tf:"snapshots_enabled"`
}

type CheckpointConfigurationObservation struct {
}

type CheckpointConfigurationParameters struct {
	CheckpointInterval *int64 `json:"checkpointInterval,omitempty" tf:"checkpoint_interval"`

	CheckpointingEnabled *bool `json:"checkpointingEnabled,omitempty" tf:"checkpointing_enabled"`

	ConfigurationType string `json:"configurationType" tf:"configuration_type"`

	MinPauseBetweenCheckpoints *int64 `json:"minPauseBetweenCheckpoints,omitempty" tf:"min_pause_between_checkpoints"`
}

type CloudwatchLoggingOptionsObservation struct {
	CloudwatchLoggingOptionId string `json:"cloudwatchLoggingOptionId" tf:"cloudwatch_logging_option_id"`
}

type CloudwatchLoggingOptionsParameters struct {
	LogStreamArn string `json:"logStreamArn" tf:"log_stream_arn"`
}

type CodeContentObservation struct {
}

type CodeContentParameters struct {
	S3ContentLocation []S3ContentLocationParameters `json:"s3ContentLocation,omitempty" tf:"s3_content_location"`

	TextContent *string `json:"textContent,omitempty" tf:"text_content"`
}

type CsvMappingParametersObservation struct {
}

type CsvMappingParametersParameters struct {
	RecordColumnDelimiter string `json:"recordColumnDelimiter" tf:"record_column_delimiter"`

	RecordRowDelimiter string `json:"recordRowDelimiter" tf:"record_row_delimiter"`
}

type DestinationSchemaObservation struct {
}

type DestinationSchemaParameters struct {
	RecordFormatType string `json:"recordFormatType" tf:"record_format_type"`
}

type EnvironmentPropertiesObservation struct {
}

type EnvironmentPropertiesParameters struct {
	PropertyGroup []PropertyGroupParameters `json:"propertyGroup" tf:"property_group"`
}

type FlinkApplicationConfigurationObservation struct {
}

type FlinkApplicationConfigurationParameters struct {
	CheckpointConfiguration []CheckpointConfigurationParameters `json:"checkpointConfiguration,omitempty" tf:"checkpoint_configuration"`

	MonitoringConfiguration []MonitoringConfigurationParameters `json:"monitoringConfiguration,omitempty" tf:"monitoring_configuration"`

	ParallelismConfiguration []ParallelismConfigurationParameters `json:"parallelismConfiguration,omitempty" tf:"parallelism_configuration"`
}

type FlinkRunConfigurationObservation struct {
}

type FlinkRunConfigurationParameters struct {
	AllowNonRestoredState *bool `json:"allowNonRestoredState,omitempty" tf:"allow_non_restored_state"`
}

type InputLambdaProcessorObservation struct {
}

type InputLambdaProcessorParameters struct {
	ResourceArn string `json:"resourceArn" tf:"resource_arn"`
}

type InputObservation struct {
	InAppStreamNames []string `json:"inAppStreamNames" tf:"in_app_stream_names"`

	InputId string `json:"inputId" tf:"input_id"`
}

type InputParallelismObservation struct {
}

type InputParallelismParameters struct {
	Count *int64 `json:"count,omitempty" tf:"count"`
}

type InputParameters struct {
	InputParallelism []InputParallelismParameters `json:"inputParallelism,omitempty" tf:"input_parallelism"`

	InputProcessingConfiguration []InputProcessingConfigurationParameters `json:"inputProcessingConfiguration,omitempty" tf:"input_processing_configuration"`

	InputSchema []InputSchemaParameters `json:"inputSchema" tf:"input_schema"`

	InputStartingPositionConfiguration []InputStartingPositionConfigurationParameters `json:"inputStartingPositionConfiguration,omitempty" tf:"input_starting_position_configuration"`

	KinesisFirehoseInput []KinesisFirehoseInputParameters `json:"kinesisFirehoseInput,omitempty" tf:"kinesis_firehose_input"`

	KinesisStreamsInput []KinesisStreamsInputParameters `json:"kinesisStreamsInput,omitempty" tf:"kinesis_streams_input"`

	NamePrefix string `json:"namePrefix" tf:"name_prefix"`
}

type InputProcessingConfigurationObservation struct {
}

type InputProcessingConfigurationParameters struct {
	InputLambdaProcessor []InputLambdaProcessorParameters `json:"inputLambdaProcessor" tf:"input_lambda_processor"`
}

type InputSchemaObservation struct {
}

type InputSchemaParameters struct {
	RecordColumn []RecordColumnParameters `json:"recordColumn" tf:"record_column"`

	RecordEncoding *string `json:"recordEncoding,omitempty" tf:"record_encoding"`

	RecordFormat []RecordFormatParameters `json:"recordFormat" tf:"record_format"`
}

type InputStartingPositionConfigurationObservation struct {
}

type InputStartingPositionConfigurationParameters struct {
	InputStartingPosition *string `json:"inputStartingPosition,omitempty" tf:"input_starting_position"`
}

type JsonMappingParametersObservation struct {
}

type JsonMappingParametersParameters struct {
	RecordRowPath string `json:"recordRowPath" tf:"record_row_path"`
}

type KinesisFirehoseInputObservation struct {
}

type KinesisFirehoseInputParameters struct {
	ResourceArn string `json:"resourceArn" tf:"resource_arn"`
}

type KinesisFirehoseOutputObservation struct {
}

type KinesisFirehoseOutputParameters struct {
	ResourceArn string `json:"resourceArn" tf:"resource_arn"`
}

type KinesisStreamsInputObservation struct {
}

type KinesisStreamsInputParameters struct {
	ResourceArn string `json:"resourceArn" tf:"resource_arn"`
}

type KinesisStreamsOutputObservation struct {
}

type KinesisStreamsOutputParameters struct {
	ResourceArn string `json:"resourceArn" tf:"resource_arn"`
}

type Kinesisanalyticsv2ApplicationObservation struct {
	Arn string `json:"arn" tf:"arn"`

	CreateTimestamp string `json:"createTimestamp" tf:"create_timestamp"`

	LastUpdateTimestamp string `json:"lastUpdateTimestamp" tf:"last_update_timestamp"`

	Status string `json:"status" tf:"status"`

	VersionId int64 `json:"versionId" tf:"version_id"`
}

type Kinesisanalyticsv2ApplicationParameters struct {
	ApplicationConfiguration []ApplicationConfigurationParameters `json:"applicationConfiguration,omitempty" tf:"application_configuration"`

	CloudwatchLoggingOptions []CloudwatchLoggingOptionsParameters `json:"cloudwatchLoggingOptions,omitempty" tf:"cloudwatch_logging_options"`

	Description *string `json:"description,omitempty" tf:"description"`

	ForceStop *bool `json:"forceStop,omitempty" tf:"force_stop"`

	Name string `json:"name" tf:"name"`

	RuntimeEnvironment string `json:"runtimeEnvironment" tf:"runtime_environment"`

	ServiceExecutionRole string `json:"serviceExecutionRole" tf:"service_execution_role"`

	StartApplication *bool `json:"startApplication,omitempty" tf:"start_application"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type LambdaOutputObservation struct {
}

type LambdaOutputParameters struct {
	ResourceArn string `json:"resourceArn" tf:"resource_arn"`
}

type MappingParametersCsvMappingParametersObservation struct {
}

type MappingParametersCsvMappingParametersParameters struct {
	RecordColumnDelimiter string `json:"recordColumnDelimiter" tf:"record_column_delimiter"`

	RecordRowDelimiter string `json:"recordRowDelimiter" tf:"record_row_delimiter"`
}

type MappingParametersJsonMappingParametersObservation struct {
}

type MappingParametersJsonMappingParametersParameters struct {
	RecordRowPath string `json:"recordRowPath" tf:"record_row_path"`
}

type MappingParametersObservation struct {
}

type MappingParametersParameters struct {
	CsvMappingParameters []CsvMappingParametersParameters `json:"csvMappingParameters,omitempty" tf:"csv_mapping_parameters"`

	JsonMappingParameters []JsonMappingParametersParameters `json:"jsonMappingParameters,omitempty" tf:"json_mapping_parameters"`
}

type MonitoringConfigurationObservation struct {
}

type MonitoringConfigurationParameters struct {
	ConfigurationType string `json:"configurationType" tf:"configuration_type"`

	LogLevel *string `json:"logLevel,omitempty" tf:"log_level"`

	MetricsLevel *string `json:"metricsLevel,omitempty" tf:"metrics_level"`
}

type OutputObservation struct {
	OutputId string `json:"outputId" tf:"output_id"`
}

type OutputParameters struct {
	DestinationSchema []DestinationSchemaParameters `json:"destinationSchema" tf:"destination_schema"`

	KinesisFirehoseOutput []KinesisFirehoseOutputParameters `json:"kinesisFirehoseOutput,omitempty" tf:"kinesis_firehose_output"`

	KinesisStreamsOutput []KinesisStreamsOutputParameters `json:"kinesisStreamsOutput,omitempty" tf:"kinesis_streams_output"`

	LambdaOutput []LambdaOutputParameters `json:"lambdaOutput,omitempty" tf:"lambda_output"`

	Name string `json:"name" tf:"name"`
}

type ParallelismConfigurationObservation struct {
}

type ParallelismConfigurationParameters struct {
	AutoScalingEnabled *bool `json:"autoScalingEnabled,omitempty" tf:"auto_scaling_enabled"`

	ConfigurationType string `json:"configurationType" tf:"configuration_type"`

	Parallelism *int64 `json:"parallelism,omitempty" tf:"parallelism"`

	ParallelismPerKpu *int64 `json:"parallelismPerKpu,omitempty" tf:"parallelism_per_kpu"`
}

type PropertyGroupObservation struct {
}

type PropertyGroupParameters struct {
	PropertyGroupId string `json:"propertyGroupId" tf:"property_group_id"`

	PropertyMap map[string]string `json:"propertyMap" tf:"property_map"`
}

type RecordColumnObservation struct {
}

type RecordColumnParameters struct {
	Mapping *string `json:"mapping,omitempty" tf:"mapping"`

	Name string `json:"name" tf:"name"`

	SqlType string `json:"sqlType" tf:"sql_type"`
}

type RecordFormatMappingParametersObservation struct {
}

type RecordFormatMappingParametersParameters struct {
	CsvMappingParameters []MappingParametersCsvMappingParametersParameters `json:"csvMappingParameters,omitempty" tf:"csv_mapping_parameters"`

	JsonMappingParameters []MappingParametersJsonMappingParametersParameters `json:"jsonMappingParameters,omitempty" tf:"json_mapping_parameters"`
}

type RecordFormatObservation struct {
}

type RecordFormatParameters struct {
	MappingParameters []MappingParametersParameters `json:"mappingParameters" tf:"mapping_parameters"`

	RecordFormatType string `json:"recordFormatType" tf:"record_format_type"`
}

type ReferenceDataSourceObservation struct {
	ReferenceId string `json:"referenceId" tf:"reference_id"`
}

type ReferenceDataSourceParameters struct {
	ReferenceSchema []ReferenceSchemaParameters `json:"referenceSchema" tf:"reference_schema"`

	S3ReferenceDataSource []S3ReferenceDataSourceParameters `json:"s3ReferenceDataSource" tf:"s3_reference_data_source"`

	TableName string `json:"tableName" tf:"table_name"`
}

type ReferenceSchemaObservation struct {
}

type ReferenceSchemaParameters struct {
	RecordColumn []ReferenceSchemaRecordColumnParameters `json:"recordColumn" tf:"record_column"`

	RecordEncoding *string `json:"recordEncoding,omitempty" tf:"record_encoding"`

	RecordFormat []ReferenceSchemaRecordFormatParameters `json:"recordFormat" tf:"record_format"`
}

type ReferenceSchemaRecordColumnObservation struct {
}

type ReferenceSchemaRecordColumnParameters struct {
	Mapping *string `json:"mapping,omitempty" tf:"mapping"`

	Name string `json:"name" tf:"name"`

	SqlType string `json:"sqlType" tf:"sql_type"`
}

type ReferenceSchemaRecordFormatObservation struct {
}

type ReferenceSchemaRecordFormatParameters struct {
	MappingParameters []RecordFormatMappingParametersParameters `json:"mappingParameters" tf:"mapping_parameters"`

	RecordFormatType string `json:"recordFormatType" tf:"record_format_type"`
}

type RunConfigurationObservation struct {
}

type RunConfigurationParameters struct {
	ApplicationRestoreConfiguration []ApplicationRestoreConfigurationParameters `json:"applicationRestoreConfiguration,omitempty" tf:"application_restore_configuration"`

	FlinkRunConfiguration []FlinkRunConfigurationParameters `json:"flinkRunConfiguration,omitempty" tf:"flink_run_configuration"`
}

type S3ContentLocationObservation struct {
}

type S3ContentLocationParameters struct {
	BucketArn string `json:"bucketArn" tf:"bucket_arn"`

	FileKey string `json:"fileKey" tf:"file_key"`

	ObjectVersion *string `json:"objectVersion,omitempty" tf:"object_version"`
}

type S3ReferenceDataSourceObservation struct {
}

type S3ReferenceDataSourceParameters struct {
	BucketArn string `json:"bucketArn" tf:"bucket_arn"`

	FileKey string `json:"fileKey" tf:"file_key"`
}

type SqlApplicationConfigurationObservation struct {
}

type SqlApplicationConfigurationParameters struct {
	Input []InputParameters `json:"input,omitempty" tf:"input"`

	Output []OutputParameters `json:"output,omitempty" tf:"output"`

	ReferenceDataSource []ReferenceDataSourceParameters `json:"referenceDataSource,omitempty" tf:"reference_data_source"`
}

type VpcConfigurationObservation struct {
	VpcConfigurationId string `json:"vpcConfigurationId" tf:"vpc_configuration_id"`

	VpcId string `json:"vpcId" tf:"vpc_id"`
}

type VpcConfigurationParameters struct {
	SecurityGroupIds []string `json:"securityGroupIds" tf:"security_group_ids"`

	SubnetIds []string `json:"subnetIds" tf:"subnet_ids"`
}

// Kinesisanalyticsv2ApplicationSpec defines the desired state of Kinesisanalyticsv2Application
type Kinesisanalyticsv2ApplicationSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       Kinesisanalyticsv2ApplicationParameters `json:"forProvider"`
}

// Kinesisanalyticsv2ApplicationStatus defines the observed state of Kinesisanalyticsv2Application.
type Kinesisanalyticsv2ApplicationStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          Kinesisanalyticsv2ApplicationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Kinesisanalyticsv2Application is the Schema for the Kinesisanalyticsv2Applications API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Kinesisanalyticsv2Application struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Kinesisanalyticsv2ApplicationSpec   `json:"spec"`
	Status            Kinesisanalyticsv2ApplicationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Kinesisanalyticsv2ApplicationList contains a list of Kinesisanalyticsv2Applications
type Kinesisanalyticsv2ApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Kinesisanalyticsv2Application `json:"items"`
}

// Repository type metadata.
var (
	Kinesisanalyticsv2ApplicationKind             = "Kinesisanalyticsv2Application"
	Kinesisanalyticsv2ApplicationGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: Kinesisanalyticsv2ApplicationKind}.String()
	Kinesisanalyticsv2ApplicationKindAPIVersion   = Kinesisanalyticsv2ApplicationKind + "." + v1alpha1.GroupVersion.String()
	Kinesisanalyticsv2ApplicationGroupVersionKind = v1alpha1.GroupVersion.WithKind(Kinesisanalyticsv2ApplicationKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &Kinesisanalyticsv2Application{}, &Kinesisanalyticsv2ApplicationList{})
}
