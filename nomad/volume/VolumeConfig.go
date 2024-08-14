// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package volume

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VolumeConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.1/docs/resources/volume#external_id Volume#external_id}
	ExternalId *string `field:"required" json:"externalId" yaml:"externalId"`
	// The display name of the volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.1/docs/resources/volume#name Volume#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of the CSI plugin that manages this volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.1/docs/resources/volume#plugin_id Volume#plugin_id}
	PluginId *string `field:"required" json:"pluginId" yaml:"pluginId"`
	// The unique ID of the volume, how jobs will refer to the volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.1/docs/resources/volume#volume_id Volume#volume_id}
	VolumeId *string `field:"required" json:"volumeId" yaml:"volumeId"`
	// Defines whether a volume should be available concurrently.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.1/docs/resources/volume#access_mode Volume#access_mode}
	AccessMode *string `field:"optional" json:"accessMode" yaml:"accessMode"`
	// The storage API that will be used by the volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.1/docs/resources/volume#attachment_mode Volume#attachment_mode}
	AttachmentMode *string `field:"optional" json:"attachmentMode" yaml:"attachmentMode"`
	// capability block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.1/docs/resources/volume#capability Volume#capability}
	Capability interface{} `field:"optional" json:"capability" yaml:"capability"`
	// An optional key-value map of strings passed directly to the CSI plugin to validate the volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.1/docs/resources/volume#context Volume#context}
	Context *map[string]*string `field:"optional" json:"context" yaml:"context"`
	// If true, the volume will be deregistered on destroy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.1/docs/resources/volume#deregister_on_destroy Volume#deregister_on_destroy}
	DeregisterOnDestroy interface{} `field:"optional" json:"deregisterOnDestroy" yaml:"deregisterOnDestroy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.1/docs/resources/volume#id Volume#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// mount_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.1/docs/resources/volume#mount_options Volume#mount_options}
	MountOptions *VolumeMountOptions `field:"optional" json:"mountOptions" yaml:"mountOptions"`
	// The namespace in which to create the volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.1/docs/resources/volume#namespace Volume#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// An optional key-value map of strings passed directly to the CSI plugin to configure the volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.1/docs/resources/volume#parameters Volume#parameters}
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// An optional key-value map of strings used as credentials for publishing and unpublishing volumes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.1/docs/resources/volume#secrets Volume#secrets}
	Secrets *map[string]*string `field:"optional" json:"secrets" yaml:"secrets"`
	// topology_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.1/docs/resources/volume#topology_request Volume#topology_request}
	TopologyRequest *VolumeTopologyRequest `field:"optional" json:"topologyRequest" yaml:"topologyRequest"`
	// The type of the volume. Currently, only 'csi' is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.1/docs/resources/volume#type Volume#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

