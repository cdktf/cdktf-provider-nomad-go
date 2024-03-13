// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-nomad-go/nomad/v9/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-nomad-go/nomad/v9/job/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/nomad/2.2.0/docs/resources/job nomad_job}.
type Job interface {
	cdktf.TerraformResource
	AllocationIds() *[]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ConsulToken() *string
	SetConsulToken(val *string)
	ConsulTokenInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	Datacenters() *[]*string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DeploymentId() *string
	DeploymentStatus() *string
	DeregisterOnDestroy() interface{}
	SetDeregisterOnDestroy(val interface{})
	DeregisterOnDestroyInput() interface{}
	DeregisterOnIdChange() interface{}
	SetDeregisterOnIdChange(val interface{})
	DeregisterOnIdChangeInput() interface{}
	Detach() interface{}
	SetDetach(val interface{})
	DetachInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Hcl1() interface{}
	SetHcl1(val interface{})
	Hcl1Input() interface{}
	Hcl2() JobHcl2OutputReference
	Hcl2Input() *JobHcl2
	Id() *string
	SetId(val *string)
	IdInput() *string
	Jobspec() *string
	SetJobspec(val *string)
	JobspecInput() *string
	Json() interface{}
	SetJson(val interface{})
	JsonInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	ModifyIndex() *string
	Name() *string
	Namespace() *string
	// The tree node.
	Node() constructs.Node
	PolicyOverride() interface{}
	SetPolicyOverride(val interface{})
	PolicyOverrideInput() interface{}
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PurgeOnDestroy() interface{}
	SetPurgeOnDestroy(val interface{})
	PurgeOnDestroyInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	ReadAllocationIds() interface{}
	SetReadAllocationIds(val interface{})
	ReadAllocationIdsInput() interface{}
	Region() *string
	RerunIfDead() interface{}
	SetRerunIfDead(val interface{})
	RerunIfDeadInput() interface{}
	Status() *string
	TaskGroups() JobTaskGroupsList
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() JobTimeoutsOutputReference
	TimeoutsInput() interface{}
	Type() *string
	VaultToken() *string
	SetVaultToken(val *string)
	VaultTokenInput() *string
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
	PutHcl2(value *JobHcl2)
	PutTimeouts(value *JobTimeouts)
	ResetConsulToken()
	ResetDeregisterOnDestroy()
	ResetDeregisterOnIdChange()
	ResetDetach()
	ResetHcl1()
	ResetHcl2()
	ResetId()
	ResetJson()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPolicyOverride()
	ResetPurgeOnDestroy()
	ResetReadAllocationIds()
	ResetRerunIfDead()
	ResetTimeouts()
	ResetVaultToken()
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

// The jsii proxy struct for Job
type jsiiProxy_Job struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Job) AllocationIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allocationIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) ConsulToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consulToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) ConsulTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consulTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Datacenters() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"datacenters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) DeploymentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) DeploymentStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) DeregisterOnDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deregisterOnDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) DeregisterOnDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deregisterOnDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) DeregisterOnIdChange() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deregisterOnIdChange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) DeregisterOnIdChangeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deregisterOnIdChangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Detach() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"detach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) DetachInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"detachInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Hcl1() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hcl1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Hcl1Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hcl1Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Hcl2() JobHcl2OutputReference {
	var returns JobHcl2OutputReference
	_jsii_.Get(
		j,
		"hcl2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Hcl2Input() *JobHcl2 {
	var returns *JobHcl2
	_jsii_.Get(
		j,
		"hcl2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Jobspec() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobspec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) JobspecInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobspecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Json() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"json",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) JsonInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) ModifyIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modifyIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) PolicyOverride() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"policyOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) PolicyOverrideInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"policyOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) PurgeOnDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"purgeOnDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) PurgeOnDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"purgeOnDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) ReadAllocationIds() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readAllocationIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) ReadAllocationIdsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readAllocationIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) RerunIfDead() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rerunIfDead",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) RerunIfDeadInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rerunIfDeadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) TaskGroups() JobTaskGroupsList {
	var returns JobTaskGroupsList
	_jsii_.Get(
		j,
		"taskGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Timeouts() JobTimeoutsOutputReference {
	var returns JobTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) VaultToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vaultToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) VaultTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vaultTokenInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/nomad/2.2.0/docs/resources/job nomad_job} Resource.
func NewJob(scope constructs.Construct, id *string, config *JobConfig) Job {
	_init_.Initialize()

	if err := validateNewJobParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Job{}

	_jsii_.Create(
		"@cdktf/provider-nomad.job.Job",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/nomad/2.2.0/docs/resources/job nomad_job} Resource.
func NewJob_Override(j Job, scope constructs.Construct, id *string, config *JobConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-nomad.job.Job",
		[]interface{}{scope, id, config},
		j,
	)
}

func (j *jsiiProxy_Job)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Job)SetConsulToken(val *string) {
	if err := j.validateSetConsulTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consulToken",
		val,
	)
}

func (j *jsiiProxy_Job)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Job)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Job)SetDeregisterOnDestroy(val interface{}) {
	if err := j.validateSetDeregisterOnDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deregisterOnDestroy",
		val,
	)
}

func (j *jsiiProxy_Job)SetDeregisterOnIdChange(val interface{}) {
	if err := j.validateSetDeregisterOnIdChangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deregisterOnIdChange",
		val,
	)
}

func (j *jsiiProxy_Job)SetDetach(val interface{}) {
	if err := j.validateSetDetachParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"detach",
		val,
	)
}

func (j *jsiiProxy_Job)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Job)SetHcl1(val interface{}) {
	if err := j.validateSetHcl1Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hcl1",
		val,
	)
}

func (j *jsiiProxy_Job)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Job)SetJobspec(val *string) {
	if err := j.validateSetJobspecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobspec",
		val,
	)
}

func (j *jsiiProxy_Job)SetJson(val interface{}) {
	if err := j.validateSetJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"json",
		val,
	)
}

func (j *jsiiProxy_Job)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Job)SetPolicyOverride(val interface{}) {
	if err := j.validateSetPolicyOverrideParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyOverride",
		val,
	)
}

func (j *jsiiProxy_Job)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Job)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Job)SetPurgeOnDestroy(val interface{}) {
	if err := j.validateSetPurgeOnDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"purgeOnDestroy",
		val,
	)
}

func (j *jsiiProxy_Job)SetReadAllocationIds(val interface{}) {
	if err := j.validateSetReadAllocationIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readAllocationIds",
		val,
	)
}

func (j *jsiiProxy_Job)SetRerunIfDead(val interface{}) {
	if err := j.validateSetRerunIfDeadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rerunIfDead",
		val,
	)
}

func (j *jsiiProxy_Job)SetVaultToken(val *string) {
	if err := j.validateSetVaultTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vaultToken",
		val,
	)
}

// Generates CDKTF code for importing a Job resource upon running "cdktf plan <stack-name>".
func Job_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateJob_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-nomad.job.Job",
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
func Job_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateJob_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-nomad.job.Job",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Job_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateJob_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-nomad.job.Job",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Job_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateJob_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-nomad.job.Job",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Job_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-nomad.job.Job",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) AddMoveTarget(moveTarget *string) {
	if err := j.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (j *jsiiProxy_Job) AddOverride(path *string, value interface{}) {
	if err := j.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (j *jsiiProxy_Job) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := j.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		j,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := j.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := j.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		j,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := j.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		j,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := j.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		j,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := j.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		j,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := j.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		j,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) GetStringAttribute(terraformAttribute *string) *string {
	if err := j.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		j,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := j.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		j,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		j,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := j.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (j *jsiiProxy_Job) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := j.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) MoveFromId(id *string) {
	if err := j.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"moveFromId",
		[]interface{}{id},
	)
}

func (j *jsiiProxy_Job) MoveTo(moveTarget *string, index interface{}) {
	if err := j.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (j *jsiiProxy_Job) MoveToId(id *string) {
	if err := j.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"moveToId",
		[]interface{}{id},
	)
}

func (j *jsiiProxy_Job) OverrideLogicalId(newLogicalId *string) {
	if err := j.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (j *jsiiProxy_Job) PutHcl2(value *JobHcl2) {
	if err := j.validatePutHcl2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putHcl2",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_Job) PutTimeouts(value *JobTimeouts) {
	if err := j.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_Job) ResetConsulToken() {
	_jsii_.InvokeVoid(
		j,
		"resetConsulToken",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetDeregisterOnDestroy() {
	_jsii_.InvokeVoid(
		j,
		"resetDeregisterOnDestroy",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetDeregisterOnIdChange() {
	_jsii_.InvokeVoid(
		j,
		"resetDeregisterOnIdChange",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetDetach() {
	_jsii_.InvokeVoid(
		j,
		"resetDetach",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetHcl1() {
	_jsii_.InvokeVoid(
		j,
		"resetHcl1",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetHcl2() {
	_jsii_.InvokeVoid(
		j,
		"resetHcl2",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetId() {
	_jsii_.InvokeVoid(
		j,
		"resetId",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetJson() {
	_jsii_.InvokeVoid(
		j,
		"resetJson",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		j,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetPolicyOverride() {
	_jsii_.InvokeVoid(
		j,
		"resetPolicyOverride",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetPurgeOnDestroy() {
	_jsii_.InvokeVoid(
		j,
		"resetPurgeOnDestroy",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetReadAllocationIds() {
	_jsii_.InvokeVoid(
		j,
		"resetReadAllocationIds",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetRerunIfDead() {
	_jsii_.InvokeVoid(
		j,
		"resetRerunIfDead",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetTimeouts() {
	_jsii_.InvokeVoid(
		j,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) ResetVaultToken() {
	_jsii_.InvokeVoid(
		j,
		"resetVaultToken",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		j,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		j,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		j,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		j,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		j,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

