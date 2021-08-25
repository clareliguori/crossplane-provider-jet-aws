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
// +groupName=app.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	apis "github.com/crossplane-contrib/provider-tf-aws/apis"
	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/app/v1alpha1"
)

type AppCookieStickinessPolicyObservation struct {
}

type AppCookieStickinessPolicyParameters struct {
	CookieName string `json:"cookieName" tf:"cookie_name"`

	LbPort int64 `json:"lbPort" tf:"lb_port"`

	LoadBalancer string `json:"loadBalancer" tf:"load_balancer"`

	Name string `json:"name" tf:"name"`
}

// AppCookieStickinessPolicySpec defines the desired state of AppCookieStickinessPolicy
type AppCookieStickinessPolicySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       AppCookieStickinessPolicyParameters `json:"forProvider"`
}

// AppCookieStickinessPolicyStatus defines the observed state of AppCookieStickinessPolicy.
type AppCookieStickinessPolicyStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          AppCookieStickinessPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AppCookieStickinessPolicy is the Schema for the AppCookieStickinessPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type AppCookieStickinessPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppCookieStickinessPolicySpec   `json:"spec"`
	Status            AppCookieStickinessPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppCookieStickinessPolicyList contains a list of AppCookieStickinessPolicys
type AppCookieStickinessPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppCookieStickinessPolicy `json:"items"`
}

// Repository type metadata.
var (
	AppCookieStickinessPolicyKind             = "AppCookieStickinessPolicy"
	AppCookieStickinessPolicyGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: AppCookieStickinessPolicyKind}.String()
	AppCookieStickinessPolicyKindAPIVersion   = AppCookieStickinessPolicyKind + "." + v1alpha1.GroupVersion.String()
	AppCookieStickinessPolicyGroupVersionKind = v1alpha1.GroupVersion.WithKind(AppCookieStickinessPolicyKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &AppCookieStickinessPolicy{}, &AppCookieStickinessPolicyList{})
}
