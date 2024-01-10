// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datanomadjob

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataNomadJobConstraintsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataNomadJobConstraintsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataNomadJobConstraintsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataNomadJobConstraintsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataNomadJobConstraintsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataNomadJobConstraintsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataNomadJobConstraintsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

