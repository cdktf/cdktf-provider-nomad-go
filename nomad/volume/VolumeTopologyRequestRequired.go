package volume


type VolumeTopologyRequestRequired struct {
	// topology block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/1.4.19/docs/resources/volume#topology Volume#topology}
	Topology interface{} `field:"required" json:"topology" yaml:"topology"`
}

