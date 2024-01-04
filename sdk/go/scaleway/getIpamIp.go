// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/lbrlabs/pulumi-scaleway/sdk/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
//	"github.com/lbrlabs/pulumi-scaleway/sdk/go/scaleway"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			pn, err := scaleway.NewVpcPrivateNetwork(ctx, "pn", nil)
//			if err != nil {
//				return err
//			}
//			nic, err := scaleway.NewInstancePrivateNic(ctx, "nic", &scaleway.InstancePrivateNicArgs{
//				ServerId:         pulumi.Any(scaleway_instance_server.Server.Id),
//				PrivateNetworkId: pn.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_ = scaleway.LookupIpamIpOutput(ctx, scaleway.GetIpamIpOutputArgs{
//				MacAddress: nic.MacAddress,
//				Type:       pulumi.String("ipv4"),
//			}, nil)
//			_ = scaleway.LookupIpamIpOutput(ctx, scaleway.GetIpamIpOutputArgs{
//				Resource: &scaleway.GetIpamIpResourceArgs{
//					Id:   nic.ID(),
//					Type: pulumi.String("instance_private_nic"),
//				},
//				Type: pulumi.String("ipv4"),
//			}, nil)
//			main, err := scaleway.NewDatabaseInstance(ctx, "main", &scaleway.DatabaseInstanceArgs{
//				NodeType:      pulumi.String("DB-DEV-S"),
//				Engine:        pulumi.String("PostgreSQL-15"),
//				IsHaCluster:   pulumi.Bool(true),
//				DisableBackup: pulumi.Bool(true),
//				UserName:      pulumi.String("my_initial_user"),
//				Password:      pulumi.String("thiZ_is_v&ry_s3cret"),
//				PrivateNetwork: &scaleway.DatabaseInstancePrivateNetworkArgs{
//					PnId: pn.ID(),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_ = scaleway.LookupIpamIpOutput(ctx, scaleway.GetIpamIpOutputArgs{
//				Resource: &scaleway.GetIpamIpResourceArgs{
//					Name: main.Name,
//					Type: pulumi.String("rdb_instance"),
//				},
//				Type: pulumi.String("ipv4"),
//			}, nil)
//			return nil
//		})
//	}
//
// ```
func LookupIpamIp(ctx *pulumi.Context, args *LookupIpamIpArgs, opts ...pulumi.InvokeOption) (*LookupIpamIpResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupIpamIpResult
	err := ctx.Invoke("scaleway:index/getIpamIp:getIpamIp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIpamIp.
type LookupIpamIpArgs struct {
	// The Mac Address linked to the IP.
	MacAddress *string `pulumi:"macAddress"`
	// The ID of the private network the IP belong to.
	PrivateNetworkId *string `pulumi:"privateNetworkId"`
	// `projectId`) The ID of the project the IP is associated with.
	ProjectId *string `pulumi:"projectId"`
	// `region`) The region in which the IP exists.
	Region *string `pulumi:"region"`
	// Filter by resource ID, type or name. If specified, `type` is required, and at least one of `id` or `name` must be set.
	Resource *GetIpamIpResource `pulumi:"resource"`
	// The tags associated with the IP.
	// As datasource only returns one IP, the search with given tags must return only one result.
	Tags []string `pulumi:"tags"`
	// The type of the resource to get the IP from. [Documentation](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/ipam/v1#pkg-constants) with type list.
	Type string `pulumi:"type"`
	// Only IPs that are zonal, and in this zone, will be returned.
	Zonal *string `pulumi:"zonal"`
}

// A collection of values returned by getIpamIp.
type LookupIpamIpResult struct {
	// The IP address
	Address string `pulumi:"address"`
	// The provider-assigned unique ID for this managed resource.
	Id               string             `pulumi:"id"`
	MacAddress       *string            `pulumi:"macAddress"`
	OrganizationId   string             `pulumi:"organizationId"`
	PrivateNetworkId *string            `pulumi:"privateNetworkId"`
	ProjectId        string             `pulumi:"projectId"`
	Region           string             `pulumi:"region"`
	Resource         *GetIpamIpResource `pulumi:"resource"`
	Tags             []string           `pulumi:"tags"`
	Type             string             `pulumi:"type"`
	Zonal            string             `pulumi:"zonal"`
}

func LookupIpamIpOutput(ctx *pulumi.Context, args LookupIpamIpOutputArgs, opts ...pulumi.InvokeOption) LookupIpamIpResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIpamIpResult, error) {
			args := v.(LookupIpamIpArgs)
			r, err := LookupIpamIp(ctx, &args, opts...)
			var s LookupIpamIpResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIpamIpResultOutput)
}

// A collection of arguments for invoking getIpamIp.
type LookupIpamIpOutputArgs struct {
	// The Mac Address linked to the IP.
	MacAddress pulumi.StringPtrInput `pulumi:"macAddress"`
	// The ID of the private network the IP belong to.
	PrivateNetworkId pulumi.StringPtrInput `pulumi:"privateNetworkId"`
	// `projectId`) The ID of the project the IP is associated with.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	// `region`) The region in which the IP exists.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// Filter by resource ID, type or name. If specified, `type` is required, and at least one of `id` or `name` must be set.
	Resource GetIpamIpResourcePtrInput `pulumi:"resource"`
	// The tags associated with the IP.
	// As datasource only returns one IP, the search with given tags must return only one result.
	Tags pulumi.StringArrayInput `pulumi:"tags"`
	// The type of the resource to get the IP from. [Documentation](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/ipam/v1#pkg-constants) with type list.
	Type pulumi.StringInput `pulumi:"type"`
	// Only IPs that are zonal, and in this zone, will be returned.
	Zonal pulumi.StringPtrInput `pulumi:"zonal"`
}

func (LookupIpamIpOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIpamIpArgs)(nil)).Elem()
}

// A collection of values returned by getIpamIp.
type LookupIpamIpResultOutput struct{ *pulumi.OutputState }

func (LookupIpamIpResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIpamIpResult)(nil)).Elem()
}

func (o LookupIpamIpResultOutput) ToLookupIpamIpResultOutput() LookupIpamIpResultOutput {
	return o
}

func (o LookupIpamIpResultOutput) ToLookupIpamIpResultOutputWithContext(ctx context.Context) LookupIpamIpResultOutput {
	return o
}

// The IP address
func (o LookupIpamIpResultOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpamIpResult) string { return v.Address }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupIpamIpResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpamIpResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIpamIpResultOutput) MacAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIpamIpResult) *string { return v.MacAddress }).(pulumi.StringPtrOutput)
}

func (o LookupIpamIpResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpamIpResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o LookupIpamIpResultOutput) PrivateNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIpamIpResult) *string { return v.PrivateNetworkId }).(pulumi.StringPtrOutput)
}

func (o LookupIpamIpResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpamIpResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o LookupIpamIpResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpamIpResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o LookupIpamIpResultOutput) Resource() GetIpamIpResourcePtrOutput {
	return o.ApplyT(func(v LookupIpamIpResult) *GetIpamIpResource { return v.Resource }).(GetIpamIpResourcePtrOutput)
}

func (o LookupIpamIpResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupIpamIpResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o LookupIpamIpResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpamIpResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupIpamIpResultOutput) Zonal() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpamIpResult) string { return v.Zonal }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIpamIpResultOutput{})
}
