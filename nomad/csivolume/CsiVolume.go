// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package csivolume

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-nomad-go/nomad/v9/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-nomad-go/nomad/v9/csivolume/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.1/docs/resources/csi_volume nomad_csi_volume}.
type CsiVolume interface {
	cdktf.TerraformResource
	Capability() CsiVolumeCapabilityList
	CapabilityInput() interface{}
	Capacity() *float64
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
	CloneId() *string
	SetCloneId(val *string)
	CloneIdInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ControllerRequired() cdktf.IResolvable
	ControllersExpected() *float64
	ControllersHealthy() *float64
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	ExternalId() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MountOptions() CsiVolumeMountOptionsOutputReference
	MountOptionsInput() *CsiVolumeMountOptions
	Name() *string
	SetName(val *string)
	NameInput() *string
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	// The tree node.
	Node() constructs.Node
	NodesExpected() *float64
	NodesHealthy() *float64
	Parameters() *map[string]*string
	SetParameters(val *map[string]*string)
	ParametersInput() *map[string]*string
	PluginId() *string
	SetPluginId(val *string)
	PluginIdInput() *string
	PluginProvider() *string
	PluginProviderVersion() *string
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
	Schedulable() cdktf.IResolvable
	Secrets() *map[string]*string
	SetSecrets(val *map[string]*string)
	SecretsInput() *map[string]*string
	SnapshotId() *string
	SetSnapshotId(val *string)
	SnapshotIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() CsiVolumeTimeoutsOutputReference
	TimeoutsInput() interface{}
	Topologies() CsiVolumeTopologiesList
	TopologyRequest() CsiVolumeTopologyRequestOutputReference
	TopologyRequestInput() *CsiVolumeTopologyRequest
	VolumeId() *string
	SetVolumeId(val *string)
	VolumeIdInput() *string
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
	PutMountOptions(value *CsiVolumeMountOptions)
	PutTimeouts(value *CsiVolumeTimeouts)
	PutTopologyRequest(value *CsiVolumeTopologyRequest)
	ResetCapacityMax()
	ResetCapacityMin()
	ResetCloneId()
	ResetId()
	ResetMountOptions()
	ResetNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParameters()
	ResetSecrets()
	ResetSnapshotId()
	ResetTimeouts()
	ResetTopologyRequest()
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

// The jsii proxy struct for CsiVolume
type jsiiProxy_CsiVolume struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CsiVolume) Capability() CsiVolumeCapabilityList {
	var returns CsiVolumeCapabilityList
	_jsii_.Get(
		j,
		"capability",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) CapabilityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"capabilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) Capacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"capacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) CapacityMax() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityMax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) CapacityMaxBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"capacityMaxBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) CapacityMaxInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityMaxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) CapacityMin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityMin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) CapacityMinBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"capacityMinBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) CapacityMinInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityMinInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) CloneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) CloneIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloneIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) ControllerRequired() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"controllerRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) ControllersExpected() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"controllersExpected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) ControllersHealthy() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"controllersHealthy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) ExternalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) MountOptions() CsiVolumeMountOptionsOutputReference {
	var returns CsiVolumeMountOptionsOutputReference
	_jsii_.Get(
		j,
		"mountOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) MountOptionsInput() *CsiVolumeMountOptions {
	var returns *CsiVolumeMountOptions
	_jsii_.Get(
		j,
		"mountOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) NodesExpected() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodesExpected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) NodesHealthy() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodesHealthy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) Parameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) ParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) PluginId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) PluginIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) PluginProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) PluginProviderVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginProviderVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) Schedulable() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"schedulable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) Secrets() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"secrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) SecretsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"secretsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) SnapshotId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) SnapshotIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) Timeouts() CsiVolumeTimeoutsOutputReference {
	var returns CsiVolumeTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) Topologies() CsiVolumeTopologiesList {
	var returns CsiVolumeTopologiesList
	_jsii_.Get(
		j,
		"topologies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) TopologyRequest() CsiVolumeTopologyRequestOutputReference {
	var returns CsiVolumeTopologyRequestOutputReference
	_jsii_.Get(
		j,
		"topologyRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) TopologyRequestInput() *CsiVolumeTopologyRequest {
	var returns *CsiVolumeTopologyRequest
	_jsii_.Get(
		j,
		"topologyRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) VolumeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolume) VolumeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.1/docs/resources/csi_volume nomad_csi_volume} Resource.
func NewCsiVolume(scope constructs.Construct, id *string, config *CsiVolumeConfig) CsiVolume {
	_init_.Initialize()

	if err := validateNewCsiVolumeParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CsiVolume{}

	_jsii_.Create(
		"@cdktf/provider-nomad.csiVolume.CsiVolume",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.1/docs/resources/csi_volume nomad_csi_volume} Resource.
func NewCsiVolume_Override(c CsiVolume, scope constructs.Construct, id *string, config *CsiVolumeConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-nomad.csiVolume.CsiVolume",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CsiVolume)SetCapacityMax(val *string) {
	if err := j.validateSetCapacityMaxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacityMax",
		val,
	)
}

func (j *jsiiProxy_CsiVolume)SetCapacityMin(val *string) {
	if err := j.validateSetCapacityMinParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacityMin",
		val,
	)
}

func (j *jsiiProxy_CsiVolume)SetCloneId(val *string) {
	if err := j.validateSetCloneIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloneId",
		val,
	)
}

func (j *jsiiProxy_CsiVolume)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CsiVolume)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CsiVolume)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CsiVolume)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CsiVolume)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CsiVolume)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CsiVolume)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CsiVolume)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_CsiVolume)SetParameters(val *map[string]*string) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_CsiVolume)SetPluginId(val *string) {
	if err := j.validateSetPluginIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pluginId",
		val,
	)
}

func (j *jsiiProxy_CsiVolume)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CsiVolume)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CsiVolume)SetSecrets(val *map[string]*string) {
	if err := j.validateSetSecretsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secrets",
		val,
	)
}

func (j *jsiiProxy_CsiVolume)SetSnapshotId(val *string) {
	if err := j.validateSetSnapshotIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotId",
		val,
	)
}

func (j *jsiiProxy_CsiVolume)SetVolumeId(val *string) {
	if err := j.validateSetVolumeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeId",
		val,
	)
}

// Generates CDKTF code for importing a CsiVolume resource upon running "cdktf plan <stack-name>".
func CsiVolume_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateCsiVolume_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-nomad.csiVolume.CsiVolume",
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
func CsiVolume_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCsiVolume_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-nomad.csiVolume.CsiVolume",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CsiVolume_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCsiVolume_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-nomad.csiVolume.CsiVolume",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CsiVolume_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCsiVolume_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-nomad.csiVolume.CsiVolume",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CsiVolume_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-nomad.csiVolume.CsiVolume",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CsiVolume) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_CsiVolume) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CsiVolume) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CsiVolume) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CsiVolume) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CsiVolume) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CsiVolume) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CsiVolume) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CsiVolume) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CsiVolume) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CsiVolume) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CsiVolume) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CsiVolume) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_CsiVolume) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CsiVolume) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CsiVolume) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_CsiVolume) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CsiVolume) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CsiVolume) PutCapability(value interface{}) {
	if err := c.validatePutCapabilityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCapability",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CsiVolume) PutMountOptions(value *CsiVolumeMountOptions) {
	if err := c.validatePutMountOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMountOptions",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CsiVolume) PutTimeouts(value *CsiVolumeTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CsiVolume) PutTopologyRequest(value *CsiVolumeTopologyRequest) {
	if err := c.validatePutTopologyRequestParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTopologyRequest",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CsiVolume) ResetCapacityMax() {
	_jsii_.InvokeVoid(
		c,
		"resetCapacityMax",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CsiVolume) ResetCapacityMin() {
	_jsii_.InvokeVoid(
		c,
		"resetCapacityMin",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CsiVolume) ResetCloneId() {
	_jsii_.InvokeVoid(
		c,
		"resetCloneId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CsiVolume) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CsiVolume) ResetMountOptions() {
	_jsii_.InvokeVoid(
		c,
		"resetMountOptions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CsiVolume) ResetNamespace() {
	_jsii_.InvokeVoid(
		c,
		"resetNamespace",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CsiVolume) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CsiVolume) ResetParameters() {
	_jsii_.InvokeVoid(
		c,
		"resetParameters",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CsiVolume) ResetSecrets() {
	_jsii_.InvokeVoid(
		c,
		"resetSecrets",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CsiVolume) ResetSnapshotId() {
	_jsii_.InvokeVoid(
		c,
		"resetSnapshotId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CsiVolume) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CsiVolume) ResetTopologyRequest() {
	_jsii_.InvokeVoid(
		c,
		"resetTopologyRequest",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CsiVolume) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CsiVolume) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CsiVolume) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CsiVolume) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CsiVolume) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CsiVolume) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

