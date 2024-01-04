// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/lbrlabs/pulumi-scaleway/sdk/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about multiple Load Balancers.
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
//			_, err := scaleway.GetLbs(ctx, &scaleway.GetLbsArgs{
//				Name: pulumi.StringRef("foobar"),
//				Zone: pulumi.StringRef("fr-par-2"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetLbs(ctx *pulumi.Context, args *GetLbsArgs, opts ...pulumi.InvokeOption) (*GetLbsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetLbsResult
	err := ctx.Invoke("scaleway:index/getLbs:getLbs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLbs.
type GetLbsArgs struct {
	// The load balancer name used as a filter. LBs with a name like it are listed.
	Name *string `pulumi:"name"`
	// The ID of the project the load-balancer is associated with.
	ProjectId *string `pulumi:"projectId"`
	// `zone`) The zone in which LBs exist.
	Zone *string `pulumi:"zone"`
}

// A collection of values returned by getLbs.
type GetLbsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// List of found LBs
	Lbs []GetLbsLb `pulumi:"lbs"`
	// The name of the load-balancer.
	Name *string `pulumi:"name"`
	// The organization ID the load-balancer is associated with.
	OrganizationId string `pulumi:"organizationId"`
	// The ID of the project the load-balancer is associated with.
	ProjectId string `pulumi:"projectId"`
	// The zone in which the load-balancer is.
	Zone string `pulumi:"zone"`
}

func GetLbsOutput(ctx *pulumi.Context, args GetLbsOutputArgs, opts ...pulumi.InvokeOption) GetLbsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetLbsResult, error) {
			args := v.(GetLbsArgs)
			r, err := GetLbs(ctx, &args, opts...)
			var s GetLbsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetLbsResultOutput)
}

// A collection of arguments for invoking getLbs.
type GetLbsOutputArgs struct {
	// The load balancer name used as a filter. LBs with a name like it are listed.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The ID of the project the load-balancer is associated with.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	// `zone`) The zone in which LBs exist.
	Zone pulumi.StringPtrInput `pulumi:"zone"`
}

func (GetLbsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLbsArgs)(nil)).Elem()
}

// A collection of values returned by getLbs.
type GetLbsResultOutput struct{ *pulumi.OutputState }

func (GetLbsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLbsResult)(nil)).Elem()
}

func (o GetLbsResultOutput) ToGetLbsResultOutput() GetLbsResultOutput {
	return o
}

func (o GetLbsResultOutput) ToGetLbsResultOutputWithContext(ctx context.Context) GetLbsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetLbsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetLbsResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of found LBs
func (o GetLbsResultOutput) Lbs() GetLbsLbArrayOutput {
	return o.ApplyT(func(v GetLbsResult) []GetLbsLb { return v.Lbs }).(GetLbsLbArrayOutput)
}

// The name of the load-balancer.
func (o GetLbsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLbsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The organization ID the load-balancer is associated with.
func (o GetLbsResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v GetLbsResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

// The ID of the project the load-balancer is associated with.
func (o GetLbsResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetLbsResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// The zone in which the load-balancer is.
func (o GetLbsResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v GetLbsResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLbsResultOutput{})
}
