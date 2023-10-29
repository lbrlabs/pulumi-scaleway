// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Scaleway.Outputs
{

    [OutputType]
    public sealed class GetVpcPrivateNetworkIpv4SubnetResult
    {
        public readonly string Address;
        public readonly string CreatedAt;
        /// <summary>
        /// The ID of the private network.
        /// </summary>
        public readonly string Id;
        public readonly int PrefixLength;
        public readonly string Subnet;
        public readonly string SubnetMask;
        public readonly string UpdatedAt;

        [OutputConstructor]
        private GetVpcPrivateNetworkIpv4SubnetResult(
            string address,

            string createdAt,

            string id,

            int prefixLength,

            string subnet,

            string subnetMask,

            string updatedAt)
        {
            Address = address;
            CreatedAt = createdAt;
            Id = id;
            PrefixLength = prefixLength;
            Subnet = subnet;
            SubnetMask = subnetMask;
            UpdatedAt = updatedAt;
        }
    }
}
