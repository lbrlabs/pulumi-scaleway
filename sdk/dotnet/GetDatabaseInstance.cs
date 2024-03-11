// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway
{
    public static class GetDatabaseInstance
    {
        /// <summary>
        /// Gets information about an RDB instance. For further information see our [developers website](https://developers.scaleway.com/en/products/rdb/api/#database-instance)
        /// </summary>
        public static Task<GetDatabaseInstanceResult> InvokeAsync(GetDatabaseInstanceArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDatabaseInstanceResult>("scaleway:index/getDatabaseInstance:getDatabaseInstance", args ?? new GetDatabaseInstanceArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about an RDB instance. For further information see our [developers website](https://developers.scaleway.com/en/products/rdb/api/#database-instance)
        /// </summary>
        public static Output<GetDatabaseInstanceResult> Invoke(GetDatabaseInstanceInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDatabaseInstanceResult>("scaleway:index/getDatabaseInstance:getDatabaseInstance", args ?? new GetDatabaseInstanceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDatabaseInstanceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The RDB instance ID.
        /// Only one of `name` and `instance_id` should be specified.
        /// </summary>
        [Input("instanceId")]
        public string? InstanceId { get; set; }

        /// <summary>
        /// The name of the RDB instance.
        /// Only one of `name` and `instance_id` should be specified.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The ID of the project the RDB instance is in. Can be used to filter instances when using `name`.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        /// <summary>
        /// `region`) The region in which the RDB instance exists.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        public GetDatabaseInstanceArgs()
        {
        }
        public static new GetDatabaseInstanceArgs Empty => new GetDatabaseInstanceArgs();
    }

    public sealed class GetDatabaseInstanceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The RDB instance ID.
        /// Only one of `name` and `instance_id` should be specified.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The name of the RDB instance.
        /// Only one of `name` and `instance_id` should be specified.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project the RDB instance is in. Can be used to filter instances when using `name`.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`) The region in which the RDB instance exists.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public GetDatabaseInstanceInvokeArgs()
        {
        }
        public static new GetDatabaseInstanceInvokeArgs Empty => new GetDatabaseInstanceInvokeArgs();
    }


    [OutputType]
    public sealed class GetDatabaseInstanceResult
    {
        public readonly bool BackupSameRegion;
        public readonly int BackupScheduleFrequency;
        public readonly int BackupScheduleRetention;
        public readonly string Certificate;
        public readonly bool DisableBackup;
        public readonly string EndpointIp;
        public readonly int EndpointPort;
        public readonly string Engine;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableDictionary<string, string> InitSettings;
        public readonly string? InstanceId;
        public readonly bool IsHaCluster;
        public readonly ImmutableArray<Outputs.GetDatabaseInstanceLoadBalancerResult> LoadBalancers;
        public readonly string? Name;
        public readonly string NodeType;
        public readonly string OrganizationId;
        public readonly string Password;
        public readonly ImmutableArray<Outputs.GetDatabaseInstancePrivateNetworkResult> PrivateNetworks;
        public readonly string? ProjectId;
        public readonly ImmutableArray<Outputs.GetDatabaseInstanceReadReplicaResult> ReadReplicas;
        public readonly string? Region;
        public readonly ImmutableDictionary<string, string> Settings;
        public readonly ImmutableArray<string> Tags;
        public readonly string UserName;
        public readonly int VolumeSizeInGb;
        public readonly string VolumeType;

        [OutputConstructor]
        private GetDatabaseInstanceResult(
            bool backupSameRegion,

            int backupScheduleFrequency,

            int backupScheduleRetention,

            string certificate,

            bool disableBackup,

            string endpointIp,

            int endpointPort,

            string engine,

            string id,

            ImmutableDictionary<string, string> initSettings,

            string? instanceId,

            bool isHaCluster,

            ImmutableArray<Outputs.GetDatabaseInstanceLoadBalancerResult> loadBalancers,

            string? name,

            string nodeType,

            string organizationId,

            string password,

            ImmutableArray<Outputs.GetDatabaseInstancePrivateNetworkResult> privateNetworks,

            string? projectId,

            ImmutableArray<Outputs.GetDatabaseInstanceReadReplicaResult> readReplicas,

            string? region,

            ImmutableDictionary<string, string> settings,

            ImmutableArray<string> tags,

            string userName,

            int volumeSizeInGb,

            string volumeType)
        {
            BackupSameRegion = backupSameRegion;
            BackupScheduleFrequency = backupScheduleFrequency;
            BackupScheduleRetention = backupScheduleRetention;
            Certificate = certificate;
            DisableBackup = disableBackup;
            EndpointIp = endpointIp;
            EndpointPort = endpointPort;
            Engine = engine;
            Id = id;
            InitSettings = initSettings;
            InstanceId = instanceId;
            IsHaCluster = isHaCluster;
            LoadBalancers = loadBalancers;
            Name = name;
            NodeType = nodeType;
            OrganizationId = organizationId;
            Password = password;
            PrivateNetworks = privateNetworks;
            ProjectId = projectId;
            ReadReplicas = readReplicas;
            Region = region;
            Settings = settings;
            Tags = tags;
            UserName = userName;
            VolumeSizeInGb = volumeSizeInGb;
            VolumeType = volumeType;
        }
    }
}
