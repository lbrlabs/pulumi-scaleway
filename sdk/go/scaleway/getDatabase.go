// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Gets information about a RDB database.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := scaleway.LookupDatabase(ctx, &scaleway.LookupDatabaseArgs{
//				InstanceId: "11111111-1111-1111-1111-111111111111",
//				Name:       "foobar",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func LookupDatabase(ctx *pulumi.Context, args *LookupDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDatabaseResult
	err := ctx.Invoke("scaleway:index/getDatabase:getDatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDatabase.
type LookupDatabaseArgs struct {
	// The RDB instance ID.
	InstanceId string `pulumi:"instanceId"`
	// The name of the RDB instance.
	Name   string  `pulumi:"name"`
	Region *string `pulumi:"region"`
}

// A collection of values returned by getDatabase.
type LookupDatabaseResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	InstanceId string `pulumi:"instanceId"`
	// Whether the database is managed or not.
	Managed bool   `pulumi:"managed"`
	Name    string `pulumi:"name"`
	// The name of the owner of the database.
	Owner  string  `pulumi:"owner"`
	Region *string `pulumi:"region"`
	// Size of the database (in bytes).
	Size string `pulumi:"size"`
}

func LookupDatabaseOutput(ctx *pulumi.Context, args LookupDatabaseOutputArgs, opts ...pulumi.InvokeOption) LookupDatabaseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatabaseResult, error) {
			args := v.(LookupDatabaseArgs)
			r, err := LookupDatabase(ctx, &args, opts...)
			var s LookupDatabaseResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDatabaseResultOutput)
}

// A collection of arguments for invoking getDatabase.
type LookupDatabaseOutputArgs struct {
	// The RDB instance ID.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// The name of the RDB instance.
	Name   pulumi.StringInput    `pulumi:"name"`
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (LookupDatabaseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseArgs)(nil)).Elem()
}

// A collection of values returned by getDatabase.
type LookupDatabaseResultOutput struct{ *pulumi.OutputState }

func (LookupDatabaseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseResult)(nil)).Elem()
}

func (o LookupDatabaseResultOutput) ToLookupDatabaseResultOutput() LookupDatabaseResultOutput {
	return o
}

func (o LookupDatabaseResultOutput) ToLookupDatabaseResultOutputWithContext(ctx context.Context) LookupDatabaseResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDatabaseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDatabaseResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

// Whether the database is managed or not.
func (o LookupDatabaseResultOutput) Managed() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDatabaseResult) bool { return v.Managed }).(pulumi.BoolOutput)
}

func (o LookupDatabaseResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Name }).(pulumi.StringOutput)
}

// The name of the owner of the database.
func (o LookupDatabaseResultOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Owner }).(pulumi.StringOutput)
}

func (o LookupDatabaseResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

// Size of the database (in bytes).
func (o LookupDatabaseResultOutput) Size() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Size }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatabaseResultOutput{})
}
