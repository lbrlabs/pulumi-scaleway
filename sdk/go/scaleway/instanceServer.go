// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type InstanceServer struct {
	pulumi.CustomResourceState

	// The additional volumes attached to the server
	AdditionalVolumeIds pulumi.StringArrayOutput `pulumi:"additionalVolumeIds"`
	// The boot type of the server
	BootType pulumi.StringPtrOutput `pulumi:"bootType"`
	// ID of the target bootscript (set boot_type to bootscript)
	BootscriptId pulumi.StringOutput `pulumi:"bootscriptId"`
	// The cloud init script associated with this server
	CloudInit pulumi.StringPtrOutput `pulumi:"cloudInit"`
	// Enable dynamic IP on the server
	EnableDynamicIp pulumi.BoolPtrOutput `pulumi:"enableDynamicIp"`
	// Determines if IPv6 is enabled for the server
	EnableIpv6 pulumi.BoolPtrOutput `pulumi:"enableIpv6"`
	// The UUID or the label of the base image used by the server
	Image pulumi.StringOutput `pulumi:"image"`
	// The ID of the reserved IP for the server
	IpId pulumi.StringPtrOutput `pulumi:"ipId"`
	// The default public IPv6 address routed to the server.
	Ipv6Address pulumi.StringOutput `pulumi:"ipv6Address"`
	// The IPv6 gateway address
	Ipv6Gateway pulumi.StringOutput `pulumi:"ipv6Gateway"`
	// The IPv6 prefix length routed to the server.
	Ipv6PrefixLength pulumi.IntOutput `pulumi:"ipv6PrefixLength"`
	// The name of the server
	Name pulumi.StringOutput `pulumi:"name"`
	// The organization_id you want to attach the resource to
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// The placement group the server is attached to
	PlacementGroupId pulumi.StringPtrOutput `pulumi:"placementGroupId"`
	// True when the placement group policy is respected
	PlacementGroupPolicyRespected pulumi.BoolOutput `pulumi:"placementGroupPolicyRespected"`
	// The Scaleway internal IP address of the server
	PrivateIp pulumi.StringOutput `pulumi:"privateIp"`
	// List of private network to connect with your instance
	PrivateNetworks InstanceServerPrivateNetworkArrayOutput `pulumi:"privateNetworks"`
	// The project_id you want to attach the resource to
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The public IPv4 address of the server
	PublicIp pulumi.StringOutput `pulumi:"publicIp"`
	// Root volume attached to the server on creation
	RootVolume InstanceServerRootVolumeOutput `pulumi:"rootVolume"`
	// The security group the server is attached to
	SecurityGroupId pulumi.StringOutput `pulumi:"securityGroupId"`
	// The state of the server should be: started, stopped, standby
	State pulumi.StringPtrOutput `pulumi:"state"`
	// The tags associated with the server
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The instance type of the server
	Type pulumi.StringOutput `pulumi:"type"`
	// The user data associated with the server
	UserData pulumi.StringMapOutput `pulumi:"userData"`
	// The zone you want to attach the resource to
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewInstanceServer registers a new resource with the given unique name, arguments, and options.
func NewInstanceServer(ctx *pulumi.Context,
	name string, args *InstanceServerArgs, opts ...pulumi.ResourceOption) (*InstanceServer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Image == nil {
		return nil, errors.New("invalid value for required argument 'Image'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource InstanceServer
	err := ctx.RegisterResource("scaleway:index/instanceServer:InstanceServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceServer gets an existing InstanceServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceServerState, opts ...pulumi.ResourceOption) (*InstanceServer, error) {
	var resource InstanceServer
	err := ctx.ReadResource("scaleway:index/instanceServer:InstanceServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceServer resources.
type instanceServerState struct {
	// The additional volumes attached to the server
	AdditionalVolumeIds []string `pulumi:"additionalVolumeIds"`
	// The boot type of the server
	BootType *string `pulumi:"bootType"`
	// ID of the target bootscript (set boot_type to bootscript)
	BootscriptId *string `pulumi:"bootscriptId"`
	// The cloud init script associated with this server
	CloudInit *string `pulumi:"cloudInit"`
	// Enable dynamic IP on the server
	EnableDynamicIp *bool `pulumi:"enableDynamicIp"`
	// Determines if IPv6 is enabled for the server
	EnableIpv6 *bool `pulumi:"enableIpv6"`
	// The UUID or the label of the base image used by the server
	Image *string `pulumi:"image"`
	// The ID of the reserved IP for the server
	IpId *string `pulumi:"ipId"`
	// The default public IPv6 address routed to the server.
	Ipv6Address *string `pulumi:"ipv6Address"`
	// The IPv6 gateway address
	Ipv6Gateway *string `pulumi:"ipv6Gateway"`
	// The IPv6 prefix length routed to the server.
	Ipv6PrefixLength *int `pulumi:"ipv6PrefixLength"`
	// The name of the server
	Name *string `pulumi:"name"`
	// The organization_id you want to attach the resource to
	OrganizationId *string `pulumi:"organizationId"`
	// The placement group the server is attached to
	PlacementGroupId *string `pulumi:"placementGroupId"`
	// True when the placement group policy is respected
	PlacementGroupPolicyRespected *bool `pulumi:"placementGroupPolicyRespected"`
	// The Scaleway internal IP address of the server
	PrivateIp *string `pulumi:"privateIp"`
	// List of private network to connect with your instance
	PrivateNetworks []InstanceServerPrivateNetwork `pulumi:"privateNetworks"`
	// The project_id you want to attach the resource to
	ProjectId *string `pulumi:"projectId"`
	// The public IPv4 address of the server
	PublicIp *string `pulumi:"publicIp"`
	// Root volume attached to the server on creation
	RootVolume *InstanceServerRootVolume `pulumi:"rootVolume"`
	// The security group the server is attached to
	SecurityGroupId *string `pulumi:"securityGroupId"`
	// The state of the server should be: started, stopped, standby
	State *string `pulumi:"state"`
	// The tags associated with the server
	Tags []string `pulumi:"tags"`
	// The instance type of the server
	Type *string `pulumi:"type"`
	// The user data associated with the server
	UserData map[string]string `pulumi:"userData"`
	// The zone you want to attach the resource to
	Zone *string `pulumi:"zone"`
}

type InstanceServerState struct {
	// The additional volumes attached to the server
	AdditionalVolumeIds pulumi.StringArrayInput
	// The boot type of the server
	BootType pulumi.StringPtrInput
	// ID of the target bootscript (set boot_type to bootscript)
	BootscriptId pulumi.StringPtrInput
	// The cloud init script associated with this server
	CloudInit pulumi.StringPtrInput
	// Enable dynamic IP on the server
	EnableDynamicIp pulumi.BoolPtrInput
	// Determines if IPv6 is enabled for the server
	EnableIpv6 pulumi.BoolPtrInput
	// The UUID or the label of the base image used by the server
	Image pulumi.StringPtrInput
	// The ID of the reserved IP for the server
	IpId pulumi.StringPtrInput
	// The default public IPv6 address routed to the server.
	Ipv6Address pulumi.StringPtrInput
	// The IPv6 gateway address
	Ipv6Gateway pulumi.StringPtrInput
	// The IPv6 prefix length routed to the server.
	Ipv6PrefixLength pulumi.IntPtrInput
	// The name of the server
	Name pulumi.StringPtrInput
	// The organization_id you want to attach the resource to
	OrganizationId pulumi.StringPtrInput
	// The placement group the server is attached to
	PlacementGroupId pulumi.StringPtrInput
	// True when the placement group policy is respected
	PlacementGroupPolicyRespected pulumi.BoolPtrInput
	// The Scaleway internal IP address of the server
	PrivateIp pulumi.StringPtrInput
	// List of private network to connect with your instance
	PrivateNetworks InstanceServerPrivateNetworkArrayInput
	// The project_id you want to attach the resource to
	ProjectId pulumi.StringPtrInput
	// The public IPv4 address of the server
	PublicIp pulumi.StringPtrInput
	// Root volume attached to the server on creation
	RootVolume InstanceServerRootVolumePtrInput
	// The security group the server is attached to
	SecurityGroupId pulumi.StringPtrInput
	// The state of the server should be: started, stopped, standby
	State pulumi.StringPtrInput
	// The tags associated with the server
	Tags pulumi.StringArrayInput
	// The instance type of the server
	Type pulumi.StringPtrInput
	// The user data associated with the server
	UserData pulumi.StringMapInput
	// The zone you want to attach the resource to
	Zone pulumi.StringPtrInput
}

func (InstanceServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceServerState)(nil)).Elem()
}

type instanceServerArgs struct {
	// The additional volumes attached to the server
	AdditionalVolumeIds []string `pulumi:"additionalVolumeIds"`
	// The boot type of the server
	BootType *string `pulumi:"bootType"`
	// ID of the target bootscript (set boot_type to bootscript)
	BootscriptId *string `pulumi:"bootscriptId"`
	// The cloud init script associated with this server
	CloudInit *string `pulumi:"cloudInit"`
	// Enable dynamic IP on the server
	EnableDynamicIp *bool `pulumi:"enableDynamicIp"`
	// Determines if IPv6 is enabled for the server
	EnableIpv6 *bool `pulumi:"enableIpv6"`
	// The UUID or the label of the base image used by the server
	Image string `pulumi:"image"`
	// The ID of the reserved IP for the server
	IpId *string `pulumi:"ipId"`
	// The name of the server
	Name *string `pulumi:"name"`
	// The placement group the server is attached to
	PlacementGroupId *string `pulumi:"placementGroupId"`
	// List of private network to connect with your instance
	PrivateNetworks []InstanceServerPrivateNetwork `pulumi:"privateNetworks"`
	// The project_id you want to attach the resource to
	ProjectId *string `pulumi:"projectId"`
	// Root volume attached to the server on creation
	RootVolume *InstanceServerRootVolume `pulumi:"rootVolume"`
	// The security group the server is attached to
	SecurityGroupId *string `pulumi:"securityGroupId"`
	// The state of the server should be: started, stopped, standby
	State *string `pulumi:"state"`
	// The tags associated with the server
	Tags []string `pulumi:"tags"`
	// The instance type of the server
	Type string `pulumi:"type"`
	// The user data associated with the server
	UserData map[string]string `pulumi:"userData"`
	// The zone you want to attach the resource to
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a InstanceServer resource.
type InstanceServerArgs struct {
	// The additional volumes attached to the server
	AdditionalVolumeIds pulumi.StringArrayInput
	// The boot type of the server
	BootType pulumi.StringPtrInput
	// ID of the target bootscript (set boot_type to bootscript)
	BootscriptId pulumi.StringPtrInput
	// The cloud init script associated with this server
	CloudInit pulumi.StringPtrInput
	// Enable dynamic IP on the server
	EnableDynamicIp pulumi.BoolPtrInput
	// Determines if IPv6 is enabled for the server
	EnableIpv6 pulumi.BoolPtrInput
	// The UUID or the label of the base image used by the server
	Image pulumi.StringInput
	// The ID of the reserved IP for the server
	IpId pulumi.StringPtrInput
	// The name of the server
	Name pulumi.StringPtrInput
	// The placement group the server is attached to
	PlacementGroupId pulumi.StringPtrInput
	// List of private network to connect with your instance
	PrivateNetworks InstanceServerPrivateNetworkArrayInput
	// The project_id you want to attach the resource to
	ProjectId pulumi.StringPtrInput
	// Root volume attached to the server on creation
	RootVolume InstanceServerRootVolumePtrInput
	// The security group the server is attached to
	SecurityGroupId pulumi.StringPtrInput
	// The state of the server should be: started, stopped, standby
	State pulumi.StringPtrInput
	// The tags associated with the server
	Tags pulumi.StringArrayInput
	// The instance type of the server
	Type pulumi.StringInput
	// The user data associated with the server
	UserData pulumi.StringMapInput
	// The zone you want to attach the resource to
	Zone pulumi.StringPtrInput
}

func (InstanceServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceServerArgs)(nil)).Elem()
}

type InstanceServerInput interface {
	pulumi.Input

	ToInstanceServerOutput() InstanceServerOutput
	ToInstanceServerOutputWithContext(ctx context.Context) InstanceServerOutput
}

func (*InstanceServer) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceServer)(nil)).Elem()
}

func (i *InstanceServer) ToInstanceServerOutput() InstanceServerOutput {
	return i.ToInstanceServerOutputWithContext(context.Background())
}

func (i *InstanceServer) ToInstanceServerOutputWithContext(ctx context.Context) InstanceServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceServerOutput)
}

type InstanceServerOutput struct{ *pulumi.OutputState }

func (InstanceServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceServer)(nil)).Elem()
}

func (o InstanceServerOutput) ToInstanceServerOutput() InstanceServerOutput {
	return o
}

func (o InstanceServerOutput) ToInstanceServerOutputWithContext(ctx context.Context) InstanceServerOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceServerInput)(nil)).Elem(), &InstanceServer{})
	pulumi.RegisterOutputType(InstanceServerOutput{})
}
