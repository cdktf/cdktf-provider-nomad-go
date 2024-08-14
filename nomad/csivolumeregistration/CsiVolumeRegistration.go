// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package csivolumeregistration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-nomad-go/nomad/v9/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-nomad-go/nomad/v9/csivolumeregistration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.1/docs/resources/csi_volume_registration nomad_csi_volume_registration}.
type CsiVolumeRegistration interface {
	cdktf.TerraformResource
	Capability() CsiVolumeRegistrationCapabilityList
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
	MountOptions() CsiVolumeRegistrationMountOptionsOutputReference
	MountOptionsInput() *CsiVolumeRegistrationMountOptions
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
	Timeouts() CsiVolumeRegistrationTimeoutsOutputReference
	TimeoutsInput() interface{}
	Topologies() CsiVolumeRegistrationTopologiesList
	TopologyRequest() CsiVolumeRegistrationTopologyRequestOutputReference
	TopologyRequestInput() *CsiVolumeRegistrationTopologyRequest
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
	PutMountOptions(value *CsiVolumeRegistrationMountOptions)
	PutTimeouts(value *CsiVolumeRegistrationTimeouts)
	PutTopologyRequest(value *CsiVolumeRegistrationTopologyRequest)
	ResetCapability()
	ResetCapacityMax()
	ResetCapacityMin()
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

// The jsii proxy struct for CsiVolumeRegistration
type jsiiProxy_CsiVolumeRegistration struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CsiVolumeRegistration) Capability() CsiVolumeRegistrationCapabilityList {
	var returns CsiVolumeRegistrationCapabilityList
	_jsii_.Get(
		j,
		"capability",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) CapabilityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"capabilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) Capacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"capacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) CapacityMax() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityMax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) CapacityMaxBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"capacityMaxBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) CapacityMaxInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityMaxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) CapacityMin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityMin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) CapacityMinBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"capacityMinBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) CapacityMinInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityMinInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) Context() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"context",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) ContextInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"contextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) ControllerRequired() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"controllerRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) ControllersExpected() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"controllersExpected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) ControllersHealthy() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"controllersHealthy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) DeregisterOnDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deregisterOnDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) DeregisterOnDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deregisterOnDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) ExternalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) ExternalIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) MountOptions() CsiVolumeRegistrationMountOptionsOutputReference {
	var returns CsiVolumeRegistrationMountOptionsOutputReference
	_jsii_.Get(
		j,
		"mountOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) MountOptionsInput() *CsiVolumeRegistrationMountOptions {
	var returns *CsiVolumeRegistrationMountOptions
	_jsii_.Get(
		j,
		"mountOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) NodesExpected() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodesExpected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) NodesHealthy() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodesHealthy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) Parameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) ParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) PluginId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) PluginIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) PluginProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) PluginProviderVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginProviderVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) Schedulable() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"schedulable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) Secrets() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"secrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) SecretsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"secretsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) Timeouts() CsiVolumeRegistrationTimeoutsOutputReference {
	var returns CsiVolumeRegistrationTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) Topologies() CsiVolumeRegistrationTopologiesList {
	var returns CsiVolumeRegistrationTopologiesList
	_jsii_.Get(
		j,
		"topologies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) TopologyRequest() CsiVolumeRegistrationTopologyRequestOutputReference {
	var returns CsiVolumeRegistrationTopologyRequestOutputReference
	_jsii_.Get(
		j,
		"topologyRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) TopologyRequestInput() *CsiVolumeRegistrationTopologyRequest {
	var returns *CsiVolumeRegistrationTopologyRequest
	_jsii_.Get(
		j,
		"topologyRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) VolumeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsiVolumeRegistration) VolumeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.1/docs/resources/csi_volume_registration nomad_csi_volume_registration} Resource.
func NewCsiVolumeRegistration(scope constructs.Construct, id *string, config *CsiVolumeRegistrationConfig) CsiVolumeRegistration {
	_init_.Initialize()

	if err := validateNewCsiVolumeRegistrationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CsiVolumeRegistration{}

	_jsii_.Create(
		"@cdktf/provider-nomad.csiVolumeRegistration.CsiVolumeRegistration",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.1/docs/resources/csi_volume_registration nomad_csi_volume_registration} Resource.
func NewCsiVolumeRegistration_Override(c CsiVolumeRegistration, scope constructs.Construct, id *string, config *CsiVolumeRegistrationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-nomad.csiVolumeRegistration.CsiVolumeRegistration",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CsiVolumeRegistration)SetCapacityMax(val *string) {
	if err := j.validateSetCapacityMaxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacityMax",
		val,
	)
}

func (j *jsiiProxy_CsiVolumeRegistration)SetCapacityMin(val *string) {
	if err := j.validateSetCapacityMinParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacityMin",
		val,
	)
}

func (j *jsiiProxy_CsiVolumeRegistration)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CsiVolumeRegistration)SetContext(val *map[string]*string) {
	if err := j.validateSetContextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"context",
		val,
	)
}

func (j *jsiiProxy_CsiVolumeRegistration)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CsiVolumeRegistration)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CsiVolumeRegistration)SetDeregisterOnDestroy(val interface{}) {
	if err := j.validateSetDeregisterOnDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deregisterOnDestroy",
		val,
	)
}

func (j *jsiiProxy_CsiVolumeRegistration)SetExternalId(val *string) {
	if err := j.validateSetExternalIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalId",
		val,
	)
}

func (j *jsiiProxy_CsiVolumeRegistration)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CsiVolumeRegistration)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CsiVolumeRegistration)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CsiVolumeRegistration)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CsiVolumeRegistration)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_CsiVolumeRegistration)SetParameters(val *map[string]*string) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_CsiVolumeRegistration)SetPluginId(val *string) {
	if err := j.validateSetPluginIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pluginId",
		val,
	)
}

func (j *jsiiProxy_CsiVolumeRegistration)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CsiVolumeRegistration)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CsiVolumeRegistration)SetSecrets(val *map[string]*string) {
	if err := j.validateSetSecretsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secrets",
		val,
	)
}

func (j *jsiiProxy_CsiVolumeRegistration)SetVolumeId(val *string) {
	if err := j.validateSetVolumeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeId",
		val,
	)
}

// Generates CDKTF code for importing a CsiVolumeRegistration resource upon running "cdktf plan <stack-name>".
func CsiVolumeRegistration_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateCsiVolumeRegistration_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-nomad.csiVolumeRegistration.CsiVolumeRegistration",
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
func CsiVolumeRegistration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCsiVolumeRegistration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-nomad.csiVolumeRegistration.CsiVolumeRegistration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CsiVolumeRegistration_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCsiVolumeRegistration_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-nomad.csiVolumeRegistration.CsiVolumeRegistration",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CsiVolumeRegistration_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCsiVolumeRegistration_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-nomad.csiVolumeRegistration.CsiVolumeRegistration",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CsiVolumeRegistration_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-nomad.csiVolumeRegistration.CsiVolumeRegistration",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CsiVolumeRegistration) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_CsiVolumeRegistration) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CsiVolumeRegistration) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CsiVolumeRegistration) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CsiVolumeRegistration) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CsiVolumeRegistration) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CsiVolumeRegistration) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CsiVolumeRegistration) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CsiVolumeRegistration) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CsiVolumeRegistration) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CsiVolumeRegistration) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CsiVolumeRegistration) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CsiVolumeRegistration) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_CsiVolumeRegistration) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CsiVolumeRegistration) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CsiVolumeRegistration) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_CsiVolumeRegistration) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CsiVolumeRegistration) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CsiVolumeRegistration) PutCapability(value interface{}) {
	if err := c.validatePutCapabilityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCapability",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CsiVolumeRegistration) PutMountOptions(value *CsiVolumeRegistrationMountOptions) {
	if err := c.validatePutMountOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMountOptions",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CsiVolumeRegistration) PutTimeouts(value *CsiVolumeRegistrationTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CsiVolumeRegistration) PutTopologyRequest(value *CsiVolumeRegistrationTopologyRequest) {
	if err := c.validatePutTopologyRequestParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTopologyRequest",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CsiVolumeRegistration) ResetCapability() {
	_jsii_.InvokeVoid(
		c,
		"resetCapability",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CsiVolumeRegistration) ResetCapacityMax() {
	_jsii_.InvokeVoid(
		c,
		"resetCapacityMax",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CsiVolumeRegistration) ResetCapacityMin() {
	_jsii_.InvokeVoid(
		c,
		"resetCapacityMin",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CsiVolumeRegistration) ResetContext() {
	_jsii_.InvokeVoid(
		c,
		"resetContext",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CsiVolumeRegistration) ResetDeregisterOnDestroy() {
	_jsii_.InvokeVoid(
		c,
		"resetDeregisterOnDestroy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CsiVolumeRegistration) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CsiVolumeRegistration) ResetMountOptions() {
	_jsii_.InvokeVoid(
		c,
		"resetMountOptions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CsiVolumeRegistration) ResetNamespace() {
	_jsii_.InvokeVoid(
		c,
		"resetNamespace",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CsiVolumeRegistration) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CsiVolumeRegistration) ResetParameters() {
	_jsii_.InvokeVoid(
		c,
		"resetParameters",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CsiVolumeRegistration) ResetSecrets() {
	_jsii_.InvokeVoid(
		c,
		"resetSecrets",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CsiVolumeRegistration) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CsiVolumeRegistration) ResetTopologyRequest() {
	_jsii_.InvokeVoid(
		c,
		"resetTopologyRequest",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CsiVolumeRegistration) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CsiVolumeRegistration) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CsiVolumeRegistration) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CsiVolumeRegistration) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CsiVolumeRegistration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CsiVolumeRegistration) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

