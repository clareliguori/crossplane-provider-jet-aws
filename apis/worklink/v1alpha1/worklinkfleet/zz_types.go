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
// +groupName=worklink.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	apis "github.com/crossplane-contrib/provider-tf-aws/apis"
	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/worklink/v1alpha1"
)

type IdentityProviderObservation struct {
}

type IdentityProviderParameters struct {
	SamlMetadata string `json:"samlMetadata" tf:"saml_metadata"`

	Type string `json:"type" tf:"type"`
}

type NetworkObservation struct {
}

type NetworkParameters struct {
	SecurityGroupIds []string `json:"securityGroupIds" tf:"security_group_ids"`

	SubnetIds []string `json:"subnetIds" tf:"subnet_ids"`

	VpcId string `json:"vpcId" tf:"vpc_id"`
}

type WorklinkFleetObservation struct {
	Arn string `json:"arn" tf:"arn"`

	CompanyCode string `json:"companyCode" tf:"company_code"`

	CreatedTime string `json:"createdTime" tf:"created_time"`

	LastUpdatedTime string `json:"lastUpdatedTime" tf:"last_updated_time"`
}

type WorklinkFleetParameters struct {
	AuditStreamArn *string `json:"auditStreamArn,omitempty" tf:"audit_stream_arn"`

	DeviceCaCertificate *string `json:"deviceCaCertificate,omitempty" tf:"device_ca_certificate"`

	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`

	IdentityProvider []IdentityProviderParameters `json:"identityProvider,omitempty" tf:"identity_provider"`

	Name string `json:"name" tf:"name"`

	Network []NetworkParameters `json:"network,omitempty" tf:"network"`

	OptimizeForEndUserLocation *bool `json:"optimizeForEndUserLocation,omitempty" tf:"optimize_for_end_user_location"`
}

// WorklinkFleetSpec defines the desired state of WorklinkFleet
type WorklinkFleetSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       WorklinkFleetParameters `json:"forProvider"`
}

// WorklinkFleetStatus defines the observed state of WorklinkFleet.
type WorklinkFleetStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          WorklinkFleetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// WorklinkFleet is the Schema for the WorklinkFleets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type WorklinkFleet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WorklinkFleetSpec   `json:"spec"`
	Status            WorklinkFleetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorklinkFleetList contains a list of WorklinkFleets
type WorklinkFleetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WorklinkFleet `json:"items"`
}

// Repository type metadata.
var (
	WorklinkFleetKind             = "WorklinkFleet"
	WorklinkFleetGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: WorklinkFleetKind}.String()
	WorklinkFleetKindAPIVersion   = WorklinkFleetKind + "." + v1alpha1.GroupVersion.String()
	WorklinkFleetGroupVersionKind = v1alpha1.GroupVersion.WithKind(WorklinkFleetKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &WorklinkFleet{}, &WorklinkFleetList{})
}
