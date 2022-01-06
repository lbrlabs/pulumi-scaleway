// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway
{
    [ScalewayResourceType("scaleway:index/iOTRoute:IOTRoute")]
    public partial class IOTRoute : Pulumi.CustomResource
    {
        /// <summary>
        /// The date and time of the creation of the IoT Route
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Database Route parameters
        /// </summary>
        [Output("database")]
        public Output<Outputs.IOTRouteDatabase?> Database { get; private set; } = null!;

        /// <summary>
        /// The ID of the route's hub
        /// </summary>
        [Output("hubId")]
        public Output<string> HubId { get; private set; } = null!;

        /// <summary>
        /// The name of the route
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The region you want to attach the resource to
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Rest Route parameters
        /// </summary>
        [Output("rest")]
        public Output<Outputs.IOTRouteRest?> Rest { get; private set; } = null!;

        /// <summary>
        /// S3 Route parameters
        /// </summary>
        [Output("s3")]
        public Output<Outputs.IOTRouteS3?> S3 { get; private set; } = null!;

        /// <summary>
        /// The Topic the route subscribes to (wildcards allowed)
        /// </summary>
        [Output("topic")]
        public Output<string> Topic { get; private set; } = null!;


        /// <summary>
        /// Create a IOTRoute resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IOTRoute(string name, IOTRouteArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/iOTRoute:IOTRoute", name, args ?? new IOTRouteArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IOTRoute(string name, Input<string> id, IOTRouteState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/iOTRoute:IOTRoute", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing IOTRoute resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IOTRoute Get(string name, Input<string> id, IOTRouteState? state = null, CustomResourceOptions? options = null)
        {
            return new IOTRoute(name, id, state, options);
        }
    }

    public sealed class IOTRouteArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Database Route parameters
        /// </summary>
        [Input("database")]
        public Input<Inputs.IOTRouteDatabaseArgs>? Database { get; set; }

        /// <summary>
        /// The ID of the route's hub
        /// </summary>
        [Input("hubId", required: true)]
        public Input<string> HubId { get; set; } = null!;

        /// <summary>
        /// The name of the route
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The region you want to attach the resource to
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Rest Route parameters
        /// </summary>
        [Input("rest")]
        public Input<Inputs.IOTRouteRestArgs>? Rest { get; set; }

        /// <summary>
        /// S3 Route parameters
        /// </summary>
        [Input("s3")]
        public Input<Inputs.IOTRouteS3Args>? S3 { get; set; }

        /// <summary>
        /// The Topic the route subscribes to (wildcards allowed)
        /// </summary>
        [Input("topic", required: true)]
        public Input<string> Topic { get; set; } = null!;

        public IOTRouteArgs()
        {
        }
    }

    public sealed class IOTRouteState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The date and time of the creation of the IoT Route
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Database Route parameters
        /// </summary>
        [Input("database")]
        public Input<Inputs.IOTRouteDatabaseGetArgs>? Database { get; set; }

        /// <summary>
        /// The ID of the route's hub
        /// </summary>
        [Input("hubId")]
        public Input<string>? HubId { get; set; }

        /// <summary>
        /// The name of the route
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The region you want to attach the resource to
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Rest Route parameters
        /// </summary>
        [Input("rest")]
        public Input<Inputs.IOTRouteRestGetArgs>? Rest { get; set; }

        /// <summary>
        /// S3 Route parameters
        /// </summary>
        [Input("s3")]
        public Input<Inputs.IOTRouteS3GetArgs>? S3 { get; set; }

        /// <summary>
        /// The Topic the route subscribes to (wildcards allowed)
        /// </summary>
        [Input("topic")]
        public Input<string>? Topic { get; set; }

        public IOTRouteState()
        {
        }
    }
}