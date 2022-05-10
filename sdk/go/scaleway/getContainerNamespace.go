// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupContainerNamespace(ctx *pulumi.Context, args *LookupContainerNamespaceArgs, opts ...pulumi.InvokeOption) (*LookupContainerNamespaceResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupContainerNamespaceResult
	err := ctx.Invoke("scaleway:index/getContainerNamespace:getContainerNamespace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getContainerNamespace.
type LookupContainerNamespaceArgs struct {
	Name        *string `pulumi:"name"`
	NamespaceId *string `pulumi:"namespaceId"`
	Region      *string `pulumi:"region"`
}

// A collection of values returned by getContainerNamespace.
type LookupContainerNamespaceResult struct {
	Description          string            `pulumi:"description"`
	EnvironmentVariables map[string]string `pulumi:"environmentVariables"`
	// The provider-assigned unique ID for this managed resource.
	Id                  string  `pulumi:"id"`
	Name                *string `pulumi:"name"`
	NamespaceId         *string `pulumi:"namespaceId"`
	OrganizationId      string  `pulumi:"organizationId"`
	ProjectId           string  `pulumi:"projectId"`
	Region              *string `pulumi:"region"`
	RegistryEndpoint    string  `pulumi:"registryEndpoint"`
	RegistryNamespaceId string  `pulumi:"registryNamespaceId"`
}

func LookupContainerNamespaceOutput(ctx *pulumi.Context, args LookupContainerNamespaceOutputArgs, opts ...pulumi.InvokeOption) LookupContainerNamespaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupContainerNamespaceResult, error) {
			args := v.(LookupContainerNamespaceArgs)
			r, err := LookupContainerNamespace(ctx, &args, opts...)
			return *r, err
		}).(LookupContainerNamespaceResultOutput)
}

// A collection of arguments for invoking getContainerNamespace.
type LookupContainerNamespaceOutputArgs struct {
	Name        pulumi.StringPtrInput `pulumi:"name"`
	NamespaceId pulumi.StringPtrInput `pulumi:"namespaceId"`
	Region      pulumi.StringPtrInput `pulumi:"region"`
}

func (LookupContainerNamespaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContainerNamespaceArgs)(nil)).Elem()
}

// A collection of values returned by getContainerNamespace.
type LookupContainerNamespaceResultOutput struct{ *pulumi.OutputState }

func (LookupContainerNamespaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContainerNamespaceResult)(nil)).Elem()
}

func (o LookupContainerNamespaceResultOutput) ToLookupContainerNamespaceResultOutput() LookupContainerNamespaceResultOutput {
	return o
}

func (o LookupContainerNamespaceResultOutput) ToLookupContainerNamespaceResultOutputWithContext(ctx context.Context) LookupContainerNamespaceResultOutput {
	return o
}

func (o LookupContainerNamespaceResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerNamespaceResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupContainerNamespaceResultOutput) EnvironmentVariables() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupContainerNamespaceResult) map[string]string { return v.EnvironmentVariables }).(pulumi.StringMapOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupContainerNamespaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerNamespaceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupContainerNamespaceResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContainerNamespaceResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupContainerNamespaceResultOutput) NamespaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContainerNamespaceResult) *string { return v.NamespaceId }).(pulumi.StringPtrOutput)
}

func (o LookupContainerNamespaceResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerNamespaceResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o LookupContainerNamespaceResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerNamespaceResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o LookupContainerNamespaceResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContainerNamespaceResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o LookupContainerNamespaceResultOutput) RegistryEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerNamespaceResult) string { return v.RegistryEndpoint }).(pulumi.StringOutput)
}

func (o LookupContainerNamespaceResultOutput) RegistryNamespaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerNamespaceResult) string { return v.RegistryNamespaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupContainerNamespaceResultOutput{})
}
