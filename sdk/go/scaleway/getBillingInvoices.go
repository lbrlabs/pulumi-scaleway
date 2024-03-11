// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Gets information about your Invoices.
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
//			_, err := scaleway.GetBillingInvoices(ctx, &scaleway.GetBillingInvoicesArgs{
//				InvoiceType: pulumi.StringRef("periodic"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetBillingInvoices(ctx *pulumi.Context, args *GetBillingInvoicesArgs, opts ...pulumi.InvokeOption) (*GetBillingInvoicesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetBillingInvoicesResult
	err := ctx.Invoke("scaleway:index/getBillingInvoices:getBillingInvoices", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBillingInvoices.
type GetBillingInvoicesArgs struct {
	// Invoices with the given type are listed. Valid values are `periodic` and `purchase`.
	InvoiceType *string `pulumi:"invoiceType"`
	// Invoices with a start date that are greater or equal to `startedAfter` are listed (RFC 3339 format).
	StartedAfter *string `pulumi:"startedAfter"`
	// Invoices with a start date that precedes `startedBefore` are listed (RFC 3339 format).
	StartedBefore *string `pulumi:"startedBefore"`
}

// A collection of values returned by getBillingInvoices.
type GetBillingInvoicesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The type of invoice.
	InvoiceType *string `pulumi:"invoiceType"`
	// List of found invoices
	Invoices       []GetBillingInvoicesInvoice `pulumi:"invoices"`
	OrganizationId string                      `pulumi:"organizationId"`
	StartedAfter   *string                     `pulumi:"startedAfter"`
	StartedBefore  *string                     `pulumi:"startedBefore"`
}

func GetBillingInvoicesOutput(ctx *pulumi.Context, args GetBillingInvoicesOutputArgs, opts ...pulumi.InvokeOption) GetBillingInvoicesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetBillingInvoicesResult, error) {
			args := v.(GetBillingInvoicesArgs)
			r, err := GetBillingInvoices(ctx, &args, opts...)
			var s GetBillingInvoicesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetBillingInvoicesResultOutput)
}

// A collection of arguments for invoking getBillingInvoices.
type GetBillingInvoicesOutputArgs struct {
	// Invoices with the given type are listed. Valid values are `periodic` and `purchase`.
	InvoiceType pulumi.StringPtrInput `pulumi:"invoiceType"`
	// Invoices with a start date that are greater or equal to `startedAfter` are listed (RFC 3339 format).
	StartedAfter pulumi.StringPtrInput `pulumi:"startedAfter"`
	// Invoices with a start date that precedes `startedBefore` are listed (RFC 3339 format).
	StartedBefore pulumi.StringPtrInput `pulumi:"startedBefore"`
}

func (GetBillingInvoicesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBillingInvoicesArgs)(nil)).Elem()
}

// A collection of values returned by getBillingInvoices.
type GetBillingInvoicesResultOutput struct{ *pulumi.OutputState }

func (GetBillingInvoicesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBillingInvoicesResult)(nil)).Elem()
}

func (o GetBillingInvoicesResultOutput) ToGetBillingInvoicesResultOutput() GetBillingInvoicesResultOutput {
	return o
}

func (o GetBillingInvoicesResultOutput) ToGetBillingInvoicesResultOutputWithContext(ctx context.Context) GetBillingInvoicesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetBillingInvoicesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetBillingInvoicesResult) string { return v.Id }).(pulumi.StringOutput)
}

// The type of invoice.
func (o GetBillingInvoicesResultOutput) InvoiceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBillingInvoicesResult) *string { return v.InvoiceType }).(pulumi.StringPtrOutput)
}

// List of found invoices
func (o GetBillingInvoicesResultOutput) Invoices() GetBillingInvoicesInvoiceArrayOutput {
	return o.ApplyT(func(v GetBillingInvoicesResult) []GetBillingInvoicesInvoice { return v.Invoices }).(GetBillingInvoicesInvoiceArrayOutput)
}

func (o GetBillingInvoicesResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v GetBillingInvoicesResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o GetBillingInvoicesResultOutput) StartedAfter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBillingInvoicesResult) *string { return v.StartedAfter }).(pulumi.StringPtrOutput)
}

func (o GetBillingInvoicesResultOutput) StartedBefore() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBillingInvoicesResult) *string { return v.StartedBefore }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetBillingInvoicesResultOutput{})
}
