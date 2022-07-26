// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type InstanceSnapshot struct {
	pulumi.CustomResourceState

	// The date and time of the creation of the snapshot
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The name of the snapshot
	Name pulumi.StringOutput `pulumi:"name"`
	// The organization_id you want to attach the resource to
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// The project_id you want to attach the resource to
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The size of the snapshot in gigabyte
	SizeInGb pulumi.IntOutput `pulumi:"sizeInGb"`
	// The tags associated with the snapshot
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The volume type of the snapshot
	Type pulumi.StringOutput `pulumi:"type"`
	// ID of the volume to take a snapshot from
	VolumeId pulumi.StringOutput `pulumi:"volumeId"`
	// The zone you want to attach the resource to
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewInstanceSnapshot registers a new resource with the given unique name, arguments, and options.
func NewInstanceSnapshot(ctx *pulumi.Context,
	name string, args *InstanceSnapshotArgs, opts ...pulumi.ResourceOption) (*InstanceSnapshot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.VolumeId == nil {
		return nil, errors.New("invalid value for required argument 'VolumeId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource InstanceSnapshot
	err := ctx.RegisterResource("scaleway:index/instanceSnapshot:InstanceSnapshot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceSnapshot gets an existing InstanceSnapshot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceSnapshot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceSnapshotState, opts ...pulumi.ResourceOption) (*InstanceSnapshot, error) {
	var resource InstanceSnapshot
	err := ctx.ReadResource("scaleway:index/instanceSnapshot:InstanceSnapshot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceSnapshot resources.
type instanceSnapshotState struct {
	// The date and time of the creation of the snapshot
	CreatedAt *string `pulumi:"createdAt"`
	// The name of the snapshot
	Name *string `pulumi:"name"`
	// The organization_id you want to attach the resource to
	OrganizationId *string `pulumi:"organizationId"`
	// The project_id you want to attach the resource to
	ProjectId *string `pulumi:"projectId"`
	// The size of the snapshot in gigabyte
	SizeInGb *int `pulumi:"sizeInGb"`
	// The tags associated with the snapshot
	Tags []string `pulumi:"tags"`
	// The volume type of the snapshot
	Type *string `pulumi:"type"`
	// ID of the volume to take a snapshot from
	VolumeId *string `pulumi:"volumeId"`
	// The zone you want to attach the resource to
	Zone *string `pulumi:"zone"`
}

type InstanceSnapshotState struct {
	// The date and time of the creation of the snapshot
	CreatedAt pulumi.StringPtrInput
	// The name of the snapshot
	Name pulumi.StringPtrInput
	// The organization_id you want to attach the resource to
	OrganizationId pulumi.StringPtrInput
	// The project_id you want to attach the resource to
	ProjectId pulumi.StringPtrInput
	// The size of the snapshot in gigabyte
	SizeInGb pulumi.IntPtrInput
	// The tags associated with the snapshot
	Tags pulumi.StringArrayInput
	// The volume type of the snapshot
	Type pulumi.StringPtrInput
	// ID of the volume to take a snapshot from
	VolumeId pulumi.StringPtrInput
	// The zone you want to attach the resource to
	Zone pulumi.StringPtrInput
}

func (InstanceSnapshotState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceSnapshotState)(nil)).Elem()
}

type instanceSnapshotArgs struct {
	// The name of the snapshot
	Name *string `pulumi:"name"`
	// The project_id you want to attach the resource to
	ProjectId *string `pulumi:"projectId"`
	// The tags associated with the snapshot
	Tags []string `pulumi:"tags"`
	// ID of the volume to take a snapshot from
	VolumeId string `pulumi:"volumeId"`
	// The zone you want to attach the resource to
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a InstanceSnapshot resource.
type InstanceSnapshotArgs struct {
	// The name of the snapshot
	Name pulumi.StringPtrInput
	// The project_id you want to attach the resource to
	ProjectId pulumi.StringPtrInput
	// The tags associated with the snapshot
	Tags pulumi.StringArrayInput
	// ID of the volume to take a snapshot from
	VolumeId pulumi.StringInput
	// The zone you want to attach the resource to
	Zone pulumi.StringPtrInput
}

func (InstanceSnapshotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceSnapshotArgs)(nil)).Elem()
}

type InstanceSnapshotInput interface {
	pulumi.Input

	ToInstanceSnapshotOutput() InstanceSnapshotOutput
	ToInstanceSnapshotOutputWithContext(ctx context.Context) InstanceSnapshotOutput
}

func (*InstanceSnapshot) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceSnapshot)(nil)).Elem()
}

func (i *InstanceSnapshot) ToInstanceSnapshotOutput() InstanceSnapshotOutput {
	return i.ToInstanceSnapshotOutputWithContext(context.Background())
}

func (i *InstanceSnapshot) ToInstanceSnapshotOutputWithContext(ctx context.Context) InstanceSnapshotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceSnapshotOutput)
}

// InstanceSnapshotArrayInput is an input type that accepts InstanceSnapshotArray and InstanceSnapshotArrayOutput values.
// You can construct a concrete instance of `InstanceSnapshotArrayInput` via:
//
//          InstanceSnapshotArray{ InstanceSnapshotArgs{...} }
type InstanceSnapshotArrayInput interface {
	pulumi.Input

	ToInstanceSnapshotArrayOutput() InstanceSnapshotArrayOutput
	ToInstanceSnapshotArrayOutputWithContext(context.Context) InstanceSnapshotArrayOutput
}

type InstanceSnapshotArray []InstanceSnapshotInput

func (InstanceSnapshotArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceSnapshot)(nil)).Elem()
}

func (i InstanceSnapshotArray) ToInstanceSnapshotArrayOutput() InstanceSnapshotArrayOutput {
	return i.ToInstanceSnapshotArrayOutputWithContext(context.Background())
}

func (i InstanceSnapshotArray) ToInstanceSnapshotArrayOutputWithContext(ctx context.Context) InstanceSnapshotArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceSnapshotArrayOutput)
}

// InstanceSnapshotMapInput is an input type that accepts InstanceSnapshotMap and InstanceSnapshotMapOutput values.
// You can construct a concrete instance of `InstanceSnapshotMapInput` via:
//
//          InstanceSnapshotMap{ "key": InstanceSnapshotArgs{...} }
type InstanceSnapshotMapInput interface {
	pulumi.Input

	ToInstanceSnapshotMapOutput() InstanceSnapshotMapOutput
	ToInstanceSnapshotMapOutputWithContext(context.Context) InstanceSnapshotMapOutput
}

type InstanceSnapshotMap map[string]InstanceSnapshotInput

func (InstanceSnapshotMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceSnapshot)(nil)).Elem()
}

func (i InstanceSnapshotMap) ToInstanceSnapshotMapOutput() InstanceSnapshotMapOutput {
	return i.ToInstanceSnapshotMapOutputWithContext(context.Background())
}

func (i InstanceSnapshotMap) ToInstanceSnapshotMapOutputWithContext(ctx context.Context) InstanceSnapshotMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceSnapshotMapOutput)
}

type InstanceSnapshotOutput struct{ *pulumi.OutputState }

func (InstanceSnapshotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceSnapshot)(nil)).Elem()
}

func (o InstanceSnapshotOutput) ToInstanceSnapshotOutput() InstanceSnapshotOutput {
	return o
}

func (o InstanceSnapshotOutput) ToInstanceSnapshotOutputWithContext(ctx context.Context) InstanceSnapshotOutput {
	return o
}

// The date and time of the creation of the snapshot
func (o InstanceSnapshotOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceSnapshot) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The name of the snapshot
func (o InstanceSnapshotOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceSnapshot) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The organization_id you want to attach the resource to
func (o InstanceSnapshotOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceSnapshot) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// The project_id you want to attach the resource to
func (o InstanceSnapshotOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceSnapshot) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The size of the snapshot in gigabyte
func (o InstanceSnapshotOutput) SizeInGb() pulumi.IntOutput {
	return o.ApplyT(func(v *InstanceSnapshot) pulumi.IntOutput { return v.SizeInGb }).(pulumi.IntOutput)
}

// The tags associated with the snapshot
func (o InstanceSnapshotOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *InstanceSnapshot) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// The volume type of the snapshot
func (o InstanceSnapshotOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceSnapshot) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// ID of the volume to take a snapshot from
func (o InstanceSnapshotOutput) VolumeId() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceSnapshot) pulumi.StringOutput { return v.VolumeId }).(pulumi.StringOutput)
}

// The zone you want to attach the resource to
func (o InstanceSnapshotOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceSnapshot) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type InstanceSnapshotArrayOutput struct{ *pulumi.OutputState }

func (InstanceSnapshotArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceSnapshot)(nil)).Elem()
}

func (o InstanceSnapshotArrayOutput) ToInstanceSnapshotArrayOutput() InstanceSnapshotArrayOutput {
	return o
}

func (o InstanceSnapshotArrayOutput) ToInstanceSnapshotArrayOutputWithContext(ctx context.Context) InstanceSnapshotArrayOutput {
	return o
}

func (o InstanceSnapshotArrayOutput) Index(i pulumi.IntInput) InstanceSnapshotOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *InstanceSnapshot {
		return vs[0].([]*InstanceSnapshot)[vs[1].(int)]
	}).(InstanceSnapshotOutput)
}

type InstanceSnapshotMapOutput struct{ *pulumi.OutputState }

func (InstanceSnapshotMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceSnapshot)(nil)).Elem()
}

func (o InstanceSnapshotMapOutput) ToInstanceSnapshotMapOutput() InstanceSnapshotMapOutput {
	return o
}

func (o InstanceSnapshotMapOutput) ToInstanceSnapshotMapOutputWithContext(ctx context.Context) InstanceSnapshotMapOutput {
	return o
}

func (o InstanceSnapshotMapOutput) MapIndex(k pulumi.StringInput) InstanceSnapshotOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *InstanceSnapshot {
		return vs[0].(map[string]*InstanceSnapshot)[vs[1].(string)]
	}).(InstanceSnapshotOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceSnapshotInput)(nil)).Elem(), &InstanceSnapshot{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceSnapshotArrayInput)(nil)).Elem(), InstanceSnapshotArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceSnapshotMapInput)(nil)).Elem(), InstanceSnapshotMap{})
	pulumi.RegisterOutputType(InstanceSnapshotOutput{})
	pulumi.RegisterOutputType(InstanceSnapshotArrayOutput{})
	pulumi.RegisterOutputType(InstanceSnapshotMapOutput{})
}
