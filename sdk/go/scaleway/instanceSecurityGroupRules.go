// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// Instance security group rules can be imported using the `{zone}/{id}`, e.g. bash
//
// ```sh
//  $ pulumi import scaleway:index/instanceSecurityGroupRules:InstanceSecurityGroupRules web fr-par-1/11111111-1111-1111-1111-111111111111
// ```
type InstanceSecurityGroupRules struct {
	pulumi.CustomResourceState

	// A list of inbound rule to add to the security group. (Structure is documented below.)
	InboundRules InstanceSecurityGroupRulesInboundRuleArrayOutput `pulumi:"inboundRules"`
	// A list of outbound rule to add to the security group. (Structure is documented below.)
	OutboundRules InstanceSecurityGroupRulesOutboundRuleArrayOutput `pulumi:"outboundRules"`
	// The ID of the security group.
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
	// A list of inbound rule to add to the security group. (Structure is documented below.)
	InboundRules []InstanceSecurityGroupRulesInboundRule `pulumi:"inboundRules"`
	// A list of outbound rule to add to the security group. (Structure is documented below.)
	OutboundRules []InstanceSecurityGroupRulesOutboundRule `pulumi:"outboundRules"`
	// The ID of the security group.
	SecurityGroupId *string `pulumi:"securityGroupId"`
}

type InstanceSecurityGroupRulesState struct {
	// A list of inbound rule to add to the security group. (Structure is documented below.)
	InboundRules InstanceSecurityGroupRulesInboundRuleArrayInput
	// A list of outbound rule to add to the security group. (Structure is documented below.)
	OutboundRules InstanceSecurityGroupRulesOutboundRuleArrayInput
	// The ID of the security group.
	SecurityGroupId pulumi.StringPtrInput
}

func (InstanceSecurityGroupRulesState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceSecurityGroupRulesState)(nil)).Elem()
}

type instanceSecurityGroupRulesArgs struct {
	// A list of inbound rule to add to the security group. (Structure is documented below.)
	InboundRules []InstanceSecurityGroupRulesInboundRule `pulumi:"inboundRules"`
	// A list of outbound rule to add to the security group. (Structure is documented below.)
	OutboundRules []InstanceSecurityGroupRulesOutboundRule `pulumi:"outboundRules"`
	// The ID of the security group.
	SecurityGroupId string `pulumi:"securityGroupId"`
}

// The set of arguments for constructing a InstanceSecurityGroupRules resource.
type InstanceSecurityGroupRulesArgs struct {
	// A list of inbound rule to add to the security group. (Structure is documented below.)
	InboundRules InstanceSecurityGroupRulesInboundRuleArrayInput
	// A list of outbound rule to add to the security group. (Structure is documented below.)
	OutboundRules InstanceSecurityGroupRulesOutboundRuleArrayInput
	// The ID of the security group.
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
	return reflect.TypeOf((*InstanceSecurityGroupRules)(nil))
}

func (i *InstanceSecurityGroupRules) ToInstanceSecurityGroupRulesOutput() InstanceSecurityGroupRulesOutput {
	return i.ToInstanceSecurityGroupRulesOutputWithContext(context.Background())
}

func (i *InstanceSecurityGroupRules) ToInstanceSecurityGroupRulesOutputWithContext(ctx context.Context) InstanceSecurityGroupRulesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceSecurityGroupRulesOutput)
}

type InstanceSecurityGroupRulesOutput struct{ *pulumi.OutputState }

func (InstanceSecurityGroupRulesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceSecurityGroupRules)(nil))
}

func (o InstanceSecurityGroupRulesOutput) ToInstanceSecurityGroupRulesOutput() InstanceSecurityGroupRulesOutput {
	return o
}

func (o InstanceSecurityGroupRulesOutput) ToInstanceSecurityGroupRulesOutputWithContext(ctx context.Context) InstanceSecurityGroupRulesOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceSecurityGroupRulesInput)(nil)).Elem(), &InstanceSecurityGroupRules{})
	pulumi.RegisterOutputType(InstanceSecurityGroupRulesOutput{})
}