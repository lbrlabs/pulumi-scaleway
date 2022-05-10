// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway
{
    [ScalewayResourceType("scaleway:index/vpcPublicGateway:VpcPublicGateway")]
    public partial class VpcPublicGateway : Pulumi.CustomResource
    {
        /// <summary>
        /// The date and time of the creation of the public gateway
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// attach an existing IP to the gateway
        /// </summary>
        [Output("ipId")]
        public Output<string> IpId { get; private set; } = null!;

        /// <summary>
        /// name of the gateway
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The organization_id you want to attach the resource to
        /// </summary>
        [Output("organizationId")]
        public Output<string> OrganizationId { get; private set; } = null!;

        /// <summary>
        /// The project_id you want to attach the resource to
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The tags associated with public gateway
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// gateway type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The date and time of the last update of the public gateway
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// override the gateway's default recursive DNS servers, if DNS features are enabled
        /// </summary>
        [Output("upstreamDnsServers")]
        public Output<ImmutableArray<string>> UpstreamDnsServers { get; private set; } = null!;

        /// <summary>
        /// The zone you want to attach the resource to
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a VpcPublicGateway resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpcPublicGateway(string name, VpcPublicGatewayArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/vpcPublicGateway:VpcPublicGateway", name, args ?? new VpcPublicGatewayArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VpcPublicGateway(string name, Input<string> id, VpcPublicGatewayState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/vpcPublicGateway:VpcPublicGateway", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/jaxxstorm/pulumi-scaleway/releases/download/${VERSION}",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing VpcPublicGateway resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpcPublicGateway Get(string name, Input<string> id, VpcPublicGatewayState? state = null, CustomResourceOptions? options = null)
        {
            return new VpcPublicGateway(name, id, state, options);
        }
    }

    public sealed class VpcPublicGatewayArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// attach an existing IP to the gateway
        /// </summary>
        [Input("ipId")]
        public Input<string>? IpId { get; set; }

        /// <summary>
        /// name of the gateway
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The project_id you want to attach the resource to
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The tags associated with public gateway
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// gateway type
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("upstreamDnsServers")]
        private InputList<string>? _upstreamDnsServers;

        /// <summary>
        /// override the gateway's default recursive DNS servers, if DNS features are enabled
        /// </summary>
        public InputList<string> UpstreamDnsServers
        {
            get => _upstreamDnsServers ?? (_upstreamDnsServers = new InputList<string>());
            set => _upstreamDnsServers = value;
        }

        /// <summary>
        /// The zone you want to attach the resource to
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public VpcPublicGatewayArgs()
        {
        }
    }

    public sealed class VpcPublicGatewayState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The date and time of the creation of the public gateway
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// attach an existing IP to the gateway
        /// </summary>
        [Input("ipId")]
        public Input<string>? IpId { get; set; }

        /// <summary>
        /// name of the gateway
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The organization_id you want to attach the resource to
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        /// <summary>
        /// The project_id you want to attach the resource to
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The tags associated with public gateway
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// gateway type
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The date and time of the last update of the public gateway
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        [Input("upstreamDnsServers")]
        private InputList<string>? _upstreamDnsServers;

        /// <summary>
        /// override the gateway's default recursive DNS servers, if DNS features are enabled
        /// </summary>
        public InputList<string> UpstreamDnsServers
        {
            get => _upstreamDnsServers ?? (_upstreamDnsServers = new InputList<string>());
            set => _upstreamDnsServers = value;
        }

        /// <summary>
        /// The zone you want to attach the resource to
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public VpcPublicGatewayState()
        {
        }
    }
}
