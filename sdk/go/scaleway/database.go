// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates and manages Scaleway RDB database.
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
// 	"github.com/pulumi/pulumi-scaleway/sdk/go/scaleway"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := scaleway.NewDatabase(ctx, "main", &scaleway.DatabaseArgs{
// 			InstanceId: pulumi.Any(scaleway_rdb_instance.Main.Id),
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
// RDB Database can be imported using the `{region}/{id}/{DBNAME}`, e.g. bash
//
// ```sh
//  $ pulumi import scaleway:index/database:Database rdb01_mydb fr-par/11111111-1111-1111-1111-111111111111/mydb
// ```
type Database struct {
	pulumi.CustomResourceState

	// UUID of the instance where to create the database.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Whether or not the database is managed or not.
	Managed pulumi.BoolOutput `pulumi:"managed"`
	// Name of the database (e.g. `my-new-database`).
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the owner of the database.
	Owner pulumi.StringOutput `pulumi:"owner"`
	// Size of the database (in bytes).
	Size pulumi.StringOutput `pulumi:"size"`
}

// NewDatabase registers a new resource with the given unique name, arguments, and options.
func NewDatabase(ctx *pulumi.Context,
	name string, args *DatabaseArgs, opts ...pulumi.ResourceOption) (*Database, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	var resource Database
	err := ctx.RegisterResource("scaleway:index/database:Database", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabase gets an existing Database resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseState, opts ...pulumi.ResourceOption) (*Database, error) {
	var resource Database
	err := ctx.ReadResource("scaleway:index/database:Database", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Database resources.
type databaseState struct {
	// UUID of the instance where to create the database.
	InstanceId *string `pulumi:"instanceId"`
	// Whether or not the database is managed or not.
	Managed *bool `pulumi:"managed"`
	// Name of the database (e.g. `my-new-database`).
	Name *string `pulumi:"name"`
	// The name of the owner of the database.
	Owner *string `pulumi:"owner"`
	// Size of the database (in bytes).
	Size *string `pulumi:"size"`
}

type DatabaseState struct {
	// UUID of the instance where to create the database.
	InstanceId pulumi.StringPtrInput
	// Whether or not the database is managed or not.
	Managed pulumi.BoolPtrInput
	// Name of the database (e.g. `my-new-database`).
	Name pulumi.StringPtrInput
	// The name of the owner of the database.
	Owner pulumi.StringPtrInput
	// Size of the database (in bytes).
	Size pulumi.StringPtrInput
}

func (DatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseState)(nil)).Elem()
}

type databaseArgs struct {
	// UUID of the instance where to create the database.
	InstanceId string `pulumi:"instanceId"`
	// Name of the database (e.g. `my-new-database`).
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a Database resource.
type DatabaseArgs struct {
	// UUID of the instance where to create the database.
	InstanceId pulumi.StringInput
	// Name of the database (e.g. `my-new-database`).
	Name pulumi.StringPtrInput
}

func (DatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseArgs)(nil)).Elem()
}

type DatabaseInput interface {
	pulumi.Input

	ToDatabaseOutput() DatabaseOutput
	ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput
}

func (*Database) ElementType() reflect.Type {
	return reflect.TypeOf((*Database)(nil))
}

func (i *Database) ToDatabaseOutput() DatabaseOutput {
	return i.ToDatabaseOutputWithContext(context.Background())
}

func (i *Database) ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseOutput)
}

type DatabaseOutput struct{ *pulumi.OutputState }

func (DatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Database)(nil))
}

func (o DatabaseOutput) ToDatabaseOutput() DatabaseOutput {
	return o
}

func (o DatabaseOutput) ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseInput)(nil)).Elem(), &Database{})
	pulumi.RegisterOutputType(DatabaseOutput{})
}