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

type OpsworksRailsAppLayerEbsVolumeObservation struct {
}

type OpsworksRailsAppLayerEbsVolumeParameters struct {

	// +kubebuilder:validation:Optional
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted"`

	// +kubebuilder:validation:Optional
	Iops *int64 `json:"iops,omitempty" tf:"iops"`

	// +kubebuilder:validation:Required
	MountPoint string `json:"mountPoint" tf:"mount_point"`

	// +kubebuilder:validation:Required
	NumberOfDisks int64 `json:"numberOfDisks" tf:"number_of_disks"`

	// +kubebuilder:validation:Optional
	RaidLevel *string `json:"raidLevel,omitempty" tf:"raid_level"`

	// +kubebuilder:validation:Required
	Size int64 `json:"size" tf:"size"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type OpsworksRailsAppLayerObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type OpsworksRailsAppLayerParameters struct {

	// +kubebuilder:validation:Optional
	AppServer *string `json:"appServer,omitempty" tf:"app_server"`

	// +kubebuilder:validation:Optional
	AutoAssignElasticIps *bool `json:"autoAssignElasticIps,omitempty" tf:"auto_assign_elastic_ips"`

	// +kubebuilder:validation:Optional
	AutoAssignPublicIps *bool `json:"autoAssignPublicIps,omitempty" tf:"auto_assign_public_ips"`

	// +kubebuilder:validation:Optional
	AutoHealing *bool `json:"autoHealing,omitempty" tf:"auto_healing"`

	// +kubebuilder:validation:Optional
	BundlerVersion *string `json:"bundlerVersion,omitempty" tf:"bundler_version"`

	// +kubebuilder:validation:Optional
	CustomConfigureRecipes []string `json:"customConfigureRecipes,omitempty" tf:"custom_configure_recipes"`

	// +kubebuilder:validation:Optional
	CustomDeployRecipes []string `json:"customDeployRecipes,omitempty" tf:"custom_deploy_recipes"`

	// +kubebuilder:validation:Optional
	CustomInstanceProfileArn *string `json:"customInstanceProfileArn,omitempty" tf:"custom_instance_profile_arn"`

	// +kubebuilder:validation:Optional
	CustomJSON *string `json:"customJson,omitempty" tf:"custom_json"`

	// +kubebuilder:validation:Optional
	CustomSecurityGroupIds []string `json:"customSecurityGroupIds,omitempty" tf:"custom_security_group_ids"`

	// +kubebuilder:validation:Optional
	CustomSetupRecipes []string `json:"customSetupRecipes,omitempty" tf:"custom_setup_recipes"`

	// +kubebuilder:validation:Optional
	CustomShutdownRecipes []string `json:"customShutdownRecipes,omitempty" tf:"custom_shutdown_recipes"`

	// +kubebuilder:validation:Optional
	CustomUndeployRecipes []string `json:"customUndeployRecipes,omitempty" tf:"custom_undeploy_recipes"`

	// +kubebuilder:validation:Optional
	DrainElbOnShutdown *bool `json:"drainElbOnShutdown,omitempty" tf:"drain_elb_on_shutdown"`

	// +kubebuilder:validation:Optional
	EbsVolume []OpsworksRailsAppLayerEbsVolumeParameters `json:"ebsVolume,omitempty" tf:"ebs_volume"`

	// +kubebuilder:validation:Optional
	ElasticLoadBalancer *string `json:"elasticLoadBalancer,omitempty" tf:"elastic_load_balancer"`

	// +kubebuilder:validation:Optional
	InstallUpdatesOnBoot *bool `json:"installUpdatesOnBoot,omitempty" tf:"install_updates_on_boot"`

	// +kubebuilder:validation:Optional
	InstanceShutdownTimeout *int64 `json:"instanceShutdownTimeout,omitempty" tf:"instance_shutdown_timeout"`

	// +kubebuilder:validation:Optional
	ManageBundler *bool `json:"manageBundler,omitempty" tf:"manage_bundler"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name"`

	// +kubebuilder:validation:Optional
	PassengerVersion *string `json:"passengerVersion,omitempty" tf:"passenger_version"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	RubyVersion *string `json:"rubyVersion,omitempty" tf:"ruby_version"`

	// +kubebuilder:validation:Optional
	RubygemsVersion *string `json:"rubygemsVersion,omitempty" tf:"rubygems_version"`

	// +kubebuilder:validation:Required
	StackID string `json:"stackId" tf:"stack_id"`

	// +kubebuilder:validation:Optional
	SystemPackages []string `json:"systemPackages,omitempty" tf:"system_packages"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	// +kubebuilder:validation:Optional
	UseEbsOptimizedInstances *bool `json:"useEbsOptimizedInstances,omitempty" tf:"use_ebs_optimized_instances"`
}

// OpsworksRailsAppLayerSpec defines the desired state of OpsworksRailsAppLayer
type OpsworksRailsAppLayerSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       OpsworksRailsAppLayerParameters `json:"forProvider"`
}

// OpsworksRailsAppLayerStatus defines the observed state of OpsworksRailsAppLayer.
type OpsworksRailsAppLayerStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          OpsworksRailsAppLayerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OpsworksRailsAppLayer is the Schema for the OpsworksRailsAppLayers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type OpsworksRailsAppLayer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OpsworksRailsAppLayerSpec   `json:"spec"`
	Status            OpsworksRailsAppLayerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OpsworksRailsAppLayerList contains a list of OpsworksRailsAppLayers
type OpsworksRailsAppLayerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OpsworksRailsAppLayer `json:"items"`
}

// Repository type metadata.
var (
	OpsworksRailsAppLayerKind             = "OpsworksRailsAppLayer"
	OpsworksRailsAppLayerGroupKind        = schema.GroupKind{Group: Group, Kind: OpsworksRailsAppLayerKind}.String()
	OpsworksRailsAppLayerKindAPIVersion   = OpsworksRailsAppLayerKind + "." + GroupVersion.String()
	OpsworksRailsAppLayerGroupVersionKind = GroupVersion.WithKind(OpsworksRailsAppLayerKind)
)

func init() {
	SchemeBuilder.Register(&OpsworksRailsAppLayer{}, &OpsworksRailsAppLayerList{})
}