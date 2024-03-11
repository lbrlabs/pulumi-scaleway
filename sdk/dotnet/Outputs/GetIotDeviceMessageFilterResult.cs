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
    public sealed class GetIotDeviceMessageFilterResult
    {
        /// <summary>
        /// Rule to restrict topics the device can publish to
        /// </summary>
        public readonly ImmutableArray<Outputs.GetIotDeviceMessageFilterPublishResult> Publishes;
        /// <summary>
        /// Rule to restrict topics the device can subscribe to
        /// </summary>
        public readonly ImmutableArray<Outputs.GetIotDeviceMessageFilterSubscribeResult> Subscribes;

        [OutputConstructor]
        private GetIotDeviceMessageFilterResult(
            ImmutableArray<Outputs.GetIotDeviceMessageFilterPublishResult> publishes,

            ImmutableArray<Outputs.GetIotDeviceMessageFilterSubscribeResult> subscribes)
        {
            Publishes = publishes;
            Subscribes = subscribes;
        }
    }
}
