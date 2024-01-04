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

// Creates and manages Scaleway Database Instances.
// For more information, see [the documentation](https://developers.scaleway.com/en/products/rdb/api).
//
// ## Examples
//
// ### Example Basic
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
//			_, err := scaleway.NewDatabaseInstance(ctx, "main", &scaleway.DatabaseInstanceArgs{
//				DisableBackup: pulumi.Bool(true),
//				Engine:        pulumi.String("PostgreSQL-11"),
//				IsHaCluster:   pulumi.Bool(true),
//				NodeType:      pulumi.String("DB-DEV-S"),
//				Password:      pulumi.String("thiZ_is_v&ry_s3cret"),
//				UserName:      pulumi.String("my_initial_user"),
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
// ### Example with Settings
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
//			_, err := scaleway.NewDatabaseInstance(ctx, "main", &scaleway.DatabaseInstanceArgs{
//				DisableBackup: pulumi.Bool(true),
//				Engine:        pulumi.String("MySQL-8"),
//				InitSettings: pulumi.StringMap{
//					"lower_case_table_names": pulumi.String("1"),
//				},
//				NodeType: pulumi.String("db-dev-s"),
//				Password: pulumi.String("thiZ_is_v&ry_s3cret"),
//				Settings: pulumi.StringMap{
//					"max_connections": pulumi.String("350"),
//				},
//				UserName: pulumi.String("my_initial_user"),
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
// ### Example with backup schedule
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
//			_, err := scaleway.NewDatabaseInstance(ctx, "main", &scaleway.DatabaseInstanceArgs{
//				BackupScheduleFrequency: pulumi.Int(24),
//				BackupScheduleRetention: pulumi.Int(7),
//				DisableBackup:           pulumi.Bool(false),
//				Engine:                  pulumi.String("PostgreSQL-11"),
//				IsHaCluster:             pulumi.Bool(true),
//				NodeType:                pulumi.String("DB-DEV-S"),
//				Password:                pulumi.String("thiZ_is_v&ry_s3cret"),
//				UserName:                pulumi.String("my_initial_user"),
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
// ## Limitations
//
// The Managed Database product is only compliant with the private network in the default availability zone (AZ).
// i.e. `fr-par-1`, `nl-ams-1`, `pl-waw-1`. To learn more, read our
// section [How to connect a PostgreSQL and MySQL Database Instance to a Private Network](https://www.scaleway.com/en/docs/managed-databases/postgresql-and-mysql/how-to/connect-database-private-network/)
//
// ## Import
//
// Database Instance can be imported using the `{region}/{id}`, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:index/databaseInstance:DatabaseInstance rdb01 fr-par/11111111-1111-1111-1111-111111111111
//
// ```
type DatabaseInstance struct {
	pulumi.CustomResourceState

	// Boolean to store logical backups in the same region as the database instance.
	BackupSameRegion pulumi.BoolOutput `pulumi:"backupSameRegion"`
	// Backup schedule frequency in hours.
	BackupScheduleFrequency pulumi.IntOutput `pulumi:"backupScheduleFrequency"`
	// Backup schedule retention in days.
	BackupScheduleRetention pulumi.IntOutput `pulumi:"backupScheduleRetention"`
	// Certificate of the database instance.
	Certificate pulumi.StringOutput `pulumi:"certificate"`
	// Disable automated backup for the database instance.
	DisableBackup pulumi.BoolPtrOutput `pulumi:"disableBackup"`
	// (Deprecated) The IP of the Database Instance.
	//
	// Deprecated: Please use the private_network or the load_balancer attribute
	EndpointIp pulumi.StringOutput `pulumi:"endpointIp"`
	// (Deprecated) The port of the Database Instance.
	EndpointPort pulumi.IntOutput `pulumi:"endpointPort"`
	// Database Instance's engine version (e.g. `PostgreSQL-11`).
	//
	// > **Important:** Updates to `engine` will recreate the Database Instance.
	Engine pulumi.StringOutput `pulumi:"engine"`
	// Map of engine settings to be set at database initialisation.
	//
	// > **Important:** Updates to `initSettings` will recreate the Database Instance.
	//
	// Please consult the [GoDoc](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@v1.0.0-beta.9/api/rdb/v1#EngineVersion) to list all available `settings` and `initSettings` for the `nodeType` of your convenience.
	InitSettings pulumi.StringMapOutput `pulumi:"initSettings"`
	// Enable or disable high availability for the database instance.
	//
	// > **Important:** Updates to `isHaCluster` will recreate the Database Instance.
	IsHaCluster pulumi.BoolPtrOutput `pulumi:"isHaCluster"`
	// List of load balancer endpoints of the database instance. A load-balancer endpoint will be set by default if no private network is.
	// This block must be defined if you want a public endpoint in addition to your private endpoint.
	LoadBalancers DatabaseInstanceLoadBalancerArrayOutput `pulumi:"loadBalancers"`
	// The name of the Database Instance.
	Name pulumi.StringOutput `pulumi:"name"`
	// The type of database instance you want to create (e.g. `db-dev-s`).
	//
	// > **Important:** Updates to `nodeType` will upgrade the Database Instance to the desired `nodeType` without any
	// interruption. Keep in mind that you cannot downgrade a Database Instance.
	//
	// > **Important:** Once your instance reaches `diskFull` status, if you are using `lssd` storage, you should upgrade the node_type,
	// and if you are using `bssd` storage, you should increase the volume size before making any other change to your instance.
	NodeType pulumi.StringOutput `pulumi:"nodeType"`
	// The organization ID the Database Instance is associated with.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// Password for the first user of the database instance.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// List of private networks endpoints of the database instance.
	PrivateNetwork DatabaseInstancePrivateNetworkPtrOutput `pulumi:"privateNetwork"`
	// `projectId`) The ID of the project the Database
	// Instance is associated with.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// List of read replicas of the database instance.
	ReadReplicas DatabaseInstanceReadReplicaArrayOutput `pulumi:"readReplicas"`
	// `region`) The region
	// in which the Database Instance should be created.
	Region pulumi.StringOutput `pulumi:"region"`
	// Map of engine settings to be set. Using this option will override default config.
	Settings pulumi.StringMapOutput `pulumi:"settings"`
	// The tags associated with the Database Instance.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// Identifier for the first user of the database instance.
	//
	// > **Important:** Updates to `userName` will recreate the Database Instance.
	UserName pulumi.StringOutput `pulumi:"userName"`
	// Volume size (in GB) when `volumeType` is set to `bssd`.
	//
	// > **Important:** Once your instance reaches `diskFull` status, you should increase the volume size before making any other change to your instance.
	VolumeSizeInGb pulumi.IntOutput `pulumi:"volumeSizeInGb"`
	// Type of volume where data are stored (`bssd` or `lssd`).
	VolumeType pulumi.StringPtrOutput `pulumi:"volumeType"`
}

// NewDatabaseInstance registers a new resource with the given unique name, arguments, and options.
func NewDatabaseInstance(ctx *pulumi.Context,
	name string, args *DatabaseInstanceArgs, opts ...pulumi.ResourceOption) (*DatabaseInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Engine == nil {
		return nil, errors.New("invalid value for required argument 'Engine'")
	}
	if args.NodeType == nil {
		return nil, errors.New("invalid value for required argument 'NodeType'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DatabaseInstance
	err := ctx.RegisterResource("scaleway:index/databaseInstance:DatabaseInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabaseInstance gets an existing DatabaseInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabaseInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseInstanceState, opts ...pulumi.ResourceOption) (*DatabaseInstance, error) {
	var resource DatabaseInstance
	err := ctx.ReadResource("scaleway:index/databaseInstance:DatabaseInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabaseInstance resources.
type databaseInstanceState struct {
	// Boolean to store logical backups in the same region as the database instance.
	BackupSameRegion *bool `pulumi:"backupSameRegion"`
	// Backup schedule frequency in hours.
	BackupScheduleFrequency *int `pulumi:"backupScheduleFrequency"`
	// Backup schedule retention in days.
	BackupScheduleRetention *int `pulumi:"backupScheduleRetention"`
	// Certificate of the database instance.
	Certificate *string `pulumi:"certificate"`
	// Disable automated backup for the database instance.
	DisableBackup *bool `pulumi:"disableBackup"`
	// (Deprecated) The IP of the Database Instance.
	//
	// Deprecated: Please use the private_network or the load_balancer attribute
	EndpointIp *string `pulumi:"endpointIp"`
	// (Deprecated) The port of the Database Instance.
	EndpointPort *int `pulumi:"endpointPort"`
	// Database Instance's engine version (e.g. `PostgreSQL-11`).
	//
	// > **Important:** Updates to `engine` will recreate the Database Instance.
	Engine *string `pulumi:"engine"`
	// Map of engine settings to be set at database initialisation.
	//
	// > **Important:** Updates to `initSettings` will recreate the Database Instance.
	//
	// Please consult the [GoDoc](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@v1.0.0-beta.9/api/rdb/v1#EngineVersion) to list all available `settings` and `initSettings` for the `nodeType` of your convenience.
	InitSettings map[string]string `pulumi:"initSettings"`
	// Enable or disable high availability for the database instance.
	//
	// > **Important:** Updates to `isHaCluster` will recreate the Database Instance.
	IsHaCluster *bool `pulumi:"isHaCluster"`
	// List of load balancer endpoints of the database instance. A load-balancer endpoint will be set by default if no private network is.
	// This block must be defined if you want a public endpoint in addition to your private endpoint.
	LoadBalancers []DatabaseInstanceLoadBalancer `pulumi:"loadBalancers"`
	// The name of the Database Instance.
	Name *string `pulumi:"name"`
	// The type of database instance you want to create (e.g. `db-dev-s`).
	//
	// > **Important:** Updates to `nodeType` will upgrade the Database Instance to the desired `nodeType` without any
	// interruption. Keep in mind that you cannot downgrade a Database Instance.
	//
	// > **Important:** Once your instance reaches `diskFull` status, if you are using `lssd` storage, you should upgrade the node_type,
	// and if you are using `bssd` storage, you should increase the volume size before making any other change to your instance.
	NodeType *string `pulumi:"nodeType"`
	// The organization ID the Database Instance is associated with.
	OrganizationId *string `pulumi:"organizationId"`
	// Password for the first user of the database instance.
	Password *string `pulumi:"password"`
	// List of private networks endpoints of the database instance.
	PrivateNetwork *DatabaseInstancePrivateNetwork `pulumi:"privateNetwork"`
	// `projectId`) The ID of the project the Database
	// Instance is associated with.
	ProjectId *string `pulumi:"projectId"`
	// List of read replicas of the database instance.
	ReadReplicas []DatabaseInstanceReadReplica `pulumi:"readReplicas"`
	// `region`) The region
	// in which the Database Instance should be created.
	Region *string `pulumi:"region"`
	// Map of engine settings to be set. Using this option will override default config.
	Settings map[string]string `pulumi:"settings"`
	// The tags associated with the Database Instance.
	Tags []string `pulumi:"tags"`
	// Identifier for the first user of the database instance.
	//
	// > **Important:** Updates to `userName` will recreate the Database Instance.
	UserName *string `pulumi:"userName"`
	// Volume size (in GB) when `volumeType` is set to `bssd`.
	//
	// > **Important:** Once your instance reaches `diskFull` status, you should increase the volume size before making any other change to your instance.
	VolumeSizeInGb *int `pulumi:"volumeSizeInGb"`
	// Type of volume where data are stored (`bssd` or `lssd`).
	VolumeType *string `pulumi:"volumeType"`
}

type DatabaseInstanceState struct {
	// Boolean to store logical backups in the same region as the database instance.
	BackupSameRegion pulumi.BoolPtrInput
	// Backup schedule frequency in hours.
	BackupScheduleFrequency pulumi.IntPtrInput
	// Backup schedule retention in days.
	BackupScheduleRetention pulumi.IntPtrInput
	// Certificate of the database instance.
	Certificate pulumi.StringPtrInput
	// Disable automated backup for the database instance.
	DisableBackup pulumi.BoolPtrInput
	// (Deprecated) The IP of the Database Instance.
	//
	// Deprecated: Please use the private_network or the load_balancer attribute
	EndpointIp pulumi.StringPtrInput
	// (Deprecated) The port of the Database Instance.
	EndpointPort pulumi.IntPtrInput
	// Database Instance's engine version (e.g. `PostgreSQL-11`).
	//
	// > **Important:** Updates to `engine` will recreate the Database Instance.
	Engine pulumi.StringPtrInput
	// Map of engine settings to be set at database initialisation.
	//
	// > **Important:** Updates to `initSettings` will recreate the Database Instance.
	//
	// Please consult the [GoDoc](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@v1.0.0-beta.9/api/rdb/v1#EngineVersion) to list all available `settings` and `initSettings` for the `nodeType` of your convenience.
	InitSettings pulumi.StringMapInput
	// Enable or disable high availability for the database instance.
	//
	// > **Important:** Updates to `isHaCluster` will recreate the Database Instance.
	IsHaCluster pulumi.BoolPtrInput
	// List of load balancer endpoints of the database instance. A load-balancer endpoint will be set by default if no private network is.
	// This block must be defined if you want a public endpoint in addition to your private endpoint.
	LoadBalancers DatabaseInstanceLoadBalancerArrayInput
	// The name of the Database Instance.
	Name pulumi.StringPtrInput
	// The type of database instance you want to create (e.g. `db-dev-s`).
	//
	// > **Important:** Updates to `nodeType` will upgrade the Database Instance to the desired `nodeType` without any
	// interruption. Keep in mind that you cannot downgrade a Database Instance.
	//
	// > **Important:** Once your instance reaches `diskFull` status, if you are using `lssd` storage, you should upgrade the node_type,
	// and if you are using `bssd` storage, you should increase the volume size before making any other change to your instance.
	NodeType pulumi.StringPtrInput
	// The organization ID the Database Instance is associated with.
	OrganizationId pulumi.StringPtrInput
	// Password for the first user of the database instance.
	Password pulumi.StringPtrInput
	// List of private networks endpoints of the database instance.
	PrivateNetwork DatabaseInstancePrivateNetworkPtrInput
	// `projectId`) The ID of the project the Database
	// Instance is associated with.
	ProjectId pulumi.StringPtrInput
	// List of read replicas of the database instance.
	ReadReplicas DatabaseInstanceReadReplicaArrayInput
	// `region`) The region
	// in which the Database Instance should be created.
	Region pulumi.StringPtrInput
	// Map of engine settings to be set. Using this option will override default config.
	Settings pulumi.StringMapInput
	// The tags associated with the Database Instance.
	Tags pulumi.StringArrayInput
	// Identifier for the first user of the database instance.
	//
	// > **Important:** Updates to `userName` will recreate the Database Instance.
	UserName pulumi.StringPtrInput
	// Volume size (in GB) when `volumeType` is set to `bssd`.
	//
	// > **Important:** Once your instance reaches `diskFull` status, you should increase the volume size before making any other change to your instance.
	VolumeSizeInGb pulumi.IntPtrInput
	// Type of volume where data are stored (`bssd` or `lssd`).
	VolumeType pulumi.StringPtrInput
}

func (DatabaseInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseInstanceState)(nil)).Elem()
}

type databaseInstanceArgs struct {
	// Boolean to store logical backups in the same region as the database instance.
	BackupSameRegion *bool `pulumi:"backupSameRegion"`
	// Backup schedule frequency in hours.
	BackupScheduleFrequency *int `pulumi:"backupScheduleFrequency"`
	// Backup schedule retention in days.
	BackupScheduleRetention *int `pulumi:"backupScheduleRetention"`
	// Disable automated backup for the database instance.
	DisableBackup *bool `pulumi:"disableBackup"`
	// Database Instance's engine version (e.g. `PostgreSQL-11`).
	//
	// > **Important:** Updates to `engine` will recreate the Database Instance.
	Engine string `pulumi:"engine"`
	// Map of engine settings to be set at database initialisation.
	//
	// > **Important:** Updates to `initSettings` will recreate the Database Instance.
	//
	// Please consult the [GoDoc](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@v1.0.0-beta.9/api/rdb/v1#EngineVersion) to list all available `settings` and `initSettings` for the `nodeType` of your convenience.
	InitSettings map[string]string `pulumi:"initSettings"`
	// Enable or disable high availability for the database instance.
	//
	// > **Important:** Updates to `isHaCluster` will recreate the Database Instance.
	IsHaCluster *bool `pulumi:"isHaCluster"`
	// List of load balancer endpoints of the database instance. A load-balancer endpoint will be set by default if no private network is.
	// This block must be defined if you want a public endpoint in addition to your private endpoint.
	LoadBalancers []DatabaseInstanceLoadBalancer `pulumi:"loadBalancers"`
	// The name of the Database Instance.
	Name *string `pulumi:"name"`
	// The type of database instance you want to create (e.g. `db-dev-s`).
	//
	// > **Important:** Updates to `nodeType` will upgrade the Database Instance to the desired `nodeType` without any
	// interruption. Keep in mind that you cannot downgrade a Database Instance.
	//
	// > **Important:** Once your instance reaches `diskFull` status, if you are using `lssd` storage, you should upgrade the node_type,
	// and if you are using `bssd` storage, you should increase the volume size before making any other change to your instance.
	NodeType string `pulumi:"nodeType"`
	// Password for the first user of the database instance.
	Password *string `pulumi:"password"`
	// List of private networks endpoints of the database instance.
	PrivateNetwork *DatabaseInstancePrivateNetwork `pulumi:"privateNetwork"`
	// `projectId`) The ID of the project the Database
	// Instance is associated with.
	ProjectId *string `pulumi:"projectId"`
	// `region`) The region
	// in which the Database Instance should be created.
	Region *string `pulumi:"region"`
	// Map of engine settings to be set. Using this option will override default config.
	Settings map[string]string `pulumi:"settings"`
	// The tags associated with the Database Instance.
	Tags []string `pulumi:"tags"`
	// Identifier for the first user of the database instance.
	//
	// > **Important:** Updates to `userName` will recreate the Database Instance.
	UserName *string `pulumi:"userName"`
	// Volume size (in GB) when `volumeType` is set to `bssd`.
	//
	// > **Important:** Once your instance reaches `diskFull` status, you should increase the volume size before making any other change to your instance.
	VolumeSizeInGb *int `pulumi:"volumeSizeInGb"`
	// Type of volume where data are stored (`bssd` or `lssd`).
	VolumeType *string `pulumi:"volumeType"`
}

// The set of arguments for constructing a DatabaseInstance resource.
type DatabaseInstanceArgs struct {
	// Boolean to store logical backups in the same region as the database instance.
	BackupSameRegion pulumi.BoolPtrInput
	// Backup schedule frequency in hours.
	BackupScheduleFrequency pulumi.IntPtrInput
	// Backup schedule retention in days.
	BackupScheduleRetention pulumi.IntPtrInput
	// Disable automated backup for the database instance.
	DisableBackup pulumi.BoolPtrInput
	// Database Instance's engine version (e.g. `PostgreSQL-11`).
	//
	// > **Important:** Updates to `engine` will recreate the Database Instance.
	Engine pulumi.StringInput
	// Map of engine settings to be set at database initialisation.
	//
	// > **Important:** Updates to `initSettings` will recreate the Database Instance.
	//
	// Please consult the [GoDoc](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@v1.0.0-beta.9/api/rdb/v1#EngineVersion) to list all available `settings` and `initSettings` for the `nodeType` of your convenience.
	InitSettings pulumi.StringMapInput
	// Enable or disable high availability for the database instance.
	//
	// > **Important:** Updates to `isHaCluster` will recreate the Database Instance.
	IsHaCluster pulumi.BoolPtrInput
	// List of load balancer endpoints of the database instance. A load-balancer endpoint will be set by default if no private network is.
	// This block must be defined if you want a public endpoint in addition to your private endpoint.
	LoadBalancers DatabaseInstanceLoadBalancerArrayInput
	// The name of the Database Instance.
	Name pulumi.StringPtrInput
	// The type of database instance you want to create (e.g. `db-dev-s`).
	//
	// > **Important:** Updates to `nodeType` will upgrade the Database Instance to the desired `nodeType` without any
	// interruption. Keep in mind that you cannot downgrade a Database Instance.
	//
	// > **Important:** Once your instance reaches `diskFull` status, if you are using `lssd` storage, you should upgrade the node_type,
	// and if you are using `bssd` storage, you should increase the volume size before making any other change to your instance.
	NodeType pulumi.StringInput
	// Password for the first user of the database instance.
	Password pulumi.StringPtrInput
	// List of private networks endpoints of the database instance.
	PrivateNetwork DatabaseInstancePrivateNetworkPtrInput
	// `projectId`) The ID of the project the Database
	// Instance is associated with.
	ProjectId pulumi.StringPtrInput
	// `region`) The region
	// in which the Database Instance should be created.
	Region pulumi.StringPtrInput
	// Map of engine settings to be set. Using this option will override default config.
	Settings pulumi.StringMapInput
	// The tags associated with the Database Instance.
	Tags pulumi.StringArrayInput
	// Identifier for the first user of the database instance.
	//
	// > **Important:** Updates to `userName` will recreate the Database Instance.
	UserName pulumi.StringPtrInput
	// Volume size (in GB) when `volumeType` is set to `bssd`.
	//
	// > **Important:** Once your instance reaches `diskFull` status, you should increase the volume size before making any other change to your instance.
	VolumeSizeInGb pulumi.IntPtrInput
	// Type of volume where data are stored (`bssd` or `lssd`).
	VolumeType pulumi.StringPtrInput
}

func (DatabaseInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseInstanceArgs)(nil)).Elem()
}

type DatabaseInstanceInput interface {
	pulumi.Input

	ToDatabaseInstanceOutput() DatabaseInstanceOutput
	ToDatabaseInstanceOutputWithContext(ctx context.Context) DatabaseInstanceOutput
}

func (*DatabaseInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseInstance)(nil)).Elem()
}

func (i *DatabaseInstance) ToDatabaseInstanceOutput() DatabaseInstanceOutput {
	return i.ToDatabaseInstanceOutputWithContext(context.Background())
}

func (i *DatabaseInstance) ToDatabaseInstanceOutputWithContext(ctx context.Context) DatabaseInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseInstanceOutput)
}

// DatabaseInstanceArrayInput is an input type that accepts DatabaseInstanceArray and DatabaseInstanceArrayOutput values.
// You can construct a concrete instance of `DatabaseInstanceArrayInput` via:
//
//	DatabaseInstanceArray{ DatabaseInstanceArgs{...} }
type DatabaseInstanceArrayInput interface {
	pulumi.Input

	ToDatabaseInstanceArrayOutput() DatabaseInstanceArrayOutput
	ToDatabaseInstanceArrayOutputWithContext(context.Context) DatabaseInstanceArrayOutput
}

type DatabaseInstanceArray []DatabaseInstanceInput

func (DatabaseInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatabaseInstance)(nil)).Elem()
}

func (i DatabaseInstanceArray) ToDatabaseInstanceArrayOutput() DatabaseInstanceArrayOutput {
	return i.ToDatabaseInstanceArrayOutputWithContext(context.Background())
}

func (i DatabaseInstanceArray) ToDatabaseInstanceArrayOutputWithContext(ctx context.Context) DatabaseInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseInstanceArrayOutput)
}

// DatabaseInstanceMapInput is an input type that accepts DatabaseInstanceMap and DatabaseInstanceMapOutput values.
// You can construct a concrete instance of `DatabaseInstanceMapInput` via:
//
//	DatabaseInstanceMap{ "key": DatabaseInstanceArgs{...} }
type DatabaseInstanceMapInput interface {
	pulumi.Input

	ToDatabaseInstanceMapOutput() DatabaseInstanceMapOutput
	ToDatabaseInstanceMapOutputWithContext(context.Context) DatabaseInstanceMapOutput
}

type DatabaseInstanceMap map[string]DatabaseInstanceInput

func (DatabaseInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatabaseInstance)(nil)).Elem()
}

func (i DatabaseInstanceMap) ToDatabaseInstanceMapOutput() DatabaseInstanceMapOutput {
	return i.ToDatabaseInstanceMapOutputWithContext(context.Background())
}

func (i DatabaseInstanceMap) ToDatabaseInstanceMapOutputWithContext(ctx context.Context) DatabaseInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseInstanceMapOutput)
}

type DatabaseInstanceOutput struct{ *pulumi.OutputState }

func (DatabaseInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseInstance)(nil)).Elem()
}

func (o DatabaseInstanceOutput) ToDatabaseInstanceOutput() DatabaseInstanceOutput {
	return o
}

func (o DatabaseInstanceOutput) ToDatabaseInstanceOutputWithContext(ctx context.Context) DatabaseInstanceOutput {
	return o
}

// Boolean to store logical backups in the same region as the database instance.
func (o DatabaseInstanceOutput) BackupSameRegion() pulumi.BoolOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.BoolOutput { return v.BackupSameRegion }).(pulumi.BoolOutput)
}

// Backup schedule frequency in hours.
func (o DatabaseInstanceOutput) BackupScheduleFrequency() pulumi.IntOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.IntOutput { return v.BackupScheduleFrequency }).(pulumi.IntOutput)
}

// Backup schedule retention in days.
func (o DatabaseInstanceOutput) BackupScheduleRetention() pulumi.IntOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.IntOutput { return v.BackupScheduleRetention }).(pulumi.IntOutput)
}

// Certificate of the database instance.
func (o DatabaseInstanceOutput) Certificate() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.StringOutput { return v.Certificate }).(pulumi.StringOutput)
}

// Disable automated backup for the database instance.
func (o DatabaseInstanceOutput) DisableBackup() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.BoolPtrOutput { return v.DisableBackup }).(pulumi.BoolPtrOutput)
}

// (Deprecated) The IP of the Database Instance.
//
// Deprecated: Please use the private_network or the load_balancer attribute
func (o DatabaseInstanceOutput) EndpointIp() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.StringOutput { return v.EndpointIp }).(pulumi.StringOutput)
}

// (Deprecated) The port of the Database Instance.
func (o DatabaseInstanceOutput) EndpointPort() pulumi.IntOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.IntOutput { return v.EndpointPort }).(pulumi.IntOutput)
}

// Database Instance's engine version (e.g. `PostgreSQL-11`).
//
// > **Important:** Updates to `engine` will recreate the Database Instance.
func (o DatabaseInstanceOutput) Engine() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.StringOutput { return v.Engine }).(pulumi.StringOutput)
}

// Map of engine settings to be set at database initialisation.
//
// > **Important:** Updates to `initSettings` will recreate the Database Instance.
//
// Please consult the [GoDoc](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@v1.0.0-beta.9/api/rdb/v1#EngineVersion) to list all available `settings` and `initSettings` for the `nodeType` of your convenience.
func (o DatabaseInstanceOutput) InitSettings() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.StringMapOutput { return v.InitSettings }).(pulumi.StringMapOutput)
}

// Enable or disable high availability for the database instance.
//
// > **Important:** Updates to `isHaCluster` will recreate the Database Instance.
func (o DatabaseInstanceOutput) IsHaCluster() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.BoolPtrOutput { return v.IsHaCluster }).(pulumi.BoolPtrOutput)
}

// List of load balancer endpoints of the database instance. A load-balancer endpoint will be set by default if no private network is.
// This block must be defined if you want a public endpoint in addition to your private endpoint.
func (o DatabaseInstanceOutput) LoadBalancers() DatabaseInstanceLoadBalancerArrayOutput {
	return o.ApplyT(func(v *DatabaseInstance) DatabaseInstanceLoadBalancerArrayOutput { return v.LoadBalancers }).(DatabaseInstanceLoadBalancerArrayOutput)
}

// The name of the Database Instance.
func (o DatabaseInstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The type of database instance you want to create (e.g. `db-dev-s`).
//
// > **Important:** Updates to `nodeType` will upgrade the Database Instance to the desired `nodeType` without any
// interruption. Keep in mind that you cannot downgrade a Database Instance.
//
// > **Important:** Once your instance reaches `diskFull` status, if you are using `lssd` storage, you should upgrade the node_type,
// and if you are using `bssd` storage, you should increase the volume size before making any other change to your instance.
func (o DatabaseInstanceOutput) NodeType() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.StringOutput { return v.NodeType }).(pulumi.StringOutput)
}

// The organization ID the Database Instance is associated with.
func (o DatabaseInstanceOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// Password for the first user of the database instance.
func (o DatabaseInstanceOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// List of private networks endpoints of the database instance.
func (o DatabaseInstanceOutput) PrivateNetwork() DatabaseInstancePrivateNetworkPtrOutput {
	return o.ApplyT(func(v *DatabaseInstance) DatabaseInstancePrivateNetworkPtrOutput { return v.PrivateNetwork }).(DatabaseInstancePrivateNetworkPtrOutput)
}

// `projectId`) The ID of the project the Database
// Instance is associated with.
func (o DatabaseInstanceOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// List of read replicas of the database instance.
func (o DatabaseInstanceOutput) ReadReplicas() DatabaseInstanceReadReplicaArrayOutput {
	return o.ApplyT(func(v *DatabaseInstance) DatabaseInstanceReadReplicaArrayOutput { return v.ReadReplicas }).(DatabaseInstanceReadReplicaArrayOutput)
}

// `region`) The region
// in which the Database Instance should be created.
func (o DatabaseInstanceOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Map of engine settings to be set. Using this option will override default config.
func (o DatabaseInstanceOutput) Settings() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.StringMapOutput { return v.Settings }).(pulumi.StringMapOutput)
}

// The tags associated with the Database Instance.
func (o DatabaseInstanceOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// Identifier for the first user of the database instance.
//
// > **Important:** Updates to `userName` will recreate the Database Instance.
func (o DatabaseInstanceOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.StringOutput { return v.UserName }).(pulumi.StringOutput)
}

// Volume size (in GB) when `volumeType` is set to `bssd`.
//
// > **Important:** Once your instance reaches `diskFull` status, you should increase the volume size before making any other change to your instance.
func (o DatabaseInstanceOutput) VolumeSizeInGb() pulumi.IntOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.IntOutput { return v.VolumeSizeInGb }).(pulumi.IntOutput)
}

// Type of volume where data are stored (`bssd` or `lssd`).
func (o DatabaseInstanceOutput) VolumeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.StringPtrOutput { return v.VolumeType }).(pulumi.StringPtrOutput)
}

type DatabaseInstanceArrayOutput struct{ *pulumi.OutputState }

func (DatabaseInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatabaseInstance)(nil)).Elem()
}

func (o DatabaseInstanceArrayOutput) ToDatabaseInstanceArrayOutput() DatabaseInstanceArrayOutput {
	return o
}

func (o DatabaseInstanceArrayOutput) ToDatabaseInstanceArrayOutputWithContext(ctx context.Context) DatabaseInstanceArrayOutput {
	return o
}

func (o DatabaseInstanceArrayOutput) Index(i pulumi.IntInput) DatabaseInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DatabaseInstance {
		return vs[0].([]*DatabaseInstance)[vs[1].(int)]
	}).(DatabaseInstanceOutput)
}

type DatabaseInstanceMapOutput struct{ *pulumi.OutputState }

func (DatabaseInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatabaseInstance)(nil)).Elem()
}

func (o DatabaseInstanceMapOutput) ToDatabaseInstanceMapOutput() DatabaseInstanceMapOutput {
	return o
}

func (o DatabaseInstanceMapOutput) ToDatabaseInstanceMapOutputWithContext(ctx context.Context) DatabaseInstanceMapOutput {
	return o
}

func (o DatabaseInstanceMapOutput) MapIndex(k pulumi.StringInput) DatabaseInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DatabaseInstance {
		return vs[0].(map[string]*DatabaseInstance)[vs[1].(string)]
	}).(DatabaseInstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseInstanceInput)(nil)).Elem(), &DatabaseInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseInstanceArrayInput)(nil)).Elem(), DatabaseInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseInstanceMapInput)(nil)).Elem(), DatabaseInstanceMap{})
	pulumi.RegisterOutputType(DatabaseInstanceOutput{})
	pulumi.RegisterOutputType(DatabaseInstanceArrayOutput{})
	pulumi.RegisterOutputType(DatabaseInstanceMapOutput{})
}
