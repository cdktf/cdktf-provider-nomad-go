// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package csivolumeregistration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CsiVolumeRegistrationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The ID of the physical volume from the storage provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.0.0/docs/resources/csi_volume_registration#external_id CsiVolumeRegistration#external_id}
	ExternalId *string `field:"required" json:"externalId" yaml:"externalId"`
	// The display name of the volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.0.0/docs/resources/csi_volume_registration#name CsiVolumeRegistration#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of the CSI plugin that manages this volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.0.0/docs/resources/csi_volume_registration#plugin_id CsiVolumeRegistration#plugin_id}
	PluginId *string `field:"required" json:"pluginId" yaml:"pluginId"`
	// The unique ID of the volume, how jobs will refer to the volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.0.0/docs/resources/csi_volume_registration#volume_id CsiVolumeRegistration#volume_id}
	VolumeId *string `field:"required" json:"volumeId" yaml:"volumeId"`
	// capability block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.0.0/docs/resources/csi_volume_registration#capability CsiVolumeRegistration#capability}
	Capability interface{} `field:"optional" json:"capability" yaml:"capability"`
	// An optional key-value map of strings passed directly to the CSI plugin to validate the volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.0.0/docs/resources/csi_volume_registration#context CsiVolumeRegistration#context}
	Context *map[string]*string `field:"optional" json:"context" yaml:"context"`
	// If true, the volume will be deregistered on destroy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.0.0/docs/resources/csi_volume_registration#deregister_on_destroy CsiVolumeRegistration#deregister_on_destroy}
	DeregisterOnDestroy interface{} `field:"optional" json:"deregisterOnDestroy" yaml:"deregisterOnDestroy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.0.0/docs/resources/csi_volume_registration#id CsiVolumeRegistration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// mount_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.0.0/docs/resources/csi_volume_registration#mount_options CsiVolumeRegistration#mount_options}
	MountOptions *CsiVolumeRegistrationMountOptions `field:"optional" json:"mountOptions" yaml:"mountOptions"`
	// The namespace in which to create the volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.0.0/docs/resources/csi_volume_registration#namespace CsiVolumeRegistration#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// An optional key-value map of strings passed directly to the CSI plugin to configure the volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.0.0/docs/resources/csi_volume_registration#parameters CsiVolumeRegistration#parameters}
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// An optional key-value map of strings used as credentials for publishing and unpublishing volumes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.0.0/docs/resources/csi_volume_registration#secrets CsiVolumeRegistration#secrets}
	Secrets *map[string]*string `field:"optional" json:"secrets" yaml:"secrets"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.0.0/docs/resources/csi_volume_registration#timeouts CsiVolumeRegistration#timeouts}
	Timeouts *CsiVolumeRegistrationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// topology_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.0.0/docs/resources/csi_volume_registration#topology_request CsiVolumeRegistration#topology_request}
	TopologyRequest *CsiVolumeRegistrationTopologyRequest `field:"optional" json:"topologyRequest" yaml:"topologyRequest"`
}

