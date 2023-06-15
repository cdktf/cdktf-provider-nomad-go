package externalvolume

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-nomad-go/nomad/v5/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-nomad-go/nomad/v5/externalvolume/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/nomad/1.4.20/docs/resources/external_volume nomad_external_volume}.
type ExternalVolume interface {
	cdktf.TerraformResource
	Capability() ExternalVolumeCapabilityList
	CapabilityInput() interface{}
	CapacityMax() *string
	SetCapacityMax(val *string)
	CapacityMaxInput() *string
	CapacityMin() *string
	SetCapacityMin(val *string)
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
	MountOptions() ExternalVolumeMountOptionsOutputReference
	MountOptionsInput() *ExternalVolumeMountOptions
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
	Topologies() ExternalVolumeTopologiesList
	TopologyRequest() ExternalVolumeTopologyRequestOutputReference
	TopologyRequestInput() *ExternalVolumeTopologyRequest
	Type() *string
	SetType(val *string)
	TypeInput() *string
	VolumeId() *string
	SetVolumeId(val *string)
	VolumeIdInput() *string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutCapability(value interface{})
	PutMountOptions(value *ExternalVolumeMountOptions)
	PutTopologyRequest(value *ExternalVolumeTopologyRequest)
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
	ResetTopologyRequest()
	ResetType()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ExternalVolume
type jsiiProxy_ExternalVolume struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ExternalVolume) Capability() ExternalVolumeCapabilityList {
	var returns ExternalVolumeCapabilityList
	_jsii_.Get(
		j,
		"capability",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) CapabilityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"capabilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) CapacityMax() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityMax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) CapacityMaxInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityMaxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) CapacityMin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityMin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) CapacityMinInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityMinInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) CloneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) CloneIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloneIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) ControllerRequired() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"controllerRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) ControllersExpected() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"controllersExpected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) ControllersHealthy() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"controllersHealthy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) MountOptions() ExternalVolumeMountOptionsOutputReference {
	var returns ExternalVolumeMountOptionsOutputReference
	_jsii_.Get(
		j,
		"mountOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) MountOptionsInput() *ExternalVolumeMountOptions {
	var returns *ExternalVolumeMountOptions
	_jsii_.Get(
		j,
		"mountOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) NodesExpected() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodesExpected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) NodesHealthy() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodesHealthy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) Parameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) ParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) PluginId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) PluginIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) PluginProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) PluginProviderVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginProviderVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) Schedulable() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"schedulable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) Secrets() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"secrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) SecretsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"secretsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) SnapshotId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) SnapshotIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) Topologies() ExternalVolumeTopologiesList {
	var returns ExternalVolumeTopologiesList
	_jsii_.Get(
		j,
		"topologies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) TopologyRequest() ExternalVolumeTopologyRequestOutputReference {
	var returns ExternalVolumeTopologyRequestOutputReference
	_jsii_.Get(
		j,
		"topologyRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) TopologyRequestInput() *ExternalVolumeTopologyRequest {
	var returns *ExternalVolumeTopologyRequest
	_jsii_.Get(
		j,
		"topologyRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) VolumeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolume) VolumeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/nomad/1.4.20/docs/resources/external_volume nomad_external_volume} Resource.
func NewExternalVolume(scope constructs.Construct, id *string, config *ExternalVolumeConfig) ExternalVolume {
	_init_.Initialize()

	if err := validateNewExternalVolumeParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ExternalVolume{}

	_jsii_.Create(
		"@cdktf/provider-nomad.externalVolume.ExternalVolume",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/nomad/1.4.20/docs/resources/external_volume nomad_external_volume} Resource.
func NewExternalVolume_Override(e ExternalVolume, scope constructs.Construct, id *string, config *ExternalVolumeConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-nomad.externalVolume.ExternalVolume",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_ExternalVolume)SetCapacityMax(val *string) {
	if err := j.validateSetCapacityMaxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacityMax",
		val,
	)
}

func (j *jsiiProxy_ExternalVolume)SetCapacityMin(val *string) {
	if err := j.validateSetCapacityMinParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacityMin",
		val,
	)
}

func (j *jsiiProxy_ExternalVolume)SetCloneId(val *string) {
	if err := j.validateSetCloneIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloneId",
		val,
	)
}

func (j *jsiiProxy_ExternalVolume)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ExternalVolume)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ExternalVolume)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ExternalVolume)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ExternalVolume)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ExternalVolume)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ExternalVolume)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ExternalVolume)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_ExternalVolume)SetParameters(val *map[string]*string) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_ExternalVolume)SetPluginId(val *string) {
	if err := j.validateSetPluginIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pluginId",
		val,
	)
}

func (j *jsiiProxy_ExternalVolume)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ExternalVolume)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ExternalVolume)SetSecrets(val *map[string]*string) {
	if err := j.validateSetSecretsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secrets",
		val,
	)
}

func (j *jsiiProxy_ExternalVolume)SetSnapshotId(val *string) {
	if err := j.validateSetSnapshotIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotId",
		val,
	)
}

func (j *jsiiProxy_ExternalVolume)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_ExternalVolume)SetVolumeId(val *string) {
	if err := j.validateSetVolumeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeId",
		val,
	)
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
func ExternalVolume_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateExternalVolume_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-nomad.externalVolume.ExternalVolume",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ExternalVolume_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateExternalVolume_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-nomad.externalVolume.ExternalVolume",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ExternalVolume_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateExternalVolume_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-nomad.externalVolume.ExternalVolume",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ExternalVolume_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-nomad.externalVolume.ExternalVolume",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_ExternalVolume) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_ExternalVolume) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalVolume) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalVolume) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalVolume) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalVolume) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalVolume) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalVolume) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalVolume) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalVolume) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalVolume) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalVolume) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_ExternalVolume) PutCapability(value interface{}) {
	if err := e.validatePutCapabilityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putCapability",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ExternalVolume) PutMountOptions(value *ExternalVolumeMountOptions) {
	if err := e.validatePutMountOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putMountOptions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ExternalVolume) PutTopologyRequest(value *ExternalVolumeTopologyRequest) {
	if err := e.validatePutTopologyRequestParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTopologyRequest",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ExternalVolume) ResetCapacityMax() {
	_jsii_.InvokeVoid(
		e,
		"resetCapacityMax",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalVolume) ResetCapacityMin() {
	_jsii_.InvokeVoid(
		e,
		"resetCapacityMin",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalVolume) ResetCloneId() {
	_jsii_.InvokeVoid(
		e,
		"resetCloneId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalVolume) ResetId() {
	_jsii_.InvokeVoid(
		e,
		"resetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalVolume) ResetMountOptions() {
	_jsii_.InvokeVoid(
		e,
		"resetMountOptions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalVolume) ResetNamespace() {
	_jsii_.InvokeVoid(
		e,
		"resetNamespace",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalVolume) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalVolume) ResetParameters() {
	_jsii_.InvokeVoid(
		e,
		"resetParameters",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalVolume) ResetSecrets() {
	_jsii_.InvokeVoid(
		e,
		"resetSecrets",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalVolume) ResetSnapshotId() {
	_jsii_.InvokeVoid(
		e,
		"resetSnapshotId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalVolume) ResetTopologyRequest() {
	_jsii_.InvokeVoid(
		e,
		"resetTopologyRequest",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalVolume) ResetType() {
	_jsii_.InvokeVoid(
		e,
		"resetType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalVolume) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalVolume) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalVolume) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalVolume) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

