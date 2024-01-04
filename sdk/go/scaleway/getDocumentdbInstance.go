// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/lbrlabs/pulumi-scaleway/sdk/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about an DocumentDB instance. For further information see our [developers website](https://www.scaleway.com/en/developers/api/document_db/)
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
//			_, err := scaleway.LookupDocumentdbInstance(ctx, &scaleway.LookupDocumentdbInstanceArgs{
//				InstanceId: pulumi.StringRef("11111111-1111-1111-1111-111111111111"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupDocumentdbInstance(ctx *pulumi.Context, args *LookupDocumentdbInstanceArgs, opts ...pulumi.InvokeOption) (*LookupDocumentdbInstanceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDocumentdbInstanceResult
	err := ctx.Invoke("scaleway:index/getDocumentdbInstance:getDocumentdbInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDocumentdbInstance.
type LookupDocumentdbInstanceArgs struct {
	// The DocumentDB instance ID.
	// Only one of `name` and `instanceId` should be specified.
	InstanceId *string `pulumi:"instanceId"`
	// The name of the DocumentDB instance.
	// Only one of `name` and `instanceId` should be specified.
	Name *string `pulumi:"name"`
	// `region`) The region in which the DocumentDB instance exists.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getDocumentdbInstance.
type LookupDocumentdbInstanceResult struct {
	Engine string `pulumi:"engine"`
	// The provider-assigned unique ID for this managed resource.
	Id               string   `pulumi:"id"`
	InstanceId       *string  `pulumi:"instanceId"`
	IsHaCluster      bool     `pulumi:"isHaCluster"`
	Name             *string  `pulumi:"name"`
	NodeType         string   `pulumi:"nodeType"`
	Password         string   `pulumi:"password"`
	ProjectId        string   `pulumi:"projectId"`
	Region           *string  `pulumi:"region"`
	Tags             []string `pulumi:"tags"`
	TelemetryEnabled bool     `pulumi:"telemetryEnabled"`
	UserName         string   `pulumi:"userName"`
	VolumeSizeInGb   int      `pulumi:"volumeSizeInGb"`
	VolumeType       string   `pulumi:"volumeType"`
}

func LookupDocumentdbInstanceOutput(ctx *pulumi.Context, args LookupDocumentdbInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupDocumentdbInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDocumentdbInstanceResult, error) {
			args := v.(LookupDocumentdbInstanceArgs)
			r, err := LookupDocumentdbInstance(ctx, &args, opts...)
			var s LookupDocumentdbInstanceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDocumentdbInstanceResultOutput)
}

// A collection of arguments for invoking getDocumentdbInstance.
type LookupDocumentdbInstanceOutputArgs struct {
	// The DocumentDB instance ID.
	// Only one of `name` and `instanceId` should be specified.
	InstanceId pulumi.StringPtrInput `pulumi:"instanceId"`
	// The name of the DocumentDB instance.
	// Only one of `name` and `instanceId` should be specified.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// `region`) The region in which the DocumentDB instance exists.
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (LookupDocumentdbInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDocumentdbInstanceArgs)(nil)).Elem()
}

// A collection of values returned by getDocumentdbInstance.
type LookupDocumentdbInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupDocumentdbInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDocumentdbInstanceResult)(nil)).Elem()
}

func (o LookupDocumentdbInstanceResultOutput) ToLookupDocumentdbInstanceResultOutput() LookupDocumentdbInstanceResultOutput {
	return o
}

func (o LookupDocumentdbInstanceResultOutput) ToLookupDocumentdbInstanceResultOutputWithContext(ctx context.Context) LookupDocumentdbInstanceResultOutput {
	return o
}

func (o LookupDocumentdbInstanceResultOutput) Engine() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDocumentdbInstanceResult) string { return v.Engine }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDocumentdbInstanceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDocumentdbInstanceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDocumentdbInstanceResultOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDocumentdbInstanceResult) *string { return v.InstanceId }).(pulumi.StringPtrOutput)
}

func (o LookupDocumentdbInstanceResultOutput) IsHaCluster() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDocumentdbInstanceResult) bool { return v.IsHaCluster }).(pulumi.BoolOutput)
}

func (o LookupDocumentdbInstanceResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDocumentdbInstanceResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupDocumentdbInstanceResultOutput) NodeType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDocumentdbInstanceResult) string { return v.NodeType }).(pulumi.StringOutput)
}

func (o LookupDocumentdbInstanceResultOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDocumentdbInstanceResult) string { return v.Password }).(pulumi.StringOutput)
}

func (o LookupDocumentdbInstanceResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDocumentdbInstanceResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o LookupDocumentdbInstanceResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDocumentdbInstanceResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o LookupDocumentdbInstanceResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDocumentdbInstanceResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o LookupDocumentdbInstanceResultOutput) TelemetryEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDocumentdbInstanceResult) bool { return v.TelemetryEnabled }).(pulumi.BoolOutput)
}

func (o LookupDocumentdbInstanceResultOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDocumentdbInstanceResult) string { return v.UserName }).(pulumi.StringOutput)
}

func (o LookupDocumentdbInstanceResultOutput) VolumeSizeInGb() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDocumentdbInstanceResult) int { return v.VolumeSizeInGb }).(pulumi.IntOutput)
}

func (o LookupDocumentdbInstanceResultOutput) VolumeType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDocumentdbInstanceResult) string { return v.VolumeType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDocumentdbInstanceResultOutput{})
}
