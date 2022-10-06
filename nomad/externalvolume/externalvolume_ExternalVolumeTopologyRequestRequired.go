package externalvolume


type ExternalVolumeTopologyRequestRequired struct {
	// topology block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/nomad/r/external_volume#topology ExternalVolume#topology}
	Topology interface{} `field:"required" json:"topology" yaml:"topology"`
}

