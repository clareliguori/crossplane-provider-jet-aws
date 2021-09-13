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

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type ConversationLogsObservation struct {
}

type ConversationLogsParameters struct {

	// +kubebuilder:validation:Required
	IamRoleArn string `json:"iamRoleArn" tf:"iam_role_arn"`

	// +kubebuilder:validation:Optional
	LogSettings []LogSettingsParameters `json:"logSettings,omitempty" tf:"log_settings"`
}

type LexBotAliasObservation struct {
	Arn string `json:"arn" tf:"arn"`

	Checksum string `json:"checksum" tf:"checksum"`

	CreatedDate string `json:"createdDate" tf:"created_date"`

	LastUpdatedDate string `json:"lastUpdatedDate" tf:"last_updated_date"`
}

type LexBotAliasParameters struct {

	// +kubebuilder:validation:Required
	BotName string `json:"botName" tf:"bot_name"`

	// +kubebuilder:validation:Required
	BotVersion string `json:"botVersion" tf:"bot_version"`

	// +kubebuilder:validation:Optional
	ConversationLogs []ConversationLogsParameters `json:"conversationLogs,omitempty" tf:"conversation_logs"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`
}

type LogSettingsObservation struct {
	ResourcePrefix string `json:"resourcePrefix" tf:"resource_prefix"`
}

type LogSettingsParameters struct {

	// +kubebuilder:validation:Required
	Destination string `json:"destination" tf:"destination"`

	// +kubebuilder:validation:Optional
	KmsKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn"`

	// +kubebuilder:validation:Required
	LogType string `json:"logType" tf:"log_type"`

	// +kubebuilder:validation:Required
	ResourceArn string `json:"resourceArn" tf:"resource_arn"`
}

// LexBotAliasSpec defines the desired state of LexBotAlias
type LexBotAliasSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       LexBotAliasParameters `json:"forProvider"`
}

// LexBotAliasStatus defines the observed state of LexBotAlias.
type LexBotAliasStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          LexBotAliasObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LexBotAlias is the Schema for the LexBotAliass API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type LexBotAlias struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LexBotAliasSpec   `json:"spec"`
	Status            LexBotAliasStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LexBotAliasList contains a list of LexBotAliass
type LexBotAliasList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LexBotAlias `json:"items"`
}

// Repository type metadata.
var (
	LexBotAliasKind             = "LexBotAlias"
	LexBotAliasGroupKind        = schema.GroupKind{Group: Group, Kind: LexBotAliasKind}.String()
	LexBotAliasKindAPIVersion   = LexBotAliasKind + "." + GroupVersion.String()
	LexBotAliasGroupVersionKind = GroupVersion.WithKind(LexBotAliasKind)
)

func init() {
	SchemeBuilder.Register(&LexBotAlias{}, &LexBotAliasList{})
}