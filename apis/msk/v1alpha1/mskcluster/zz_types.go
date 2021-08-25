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
// +groupName=msk.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	apis "github.com/crossplane-contrib/provider-tf-aws/apis"
	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/msk/v1alpha1"
)

type BrokerLogsObservation struct {
}

type BrokerLogsParameters struct {
	CloudwatchLogs []CloudwatchLogsParameters `json:"cloudwatchLogs,omitempty" tf:"cloudwatch_logs"`

	Firehose []FirehoseParameters `json:"firehose,omitempty" tf:"firehose"`

	S3 []S3Parameters `json:"s3,omitempty" tf:"s3"`
}

type BrokerNodeGroupInfoObservation struct {
}

type BrokerNodeGroupInfoParameters struct {
	AzDistribution *string `json:"azDistribution,omitempty" tf:"az_distribution"`

	ClientSubnets []string `json:"clientSubnets" tf:"client_subnets"`

	EbsVolumeSize int64 `json:"ebsVolumeSize" tf:"ebs_volume_size"`

	InstanceType string `json:"instanceType" tf:"instance_type"`

	SecurityGroups []string `json:"securityGroups" tf:"security_groups"`
}

type ClientAuthenticationObservation struct {
}

type ClientAuthenticationParameters struct {
	Sasl []SaslParameters `json:"sasl,omitempty" tf:"sasl"`

	Tls []TlsParameters `json:"tls,omitempty" tf:"tls"`
}

type CloudwatchLogsObservation struct {
}

type CloudwatchLogsParameters struct {
	Enabled bool `json:"enabled" tf:"enabled"`

	LogGroup *string `json:"logGroup,omitempty" tf:"log_group"`
}

type ConfigurationInfoObservation struct {
}

type ConfigurationInfoParameters struct {
	Arn string `json:"arn" tf:"arn"`

	Revision int64 `json:"revision" tf:"revision"`
}

type EncryptionInTransitObservation struct {
}

type EncryptionInTransitParameters struct {
	ClientBroker *string `json:"clientBroker,omitempty" tf:"client_broker"`

	InCluster *bool `json:"inCluster,omitempty" tf:"in_cluster"`
}

type EncryptionInfoObservation struct {
}

type EncryptionInfoParameters struct {
	EncryptionAtRestKmsKeyArn *string `json:"encryptionAtRestKmsKeyArn,omitempty" tf:"encryption_at_rest_kms_key_arn"`

	EncryptionInTransit []EncryptionInTransitParameters `json:"encryptionInTransit,omitempty" tf:"encryption_in_transit"`
}

type FirehoseObservation struct {
}

type FirehoseParameters struct {
	DeliveryStream *string `json:"deliveryStream,omitempty" tf:"delivery_stream"`

	Enabled bool `json:"enabled" tf:"enabled"`
}

type JmxExporterObservation struct {
}

type JmxExporterParameters struct {
	EnabledInBroker bool `json:"enabledInBroker" tf:"enabled_in_broker"`
}

type LoggingInfoObservation struct {
}

type LoggingInfoParameters struct {
	BrokerLogs []BrokerLogsParameters `json:"brokerLogs" tf:"broker_logs"`
}

type MskClusterObservation struct {
	Arn string `json:"arn" tf:"arn"`

	BootstrapBrokers string `json:"bootstrapBrokers" tf:"bootstrap_brokers"`

	BootstrapBrokersSaslIam string `json:"bootstrapBrokersSaslIam" tf:"bootstrap_brokers_sasl_iam"`

	BootstrapBrokersSaslScram string `json:"bootstrapBrokersSaslScram" tf:"bootstrap_brokers_sasl_scram"`

	BootstrapBrokersTls string `json:"bootstrapBrokersTls" tf:"bootstrap_brokers_tls"`

	CurrentVersion string `json:"currentVersion" tf:"current_version"`

	ZookeeperConnectString string `json:"zookeeperConnectString" tf:"zookeeper_connect_string"`
}

type MskClusterParameters struct {
	BrokerNodeGroupInfo []BrokerNodeGroupInfoParameters `json:"brokerNodeGroupInfo" tf:"broker_node_group_info"`

	ClientAuthentication []ClientAuthenticationParameters `json:"clientAuthentication,omitempty" tf:"client_authentication"`

	ClusterName string `json:"clusterName" tf:"cluster_name"`

	ConfigurationInfo []ConfigurationInfoParameters `json:"configurationInfo,omitempty" tf:"configuration_info"`

	EncryptionInfo []EncryptionInfoParameters `json:"encryptionInfo,omitempty" tf:"encryption_info"`

	EnhancedMonitoring *string `json:"enhancedMonitoring,omitempty" tf:"enhanced_monitoring"`

	KafkaVersion string `json:"kafkaVersion" tf:"kafka_version"`

	LoggingInfo []LoggingInfoParameters `json:"loggingInfo,omitempty" tf:"logging_info"`

	NumberOfBrokerNodes int64 `json:"numberOfBrokerNodes" tf:"number_of_broker_nodes"`

	OpenMonitoring []OpenMonitoringParameters `json:"openMonitoring,omitempty" tf:"open_monitoring"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type NodeExporterObservation struct {
}

type NodeExporterParameters struct {
	EnabledInBroker bool `json:"enabledInBroker" tf:"enabled_in_broker"`
}

type OpenMonitoringObservation struct {
}

type OpenMonitoringParameters struct {
	Prometheus []PrometheusParameters `json:"prometheus" tf:"prometheus"`
}

type PrometheusObservation struct {
}

type PrometheusParameters struct {
	JmxExporter []JmxExporterParameters `json:"jmxExporter,omitempty" tf:"jmx_exporter"`

	NodeExporter []NodeExporterParameters `json:"nodeExporter,omitempty" tf:"node_exporter"`
}

type S3Observation struct {
}

type S3Parameters struct {
	Bucket *string `json:"bucket,omitempty" tf:"bucket"`

	Enabled bool `json:"enabled" tf:"enabled"`

	Prefix *string `json:"prefix,omitempty" tf:"prefix"`
}

type SaslObservation struct {
}

type SaslParameters struct {
	Iam *bool `json:"iam,omitempty" tf:"iam"`

	Scram *bool `json:"scram,omitempty" tf:"scram"`
}

type TlsObservation struct {
}

type TlsParameters struct {
	CertificateAuthorityArns []string `json:"certificateAuthorityArns,omitempty" tf:"certificate_authority_arns"`
}

// MskClusterSpec defines the desired state of MskCluster
type MskClusterSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       MskClusterParameters `json:"forProvider"`
}

// MskClusterStatus defines the observed state of MskCluster.
type MskClusterStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          MskClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MskCluster is the Schema for the MskClusters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type MskCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MskClusterSpec   `json:"spec"`
	Status            MskClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MskClusterList contains a list of MskClusters
type MskClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MskCluster `json:"items"`
}

// Repository type metadata.
var (
	MskClusterKind             = "MskCluster"
	MskClusterGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: MskClusterKind}.String()
	MskClusterKindAPIVersion   = MskClusterKind + "." + v1alpha1.GroupVersion.String()
	MskClusterGroupVersionKind = v1alpha1.GroupVersion.WithKind(MskClusterKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &MskCluster{}, &MskClusterList{})
}
