// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package nodepool


type NodePoolSchedulerConfig struct {
	// If true, the node pool will have memory oversubscription enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.4.0/docs/resources/node_pool#memory_oversubscription NodePool#memory_oversubscription}
	MemoryOversubscription *string `field:"optional" json:"memoryOversubscription" yaml:"memoryOversubscription"`
	// The scheduler algorithm to use in the node pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.4.0/docs/resources/node_pool#scheduler_algorithm NodePool#scheduler_algorithm}
	SchedulerAlgorithm *string `field:"optional" json:"schedulerAlgorithm" yaml:"schedulerAlgorithm"`
}

