// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway
{
    [ScalewayResourceType("scaleway:index/databaseInstance:DatabaseInstance")]
    public partial class DatabaseInstance : Pulumi.CustomResource
    {
        /// <summary>
        /// Boolean to store logical backups in the same region as the database instance
        /// </summary>
        [Output("backupSameRegion")]
        public Output<bool> BackupSameRegion { get; private set; } = null!;

        /// <summary>
        /// Backup schedule frequency in hours
        /// </summary>
        [Output("backupScheduleFrequency")]
        public Output<int> BackupScheduleFrequency { get; private set; } = null!;

        /// <summary>
        /// Backup schedule retention in days
        /// </summary>
        [Output("backupScheduleRetention")]
        public Output<int> BackupScheduleRetention { get; private set; } = null!;

        /// <summary>
        /// Certificate of the database instance
        /// </summary>
        [Output("certificate")]
        public Output<string> Certificate { get; private set; } = null!;

        /// <summary>
        /// Disable automated backup for the database instance
        /// </summary>
        [Output("disableBackup")]
        public Output<bool?> DisableBackup { get; private set; } = null!;

        /// <summary>
        /// Endpoint IP of the database instance
        /// </summary>
        [Output("endpointIp")]
        public Output<string> EndpointIp { get; private set; } = null!;

        /// <summary>
        /// Endpoint port of the database instance
        /// </summary>
        [Output("endpointPort")]
        public Output<int> EndpointPort { get; private set; } = null!;

        /// <summary>
        /// Database's engine version id
        /// </summary>
        [Output("engine")]
        public Output<string> Engine { get; private set; } = null!;

        /// <summary>
        /// Enable or disable high availability for the database instance
        /// </summary>
        [Output("isHaCluster")]
        public Output<bool?> IsHaCluster { get; private set; } = null!;

        /// <summary>
        /// Load balancer of the database instance
        /// </summary>
        [Output("loadBalancers")]
        public Output<ImmutableArray<Outputs.DatabaseInstanceLoadBalancer>> LoadBalancers { get; private set; } = null!;

        /// <summary>
        /// Name of the database instance
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The type of database instance you want to create
        /// </summary>
        [Output("nodeType")]
        public Output<string> NodeType { get; private set; } = null!;

        /// <summary>
        /// The organization_id you want to attach the resource to
        /// </summary>
        [Output("organizationId")]
        public Output<string> OrganizationId { get; private set; } = null!;

        /// <summary>
        /// Password for the first user of the database instance
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// List of private network to expose your database instance
        /// </summary>
        [Output("privateNetwork")]
        public Output<Outputs.DatabaseInstancePrivateNetwork?> PrivateNetwork { get; private set; } = null!;

        /// <summary>
        /// The project_id you want to attach the resource to
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Read replicas of the database instance
        /// </summary>
        [Output("readReplicas")]
        public Output<ImmutableArray<Outputs.DatabaseInstanceReadReplica>> ReadReplicas { get; private set; } = null!;

        /// <summary>
        /// The region you want to attach the resource to
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Map of engine settings to be set.
        /// </summary>
        [Output("settings")]
        public Output<ImmutableDictionary<string, string>> Settings { get; private set; } = null!;

        /// <summary>
        /// List of tags ["tag1", "tag2", ...] attached to a database instance
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// Identifier for the first user of the database instance
        /// </summary>
        [Output("userName")]
        public Output<string?> UserName { get; private set; } = null!;

        /// <summary>
        /// Volume size (in GB) when volume_type is not lssd
        /// </summary>
        [Output("volumeSizeInGb")]
        public Output<int> VolumeSizeInGb { get; private set; } = null!;

        /// <summary>
        /// Type of volume where data are stored
        /// </summary>
        [Output("volumeType")]
        public Output<string?> VolumeType { get; private set; } = null!;


        /// <summary>
        /// Create a DatabaseInstance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DatabaseInstance(string name, DatabaseInstanceArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/databaseInstance:DatabaseInstance", name, args ?? new DatabaseInstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DatabaseInstance(string name, Input<string> id, DatabaseInstanceState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/databaseInstance:DatabaseInstance", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/pulumiverse/pulumi-scaleway/releases/download/${VERSION}",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DatabaseInstance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DatabaseInstance Get(string name, Input<string> id, DatabaseInstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new DatabaseInstance(name, id, state, options);
        }
    }

    public sealed class DatabaseInstanceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Boolean to store logical backups in the same region as the database instance
        /// </summary>
        [Input("backupSameRegion")]
        public Input<bool>? BackupSameRegion { get; set; }

        /// <summary>
        /// Backup schedule frequency in hours
        /// </summary>
        [Input("backupScheduleFrequency")]
        public Input<int>? BackupScheduleFrequency { get; set; }

        /// <summary>
        /// Backup schedule retention in days
        /// </summary>
        [Input("backupScheduleRetention")]
        public Input<int>? BackupScheduleRetention { get; set; }

        /// <summary>
        /// Disable automated backup for the database instance
        /// </summary>
        [Input("disableBackup")]
        public Input<bool>? DisableBackup { get; set; }

        /// <summary>
        /// Database's engine version id
        /// </summary>
        [Input("engine", required: true)]
        public Input<string> Engine { get; set; } = null!;

        /// <summary>
        /// Enable or disable high availability for the database instance
        /// </summary>
        [Input("isHaCluster")]
        public Input<bool>? IsHaCluster { get; set; }

        /// <summary>
        /// Name of the database instance
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The type of database instance you want to create
        /// </summary>
        [Input("nodeType", required: true)]
        public Input<string> NodeType { get; set; } = null!;

        /// <summary>
        /// Password for the first user of the database instance
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// List of private network to expose your database instance
        /// </summary>
        [Input("privateNetwork")]
        public Input<Inputs.DatabaseInstancePrivateNetworkArgs>? PrivateNetwork { get; set; }

        /// <summary>
        /// The project_id you want to attach the resource to
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The region you want to attach the resource to
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("settings")]
        private InputMap<string>? _settings;

        /// <summary>
        /// Map of engine settings to be set.
        /// </summary>
        public InputMap<string> Settings
        {
            get => _settings ?? (_settings = new InputMap<string>());
            set => _settings = value;
        }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// List of tags ["tag1", "tag2", ...] attached to a database instance
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Identifier for the first user of the database instance
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        /// <summary>
        /// Volume size (in GB) when volume_type is not lssd
        /// </summary>
        [Input("volumeSizeInGb")]
        public Input<int>? VolumeSizeInGb { get; set; }

        /// <summary>
        /// Type of volume where data are stored
        /// </summary>
        [Input("volumeType")]
        public Input<string>? VolumeType { get; set; }

        public DatabaseInstanceArgs()
        {
        }
    }

    public sealed class DatabaseInstanceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Boolean to store logical backups in the same region as the database instance
        /// </summary>
        [Input("backupSameRegion")]
        public Input<bool>? BackupSameRegion { get; set; }

        /// <summary>
        /// Backup schedule frequency in hours
        /// </summary>
        [Input("backupScheduleFrequency")]
        public Input<int>? BackupScheduleFrequency { get; set; }

        /// <summary>
        /// Backup schedule retention in days
        /// </summary>
        [Input("backupScheduleRetention")]
        public Input<int>? BackupScheduleRetention { get; set; }

        /// <summary>
        /// Certificate of the database instance
        /// </summary>
        [Input("certificate")]
        public Input<string>? Certificate { get; set; }

        /// <summary>
        /// Disable automated backup for the database instance
        /// </summary>
        [Input("disableBackup")]
        public Input<bool>? DisableBackup { get; set; }

        /// <summary>
        /// Endpoint IP of the database instance
        /// </summary>
        [Input("endpointIp")]
        public Input<string>? EndpointIp { get; set; }

        /// <summary>
        /// Endpoint port of the database instance
        /// </summary>
        [Input("endpointPort")]
        public Input<int>? EndpointPort { get; set; }

        /// <summary>
        /// Database's engine version id
        /// </summary>
        [Input("engine")]
        public Input<string>? Engine { get; set; }

        /// <summary>
        /// Enable or disable high availability for the database instance
        /// </summary>
        [Input("isHaCluster")]
        public Input<bool>? IsHaCluster { get; set; }

        [Input("loadBalancers")]
        private InputList<Inputs.DatabaseInstanceLoadBalancerGetArgs>? _loadBalancers;

        /// <summary>
        /// Load balancer of the database instance
        /// </summary>
        public InputList<Inputs.DatabaseInstanceLoadBalancerGetArgs> LoadBalancers
        {
            get => _loadBalancers ?? (_loadBalancers = new InputList<Inputs.DatabaseInstanceLoadBalancerGetArgs>());
            set => _loadBalancers = value;
        }

        /// <summary>
        /// Name of the database instance
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The type of database instance you want to create
        /// </summary>
        [Input("nodeType")]
        public Input<string>? NodeType { get; set; }

        /// <summary>
        /// The organization_id you want to attach the resource to
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        /// <summary>
        /// Password for the first user of the database instance
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// List of private network to expose your database instance
        /// </summary>
        [Input("privateNetwork")]
        public Input<Inputs.DatabaseInstancePrivateNetworkGetArgs>? PrivateNetwork { get; set; }

        /// <summary>
        /// The project_id you want to attach the resource to
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("readReplicas")]
        private InputList<Inputs.DatabaseInstanceReadReplicaGetArgs>? _readReplicas;

        /// <summary>
        /// Read replicas of the database instance
        /// </summary>
        public InputList<Inputs.DatabaseInstanceReadReplicaGetArgs> ReadReplicas
        {
            get => _readReplicas ?? (_readReplicas = new InputList<Inputs.DatabaseInstanceReadReplicaGetArgs>());
            set => _readReplicas = value;
        }

        /// <summary>
        /// The region you want to attach the resource to
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("settings")]
        private InputMap<string>? _settings;

        /// <summary>
        /// Map of engine settings to be set.
        /// </summary>
        public InputMap<string> Settings
        {
            get => _settings ?? (_settings = new InputMap<string>());
            set => _settings = value;
        }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// List of tags ["tag1", "tag2", ...] attached to a database instance
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Identifier for the first user of the database instance
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        /// <summary>
        /// Volume size (in GB) when volume_type is not lssd
        /// </summary>
        [Input("volumeSizeInGb")]
        public Input<int>? VolumeSizeInGb { get; set; }

        /// <summary>
        /// Type of volume where data are stored
        /// </summary>
        [Input("volumeType")]
        public Input<string>? VolumeType { get; set; }

        public DatabaseInstanceState()
        {
        }
    }
}
