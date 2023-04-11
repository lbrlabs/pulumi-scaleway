// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates and manages Scaleway Database read replicas.
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
//			instance, err := scaleway.NewDatabaseInstance(ctx, "instance", &scaleway.DatabaseInstanceArgs{
//				NodeType:      pulumi.String("db-dev-s"),
//				Engine:        pulumi.String("PostgreSQL-14"),
//				IsHaCluster:   pulumi.Bool(false),
//				DisableBackup: pulumi.Bool(true),
//				UserName:      pulumi.String("my_initial_user"),
//				Password:      pulumi.String("thiZ_is_v&ry_s3cret"),
//				Tags: pulumi.StringArray{
//					pulumi.String("terraform-test"),
//					pulumi.String("scaleway_rdb_read_replica"),
//					pulumi.String("minimal"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewDatabaseReadReplica(ctx, "replica", &scaleway.DatabaseReadReplicaArgs{
//				InstanceId:   instance.ID(),
//				DirectAccess: nil,
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
// ### Private network
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
//			instance, err := scaleway.NewDatabaseInstance(ctx, "instance", &scaleway.DatabaseInstanceArgs{
//				NodeType:      pulumi.String("db-dev-s"),
//				Engine:        pulumi.String("PostgreSQL-14"),
//				IsHaCluster:   pulumi.Bool(false),
//				DisableBackup: pulumi.Bool(true),
//				UserName:      pulumi.String("my_initial_user"),
//				Password:      pulumi.String("thiZ_is_v&ry_s3cret"),
//			})
//			if err != nil {
//				return err
//			}
//			pn, err := scaleway.NewVpcPrivateNetwork(ctx, "pn", nil)
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewDatabaseReadReplica(ctx, "replica", &scaleway.DatabaseReadReplicaArgs{
//				InstanceId: instance.ID(),
//				PrivateNetwork: &scaleway.DatabaseReadReplicaPrivateNetworkArgs{
//					PrivateNetworkId: pn.ID(),
//					ServiceIp:        pulumi.String("192.168.1.254/24"),
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
// ## Import
//
// Database Read replica can be imported using the `{region}/{id}`, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:index/databaseReadReplica:DatabaseReadReplica rr fr-par/11111111-1111-1111-1111-111111111111
//
// ```
type DatabaseReadReplica struct {
	pulumi.CustomResourceState

	// Creates a direct access endpoint to rdb replica.
	DirectAccess DatabaseReadReplicaDirectAccessPtrOutput `pulumi:"directAccess"`
	// UUID of the rdb instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Create an endpoint in a private network.
	PrivateNetwork DatabaseReadReplicaPrivateNetworkPtrOutput `pulumi:"privateNetwork"`
	// `region`) The region
	// in which the Database read replica should be created.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewDatabaseReadReplica registers a new resource with the given unique name, arguments, and options.
func NewDatabaseReadReplica(ctx *pulumi.Context,
	name string, args *DatabaseReadReplicaArgs, opts ...pulumi.ResourceOption) (*DatabaseReadReplica, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource DatabaseReadReplica
	err := ctx.RegisterResource("scaleway:index/databaseReadReplica:DatabaseReadReplica", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabaseReadReplica gets an existing DatabaseReadReplica resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabaseReadReplica(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseReadReplicaState, opts ...pulumi.ResourceOption) (*DatabaseReadReplica, error) {
	var resource DatabaseReadReplica
	err := ctx.ReadResource("scaleway:index/databaseReadReplica:DatabaseReadReplica", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabaseReadReplica resources.
type databaseReadReplicaState struct {
	// Creates a direct access endpoint to rdb replica.
	DirectAccess *DatabaseReadReplicaDirectAccess `pulumi:"directAccess"`
	// UUID of the rdb instance.
	InstanceId *string `pulumi:"instanceId"`
	// Create an endpoint in a private network.
	PrivateNetwork *DatabaseReadReplicaPrivateNetwork `pulumi:"privateNetwork"`
	// `region`) The region
	// in which the Database read replica should be created.
	Region *string `pulumi:"region"`
}

type DatabaseReadReplicaState struct {
	// Creates a direct access endpoint to rdb replica.
	DirectAccess DatabaseReadReplicaDirectAccessPtrInput
	// UUID of the rdb instance.
	InstanceId pulumi.StringPtrInput
	// Create an endpoint in a private network.
	PrivateNetwork DatabaseReadReplicaPrivateNetworkPtrInput
	// `region`) The region
	// in which the Database read replica should be created.
	Region pulumi.StringPtrInput
}

func (DatabaseReadReplicaState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseReadReplicaState)(nil)).Elem()
}

type databaseReadReplicaArgs struct {
	// Creates a direct access endpoint to rdb replica.
	DirectAccess *DatabaseReadReplicaDirectAccess `pulumi:"directAccess"`
	// UUID of the rdb instance.
	InstanceId string `pulumi:"instanceId"`
	// Create an endpoint in a private network.
	PrivateNetwork *DatabaseReadReplicaPrivateNetwork `pulumi:"privateNetwork"`
	// `region`) The region
	// in which the Database read replica should be created.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a DatabaseReadReplica resource.
type DatabaseReadReplicaArgs struct {
	// Creates a direct access endpoint to rdb replica.
	DirectAccess DatabaseReadReplicaDirectAccessPtrInput
	// UUID of the rdb instance.
	InstanceId pulumi.StringInput
	// Create an endpoint in a private network.
	PrivateNetwork DatabaseReadReplicaPrivateNetworkPtrInput
	// `region`) The region
	// in which the Database read replica should be created.
	Region pulumi.StringPtrInput
}

func (DatabaseReadReplicaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseReadReplicaArgs)(nil)).Elem()
}

type DatabaseReadReplicaInput interface {
	pulumi.Input

	ToDatabaseReadReplicaOutput() DatabaseReadReplicaOutput
	ToDatabaseReadReplicaOutputWithContext(ctx context.Context) DatabaseReadReplicaOutput
}

func (*DatabaseReadReplica) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseReadReplica)(nil)).Elem()
}

func (i *DatabaseReadReplica) ToDatabaseReadReplicaOutput() DatabaseReadReplicaOutput {
	return i.ToDatabaseReadReplicaOutputWithContext(context.Background())
}

func (i *DatabaseReadReplica) ToDatabaseReadReplicaOutputWithContext(ctx context.Context) DatabaseReadReplicaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseReadReplicaOutput)
}

// DatabaseReadReplicaArrayInput is an input type that accepts DatabaseReadReplicaArray and DatabaseReadReplicaArrayOutput values.
// You can construct a concrete instance of `DatabaseReadReplicaArrayInput` via:
//
//	DatabaseReadReplicaArray{ DatabaseReadReplicaArgs{...} }
type DatabaseReadReplicaArrayInput interface {
	pulumi.Input

	ToDatabaseReadReplicaArrayOutput() DatabaseReadReplicaArrayOutput
	ToDatabaseReadReplicaArrayOutputWithContext(context.Context) DatabaseReadReplicaArrayOutput
}

type DatabaseReadReplicaArray []DatabaseReadReplicaInput

func (DatabaseReadReplicaArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatabaseReadReplica)(nil)).Elem()
}

func (i DatabaseReadReplicaArray) ToDatabaseReadReplicaArrayOutput() DatabaseReadReplicaArrayOutput {
	return i.ToDatabaseReadReplicaArrayOutputWithContext(context.Background())
}

func (i DatabaseReadReplicaArray) ToDatabaseReadReplicaArrayOutputWithContext(ctx context.Context) DatabaseReadReplicaArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseReadReplicaArrayOutput)
}

// DatabaseReadReplicaMapInput is an input type that accepts DatabaseReadReplicaMap and DatabaseReadReplicaMapOutput values.
// You can construct a concrete instance of `DatabaseReadReplicaMapInput` via:
//
//	DatabaseReadReplicaMap{ "key": DatabaseReadReplicaArgs{...} }
type DatabaseReadReplicaMapInput interface {
	pulumi.Input

	ToDatabaseReadReplicaMapOutput() DatabaseReadReplicaMapOutput
	ToDatabaseReadReplicaMapOutputWithContext(context.Context) DatabaseReadReplicaMapOutput
}

type DatabaseReadReplicaMap map[string]DatabaseReadReplicaInput

func (DatabaseReadReplicaMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatabaseReadReplica)(nil)).Elem()
}

func (i DatabaseReadReplicaMap) ToDatabaseReadReplicaMapOutput() DatabaseReadReplicaMapOutput {
	return i.ToDatabaseReadReplicaMapOutputWithContext(context.Background())
}

func (i DatabaseReadReplicaMap) ToDatabaseReadReplicaMapOutputWithContext(ctx context.Context) DatabaseReadReplicaMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseReadReplicaMapOutput)
}

type DatabaseReadReplicaOutput struct{ *pulumi.OutputState }

func (DatabaseReadReplicaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseReadReplica)(nil)).Elem()
}

func (o DatabaseReadReplicaOutput) ToDatabaseReadReplicaOutput() DatabaseReadReplicaOutput {
	return o
}

func (o DatabaseReadReplicaOutput) ToDatabaseReadReplicaOutputWithContext(ctx context.Context) DatabaseReadReplicaOutput {
	return o
}

// Creates a direct access endpoint to rdb replica.
func (o DatabaseReadReplicaOutput) DirectAccess() DatabaseReadReplicaDirectAccessPtrOutput {
	return o.ApplyT(func(v *DatabaseReadReplica) DatabaseReadReplicaDirectAccessPtrOutput { return v.DirectAccess }).(DatabaseReadReplicaDirectAccessPtrOutput)
}

// UUID of the rdb instance.
func (o DatabaseReadReplicaOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseReadReplica) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Create an endpoint in a private network.
func (o DatabaseReadReplicaOutput) PrivateNetwork() DatabaseReadReplicaPrivateNetworkPtrOutput {
	return o.ApplyT(func(v *DatabaseReadReplica) DatabaseReadReplicaPrivateNetworkPtrOutput { return v.PrivateNetwork }).(DatabaseReadReplicaPrivateNetworkPtrOutput)
}

// `region`) The region
// in which the Database read replica should be created.
func (o DatabaseReadReplicaOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseReadReplica) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

type DatabaseReadReplicaArrayOutput struct{ *pulumi.OutputState }

func (DatabaseReadReplicaArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatabaseReadReplica)(nil)).Elem()
}

func (o DatabaseReadReplicaArrayOutput) ToDatabaseReadReplicaArrayOutput() DatabaseReadReplicaArrayOutput {
	return o
}

func (o DatabaseReadReplicaArrayOutput) ToDatabaseReadReplicaArrayOutputWithContext(ctx context.Context) DatabaseReadReplicaArrayOutput {
	return o
}

func (o DatabaseReadReplicaArrayOutput) Index(i pulumi.IntInput) DatabaseReadReplicaOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DatabaseReadReplica {
		return vs[0].([]*DatabaseReadReplica)[vs[1].(int)]
	}).(DatabaseReadReplicaOutput)
}

type DatabaseReadReplicaMapOutput struct{ *pulumi.OutputState }

func (DatabaseReadReplicaMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatabaseReadReplica)(nil)).Elem()
}

func (o DatabaseReadReplicaMapOutput) ToDatabaseReadReplicaMapOutput() DatabaseReadReplicaMapOutput {
	return o
}

func (o DatabaseReadReplicaMapOutput) ToDatabaseReadReplicaMapOutputWithContext(ctx context.Context) DatabaseReadReplicaMapOutput {
	return o
}

func (o DatabaseReadReplicaMapOutput) MapIndex(k pulumi.StringInput) DatabaseReadReplicaOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DatabaseReadReplica {
		return vs[0].(map[string]*DatabaseReadReplica)[vs[1].(string)]
	}).(DatabaseReadReplicaOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseReadReplicaInput)(nil)).Elem(), &DatabaseReadReplica{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseReadReplicaArrayInput)(nil)).Elem(), DatabaseReadReplicaArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseReadReplicaMapInput)(nil)).Elem(), DatabaseReadReplicaMap{})
	pulumi.RegisterOutputType(DatabaseReadReplicaOutput{})
	pulumi.RegisterOutputType(DatabaseReadReplicaArrayOutput{})
	pulumi.RegisterOutputType(DatabaseReadReplicaMapOutput{})
}
