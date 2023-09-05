// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datanomadnamespace

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataNomadNamespaceNodePoolConfigList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataNomadNamespaceNodePoolConfigList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataNomadNamespaceNodePoolConfigList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataNomadNamespaceNodePoolConfigList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataNomadNamespaceNodePoolConfigList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataNomadNamespaceNodePoolConfigListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

