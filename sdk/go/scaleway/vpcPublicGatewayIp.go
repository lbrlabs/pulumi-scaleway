// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/lbrlabs/pulumi-scaleway/sdk/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Creates and manages Scaleway VPC Public Gateway IP.
// For more information, see [the documentation](https://developers.scaleway.com/en/products/vpc-gw/api/v1/#ips-268151).
//
// ## Example
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
//			main, err := scaleway.NewVpcPublicGatewayIp(ctx, "main", &scaleway.VpcPublicGatewayIpArgs{
//				Reverse: pulumi.String("tf.example.com"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewDomainRecord(ctx, "tfA", &scaleway.DomainRecordArgs{
//				Data:     main.Address,
//				DnsZone:  pulumi.String("example.com"),
//				Priority: pulumi.Int(1),
//				Ttl:      pulumi.Int(3600),
//				Type:     pulumi.String("A"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Public gateway can be imported using the `{zone}/{id}`, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:index/vpcPublicGatewayIp:VpcPublicGatewayIp main fr-par-1/11111111-1111-1111-1111-111111111111
//
// ```
type VpcPublicGatewayIp struct {
	pulumi.CustomResourceState

	// The IP address itself.
	Address pulumi.StringOutput `pulumi:"address"`
	// The date and time of the creation of the public gateway ip.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The organization ID the public gateway ip is associated with.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// `projectId`) The ID of the project the public gateway ip is associated with.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The reverse domain name for the IP address
	Reverse pulumi.StringOutput `pulumi:"reverse"`
	// The tags associated with the public gateway IP.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The date and time of the last update of the public gateway ip.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// `zone`) The zone in which the public gateway ip should be created.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewVpcPublicGatewayIp registers a new resource with the given unique name, arguments, and options.
func NewVpcPublicGatewayIp(ctx *pulumi.Context,
	name string, args *VpcPublicGatewayIpArgs, opts ...pulumi.ResourceOption) (*VpcPublicGatewayIp, error) {
	if args == nil {
		args = &VpcPublicGatewayIpArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VpcPublicGatewayIp
	err := ctx.RegisterResource("scaleway:index/vpcPublicGatewayIp:VpcPublicGatewayIp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcPublicGatewayIp gets an existing VpcPublicGatewayIp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcPublicGatewayIp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcPublicGatewayIpState, opts ...pulumi.ResourceOption) (*VpcPublicGatewayIp, error) {
	var resource VpcPublicGatewayIp
	err := ctx.ReadResource("scaleway:index/vpcPublicGatewayIp:VpcPublicGatewayIp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcPublicGatewayIp resources.
type vpcPublicGatewayIpState struct {
	// The IP address itself.
	Address *string `pulumi:"address"`
	// The date and time of the creation of the public gateway ip.
	CreatedAt *string `pulumi:"createdAt"`
	// The organization ID the public gateway ip is associated with.
	OrganizationId *string `pulumi:"organizationId"`
	// `projectId`) The ID of the project the public gateway ip is associated with.
	ProjectId *string `pulumi:"projectId"`
	// The reverse domain name for the IP address
	Reverse *string `pulumi:"reverse"`
	// The tags associated with the public gateway IP.
	Tags []string `pulumi:"tags"`
	// The date and time of the last update of the public gateway ip.
	UpdatedAt *string `pulumi:"updatedAt"`
	// `zone`) The zone in which the public gateway ip should be created.
	Zone *string `pulumi:"zone"`
}

type VpcPublicGatewayIpState struct {
	// The IP address itself.
	Address pulumi.StringPtrInput
	// The date and time of the creation of the public gateway ip.
	CreatedAt pulumi.StringPtrInput
	// The organization ID the public gateway ip is associated with.
	OrganizationId pulumi.StringPtrInput
	// `projectId`) The ID of the project the public gateway ip is associated with.
	ProjectId pulumi.StringPtrInput
	// The reverse domain name for the IP address
	Reverse pulumi.StringPtrInput
	// The tags associated with the public gateway IP.
	Tags pulumi.StringArrayInput
	// The date and time of the last update of the public gateway ip.
	UpdatedAt pulumi.StringPtrInput
	// `zone`) The zone in which the public gateway ip should be created.
	Zone pulumi.StringPtrInput
}

func (VpcPublicGatewayIpState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcPublicGatewayIpState)(nil)).Elem()
}

type vpcPublicGatewayIpArgs struct {
	// `projectId`) The ID of the project the public gateway ip is associated with.
	ProjectId *string `pulumi:"projectId"`
	// The reverse domain name for the IP address
	Reverse *string `pulumi:"reverse"`
	// The tags associated with the public gateway IP.
	Tags []string `pulumi:"tags"`
	// `zone`) The zone in which the public gateway ip should be created.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a VpcPublicGatewayIp resource.
type VpcPublicGatewayIpArgs struct {
	// `projectId`) The ID of the project the public gateway ip is associated with.
	ProjectId pulumi.StringPtrInput
	// The reverse domain name for the IP address
	Reverse pulumi.StringPtrInput
	// The tags associated with the public gateway IP.
	Tags pulumi.StringArrayInput
	// `zone`) The zone in which the public gateway ip should be created.
	Zone pulumi.StringPtrInput
}

func (VpcPublicGatewayIpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcPublicGatewayIpArgs)(nil)).Elem()
}

type VpcPublicGatewayIpInput interface {
	pulumi.Input

	ToVpcPublicGatewayIpOutput() VpcPublicGatewayIpOutput
	ToVpcPublicGatewayIpOutputWithContext(ctx context.Context) VpcPublicGatewayIpOutput
}

func (*VpcPublicGatewayIp) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcPublicGatewayIp)(nil)).Elem()
}

func (i *VpcPublicGatewayIp) ToVpcPublicGatewayIpOutput() VpcPublicGatewayIpOutput {
	return i.ToVpcPublicGatewayIpOutputWithContext(context.Background())
}

func (i *VpcPublicGatewayIp) ToVpcPublicGatewayIpOutputWithContext(ctx context.Context) VpcPublicGatewayIpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcPublicGatewayIpOutput)
}

func (i *VpcPublicGatewayIp) ToOutput(ctx context.Context) pulumix.Output[*VpcPublicGatewayIp] {
	return pulumix.Output[*VpcPublicGatewayIp]{
		OutputState: i.ToVpcPublicGatewayIpOutputWithContext(ctx).OutputState,
	}
}

// VpcPublicGatewayIpArrayInput is an input type that accepts VpcPublicGatewayIpArray and VpcPublicGatewayIpArrayOutput values.
// You can construct a concrete instance of `VpcPublicGatewayIpArrayInput` via:
//
//	VpcPublicGatewayIpArray{ VpcPublicGatewayIpArgs{...} }
type VpcPublicGatewayIpArrayInput interface {
	pulumi.Input

	ToVpcPublicGatewayIpArrayOutput() VpcPublicGatewayIpArrayOutput
	ToVpcPublicGatewayIpArrayOutputWithContext(context.Context) VpcPublicGatewayIpArrayOutput
}

type VpcPublicGatewayIpArray []VpcPublicGatewayIpInput

func (VpcPublicGatewayIpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcPublicGatewayIp)(nil)).Elem()
}

func (i VpcPublicGatewayIpArray) ToVpcPublicGatewayIpArrayOutput() VpcPublicGatewayIpArrayOutput {
	return i.ToVpcPublicGatewayIpArrayOutputWithContext(context.Background())
}

func (i VpcPublicGatewayIpArray) ToVpcPublicGatewayIpArrayOutputWithContext(ctx context.Context) VpcPublicGatewayIpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcPublicGatewayIpArrayOutput)
}

func (i VpcPublicGatewayIpArray) ToOutput(ctx context.Context) pulumix.Output[[]*VpcPublicGatewayIp] {
	return pulumix.Output[[]*VpcPublicGatewayIp]{
		OutputState: i.ToVpcPublicGatewayIpArrayOutputWithContext(ctx).OutputState,
	}
}

// VpcPublicGatewayIpMapInput is an input type that accepts VpcPublicGatewayIpMap and VpcPublicGatewayIpMapOutput values.
// You can construct a concrete instance of `VpcPublicGatewayIpMapInput` via:
//
//	VpcPublicGatewayIpMap{ "key": VpcPublicGatewayIpArgs{...} }
type VpcPublicGatewayIpMapInput interface {
	pulumi.Input

	ToVpcPublicGatewayIpMapOutput() VpcPublicGatewayIpMapOutput
	ToVpcPublicGatewayIpMapOutputWithContext(context.Context) VpcPublicGatewayIpMapOutput
}

type VpcPublicGatewayIpMap map[string]VpcPublicGatewayIpInput

func (VpcPublicGatewayIpMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcPublicGatewayIp)(nil)).Elem()
}

func (i VpcPublicGatewayIpMap) ToVpcPublicGatewayIpMapOutput() VpcPublicGatewayIpMapOutput {
	return i.ToVpcPublicGatewayIpMapOutputWithContext(context.Background())
}

func (i VpcPublicGatewayIpMap) ToVpcPublicGatewayIpMapOutputWithContext(ctx context.Context) VpcPublicGatewayIpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcPublicGatewayIpMapOutput)
}

func (i VpcPublicGatewayIpMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*VpcPublicGatewayIp] {
	return pulumix.Output[map[string]*VpcPublicGatewayIp]{
		OutputState: i.ToVpcPublicGatewayIpMapOutputWithContext(ctx).OutputState,
	}
}

type VpcPublicGatewayIpOutput struct{ *pulumi.OutputState }

func (VpcPublicGatewayIpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcPublicGatewayIp)(nil)).Elem()
}

func (o VpcPublicGatewayIpOutput) ToVpcPublicGatewayIpOutput() VpcPublicGatewayIpOutput {
	return o
}

func (o VpcPublicGatewayIpOutput) ToVpcPublicGatewayIpOutputWithContext(ctx context.Context) VpcPublicGatewayIpOutput {
	return o
}

func (o VpcPublicGatewayIpOutput) ToOutput(ctx context.Context) pulumix.Output[*VpcPublicGatewayIp] {
	return pulumix.Output[*VpcPublicGatewayIp]{
		OutputState: o.OutputState,
	}
}

// The IP address itself.
func (o VpcPublicGatewayIpOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGatewayIp) pulumi.StringOutput { return v.Address }).(pulumi.StringOutput)
}

// The date and time of the creation of the public gateway ip.
func (o VpcPublicGatewayIpOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGatewayIp) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The organization ID the public gateway ip is associated with.
func (o VpcPublicGatewayIpOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGatewayIp) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// `projectId`) The ID of the project the public gateway ip is associated with.
func (o VpcPublicGatewayIpOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGatewayIp) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The reverse domain name for the IP address
func (o VpcPublicGatewayIpOutput) Reverse() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGatewayIp) pulumi.StringOutput { return v.Reverse }).(pulumi.StringOutput)
}

// The tags associated with the public gateway IP.
func (o VpcPublicGatewayIpOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VpcPublicGatewayIp) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// The date and time of the last update of the public gateway ip.
func (o VpcPublicGatewayIpOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGatewayIp) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// `zone`) The zone in which the public gateway ip should be created.
func (o VpcPublicGatewayIpOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGatewayIp) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type VpcPublicGatewayIpArrayOutput struct{ *pulumi.OutputState }

func (VpcPublicGatewayIpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcPublicGatewayIp)(nil)).Elem()
}

func (o VpcPublicGatewayIpArrayOutput) ToVpcPublicGatewayIpArrayOutput() VpcPublicGatewayIpArrayOutput {
	return o
}

func (o VpcPublicGatewayIpArrayOutput) ToVpcPublicGatewayIpArrayOutputWithContext(ctx context.Context) VpcPublicGatewayIpArrayOutput {
	return o
}

func (o VpcPublicGatewayIpArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*VpcPublicGatewayIp] {
	return pulumix.Output[[]*VpcPublicGatewayIp]{
		OutputState: o.OutputState,
	}
}

func (o VpcPublicGatewayIpArrayOutput) Index(i pulumi.IntInput) VpcPublicGatewayIpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpcPublicGatewayIp {
		return vs[0].([]*VpcPublicGatewayIp)[vs[1].(int)]
	}).(VpcPublicGatewayIpOutput)
}

type VpcPublicGatewayIpMapOutput struct{ *pulumi.OutputState }

func (VpcPublicGatewayIpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcPublicGatewayIp)(nil)).Elem()
}

func (o VpcPublicGatewayIpMapOutput) ToVpcPublicGatewayIpMapOutput() VpcPublicGatewayIpMapOutput {
	return o
}

func (o VpcPublicGatewayIpMapOutput) ToVpcPublicGatewayIpMapOutputWithContext(ctx context.Context) VpcPublicGatewayIpMapOutput {
	return o
}

func (o VpcPublicGatewayIpMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*VpcPublicGatewayIp] {
	return pulumix.Output[map[string]*VpcPublicGatewayIp]{
		OutputState: o.OutputState,
	}
}

func (o VpcPublicGatewayIpMapOutput) MapIndex(k pulumi.StringInput) VpcPublicGatewayIpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpcPublicGatewayIp {
		return vs[0].(map[string]*VpcPublicGatewayIp)[vs[1].(string)]
	}).(VpcPublicGatewayIpOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcPublicGatewayIpInput)(nil)).Elem(), &VpcPublicGatewayIp{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcPublicGatewayIpArrayInput)(nil)).Elem(), VpcPublicGatewayIpArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcPublicGatewayIpMapInput)(nil)).Elem(), VpcPublicGatewayIpMap{})
	pulumi.RegisterOutputType(VpcPublicGatewayIpOutput{})
	pulumi.RegisterOutputType(VpcPublicGatewayIpArrayOutput{})
	pulumi.RegisterOutputType(VpcPublicGatewayIpMapOutput{})
}
