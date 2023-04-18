package externalvolume


type ExternalVolumeTopologyRequestPreferred struct {
	// topology block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/1.4.19/docs/resources/external_volume#topology ExternalVolume#topology}
	Topology interface{} `field:"required" json:"topology" yaml:"topology"`
}

