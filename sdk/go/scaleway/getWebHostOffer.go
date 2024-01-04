// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/lbrlabs/pulumi-scaleway/sdk/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about a webhosting offer.
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
//			_, err := scaleway.GetWebHostOffer(ctx, &scaleway.GetWebHostOfferArgs{
//				Name: pulumi.StringRef("performance"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.GetWebHostOffer(ctx, &scaleway.GetWebHostOfferArgs{
//				OfferId: pulumi.StringRef("de2426b4-a9e9-11ec-b909-0242ac120002"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetWebHostOffer(ctx *pulumi.Context, args *GetWebHostOfferArgs, opts ...pulumi.InvokeOption) (*GetWebHostOfferResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetWebHostOfferResult
	err := ctx.Invoke("scaleway:index/getWebHostOffer:getWebHostOffer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getWebHostOffer.
type GetWebHostOfferArgs struct {
	// The offer name. Only one of `name` and `offerId` should be specified.
	Name *string `pulumi:"name"`
	// The offer id. Only one of `name` and `offerId` should be specified.
	OfferId *string `pulumi:"offerId"`
	// `region`) The region in which offer exists.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getWebHostOffer.
type GetWebHostOfferResult struct {
	// The unique identifier used for billing.
	BillingOperationPath string `pulumi:"billingOperationPath"`
	// The provider-assigned unique ID for this managed resource.
	Id      string  `pulumi:"id"`
	Name    *string `pulumi:"name"`
	OfferId *string `pulumi:"offerId"`
	// The offer price.
	Price string `pulumi:"price"`
	// The offer product.
	Products []GetWebHostOfferProduct `pulumi:"products"`
	Region   string                   `pulumi:"region"`
}

func GetWebHostOfferOutput(ctx *pulumi.Context, args GetWebHostOfferOutputArgs, opts ...pulumi.InvokeOption) GetWebHostOfferResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetWebHostOfferResult, error) {
			args := v.(GetWebHostOfferArgs)
			r, err := GetWebHostOffer(ctx, &args, opts...)
			var s GetWebHostOfferResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetWebHostOfferResultOutput)
}

// A collection of arguments for invoking getWebHostOffer.
type GetWebHostOfferOutputArgs struct {
	// The offer name. Only one of `name` and `offerId` should be specified.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The offer id. Only one of `name` and `offerId` should be specified.
	OfferId pulumi.StringPtrInput `pulumi:"offerId"`
	// `region`) The region in which offer exists.
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (GetWebHostOfferOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetWebHostOfferArgs)(nil)).Elem()
}

// A collection of values returned by getWebHostOffer.
type GetWebHostOfferResultOutput struct{ *pulumi.OutputState }

func (GetWebHostOfferResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetWebHostOfferResult)(nil)).Elem()
}

func (o GetWebHostOfferResultOutput) ToGetWebHostOfferResultOutput() GetWebHostOfferResultOutput {
	return o
}

func (o GetWebHostOfferResultOutput) ToGetWebHostOfferResultOutputWithContext(ctx context.Context) GetWebHostOfferResultOutput {
	return o
}

// The unique identifier used for billing.
func (o GetWebHostOfferResultOutput) BillingOperationPath() pulumi.StringOutput {
	return o.ApplyT(func(v GetWebHostOfferResult) string { return v.BillingOperationPath }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetWebHostOfferResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetWebHostOfferResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetWebHostOfferResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetWebHostOfferResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetWebHostOfferResultOutput) OfferId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetWebHostOfferResult) *string { return v.OfferId }).(pulumi.StringPtrOutput)
}

// The offer price.
func (o GetWebHostOfferResultOutput) Price() pulumi.StringOutput {
	return o.ApplyT(func(v GetWebHostOfferResult) string { return v.Price }).(pulumi.StringOutput)
}

// The offer product.
func (o GetWebHostOfferResultOutput) Products() GetWebHostOfferProductArrayOutput {
	return o.ApplyT(func(v GetWebHostOfferResult) []GetWebHostOfferProduct { return v.Products }).(GetWebHostOfferProductArrayOutput)
}

func (o GetWebHostOfferResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetWebHostOfferResult) string { return v.Region }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetWebHostOfferResultOutput{})
}
