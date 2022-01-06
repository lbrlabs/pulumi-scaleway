// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates and manages Scaleway Load-Balancers IPs.
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
// 	"github.com/pulumi/pulumi-scaleway/sdk/go/scaleway"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := scaleway.NewLoadbalancerIP(ctx, "ip", &scaleway.LoadbalancerIPArgs{
// 			Reverse: pulumi.String("my-reverse.com"),
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
// IPs can be imported using the `{zone}/{id}`, e.g. bash
//
// ```sh
//  $ pulumi import scaleway:index/loadbalancerIP:LoadbalancerIP ip01 fr-par-1/11111111-1111-1111-1111-111111111111
// ```
type LoadbalancerIP struct {
	pulumi.CustomResourceState

	// The IP Address
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// The associated load-balance ID if any
	LbId pulumi.StringOutput `pulumi:"lbId"`
	// The organization_id you want to attach the resource to
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// The project_id you want to attach the resource to
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The region of the resource
	Region pulumi.StringOutput `pulumi:"region"`
	// The reverse domain associated with this IP.
	Reverse pulumi.StringOutput `pulumi:"reverse"`
	// The zone you want to attach the resource to
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewLoadbalancerIP registers a new resource with the given unique name, arguments, and options.
func NewLoadbalancerIP(ctx *pulumi.Context,
	name string, args *LoadbalancerIPArgs, opts ...pulumi.ResourceOption) (*LoadbalancerIP, error) {
	if args == nil {
		args = &LoadbalancerIPArgs{}
	}

	var resource LoadbalancerIP
	err := ctx.RegisterResource("scaleway:index/loadbalancerIP:LoadbalancerIP", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLoadbalancerIP gets an existing LoadbalancerIP resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLoadbalancerIP(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LoadbalancerIPState, opts ...pulumi.ResourceOption) (*LoadbalancerIP, error) {
	var resource LoadbalancerIP
	err := ctx.ReadResource("scaleway:index/loadbalancerIP:LoadbalancerIP", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LoadbalancerIP resources.
type loadbalancerIPState struct {
	// The IP Address
	IpAddress *string `pulumi:"ipAddress"`
	// The associated load-balance ID if any
	LbId *string `pulumi:"lbId"`
	// The organization_id you want to attach the resource to
	OrganizationId *string `pulumi:"organizationId"`
	// The project_id you want to attach the resource to
	ProjectId *string `pulumi:"projectId"`
	// The region of the resource
	Region *string `pulumi:"region"`
	// The reverse domain associated with this IP.
	Reverse *string `pulumi:"reverse"`
	// The zone you want to attach the resource to
	Zone *string `pulumi:"zone"`
}

type LoadbalancerIPState struct {
	// The IP Address
	IpAddress pulumi.StringPtrInput
	// The associated load-balance ID if any
	LbId pulumi.StringPtrInput
	// The organization_id you want to attach the resource to
	OrganizationId pulumi.StringPtrInput
	// The project_id you want to attach the resource to
	ProjectId pulumi.StringPtrInput
	// The region of the resource
	Region pulumi.StringPtrInput
	// The reverse domain associated with this IP.
	Reverse pulumi.StringPtrInput
	// The zone you want to attach the resource to
	Zone pulumi.StringPtrInput
}

func (LoadbalancerIPState) ElementType() reflect.Type {
	return reflect.TypeOf((*loadbalancerIPState)(nil)).Elem()
}

type loadbalancerIPArgs struct {
	// The project_id you want to attach the resource to
	ProjectId *string `pulumi:"projectId"`
	// The reverse domain associated with this IP.
	Reverse *string `pulumi:"reverse"`
	// The zone you want to attach the resource to
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a LoadbalancerIP resource.
type LoadbalancerIPArgs struct {
	// The project_id you want to attach the resource to
	ProjectId pulumi.StringPtrInput
	// The reverse domain associated with this IP.
	Reverse pulumi.StringPtrInput
	// The zone you want to attach the resource to
	Zone pulumi.StringPtrInput
}

func (LoadbalancerIPArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*loadbalancerIPArgs)(nil)).Elem()
}

type LoadbalancerIPInput interface {
	pulumi.Input

	ToLoadbalancerIPOutput() LoadbalancerIPOutput
	ToLoadbalancerIPOutputWithContext(ctx context.Context) LoadbalancerIPOutput
}

func (*LoadbalancerIP) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadbalancerIP)(nil)).Elem()
}

func (i *LoadbalancerIP) ToLoadbalancerIPOutput() LoadbalancerIPOutput {
	return i.ToLoadbalancerIPOutputWithContext(context.Background())
}

func (i *LoadbalancerIP) ToLoadbalancerIPOutputWithContext(ctx context.Context) LoadbalancerIPOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadbalancerIPOutput)
}

type LoadbalancerIPOutput struct{ *pulumi.OutputState }

func (LoadbalancerIPOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadbalancerIP)(nil)).Elem()
}

func (o LoadbalancerIPOutput) ToLoadbalancerIPOutput() LoadbalancerIPOutput {
	return o
}

func (o LoadbalancerIPOutput) ToLoadbalancerIPOutputWithContext(ctx context.Context) LoadbalancerIPOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LoadbalancerIPInput)(nil)).Elem(), &LoadbalancerIP{})
	pulumi.RegisterOutputType(LoadbalancerIPOutput{})
}