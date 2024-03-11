// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Gets information about a function namespace.
//
// ## Example Usage
//
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
//			_, err := scaleway.LookupFunctionNamespace(ctx, &scaleway.LookupFunctionNamespaceArgs{
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
func LookupFunctionNamespace(ctx *pulumi.Context, args *LookupFunctionNamespaceArgs, opts ...pulumi.InvokeOption) (*LookupFunctionNamespaceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFunctionNamespaceResult
	err := ctx.Invoke("scaleway:index/getFunctionNamespace:getFunctionNamespace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFunctionNamespace.
type LookupFunctionNamespaceArgs struct {
	// The namespace name.
	// Only one of `name` and `namespaceId` should be specified.
	Name *string `pulumi:"name"`
	// The namespace id.
	// Only one of `name` and `namespaceId` should be specified.
	NamespaceId *string `pulumi:"namespaceId"`
	// `region`) The region in which the namespace exists.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getFunctionNamespace.
type LookupFunctionNamespaceResult struct {
	// The description of the namespace.
	Description string `pulumi:"description"`
	// The environment variables of the namespace.
	EnvironmentVariables map[string]string `pulumi:"environmentVariables"`
	// The provider-assigned unique ID for this managed resource.
	Id          string  `pulumi:"id"`
	Name        *string `pulumi:"name"`
	NamespaceId *string `pulumi:"namespaceId"`
	// The organization ID the namespace is associated with.
	OrganizationId string  `pulumi:"organizationId"`
	ProjectId      string  `pulumi:"projectId"`
	Region         *string `pulumi:"region"`
	// The registry endpoint of the namespace.
	RegistryEndpoint string `pulumi:"registryEndpoint"`
	// The registry namespace ID of the namespace.
	RegistryNamespaceId        string            `pulumi:"registryNamespaceId"`
	SecretEnvironmentVariables map[string]string `pulumi:"secretEnvironmentVariables"`
}

func LookupFunctionNamespaceOutput(ctx *pulumi.Context, args LookupFunctionNamespaceOutputArgs, opts ...pulumi.InvokeOption) LookupFunctionNamespaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFunctionNamespaceResult, error) {
			args := v.(LookupFunctionNamespaceArgs)
			r, err := LookupFunctionNamespace(ctx, &args, opts...)
			var s LookupFunctionNamespaceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFunctionNamespaceResultOutput)
}

// A collection of arguments for invoking getFunctionNamespace.
type LookupFunctionNamespaceOutputArgs struct {
	// The namespace name.
	// Only one of `name` and `namespaceId` should be specified.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The namespace id.
	// Only one of `name` and `namespaceId` should be specified.
	NamespaceId pulumi.StringPtrInput `pulumi:"namespaceId"`
	// `region`) The region in which the namespace exists.
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (LookupFunctionNamespaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFunctionNamespaceArgs)(nil)).Elem()
}

// A collection of values returned by getFunctionNamespace.
type LookupFunctionNamespaceResultOutput struct{ *pulumi.OutputState }

func (LookupFunctionNamespaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFunctionNamespaceResult)(nil)).Elem()
}

func (o LookupFunctionNamespaceResultOutput) ToLookupFunctionNamespaceResultOutput() LookupFunctionNamespaceResultOutput {
	return o
}

func (o LookupFunctionNamespaceResultOutput) ToLookupFunctionNamespaceResultOutputWithContext(ctx context.Context) LookupFunctionNamespaceResultOutput {
	return o
}

// The description of the namespace.
func (o LookupFunctionNamespaceResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionNamespaceResult) string { return v.Description }).(pulumi.StringOutput)
}

// The environment variables of the namespace.
func (o LookupFunctionNamespaceResultOutput) EnvironmentVariables() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupFunctionNamespaceResult) map[string]string { return v.EnvironmentVariables }).(pulumi.StringMapOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFunctionNamespaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionNamespaceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFunctionNamespaceResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFunctionNamespaceResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupFunctionNamespaceResultOutput) NamespaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFunctionNamespaceResult) *string { return v.NamespaceId }).(pulumi.StringPtrOutput)
}

// The organization ID the namespace is associated with.
func (o LookupFunctionNamespaceResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionNamespaceResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o LookupFunctionNamespaceResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionNamespaceResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o LookupFunctionNamespaceResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFunctionNamespaceResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

// The registry endpoint of the namespace.
func (o LookupFunctionNamespaceResultOutput) RegistryEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionNamespaceResult) string { return v.RegistryEndpoint }).(pulumi.StringOutput)
}

// The registry namespace ID of the namespace.
func (o LookupFunctionNamespaceResultOutput) RegistryNamespaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionNamespaceResult) string { return v.RegistryNamespaceId }).(pulumi.StringOutput)
}

func (o LookupFunctionNamespaceResultOutput) SecretEnvironmentVariables() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupFunctionNamespaceResult) map[string]string { return v.SecretEnvironmentVariables }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFunctionNamespaceResultOutput{})
}
