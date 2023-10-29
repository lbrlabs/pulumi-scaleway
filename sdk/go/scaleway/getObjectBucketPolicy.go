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

// Gets information about the Bucket's policy.
// For more information, see [the documentation](https://www.scaleway.com/en/docs/object-storage-feature/).
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
//			_, err := scaleway.LookupObjectBucketPolicy(ctx, &scaleway.LookupObjectBucketPolicyArgs{
//				Bucket: "bucket.test.com",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupObjectBucketPolicy(ctx *pulumi.Context, args *LookupObjectBucketPolicyArgs, opts ...pulumi.InvokeOption) (*LookupObjectBucketPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupObjectBucketPolicyResult
	err := ctx.Invoke("scaleway:index/getObjectBucketPolicy:getObjectBucketPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getObjectBucketPolicy.
type LookupObjectBucketPolicyArgs struct {
	// The bucket name.
	Bucket string `pulumi:"bucket"`
	// `projectId`) The ID of the project the bucket is associated with.
	ProjectId *string `pulumi:"projectId"`
	// `region`) The region in which the Object Storage exists.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getObjectBucketPolicy.
type LookupObjectBucketPolicyResult struct {
	Bucket string `pulumi:"bucket"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The bucket's policy in JSON format.
	Policy    string  `pulumi:"policy"`
	ProjectId *string `pulumi:"projectId"`
	Region    *string `pulumi:"region"`
}

func LookupObjectBucketPolicyOutput(ctx *pulumi.Context, args LookupObjectBucketPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupObjectBucketPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupObjectBucketPolicyResult, error) {
			args := v.(LookupObjectBucketPolicyArgs)
			r, err := LookupObjectBucketPolicy(ctx, &args, opts...)
			var s LookupObjectBucketPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupObjectBucketPolicyResultOutput)
}

// A collection of arguments for invoking getObjectBucketPolicy.
type LookupObjectBucketPolicyOutputArgs struct {
	// The bucket name.
	Bucket pulumi.StringInput `pulumi:"bucket"`
	// `projectId`) The ID of the project the bucket is associated with.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	// `region`) The region in which the Object Storage exists.
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (LookupObjectBucketPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupObjectBucketPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getObjectBucketPolicy.
type LookupObjectBucketPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupObjectBucketPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupObjectBucketPolicyResult)(nil)).Elem()
}

func (o LookupObjectBucketPolicyResultOutput) ToLookupObjectBucketPolicyResultOutput() LookupObjectBucketPolicyResultOutput {
	return o
}

func (o LookupObjectBucketPolicyResultOutput) ToLookupObjectBucketPolicyResultOutputWithContext(ctx context.Context) LookupObjectBucketPolicyResultOutput {
	return o
}

func (o LookupObjectBucketPolicyResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupObjectBucketPolicyResult] {
	return pulumix.Output[LookupObjectBucketPolicyResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupObjectBucketPolicyResultOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v LookupObjectBucketPolicyResult) string { return v.Bucket }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupObjectBucketPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupObjectBucketPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// The bucket's policy in JSON format.
func (o LookupObjectBucketPolicyResultOutput) Policy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupObjectBucketPolicyResult) string { return v.Policy }).(pulumi.StringOutput)
}

func (o LookupObjectBucketPolicyResultOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupObjectBucketPolicyResult) *string { return v.ProjectId }).(pulumi.StringPtrOutput)
}

func (o LookupObjectBucketPolicyResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupObjectBucketPolicyResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupObjectBucketPolicyResultOutput{})
}
