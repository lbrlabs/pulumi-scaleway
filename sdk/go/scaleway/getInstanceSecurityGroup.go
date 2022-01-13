// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about a Security Group.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-scaleway/sdk/go/scaleway"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "11111111-1111-1111-1111-111111111111"
// 		_, err := scaleway.LookupInstanceSecurityGroup(ctx, &GetInstanceSecurityGroupArgs{
// 			SecurityGroupId: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupInstanceSecurityGroup(ctx *pulumi.Context, args *LookupInstanceSecurityGroupArgs, opts ...pulumi.InvokeOption) (*LookupInstanceSecurityGroupResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupInstanceSecurityGroupResult
	err := ctx.Invoke("scaleway:index/getInstanceSecurityGroup:getInstanceSecurityGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstanceSecurityGroup.
type LookupInstanceSecurityGroupArgs struct {
	// The security group name. Only one of `name` and `securityGroupId` should be specified.
	Name *string `pulumi:"name"`
	// The security group id. Only one of `name` and `securityGroupId` should be specified.
	SecurityGroupId *string `pulumi:"securityGroupId"`
	// `zone`) The zone in which the security group exists.
	Zone *string `pulumi:"zone"`
}

// A collection of values returned by getInstanceSecurityGroup.
type LookupInstanceSecurityGroupResult struct {
	Description           string `pulumi:"description"`
	EnableDefaultSecurity bool   `pulumi:"enableDefaultSecurity"`
	ExternalRules         bool   `pulumi:"externalRules"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The default policy on incoming traffic. Possible values are: `accept` or `drop`.
	InboundDefaultPolicy string `pulumi:"inboundDefaultPolicy"`
	// A list of inbound rule to add to the security group. (Structure is documented below.)
	InboundRules []GetInstanceSecurityGroupInboundRule `pulumi:"inboundRules"`
	Name         *string                               `pulumi:"name"`
	// The ID of the organization the security group is associated with.
	OrganizationId string `pulumi:"organizationId"`
	// The default policy on outgoing traffic. Possible values are: `accept` or `drop`.
	OutboundDefaultPolicy string `pulumi:"outboundDefaultPolicy"`
	// A list of outbound rule to add to the security group. (Structure is documented below.)
	OutboundRules []GetInstanceSecurityGroupOutboundRule `pulumi:"outboundRules"`
	// The ID of the project the security group is associated with.
	ProjectId       string  `pulumi:"projectId"`
	SecurityGroupId *string `pulumi:"securityGroupId"`
	Stateful        bool    `pulumi:"stateful"`
	Zone            *string `pulumi:"zone"`
}

func LookupInstanceSecurityGroupOutput(ctx *pulumi.Context, args LookupInstanceSecurityGroupOutputArgs, opts ...pulumi.InvokeOption) LookupInstanceSecurityGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInstanceSecurityGroupResult, error) {
			args := v.(LookupInstanceSecurityGroupArgs)
			r, err := LookupInstanceSecurityGroup(ctx, &args, opts...)
			return *r, err
		}).(LookupInstanceSecurityGroupResultOutput)
}

// A collection of arguments for invoking getInstanceSecurityGroup.
type LookupInstanceSecurityGroupOutputArgs struct {
	// The security group name. Only one of `name` and `securityGroupId` should be specified.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The security group id. Only one of `name` and `securityGroupId` should be specified.
	SecurityGroupId pulumi.StringPtrInput `pulumi:"securityGroupId"`
	// `zone`) The zone in which the security group exists.
	Zone pulumi.StringPtrInput `pulumi:"zone"`
}

func (LookupInstanceSecurityGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceSecurityGroupArgs)(nil)).Elem()
}

// A collection of values returned by getInstanceSecurityGroup.
type LookupInstanceSecurityGroupResultOutput struct{ *pulumi.OutputState }

func (LookupInstanceSecurityGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceSecurityGroupResult)(nil)).Elem()
}

func (o LookupInstanceSecurityGroupResultOutput) ToLookupInstanceSecurityGroupResultOutput() LookupInstanceSecurityGroupResultOutput {
	return o
}

func (o LookupInstanceSecurityGroupResultOutput) ToLookupInstanceSecurityGroupResultOutputWithContext(ctx context.Context) LookupInstanceSecurityGroupResultOutput {
	return o
}

func (o LookupInstanceSecurityGroupResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceSecurityGroupResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupInstanceSecurityGroupResultOutput) EnableDefaultSecurity() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInstanceSecurityGroupResult) bool { return v.EnableDefaultSecurity }).(pulumi.BoolOutput)
}

func (o LookupInstanceSecurityGroupResultOutput) ExternalRules() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInstanceSecurityGroupResult) bool { return v.ExternalRules }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupInstanceSecurityGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceSecurityGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// The default policy on incoming traffic. Possible values are: `accept` or `drop`.
func (o LookupInstanceSecurityGroupResultOutput) InboundDefaultPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceSecurityGroupResult) string { return v.InboundDefaultPolicy }).(pulumi.StringOutput)
}

// A list of inbound rule to add to the security group. (Structure is documented below.)
func (o LookupInstanceSecurityGroupResultOutput) InboundRules() GetInstanceSecurityGroupInboundRuleArrayOutput {
	return o.ApplyT(func(v LookupInstanceSecurityGroupResult) []GetInstanceSecurityGroupInboundRule { return v.InboundRules }).(GetInstanceSecurityGroupInboundRuleArrayOutput)
}

func (o LookupInstanceSecurityGroupResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceSecurityGroupResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The ID of the organization the security group is associated with.
func (o LookupInstanceSecurityGroupResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceSecurityGroupResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

// The default policy on outgoing traffic. Possible values are: `accept` or `drop`.
func (o LookupInstanceSecurityGroupResultOutput) OutboundDefaultPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceSecurityGroupResult) string { return v.OutboundDefaultPolicy }).(pulumi.StringOutput)
}

// A list of outbound rule to add to the security group. (Structure is documented below.)
func (o LookupInstanceSecurityGroupResultOutput) OutboundRules() GetInstanceSecurityGroupOutboundRuleArrayOutput {
	return o.ApplyT(func(v LookupInstanceSecurityGroupResult) []GetInstanceSecurityGroupOutboundRule {
		return v.OutboundRules
	}).(GetInstanceSecurityGroupOutboundRuleArrayOutput)
}

// The ID of the project the security group is associated with.
func (o LookupInstanceSecurityGroupResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceSecurityGroupResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o LookupInstanceSecurityGroupResultOutput) SecurityGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceSecurityGroupResult) *string { return v.SecurityGroupId }).(pulumi.StringPtrOutput)
}

func (o LookupInstanceSecurityGroupResultOutput) Stateful() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInstanceSecurityGroupResult) bool { return v.Stateful }).(pulumi.BoolOutput)
}

func (o LookupInstanceSecurityGroupResultOutput) Zone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceSecurityGroupResult) *string { return v.Zone }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInstanceSecurityGroupResultOutput{})
}
