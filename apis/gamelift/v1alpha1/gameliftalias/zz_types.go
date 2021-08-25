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
// +groupName=gamelift.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	apis "github.com/crossplane-contrib/provider-tf-aws/apis"
	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/gamelift/v1alpha1"
)

type GameliftAliasObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type GameliftAliasParameters struct {
	Description *string `json:"description,omitempty" tf:"description"`

	Name string `json:"name" tf:"name"`

	RoutingStrategy []RoutingStrategyParameters `json:"routingStrategy" tf:"routing_strategy"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type RoutingStrategyObservation struct {
}

type RoutingStrategyParameters struct {
	FleetId *string `json:"fleetId,omitempty" tf:"fleet_id"`

	Message *string `json:"message,omitempty" tf:"message"`

	Type string `json:"type" tf:"type"`
}

// GameliftAliasSpec defines the desired state of GameliftAlias
type GameliftAliasSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       GameliftAliasParameters `json:"forProvider"`
}

// GameliftAliasStatus defines the observed state of GameliftAlias.
type GameliftAliasStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          GameliftAliasObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GameliftAlias is the Schema for the GameliftAliass API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type GameliftAlias struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GameliftAliasSpec   `json:"spec"`
	Status            GameliftAliasStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GameliftAliasList contains a list of GameliftAliass
type GameliftAliasList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GameliftAlias `json:"items"`
}

// Repository type metadata.
var (
	GameliftAliasKind             = "GameliftAlias"
	GameliftAliasGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: GameliftAliasKind}.String()
	GameliftAliasKindAPIVersion   = GameliftAliasKind + "." + v1alpha1.GroupVersion.String()
	GameliftAliasGroupVersionKind = v1alpha1.GroupVersion.WithKind(GameliftAliasKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &GameliftAlias{}, &GameliftAliasList{})
}
