package volume


type VolumeMountOptions struct {
	// The file system type.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/nomad/r/volume#fs_type Volume#fs_type}
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// The flags passed to mount.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/nomad/r/volume#mount_flags Volume#mount_flags}
	MountFlags *[]*string `field:"optional" json:"mountFlags" yaml:"mountFlags"`
}

