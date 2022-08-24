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
    [ScalewayResourceType("scaleway:index/redisCluster:RedisCluster")]
    public partial class RedisCluster : Pulumi.CustomResource
    {
        /// <summary>
        /// List of acl rules.
        /// </summary>
        [Output("acls")]
        public Output<ImmutableArray<Outputs.RedisClusterAcl>> Acls { get; private set; } = null!;

        /// <summary>
        /// Number of nodes for the cluster.
        /// </summary>
        [Output("clusterSize")]
        public Output<int> ClusterSize { get; private set; } = null!;

        /// <summary>
        /// The date and time of the creation of the Redis cluster
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Name of the redis cluster
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Type of node to use for the cluster
        /// </summary>
        [Output("nodeType")]
        public Output<string> NodeType { get; private set; } = null!;

        /// <summary>
        /// Password of the user
        /// </summary>
        [Output("password")]
        public Output<string> Password { get; private set; } = null!;

        /// <summary>
        /// The project_id you want to attach the resource to
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Map of settings to define for the cluster.
        /// </summary>
        [Output("settings")]
        public Output<ImmutableDictionary<string, string>?> Settings { get; private set; } = null!;

        /// <summary>
        /// List of tags ["tag1", "tag2", ...] attached to a redis cluster
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// Whether or not TLS is enabled.
        /// </summary>
        [Output("tlsEnabled")]
        public Output<bool?> TlsEnabled { get; private set; } = null!;

        /// <summary>
        /// The date and time of the last update of the Redis cluster
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// Name of the user created when the cluster is created
        /// </summary>
        [Output("userName")]
        public Output<string> UserName { get; private set; } = null!;

        /// <summary>
        /// Redis version of the cluster
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;

        /// <summary>
        /// The zone you want to attach the resource to
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a RedisCluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RedisCluster(string name, RedisClusterArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/redisCluster:RedisCluster", name, args ?? new RedisClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RedisCluster(string name, Input<string> id, RedisClusterState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/redisCluster:RedisCluster", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing RedisCluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RedisCluster Get(string name, Input<string> id, RedisClusterState? state = null, CustomResourceOptions? options = null)
        {
            return new RedisCluster(name, id, state, options);
        }
    }

    public sealed class RedisClusterArgs : Pulumi.ResourceArgs
    {
        [Input("acls")]
        private InputList<Inputs.RedisClusterAclArgs>? _acls;

        /// <summary>
        /// List of acl rules.
        /// </summary>
        public InputList<Inputs.RedisClusterAclArgs> Acls
        {
            get => _acls ?? (_acls = new InputList<Inputs.RedisClusterAclArgs>());
            set => _acls = value;
        }

        /// <summary>
        /// Number of nodes for the cluster.
        /// </summary>
        [Input("clusterSize")]
        public Input<int>? ClusterSize { get; set; }

        /// <summary>
        /// Name of the redis cluster
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Type of node to use for the cluster
        /// </summary>
        [Input("nodeType", required: true)]
        public Input<string> NodeType { get; set; } = null!;

        /// <summary>
        /// Password of the user
        /// </summary>
        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        /// <summary>
        /// The project_id you want to attach the resource to
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("settings")]
        private InputMap<string>? _settings;

        /// <summary>
        /// Map of settings to define for the cluster.
        /// </summary>
        public InputMap<string> Settings
        {
            get => _settings ?? (_settings = new InputMap<string>());
            set => _settings = value;
        }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// List of tags ["tag1", "tag2", ...] attached to a redis cluster
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Whether or not TLS is enabled.
        /// </summary>
        [Input("tlsEnabled")]
        public Input<bool>? TlsEnabled { get; set; }

        /// <summary>
        /// Name of the user created when the cluster is created
        /// </summary>
        [Input("userName", required: true)]
        public Input<string> UserName { get; set; } = null!;

        /// <summary>
        /// Redis version of the cluster
        /// </summary>
        [Input("version", required: true)]
        public Input<string> Version { get; set; } = null!;

        /// <summary>
        /// The zone you want to attach the resource to
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public RedisClusterArgs()
        {
        }
    }

    public sealed class RedisClusterState : Pulumi.ResourceArgs
    {
        [Input("acls")]
        private InputList<Inputs.RedisClusterAclGetArgs>? _acls;

        /// <summary>
        /// List of acl rules.
        /// </summary>
        public InputList<Inputs.RedisClusterAclGetArgs> Acls
        {
            get => _acls ?? (_acls = new InputList<Inputs.RedisClusterAclGetArgs>());
            set => _acls = value;
        }

        /// <summary>
        /// Number of nodes for the cluster.
        /// </summary>
        [Input("clusterSize")]
        public Input<int>? ClusterSize { get; set; }

        /// <summary>
        /// The date and time of the creation of the Redis cluster
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Name of the redis cluster
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Type of node to use for the cluster
        /// </summary>
        [Input("nodeType")]
        public Input<string>? NodeType { get; set; }

        /// <summary>
        /// Password of the user
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// The project_id you want to attach the resource to
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("settings")]
        private InputMap<string>? _settings;

        /// <summary>
        /// Map of settings to define for the cluster.
        /// </summary>
        public InputMap<string> Settings
        {
            get => _settings ?? (_settings = new InputMap<string>());
            set => _settings = value;
        }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// List of tags ["tag1", "tag2", ...] attached to a redis cluster
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Whether or not TLS is enabled.
        /// </summary>
        [Input("tlsEnabled")]
        public Input<bool>? TlsEnabled { get; set; }

        /// <summary>
        /// The date and time of the last update of the Redis cluster
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        /// <summary>
        /// Name of the user created when the cluster is created
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        /// <summary>
        /// Redis version of the cluster
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        /// <summary>
        /// The zone you want to attach the resource to
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public RedisClusterState()
        {
        }
    }
}
