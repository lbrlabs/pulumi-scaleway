// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about a registry namespace.
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
//			_, err := scaleway.LookupRegistryNamespace(ctx, &scaleway.LookupRegistryNamespaceArgs{
//				NamespaceId: pulumi.StringRef("11111111-1111-1111-1111-111111111111"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupRegistryNamespace(ctx *pulumi.Context, args *LookupRegistryNamespaceArgs, opts ...pulumi.InvokeOption) (*LookupRegistryNamespaceResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupRegistryNamespaceResult
	err := ctx.Invoke("scaleway:index/getRegistryNamespace:getRegistryNamespace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRegistryNamespace.
type LookupRegistryNamespaceArgs struct {
	// The namespace name.
	// Only one of `name` and `namespaceId` should be specified.
	Name *string `pulumi:"name"`
	// The namespace id.
	// Only one of `name` and `namespaceId` should be specified.
	NamespaceId *string `pulumi:"namespaceId"`
	// `region`) The region in which the namespace exists.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getRegistryNamespace.
type LookupRegistryNamespaceResult struct {
	Description string `pulumi:"description"`
	// The endpoint of the Registry Namespace.
	Endpoint string `pulumi:"endpoint"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The Namespace Privacy Policy: whether or not the images are public.
	IsPublic    bool    `pulumi:"isPublic"`
	Name        *string `pulumi:"name"`
	NamespaceId *string `pulumi:"namespaceId"`
	// The organization ID the namespace is associated with.
	OrganizationId string  `pulumi:"organizationId"`
	ProjectId      string  `pulumi:"projectId"`
	Region         *string `pulumi:"region"`
}

func LookupRegistryNamespaceOutput(ctx *pulumi.Context, args LookupRegistryNamespaceOutputArgs, opts ...pulumi.InvokeOption) LookupRegistryNamespaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRegistryNamespaceResult, error) {
			args := v.(LookupRegistryNamespaceArgs)
			r, err := LookupRegistryNamespace(ctx, &args, opts...)
			var s LookupRegistryNamespaceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRegistryNamespaceResultOutput)
}

// A collection of arguments for invoking getRegistryNamespace.
type LookupRegistryNamespaceOutputArgs struct {
	// The namespace name.
	// Only one of `name` and `namespaceId` should be specified.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The namespace id.
	// Only one of `name` and `namespaceId` should be specified.
	NamespaceId pulumi.StringPtrInput `pulumi:"namespaceId"`
	// `region`) The region in which the namespace exists.
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (LookupRegistryNamespaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryNamespaceArgs)(nil)).Elem()
}

// A collection of values returned by getRegistryNamespace.
type LookupRegistryNamespaceResultOutput struct{ *pulumi.OutputState }

func (LookupRegistryNamespaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryNamespaceResult)(nil)).Elem()
}

func (o LookupRegistryNamespaceResultOutput) ToLookupRegistryNamespaceResultOutput() LookupRegistryNamespaceResultOutput {
	return o
}

func (o LookupRegistryNamespaceResultOutput) ToLookupRegistryNamespaceResultOutputWithContext(ctx context.Context) LookupRegistryNamespaceResultOutput {
	return o
}

func (o LookupRegistryNamespaceResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryNamespaceResult) string { return v.Description }).(pulumi.StringOutput)
}

// The endpoint of the Registry Namespace.
func (o LookupRegistryNamespaceResultOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryNamespaceResult) string { return v.Endpoint }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRegistryNamespaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryNamespaceResult) string { return v.Id }).(pulumi.StringOutput)
}

// The Namespace Privacy Policy: whether or not the images are public.
func (o LookupRegistryNamespaceResultOutput) IsPublic() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRegistryNamespaceResult) bool { return v.IsPublic }).(pulumi.BoolOutput)
}

func (o LookupRegistryNamespaceResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegistryNamespaceResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupRegistryNamespaceResultOutput) NamespaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegistryNamespaceResult) *string { return v.NamespaceId }).(pulumi.StringPtrOutput)
}

// The organization ID the namespace is associated with.
func (o LookupRegistryNamespaceResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryNamespaceResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o LookupRegistryNamespaceResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryNamespaceResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o LookupRegistryNamespaceResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegistryNamespaceResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRegistryNamespaceResultOutput{})
}
