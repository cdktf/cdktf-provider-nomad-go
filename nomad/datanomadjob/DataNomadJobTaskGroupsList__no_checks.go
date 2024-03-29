// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datanomadjob

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataNomadJobTaskGroupsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataNomadJobTaskGroupsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataNomadJobTaskGroupsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataNomadJobTaskGroupsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataNomadJobTaskGroupsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataNomadJobTaskGroupsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataNomadJobTaskGroupsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

