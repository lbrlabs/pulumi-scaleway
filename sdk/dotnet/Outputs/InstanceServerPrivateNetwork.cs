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
    public sealed class InstanceServerPrivateNetwork
    {
        /// <summary>
        /// MAC address of the NIC
        /// </summary>
        public readonly string? MacAddress;
        /// <summary>
        /// The Private Network ID
        /// </summary>
        public readonly string PnId;
        /// <summary>
        /// The private NIC state
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// `zone`) The zone in which the server should be created.
        /// </summary>
        public readonly string? Zone;

        [OutputConstructor]
        private InstanceServerPrivateNetwork(
            string? macAddress,

            string pnId,

            string? status,

            string? zone)
        {
            MacAddress = macAddress;
            PnId = pnId;
            Status = status;
            Zone = zone;
        }
    }
}
