// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VpcPublicGatewayDhcpReservation struct {
	pulumi.CustomResourceState

	// The configuration creation date.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The ID of the owning GatewayNetwork (UUID format).
	GatewayNetworkId pulumi.StringOutput `pulumi:"gatewayNetworkId"`
	// The Hostname of the client machine.
	Hostname pulumi.StringOutput `pulumi:"hostname"`
	// The IP address to give to the machine (IPv4 address).
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// The MAC address to give a static entry to.
	MacAddress pulumi.StringOutput `pulumi:"macAddress"`
	// The reservation type, either static (DHCP reservation) or dynamic (DHCP lease). Possible values are reservation and
	// lease
	Type pulumi.StringOutput `pulumi:"type"`
	// The configuration last modification date.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// The zone you want to attach the resource to
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewVpcPublicGatewayDhcpReservation registers a new resource with the given unique name, arguments, and options.
func NewVpcPublicGatewayDhcpReservation(ctx *pulumi.Context,
	name string, args *VpcPublicGatewayDhcpReservationArgs, opts ...pulumi.ResourceOption) (*VpcPublicGatewayDhcpReservation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GatewayNetworkId == nil {
		return nil, errors.New("invalid value for required argument 'GatewayNetworkId'")
	}
	if args.IpAddress == nil {
		return nil, errors.New("invalid value for required argument 'IpAddress'")
	}
	if args.MacAddress == nil {
		return nil, errors.New("invalid value for required argument 'MacAddress'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource VpcPublicGatewayDhcpReservation
	err := ctx.RegisterResource("scaleway:index/vpcPublicGatewayDhcpReservation:VpcPublicGatewayDhcpReservation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcPublicGatewayDhcpReservation gets an existing VpcPublicGatewayDhcpReservation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcPublicGatewayDhcpReservation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcPublicGatewayDhcpReservationState, opts ...pulumi.ResourceOption) (*VpcPublicGatewayDhcpReservation, error) {
	var resource VpcPublicGatewayDhcpReservation
	err := ctx.ReadResource("scaleway:index/vpcPublicGatewayDhcpReservation:VpcPublicGatewayDhcpReservation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcPublicGatewayDhcpReservation resources.
type vpcPublicGatewayDhcpReservationState struct {
	// The configuration creation date.
	CreatedAt *string `pulumi:"createdAt"`
	// The ID of the owning GatewayNetwork (UUID format).
	GatewayNetworkId *string `pulumi:"gatewayNetworkId"`
	// The Hostname of the client machine.
	Hostname *string `pulumi:"hostname"`
	// The IP address to give to the machine (IPv4 address).
	IpAddress *string `pulumi:"ipAddress"`
	// The MAC address to give a static entry to.
	MacAddress *string `pulumi:"macAddress"`
	// The reservation type, either static (DHCP reservation) or dynamic (DHCP lease). Possible values are reservation and
	// lease
	Type *string `pulumi:"type"`
	// The configuration last modification date.
	UpdatedAt *string `pulumi:"updatedAt"`
	// The zone you want to attach the resource to
	Zone *string `pulumi:"zone"`
}

type VpcPublicGatewayDhcpReservationState struct {
	// The configuration creation date.
	CreatedAt pulumi.StringPtrInput
	// The ID of the owning GatewayNetwork (UUID format).
	GatewayNetworkId pulumi.StringPtrInput
	// The Hostname of the client machine.
	Hostname pulumi.StringPtrInput
	// The IP address to give to the machine (IPv4 address).
	IpAddress pulumi.StringPtrInput
	// The MAC address to give a static entry to.
	MacAddress pulumi.StringPtrInput
	// The reservation type, either static (DHCP reservation) or dynamic (DHCP lease). Possible values are reservation and
	// lease
	Type pulumi.StringPtrInput
	// The configuration last modification date.
	UpdatedAt pulumi.StringPtrInput
	// The zone you want to attach the resource to
	Zone pulumi.StringPtrInput
}

func (VpcPublicGatewayDhcpReservationState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcPublicGatewayDhcpReservationState)(nil)).Elem()
}

type vpcPublicGatewayDhcpReservationArgs struct {
	// The ID of the owning GatewayNetwork (UUID format).
	GatewayNetworkId string `pulumi:"gatewayNetworkId"`
	// The IP address to give to the machine (IPv4 address).
	IpAddress string `pulumi:"ipAddress"`
	// The MAC address to give a static entry to.
	MacAddress string `pulumi:"macAddress"`
	// The zone you want to attach the resource to
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a VpcPublicGatewayDhcpReservation resource.
type VpcPublicGatewayDhcpReservationArgs struct {
	// The ID of the owning GatewayNetwork (UUID format).
	GatewayNetworkId pulumi.StringInput
	// The IP address to give to the machine (IPv4 address).
	IpAddress pulumi.StringInput
	// The MAC address to give a static entry to.
	MacAddress pulumi.StringInput
	// The zone you want to attach the resource to
	Zone pulumi.StringPtrInput
}

func (VpcPublicGatewayDhcpReservationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcPublicGatewayDhcpReservationArgs)(nil)).Elem()
}

type VpcPublicGatewayDhcpReservationInput interface {
	pulumi.Input

	ToVpcPublicGatewayDhcpReservationOutput() VpcPublicGatewayDhcpReservationOutput
	ToVpcPublicGatewayDhcpReservationOutputWithContext(ctx context.Context) VpcPublicGatewayDhcpReservationOutput
}

func (*VpcPublicGatewayDhcpReservation) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcPublicGatewayDhcpReservation)(nil)).Elem()
}

func (i *VpcPublicGatewayDhcpReservation) ToVpcPublicGatewayDhcpReservationOutput() VpcPublicGatewayDhcpReservationOutput {
	return i.ToVpcPublicGatewayDhcpReservationOutputWithContext(context.Background())
}

func (i *VpcPublicGatewayDhcpReservation) ToVpcPublicGatewayDhcpReservationOutputWithContext(ctx context.Context) VpcPublicGatewayDhcpReservationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcPublicGatewayDhcpReservationOutput)
}

type VpcPublicGatewayDhcpReservationOutput struct{ *pulumi.OutputState }

func (VpcPublicGatewayDhcpReservationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcPublicGatewayDhcpReservation)(nil)).Elem()
}

func (o VpcPublicGatewayDhcpReservationOutput) ToVpcPublicGatewayDhcpReservationOutput() VpcPublicGatewayDhcpReservationOutput {
	return o
}

func (o VpcPublicGatewayDhcpReservationOutput) ToVpcPublicGatewayDhcpReservationOutputWithContext(ctx context.Context) VpcPublicGatewayDhcpReservationOutput {
	return o
}

// The configuration creation date.
func (o VpcPublicGatewayDhcpReservationOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGatewayDhcpReservation) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The ID of the owning GatewayNetwork (UUID format).
func (o VpcPublicGatewayDhcpReservationOutput) GatewayNetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGatewayDhcpReservation) pulumi.StringOutput { return v.GatewayNetworkId }).(pulumi.StringOutput)
}

// The Hostname of the client machine.
func (o VpcPublicGatewayDhcpReservationOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGatewayDhcpReservation) pulumi.StringOutput { return v.Hostname }).(pulumi.StringOutput)
}

// The IP address to give to the machine (IPv4 address).
func (o VpcPublicGatewayDhcpReservationOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGatewayDhcpReservation) pulumi.StringOutput { return v.IpAddress }).(pulumi.StringOutput)
}

// The MAC address to give a static entry to.
func (o VpcPublicGatewayDhcpReservationOutput) MacAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGatewayDhcpReservation) pulumi.StringOutput { return v.MacAddress }).(pulumi.StringOutput)
}

// The reservation type, either static (DHCP reservation) or dynamic (DHCP lease). Possible values are reservation and
// lease
func (o VpcPublicGatewayDhcpReservationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGatewayDhcpReservation) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The configuration last modification date.
func (o VpcPublicGatewayDhcpReservationOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGatewayDhcpReservation) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// The zone you want to attach the resource to
func (o VpcPublicGatewayDhcpReservationOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGatewayDhcpReservation) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcPublicGatewayDhcpReservationInput)(nil)).Elem(), &VpcPublicGatewayDhcpReservation{})
	pulumi.RegisterOutputType(VpcPublicGatewayDhcpReservationOutput{})
}
