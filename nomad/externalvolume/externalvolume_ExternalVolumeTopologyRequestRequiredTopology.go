package externalvolume


type ExternalVolumeTopologyRequestRequiredTopology struct {
	// Define the attributes for the topology request.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/nomad/r/external_volume#segments ExternalVolume#segments}
	Segments *map[string]*string `field:"required" json:"segments" yaml:"segments"`
}

