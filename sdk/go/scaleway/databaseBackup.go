// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"errors"
	"github.com/lbrlabs/pulumi-scaleway/sdk/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates and manages Scaleway RDB database backup.
// For more information, see [the documentation](https://developers.scaleway.com/en/products/rdb/api).
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
//			_, err := scaleway.NewDatabaseBackup(ctx, "main", &scaleway.DatabaseBackupArgs{
//				InstanceId:   pulumi.Any(data.Scaleway_rdb_instance.Main.Id),
//				DatabaseName: pulumi.Any(data.Scaleway_rdb_database.Main.Name),
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
// ### With expiration
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
//			_, err := scaleway.NewDatabaseBackup(ctx, "main", &scaleway.DatabaseBackupArgs{
//				InstanceId:   pulumi.Any(data.Scaleway_rdb_instance.Main.Id),
//				DatabaseName: pulumi.Any(data.Scaleway_rdb_database.Main.Name),
//				ExpiresAt:    pulumi.String("2022-06-16T07:48:44Z"),
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
// RDB Database can be imported using the `{region}/{id}`, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:index/databaseBackup:DatabaseBackup mybackup fr-par/11111111-1111-1111-1111-111111111111
//
// ```
type DatabaseBackup struct {
	pulumi.CustomResourceState

	// Creation date (Format ISO 8601).
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Name of the database of this backup.
	DatabaseName pulumi.StringOutput `pulumi:"databaseName"`
	// Expiration date (Format ISO 8601).
	//
	// > **Important:** `expiresAt` cannot be removed after being set.
	ExpiresAt pulumi.StringPtrOutput `pulumi:"expiresAt"`
	// UUID of the rdb instance.
	//
	// > **Important:** Updates to `instanceId` will recreate the Backup.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Name of the instance of the backup.
	InstanceName pulumi.StringOutput `pulumi:"instanceName"`
	// Name of the database (e.g. `my-database`).
	Name pulumi.StringOutput `pulumi:"name"`
	// `region`) The region in which the resource exists.
	Region pulumi.StringOutput `pulumi:"region"`
	// Size of the backup (in bytes).
	Size pulumi.IntOutput `pulumi:"size"`
	// Updated date (Format ISO 8601).
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewDatabaseBackup registers a new resource with the given unique name, arguments, and options.
func NewDatabaseBackup(ctx *pulumi.Context,
	name string, args *DatabaseBackupArgs, opts ...pulumi.ResourceOption) (*DatabaseBackup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DatabaseBackup
	err := ctx.RegisterResource("scaleway:index/databaseBackup:DatabaseBackup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabaseBackup gets an existing DatabaseBackup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabaseBackup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseBackupState, opts ...pulumi.ResourceOption) (*DatabaseBackup, error) {
	var resource DatabaseBackup
	err := ctx.ReadResource("scaleway:index/databaseBackup:DatabaseBackup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabaseBackup resources.
type databaseBackupState struct {
	// Creation date (Format ISO 8601).
	CreatedAt *string `pulumi:"createdAt"`
	// Name of the database of this backup.
	DatabaseName *string `pulumi:"databaseName"`
	// Expiration date (Format ISO 8601).
	//
	// > **Important:** `expiresAt` cannot be removed after being set.
	ExpiresAt *string `pulumi:"expiresAt"`
	// UUID of the rdb instance.
	//
	// > **Important:** Updates to `instanceId` will recreate the Backup.
	InstanceId *string `pulumi:"instanceId"`
	// Name of the instance of the backup.
	InstanceName *string `pulumi:"instanceName"`
	// Name of the database (e.g. `my-database`).
	Name *string `pulumi:"name"`
	// `region`) The region in which the resource exists.
	Region *string `pulumi:"region"`
	// Size of the backup (in bytes).
	Size *int `pulumi:"size"`
	// Updated date (Format ISO 8601).
	UpdatedAt *string `pulumi:"updatedAt"`
}

type DatabaseBackupState struct {
	// Creation date (Format ISO 8601).
	CreatedAt pulumi.StringPtrInput
	// Name of the database of this backup.
	DatabaseName pulumi.StringPtrInput
	// Expiration date (Format ISO 8601).
	//
	// > **Important:** `expiresAt` cannot be removed after being set.
	ExpiresAt pulumi.StringPtrInput
	// UUID of the rdb instance.
	//
	// > **Important:** Updates to `instanceId` will recreate the Backup.
	InstanceId pulumi.StringPtrInput
	// Name of the instance of the backup.
	InstanceName pulumi.StringPtrInput
	// Name of the database (e.g. `my-database`).
	Name pulumi.StringPtrInput
	// `region`) The region in which the resource exists.
	Region pulumi.StringPtrInput
	// Size of the backup (in bytes).
	Size pulumi.IntPtrInput
	// Updated date (Format ISO 8601).
	UpdatedAt pulumi.StringPtrInput
}

func (DatabaseBackupState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseBackupState)(nil)).Elem()
}

type databaseBackupArgs struct {
	// Name of the database of this backup.
	DatabaseName string `pulumi:"databaseName"`
	// Expiration date (Format ISO 8601).
	//
	// > **Important:** `expiresAt` cannot be removed after being set.
	ExpiresAt *string `pulumi:"expiresAt"`
	// UUID of the rdb instance.
	//
	// > **Important:** Updates to `instanceId` will recreate the Backup.
	InstanceId string `pulumi:"instanceId"`
	// Name of the database (e.g. `my-database`).
	Name *string `pulumi:"name"`
	// `region`) The region in which the resource exists.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a DatabaseBackup resource.
type DatabaseBackupArgs struct {
	// Name of the database of this backup.
	DatabaseName pulumi.StringInput
	// Expiration date (Format ISO 8601).
	//
	// > **Important:** `expiresAt` cannot be removed after being set.
	ExpiresAt pulumi.StringPtrInput
	// UUID of the rdb instance.
	//
	// > **Important:** Updates to `instanceId` will recreate the Backup.
	InstanceId pulumi.StringInput
	// Name of the database (e.g. `my-database`).
	Name pulumi.StringPtrInput
	// `region`) The region in which the resource exists.
	Region pulumi.StringPtrInput
}

func (DatabaseBackupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseBackupArgs)(nil)).Elem()
}

type DatabaseBackupInput interface {
	pulumi.Input

	ToDatabaseBackupOutput() DatabaseBackupOutput
	ToDatabaseBackupOutputWithContext(ctx context.Context) DatabaseBackupOutput
}

func (*DatabaseBackup) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseBackup)(nil)).Elem()
}

func (i *DatabaseBackup) ToDatabaseBackupOutput() DatabaseBackupOutput {
	return i.ToDatabaseBackupOutputWithContext(context.Background())
}

func (i *DatabaseBackup) ToDatabaseBackupOutputWithContext(ctx context.Context) DatabaseBackupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseBackupOutput)
}

// DatabaseBackupArrayInput is an input type that accepts DatabaseBackupArray and DatabaseBackupArrayOutput values.
// You can construct a concrete instance of `DatabaseBackupArrayInput` via:
//
//	DatabaseBackupArray{ DatabaseBackupArgs{...} }
type DatabaseBackupArrayInput interface {
	pulumi.Input

	ToDatabaseBackupArrayOutput() DatabaseBackupArrayOutput
	ToDatabaseBackupArrayOutputWithContext(context.Context) DatabaseBackupArrayOutput
}

type DatabaseBackupArray []DatabaseBackupInput

func (DatabaseBackupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatabaseBackup)(nil)).Elem()
}

func (i DatabaseBackupArray) ToDatabaseBackupArrayOutput() DatabaseBackupArrayOutput {
	return i.ToDatabaseBackupArrayOutputWithContext(context.Background())
}

func (i DatabaseBackupArray) ToDatabaseBackupArrayOutputWithContext(ctx context.Context) DatabaseBackupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseBackupArrayOutput)
}

// DatabaseBackupMapInput is an input type that accepts DatabaseBackupMap and DatabaseBackupMapOutput values.
// You can construct a concrete instance of `DatabaseBackupMapInput` via:
//
//	DatabaseBackupMap{ "key": DatabaseBackupArgs{...} }
type DatabaseBackupMapInput interface {
	pulumi.Input

	ToDatabaseBackupMapOutput() DatabaseBackupMapOutput
	ToDatabaseBackupMapOutputWithContext(context.Context) DatabaseBackupMapOutput
}

type DatabaseBackupMap map[string]DatabaseBackupInput

func (DatabaseBackupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatabaseBackup)(nil)).Elem()
}

func (i DatabaseBackupMap) ToDatabaseBackupMapOutput() DatabaseBackupMapOutput {
	return i.ToDatabaseBackupMapOutputWithContext(context.Background())
}

func (i DatabaseBackupMap) ToDatabaseBackupMapOutputWithContext(ctx context.Context) DatabaseBackupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseBackupMapOutput)
}

type DatabaseBackupOutput struct{ *pulumi.OutputState }

func (DatabaseBackupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseBackup)(nil)).Elem()
}

func (o DatabaseBackupOutput) ToDatabaseBackupOutput() DatabaseBackupOutput {
	return o
}

func (o DatabaseBackupOutput) ToDatabaseBackupOutputWithContext(ctx context.Context) DatabaseBackupOutput {
	return o
}

// Creation date (Format ISO 8601).
func (o DatabaseBackupOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseBackup) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Name of the database of this backup.
func (o DatabaseBackupOutput) DatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseBackup) pulumi.StringOutput { return v.DatabaseName }).(pulumi.StringOutput)
}

// Expiration date (Format ISO 8601).
//
// > **Important:** `expiresAt` cannot be removed after being set.
func (o DatabaseBackupOutput) ExpiresAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseBackup) pulumi.StringPtrOutput { return v.ExpiresAt }).(pulumi.StringPtrOutput)
}

// UUID of the rdb instance.
//
// > **Important:** Updates to `instanceId` will recreate the Backup.
func (o DatabaseBackupOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseBackup) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Name of the instance of the backup.
func (o DatabaseBackupOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseBackup) pulumi.StringOutput { return v.InstanceName }).(pulumi.StringOutput)
}

// Name of the database (e.g. `my-database`).
func (o DatabaseBackupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseBackup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// `region`) The region in which the resource exists.
func (o DatabaseBackupOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseBackup) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Size of the backup (in bytes).
func (o DatabaseBackupOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v *DatabaseBackup) pulumi.IntOutput { return v.Size }).(pulumi.IntOutput)
}

// Updated date (Format ISO 8601).
func (o DatabaseBackupOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseBackup) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type DatabaseBackupArrayOutput struct{ *pulumi.OutputState }

func (DatabaseBackupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatabaseBackup)(nil)).Elem()
}

func (o DatabaseBackupArrayOutput) ToDatabaseBackupArrayOutput() DatabaseBackupArrayOutput {
	return o
}

func (o DatabaseBackupArrayOutput) ToDatabaseBackupArrayOutputWithContext(ctx context.Context) DatabaseBackupArrayOutput {
	return o
}

func (o DatabaseBackupArrayOutput) Index(i pulumi.IntInput) DatabaseBackupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DatabaseBackup {
		return vs[0].([]*DatabaseBackup)[vs[1].(int)]
	}).(DatabaseBackupOutput)
}

type DatabaseBackupMapOutput struct{ *pulumi.OutputState }

func (DatabaseBackupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatabaseBackup)(nil)).Elem()
}

func (o DatabaseBackupMapOutput) ToDatabaseBackupMapOutput() DatabaseBackupMapOutput {
	return o
}

func (o DatabaseBackupMapOutput) ToDatabaseBackupMapOutputWithContext(ctx context.Context) DatabaseBackupMapOutput {
	return o
}

func (o DatabaseBackupMapOutput) MapIndex(k pulumi.StringInput) DatabaseBackupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DatabaseBackup {
		return vs[0].(map[string]*DatabaseBackup)[vs[1].(string)]
	}).(DatabaseBackupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseBackupInput)(nil)).Elem(), &DatabaseBackup{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseBackupArrayInput)(nil)).Elem(), DatabaseBackupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseBackupMapInput)(nil)).Elem(), DatabaseBackupMap{})
	pulumi.RegisterOutputType(DatabaseBackupOutput{})
	pulumi.RegisterOutputType(DatabaseBackupArrayOutput{})
	pulumi.RegisterOutputType(DatabaseBackupMapOutput{})
}
