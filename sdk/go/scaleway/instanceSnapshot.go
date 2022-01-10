// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates and manages Scaleway Compute Snapshots.
// For more information, see [the documentation](https://developers.scaleway.com/en/products/instance/api/#snapshots-756fae).
//
// ## Example
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
// 		_, err := scaleway.NewInstanceSnapshot(ctx, "main", &scaleway.InstanceSnapshotArgs{
// 			VolumeId: pulumi.String("11111111-1111-1111-1111-111111111111"),
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
// Snapshots can be imported using the `{zone}/{id}`, e.g. bash
//
// ```sh
//  $ pulumi import scaleway:index/instanceSnapshot:InstanceSnapshot main fr-par-1/11111111-1111-1111-1111-111111111111
// ```
type InstanceSnapshot struct {
	pulumi.CustomResourceState

	// The snapshot creation time.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The name of the snapshot. If not provided it will be randomly generated.
	Name pulumi.StringOutput `pulumi:"name"`
	// The organization ID the snapshot is associated with.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// `projectId`) The ID of the project the snapshot is associated with.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// (Optional) The size of the snapshot.
	SizeInGb pulumi.IntOutput `pulumi:"sizeInGb"`
	// The type of the snapshot. The possible values are: `bSsd` (Block SSD), `lSsd` (Local SSD).
	Type pulumi.StringOutput `pulumi:"type"`
	// The ID of the volume to take a snapshot from.
	VolumeId pulumi.StringOutput `pulumi:"volumeId"`
	// `zone`) The zone in which the snapshot should be created.
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
	// The snapshot creation time.
	CreatedAt *string `pulumi:"createdAt"`
	// The name of the snapshot. If not provided it will be randomly generated.
	Name *string `pulumi:"name"`
	// The organization ID the snapshot is associated with.
	OrganizationId *string `pulumi:"organizationId"`
	// `projectId`) The ID of the project the snapshot is associated with.
	ProjectId *string `pulumi:"projectId"`
	// (Optional) The size of the snapshot.
	SizeInGb *int `pulumi:"sizeInGb"`
	// The type of the snapshot. The possible values are: `bSsd` (Block SSD), `lSsd` (Local SSD).
	Type *string `pulumi:"type"`
	// The ID of the volume to take a snapshot from.
	VolumeId *string `pulumi:"volumeId"`
	// `zone`) The zone in which the snapshot should be created.
	Zone *string `pulumi:"zone"`
}

type InstanceSnapshotState struct {
	// The snapshot creation time.
	CreatedAt pulumi.StringPtrInput
	// The name of the snapshot. If not provided it will be randomly generated.
	Name pulumi.StringPtrInput
	// The organization ID the snapshot is associated with.
	OrganizationId pulumi.StringPtrInput
	// `projectId`) The ID of the project the snapshot is associated with.
	ProjectId pulumi.StringPtrInput
	// (Optional) The size of the snapshot.
	SizeInGb pulumi.IntPtrInput
	// The type of the snapshot. The possible values are: `bSsd` (Block SSD), `lSsd` (Local SSD).
	Type pulumi.StringPtrInput
	// The ID of the volume to take a snapshot from.
	VolumeId pulumi.StringPtrInput
	// `zone`) The zone in which the snapshot should be created.
	Zone pulumi.StringPtrInput
}

func (InstanceSnapshotState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceSnapshotState)(nil)).Elem()
}

type instanceSnapshotArgs struct {
	// The name of the snapshot. If not provided it will be randomly generated.
	Name *string `pulumi:"name"`
	// `projectId`) The ID of the project the snapshot is associated with.
	ProjectId *string `pulumi:"projectId"`
	// The ID of the volume to take a snapshot from.
	VolumeId string `pulumi:"volumeId"`
	// `zone`) The zone in which the snapshot should be created.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a InstanceSnapshot resource.
type InstanceSnapshotArgs struct {
	// The name of the snapshot. If not provided it will be randomly generated.
	Name pulumi.StringPtrInput
	// `projectId`) The ID of the project the snapshot is associated with.
	ProjectId pulumi.StringPtrInput
	// The ID of the volume to take a snapshot from.
	VolumeId pulumi.StringInput
	// `zone`) The zone in which the snapshot should be created.
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
	return reflect.TypeOf((*InstanceSnapshot)(nil))
}

func (i *InstanceSnapshot) ToInstanceSnapshotOutput() InstanceSnapshotOutput {
	return i.ToInstanceSnapshotOutputWithContext(context.Background())
}

func (i *InstanceSnapshot) ToInstanceSnapshotOutputWithContext(ctx context.Context) InstanceSnapshotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceSnapshotOutput)
}

type InstanceSnapshotOutput struct{ *pulumi.OutputState }

func (InstanceSnapshotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceSnapshot)(nil))
}

func (o InstanceSnapshotOutput) ToInstanceSnapshotOutput() InstanceSnapshotOutput {
	return o
}

func (o InstanceSnapshotOutput) ToInstanceSnapshotOutputWithContext(ctx context.Context) InstanceSnapshotOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceSnapshotInput)(nil)).Elem(), &InstanceSnapshot{})
	pulumi.RegisterOutputType(InstanceSnapshotOutput{})
}
