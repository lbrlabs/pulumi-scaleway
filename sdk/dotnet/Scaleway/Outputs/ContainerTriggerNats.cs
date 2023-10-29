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
    public sealed class ContainerTriggerNats
    {
        /// <summary>
        /// ID of the mnq nats account.
        /// </summary>
        public readonly string? AccountId;
        /// <summary>
        /// ID of the project that contain the mnq nats account, defaults to provider's project
        /// </summary>
        public readonly string? ProjectId;
        /// <summary>
        /// `region`). The region in which the namespace should be created.
        /// </summary>
        public readonly string? Region;
        /// <summary>
        /// The subject to listen to
        /// </summary>
        public readonly string Subject;

        [OutputConstructor]
        private ContainerTriggerNats(
            string? accountId,

            string? projectId,

            string? region,

            string subject)
        {
            AccountId = accountId;
            ProjectId = projectId;
            Region = region;
            Subject = subject;
        }
    }
}
