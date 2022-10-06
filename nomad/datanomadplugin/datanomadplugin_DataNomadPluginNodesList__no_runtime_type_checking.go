//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package datanomadplugin

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataNomadPluginNodesList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataNomadPluginNodesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataNomadPluginNodesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataNomadPluginNodesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataNomadPluginNodesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataNomadPluginNodesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

