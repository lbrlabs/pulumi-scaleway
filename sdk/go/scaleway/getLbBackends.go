// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/lbrlabs/pulumi-scaleway/sdk/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets information about multiple Load Balancer Backends.
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
//			_, err := scaleway.GetLbBackends(ctx, &scaleway.GetLbBackendsArgs{
//				LbId: scaleway_lb.Lb01.Id,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.GetLbBackends(ctx, &scaleway.GetLbBackendsArgs{
//				LbId: scaleway_lb.Lb01.Id,
//				Name: pulumi.StringRef("tf-backend-datasource"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetLbBackends(ctx *pulumi.Context, args *GetLbBackendsArgs, opts ...pulumi.InvokeOption) (*GetLbBackendsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetLbBackendsResult
	err := ctx.Invoke("scaleway:index/getLbBackends:getLbBackends", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLbBackends.
type GetLbBackendsArgs struct {
	// The load-balancer ID this backend is attached to. backends with a LB ID like it are listed.
	LbId string `pulumi:"lbId"`
	// The backend name used as filter. Backends with a name like it are listed.
	Name      *string `pulumi:"name"`
	ProjectId *string `pulumi:"projectId"`
	// `zone`) The zone in which backends exist.
	Zone *string `pulumi:"zone"`
}

// A collection of values returned by getLbBackends.
type GetLbBackendsResult struct {
	// List of found backends
	Backends []GetLbBackendsBackend `pulumi:"backends"`
	// The provider-assigned unique ID for this managed resource.
	Id             string  `pulumi:"id"`
	LbId           string  `pulumi:"lbId"`
	Name           *string `pulumi:"name"`
	OrganizationId string  `pulumi:"organizationId"`
	ProjectId      string  `pulumi:"projectId"`
	Zone           string  `pulumi:"zone"`
}

func GetLbBackendsOutput(ctx *pulumi.Context, args GetLbBackendsOutputArgs, opts ...pulumi.InvokeOption) GetLbBackendsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetLbBackendsResult, error) {
			args := v.(GetLbBackendsArgs)
			r, err := GetLbBackends(ctx, &args, opts...)
			var s GetLbBackendsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetLbBackendsResultOutput)
}

// A collection of arguments for invoking getLbBackends.
type GetLbBackendsOutputArgs struct {
	// The load-balancer ID this backend is attached to. backends with a LB ID like it are listed.
	LbId pulumi.StringInput `pulumi:"lbId"`
	// The backend name used as filter. Backends with a name like it are listed.
	Name      pulumi.StringPtrInput `pulumi:"name"`
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	// `zone`) The zone in which backends exist.
	Zone pulumi.StringPtrInput `pulumi:"zone"`
}

func (GetLbBackendsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLbBackendsArgs)(nil)).Elem()
}

// A collection of values returned by getLbBackends.
type GetLbBackendsResultOutput struct{ *pulumi.OutputState }

func (GetLbBackendsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLbBackendsResult)(nil)).Elem()
}

func (o GetLbBackendsResultOutput) ToGetLbBackendsResultOutput() GetLbBackendsResultOutput {
	return o
}

func (o GetLbBackendsResultOutput) ToGetLbBackendsResultOutputWithContext(ctx context.Context) GetLbBackendsResultOutput {
	return o
}

func (o GetLbBackendsResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetLbBackendsResult] {
	return pulumix.Output[GetLbBackendsResult]{
		OutputState: o.OutputState,
	}
}

// List of found backends
func (o GetLbBackendsResultOutput) Backends() GetLbBackendsBackendArrayOutput {
	return o.ApplyT(func(v GetLbBackendsResult) []GetLbBackendsBackend { return v.Backends }).(GetLbBackendsBackendArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetLbBackendsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetLbBackendsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetLbBackendsResultOutput) LbId() pulumi.StringOutput {
	return o.ApplyT(func(v GetLbBackendsResult) string { return v.LbId }).(pulumi.StringOutput)
}

func (o GetLbBackendsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLbBackendsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetLbBackendsResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v GetLbBackendsResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o GetLbBackendsResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetLbBackendsResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o GetLbBackendsResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v GetLbBackendsResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLbBackendsResultOutput{})
}
