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
    [ScalewayResourceType("scaleway:index/vpcPublicGatewayIp:VpcPublicGatewayIp")]
    public partial class VpcPublicGatewayIp : Pulumi.CustomResource
    {
        /// <summary>
        /// the IP itself
        /// </summary>
        [Output("address")]
        public Output<string> Address { get; private set; } = null!;

        /// <summary>
        /// The date and time of the creation of the public gateway IP
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

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
        /// reverse domain name for the IP address
        /// </summary>
        [Output("reverse")]
        public Output<string> Reverse { get; private set; } = null!;

        /// <summary>
        /// The tags associated with public gateway IP
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// The date and time of the last update of the public gateway IP
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// The zone you want to attach the resource to
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a VpcPublicGatewayIp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpcPublicGatewayIp(string name, VpcPublicGatewayIpArgs? args = null, CustomResourceOptions? options = null)
            : base("scaleway:index/vpcPublicGatewayIp:VpcPublicGatewayIp", name, args ?? new VpcPublicGatewayIpArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VpcPublicGatewayIp(string name, Input<string> id, VpcPublicGatewayIpState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/vpcPublicGatewayIp:VpcPublicGatewayIp", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VpcPublicGatewayIp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpcPublicGatewayIp Get(string name, Input<string> id, VpcPublicGatewayIpState? state = null, CustomResourceOptions? options = null)
        {
            return new VpcPublicGatewayIp(name, id, state, options);
        }
    }

    public sealed class VpcPublicGatewayIpArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The project_id you want to attach the resource to
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// reverse domain name for the IP address
        /// </summary>
        [Input("reverse")]
        public Input<string>? Reverse { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The tags associated with public gateway IP
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The zone you want to attach the resource to
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public VpcPublicGatewayIpArgs()
        {
        }
    }

    public sealed class VpcPublicGatewayIpState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// the IP itself
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// The date and time of the creation of the public gateway IP
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

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

        /// <summary>
        /// reverse domain name for the IP address
        /// </summary>
        [Input("reverse")]
        public Input<string>? Reverse { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The tags associated with public gateway IP
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The date and time of the last update of the public gateway IP
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        /// <summary>
        /// The zone you want to attach the resource to
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public VpcPublicGatewayIpState()
        {
        }
    }
}
