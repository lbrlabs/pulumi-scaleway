// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about a baremetal server.
// For more information, see [the documentation](https://developers.scaleway.com/en/products/baremetal/api).
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
//			_, err = scaleway.LookupBaremetalServer(ctx, &scaleway.LookupBaremetalServerArgs{
//				Name: pulumi.StringRef("foobar"),
//				Zone: pulumi.StringRef("fr-par-2"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.LookupBaremetalServer(ctx, &scaleway.LookupBaremetalServerArgs{
//				ServerId: pulumi.StringRef("11111111-1111-1111-1111-111111111111"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupBaremetalServer(ctx *pulumi.Context, args *LookupBaremetalServerArgs, opts ...pulumi.InvokeOption) (*LookupBaremetalServerResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupBaremetalServerResult
	err := ctx.Invoke("scaleway:index/getBaremetalServer:getBaremetalServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBaremetalServer.
type LookupBaremetalServerArgs struct {
	// The server name. Only one of `name` and `serverId` should be specified.
	Name     *string `pulumi:"name"`
	ServerId *string `pulumi:"serverId"`
	// `zone`) The zone in which the server exists.
	Zone *string `pulumi:"zone"`
}

// A collection of values returned by getBaremetalServer.
type LookupBaremetalServerResult struct {
	Description string `pulumi:"description"`
	Domain      string `pulumi:"domain"`
	Hostname    string `pulumi:"hostname"`
	// The provider-assigned unique ID for this managed resource.
	Id                       string                     `pulumi:"id"`
	Ips                      []GetBaremetalServerIp     `pulumi:"ips"`
	Name                     *string                    `pulumi:"name"`
	Offer                    string                     `pulumi:"offer"`
	OfferId                  string                     `pulumi:"offerId"`
	Options                  []GetBaremetalServerOption `pulumi:"options"`
	OrganizationId           string                     `pulumi:"organizationId"`
	Os                       string                     `pulumi:"os"`
	OsId                     string                     `pulumi:"osId"`
	Password                 string                     `pulumi:"password"`
	ProjectId                string                     `pulumi:"projectId"`
	ReinstallOnConfigChanges bool                       `pulumi:"reinstallOnConfigChanges"`
	ServerId                 *string                    `pulumi:"serverId"`
	ServicePassword          string                     `pulumi:"servicePassword"`
	ServiceUser              string                     `pulumi:"serviceUser"`
	SshKeyIds                []string                   `pulumi:"sshKeyIds"`
	Tags                     []string                   `pulumi:"tags"`
	User                     string                     `pulumi:"user"`
	Zone                     *string                    `pulumi:"zone"`
}

func LookupBaremetalServerOutput(ctx *pulumi.Context, args LookupBaremetalServerOutputArgs, opts ...pulumi.InvokeOption) LookupBaremetalServerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBaremetalServerResult, error) {
			args := v.(LookupBaremetalServerArgs)
			r, err := LookupBaremetalServer(ctx, &args, opts...)
			var s LookupBaremetalServerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBaremetalServerResultOutput)
}

// A collection of arguments for invoking getBaremetalServer.
type LookupBaremetalServerOutputArgs struct {
	// The server name. Only one of `name` and `serverId` should be specified.
	Name     pulumi.StringPtrInput `pulumi:"name"`
	ServerId pulumi.StringPtrInput `pulumi:"serverId"`
	// `zone`) The zone in which the server exists.
	Zone pulumi.StringPtrInput `pulumi:"zone"`
}

func (LookupBaremetalServerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBaremetalServerArgs)(nil)).Elem()
}

// A collection of values returned by getBaremetalServer.
type LookupBaremetalServerResultOutput struct{ *pulumi.OutputState }

func (LookupBaremetalServerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBaremetalServerResult)(nil)).Elem()
}

func (o LookupBaremetalServerResultOutput) ToLookupBaremetalServerResultOutput() LookupBaremetalServerResultOutput {
	return o
}

func (o LookupBaremetalServerResultOutput) ToLookupBaremetalServerResultOutputWithContext(ctx context.Context) LookupBaremetalServerResultOutput {
	return o
}

func (o LookupBaremetalServerResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBaremetalServerResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupBaremetalServerResultOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBaremetalServerResult) string { return v.Domain }).(pulumi.StringOutput)
}

func (o LookupBaremetalServerResultOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBaremetalServerResult) string { return v.Hostname }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupBaremetalServerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBaremetalServerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBaremetalServerResultOutput) Ips() GetBaremetalServerIpArrayOutput {
	return o.ApplyT(func(v LookupBaremetalServerResult) []GetBaremetalServerIp { return v.Ips }).(GetBaremetalServerIpArrayOutput)
}

func (o LookupBaremetalServerResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBaremetalServerResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupBaremetalServerResultOutput) Offer() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBaremetalServerResult) string { return v.Offer }).(pulumi.StringOutput)
}

func (o LookupBaremetalServerResultOutput) OfferId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBaremetalServerResult) string { return v.OfferId }).(pulumi.StringOutput)
}

func (o LookupBaremetalServerResultOutput) Options() GetBaremetalServerOptionArrayOutput {
	return o.ApplyT(func(v LookupBaremetalServerResult) []GetBaremetalServerOption { return v.Options }).(GetBaremetalServerOptionArrayOutput)
}

func (o LookupBaremetalServerResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBaremetalServerResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o LookupBaremetalServerResultOutput) Os() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBaremetalServerResult) string { return v.Os }).(pulumi.StringOutput)
}

func (o LookupBaremetalServerResultOutput) OsId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBaremetalServerResult) string { return v.OsId }).(pulumi.StringOutput)
}

func (o LookupBaremetalServerResultOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBaremetalServerResult) string { return v.Password }).(pulumi.StringOutput)
}

func (o LookupBaremetalServerResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBaremetalServerResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o LookupBaremetalServerResultOutput) ReinstallOnConfigChanges() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupBaremetalServerResult) bool { return v.ReinstallOnConfigChanges }).(pulumi.BoolOutput)
}

func (o LookupBaremetalServerResultOutput) ServerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBaremetalServerResult) *string { return v.ServerId }).(pulumi.StringPtrOutput)
}

func (o LookupBaremetalServerResultOutput) ServicePassword() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBaremetalServerResult) string { return v.ServicePassword }).(pulumi.StringOutput)
}

func (o LookupBaremetalServerResultOutput) ServiceUser() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBaremetalServerResult) string { return v.ServiceUser }).(pulumi.StringOutput)
}

func (o LookupBaremetalServerResultOutput) SshKeyIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupBaremetalServerResult) []string { return v.SshKeyIds }).(pulumi.StringArrayOutput)
}

func (o LookupBaremetalServerResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupBaremetalServerResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o LookupBaremetalServerResultOutput) User() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBaremetalServerResult) string { return v.User }).(pulumi.StringOutput)
}

func (o LookupBaremetalServerResultOutput) Zone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBaremetalServerResult) *string { return v.Zone }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBaremetalServerResultOutput{})
}
