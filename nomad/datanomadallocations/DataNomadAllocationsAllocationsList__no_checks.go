// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datanomadallocations

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataNomadAllocationsAllocationsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataNomadAllocationsAllocationsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataNomadAllocationsAllocationsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataNomadAllocationsAllocationsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataNomadAllocationsAllocationsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataNomadAllocationsAllocationsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataNomadAllocationsAllocationsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

