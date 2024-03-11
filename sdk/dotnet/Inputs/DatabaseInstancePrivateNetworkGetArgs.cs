// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway.Inputs
{

    public sealed class DatabaseInstancePrivateNetworkGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the endpoint.
        /// </summary>
        [Input("endpointId")]
        public Input<string>? EndpointId { get; set; }

        /// <summary>
        /// Hostname of the endpoint.
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        /// <summary>
        /// IPv4 address on the network.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// The IP with the given mask within the private subnet
        /// </summary>
        [Input("ipNet")]
        public Input<string>? IpNet { get; set; }

        /// <summary>
        /// The name of the Database Instance.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The private network ID
        /// </summary>
        [Input("pnId", required: true)]
        public Input<string> PnId { get; set; } = null!;

        /// <summary>
        /// Port in the Private Network.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The zone you want to attach the resource to
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public DatabaseInstancePrivateNetworkGetArgs()
        {
        }
        public static new DatabaseInstancePrivateNetworkGetArgs Empty => new DatabaseInstancePrivateNetworkGetArgs();
    }
}
