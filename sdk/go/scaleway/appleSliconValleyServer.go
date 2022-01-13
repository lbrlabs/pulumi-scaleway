// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AppleSliconValleyServer struct {
	pulumi.CustomResourceState

	// The date and time of the creation of the server
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The minimal date and time on which you can delete this server due to Apple licence
	DeletableAt pulumi.StringOutput `pulumi:"deletableAt"`
	// IPv4 address of the server
	Ip pulumi.StringOutput `pulumi:"ip"`
	// Name of the server
	Name pulumi.StringOutput `pulumi:"name"`
	// The organization_id you want to attach the resource to
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// The project_id you want to attach the resource to
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Type of the server
	Type pulumi.StringOutput `pulumi:"type"`
	// The date and time of the last update of the server
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// VNC url use to connect remotely to the desktop GUI
	VncUrl pulumi.StringOutput `pulumi:"vncUrl"`
	// The zone you want to attach the resource to
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewAppleSliconValleyServer registers a new resource with the given unique name, arguments, and options.
func NewAppleSliconValleyServer(ctx *pulumi.Context,
	name string, args *AppleSliconValleyServerArgs, opts ...pulumi.ResourceOption) (*AppleSliconValleyServer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource AppleSliconValleyServer
	err := ctx.RegisterResource("scaleway:index/appleSliconValleyServer:AppleSliconValleyServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppleSliconValleyServer gets an existing AppleSliconValleyServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppleSliconValleyServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppleSliconValleyServerState, opts ...pulumi.ResourceOption) (*AppleSliconValleyServer, error) {
	var resource AppleSliconValleyServer
	err := ctx.ReadResource("scaleway:index/appleSliconValleyServer:AppleSliconValleyServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppleSliconValleyServer resources.
type appleSliconValleyServerState struct {
	// The date and time of the creation of the server
	CreatedAt *string `pulumi:"createdAt"`
	// The minimal date and time on which you can delete this server due to Apple licence
	DeletableAt *string `pulumi:"deletableAt"`
	// IPv4 address of the server
	Ip *string `pulumi:"ip"`
	// Name of the server
	Name *string `pulumi:"name"`
	// The organization_id you want to attach the resource to
	OrganizationId *string `pulumi:"organizationId"`
	// The project_id you want to attach the resource to
	ProjectId *string `pulumi:"projectId"`
	// Type of the server
	Type *string `pulumi:"type"`
	// The date and time of the last update of the server
	UpdatedAt *string `pulumi:"updatedAt"`
	// VNC url use to connect remotely to the desktop GUI
	VncUrl *string `pulumi:"vncUrl"`
	// The zone you want to attach the resource to
	Zone *string `pulumi:"zone"`
}

type AppleSliconValleyServerState struct {
	// The date and time of the creation of the server
	CreatedAt pulumi.StringPtrInput
	// The minimal date and time on which you can delete this server due to Apple licence
	DeletableAt pulumi.StringPtrInput
	// IPv4 address of the server
	Ip pulumi.StringPtrInput
	// Name of the server
	Name pulumi.StringPtrInput
	// The organization_id you want to attach the resource to
	OrganizationId pulumi.StringPtrInput
	// The project_id you want to attach the resource to
	ProjectId pulumi.StringPtrInput
	// Type of the server
	Type pulumi.StringPtrInput
	// The date and time of the last update of the server
	UpdatedAt pulumi.StringPtrInput
	// VNC url use to connect remotely to the desktop GUI
	VncUrl pulumi.StringPtrInput
	// The zone you want to attach the resource to
	Zone pulumi.StringPtrInput
}

func (AppleSliconValleyServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*appleSliconValleyServerState)(nil)).Elem()
}

type appleSliconValleyServerArgs struct {
	// Name of the server
	Name *string `pulumi:"name"`
	// The project_id you want to attach the resource to
	ProjectId *string `pulumi:"projectId"`
	// Type of the server
	Type string `pulumi:"type"`
	// The zone you want to attach the resource to
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a AppleSliconValleyServer resource.
type AppleSliconValleyServerArgs struct {
	// Name of the server
	Name pulumi.StringPtrInput
	// The project_id you want to attach the resource to
	ProjectId pulumi.StringPtrInput
	// Type of the server
	Type pulumi.StringInput
	// The zone you want to attach the resource to
	Zone pulumi.StringPtrInput
}

func (AppleSliconValleyServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appleSliconValleyServerArgs)(nil)).Elem()
}

type AppleSliconValleyServerInput interface {
	pulumi.Input

	ToAppleSliconValleyServerOutput() AppleSliconValleyServerOutput
	ToAppleSliconValleyServerOutputWithContext(ctx context.Context) AppleSliconValleyServerOutput
}

func (*AppleSliconValleyServer) ElementType() reflect.Type {
	return reflect.TypeOf((**AppleSliconValleyServer)(nil)).Elem()
}

func (i *AppleSliconValleyServer) ToAppleSliconValleyServerOutput() AppleSliconValleyServerOutput {
	return i.ToAppleSliconValleyServerOutputWithContext(context.Background())
}

func (i *AppleSliconValleyServer) ToAppleSliconValleyServerOutputWithContext(ctx context.Context) AppleSliconValleyServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppleSliconValleyServerOutput)
}

type AppleSliconValleyServerOutput struct{ *pulumi.OutputState }

func (AppleSliconValleyServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppleSliconValleyServer)(nil)).Elem()
}

func (o AppleSliconValleyServerOutput) ToAppleSliconValleyServerOutput() AppleSliconValleyServerOutput {
	return o
}

func (o AppleSliconValleyServerOutput) ToAppleSliconValleyServerOutputWithContext(ctx context.Context) AppleSliconValleyServerOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppleSliconValleyServerInput)(nil)).Elem(), &AppleSliconValleyServer{})
	pulumi.RegisterOutputType(AppleSliconValleyServerOutput{})
}
