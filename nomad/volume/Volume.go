// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package volume

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-nomad-go/nomad/v9/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-nomad-go/nomad/v9/volume/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.1/docs/resources/volume nomad_volume}.
type Volume interface {
	cdktf.TerraformResource
	AccessMode() *string
	SetAccessMode(val *string)
	AccessModeInput() *string
	AttachmentMode() *string
	SetAttachmentMode(val *string)
	AttachmentModeInput() *string
	Capability() VolumeCapabilityList
	CapabilityInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	Context() *map[string]*string
	SetContext(val *map[string]*string)
	ContextInput() *map[string]*string
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
	DeregisterOnDestroy() interface{}
	SetDeregisterOnDestroy(val interface{})
	DeregisterOnDestroyInput() interface{}
	ExternalId() *string
	SetExternalId(val *string)
	ExternalIdInput() *string
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
	MountOptions() VolumeMountOptionsOutputReference
	MountOptionsInput() *VolumeMountOptions
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Topologies() VolumeTopologiesList
	TopologyRequest() VolumeTopologyRequestOutputReference
	TopologyRequestInput() *VolumeTopologyRequest
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	PutMountOptions(value *VolumeMountOptions)
	PutTopologyRequest(value *VolumeTopologyRequest)
	ResetAccessMode()
	ResetAttachmentMode()
	ResetCapability()
	ResetContext()
	ResetDeregisterOnDestroy()
	ResetId()
	ResetMountOptions()
	ResetNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParameters()
	ResetSecrets()
	ResetTopologyRequest()
	ResetType()
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

// The jsii proxy struct for Volume
type jsiiProxy_Volume struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Volume) AccessMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) AccessModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) AttachmentMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attachmentMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) AttachmentModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attachmentModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) Capability() VolumeCapabilityList {
	var returns VolumeCapabilityList
	_jsii_.Get(
		j,
		"capability",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) CapabilityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"capabilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) Context() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"context",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) ContextInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"contextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) ControllerRequired() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"controllerRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) ControllersExpected() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"controllersExpected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) ControllersHealthy() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"controllersHealthy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) DeregisterOnDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deregisterOnDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) DeregisterOnDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deregisterOnDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) ExternalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) ExternalIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) MountOptions() VolumeMountOptionsOutputReference {
	var returns VolumeMountOptionsOutputReference
	_jsii_.Get(
		j,
		"mountOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) MountOptionsInput() *VolumeMountOptions {
	var returns *VolumeMountOptions
	_jsii_.Get(
		j,
		"mountOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) NodesExpected() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodesExpected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) NodesHealthy() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodesHealthy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) Parameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) ParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) PluginId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) PluginIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) PluginProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) PluginProviderVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginProviderVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) Schedulable() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"schedulable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) Secrets() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"secrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) SecretsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"secretsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) Topologies() VolumeTopologiesList {
	var returns VolumeTopologiesList
	_jsii_.Get(
		j,
		"topologies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) TopologyRequest() VolumeTopologyRequestOutputReference {
	var returns VolumeTopologyRequestOutputReference
	_jsii_.Get(
		j,
		"topologyRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) TopologyRequestInput() *VolumeTopologyRequest {
	var returns *VolumeTopologyRequest
	_jsii_.Get(
		j,
		"topologyRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) VolumeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) VolumeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.1/docs/resources/volume nomad_volume} Resource.
func NewVolume(scope constructs.Construct, id *string, config *VolumeConfig) Volume {
	_init_.Initialize()

	if err := validateNewVolumeParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Volume{}

	_jsii_.Create(
		"@cdktf/provider-nomad.volume.Volume",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.1/docs/resources/volume nomad_volume} Resource.
func NewVolume_Override(v Volume, scope constructs.Construct, id *string, config *VolumeConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-nomad.volume.Volume",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_Volume)SetAccessMode(val *string) {
	if err := j.validateSetAccessModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessMode",
		val,
	)
}

func (j *jsiiProxy_Volume)SetAttachmentMode(val *string) {
	if err := j.validateSetAttachmentModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attachmentMode",
		val,
	)
}

func (j *jsiiProxy_Volume)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Volume)SetContext(val *map[string]*string) {
	if err := j.validateSetContextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"context",
		val,
	)
}

func (j *jsiiProxy_Volume)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Volume)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Volume)SetDeregisterOnDestroy(val interface{}) {
	if err := j.validateSetDeregisterOnDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deregisterOnDestroy",
		val,
	)
}

func (j *jsiiProxy_Volume)SetExternalId(val *string) {
	if err := j.validateSetExternalIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalId",
		val,
	)
}

func (j *jsiiProxy_Volume)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Volume)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Volume)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Volume)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Volume)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_Volume)SetParameters(val *map[string]*string) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_Volume)SetPluginId(val *string) {
	if err := j.validateSetPluginIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pluginId",
		val,
	)
}

func (j *jsiiProxy_Volume)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Volume)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Volume)SetSecrets(val *map[string]*string) {
	if err := j.validateSetSecretsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secrets",
		val,
	)
}

func (j *jsiiProxy_Volume)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_Volume)SetVolumeId(val *string) {
	if err := j.validateSetVolumeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeId",
		val,
	)
}

// Generates CDKTF code for importing a Volume resource upon running "cdktf plan <stack-name>".
func Volume_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateVolume_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-nomad.volume.Volume",
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
func Volume_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVolume_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-nomad.volume.Volume",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Volume_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVolume_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-nomad.volume.Volume",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Volume_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVolume_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-nomad.volume.Volume",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Volume_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-nomad.volume.Volume",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_Volume) AddMoveTarget(moveTarget *string) {
	if err := v.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (v *jsiiProxy_Volume) AddOverride(path *string, value interface{}) {
	if err := v.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_Volume) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_Volume) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_Volume) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_Volume) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_Volume) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_Volume) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_Volume) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_Volume) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_Volume) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_Volume) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_Volume) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := v.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (v *jsiiProxy_Volume) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_Volume) MoveFromId(id *string) {
	if err := v.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveFromId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_Volume) MoveTo(moveTarget *string, index interface{}) {
	if err := v.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (v *jsiiProxy_Volume) MoveToId(id *string) {
	if err := v.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveToId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_Volume) OverrideLogicalId(newLogicalId *string) {
	if err := v.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_Volume) PutCapability(value interface{}) {
	if err := v.validatePutCapabilityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putCapability",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_Volume) PutMountOptions(value *VolumeMountOptions) {
	if err := v.validatePutMountOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putMountOptions",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_Volume) PutTopologyRequest(value *VolumeTopologyRequest) {
	if err := v.validatePutTopologyRequestParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putTopologyRequest",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_Volume) ResetAccessMode() {
	_jsii_.InvokeVoid(
		v,
		"resetAccessMode",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Volume) ResetAttachmentMode() {
	_jsii_.InvokeVoid(
		v,
		"resetAttachmentMode",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Volume) ResetCapability() {
	_jsii_.InvokeVoid(
		v,
		"resetCapability",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Volume) ResetContext() {
	_jsii_.InvokeVoid(
		v,
		"resetContext",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Volume) ResetDeregisterOnDestroy() {
	_jsii_.InvokeVoid(
		v,
		"resetDeregisterOnDestroy",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Volume) ResetId() {
	_jsii_.InvokeVoid(
		v,
		"resetId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Volume) ResetMountOptions() {
	_jsii_.InvokeVoid(
		v,
		"resetMountOptions",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Volume) ResetNamespace() {
	_jsii_.InvokeVoid(
		v,
		"resetNamespace",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Volume) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Volume) ResetParameters() {
	_jsii_.InvokeVoid(
		v,
		"resetParameters",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Volume) ResetSecrets() {
	_jsii_.InvokeVoid(
		v,
		"resetSecrets",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Volume) ResetTopologyRequest() {
	_jsii_.InvokeVoid(
		v,
		"resetTopologyRequest",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Volume) ResetType() {
	_jsii_.InvokeVoid(
		v,
		"resetType",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Volume) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_Volume) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_Volume) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_Volume) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_Volume) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_Volume) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

