package volume


type VolumeTopologyRequest struct {
	// required block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/1.4.20/docs/resources/volume#required Volume#required}
	Required *VolumeTopologyRequestRequired `field:"optional" json:"required" yaml:"required"`
}

