package externalvolume


type ExternalVolumeTopologyRequest struct {
	// preferred block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/nomad/r/external_volume#preferred ExternalVolume#preferred}
	Preferred *ExternalVolumeTopologyRequestPreferred `field:"optional" json:"preferred" yaml:"preferred"`
	// required block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/nomad/r/external_volume#required ExternalVolume#required}
	Required *ExternalVolumeTopologyRequestRequired `field:"optional" json:"required" yaml:"required"`
}

