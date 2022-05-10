// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

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
	Name            *string `pulumi:"name"`
	SecurityGroupId *string `pulumi:"securityGroupId"`
	Zone            *string `pulumi:"zone"`
}

// A collection of values returned by getInstanceSecurityGroup.
type LookupInstanceSecurityGroupResult struct {
	Description           string `pulumi:"description"`
	EnableDefaultSecurity bool   `pulumi:"enableDefaultSecurity"`
	ExternalRules         bool   `pulumi:"externalRules"`
	// The provider-assigned unique ID for this managed resource.
	Id                    string                                 `pulumi:"id"`
	InboundDefaultPolicy  string                                 `pulumi:"inboundDefaultPolicy"`
	InboundRules          []GetInstanceSecurityGroupInboundRule  `pulumi:"inboundRules"`
	Name                  *string                                `pulumi:"name"`
	OrganizationId        string                                 `pulumi:"organizationId"`
	OutboundDefaultPolicy string                                 `pulumi:"outboundDefaultPolicy"`
	OutboundRules         []GetInstanceSecurityGroupOutboundRule `pulumi:"outboundRules"`
	ProjectId             string                                 `pulumi:"projectId"`
	SecurityGroupId       *string                                `pulumi:"securityGroupId"`
	Stateful              bool                                   `pulumi:"stateful"`
	Tags                  []string                               `pulumi:"tags"`
	Zone                  *string                                `pulumi:"zone"`
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
	Name            pulumi.StringPtrInput `pulumi:"name"`
	SecurityGroupId pulumi.StringPtrInput `pulumi:"securityGroupId"`
	Zone            pulumi.StringPtrInput `pulumi:"zone"`
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

func (o LookupInstanceSecurityGroupResultOutput) InboundDefaultPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceSecurityGroupResult) string { return v.InboundDefaultPolicy }).(pulumi.StringOutput)
}

func (o LookupInstanceSecurityGroupResultOutput) InboundRules() GetInstanceSecurityGroupInboundRuleArrayOutput {
	return o.ApplyT(func(v LookupInstanceSecurityGroupResult) []GetInstanceSecurityGroupInboundRule { return v.InboundRules }).(GetInstanceSecurityGroupInboundRuleArrayOutput)
}

func (o LookupInstanceSecurityGroupResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceSecurityGroupResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupInstanceSecurityGroupResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceSecurityGroupResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o LookupInstanceSecurityGroupResultOutput) OutboundDefaultPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceSecurityGroupResult) string { return v.OutboundDefaultPolicy }).(pulumi.StringOutput)
}

func (o LookupInstanceSecurityGroupResultOutput) OutboundRules() GetInstanceSecurityGroupOutboundRuleArrayOutput {
	return o.ApplyT(func(v LookupInstanceSecurityGroupResult) []GetInstanceSecurityGroupOutboundRule {
		return v.OutboundRules
	}).(GetInstanceSecurityGroupOutboundRuleArrayOutput)
}

func (o LookupInstanceSecurityGroupResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceSecurityGroupResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o LookupInstanceSecurityGroupResultOutput) SecurityGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceSecurityGroupResult) *string { return v.SecurityGroupId }).(pulumi.StringPtrOutput)
}

func (o LookupInstanceSecurityGroupResultOutput) Stateful() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInstanceSecurityGroupResult) bool { return v.Stateful }).(pulumi.BoolOutput)
}

func (o LookupInstanceSecurityGroupResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupInstanceSecurityGroupResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o LookupInstanceSecurityGroupResultOutput) Zone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceSecurityGroupResult) *string { return v.Zone }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInstanceSecurityGroupResultOutput{})
}
