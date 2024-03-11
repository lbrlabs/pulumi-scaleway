// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Gets information about IP managed by IPAM service. IPAM service is used for dhcp bundled in VPCs' private networks.
//
// ## Examples
//
// ### Instance Private Network IP
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
//			nic, err := scaleway.NewInstancePrivateNic(ctx, "nic", &scaleway.InstancePrivateNicArgs{
//				ServerId:         pulumi.Any(scaleway_instance_server.Server.Id),
//				PrivateNetworkId: pulumi.Any(scaleway_vpc_private_network.Pn.Id),
//			})
//			if err != nil {
//				return err
//			}
//			_ = scaleway.GetIpamIpOutput(ctx, scaleway.GetIpamIpOutputArgs{
//				MacAddress: nic.MacAddress,
//				Type:       pulumi.String("ipv4"),
//			}, nil)
//			_ = scaleway.GetIpamIpOutput(ctx, scaleway.GetIpamIpOutputArgs{
//				Resource: &scaleway.GetIpamIpResourceArgs{
//					Id:   nic.ID(),
//					Type: pulumi.String("instance_private_nic"),
//				},
//				Type: pulumi.String("ipv4"),
//			}, nil)
//			return nil
//		})
//	}
//
// ```
func GetIpamIp(ctx *pulumi.Context, args *GetIpamIpArgs, opts ...pulumi.InvokeOption) (*GetIpamIpResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetIpamIpResult
	err := ctx.Invoke("scaleway:index/getIpamIp:getIpamIp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIpamIp.
type GetIpamIpArgs struct {
	// The Mac Address linked to the IP.
	MacAddress *string `pulumi:"macAddress"`
	// The ID of the private network the IP belong to.
	PrivateNetworkId *string `pulumi:"privateNetworkId"`
	// `region`) The region in which the IP exists.
	Region *string `pulumi:"region"`
	// Filter by resource ID and type, both attributes must be set
	Resource *GetIpamIpResource `pulumi:"resource"`
	// The type of the resource to get the IP from. [Documentation](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/ipam/v1alpha1#pkg-constants) with type list.
	Type string `pulumi:"type"`
}

// A collection of values returned by getIpamIp.
type GetIpamIpResult struct {
	// The IP address
	Address string `pulumi:"address"`
	// The provider-assigned unique ID for this managed resource.
	Id               string             `pulumi:"id"`
	MacAddress       *string            `pulumi:"macAddress"`
	PrivateNetworkId *string            `pulumi:"privateNetworkId"`
	Region           string             `pulumi:"region"`
	Resource         *GetIpamIpResource `pulumi:"resource"`
	Type             string             `pulumi:"type"`
}

func GetIpamIpOutput(ctx *pulumi.Context, args GetIpamIpOutputArgs, opts ...pulumi.InvokeOption) GetIpamIpResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetIpamIpResult, error) {
			args := v.(GetIpamIpArgs)
			r, err := GetIpamIp(ctx, &args, opts...)
			var s GetIpamIpResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetIpamIpResultOutput)
}

// A collection of arguments for invoking getIpamIp.
type GetIpamIpOutputArgs struct {
	// The Mac Address linked to the IP.
	MacAddress pulumi.StringPtrInput `pulumi:"macAddress"`
	// The ID of the private network the IP belong to.
	PrivateNetworkId pulumi.StringPtrInput `pulumi:"privateNetworkId"`
	// `region`) The region in which the IP exists.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// Filter by resource ID and type, both attributes must be set
	Resource GetIpamIpResourcePtrInput `pulumi:"resource"`
	// The type of the resource to get the IP from. [Documentation](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/ipam/v1alpha1#pkg-constants) with type list.
	Type pulumi.StringInput `pulumi:"type"`
}

func (GetIpamIpOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIpamIpArgs)(nil)).Elem()
}

// A collection of values returned by getIpamIp.
type GetIpamIpResultOutput struct{ *pulumi.OutputState }

func (GetIpamIpResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIpamIpResult)(nil)).Elem()
}

func (o GetIpamIpResultOutput) ToGetIpamIpResultOutput() GetIpamIpResultOutput {
	return o
}

func (o GetIpamIpResultOutput) ToGetIpamIpResultOutputWithContext(ctx context.Context) GetIpamIpResultOutput {
	return o
}

// The IP address
func (o GetIpamIpResultOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v GetIpamIpResult) string { return v.Address }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetIpamIpResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetIpamIpResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetIpamIpResultOutput) MacAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIpamIpResult) *string { return v.MacAddress }).(pulumi.StringPtrOutput)
}

func (o GetIpamIpResultOutput) PrivateNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIpamIpResult) *string { return v.PrivateNetworkId }).(pulumi.StringPtrOutput)
}

func (o GetIpamIpResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetIpamIpResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o GetIpamIpResultOutput) Resource() GetIpamIpResourcePtrOutput {
	return o.ApplyT(func(v GetIpamIpResult) *GetIpamIpResource { return v.Resource }).(GetIpamIpResourcePtrOutput)
}

func (o GetIpamIpResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetIpamIpResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIpamIpResultOutput{})
}