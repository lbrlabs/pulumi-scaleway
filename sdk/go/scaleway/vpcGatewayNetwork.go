// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VpcGatewayNetwork struct {
	pulumi.CustomResourceState

	// Remove DHCP config on this network on destroy
	CleanupDhcp pulumi.BoolPtrOutput `pulumi:"cleanupDhcp"`
	// The date and time of the creation of the gateway network
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The ID of the public gateway DHCP config
	DhcpId pulumi.StringPtrOutput `pulumi:"dhcpId"`
	// Enable DHCP config on this network
	EnableDhcp pulumi.BoolPtrOutput `pulumi:"enableDhcp"`
	// Enable masquerade on this network
	EnableMasquerade pulumi.BoolPtrOutput `pulumi:"enableMasquerade"`
	// The ID of the public gateway where connect to
	GatewayId pulumi.StringOutput `pulumi:"gatewayId"`
	// The mac address on this network
	MacAddress pulumi.StringOutput `pulumi:"macAddress"`
	// The ID of the private network where connect to
	PrivateNetworkId pulumi.StringOutput `pulumi:"privateNetworkId"`
	// The static IP address in CIDR on this network
	StaticAddress pulumi.StringPtrOutput `pulumi:"staticAddress"`
	// The date and time of the last update of the gateway network
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// The zone you want to attach the resource to
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewVpcGatewayNetwork registers a new resource with the given unique name, arguments, and options.
func NewVpcGatewayNetwork(ctx *pulumi.Context,
	name string, args *VpcGatewayNetworkArgs, opts ...pulumi.ResourceOption) (*VpcGatewayNetwork, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GatewayId == nil {
		return nil, errors.New("invalid value for required argument 'GatewayId'")
	}
	if args.PrivateNetworkId == nil {
		return nil, errors.New("invalid value for required argument 'PrivateNetworkId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource VpcGatewayNetwork
	err := ctx.RegisterResource("scaleway:index/vpcGatewayNetwork:VpcGatewayNetwork", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcGatewayNetwork gets an existing VpcGatewayNetwork resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcGatewayNetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcGatewayNetworkState, opts ...pulumi.ResourceOption) (*VpcGatewayNetwork, error) {
	var resource VpcGatewayNetwork
	err := ctx.ReadResource("scaleway:index/vpcGatewayNetwork:VpcGatewayNetwork", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcGatewayNetwork resources.
type vpcGatewayNetworkState struct {
	// Remove DHCP config on this network on destroy
	CleanupDhcp *bool `pulumi:"cleanupDhcp"`
	// The date and time of the creation of the gateway network
	CreatedAt *string `pulumi:"createdAt"`
	// The ID of the public gateway DHCP config
	DhcpId *string `pulumi:"dhcpId"`
	// Enable DHCP config on this network
	EnableDhcp *bool `pulumi:"enableDhcp"`
	// Enable masquerade on this network
	EnableMasquerade *bool `pulumi:"enableMasquerade"`
	// The ID of the public gateway where connect to
	GatewayId *string `pulumi:"gatewayId"`
	// The mac address on this network
	MacAddress *string `pulumi:"macAddress"`
	// The ID of the private network where connect to
	PrivateNetworkId *string `pulumi:"privateNetworkId"`
	// The static IP address in CIDR on this network
	StaticAddress *string `pulumi:"staticAddress"`
	// The date and time of the last update of the gateway network
	UpdatedAt *string `pulumi:"updatedAt"`
	// The zone you want to attach the resource to
	Zone *string `pulumi:"zone"`
}

type VpcGatewayNetworkState struct {
	// Remove DHCP config on this network on destroy
	CleanupDhcp pulumi.BoolPtrInput
	// The date and time of the creation of the gateway network
	CreatedAt pulumi.StringPtrInput
	// The ID of the public gateway DHCP config
	DhcpId pulumi.StringPtrInput
	// Enable DHCP config on this network
	EnableDhcp pulumi.BoolPtrInput
	// Enable masquerade on this network
	EnableMasquerade pulumi.BoolPtrInput
	// The ID of the public gateway where connect to
	GatewayId pulumi.StringPtrInput
	// The mac address on this network
	MacAddress pulumi.StringPtrInput
	// The ID of the private network where connect to
	PrivateNetworkId pulumi.StringPtrInput
	// The static IP address in CIDR on this network
	StaticAddress pulumi.StringPtrInput
	// The date and time of the last update of the gateway network
	UpdatedAt pulumi.StringPtrInput
	// The zone you want to attach the resource to
	Zone pulumi.StringPtrInput
}

func (VpcGatewayNetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcGatewayNetworkState)(nil)).Elem()
}

type vpcGatewayNetworkArgs struct {
	// Remove DHCP config on this network on destroy
	CleanupDhcp *bool `pulumi:"cleanupDhcp"`
	// The ID of the public gateway DHCP config
	DhcpId *string `pulumi:"dhcpId"`
	// Enable DHCP config on this network
	EnableDhcp *bool `pulumi:"enableDhcp"`
	// Enable masquerade on this network
	EnableMasquerade *bool `pulumi:"enableMasquerade"`
	// The ID of the public gateway where connect to
	GatewayId string `pulumi:"gatewayId"`
	// The ID of the private network where connect to
	PrivateNetworkId string `pulumi:"privateNetworkId"`
	// The static IP address in CIDR on this network
	StaticAddress *string `pulumi:"staticAddress"`
	// The zone you want to attach the resource to
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a VpcGatewayNetwork resource.
type VpcGatewayNetworkArgs struct {
	// Remove DHCP config on this network on destroy
	CleanupDhcp pulumi.BoolPtrInput
	// The ID of the public gateway DHCP config
	DhcpId pulumi.StringPtrInput
	// Enable DHCP config on this network
	EnableDhcp pulumi.BoolPtrInput
	// Enable masquerade on this network
	EnableMasquerade pulumi.BoolPtrInput
	// The ID of the public gateway where connect to
	GatewayId pulumi.StringInput
	// The ID of the private network where connect to
	PrivateNetworkId pulumi.StringInput
	// The static IP address in CIDR on this network
	StaticAddress pulumi.StringPtrInput
	// The zone you want to attach the resource to
	Zone pulumi.StringPtrInput
}

func (VpcGatewayNetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcGatewayNetworkArgs)(nil)).Elem()
}

type VpcGatewayNetworkInput interface {
	pulumi.Input

	ToVpcGatewayNetworkOutput() VpcGatewayNetworkOutput
	ToVpcGatewayNetworkOutputWithContext(ctx context.Context) VpcGatewayNetworkOutput
}

func (*VpcGatewayNetwork) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcGatewayNetwork)(nil)).Elem()
}

func (i *VpcGatewayNetwork) ToVpcGatewayNetworkOutput() VpcGatewayNetworkOutput {
	return i.ToVpcGatewayNetworkOutputWithContext(context.Background())
}

func (i *VpcGatewayNetwork) ToVpcGatewayNetworkOutputWithContext(ctx context.Context) VpcGatewayNetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcGatewayNetworkOutput)
}

type VpcGatewayNetworkOutput struct{ *pulumi.OutputState }

func (VpcGatewayNetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcGatewayNetwork)(nil)).Elem()
}

func (o VpcGatewayNetworkOutput) ToVpcGatewayNetworkOutput() VpcGatewayNetworkOutput {
	return o
}

func (o VpcGatewayNetworkOutput) ToVpcGatewayNetworkOutputWithContext(ctx context.Context) VpcGatewayNetworkOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcGatewayNetworkInput)(nil)).Elem(), &VpcGatewayNetwork{})
	pulumi.RegisterOutputType(VpcGatewayNetworkOutput{})
}
