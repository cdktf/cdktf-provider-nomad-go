package volume


type VolumeTopologyRequest struct {
	// required block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/nomad/r/volume#required Volume#required}
	Required *VolumeTopologyRequestRequired `field:"optional" json:"required" yaml:"required"`
}

