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

    public sealed class DatabaseInstancePrivateNetworkArgs : Pulumi.ResourceArgs
    {
        [Input("endpointId")]
        public Input<string>? EndpointId { get; set; }

        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        [Input("ip")]
        public Input<string>? Ip { get; set; }

        [Input("ipNet", required: true)]
        public Input<string> IpNet { get; set; } = null!;

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("pnId", required: true)]
        public Input<string> PnId { get; set; } = null!;

        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public DatabaseInstancePrivateNetworkArgs()
        {
        }
    }
}
