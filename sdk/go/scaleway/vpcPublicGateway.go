// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates and manages Scaleway VPC Public Gateway.
// For more information, see [the documentation](https://developers.scaleway.com/en/products/vpc-gw/api/v1).
//
// ## Example
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-scaleway/sdk/go/scaleway"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := scaleway.NewVpcPublicGateway(ctx, "main", &scaleway.VpcPublicGatewayArgs{
// 			Tags: pulumi.StringArray{
// 				pulumi.String("demo"),
// 				pulumi.String("terraform"),
// 			},
// 			Type: pulumi.String("VPC-GW-S"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Public gateway can be imported using the `{zone}/{id}`, e.g. bash
//
// ```sh
//  $ pulumi import scaleway:index/vpcPublicGateway:VpcPublicGateway main fr-par-1/11111111-1111-1111-1111-111111111111
// ```
type VpcPublicGateway struct {
	pulumi.CustomResourceState

	// The date and time of the creation of the public gateway.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// attach an existing flexible IP to the gateway
	IpId pulumi.StringOutput `pulumi:"ipId"`
	// The name of the public gateway. If not provided it will be randomly generated.
	Name pulumi.StringOutput `pulumi:"name"`
	// The organization ID the public gateway is associated with.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// `projectId`) The ID of the project the public gateway is associated with.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The tags associated with the public gateway.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The gateway type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The date and time of the last update of the public gateway.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// override the gateway's default recursive DNS servers, if DNS features are enabled.
	UpstreamDnsServers pulumi.StringArrayOutput `pulumi:"upstreamDnsServers"`
	// `zone`) The zone in which the public gateway should be created.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewVpcPublicGateway registers a new resource with the given unique name, arguments, and options.
func NewVpcPublicGateway(ctx *pulumi.Context,
	name string, args *VpcPublicGatewayArgs, opts ...pulumi.ResourceOption) (*VpcPublicGateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource VpcPublicGateway
	err := ctx.RegisterResource("scaleway:index/vpcPublicGateway:VpcPublicGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcPublicGateway gets an existing VpcPublicGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcPublicGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcPublicGatewayState, opts ...pulumi.ResourceOption) (*VpcPublicGateway, error) {
	var resource VpcPublicGateway
	err := ctx.ReadResource("scaleway:index/vpcPublicGateway:VpcPublicGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcPublicGateway resources.
type vpcPublicGatewayState struct {
	// The date and time of the creation of the public gateway.
	CreatedAt *string `pulumi:"createdAt"`
	// attach an existing flexible IP to the gateway
	IpId *string `pulumi:"ipId"`
	// The name of the public gateway. If not provided it will be randomly generated.
	Name *string `pulumi:"name"`
	// The organization ID the public gateway is associated with.
	OrganizationId *string `pulumi:"organizationId"`
	// `projectId`) The ID of the project the public gateway is associated with.
	ProjectId *string `pulumi:"projectId"`
	// The tags associated with the public gateway.
	Tags []string `pulumi:"tags"`
	// The gateway type.
	Type *string `pulumi:"type"`
	// The date and time of the last update of the public gateway.
	UpdatedAt *string `pulumi:"updatedAt"`
	// override the gateway's default recursive DNS servers, if DNS features are enabled.
	UpstreamDnsServers []string `pulumi:"upstreamDnsServers"`
	// `zone`) The zone in which the public gateway should be created.
	Zone *string `pulumi:"zone"`
}

type VpcPublicGatewayState struct {
	// The date and time of the creation of the public gateway.
	CreatedAt pulumi.StringPtrInput
	// attach an existing flexible IP to the gateway
	IpId pulumi.StringPtrInput
	// The name of the public gateway. If not provided it will be randomly generated.
	Name pulumi.StringPtrInput
	// The organization ID the public gateway is associated with.
	OrganizationId pulumi.StringPtrInput
	// `projectId`) The ID of the project the public gateway is associated with.
	ProjectId pulumi.StringPtrInput
	// The tags associated with the public gateway.
	Tags pulumi.StringArrayInput
	// The gateway type.
	Type pulumi.StringPtrInput
	// The date and time of the last update of the public gateway.
	UpdatedAt pulumi.StringPtrInput
	// override the gateway's default recursive DNS servers, if DNS features are enabled.
	UpstreamDnsServers pulumi.StringArrayInput
	// `zone`) The zone in which the public gateway should be created.
	Zone pulumi.StringPtrInput
}

func (VpcPublicGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcPublicGatewayState)(nil)).Elem()
}

type vpcPublicGatewayArgs struct {
	// attach an existing flexible IP to the gateway
	IpId *string `pulumi:"ipId"`
	// The name of the public gateway. If not provided it will be randomly generated.
	Name *string `pulumi:"name"`
	// `projectId`) The ID of the project the public gateway is associated with.
	ProjectId *string `pulumi:"projectId"`
	// The tags associated with the public gateway.
	Tags []string `pulumi:"tags"`
	// The gateway type.
	Type string `pulumi:"type"`
	// override the gateway's default recursive DNS servers, if DNS features are enabled.
	UpstreamDnsServers []string `pulumi:"upstreamDnsServers"`
	// `zone`) The zone in which the public gateway should be created.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a VpcPublicGateway resource.
type VpcPublicGatewayArgs struct {
	// attach an existing flexible IP to the gateway
	IpId pulumi.StringPtrInput
	// The name of the public gateway. If not provided it will be randomly generated.
	Name pulumi.StringPtrInput
	// `projectId`) The ID of the project the public gateway is associated with.
	ProjectId pulumi.StringPtrInput
	// The tags associated with the public gateway.
	Tags pulumi.StringArrayInput
	// The gateway type.
	Type pulumi.StringInput
	// override the gateway's default recursive DNS servers, if DNS features are enabled.
	UpstreamDnsServers pulumi.StringArrayInput
	// `zone`) The zone in which the public gateway should be created.
	Zone pulumi.StringPtrInput
}

func (VpcPublicGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcPublicGatewayArgs)(nil)).Elem()
}

type VpcPublicGatewayInput interface {
	pulumi.Input

	ToVpcPublicGatewayOutput() VpcPublicGatewayOutput
	ToVpcPublicGatewayOutputWithContext(ctx context.Context) VpcPublicGatewayOutput
}

func (*VpcPublicGateway) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcPublicGateway)(nil)).Elem()
}

func (i *VpcPublicGateway) ToVpcPublicGatewayOutput() VpcPublicGatewayOutput {
	return i.ToVpcPublicGatewayOutputWithContext(context.Background())
}

func (i *VpcPublicGateway) ToVpcPublicGatewayOutputWithContext(ctx context.Context) VpcPublicGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcPublicGatewayOutput)
}

type VpcPublicGatewayOutput struct{ *pulumi.OutputState }

func (VpcPublicGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcPublicGateway)(nil)).Elem()
}

func (o VpcPublicGatewayOutput) ToVpcPublicGatewayOutput() VpcPublicGatewayOutput {
	return o
}

func (o VpcPublicGatewayOutput) ToVpcPublicGatewayOutputWithContext(ctx context.Context) VpcPublicGatewayOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcPublicGatewayInput)(nil)).Elem(), &VpcPublicGateway{})
	pulumi.RegisterOutputType(VpcPublicGatewayOutput{})
}
