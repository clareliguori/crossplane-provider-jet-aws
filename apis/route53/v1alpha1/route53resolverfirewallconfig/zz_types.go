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
// +groupName=route53.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	apis "github.com/crossplane-contrib/provider-tf-aws/apis"
	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/route53/v1alpha1"
)

type Route53ResolverFirewallConfigObservation struct {
	OwnerId string `json:"ownerId" tf:"owner_id"`
}

type Route53ResolverFirewallConfigParameters struct {
	FirewallFailOpen *string `json:"firewallFailOpen,omitempty" tf:"firewall_fail_open"`

	ResourceId string `json:"resourceId" tf:"resource_id"`
}

// Route53ResolverFirewallConfigSpec defines the desired state of Route53ResolverFirewallConfig
type Route53ResolverFirewallConfigSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       Route53ResolverFirewallConfigParameters `json:"forProvider"`
}

// Route53ResolverFirewallConfigStatus defines the observed state of Route53ResolverFirewallConfig.
type Route53ResolverFirewallConfigStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          Route53ResolverFirewallConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Route53ResolverFirewallConfig is the Schema for the Route53ResolverFirewallConfigs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Route53ResolverFirewallConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Route53ResolverFirewallConfigSpec   `json:"spec"`
	Status            Route53ResolverFirewallConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Route53ResolverFirewallConfigList contains a list of Route53ResolverFirewallConfigs
type Route53ResolverFirewallConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Route53ResolverFirewallConfig `json:"items"`
}

// Repository type metadata.
var (
	Route53ResolverFirewallConfigKind             = "Route53ResolverFirewallConfig"
	Route53ResolverFirewallConfigGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: Route53ResolverFirewallConfigKind}.String()
	Route53ResolverFirewallConfigKindAPIVersion   = Route53ResolverFirewallConfigKind + "." + v1alpha1.GroupVersion.String()
	Route53ResolverFirewallConfigGroupVersionKind = v1alpha1.GroupVersion.WithKind(Route53ResolverFirewallConfigKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &Route53ResolverFirewallConfig{}, &Route53ResolverFirewallConfigList{})
}
