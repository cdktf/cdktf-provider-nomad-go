package volume


type VolumeTopologyRequestRequired struct {
	// topology block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/nomad/r/volume#topology Volume#topology}
	Topology interface{} `field:"required" json:"topology" yaml:"topology"`
}

