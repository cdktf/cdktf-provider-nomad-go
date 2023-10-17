// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package variable

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_Variable) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (v *jsiiProxy_Variable) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (v *jsiiProxy_Variable) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (v *jsiiProxy_Variable) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (v *jsiiProxy_Variable) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (v *jsiiProxy_Variable) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (v *jsiiProxy_Variable) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (v *jsiiProxy_Variable) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (v *jsiiProxy_Variable) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (v *jsiiProxy_Variable) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (v *jsiiProxy_Variable) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (v *jsiiProxy_Variable) validateImportFromParameters(id *string) error {
	return nil
}

func (v *jsiiProxy_Variable) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (v *jsiiProxy_Variable) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (v *jsiiProxy_Variable) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateVariable_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateVariable_IsConstructParameters(x interface{}) error {
	return nil
}

func validateVariable_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateVariable_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Variable) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Variable) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Variable) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Variable) validateSetItemsParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_Variable) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Variable) validateSetNamespaceParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Variable) validateSetPathParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Variable) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func validateNewVariableParameters(scope constructs.Construct, id *string, config *VariableConfig) error {
	return nil
}

