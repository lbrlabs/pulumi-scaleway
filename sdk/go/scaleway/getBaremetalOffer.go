// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetBaremetalOffer(ctx *pulumi.Context, args *GetBaremetalOfferArgs, opts ...pulumi.InvokeOption) (*GetBaremetalOfferResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetBaremetalOfferResult
	err := ctx.Invoke("scaleway:index/getBaremetalOffer:getBaremetalOffer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBaremetalOffer.
type GetBaremetalOfferArgs struct {
	IncludeDisabled *bool   `pulumi:"includeDisabled"`
	Name            *string `pulumi:"name"`
	OfferId         *string `pulumi:"offerId"`
	Zone            *string `pulumi:"zone"`
}

// A collection of values returned by getBaremetalOffer.
type GetBaremetalOfferResult struct {
	Bandwidth       int                     `pulumi:"bandwidth"`
	CommercialRange string                  `pulumi:"commercialRange"`
	Cpu             GetBaremetalOfferCpu    `pulumi:"cpu"`
	Disks           []GetBaremetalOfferDisk `pulumi:"disks"`
	// The provider-assigned unique ID for this managed resource.
	Id              string                    `pulumi:"id"`
	IncludeDisabled *bool                     `pulumi:"includeDisabled"`
	Memories        []GetBaremetalOfferMemory `pulumi:"memories"`
	Name            *string                   `pulumi:"name"`
	OfferId         *string                   `pulumi:"offerId"`
	Stock           string                    `pulumi:"stock"`
	Zone            string                    `pulumi:"zone"`
}

func GetBaremetalOfferOutput(ctx *pulumi.Context, args GetBaremetalOfferOutputArgs, opts ...pulumi.InvokeOption) GetBaremetalOfferResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetBaremetalOfferResult, error) {
			args := v.(GetBaremetalOfferArgs)
			r, err := GetBaremetalOffer(ctx, &args, opts...)
			return *r, err
		}).(GetBaremetalOfferResultOutput)
}

// A collection of arguments for invoking getBaremetalOffer.
type GetBaremetalOfferOutputArgs struct {
	IncludeDisabled pulumi.BoolPtrInput   `pulumi:"includeDisabled"`
	Name            pulumi.StringPtrInput `pulumi:"name"`
	OfferId         pulumi.StringPtrInput `pulumi:"offerId"`
	Zone            pulumi.StringPtrInput `pulumi:"zone"`
}

func (GetBaremetalOfferOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBaremetalOfferArgs)(nil)).Elem()
}

// A collection of values returned by getBaremetalOffer.
type GetBaremetalOfferResultOutput struct{ *pulumi.OutputState }

func (GetBaremetalOfferResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBaremetalOfferResult)(nil)).Elem()
}

func (o GetBaremetalOfferResultOutput) ToGetBaremetalOfferResultOutput() GetBaremetalOfferResultOutput {
	return o
}

func (o GetBaremetalOfferResultOutput) ToGetBaremetalOfferResultOutputWithContext(ctx context.Context) GetBaremetalOfferResultOutput {
	return o
}

func (o GetBaremetalOfferResultOutput) Bandwidth() pulumi.IntOutput {
	return o.ApplyT(func(v GetBaremetalOfferResult) int { return v.Bandwidth }).(pulumi.IntOutput)
}

func (o GetBaremetalOfferResultOutput) CommercialRange() pulumi.StringOutput {
	return o.ApplyT(func(v GetBaremetalOfferResult) string { return v.CommercialRange }).(pulumi.StringOutput)
}

func (o GetBaremetalOfferResultOutput) Cpu() GetBaremetalOfferCpuOutput {
	return o.ApplyT(func(v GetBaremetalOfferResult) GetBaremetalOfferCpu { return v.Cpu }).(GetBaremetalOfferCpuOutput)
}

func (o GetBaremetalOfferResultOutput) Disks() GetBaremetalOfferDiskArrayOutput {
	return o.ApplyT(func(v GetBaremetalOfferResult) []GetBaremetalOfferDisk { return v.Disks }).(GetBaremetalOfferDiskArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetBaremetalOfferResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetBaremetalOfferResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetBaremetalOfferResultOutput) IncludeDisabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetBaremetalOfferResult) *bool { return v.IncludeDisabled }).(pulumi.BoolPtrOutput)
}

func (o GetBaremetalOfferResultOutput) Memories() GetBaremetalOfferMemoryArrayOutput {
	return o.ApplyT(func(v GetBaremetalOfferResult) []GetBaremetalOfferMemory { return v.Memories }).(GetBaremetalOfferMemoryArrayOutput)
}

func (o GetBaremetalOfferResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBaremetalOfferResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetBaremetalOfferResultOutput) OfferId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBaremetalOfferResult) *string { return v.OfferId }).(pulumi.StringPtrOutput)
}

func (o GetBaremetalOfferResultOutput) Stock() pulumi.StringOutput {
	return o.ApplyT(func(v GetBaremetalOfferResult) string { return v.Stock }).(pulumi.StringOutput)
}

func (o GetBaremetalOfferResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v GetBaremetalOfferResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetBaremetalOfferResultOutput{})
}
