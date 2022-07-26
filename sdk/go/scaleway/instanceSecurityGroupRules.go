// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type InstanceSecurityGroupRules struct {
	pulumi.CustomResourceState

	// Inbound rules for this set of security group rules
	InboundRules InstanceSecurityGroupRulesInboundRuleArrayOutput `pulumi:"inboundRules"`
	// Outbound rules for this set of security group rules
	OutboundRules InstanceSecurityGroupRulesOutboundRuleArrayOutput `pulumi:"outboundRules"`
	// The security group associated with this volume
	SecurityGroupId pulumi.StringOutput `pulumi:"securityGroupId"`
}

// NewInstanceSecurityGroupRules registers a new resource with the given unique name, arguments, and options.
func NewInstanceSecurityGroupRules(ctx *pulumi.Context,
	name string, args *InstanceSecurityGroupRulesArgs, opts ...pulumi.ResourceOption) (*InstanceSecurityGroupRules, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SecurityGroupId == nil {
		return nil, errors.New("invalid value for required argument 'SecurityGroupId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource InstanceSecurityGroupRules
	err := ctx.RegisterResource("scaleway:index/instanceSecurityGroupRules:InstanceSecurityGroupRules", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceSecurityGroupRules gets an existing InstanceSecurityGroupRules resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceSecurityGroupRules(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceSecurityGroupRulesState, opts ...pulumi.ResourceOption) (*InstanceSecurityGroupRules, error) {
	var resource InstanceSecurityGroupRules
	err := ctx.ReadResource("scaleway:index/instanceSecurityGroupRules:InstanceSecurityGroupRules", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceSecurityGroupRules resources.
type instanceSecurityGroupRulesState struct {
	// Inbound rules for this set of security group rules
	InboundRules []InstanceSecurityGroupRulesInboundRule `pulumi:"inboundRules"`
	// Outbound rules for this set of security group rules
	OutboundRules []InstanceSecurityGroupRulesOutboundRule `pulumi:"outboundRules"`
	// The security group associated with this volume
	SecurityGroupId *string `pulumi:"securityGroupId"`
}

type InstanceSecurityGroupRulesState struct {
	// Inbound rules for this set of security group rules
	InboundRules InstanceSecurityGroupRulesInboundRuleArrayInput
	// Outbound rules for this set of security group rules
	OutboundRules InstanceSecurityGroupRulesOutboundRuleArrayInput
	// The security group associated with this volume
	SecurityGroupId pulumi.StringPtrInput
}

func (InstanceSecurityGroupRulesState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceSecurityGroupRulesState)(nil)).Elem()
}

type instanceSecurityGroupRulesArgs struct {
	// Inbound rules for this set of security group rules
	InboundRules []InstanceSecurityGroupRulesInboundRule `pulumi:"inboundRules"`
	// Outbound rules for this set of security group rules
	OutboundRules []InstanceSecurityGroupRulesOutboundRule `pulumi:"outboundRules"`
	// The security group associated with this volume
	SecurityGroupId string `pulumi:"securityGroupId"`
}

// The set of arguments for constructing a InstanceSecurityGroupRules resource.
type InstanceSecurityGroupRulesArgs struct {
	// Inbound rules for this set of security group rules
	InboundRules InstanceSecurityGroupRulesInboundRuleArrayInput
	// Outbound rules for this set of security group rules
	OutboundRules InstanceSecurityGroupRulesOutboundRuleArrayInput
	// The security group associated with this volume
	SecurityGroupId pulumi.StringInput
}

func (InstanceSecurityGroupRulesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceSecurityGroupRulesArgs)(nil)).Elem()
}

type InstanceSecurityGroupRulesInput interface {
	pulumi.Input

	ToInstanceSecurityGroupRulesOutput() InstanceSecurityGroupRulesOutput
	ToInstanceSecurityGroupRulesOutputWithContext(ctx context.Context) InstanceSecurityGroupRulesOutput
}

func (*InstanceSecurityGroupRules) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceSecurityGroupRules)(nil)).Elem()
}

func (i *InstanceSecurityGroupRules) ToInstanceSecurityGroupRulesOutput() InstanceSecurityGroupRulesOutput {
	return i.ToInstanceSecurityGroupRulesOutputWithContext(context.Background())
}

func (i *InstanceSecurityGroupRules) ToInstanceSecurityGroupRulesOutputWithContext(ctx context.Context) InstanceSecurityGroupRulesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceSecurityGroupRulesOutput)
}

// InstanceSecurityGroupRulesArrayInput is an input type that accepts InstanceSecurityGroupRulesArray and InstanceSecurityGroupRulesArrayOutput values.
// You can construct a concrete instance of `InstanceSecurityGroupRulesArrayInput` via:
//
//          InstanceSecurityGroupRulesArray{ InstanceSecurityGroupRulesArgs{...} }
type InstanceSecurityGroupRulesArrayInput interface {
	pulumi.Input

	ToInstanceSecurityGroupRulesArrayOutput() InstanceSecurityGroupRulesArrayOutput
	ToInstanceSecurityGroupRulesArrayOutputWithContext(context.Context) InstanceSecurityGroupRulesArrayOutput
}

type InstanceSecurityGroupRulesArray []InstanceSecurityGroupRulesInput

func (InstanceSecurityGroupRulesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceSecurityGroupRules)(nil)).Elem()
}

func (i InstanceSecurityGroupRulesArray) ToInstanceSecurityGroupRulesArrayOutput() InstanceSecurityGroupRulesArrayOutput {
	return i.ToInstanceSecurityGroupRulesArrayOutputWithContext(context.Background())
}

func (i InstanceSecurityGroupRulesArray) ToInstanceSecurityGroupRulesArrayOutputWithContext(ctx context.Context) InstanceSecurityGroupRulesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceSecurityGroupRulesArrayOutput)
}

// InstanceSecurityGroupRulesMapInput is an input type that accepts InstanceSecurityGroupRulesMap and InstanceSecurityGroupRulesMapOutput values.
// You can construct a concrete instance of `InstanceSecurityGroupRulesMapInput` via:
//
//          InstanceSecurityGroupRulesMap{ "key": InstanceSecurityGroupRulesArgs{...} }
type InstanceSecurityGroupRulesMapInput interface {
	pulumi.Input

	ToInstanceSecurityGroupRulesMapOutput() InstanceSecurityGroupRulesMapOutput
	ToInstanceSecurityGroupRulesMapOutputWithContext(context.Context) InstanceSecurityGroupRulesMapOutput
}

type InstanceSecurityGroupRulesMap map[string]InstanceSecurityGroupRulesInput

func (InstanceSecurityGroupRulesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceSecurityGroupRules)(nil)).Elem()
}

func (i InstanceSecurityGroupRulesMap) ToInstanceSecurityGroupRulesMapOutput() InstanceSecurityGroupRulesMapOutput {
	return i.ToInstanceSecurityGroupRulesMapOutputWithContext(context.Background())
}

func (i InstanceSecurityGroupRulesMap) ToInstanceSecurityGroupRulesMapOutputWithContext(ctx context.Context) InstanceSecurityGroupRulesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceSecurityGroupRulesMapOutput)
}

type InstanceSecurityGroupRulesOutput struct{ *pulumi.OutputState }

func (InstanceSecurityGroupRulesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceSecurityGroupRules)(nil)).Elem()
}

func (o InstanceSecurityGroupRulesOutput) ToInstanceSecurityGroupRulesOutput() InstanceSecurityGroupRulesOutput {
	return o
}

func (o InstanceSecurityGroupRulesOutput) ToInstanceSecurityGroupRulesOutputWithContext(ctx context.Context) InstanceSecurityGroupRulesOutput {
	return o
}

// Inbound rules for this set of security group rules
func (o InstanceSecurityGroupRulesOutput) InboundRules() InstanceSecurityGroupRulesInboundRuleArrayOutput {
	return o.ApplyT(func(v *InstanceSecurityGroupRules) InstanceSecurityGroupRulesInboundRuleArrayOutput {
		return v.InboundRules
	}).(InstanceSecurityGroupRulesInboundRuleArrayOutput)
}

// Outbound rules for this set of security group rules
func (o InstanceSecurityGroupRulesOutput) OutboundRules() InstanceSecurityGroupRulesOutboundRuleArrayOutput {
	return o.ApplyT(func(v *InstanceSecurityGroupRules) InstanceSecurityGroupRulesOutboundRuleArrayOutput {
		return v.OutboundRules
	}).(InstanceSecurityGroupRulesOutboundRuleArrayOutput)
}

// The security group associated with this volume
func (o InstanceSecurityGroupRulesOutput) SecurityGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceSecurityGroupRules) pulumi.StringOutput { return v.SecurityGroupId }).(pulumi.StringOutput)
}

type InstanceSecurityGroupRulesArrayOutput struct{ *pulumi.OutputState }

func (InstanceSecurityGroupRulesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceSecurityGroupRules)(nil)).Elem()
}

func (o InstanceSecurityGroupRulesArrayOutput) ToInstanceSecurityGroupRulesArrayOutput() InstanceSecurityGroupRulesArrayOutput {
	return o
}

func (o InstanceSecurityGroupRulesArrayOutput) ToInstanceSecurityGroupRulesArrayOutputWithContext(ctx context.Context) InstanceSecurityGroupRulesArrayOutput {
	return o
}

func (o InstanceSecurityGroupRulesArrayOutput) Index(i pulumi.IntInput) InstanceSecurityGroupRulesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *InstanceSecurityGroupRules {
		return vs[0].([]*InstanceSecurityGroupRules)[vs[1].(int)]
	}).(InstanceSecurityGroupRulesOutput)
}

type InstanceSecurityGroupRulesMapOutput struct{ *pulumi.OutputState }

func (InstanceSecurityGroupRulesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceSecurityGroupRules)(nil)).Elem()
}

func (o InstanceSecurityGroupRulesMapOutput) ToInstanceSecurityGroupRulesMapOutput() InstanceSecurityGroupRulesMapOutput {
	return o
}

func (o InstanceSecurityGroupRulesMapOutput) ToInstanceSecurityGroupRulesMapOutputWithContext(ctx context.Context) InstanceSecurityGroupRulesMapOutput {
	return o
}

func (o InstanceSecurityGroupRulesMapOutput) MapIndex(k pulumi.StringInput) InstanceSecurityGroupRulesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *InstanceSecurityGroupRules {
		return vs[0].(map[string]*InstanceSecurityGroupRules)[vs[1].(string)]
	}).(InstanceSecurityGroupRulesOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceSecurityGroupRulesInput)(nil)).Elem(), &InstanceSecurityGroupRules{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceSecurityGroupRulesArrayInput)(nil)).Elem(), InstanceSecurityGroupRulesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceSecurityGroupRulesMapInput)(nil)).Elem(), InstanceSecurityGroupRulesMap{})
	pulumi.RegisterOutputType(InstanceSecurityGroupRulesOutput{})
	pulumi.RegisterOutputType(InstanceSecurityGroupRulesArrayOutput{})
	pulumi.RegisterOutputType(InstanceSecurityGroupRulesMapOutput{})
}
