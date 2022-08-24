// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type InstancePlacementGroup struct {
	pulumi.CustomResourceState

	// The name of the placement group
	Name pulumi.StringOutput `pulumi:"name"`
	// The organization_id you want to attach the resource to
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// One of the two policy_mode may be selected: enforced or optional.
	PolicyMode pulumi.StringPtrOutput `pulumi:"policyMode"`
	// Is true when the policy is respected.
	PolicyRespected pulumi.BoolOutput `pulumi:"policyRespected"`
	// The operating mode is selected by a policy_type
	PolicyType pulumi.StringPtrOutput `pulumi:"policyType"`
	// The project_id you want to attach the resource to
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The tags associated with the placement group
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The zone you want to attach the resource to
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewInstancePlacementGroup registers a new resource with the given unique name, arguments, and options.
func NewInstancePlacementGroup(ctx *pulumi.Context,
	name string, args *InstancePlacementGroupArgs, opts ...pulumi.ResourceOption) (*InstancePlacementGroup, error) {
	if args == nil {
		args = &InstancePlacementGroupArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource InstancePlacementGroup
	err := ctx.RegisterResource("scaleway:index/instancePlacementGroup:InstancePlacementGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstancePlacementGroup gets an existing InstancePlacementGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstancePlacementGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstancePlacementGroupState, opts ...pulumi.ResourceOption) (*InstancePlacementGroup, error) {
	var resource InstancePlacementGroup
	err := ctx.ReadResource("scaleway:index/instancePlacementGroup:InstancePlacementGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstancePlacementGroup resources.
type instancePlacementGroupState struct {
	// The name of the placement group
	Name *string `pulumi:"name"`
	// The organization_id you want to attach the resource to
	OrganizationId *string `pulumi:"organizationId"`
	// One of the two policy_mode may be selected: enforced or optional.
	PolicyMode *string `pulumi:"policyMode"`
	// Is true when the policy is respected.
	PolicyRespected *bool `pulumi:"policyRespected"`
	// The operating mode is selected by a policy_type
	PolicyType *string `pulumi:"policyType"`
	// The project_id you want to attach the resource to
	ProjectId *string `pulumi:"projectId"`
	// The tags associated with the placement group
	Tags []string `pulumi:"tags"`
	// The zone you want to attach the resource to
	Zone *string `pulumi:"zone"`
}

type InstancePlacementGroupState struct {
	// The name of the placement group
	Name pulumi.StringPtrInput
	// The organization_id you want to attach the resource to
	OrganizationId pulumi.StringPtrInput
	// One of the two policy_mode may be selected: enforced or optional.
	PolicyMode pulumi.StringPtrInput
	// Is true when the policy is respected.
	PolicyRespected pulumi.BoolPtrInput
	// The operating mode is selected by a policy_type
	PolicyType pulumi.StringPtrInput
	// The project_id you want to attach the resource to
	ProjectId pulumi.StringPtrInput
	// The tags associated with the placement group
	Tags pulumi.StringArrayInput
	// The zone you want to attach the resource to
	Zone pulumi.StringPtrInput
}

func (InstancePlacementGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*instancePlacementGroupState)(nil)).Elem()
}

type instancePlacementGroupArgs struct {
	// The name of the placement group
	Name *string `pulumi:"name"`
	// One of the two policy_mode may be selected: enforced or optional.
	PolicyMode *string `pulumi:"policyMode"`
	// The operating mode is selected by a policy_type
	PolicyType *string `pulumi:"policyType"`
	// The project_id you want to attach the resource to
	ProjectId *string `pulumi:"projectId"`
	// The tags associated with the placement group
	Tags []string `pulumi:"tags"`
	// The zone you want to attach the resource to
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a InstancePlacementGroup resource.
type InstancePlacementGroupArgs struct {
	// The name of the placement group
	Name pulumi.StringPtrInput
	// One of the two policy_mode may be selected: enforced or optional.
	PolicyMode pulumi.StringPtrInput
	// The operating mode is selected by a policy_type
	PolicyType pulumi.StringPtrInput
	// The project_id you want to attach the resource to
	ProjectId pulumi.StringPtrInput
	// The tags associated with the placement group
	Tags pulumi.StringArrayInput
	// The zone you want to attach the resource to
	Zone pulumi.StringPtrInput
}

func (InstancePlacementGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instancePlacementGroupArgs)(nil)).Elem()
}

type InstancePlacementGroupInput interface {
	pulumi.Input

	ToInstancePlacementGroupOutput() InstancePlacementGroupOutput
	ToInstancePlacementGroupOutputWithContext(ctx context.Context) InstancePlacementGroupOutput
}

func (*InstancePlacementGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**InstancePlacementGroup)(nil)).Elem()
}

func (i *InstancePlacementGroup) ToInstancePlacementGroupOutput() InstancePlacementGroupOutput {
	return i.ToInstancePlacementGroupOutputWithContext(context.Background())
}

func (i *InstancePlacementGroup) ToInstancePlacementGroupOutputWithContext(ctx context.Context) InstancePlacementGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstancePlacementGroupOutput)
}

// InstancePlacementGroupArrayInput is an input type that accepts InstancePlacementGroupArray and InstancePlacementGroupArrayOutput values.
// You can construct a concrete instance of `InstancePlacementGroupArrayInput` via:
//
//	InstancePlacementGroupArray{ InstancePlacementGroupArgs{...} }
type InstancePlacementGroupArrayInput interface {
	pulumi.Input

	ToInstancePlacementGroupArrayOutput() InstancePlacementGroupArrayOutput
	ToInstancePlacementGroupArrayOutputWithContext(context.Context) InstancePlacementGroupArrayOutput
}

type InstancePlacementGroupArray []InstancePlacementGroupInput

func (InstancePlacementGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstancePlacementGroup)(nil)).Elem()
}

func (i InstancePlacementGroupArray) ToInstancePlacementGroupArrayOutput() InstancePlacementGroupArrayOutput {
	return i.ToInstancePlacementGroupArrayOutputWithContext(context.Background())
}

func (i InstancePlacementGroupArray) ToInstancePlacementGroupArrayOutputWithContext(ctx context.Context) InstancePlacementGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstancePlacementGroupArrayOutput)
}

// InstancePlacementGroupMapInput is an input type that accepts InstancePlacementGroupMap and InstancePlacementGroupMapOutput values.
// You can construct a concrete instance of `InstancePlacementGroupMapInput` via:
//
//	InstancePlacementGroupMap{ "key": InstancePlacementGroupArgs{...} }
type InstancePlacementGroupMapInput interface {
	pulumi.Input

	ToInstancePlacementGroupMapOutput() InstancePlacementGroupMapOutput
	ToInstancePlacementGroupMapOutputWithContext(context.Context) InstancePlacementGroupMapOutput
}

type InstancePlacementGroupMap map[string]InstancePlacementGroupInput

func (InstancePlacementGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstancePlacementGroup)(nil)).Elem()
}

func (i InstancePlacementGroupMap) ToInstancePlacementGroupMapOutput() InstancePlacementGroupMapOutput {
	return i.ToInstancePlacementGroupMapOutputWithContext(context.Background())
}

func (i InstancePlacementGroupMap) ToInstancePlacementGroupMapOutputWithContext(ctx context.Context) InstancePlacementGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstancePlacementGroupMapOutput)
}

type InstancePlacementGroupOutput struct{ *pulumi.OutputState }

func (InstancePlacementGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstancePlacementGroup)(nil)).Elem()
}

func (o InstancePlacementGroupOutput) ToInstancePlacementGroupOutput() InstancePlacementGroupOutput {
	return o
}

func (o InstancePlacementGroupOutput) ToInstancePlacementGroupOutputWithContext(ctx context.Context) InstancePlacementGroupOutput {
	return o
}

// The name of the placement group
func (o InstancePlacementGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *InstancePlacementGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The organization_id you want to attach the resource to
func (o InstancePlacementGroupOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *InstancePlacementGroup) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// One of the two policy_mode may be selected: enforced or optional.
func (o InstancePlacementGroupOutput) PolicyMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstancePlacementGroup) pulumi.StringPtrOutput { return v.PolicyMode }).(pulumi.StringPtrOutput)
}

// Is true when the policy is respected.
func (o InstancePlacementGroupOutput) PolicyRespected() pulumi.BoolOutput {
	return o.ApplyT(func(v *InstancePlacementGroup) pulumi.BoolOutput { return v.PolicyRespected }).(pulumi.BoolOutput)
}

// The operating mode is selected by a policy_type
func (o InstancePlacementGroupOutput) PolicyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstancePlacementGroup) pulumi.StringPtrOutput { return v.PolicyType }).(pulumi.StringPtrOutput)
}

// The project_id you want to attach the resource to
func (o InstancePlacementGroupOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *InstancePlacementGroup) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The tags associated with the placement group
func (o InstancePlacementGroupOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *InstancePlacementGroup) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// The zone you want to attach the resource to
func (o InstancePlacementGroupOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *InstancePlacementGroup) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type InstancePlacementGroupArrayOutput struct{ *pulumi.OutputState }

func (InstancePlacementGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstancePlacementGroup)(nil)).Elem()
}

func (o InstancePlacementGroupArrayOutput) ToInstancePlacementGroupArrayOutput() InstancePlacementGroupArrayOutput {
	return o
}

func (o InstancePlacementGroupArrayOutput) ToInstancePlacementGroupArrayOutputWithContext(ctx context.Context) InstancePlacementGroupArrayOutput {
	return o
}

func (o InstancePlacementGroupArrayOutput) Index(i pulumi.IntInput) InstancePlacementGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *InstancePlacementGroup {
		return vs[0].([]*InstancePlacementGroup)[vs[1].(int)]
	}).(InstancePlacementGroupOutput)
}

type InstancePlacementGroupMapOutput struct{ *pulumi.OutputState }

func (InstancePlacementGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstancePlacementGroup)(nil)).Elem()
}

func (o InstancePlacementGroupMapOutput) ToInstancePlacementGroupMapOutput() InstancePlacementGroupMapOutput {
	return o
}

func (o InstancePlacementGroupMapOutput) ToInstancePlacementGroupMapOutputWithContext(ctx context.Context) InstancePlacementGroupMapOutput {
	return o
}

func (o InstancePlacementGroupMapOutput) MapIndex(k pulumi.StringInput) InstancePlacementGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *InstancePlacementGroup {
		return vs[0].(map[string]*InstancePlacementGroup)[vs[1].(string)]
	}).(InstancePlacementGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstancePlacementGroupInput)(nil)).Elem(), &InstancePlacementGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstancePlacementGroupArrayInput)(nil)).Elem(), InstancePlacementGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstancePlacementGroupMapInput)(nil)).Elem(), InstancePlacementGroupMap{})
	pulumi.RegisterOutputType(InstancePlacementGroupOutput{})
	pulumi.RegisterOutputType(InstancePlacementGroupArrayOutput{})
	pulumi.RegisterOutputType(InstancePlacementGroupMapOutput{})
}
