// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway.Outputs
{

    [OutputType]
    public sealed class BaremetalServerIpv4
    {
        /// <summary>
        /// The address of the IPv6.
        /// </summary>
        public readonly string? Address;
        /// <summary>
        /// The id of the private network to attach.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The reverse of the IPv6.
        /// </summary>
        public readonly string? Reverse;
        /// <summary>
        /// The type of the IPv6.
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private BaremetalServerIpv4(
            string? address,

            string? id,

            string? reverse,

            string? version)
        {
            Address = address;
            Id = id;
            Reverse = reverse;
            Version = version;
        }
    }
}