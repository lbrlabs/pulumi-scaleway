// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway
{
    [ScalewayResourceType("scaleway:index/loadbalancerRoute:LoadbalancerRoute")]
    public partial class LoadbalancerRoute : Pulumi.CustomResource
    {
        /// <summary>
        /// The backend ID destination of redirection
        /// </summary>
        [Output("backendId")]
        public Output<string> BackendId { get; private set; } = null!;

        /// <summary>
        /// The frontend ID origin of redirection
        /// </summary>
        [Output("frontendId")]
        public Output<string> FrontendId { get; private set; } = null!;

        /// <summary>
        /// The domain to match against
        /// </summary>
        [Output("matchSni")]
        public Output<string?> MatchSni { get; private set; } = null!;


        /// <summary>
        /// Create a LoadbalancerRoute resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LoadbalancerRoute(string name, LoadbalancerRouteArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/loadbalancerRoute:LoadbalancerRoute", name, args ?? new LoadbalancerRouteArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LoadbalancerRoute(string name, Input<string> id, LoadbalancerRouteState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/loadbalancerRoute:LoadbalancerRoute", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LoadbalancerRoute resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LoadbalancerRoute Get(string name, Input<string> id, LoadbalancerRouteState? state = null, CustomResourceOptions? options = null)
        {
            return new LoadbalancerRoute(name, id, state, options);
        }
    }

    public sealed class LoadbalancerRouteArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The backend ID destination of redirection
        /// </summary>
        [Input("backendId", required: true)]
        public Input<string> BackendId { get; set; } = null!;

        /// <summary>
        /// The frontend ID origin of redirection
        /// </summary>
        [Input("frontendId", required: true)]
        public Input<string> FrontendId { get; set; } = null!;

        /// <summary>
        /// The domain to match against
        /// </summary>
        [Input("matchSni")]
        public Input<string>? MatchSni { get; set; }

        public LoadbalancerRouteArgs()
        {
        }
    }

    public sealed class LoadbalancerRouteState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The backend ID destination of redirection
        /// </summary>
        [Input("backendId")]
        public Input<string>? BackendId { get; set; }

        /// <summary>
        /// The frontend ID origin of redirection
        /// </summary>
        [Input("frontendId")]
        public Input<string>? FrontendId { get; set; }

        /// <summary>
        /// The domain to match against
        /// </summary>
        [Input("matchSni")]
        public Input<string>? MatchSni { get; set; }

        public LoadbalancerRouteState()
        {
        }
    }
}
