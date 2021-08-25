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
// +groupName=eks.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	apis "github.com/crossplane-contrib/provider-tf-aws/apis"
	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/eks/v1alpha1"
)

type CertificateAuthorityObservation struct {
	Data string `json:"data" tf:"data"`
}

type CertificateAuthorityParameters struct {
}

type EksClusterObservation struct {
	Arn string `json:"arn" tf:"arn"`

	CertificateAuthority []CertificateAuthorityObservation `json:"certificateAuthority" tf:"certificate_authority"`

	CreatedAt string `json:"createdAt" tf:"created_at"`

	Endpoint string `json:"endpoint" tf:"endpoint"`

	Identity []IdentityObservation `json:"identity" tf:"identity"`

	PlatformVersion string `json:"platformVersion" tf:"platform_version"`

	Status string `json:"status" tf:"status"`
}

type EksClusterParameters struct {
	EnabledClusterLogTypes []string `json:"enabledClusterLogTypes,omitempty" tf:"enabled_cluster_log_types"`

	EncryptionConfig []EncryptionConfigParameters `json:"encryptionConfig,omitempty" tf:"encryption_config"`

	KubernetesNetworkConfig []KubernetesNetworkConfigParameters `json:"kubernetesNetworkConfig,omitempty" tf:"kubernetes_network_config"`

	Name string `json:"name" tf:"name"`

	RoleArn string `json:"roleArn" tf:"role_arn"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	Version *string `json:"version,omitempty" tf:"version"`

	VpcConfig []VpcConfigParameters `json:"vpcConfig" tf:"vpc_config"`
}

type EncryptionConfigObservation struct {
}

type EncryptionConfigParameters struct {
	Provider []ProviderParameters `json:"provider" tf:"provider"`

	Resources []string `json:"resources" tf:"resources"`
}

type IdentityObservation struct {
	Oidc []OidcObservation `json:"oidc" tf:"oidc"`
}

type IdentityParameters struct {
}

type KubernetesNetworkConfigObservation struct {
}

type KubernetesNetworkConfigParameters struct {
	ServiceIpv4Cidr *string `json:"serviceIpv4Cidr,omitempty" tf:"service_ipv4_cidr"`
}

type OidcObservation struct {
	Issuer string `json:"issuer" tf:"issuer"`
}

type OidcParameters struct {
}

type ProviderObservation struct {
}

type ProviderParameters struct {
	KeyArn string `json:"keyArn" tf:"key_arn"`
}

type VpcConfigObservation struct {
	ClusterSecurityGroupId string `json:"clusterSecurityGroupId" tf:"cluster_security_group_id"`

	VpcId string `json:"vpcId" tf:"vpc_id"`
}

type VpcConfigParameters struct {
	EndpointPrivateAccess *bool `json:"endpointPrivateAccess,omitempty" tf:"endpoint_private_access"`

	EndpointPublicAccess *bool `json:"endpointPublicAccess,omitempty" tf:"endpoint_public_access"`

	PublicAccessCidrs []string `json:"publicAccessCidrs,omitempty" tf:"public_access_cidrs"`

	SecurityGroupIds []string `json:"securityGroupIds,omitempty" tf:"security_group_ids"`

	SubnetIds []string `json:"subnetIds" tf:"subnet_ids"`
}

// EksClusterSpec defines the desired state of EksCluster
type EksClusterSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       EksClusterParameters `json:"forProvider"`
}

// EksClusterStatus defines the observed state of EksCluster.
type EksClusterStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          EksClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EksCluster is the Schema for the EksClusters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type EksCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EksClusterSpec   `json:"spec"`
	Status            EksClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EksClusterList contains a list of EksClusters
type EksClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EksCluster `json:"items"`
}

// Repository type metadata.
var (
	EksClusterKind             = "EksCluster"
	EksClusterGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: EksClusterKind}.String()
	EksClusterKindAPIVersion   = EksClusterKind + "." + v1alpha1.GroupVersion.String()
	EksClusterGroupVersionKind = v1alpha1.GroupVersion.WithKind(EksClusterKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &EksCluster{}, &EksClusterList{})
}
