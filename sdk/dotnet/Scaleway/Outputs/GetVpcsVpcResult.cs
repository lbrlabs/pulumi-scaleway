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
    public sealed class GetVpcsVpcResult
    {
        /// <summary>
        /// Date and time of VPC's creation (RFC 3339 format).
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// The associated VPC ID.
        /// &gt; **Important:** VPCs' IDs are regional, which means they are of the form `{region}/{id}`, e.g. `fr-par/11111111-1111-1111-1111-111111111111
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Defines whether the VPC is the default one for its Project.
        /// </summary>
        public readonly bool IsDefault;
        /// <summary>
        /// The VPC name used as filter. VPCs with a name like it are listed.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The organization ID the VPC is associated with.
        /// </summary>
        public readonly string OrganizationId;
        /// <summary>
        /// The ID of the project the VPC is associated with.
        /// </summary>
        public readonly string ProjectId;
        /// <summary>
        /// `region`). The region in which vpcs exist.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// List of tags used as filter. VPCs with these exact tags are listed.
        /// </summary>
        public readonly ImmutableArray<string> Tags;
        public readonly string UpdateAt;

        [OutputConstructor]
        private GetVpcsVpcResult(
            string createdAt,

            string id,

            bool isDefault,

            string name,

            string organizationId,

            string projectId,

            string region,

            ImmutableArray<string> tags,

            string updateAt)
        {
            CreatedAt = createdAt;
            Id = id;
            IsDefault = isDefault;
            Name = name;
            OrganizationId = organizationId;
            ProjectId = projectId;
            Region = region;
            Tags = tags;
            UpdateAt = updateAt;
        }
    }
}
