// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Scaleway.Inputs
{

    public sealed class IpamIpSourceGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The private network the IP lives in if the IP is a private IP.
        /// </summary>
        [Input("privateNetworkId")]
        public Input<string>? PrivateNetworkId { get; set; }

        /// <summary>
        /// The private network subnet the IP lives in if the IP is a private IP in a private network.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        /// <summary>
        /// The zone the IP lives in if the IP is a public zoned one
        /// </summary>
        [Input("zonal")]
        public Input<string>? Zonal { get; set; }

        public IpamIpSourceGetArgs()
        {
        }
        public static new IpamIpSourceGetArgs Empty => new IpamIpSourceGetArgs();
    }
}
