// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about an instance IP.
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
// 		opt0 := "fr-par-1/11111111-1111-1111-1111-111111111111"
// 		_, err := scaleway.LookupInstanceIp(ctx, &GetInstanceIpArgs{
// 			Id: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupInstanceIp(ctx *pulumi.Context, args *LookupInstanceIpArgs, opts ...pulumi.InvokeOption) (*LookupInstanceIpResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupInstanceIpResult
	err := ctx.Invoke("scaleway:index/getInstanceIp:getInstanceIp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstanceIp.
type LookupInstanceIpArgs struct {
	// The IPv4 address to retrieve
	// Only one of `address` and `id` should be specified.
	Address *string `pulumi:"address"`
	// The ID of the IP address to retrieve
	// Only one of `address` and `id` should be specified.
	Id *string `pulumi:"id"`
}

// A collection of values returned by getInstanceIp.
type LookupInstanceIpResult struct {
	// The IP address.
	Address *string `pulumi:"address"`
	// The ID of the IP.
	Id *string `pulumi:"id"`
	// The organization ID the IP is associated with.
	OrganizationId string `pulumi:"organizationId"`
	ProjectId      string `pulumi:"projectId"`
	// The reverse dns attached to this IP
	Reverse  string `pulumi:"reverse"`
	ServerId string `pulumi:"serverId"`
	Zone     string `pulumi:"zone"`
}

func LookupInstanceIpOutput(ctx *pulumi.Context, args LookupInstanceIpOutputArgs, opts ...pulumi.InvokeOption) LookupInstanceIpResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInstanceIpResult, error) {
			args := v.(LookupInstanceIpArgs)
			r, err := LookupInstanceIp(ctx, &args, opts...)
			return *r, err
		}).(LookupInstanceIpResultOutput)
}

// A collection of arguments for invoking getInstanceIp.
type LookupInstanceIpOutputArgs struct {
	// The IPv4 address to retrieve
	// Only one of `address` and `id` should be specified.
	Address pulumi.StringPtrInput `pulumi:"address"`
	// The ID of the IP address to retrieve
	// Only one of `address` and `id` should be specified.
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (LookupInstanceIpOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceIpArgs)(nil)).Elem()
}

// A collection of values returned by getInstanceIp.
type LookupInstanceIpResultOutput struct{ *pulumi.OutputState }

func (LookupInstanceIpResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceIpResult)(nil)).Elem()
}

func (o LookupInstanceIpResultOutput) ToLookupInstanceIpResultOutput() LookupInstanceIpResultOutput {
	return o
}

func (o LookupInstanceIpResultOutput) ToLookupInstanceIpResultOutputWithContext(ctx context.Context) LookupInstanceIpResultOutput {
	return o
}

// The IP address.
func (o LookupInstanceIpResultOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceIpResult) *string { return v.Address }).(pulumi.StringPtrOutput)
}

// The ID of the IP.
func (o LookupInstanceIpResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceIpResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The organization ID the IP is associated with.
func (o LookupInstanceIpResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceIpResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o LookupInstanceIpResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceIpResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// The reverse dns attached to this IP
func (o LookupInstanceIpResultOutput) Reverse() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceIpResult) string { return v.Reverse }).(pulumi.StringOutput)
}

func (o LookupInstanceIpResultOutput) ServerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceIpResult) string { return v.ServerId }).(pulumi.StringOutput)
}

func (o LookupInstanceIpResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceIpResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInstanceIpResultOutput{})
}
