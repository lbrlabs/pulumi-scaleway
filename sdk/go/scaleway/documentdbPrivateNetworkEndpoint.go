// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"errors"
	"github.com/lbrlabs/pulumi-scaleway/sdk/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// ## Import
//
// Database Instance Endpoint can be imported using the `{region}/{endpoint_id}`, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:index/documentdbPrivateNetworkEndpoint:DocumentdbPrivateNetworkEndpoint end fr-par/11111111-1111-1111-1111-111111111111
//
// ```
type DocumentdbPrivateNetworkEndpoint struct {
	pulumi.CustomResourceState

	// Hostname of the endpoint.
	Hostname pulumi.StringOutput `pulumi:"hostname"`
	// UUID of the documentdb instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// IPv4 address on the network.
	Ip pulumi.StringOutput `pulumi:"ip"`
	// The IP network address within the private subnet. This must be an IPv4 address with a
	// CIDR notation. The IP network address within the private subnet is determined by the IP Address Management (IPAM)
	// service if not set.
	IpNet pulumi.StringOutput `pulumi:"ipNet"`
	// Name of the endpoint.
	Name pulumi.StringOutput `pulumi:"name"`
	// Port in the Private Network.
	Port pulumi.IntOutput `pulumi:"port"`
	// The ID of the private network.
	PrivateNetworkId pulumi.StringOutput `pulumi:"privateNetworkId"`
	// The region you want to attach the resource to
	Region pulumi.StringOutput `pulumi:"region"`
	// The zone you want to attach the resource to
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewDocumentdbPrivateNetworkEndpoint registers a new resource with the given unique name, arguments, and options.
func NewDocumentdbPrivateNetworkEndpoint(ctx *pulumi.Context,
	name string, args *DocumentdbPrivateNetworkEndpointArgs, opts ...pulumi.ResourceOption) (*DocumentdbPrivateNetworkEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.PrivateNetworkId == nil {
		return nil, errors.New("invalid value for required argument 'PrivateNetworkId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DocumentdbPrivateNetworkEndpoint
	err := ctx.RegisterResource("scaleway:index/documentdbPrivateNetworkEndpoint:DocumentdbPrivateNetworkEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDocumentdbPrivateNetworkEndpoint gets an existing DocumentdbPrivateNetworkEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDocumentdbPrivateNetworkEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DocumentdbPrivateNetworkEndpointState, opts ...pulumi.ResourceOption) (*DocumentdbPrivateNetworkEndpoint, error) {
	var resource DocumentdbPrivateNetworkEndpoint
	err := ctx.ReadResource("scaleway:index/documentdbPrivateNetworkEndpoint:DocumentdbPrivateNetworkEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DocumentdbPrivateNetworkEndpoint resources.
type documentdbPrivateNetworkEndpointState struct {
	// Hostname of the endpoint.
	Hostname *string `pulumi:"hostname"`
	// UUID of the documentdb instance.
	InstanceId *string `pulumi:"instanceId"`
	// IPv4 address on the network.
	Ip *string `pulumi:"ip"`
	// The IP network address within the private subnet. This must be an IPv4 address with a
	// CIDR notation. The IP network address within the private subnet is determined by the IP Address Management (IPAM)
	// service if not set.
	IpNet *string `pulumi:"ipNet"`
	// Name of the endpoint.
	Name *string `pulumi:"name"`
	// Port in the Private Network.
	Port *int `pulumi:"port"`
	// The ID of the private network.
	PrivateNetworkId *string `pulumi:"privateNetworkId"`
	// The region you want to attach the resource to
	Region *string `pulumi:"region"`
	// The zone you want to attach the resource to
	Zone *string `pulumi:"zone"`
}

type DocumentdbPrivateNetworkEndpointState struct {
	// Hostname of the endpoint.
	Hostname pulumi.StringPtrInput
	// UUID of the documentdb instance.
	InstanceId pulumi.StringPtrInput
	// IPv4 address on the network.
	Ip pulumi.StringPtrInput
	// The IP network address within the private subnet. This must be an IPv4 address with a
	// CIDR notation. The IP network address within the private subnet is determined by the IP Address Management (IPAM)
	// service if not set.
	IpNet pulumi.StringPtrInput
	// Name of the endpoint.
	Name pulumi.StringPtrInput
	// Port in the Private Network.
	Port pulumi.IntPtrInput
	// The ID of the private network.
	PrivateNetworkId pulumi.StringPtrInput
	// The region you want to attach the resource to
	Region pulumi.StringPtrInput
	// The zone you want to attach the resource to
	Zone pulumi.StringPtrInput
}

func (DocumentdbPrivateNetworkEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*documentdbPrivateNetworkEndpointState)(nil)).Elem()
}

type documentdbPrivateNetworkEndpointArgs struct {
	// UUID of the documentdb instance.
	InstanceId string `pulumi:"instanceId"`
	// The IP network address within the private subnet. This must be an IPv4 address with a
	// CIDR notation. The IP network address within the private subnet is determined by the IP Address Management (IPAM)
	// service if not set.
	IpNet *string `pulumi:"ipNet"`
	// Port in the Private Network.
	Port *int `pulumi:"port"`
	// The ID of the private network.
	PrivateNetworkId string `pulumi:"privateNetworkId"`
	// The region you want to attach the resource to
	Region *string `pulumi:"region"`
	// The zone you want to attach the resource to
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a DocumentdbPrivateNetworkEndpoint resource.
type DocumentdbPrivateNetworkEndpointArgs struct {
	// UUID of the documentdb instance.
	InstanceId pulumi.StringInput
	// The IP network address within the private subnet. This must be an IPv4 address with a
	// CIDR notation. The IP network address within the private subnet is determined by the IP Address Management (IPAM)
	// service if not set.
	IpNet pulumi.StringPtrInput
	// Port in the Private Network.
	Port pulumi.IntPtrInput
	// The ID of the private network.
	PrivateNetworkId pulumi.StringInput
	// The region you want to attach the resource to
	Region pulumi.StringPtrInput
	// The zone you want to attach the resource to
	Zone pulumi.StringPtrInput
}

func (DocumentdbPrivateNetworkEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*documentdbPrivateNetworkEndpointArgs)(nil)).Elem()
}

type DocumentdbPrivateNetworkEndpointInput interface {
	pulumi.Input

	ToDocumentdbPrivateNetworkEndpointOutput() DocumentdbPrivateNetworkEndpointOutput
	ToDocumentdbPrivateNetworkEndpointOutputWithContext(ctx context.Context) DocumentdbPrivateNetworkEndpointOutput
}

func (*DocumentdbPrivateNetworkEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**DocumentdbPrivateNetworkEndpoint)(nil)).Elem()
}

func (i *DocumentdbPrivateNetworkEndpoint) ToDocumentdbPrivateNetworkEndpointOutput() DocumentdbPrivateNetworkEndpointOutput {
	return i.ToDocumentdbPrivateNetworkEndpointOutputWithContext(context.Background())
}

func (i *DocumentdbPrivateNetworkEndpoint) ToDocumentdbPrivateNetworkEndpointOutputWithContext(ctx context.Context) DocumentdbPrivateNetworkEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentdbPrivateNetworkEndpointOutput)
}

func (i *DocumentdbPrivateNetworkEndpoint) ToOutput(ctx context.Context) pulumix.Output[*DocumentdbPrivateNetworkEndpoint] {
	return pulumix.Output[*DocumentdbPrivateNetworkEndpoint]{
		OutputState: i.ToDocumentdbPrivateNetworkEndpointOutputWithContext(ctx).OutputState,
	}
}

// DocumentdbPrivateNetworkEndpointArrayInput is an input type that accepts DocumentdbPrivateNetworkEndpointArray and DocumentdbPrivateNetworkEndpointArrayOutput values.
// You can construct a concrete instance of `DocumentdbPrivateNetworkEndpointArrayInput` via:
//
//	DocumentdbPrivateNetworkEndpointArray{ DocumentdbPrivateNetworkEndpointArgs{...} }
type DocumentdbPrivateNetworkEndpointArrayInput interface {
	pulumi.Input

	ToDocumentdbPrivateNetworkEndpointArrayOutput() DocumentdbPrivateNetworkEndpointArrayOutput
	ToDocumentdbPrivateNetworkEndpointArrayOutputWithContext(context.Context) DocumentdbPrivateNetworkEndpointArrayOutput
}

type DocumentdbPrivateNetworkEndpointArray []DocumentdbPrivateNetworkEndpointInput

func (DocumentdbPrivateNetworkEndpointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DocumentdbPrivateNetworkEndpoint)(nil)).Elem()
}

func (i DocumentdbPrivateNetworkEndpointArray) ToDocumentdbPrivateNetworkEndpointArrayOutput() DocumentdbPrivateNetworkEndpointArrayOutput {
	return i.ToDocumentdbPrivateNetworkEndpointArrayOutputWithContext(context.Background())
}

func (i DocumentdbPrivateNetworkEndpointArray) ToDocumentdbPrivateNetworkEndpointArrayOutputWithContext(ctx context.Context) DocumentdbPrivateNetworkEndpointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentdbPrivateNetworkEndpointArrayOutput)
}

func (i DocumentdbPrivateNetworkEndpointArray) ToOutput(ctx context.Context) pulumix.Output[[]*DocumentdbPrivateNetworkEndpoint] {
	return pulumix.Output[[]*DocumentdbPrivateNetworkEndpoint]{
		OutputState: i.ToDocumentdbPrivateNetworkEndpointArrayOutputWithContext(ctx).OutputState,
	}
}

// DocumentdbPrivateNetworkEndpointMapInput is an input type that accepts DocumentdbPrivateNetworkEndpointMap and DocumentdbPrivateNetworkEndpointMapOutput values.
// You can construct a concrete instance of `DocumentdbPrivateNetworkEndpointMapInput` via:
//
//	DocumentdbPrivateNetworkEndpointMap{ "key": DocumentdbPrivateNetworkEndpointArgs{...} }
type DocumentdbPrivateNetworkEndpointMapInput interface {
	pulumi.Input

	ToDocumentdbPrivateNetworkEndpointMapOutput() DocumentdbPrivateNetworkEndpointMapOutput
	ToDocumentdbPrivateNetworkEndpointMapOutputWithContext(context.Context) DocumentdbPrivateNetworkEndpointMapOutput
}

type DocumentdbPrivateNetworkEndpointMap map[string]DocumentdbPrivateNetworkEndpointInput

func (DocumentdbPrivateNetworkEndpointMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DocumentdbPrivateNetworkEndpoint)(nil)).Elem()
}

func (i DocumentdbPrivateNetworkEndpointMap) ToDocumentdbPrivateNetworkEndpointMapOutput() DocumentdbPrivateNetworkEndpointMapOutput {
	return i.ToDocumentdbPrivateNetworkEndpointMapOutputWithContext(context.Background())
}

func (i DocumentdbPrivateNetworkEndpointMap) ToDocumentdbPrivateNetworkEndpointMapOutputWithContext(ctx context.Context) DocumentdbPrivateNetworkEndpointMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentdbPrivateNetworkEndpointMapOutput)
}

func (i DocumentdbPrivateNetworkEndpointMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*DocumentdbPrivateNetworkEndpoint] {
	return pulumix.Output[map[string]*DocumentdbPrivateNetworkEndpoint]{
		OutputState: i.ToDocumentdbPrivateNetworkEndpointMapOutputWithContext(ctx).OutputState,
	}
}

type DocumentdbPrivateNetworkEndpointOutput struct{ *pulumi.OutputState }

func (DocumentdbPrivateNetworkEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DocumentdbPrivateNetworkEndpoint)(nil)).Elem()
}

func (o DocumentdbPrivateNetworkEndpointOutput) ToDocumentdbPrivateNetworkEndpointOutput() DocumentdbPrivateNetworkEndpointOutput {
	return o
}

func (o DocumentdbPrivateNetworkEndpointOutput) ToDocumentdbPrivateNetworkEndpointOutputWithContext(ctx context.Context) DocumentdbPrivateNetworkEndpointOutput {
	return o
}

func (o DocumentdbPrivateNetworkEndpointOutput) ToOutput(ctx context.Context) pulumix.Output[*DocumentdbPrivateNetworkEndpoint] {
	return pulumix.Output[*DocumentdbPrivateNetworkEndpoint]{
		OutputState: o.OutputState,
	}
}

// Hostname of the endpoint.
func (o DocumentdbPrivateNetworkEndpointOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentdbPrivateNetworkEndpoint) pulumi.StringOutput { return v.Hostname }).(pulumi.StringOutput)
}

// UUID of the documentdb instance.
func (o DocumentdbPrivateNetworkEndpointOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentdbPrivateNetworkEndpoint) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// IPv4 address on the network.
func (o DocumentdbPrivateNetworkEndpointOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentdbPrivateNetworkEndpoint) pulumi.StringOutput { return v.Ip }).(pulumi.StringOutput)
}

// The IP network address within the private subnet. This must be an IPv4 address with a
// CIDR notation. The IP network address within the private subnet is determined by the IP Address Management (IPAM)
// service if not set.
func (o DocumentdbPrivateNetworkEndpointOutput) IpNet() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentdbPrivateNetworkEndpoint) pulumi.StringOutput { return v.IpNet }).(pulumi.StringOutput)
}

// Name of the endpoint.
func (o DocumentdbPrivateNetworkEndpointOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentdbPrivateNetworkEndpoint) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Port in the Private Network.
func (o DocumentdbPrivateNetworkEndpointOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *DocumentdbPrivateNetworkEndpoint) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

// The ID of the private network.
func (o DocumentdbPrivateNetworkEndpointOutput) PrivateNetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentdbPrivateNetworkEndpoint) pulumi.StringOutput { return v.PrivateNetworkId }).(pulumi.StringOutput)
}

// The region you want to attach the resource to
func (o DocumentdbPrivateNetworkEndpointOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentdbPrivateNetworkEndpoint) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The zone you want to attach the resource to
func (o DocumentdbPrivateNetworkEndpointOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentdbPrivateNetworkEndpoint) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type DocumentdbPrivateNetworkEndpointArrayOutput struct{ *pulumi.OutputState }

func (DocumentdbPrivateNetworkEndpointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DocumentdbPrivateNetworkEndpoint)(nil)).Elem()
}

func (o DocumentdbPrivateNetworkEndpointArrayOutput) ToDocumentdbPrivateNetworkEndpointArrayOutput() DocumentdbPrivateNetworkEndpointArrayOutput {
	return o
}

func (o DocumentdbPrivateNetworkEndpointArrayOutput) ToDocumentdbPrivateNetworkEndpointArrayOutputWithContext(ctx context.Context) DocumentdbPrivateNetworkEndpointArrayOutput {
	return o
}

func (o DocumentdbPrivateNetworkEndpointArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*DocumentdbPrivateNetworkEndpoint] {
	return pulumix.Output[[]*DocumentdbPrivateNetworkEndpoint]{
		OutputState: o.OutputState,
	}
}

func (o DocumentdbPrivateNetworkEndpointArrayOutput) Index(i pulumi.IntInput) DocumentdbPrivateNetworkEndpointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DocumentdbPrivateNetworkEndpoint {
		return vs[0].([]*DocumentdbPrivateNetworkEndpoint)[vs[1].(int)]
	}).(DocumentdbPrivateNetworkEndpointOutput)
}

type DocumentdbPrivateNetworkEndpointMapOutput struct{ *pulumi.OutputState }

func (DocumentdbPrivateNetworkEndpointMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DocumentdbPrivateNetworkEndpoint)(nil)).Elem()
}

func (o DocumentdbPrivateNetworkEndpointMapOutput) ToDocumentdbPrivateNetworkEndpointMapOutput() DocumentdbPrivateNetworkEndpointMapOutput {
	return o
}

func (o DocumentdbPrivateNetworkEndpointMapOutput) ToDocumentdbPrivateNetworkEndpointMapOutputWithContext(ctx context.Context) DocumentdbPrivateNetworkEndpointMapOutput {
	return o
}

func (o DocumentdbPrivateNetworkEndpointMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*DocumentdbPrivateNetworkEndpoint] {
	return pulumix.Output[map[string]*DocumentdbPrivateNetworkEndpoint]{
		OutputState: o.OutputState,
	}
}

func (o DocumentdbPrivateNetworkEndpointMapOutput) MapIndex(k pulumi.StringInput) DocumentdbPrivateNetworkEndpointOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DocumentdbPrivateNetworkEndpoint {
		return vs[0].(map[string]*DocumentdbPrivateNetworkEndpoint)[vs[1].(string)]
	}).(DocumentdbPrivateNetworkEndpointOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DocumentdbPrivateNetworkEndpointInput)(nil)).Elem(), &DocumentdbPrivateNetworkEndpoint{})
	pulumi.RegisterInputType(reflect.TypeOf((*DocumentdbPrivateNetworkEndpointArrayInput)(nil)).Elem(), DocumentdbPrivateNetworkEndpointArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DocumentdbPrivateNetworkEndpointMapInput)(nil)).Elem(), DocumentdbPrivateNetworkEndpointMap{})
	pulumi.RegisterOutputType(DocumentdbPrivateNetworkEndpointOutput{})
	pulumi.RegisterOutputType(DocumentdbPrivateNetworkEndpointArrayOutput{})
	pulumi.RegisterOutputType(DocumentdbPrivateNetworkEndpointMapOutput{})
}
