// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dynamichostvolume

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-nomad-go/nomad/v10/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-nomad-go/nomad/v10/dynamichostvolume/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/nomad/2.5.1/docs/resources/dynamic_host_volume nomad_dynamic_host_volume}.
type DynamicHostVolume interface {
	cdktf.TerraformResource
	Capability() DynamicHostVolumeCapabilityList
	CapabilityInput() interface{}
	Capacity() *string
	CapacityBytes() *float64
	CapacityMax() *string
	SetCapacityMax(val *string)
	CapacityMaxBytes() *float64
	CapacityMaxInput() *string
	CapacityMin() *string
	SetCapacityMin(val *string)
	CapacityMinBytes() *float64
	CapacityMinInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	Constraint() DynamicHostVolumeConstraintList
	ConstraintInput() interface{}
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HostPath() *string
	Id() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	// The tree node.
	Node() constructs.Node
	NodeId() *string
	SetNodeId(val *string)
	NodeIdInput() *string
	NodePool() *string
	SetNodePool(val *string)
	NodePoolInput() *string
	Parameters() *map[string]*string
	SetParameters(val *map[string]*string)
	ParametersInput() *map[string]*string
	PluginId() *string
	SetPluginId(val *string)
	PluginIdInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	State() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutCapability(value interface{})
	PutConstraint(value interface{})
	ResetCapacityMax()
	ResetCapacityMin()
	ResetConstraint()
	ResetNamespace()
	ResetNodeId()
	ResetNodePool()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParameters()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DynamicHostVolume
type jsiiProxy_DynamicHostVolume struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DynamicHostVolume) Capability() DynamicHostVolumeCapabilityList {
	var returns DynamicHostVolumeCapabilityList
	_jsii_.Get(
		j,
		"capability",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamicHostVolume) CapabilityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"capabilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamicHostVolume) Capacity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamicHostVolume) CapacityBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"capacityBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamicHostVolume) CapacityMax() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityMax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamicHostVolume) CapacityMaxBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"capacityMaxBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamicHostVolume) CapacityMaxInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityMaxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamicHostVolume) CapacityMin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityMin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamicHostVolume) CapacityMinBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"capacityMinBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamicHostVolume) CapacityMinInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityMinInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamicHostVolume) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamicHostVolume) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamicHostVolume) Constraint() DynamicHostVolumeConstraintList {
	var returns DynamicHostVolumeConstraintList
	_jsii_.Get(
		j,
		"constraint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamicHostVolume) ConstraintInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"constraintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamicHostVolume) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamicHostVolume) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamicHostVolume) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamicHostVolume) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamicHostVolume) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamicHostVolume) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamicHostVolume) HostPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamicHostVolume) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamicHostVolume) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamicHostVolume) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamicHostVolume) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamicHostVolume) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamicHostVolume) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamicHostVolume) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamicHostVolume) NodeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamicHostVolume) NodeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamicHostVolume) NodePool() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodePool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamicHostVolume) NodePoolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodePoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamicHostVolume) Parameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamicHostVolume) ParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamicHostVolume) PluginId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamicHostVolume) PluginIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamicHostVolume) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamicHostVolume) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamicHostVolume) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamicHostVolume) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamicHostVolume) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamicHostVolume) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamicHostVolume) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/nomad/2.5.1/docs/resources/dynamic_host_volume nomad_dynamic_host_volume} Resource.
func NewDynamicHostVolume(scope constructs.Construct, id *string, config *DynamicHostVolumeConfig) DynamicHostVolume {
	_init_.Initialize()

	if err := validateNewDynamicHostVolumeParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DynamicHostVolume{}

	_jsii_.Create(
		"@cdktf/provider-nomad.dynamicHostVolume.DynamicHostVolume",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/nomad/2.5.1/docs/resources/dynamic_host_volume nomad_dynamic_host_volume} Resource.
func NewDynamicHostVolume_Override(d DynamicHostVolume, scope constructs.Construct, id *string, config *DynamicHostVolumeConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-nomad.dynamicHostVolume.DynamicHostVolume",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DynamicHostVolume)SetCapacityMax(val *string) {
	if err := j.validateSetCapacityMaxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacityMax",
		val,
	)
}

func (j *jsiiProxy_DynamicHostVolume)SetCapacityMin(val *string) {
	if err := j.validateSetCapacityMinParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacityMin",
		val,
	)
}

func (j *jsiiProxy_DynamicHostVolume)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DynamicHostVolume)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DynamicHostVolume)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DynamicHostVolume)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DynamicHostVolume)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DynamicHostVolume)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DynamicHostVolume)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_DynamicHostVolume)SetNodeId(val *string) {
	if err := j.validateSetNodeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeId",
		val,
	)
}

func (j *jsiiProxy_DynamicHostVolume)SetNodePool(val *string) {
	if err := j.validateSetNodePoolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodePool",
		val,
	)
}

func (j *jsiiProxy_DynamicHostVolume)SetParameters(val *map[string]*string) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_DynamicHostVolume)SetPluginId(val *string) {
	if err := j.validateSetPluginIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pluginId",
		val,
	)
}

func (j *jsiiProxy_DynamicHostVolume)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DynamicHostVolume)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a DynamicHostVolume resource upon running "cdktf plan <stack-name>".
func DynamicHostVolume_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDynamicHostVolume_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-nomad.dynamicHostVolume.DynamicHostVolume",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func DynamicHostVolume_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDynamicHostVolume_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-nomad.dynamicHostVolume.DynamicHostVolume",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DynamicHostVolume_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDynamicHostVolume_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-nomad.dynamicHostVolume.DynamicHostVolume",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DynamicHostVolume_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDynamicHostVolume_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-nomad.dynamicHostVolume.DynamicHostVolume",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DynamicHostVolume_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-nomad.dynamicHostVolume.DynamicHostVolume",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DynamicHostVolume) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DynamicHostVolume) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DynamicHostVolume) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamicHostVolume) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamicHostVolume) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamicHostVolume) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamicHostVolume) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamicHostVolume) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamicHostVolume) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamicHostVolume) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamicHostVolume) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamicHostVolume) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamicHostVolume) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DynamicHostVolume) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamicHostVolume) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DynamicHostVolume) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DynamicHostVolume) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DynamicHostVolume) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DynamicHostVolume) PutCapability(value interface{}) {
	if err := d.validatePutCapabilityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCapability",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamicHostVolume) PutConstraint(value interface{}) {
	if err := d.validatePutConstraintParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putConstraint",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamicHostVolume) ResetCapacityMax() {
	_jsii_.InvokeVoid(
		d,
		"resetCapacityMax",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamicHostVolume) ResetCapacityMin() {
	_jsii_.InvokeVoid(
		d,
		"resetCapacityMin",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamicHostVolume) ResetConstraint() {
	_jsii_.InvokeVoid(
		d,
		"resetConstraint",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamicHostVolume) ResetNamespace() {
	_jsii_.InvokeVoid(
		d,
		"resetNamespace",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamicHostVolume) ResetNodeId() {
	_jsii_.InvokeVoid(
		d,
		"resetNodeId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamicHostVolume) ResetNodePool() {
	_jsii_.InvokeVoid(
		d,
		"resetNodePool",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamicHostVolume) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamicHostVolume) ResetParameters() {
	_jsii_.InvokeVoid(
		d,
		"resetParameters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamicHostVolume) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamicHostVolume) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamicHostVolume) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamicHostVolume) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamicHostVolume) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamicHostVolume) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

