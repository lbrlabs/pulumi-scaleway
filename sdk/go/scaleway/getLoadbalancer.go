// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about a Load Balancer.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/lbrlabs/pulumi-scaleway/sdk/go/scaleway"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err = scaleway.LookupLoadbalancer(ctx, &GetLoadbalancerArgs{
//				Name: pulumi.StringRef("foobar"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.LookupLoadbalancer(ctx, &GetLoadbalancerArgs{
//				LbId: pulumi.StringRef("11111111-1111-1111-1111-111111111111"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupLoadbalancer(ctx *pulumi.Context, args *LookupLoadbalancerArgs, opts ...pulumi.InvokeOption) (*LookupLoadbalancerResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupLoadbalancerResult
	err := ctx.Invoke("scaleway:index/getLoadbalancer:getLoadbalancer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLoadbalancer.
type LookupLoadbalancerArgs struct {
	LbId *string `pulumi:"lbId"`
	// The IP address.
	Name      *string `pulumi:"name"`
	ReleaseIp *bool   `pulumi:"releaseIp"`
	// (Defaults to provider `region`) The region in which the LB exists.
	Zone *string `pulumi:"zone"`
}

// A collection of values returned by getLoadbalancer.
type LookupLoadbalancerResult struct {
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The load-balancer public IP Address.
	IpAddress       string                          `pulumi:"ipAddress"`
	IpId            string                          `pulumi:"ipId"`
	LbId            *string                         `pulumi:"lbId"`
	Name            *string                         `pulumi:"name"`
	OrganizationId  string                          `pulumi:"organizationId"`
	PrivateNetworks []GetLoadbalancerPrivateNetwork `pulumi:"privateNetworks"`
	// (Defaults to provider `projectId`) The ID of the project the LB is associated with.
	ProjectId             string `pulumi:"projectId"`
	Region                string `pulumi:"region"`
	ReleaseIp             *bool  `pulumi:"releaseIp"`
	SslCompatibilityLevel string `pulumi:"sslCompatibilityLevel"`
	// The tags associated with the load-balancers.
	Tags []string `pulumi:"tags"`
	// The type of the load-balancer.
	Type string `pulumi:"type"`
	// (Defaults to provider `region`) The region in which the LB exists.
	Zone *string `pulumi:"zone"`
}

func LookupLoadbalancerOutput(ctx *pulumi.Context, args LookupLoadbalancerOutputArgs, opts ...pulumi.InvokeOption) LookupLoadbalancerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLoadbalancerResult, error) {
			args := v.(LookupLoadbalancerArgs)
			r, err := LookupLoadbalancer(ctx, &args, opts...)
			var s LookupLoadbalancerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLoadbalancerResultOutput)
}

// A collection of arguments for invoking getLoadbalancer.
type LookupLoadbalancerOutputArgs struct {
	LbId pulumi.StringPtrInput `pulumi:"lbId"`
	// The IP address.
	Name      pulumi.StringPtrInput `pulumi:"name"`
	ReleaseIp pulumi.BoolPtrInput   `pulumi:"releaseIp"`
	// (Defaults to provider `region`) The region in which the LB exists.
	Zone pulumi.StringPtrInput `pulumi:"zone"`
}

func (LookupLoadbalancerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLoadbalancerArgs)(nil)).Elem()
}

// A collection of values returned by getLoadbalancer.
type LookupLoadbalancerResultOutput struct{ *pulumi.OutputState }

func (LookupLoadbalancerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLoadbalancerResult)(nil)).Elem()
}

func (o LookupLoadbalancerResultOutput) ToLookupLoadbalancerResultOutput() LookupLoadbalancerResultOutput {
	return o
}

func (o LookupLoadbalancerResultOutput) ToLookupLoadbalancerResultOutputWithContext(ctx context.Context) LookupLoadbalancerResultOutput {
	return o
}

func (o LookupLoadbalancerResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoadbalancerResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupLoadbalancerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoadbalancerResult) string { return v.Id }).(pulumi.StringOutput)
}

// The load-balancer public IP Address.
func (o LookupLoadbalancerResultOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoadbalancerResult) string { return v.IpAddress }).(pulumi.StringOutput)
}

func (o LookupLoadbalancerResultOutput) IpId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoadbalancerResult) string { return v.IpId }).(pulumi.StringOutput)
}

func (o LookupLoadbalancerResultOutput) LbId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLoadbalancerResult) *string { return v.LbId }).(pulumi.StringPtrOutput)
}

func (o LookupLoadbalancerResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLoadbalancerResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupLoadbalancerResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoadbalancerResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o LookupLoadbalancerResultOutput) PrivateNetworks() GetLoadbalancerPrivateNetworkArrayOutput {
	return o.ApplyT(func(v LookupLoadbalancerResult) []GetLoadbalancerPrivateNetwork { return v.PrivateNetworks }).(GetLoadbalancerPrivateNetworkArrayOutput)
}

// (Defaults to provider `projectId`) The ID of the project the LB is associated with.
func (o LookupLoadbalancerResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoadbalancerResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o LookupLoadbalancerResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoadbalancerResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o LookupLoadbalancerResultOutput) ReleaseIp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLoadbalancerResult) *bool { return v.ReleaseIp }).(pulumi.BoolPtrOutput)
}

func (o LookupLoadbalancerResultOutput) SslCompatibilityLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoadbalancerResult) string { return v.SslCompatibilityLevel }).(pulumi.StringOutput)
}

// The tags associated with the load-balancers.
func (o LookupLoadbalancerResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLoadbalancerResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

// The type of the load-balancer.
func (o LookupLoadbalancerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoadbalancerResult) string { return v.Type }).(pulumi.StringOutput)
}

// (Defaults to provider `region`) The region in which the LB exists.
func (o LookupLoadbalancerResultOutput) Zone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLoadbalancerResult) *string { return v.Zone }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLoadbalancerResultOutput{})
}
