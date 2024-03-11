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
    public sealed class GetIpamIpsIpResult
    {
        /// <summary>
        /// The Scaleway internal IP address of the server.
        /// </summary>
        public readonly string Address;
        /// <summary>
        /// The date and time of the creation of the IP.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// The ID of the resource that the IP is bound to.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The ID of the project used as filter.
        /// </summary>
        public readonly string ProjectId;
        /// <summary>
        /// The region used as filter.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// Filter by resource ID, type or name.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetIpamIpsIpResourceResult> Resources;
        /// <summary>
        /// The tags used as filter.
        /// </summary>
        public readonly ImmutableArray<string> Tags;
        /// <summary>
        /// The date and time of the last update of the IP.
        /// </summary>
        public readonly string UpdatedAt;
        /// <summary>
        /// The zone in which the IP is.
        /// </summary>
        public readonly string Zone;

        [OutputConstructor]
        private GetIpamIpsIpResult(
            string address,

            string createdAt,

            string id,

            string projectId,

            string region,

            ImmutableArray<Outputs.GetIpamIpsIpResourceResult> resources,

            ImmutableArray<string> tags,

            string updatedAt,

            string zone)
        {
            Address = address;
            CreatedAt = createdAt;
            Id = id;
            ProjectId = projectId;
            Region = region;
            Resources = resources;
            Tags = tags;
            UpdatedAt = updatedAt;
            Zone = zone;
        }
    }
}
