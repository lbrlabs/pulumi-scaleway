// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Creates and manages Scaleway Compute Baremetal servers. For more information, see [the documentation](https://developers.scaleway.com/en/products/baremetal/api).
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
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			main, err := scaleway.LookupAccountSshKey(ctx, &scaleway.LookupAccountSshKeyArgs{
//				Name: pulumi.StringRef("main"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewBaremetalServer(ctx, "base", &scaleway.BaremetalServerArgs{
//				Zone:  pulumi.String("fr-par-2"),
//				Offer: pulumi.String("GP-BM1-S"),
//				Os:    pulumi.String("d17d6872-0412-45d9-a198-af82c34d3c5c"),
//				SshKeyIds: pulumi.StringArray{
//					*pulumi.String(main.Id),
//				},
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
// ### Without install config
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
//			myOffer, err := scaleway.GetBaremetalOffer(ctx, &scaleway.GetBaremetalOfferArgs{
//				Zone: pulumi.StringRef("fr-par-2"),
//				Name: pulumi.StringRef("EM-B112X-SSD"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewBaremetalServer(ctx, "base", &scaleway.BaremetalServerArgs{
//				Zone:                   pulumi.String("fr-par-2"),
//				Offer:                  *pulumi.String(myOffer.OfferId),
//				InstallConfigAfterward: pulumi.Bool(true),
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
// Baremetal servers can be imported using the `{zone}/{id}`, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:index/baremetalServer:BaremetalServer web fr-par-2/11111111-1111-1111-1111-111111111111
//
// ```
type BaremetalServer struct {
	pulumi.CustomResourceState

	// A description for the server.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The domain of the server.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// The hostname of the server.
	Hostname pulumi.StringPtrOutput `pulumi:"hostname"`
	// If True, this boolean allows to create a server without the install config if you want to provide it later.
	InstallConfigAfterward pulumi.BoolPtrOutput `pulumi:"installConfigAfterward"`
	// (List of) The IPs of the server.
	Ips BaremetalServerIpArrayOutput `pulumi:"ips"`
	// (List of) The IPv4 addresses of the server.
	Ipv4s BaremetalServerIpv4ArrayOutput `pulumi:"ipv4s"`
	// (List of) The IPv6 addresses of the server.
	Ipv6s BaremetalServerIpv6ArrayOutput `pulumi:"ipv6s"`
	// The name of the server.
	Name pulumi.StringOutput `pulumi:"name"`
	// The offer name or UUID of the baremetal server.
	// Use [this endpoint](https://developers.scaleway.com/en/products/baremetal/api/#get-334154) to find the right offer.
	//
	// > **Important:** Updates to `offer` will recreate the server.
	Offer pulumi.StringOutput `pulumi:"offer"`
	// The ID of the offer.
	OfferId pulumi.StringOutput `pulumi:"offerId"`
	// The name of the offer.
	OfferName pulumi.StringOutput `pulumi:"offerName"`
	// The options to enable on the server.
	// > The `options` block supports:
	Options BaremetalServerOptionArrayOutput `pulumi:"options"`
	// The organization ID the server is associated with.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// The UUID of the os to install on the server.
	// Use [this endpoint](https://developers.scaleway.com/en/products/baremetal/api/#get-87598a) to find the right OS ID.
	// > **Important:** Updates to `os` will reinstall the server.
	Os pulumi.StringPtrOutput `pulumi:"os"`
	// The name of the os.
	OsName pulumi.StringOutput `pulumi:"osName"`
	// Password used for the installation. May be required depending on used os.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// The private networks to attach to the server. For more information, see [the documentation](https://www.scaleway.com/en/docs/compute/elastic-metal/how-to/use-private-networks/)
	PrivateNetworks BaremetalServerPrivateNetworkArrayOutput `pulumi:"privateNetworks"`
	// `projectId`) The ID of the project the server is associated with.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// If True, this boolean allows to reinstall the server on install config changes.
	// > **Important:** Updates to `sshKeyIds`, `user`, `password`, `serviceUser` or `servicePassword` will not take effect on the server, it requires to reinstall it. To do so please set 'reinstall_on_config_changes' argument to true.
	ReinstallOnConfigChanges pulumi.BoolPtrOutput `pulumi:"reinstallOnConfigChanges"`
	// Password used for the service to install. May be required depending on used os.
	ServicePassword pulumi.StringPtrOutput `pulumi:"servicePassword"`
	// User used for the service to install.
	ServiceUser pulumi.StringOutput `pulumi:"serviceUser"`
	// List of SSH keys allowed to connect to the server.
	SshKeyIds pulumi.StringArrayOutput `pulumi:"sshKeyIds"`
	// The tags associated with the server.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// User used for the installation.
	User pulumi.StringOutput `pulumi:"user"`
	// `zone`) The zone in which the server should be created.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewBaremetalServer registers a new resource with the given unique name, arguments, and options.
func NewBaremetalServer(ctx *pulumi.Context,
	name string, args *BaremetalServerArgs, opts ...pulumi.ResourceOption) (*BaremetalServer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Offer == nil {
		return nil, errors.New("invalid value for required argument 'Offer'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	if args.ServicePassword != nil {
		args.ServicePassword = pulumi.ToSecret(args.ServicePassword).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
		"servicePassword",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BaremetalServer
	err := ctx.RegisterResource("scaleway:index/baremetalServer:BaremetalServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBaremetalServer gets an existing BaremetalServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBaremetalServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BaremetalServerState, opts ...pulumi.ResourceOption) (*BaremetalServer, error) {
	var resource BaremetalServer
	err := ctx.ReadResource("scaleway:index/baremetalServer:BaremetalServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BaremetalServer resources.
type baremetalServerState struct {
	// A description for the server.
	Description *string `pulumi:"description"`
	// The domain of the server.
	Domain *string `pulumi:"domain"`
	// The hostname of the server.
	Hostname *string `pulumi:"hostname"`
	// If True, this boolean allows to create a server without the install config if you want to provide it later.
	InstallConfigAfterward *bool `pulumi:"installConfigAfterward"`
	// (List of) The IPs of the server.
	Ips []BaremetalServerIp `pulumi:"ips"`
	// (List of) The IPv4 addresses of the server.
	Ipv4s []BaremetalServerIpv4 `pulumi:"ipv4s"`
	// (List of) The IPv6 addresses of the server.
	Ipv6s []BaremetalServerIpv6 `pulumi:"ipv6s"`
	// The name of the server.
	Name *string `pulumi:"name"`
	// The offer name or UUID of the baremetal server.
	// Use [this endpoint](https://developers.scaleway.com/en/products/baremetal/api/#get-334154) to find the right offer.
	//
	// > **Important:** Updates to `offer` will recreate the server.
	Offer *string `pulumi:"offer"`
	// The ID of the offer.
	OfferId *string `pulumi:"offerId"`
	// The name of the offer.
	OfferName *string `pulumi:"offerName"`
	// The options to enable on the server.
	// > The `options` block supports:
	Options []BaremetalServerOption `pulumi:"options"`
	// The organization ID the server is associated with.
	OrganizationId *string `pulumi:"organizationId"`
	// The UUID of the os to install on the server.
	// Use [this endpoint](https://developers.scaleway.com/en/products/baremetal/api/#get-87598a) to find the right OS ID.
	// > **Important:** Updates to `os` will reinstall the server.
	Os *string `pulumi:"os"`
	// The name of the os.
	OsName *string `pulumi:"osName"`
	// Password used for the installation. May be required depending on used os.
	Password *string `pulumi:"password"`
	// The private networks to attach to the server. For more information, see [the documentation](https://www.scaleway.com/en/docs/compute/elastic-metal/how-to/use-private-networks/)
	PrivateNetworks []BaremetalServerPrivateNetwork `pulumi:"privateNetworks"`
	// `projectId`) The ID of the project the server is associated with.
	ProjectId *string `pulumi:"projectId"`
	// If True, this boolean allows to reinstall the server on install config changes.
	// > **Important:** Updates to `sshKeyIds`, `user`, `password`, `serviceUser` or `servicePassword` will not take effect on the server, it requires to reinstall it. To do so please set 'reinstall_on_config_changes' argument to true.
	ReinstallOnConfigChanges *bool `pulumi:"reinstallOnConfigChanges"`
	// Password used for the service to install. May be required depending on used os.
	ServicePassword *string `pulumi:"servicePassword"`
	// User used for the service to install.
	ServiceUser *string `pulumi:"serviceUser"`
	// List of SSH keys allowed to connect to the server.
	SshKeyIds []string `pulumi:"sshKeyIds"`
	// The tags associated with the server.
	Tags []string `pulumi:"tags"`
	// User used for the installation.
	User *string `pulumi:"user"`
	// `zone`) The zone in which the server should be created.
	Zone *string `pulumi:"zone"`
}

type BaremetalServerState struct {
	// A description for the server.
	Description pulumi.StringPtrInput
	// The domain of the server.
	Domain pulumi.StringPtrInput
	// The hostname of the server.
	Hostname pulumi.StringPtrInput
	// If True, this boolean allows to create a server without the install config if you want to provide it later.
	InstallConfigAfterward pulumi.BoolPtrInput
	// (List of) The IPs of the server.
	Ips BaremetalServerIpArrayInput
	// (List of) The IPv4 addresses of the server.
	Ipv4s BaremetalServerIpv4ArrayInput
	// (List of) The IPv6 addresses of the server.
	Ipv6s BaremetalServerIpv6ArrayInput
	// The name of the server.
	Name pulumi.StringPtrInput
	// The offer name or UUID of the baremetal server.
	// Use [this endpoint](https://developers.scaleway.com/en/products/baremetal/api/#get-334154) to find the right offer.
	//
	// > **Important:** Updates to `offer` will recreate the server.
	Offer pulumi.StringPtrInput
	// The ID of the offer.
	OfferId pulumi.StringPtrInput
	// The name of the offer.
	OfferName pulumi.StringPtrInput
	// The options to enable on the server.
	// > The `options` block supports:
	Options BaremetalServerOptionArrayInput
	// The organization ID the server is associated with.
	OrganizationId pulumi.StringPtrInput
	// The UUID of the os to install on the server.
	// Use [this endpoint](https://developers.scaleway.com/en/products/baremetal/api/#get-87598a) to find the right OS ID.
	// > **Important:** Updates to `os` will reinstall the server.
	Os pulumi.StringPtrInput
	// The name of the os.
	OsName pulumi.StringPtrInput
	// Password used for the installation. May be required depending on used os.
	Password pulumi.StringPtrInput
	// The private networks to attach to the server. For more information, see [the documentation](https://www.scaleway.com/en/docs/compute/elastic-metal/how-to/use-private-networks/)
	PrivateNetworks BaremetalServerPrivateNetworkArrayInput
	// `projectId`) The ID of the project the server is associated with.
	ProjectId pulumi.StringPtrInput
	// If True, this boolean allows to reinstall the server on install config changes.
	// > **Important:** Updates to `sshKeyIds`, `user`, `password`, `serviceUser` or `servicePassword` will not take effect on the server, it requires to reinstall it. To do so please set 'reinstall_on_config_changes' argument to true.
	ReinstallOnConfigChanges pulumi.BoolPtrInput
	// Password used for the service to install. May be required depending on used os.
	ServicePassword pulumi.StringPtrInput
	// User used for the service to install.
	ServiceUser pulumi.StringPtrInput
	// List of SSH keys allowed to connect to the server.
	SshKeyIds pulumi.StringArrayInput
	// The tags associated with the server.
	Tags pulumi.StringArrayInput
	// User used for the installation.
	User pulumi.StringPtrInput
	// `zone`) The zone in which the server should be created.
	Zone pulumi.StringPtrInput
}

func (BaremetalServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*baremetalServerState)(nil)).Elem()
}

type baremetalServerArgs struct {
	// A description for the server.
	Description *string `pulumi:"description"`
	// The hostname of the server.
	Hostname *string `pulumi:"hostname"`
	// If True, this boolean allows to create a server without the install config if you want to provide it later.
	InstallConfigAfterward *bool `pulumi:"installConfigAfterward"`
	// The name of the server.
	Name *string `pulumi:"name"`
	// The offer name or UUID of the baremetal server.
	// Use [this endpoint](https://developers.scaleway.com/en/products/baremetal/api/#get-334154) to find the right offer.
	//
	// > **Important:** Updates to `offer` will recreate the server.
	Offer string `pulumi:"offer"`
	// The options to enable on the server.
	// > The `options` block supports:
	Options []BaremetalServerOption `pulumi:"options"`
	// The UUID of the os to install on the server.
	// Use [this endpoint](https://developers.scaleway.com/en/products/baremetal/api/#get-87598a) to find the right OS ID.
	// > **Important:** Updates to `os` will reinstall the server.
	Os *string `pulumi:"os"`
	// Password used for the installation. May be required depending on used os.
	Password *string `pulumi:"password"`
	// The private networks to attach to the server. For more information, see [the documentation](https://www.scaleway.com/en/docs/compute/elastic-metal/how-to/use-private-networks/)
	PrivateNetworks []BaremetalServerPrivateNetwork `pulumi:"privateNetworks"`
	// `projectId`) The ID of the project the server is associated with.
	ProjectId *string `pulumi:"projectId"`
	// If True, this boolean allows to reinstall the server on install config changes.
	// > **Important:** Updates to `sshKeyIds`, `user`, `password`, `serviceUser` or `servicePassword` will not take effect on the server, it requires to reinstall it. To do so please set 'reinstall_on_config_changes' argument to true.
	ReinstallOnConfigChanges *bool `pulumi:"reinstallOnConfigChanges"`
	// Password used for the service to install. May be required depending on used os.
	ServicePassword *string `pulumi:"servicePassword"`
	// User used for the service to install.
	ServiceUser *string `pulumi:"serviceUser"`
	// List of SSH keys allowed to connect to the server.
	SshKeyIds []string `pulumi:"sshKeyIds"`
	// The tags associated with the server.
	Tags []string `pulumi:"tags"`
	// User used for the installation.
	User *string `pulumi:"user"`
	// `zone`) The zone in which the server should be created.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a BaremetalServer resource.
type BaremetalServerArgs struct {
	// A description for the server.
	Description pulumi.StringPtrInput
	// The hostname of the server.
	Hostname pulumi.StringPtrInput
	// If True, this boolean allows to create a server without the install config if you want to provide it later.
	InstallConfigAfterward pulumi.BoolPtrInput
	// The name of the server.
	Name pulumi.StringPtrInput
	// The offer name or UUID of the baremetal server.
	// Use [this endpoint](https://developers.scaleway.com/en/products/baremetal/api/#get-334154) to find the right offer.
	//
	// > **Important:** Updates to `offer` will recreate the server.
	Offer pulumi.StringInput
	// The options to enable on the server.
	// > The `options` block supports:
	Options BaremetalServerOptionArrayInput
	// The UUID of the os to install on the server.
	// Use [this endpoint](https://developers.scaleway.com/en/products/baremetal/api/#get-87598a) to find the right OS ID.
	// > **Important:** Updates to `os` will reinstall the server.
	Os pulumi.StringPtrInput
	// Password used for the installation. May be required depending on used os.
	Password pulumi.StringPtrInput
	// The private networks to attach to the server. For more information, see [the documentation](https://www.scaleway.com/en/docs/compute/elastic-metal/how-to/use-private-networks/)
	PrivateNetworks BaremetalServerPrivateNetworkArrayInput
	// `projectId`) The ID of the project the server is associated with.
	ProjectId pulumi.StringPtrInput
	// If True, this boolean allows to reinstall the server on install config changes.
	// > **Important:** Updates to `sshKeyIds`, `user`, `password`, `serviceUser` or `servicePassword` will not take effect on the server, it requires to reinstall it. To do so please set 'reinstall_on_config_changes' argument to true.
	ReinstallOnConfigChanges pulumi.BoolPtrInput
	// Password used for the service to install. May be required depending on used os.
	ServicePassword pulumi.StringPtrInput
	// User used for the service to install.
	ServiceUser pulumi.StringPtrInput
	// List of SSH keys allowed to connect to the server.
	SshKeyIds pulumi.StringArrayInput
	// The tags associated with the server.
	Tags pulumi.StringArrayInput
	// User used for the installation.
	User pulumi.StringPtrInput
	// `zone`) The zone in which the server should be created.
	Zone pulumi.StringPtrInput
}

func (BaremetalServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*baremetalServerArgs)(nil)).Elem()
}

type BaremetalServerInput interface {
	pulumi.Input

	ToBaremetalServerOutput() BaremetalServerOutput
	ToBaremetalServerOutputWithContext(ctx context.Context) BaremetalServerOutput
}

func (*BaremetalServer) ElementType() reflect.Type {
	return reflect.TypeOf((**BaremetalServer)(nil)).Elem()
}

func (i *BaremetalServer) ToBaremetalServerOutput() BaremetalServerOutput {
	return i.ToBaremetalServerOutputWithContext(context.Background())
}

func (i *BaremetalServer) ToBaremetalServerOutputWithContext(ctx context.Context) BaremetalServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BaremetalServerOutput)
}

// BaremetalServerArrayInput is an input type that accepts BaremetalServerArray and BaremetalServerArrayOutput values.
// You can construct a concrete instance of `BaremetalServerArrayInput` via:
//
//	BaremetalServerArray{ BaremetalServerArgs{...} }
type BaremetalServerArrayInput interface {
	pulumi.Input

	ToBaremetalServerArrayOutput() BaremetalServerArrayOutput
	ToBaremetalServerArrayOutputWithContext(context.Context) BaremetalServerArrayOutput
}

type BaremetalServerArray []BaremetalServerInput

func (BaremetalServerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BaremetalServer)(nil)).Elem()
}

func (i BaremetalServerArray) ToBaremetalServerArrayOutput() BaremetalServerArrayOutput {
	return i.ToBaremetalServerArrayOutputWithContext(context.Background())
}

func (i BaremetalServerArray) ToBaremetalServerArrayOutputWithContext(ctx context.Context) BaremetalServerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BaremetalServerArrayOutput)
}

// BaremetalServerMapInput is an input type that accepts BaremetalServerMap and BaremetalServerMapOutput values.
// You can construct a concrete instance of `BaremetalServerMapInput` via:
//
//	BaremetalServerMap{ "key": BaremetalServerArgs{...} }
type BaremetalServerMapInput interface {
	pulumi.Input

	ToBaremetalServerMapOutput() BaremetalServerMapOutput
	ToBaremetalServerMapOutputWithContext(context.Context) BaremetalServerMapOutput
}

type BaremetalServerMap map[string]BaremetalServerInput

func (BaremetalServerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BaremetalServer)(nil)).Elem()
}

func (i BaremetalServerMap) ToBaremetalServerMapOutput() BaremetalServerMapOutput {
	return i.ToBaremetalServerMapOutputWithContext(context.Background())
}

func (i BaremetalServerMap) ToBaremetalServerMapOutputWithContext(ctx context.Context) BaremetalServerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BaremetalServerMapOutput)
}

type BaremetalServerOutput struct{ *pulumi.OutputState }

func (BaremetalServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BaremetalServer)(nil)).Elem()
}

func (o BaremetalServerOutput) ToBaremetalServerOutput() BaremetalServerOutput {
	return o
}

func (o BaremetalServerOutput) ToBaremetalServerOutputWithContext(ctx context.Context) BaremetalServerOutput {
	return o
}

// A description for the server.
func (o BaremetalServerOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BaremetalServer) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The domain of the server.
func (o BaremetalServerOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v *BaremetalServer) pulumi.StringOutput { return v.Domain }).(pulumi.StringOutput)
}

// The hostname of the server.
func (o BaremetalServerOutput) Hostname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BaremetalServer) pulumi.StringPtrOutput { return v.Hostname }).(pulumi.StringPtrOutput)
}

// If True, this boolean allows to create a server without the install config if you want to provide it later.
func (o BaremetalServerOutput) InstallConfigAfterward() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BaremetalServer) pulumi.BoolPtrOutput { return v.InstallConfigAfterward }).(pulumi.BoolPtrOutput)
}

// (List of) The IPs of the server.
func (o BaremetalServerOutput) Ips() BaremetalServerIpArrayOutput {
	return o.ApplyT(func(v *BaremetalServer) BaremetalServerIpArrayOutput { return v.Ips }).(BaremetalServerIpArrayOutput)
}

// (List of) The IPv4 addresses of the server.
func (o BaremetalServerOutput) Ipv4s() BaremetalServerIpv4ArrayOutput {
	return o.ApplyT(func(v *BaremetalServer) BaremetalServerIpv4ArrayOutput { return v.Ipv4s }).(BaremetalServerIpv4ArrayOutput)
}

// (List of) The IPv6 addresses of the server.
func (o BaremetalServerOutput) Ipv6s() BaremetalServerIpv6ArrayOutput {
	return o.ApplyT(func(v *BaremetalServer) BaremetalServerIpv6ArrayOutput { return v.Ipv6s }).(BaremetalServerIpv6ArrayOutput)
}

// The name of the server.
func (o BaremetalServerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BaremetalServer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The offer name or UUID of the baremetal server.
// Use [this endpoint](https://developers.scaleway.com/en/products/baremetal/api/#get-334154) to find the right offer.
//
// > **Important:** Updates to `offer` will recreate the server.
func (o BaremetalServerOutput) Offer() pulumi.StringOutput {
	return o.ApplyT(func(v *BaremetalServer) pulumi.StringOutput { return v.Offer }).(pulumi.StringOutput)
}

// The ID of the offer.
func (o BaremetalServerOutput) OfferId() pulumi.StringOutput {
	return o.ApplyT(func(v *BaremetalServer) pulumi.StringOutput { return v.OfferId }).(pulumi.StringOutput)
}

// The name of the offer.
func (o BaremetalServerOutput) OfferName() pulumi.StringOutput {
	return o.ApplyT(func(v *BaremetalServer) pulumi.StringOutput { return v.OfferName }).(pulumi.StringOutput)
}

// The options to enable on the server.
// > The `options` block supports:
func (o BaremetalServerOutput) Options() BaremetalServerOptionArrayOutput {
	return o.ApplyT(func(v *BaremetalServer) BaremetalServerOptionArrayOutput { return v.Options }).(BaremetalServerOptionArrayOutput)
}

// The organization ID the server is associated with.
func (o BaremetalServerOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *BaremetalServer) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// The UUID of the os to install on the server.
// Use [this endpoint](https://developers.scaleway.com/en/products/baremetal/api/#get-87598a) to find the right OS ID.
// > **Important:** Updates to `os` will reinstall the server.
func (o BaremetalServerOutput) Os() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BaremetalServer) pulumi.StringPtrOutput { return v.Os }).(pulumi.StringPtrOutput)
}

// The name of the os.
func (o BaremetalServerOutput) OsName() pulumi.StringOutput {
	return o.ApplyT(func(v *BaremetalServer) pulumi.StringOutput { return v.OsName }).(pulumi.StringOutput)
}

// Password used for the installation. May be required depending on used os.
func (o BaremetalServerOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BaremetalServer) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// The private networks to attach to the server. For more information, see [the documentation](https://www.scaleway.com/en/docs/compute/elastic-metal/how-to/use-private-networks/)
func (o BaremetalServerOutput) PrivateNetworks() BaremetalServerPrivateNetworkArrayOutput {
	return o.ApplyT(func(v *BaremetalServer) BaremetalServerPrivateNetworkArrayOutput { return v.PrivateNetworks }).(BaremetalServerPrivateNetworkArrayOutput)
}

// `projectId`) The ID of the project the server is associated with.
func (o BaremetalServerOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *BaremetalServer) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// If True, this boolean allows to reinstall the server on install config changes.
// > **Important:** Updates to `sshKeyIds`, `user`, `password`, `serviceUser` or `servicePassword` will not take effect on the server, it requires to reinstall it. To do so please set 'reinstall_on_config_changes' argument to true.
func (o BaremetalServerOutput) ReinstallOnConfigChanges() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BaremetalServer) pulumi.BoolPtrOutput { return v.ReinstallOnConfigChanges }).(pulumi.BoolPtrOutput)
}

// Password used for the service to install. May be required depending on used os.
func (o BaremetalServerOutput) ServicePassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BaremetalServer) pulumi.StringPtrOutput { return v.ServicePassword }).(pulumi.StringPtrOutput)
}

// User used for the service to install.
func (o BaremetalServerOutput) ServiceUser() pulumi.StringOutput {
	return o.ApplyT(func(v *BaremetalServer) pulumi.StringOutput { return v.ServiceUser }).(pulumi.StringOutput)
}

// List of SSH keys allowed to connect to the server.
func (o BaremetalServerOutput) SshKeyIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BaremetalServer) pulumi.StringArrayOutput { return v.SshKeyIds }).(pulumi.StringArrayOutput)
}

// The tags associated with the server.
func (o BaremetalServerOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BaremetalServer) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// User used for the installation.
func (o BaremetalServerOutput) User() pulumi.StringOutput {
	return o.ApplyT(func(v *BaremetalServer) pulumi.StringOutput { return v.User }).(pulumi.StringOutput)
}

// `zone`) The zone in which the server should be created.
func (o BaremetalServerOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *BaremetalServer) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type BaremetalServerArrayOutput struct{ *pulumi.OutputState }

func (BaremetalServerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BaremetalServer)(nil)).Elem()
}

func (o BaremetalServerArrayOutput) ToBaremetalServerArrayOutput() BaremetalServerArrayOutput {
	return o
}

func (o BaremetalServerArrayOutput) ToBaremetalServerArrayOutputWithContext(ctx context.Context) BaremetalServerArrayOutput {
	return o
}

func (o BaremetalServerArrayOutput) Index(i pulumi.IntInput) BaremetalServerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BaremetalServer {
		return vs[0].([]*BaremetalServer)[vs[1].(int)]
	}).(BaremetalServerOutput)
}

type BaremetalServerMapOutput struct{ *pulumi.OutputState }

func (BaremetalServerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BaremetalServer)(nil)).Elem()
}

func (o BaremetalServerMapOutput) ToBaremetalServerMapOutput() BaremetalServerMapOutput {
	return o
}

func (o BaremetalServerMapOutput) ToBaremetalServerMapOutputWithContext(ctx context.Context) BaremetalServerMapOutput {
	return o
}

func (o BaremetalServerMapOutput) MapIndex(k pulumi.StringInput) BaremetalServerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BaremetalServer {
		return vs[0].(map[string]*BaremetalServer)[vs[1].(string)]
	}).(BaremetalServerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BaremetalServerInput)(nil)).Elem(), &BaremetalServer{})
	pulumi.RegisterInputType(reflect.TypeOf((*BaremetalServerArrayInput)(nil)).Elem(), BaremetalServerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BaremetalServerMapInput)(nil)).Elem(), BaremetalServerMap{})
	pulumi.RegisterOutputType(BaremetalServerOutput{})
	pulumi.RegisterOutputType(BaremetalServerArrayOutput{})
	pulumi.RegisterOutputType(BaremetalServerMapOutput{})
}
