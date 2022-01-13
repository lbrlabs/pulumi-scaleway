// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Outputs
{

    [OutputType]
    public sealed class GetDatabaseInstancePrivateNetworkResult
    {
        public readonly string EndpointId;
        public readonly string Hostname;
        public readonly string Ip;
        public readonly string IpNet;
        /// <summary>
        /// The name of the RDB instance.
        /// Only one of `name` and `instance_id` should be specified.
        /// </summary>
        public readonly string Name;
        public readonly string PnId;
        public readonly int Port;
        public readonly string Zone;

        [OutputConstructor]
        private GetDatabaseInstancePrivateNetworkResult(
            string endpointId,

            string hostname,

            string ip,

            string ipNet,

            string name,

            string pnId,

            int port,

            string zone)
        {
            EndpointId = endpointId;
            Hostname = hostname;
            Ip = ip;
            IpNet = ipNet;
            Name = name;
            PnId = pnId;
            Port = port;
            Zone = zone;
        }
    }
}