// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package nodepool

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NodePool) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (n *jsiiProxy_NodePool) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (n *jsiiProxy_NodePool) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (n *jsiiProxy_NodePool) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (n *jsiiProxy_NodePool) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (n *jsiiProxy_NodePool) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (n *jsiiProxy_NodePool) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (n *jsiiProxy_NodePool) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (n *jsiiProxy_NodePool) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (n *jsiiProxy_NodePool) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (n *jsiiProxy_NodePool) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (n *jsiiProxy_NodePool) validateImportFromParameters(id *string) error {
	return nil
}

func (n *jsiiProxy_NodePool) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (n *jsiiProxy_NodePool) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (n *jsiiProxy_NodePool) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (n *jsiiProxy_NodePool) validatePutSchedulerConfigParameters(value *NodePoolSchedulerConfig) error {
	return nil
}

func validateNodePool_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateNodePool_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNodePool_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateNodePool_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_NodePool) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NodePool) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NodePool) validateSetDescriptionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NodePool) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NodePool) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_NodePool) validateSetMetaParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_NodePool) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NodePool) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func validateNewNodePoolParameters(scope constructs.Construct, id *string, config *NodePoolConfig) error {
	return nil
}

