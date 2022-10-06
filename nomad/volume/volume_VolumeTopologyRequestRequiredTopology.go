package volume


type VolumeTopologyRequestRequiredTopology struct {
	// Define attributes for the topology request.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/nomad/r/volume#segments Volume#segments}
	Segments *map[string]*string `field:"required" json:"segments" yaml:"segments"`
}

