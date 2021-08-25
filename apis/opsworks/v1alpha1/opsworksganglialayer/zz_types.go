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
// +groupName=opsworks.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	apis "github.com/crossplane-contrib/provider-tf-aws/apis"
	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/opsworks/v1alpha1"
)

type EbsVolumeObservation struct {
}

type EbsVolumeParameters struct {
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted"`

	Iops *int64 `json:"iops,omitempty" tf:"iops"`

	MountPoint string `json:"mountPoint" tf:"mount_point"`

	NumberOfDisks int64 `json:"numberOfDisks" tf:"number_of_disks"`

	RaidLevel *string `json:"raidLevel,omitempty" tf:"raid_level"`

	Size int64 `json:"size" tf:"size"`

	Type *string `json:"type,omitempty" tf:"type"`
}

type OpsworksGangliaLayerObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type OpsworksGangliaLayerParameters struct {
	AutoAssignElasticIps *bool `json:"autoAssignElasticIps,omitempty" tf:"auto_assign_elastic_ips"`

	AutoAssignPublicIps *bool `json:"autoAssignPublicIps,omitempty" tf:"auto_assign_public_ips"`

	AutoHealing *bool `json:"autoHealing,omitempty" tf:"auto_healing"`

	CustomConfigureRecipes []string `json:"customConfigureRecipes,omitempty" tf:"custom_configure_recipes"`

	CustomDeployRecipes []string `json:"customDeployRecipes,omitempty" tf:"custom_deploy_recipes"`

	CustomInstanceProfileArn *string `json:"customInstanceProfileArn,omitempty" tf:"custom_instance_profile_arn"`

	CustomJson *string `json:"customJson,omitempty" tf:"custom_json"`

	CustomSecurityGroupIds []string `json:"customSecurityGroupIds,omitempty" tf:"custom_security_group_ids"`

	CustomSetupRecipes []string `json:"customSetupRecipes,omitempty" tf:"custom_setup_recipes"`

	CustomShutdownRecipes []string `json:"customShutdownRecipes,omitempty" tf:"custom_shutdown_recipes"`

	CustomUndeployRecipes []string `json:"customUndeployRecipes,omitempty" tf:"custom_undeploy_recipes"`

	DrainElbOnShutdown *bool `json:"drainElbOnShutdown,omitempty" tf:"drain_elb_on_shutdown"`

	EbsVolume []EbsVolumeParameters `json:"ebsVolume,omitempty" tf:"ebs_volume"`

	ElasticLoadBalancer *string `json:"elasticLoadBalancer,omitempty" tf:"elastic_load_balancer"`

	InstallUpdatesOnBoot *bool `json:"installUpdatesOnBoot,omitempty" tf:"install_updates_on_boot"`

	InstanceShutdownTimeout *int64 `json:"instanceShutdownTimeout,omitempty" tf:"instance_shutdown_timeout"`

	Name *string `json:"name,omitempty" tf:"name"`

	Password string `json:"password" tf:"password"`

	StackId string `json:"stackId" tf:"stack_id"`

	SystemPackages []string `json:"systemPackages,omitempty" tf:"system_packages"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	Url *string `json:"url,omitempty" tf:"url"`

	UseEbsOptimizedInstances *bool `json:"useEbsOptimizedInstances,omitempty" tf:"use_ebs_optimized_instances"`

	Username *string `json:"username,omitempty" tf:"username"`
}

// OpsworksGangliaLayerSpec defines the desired state of OpsworksGangliaLayer
type OpsworksGangliaLayerSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       OpsworksGangliaLayerParameters `json:"forProvider"`
}

// OpsworksGangliaLayerStatus defines the observed state of OpsworksGangliaLayer.
type OpsworksGangliaLayerStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          OpsworksGangliaLayerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OpsworksGangliaLayer is the Schema for the OpsworksGangliaLayers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type OpsworksGangliaLayer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OpsworksGangliaLayerSpec   `json:"spec"`
	Status            OpsworksGangliaLayerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OpsworksGangliaLayerList contains a list of OpsworksGangliaLayers
type OpsworksGangliaLayerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OpsworksGangliaLayer `json:"items"`
}

// Repository type metadata.
var (
	OpsworksGangliaLayerKind             = "OpsworksGangliaLayer"
	OpsworksGangliaLayerGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: OpsworksGangliaLayerKind}.String()
	OpsworksGangliaLayerKindAPIVersion   = OpsworksGangliaLayerKind + "." + v1alpha1.GroupVersion.String()
	OpsworksGangliaLayerGroupVersionKind = v1alpha1.GroupVersion.WithKind(OpsworksGangliaLayerKind)
)

func init() {
	apis.SchemaMap[v1alpha1.GroupVersion] = append(apis.SchemaMap[v1alpha1.GroupVersion], &OpsworksGangliaLayer{}, &OpsworksGangliaLayerList{})
}
