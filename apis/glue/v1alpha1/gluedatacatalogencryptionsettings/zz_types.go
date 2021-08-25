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
// +groupName=glue.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	apis "github.com/crossplane-contrib/provider-tf-aws/apis"
	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/glue/v1alpha1"
)

type ConnectionPasswordEncryptionObservation struct {
}

type ConnectionPasswordEncryptionParameters struct {
	AwsKmsKeyId *string `json:"awsKmsKeyId,omitempty" tf:"aws_kms_key_id"`

	ReturnConnectionPasswordEncrypted bool `json:"returnConnectionPasswordEncrypted" tf:"return_connection_password_encrypted"`
}

type DataCatalogEncryptionSettingsObservation struct {
}

type DataCatalogEncryptionSettingsParameters struct {
	ConnectionPasswordEncryption []ConnectionPasswordEncryptionParameters `json:"connectionPasswordEncryption" tf:"connection_password_encryption"`

	EncryptionAtRest []EncryptionAtRestParameters `json:"encryptionAtRest" tf:"encryption_at_rest"`
}

type EncryptionAtRestObservation struct {
}

type EncryptionAtRestParameters struct {
	CatalogEncryptionMode string `json:"catalogEncryptionMode" tf:"catalog_encryption_mode"`

	SseAwsKmsKeyId *string `json:"sseAwsKmsKeyId,omitempty" tf:"sse_aws_kms_key_id"`
}

type GlueDataCatalogEncryptionSettingsObservation struct {
}

type GlueDataCatalogEncryptionSettingsParameters struct {
	CatalogId *string `json:"catalogId,omitempty" tf:"catalog_id"`

	DataCatalogEncryptionSettings []DataCatalogEncryptionSettingsParameters `json:"dataCatalogEncryptionSettings" tf:"data_catalog_encryption_settings"`
}

// GlueDataCatalogEncryptionSettingsSpec defines the desired state of GlueDataCatalogEncryptionSettings
type GlueDataCatalogEncryptionSettingsSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       GlueDataCatalogEncryptionSettingsParameters `json:"forProvider"`
}

// GlueDataCatalogEncryptionSettingsStatus defines the observed state of GlueDataCatalogEncryptionSettings.
type GlueDataCatalogEncryptionSettingsStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          GlueDataCatalogEncryptionSettingsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GlueDataCatalogEncryptionSettings is the Schema for the GlueDataCatalogEncryptionSettingss API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type GlueDataCatalogEncryptionSettings struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GlueDataCatalogEncryptionSettingsSpec   `json:"spec"`
	Status            GlueDataCatalogEncryptionSettingsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GlueDataCatalogEncryptionSettingsList contains a list of GlueDataCatalogEncryptionSettingss
type GlueDataCatalogEncryptionSettingsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlueDataCatalogEncryptionSettings `json:"items"`
}

// Repository type metadata.
var (
	GlueDataCatalogEncryptionSettingsKind             = "GlueDataCatalogEncryptionSettings"
	GlueDataCatalogEncryptionSettingsGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: GlueDataCatalogEncryptionSettingsKind}.String()
	GlueDataCatalogEncryptionSettingsKindAPIVersion   = GlueDataCatalogEncryptionSettingsKind + "." + v1alpha1.GroupVersion.String()
	GlueDataCatalogEncryptionSettingsGroupVersionKind = v1alpha1.GroupVersion.WithKind(GlueDataCatalogEncryptionSettingsKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &GlueDataCatalogEncryptionSettings{}, &GlueDataCatalogEncryptionSettingsList{})
}
