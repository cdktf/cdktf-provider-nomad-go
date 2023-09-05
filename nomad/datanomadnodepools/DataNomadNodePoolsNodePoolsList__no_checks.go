// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datanomadnodepools

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataNomadNodePoolsNodePoolsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataNomadNodePoolsNodePoolsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataNomadNodePoolsNodePoolsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataNomadNodePoolsNodePoolsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataNomadNodePoolsNodePoolsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataNomadNodePoolsNodePoolsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

