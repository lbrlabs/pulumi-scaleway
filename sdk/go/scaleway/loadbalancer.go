// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates and manages Scaleway Load-Balancers.
// For more information, see [the documentation](https://developers.scaleway.com/en/products/lb/zoned_api).
//
// ## Examples
//
// ### Basic
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
//			main, err := scaleway.NewLoadbalancerIp(ctx, "main", &scaleway.LoadbalancerIpArgs{
//				Zone: pulumi.String("fr-par-1"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewLoadbalancer(ctx, "base", &scaleway.LoadbalancerArgs{
//				IpId: main.ID(),
//				Zone: main.Zone,
//				Type: pulumi.String("LB-S"),
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
// ### IP for Public Gateway
// resource "scaleway_vpc_public_gateway_ip" "main" {
// }
//
// ### Scaleway Private Network
// resource scaleway_vpc_private_network main {
// }
//
// ### VPC Public Gateway Network
//
//	resource "scaleway_vpc_public_gateway" "main" {
//	    name  = "tf-test-public-gw"
//	    type  = "VPC-GW-S"
//	    ip_id = scaleway_vpc_public_gateway_ip.main.id
//	}
//
// ### VPC Public Gateway Network DHCP config
//
//	resource "scaleway_vpc_public_gateway_dhcp" "main" {
//	    subnet = "10.0.0.0/24"
//	}
//
// ### VPC Gateway Network
//
//	resource "scaleway_vpc_gateway_network" "main" {
//	    gateway_id         = scaleway_vpc_public_gateway.main.id
//	    private_network_id = scaleway_vpc_private_network.main.id
//	    dhcp_id            = scaleway_vpc_public_gateway_dhcp.main.id
//	    cleanup_dhcp       = true
//	    enable_masquerade  = true
//	}
//
// ### Scaleway Instance
//
//	resource "scaleway_instance_server" "main" {
//	    name        = "Scaleway Terraform Provider"
//	    type        = "DEV1-S"
//	    image       = "debian_bullseye"
//	    enable_ipv6 = false
//
//	    private_network {
//	        pn_id = scaleway_vpc_private_network.main.id
//	    }
//	}
//
// ### IP for LB IP
// resource scaleway_lb_ip main {
// }
//
// ### Scaleway Private Network
//
//	resource scaleway_vpc_private_network "main" {
//	    name = "private network with static config"
//	}
//
// ## Migration
//
// In order to migrate to other types you can check the migration up or down via our CLI `scw lb lb-types list`.
// this change will not recreate your Load Balancer.
//
// Please check our [documentation](https://developers.scaleway.com/en/products/lb/zoned_api/#post-355592) for further details
//
// ## IP ID
//
// Since v1.15.0, `ipId` is a required field. This means that now a separate `LoadbalancerIp` is required.
// When importing, the IP needs to be imported as well as the LB.
// When upgrading to v1.15.0, you will need to create a new `LoadbalancerIp` resource and import it.
//
// For instance, if you had the following:
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
//			_, err := scaleway.NewLoadbalancer(ctx, "main", &scaleway.LoadbalancerArgs{
//				Type: pulumi.String("LB-S"),
//				Zone: pulumi.String("fr-par-1"),
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
// You will need to update it to:
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
//			mainLoadbalancerIp, err := scaleway.NewLoadbalancerIp(ctx, "mainLoadbalancerIp", nil)
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewLoadbalancer(ctx, "mainLoadbalancer", &scaleway.LoadbalancerArgs{
//				IpId:      mainLoadbalancerIp.ID(),
//				Zone:      pulumi.String("fr-par-1"),
//				Type:      pulumi.String("LB-S"),
//				ReleaseIp: pulumi.Bool(false),
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
// Load-Balancer can be imported using the `{zone}/{id}`, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:index/loadbalancer:Loadbalancer main fr-par-1/11111111-1111-1111-1111-111111111111
//
// ```
//
//	Be aware that you will also need to import the `scaleway_lb_ip` resource.
type Loadbalancer struct {
	pulumi.CustomResourceState

	// The description of the load-balancer.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The load-balance public IP Address
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// The ID of the associated LB IP. See below.
	IpId pulumi.StringOutput `pulumi:"ipId"`
	// The name of the load-balancer.
	Name pulumi.StringOutput `pulumi:"name"`
	// The organization ID the load-balancer is associated with.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// List of private network to connect with your load balancer
	PrivateNetworks LoadbalancerPrivateNetworkArrayOutput `pulumi:"privateNetworks"`
	// `projectId`) The ID of the project the load-balancer is associated with.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The region of the resource
	Region pulumi.StringOutput `pulumi:"region"`
	// The releaseIp allow release the ip address associated with the load-balancers.
	//
	// Deprecated: The resource ip will be destroyed by it's own resource. Please set this to `false`
	ReleaseIp pulumi.BoolPtrOutput `pulumi:"releaseIp"`
	// Enforces minimal SSL version (in SSL/TLS offloading context). Please check [possible values](https://developers.scaleway.com/en/products/lb/zoned_api/#ssl-compatibility-level-442f99).
	SslCompatibilityLevel pulumi.StringPtrOutput `pulumi:"sslCompatibilityLevel"`
	// The tags associated with the load-balancers.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The type of the load-balancer. Please check the migration section to upgrade the type
	Type pulumi.StringOutput `pulumi:"type"`
	// `zone`) The zone of the load-balancer.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewLoadbalancer registers a new resource with the given unique name, arguments, and options.
func NewLoadbalancer(ctx *pulumi.Context,
	name string, args *LoadbalancerArgs, opts ...pulumi.ResourceOption) (*Loadbalancer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IpId == nil {
		return nil, errors.New("invalid value for required argument 'IpId'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Loadbalancer
	err := ctx.RegisterResource("scaleway:index/loadbalancer:Loadbalancer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLoadbalancer gets an existing Loadbalancer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLoadbalancer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LoadbalancerState, opts ...pulumi.ResourceOption) (*Loadbalancer, error) {
	var resource Loadbalancer
	err := ctx.ReadResource("scaleway:index/loadbalancer:Loadbalancer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Loadbalancer resources.
type loadbalancerState struct {
	// The description of the load-balancer.
	Description *string `pulumi:"description"`
	// The load-balance public IP Address
	IpAddress *string `pulumi:"ipAddress"`
	// The ID of the associated LB IP. See below.
	IpId *string `pulumi:"ipId"`
	// The name of the load-balancer.
	Name *string `pulumi:"name"`
	// The organization ID the load-balancer is associated with.
	OrganizationId *string `pulumi:"organizationId"`
	// List of private network to connect with your load balancer
	PrivateNetworks []LoadbalancerPrivateNetwork `pulumi:"privateNetworks"`
	// `projectId`) The ID of the project the load-balancer is associated with.
	ProjectId *string `pulumi:"projectId"`
	// The region of the resource
	Region *string `pulumi:"region"`
	// The releaseIp allow release the ip address associated with the load-balancers.
	//
	// Deprecated: The resource ip will be destroyed by it's own resource. Please set this to `false`
	ReleaseIp *bool `pulumi:"releaseIp"`
	// Enforces minimal SSL version (in SSL/TLS offloading context). Please check [possible values](https://developers.scaleway.com/en/products/lb/zoned_api/#ssl-compatibility-level-442f99).
	SslCompatibilityLevel *string `pulumi:"sslCompatibilityLevel"`
	// The tags associated with the load-balancers.
	Tags []string `pulumi:"tags"`
	// The type of the load-balancer. Please check the migration section to upgrade the type
	Type *string `pulumi:"type"`
	// `zone`) The zone of the load-balancer.
	Zone *string `pulumi:"zone"`
}

type LoadbalancerState struct {
	// The description of the load-balancer.
	Description pulumi.StringPtrInput
	// The load-balance public IP Address
	IpAddress pulumi.StringPtrInput
	// The ID of the associated LB IP. See below.
	IpId pulumi.StringPtrInput
	// The name of the load-balancer.
	Name pulumi.StringPtrInput
	// The organization ID the load-balancer is associated with.
	OrganizationId pulumi.StringPtrInput
	// List of private network to connect with your load balancer
	PrivateNetworks LoadbalancerPrivateNetworkArrayInput
	// `projectId`) The ID of the project the load-balancer is associated with.
	ProjectId pulumi.StringPtrInput
	// The region of the resource
	Region pulumi.StringPtrInput
	// The releaseIp allow release the ip address associated with the load-balancers.
	//
	// Deprecated: The resource ip will be destroyed by it's own resource. Please set this to `false`
	ReleaseIp pulumi.BoolPtrInput
	// Enforces minimal SSL version (in SSL/TLS offloading context). Please check [possible values](https://developers.scaleway.com/en/products/lb/zoned_api/#ssl-compatibility-level-442f99).
	SslCompatibilityLevel pulumi.StringPtrInput
	// The tags associated with the load-balancers.
	Tags pulumi.StringArrayInput
	// The type of the load-balancer. Please check the migration section to upgrade the type
	Type pulumi.StringPtrInput
	// `zone`) The zone of the load-balancer.
	Zone pulumi.StringPtrInput
}

func (LoadbalancerState) ElementType() reflect.Type {
	return reflect.TypeOf((*loadbalancerState)(nil)).Elem()
}

type loadbalancerArgs struct {
	// The description of the load-balancer.
	Description *string `pulumi:"description"`
	// The ID of the associated LB IP. See below.
	IpId string `pulumi:"ipId"`
	// The name of the load-balancer.
	Name *string `pulumi:"name"`
	// List of private network to connect with your load balancer
	PrivateNetworks []LoadbalancerPrivateNetwork `pulumi:"privateNetworks"`
	// `projectId`) The ID of the project the load-balancer is associated with.
	ProjectId *string `pulumi:"projectId"`
	// The releaseIp allow release the ip address associated with the load-balancers.
	//
	// Deprecated: The resource ip will be destroyed by it's own resource. Please set this to `false`
	ReleaseIp *bool `pulumi:"releaseIp"`
	// Enforces minimal SSL version (in SSL/TLS offloading context). Please check [possible values](https://developers.scaleway.com/en/products/lb/zoned_api/#ssl-compatibility-level-442f99).
	SslCompatibilityLevel *string `pulumi:"sslCompatibilityLevel"`
	// The tags associated with the load-balancers.
	Tags []string `pulumi:"tags"`
	// The type of the load-balancer. Please check the migration section to upgrade the type
	Type string `pulumi:"type"`
	// `zone`) The zone of the load-balancer.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a Loadbalancer resource.
type LoadbalancerArgs struct {
	// The description of the load-balancer.
	Description pulumi.StringPtrInput
	// The ID of the associated LB IP. See below.
	IpId pulumi.StringInput
	// The name of the load-balancer.
	Name pulumi.StringPtrInput
	// List of private network to connect with your load balancer
	PrivateNetworks LoadbalancerPrivateNetworkArrayInput
	// `projectId`) The ID of the project the load-balancer is associated with.
	ProjectId pulumi.StringPtrInput
	// The releaseIp allow release the ip address associated with the load-balancers.
	//
	// Deprecated: The resource ip will be destroyed by it's own resource. Please set this to `false`
	ReleaseIp pulumi.BoolPtrInput
	// Enforces minimal SSL version (in SSL/TLS offloading context). Please check [possible values](https://developers.scaleway.com/en/products/lb/zoned_api/#ssl-compatibility-level-442f99).
	SslCompatibilityLevel pulumi.StringPtrInput
	// The tags associated with the load-balancers.
	Tags pulumi.StringArrayInput
	// The type of the load-balancer. Please check the migration section to upgrade the type
	Type pulumi.StringInput
	// `zone`) The zone of the load-balancer.
	Zone pulumi.StringPtrInput
}

func (LoadbalancerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*loadbalancerArgs)(nil)).Elem()
}

type LoadbalancerInput interface {
	pulumi.Input

	ToLoadbalancerOutput() LoadbalancerOutput
	ToLoadbalancerOutputWithContext(ctx context.Context) LoadbalancerOutput
}

func (*Loadbalancer) ElementType() reflect.Type {
	return reflect.TypeOf((**Loadbalancer)(nil)).Elem()
}

func (i *Loadbalancer) ToLoadbalancerOutput() LoadbalancerOutput {
	return i.ToLoadbalancerOutputWithContext(context.Background())
}

func (i *Loadbalancer) ToLoadbalancerOutputWithContext(ctx context.Context) LoadbalancerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadbalancerOutput)
}

// LoadbalancerArrayInput is an input type that accepts LoadbalancerArray and LoadbalancerArrayOutput values.
// You can construct a concrete instance of `LoadbalancerArrayInput` via:
//
//	LoadbalancerArray{ LoadbalancerArgs{...} }
type LoadbalancerArrayInput interface {
	pulumi.Input

	ToLoadbalancerArrayOutput() LoadbalancerArrayOutput
	ToLoadbalancerArrayOutputWithContext(context.Context) LoadbalancerArrayOutput
}

type LoadbalancerArray []LoadbalancerInput

func (LoadbalancerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Loadbalancer)(nil)).Elem()
}

func (i LoadbalancerArray) ToLoadbalancerArrayOutput() LoadbalancerArrayOutput {
	return i.ToLoadbalancerArrayOutputWithContext(context.Background())
}

func (i LoadbalancerArray) ToLoadbalancerArrayOutputWithContext(ctx context.Context) LoadbalancerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadbalancerArrayOutput)
}

// LoadbalancerMapInput is an input type that accepts LoadbalancerMap and LoadbalancerMapOutput values.
// You can construct a concrete instance of `LoadbalancerMapInput` via:
//
//	LoadbalancerMap{ "key": LoadbalancerArgs{...} }
type LoadbalancerMapInput interface {
	pulumi.Input

	ToLoadbalancerMapOutput() LoadbalancerMapOutput
	ToLoadbalancerMapOutputWithContext(context.Context) LoadbalancerMapOutput
}

type LoadbalancerMap map[string]LoadbalancerInput

func (LoadbalancerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Loadbalancer)(nil)).Elem()
}

func (i LoadbalancerMap) ToLoadbalancerMapOutput() LoadbalancerMapOutput {
	return i.ToLoadbalancerMapOutputWithContext(context.Background())
}

func (i LoadbalancerMap) ToLoadbalancerMapOutputWithContext(ctx context.Context) LoadbalancerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadbalancerMapOutput)
}

type LoadbalancerOutput struct{ *pulumi.OutputState }

func (LoadbalancerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Loadbalancer)(nil)).Elem()
}

func (o LoadbalancerOutput) ToLoadbalancerOutput() LoadbalancerOutput {
	return o
}

func (o LoadbalancerOutput) ToLoadbalancerOutputWithContext(ctx context.Context) LoadbalancerOutput {
	return o
}

// The description of the load-balancer.
func (o LoadbalancerOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The load-balance public IP Address
func (o LoadbalancerOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.StringOutput { return v.IpAddress }).(pulumi.StringOutput)
}

// The ID of the associated LB IP. See below.
func (o LoadbalancerOutput) IpId() pulumi.StringOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.StringOutput { return v.IpId }).(pulumi.StringOutput)
}

// The name of the load-balancer.
func (o LoadbalancerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The organization ID the load-balancer is associated with.
func (o LoadbalancerOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// List of private network to connect with your load balancer
func (o LoadbalancerOutput) PrivateNetworks() LoadbalancerPrivateNetworkArrayOutput {
	return o.ApplyT(func(v *Loadbalancer) LoadbalancerPrivateNetworkArrayOutput { return v.PrivateNetworks }).(LoadbalancerPrivateNetworkArrayOutput)
}

// `projectId`) The ID of the project the load-balancer is associated with.
func (o LoadbalancerOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The region of the resource
func (o LoadbalancerOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The releaseIp allow release the ip address associated with the load-balancers.
//
// Deprecated: The resource ip will be destroyed by it's own resource. Please set this to `false`
func (o LoadbalancerOutput) ReleaseIp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.BoolPtrOutput { return v.ReleaseIp }).(pulumi.BoolPtrOutput)
}

// Enforces minimal SSL version (in SSL/TLS offloading context). Please check [possible values](https://developers.scaleway.com/en/products/lb/zoned_api/#ssl-compatibility-level-442f99).
func (o LoadbalancerOutput) SslCompatibilityLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.StringPtrOutput { return v.SslCompatibilityLevel }).(pulumi.StringPtrOutput)
}

// The tags associated with the load-balancers.
func (o LoadbalancerOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// The type of the load-balancer. Please check the migration section to upgrade the type
func (o LoadbalancerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// `zone`) The zone of the load-balancer.
func (o LoadbalancerOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type LoadbalancerArrayOutput struct{ *pulumi.OutputState }

func (LoadbalancerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Loadbalancer)(nil)).Elem()
}

func (o LoadbalancerArrayOutput) ToLoadbalancerArrayOutput() LoadbalancerArrayOutput {
	return o
}

func (o LoadbalancerArrayOutput) ToLoadbalancerArrayOutputWithContext(ctx context.Context) LoadbalancerArrayOutput {
	return o
}

func (o LoadbalancerArrayOutput) Index(i pulumi.IntInput) LoadbalancerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Loadbalancer {
		return vs[0].([]*Loadbalancer)[vs[1].(int)]
	}).(LoadbalancerOutput)
}

type LoadbalancerMapOutput struct{ *pulumi.OutputState }

func (LoadbalancerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Loadbalancer)(nil)).Elem()
}

func (o LoadbalancerMapOutput) ToLoadbalancerMapOutput() LoadbalancerMapOutput {
	return o
}

func (o LoadbalancerMapOutput) ToLoadbalancerMapOutputWithContext(ctx context.Context) LoadbalancerMapOutput {
	return o
}

func (o LoadbalancerMapOutput) MapIndex(k pulumi.StringInput) LoadbalancerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Loadbalancer {
		return vs[0].(map[string]*Loadbalancer)[vs[1].(string)]
	}).(LoadbalancerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LoadbalancerInput)(nil)).Elem(), &Loadbalancer{})
	pulumi.RegisterInputType(reflect.TypeOf((*LoadbalancerArrayInput)(nil)).Elem(), LoadbalancerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LoadbalancerMapInput)(nil)).Elem(), LoadbalancerMap{})
	pulumi.RegisterOutputType(LoadbalancerOutput{})
	pulumi.RegisterOutputType(LoadbalancerArrayOutput{})
	pulumi.RegisterOutputType(LoadbalancerMapOutput{})
}
