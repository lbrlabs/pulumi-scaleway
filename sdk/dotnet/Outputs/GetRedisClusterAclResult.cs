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
    public sealed class GetRedisClusterAclResult
    {
        public readonly string Description;
        public readonly string Id;
        public readonly string Ip;

        [OutputConstructor]
        private GetRedisClusterAclResult(
            string description,

            string id,

            string ip)
        {
            Description = description;
            Id = id;
            Ip = ip;
        }
    }
}
