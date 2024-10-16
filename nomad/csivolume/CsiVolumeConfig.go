// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package csivolume

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CsiVolumeConfig struct {
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
	// capability block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.4.0/docs/resources/csi_volume#capability CsiVolume#capability}
	Capability interface{} `field:"required" json:"capability" yaml:"capability"`
	// The display name of the volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.4.0/docs/resources/csi_volume#name CsiVolume#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of the CSI plugin that manages this volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.4.0/docs/resources/csi_volume#plugin_id CsiVolume#plugin_id}
	PluginId *string `field:"required" json:"pluginId" yaml:"pluginId"`
	// The unique ID of the volume, how jobs will refer to the volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.4.0/docs/resources/csi_volume#volume_id CsiVolume#volume_id}
	VolumeId *string `field:"required" json:"volumeId" yaml:"volumeId"`
	// Defines how large the volume can be.
	//
	// The storage provider may return a volume that is smaller than this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.4.0/docs/resources/csi_volume#capacity_max CsiVolume#capacity_max}
	CapacityMax *string `field:"optional" json:"capacityMax" yaml:"capacityMax"`
	// Defines how small the volume can be.
	//
	// The storage provider may return a volume that is larger than this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.4.0/docs/resources/csi_volume#capacity_min CsiVolume#capacity_min}
	CapacityMin *string `field:"optional" json:"capacityMin" yaml:"capacityMin"`
	// The volume ID to clone when creating this volume. Storage provider must support cloning. Conflicts with 'snapshot_id'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.4.0/docs/resources/csi_volume#clone_id CsiVolume#clone_id}
	CloneId *string `field:"optional" json:"cloneId" yaml:"cloneId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.4.0/docs/resources/csi_volume#id CsiVolume#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// mount_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.4.0/docs/resources/csi_volume#mount_options CsiVolume#mount_options}
	MountOptions *CsiVolumeMountOptions `field:"optional" json:"mountOptions" yaml:"mountOptions"`
	// The namespace in which to create the volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.4.0/docs/resources/csi_volume#namespace CsiVolume#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// An optional key-value map of strings passed directly to the CSI plugin to configure the volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.4.0/docs/resources/csi_volume#parameters CsiVolume#parameters}
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// An optional key-value map of strings used as credentials for publishing and unpublishing volumes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.4.0/docs/resources/csi_volume#secrets CsiVolume#secrets}
	Secrets *map[string]*string `field:"optional" json:"secrets" yaml:"secrets"`
	// The snapshot ID to restore when creating this volume. Storage provider must support snapshots. Conflicts with 'clone_id'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.4.0/docs/resources/csi_volume#snapshot_id CsiVolume#snapshot_id}
	SnapshotId *string `field:"optional" json:"snapshotId" yaml:"snapshotId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.4.0/docs/resources/csi_volume#timeouts CsiVolume#timeouts}
	Timeouts *CsiVolumeTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// topology_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.4.0/docs/resources/csi_volume#topology_request CsiVolume#topology_request}
	TopologyRequest *CsiVolumeTopologyRequest `field:"optional" json:"topologyRequest" yaml:"topologyRequest"`
}

