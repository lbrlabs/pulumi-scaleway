// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway
{
    public static class GetLoadbalancerIp
    {
        public static Task<GetLoadbalancerIpResult> InvokeAsync(GetLoadbalancerIpArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetLoadbalancerIpResult>("scaleway:index/getLoadbalancerIp:getLoadbalancerIp", args ?? new GetLoadbalancerIpArgs(), options.WithDefaults());

        public static Output<GetLoadbalancerIpResult> Invoke(GetLoadbalancerIpInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetLoadbalancerIpResult>("scaleway:index/getLoadbalancerIp:getLoadbalancerIp", args ?? new GetLoadbalancerIpInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLoadbalancerIpArgs : Pulumi.InvokeArgs
    {
        [Input("ipAddress")]
        public string? IpAddress { get; set; }

        [Input("ipId")]
        public string? IpId { get; set; }

        public GetLoadbalancerIpArgs()
        {
        }
    }

    public sealed class GetLoadbalancerIpInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        [Input("ipId")]
        public Input<string>? IpId { get; set; }

        public GetLoadbalancerIpInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetLoadbalancerIpResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? IpAddress;
        public readonly string? IpId;
        public readonly string LbId;
        public readonly string OrganizationId;
        public readonly string ProjectId;
        public readonly string Region;
        public readonly string Reverse;
        public readonly string Zone;

        [OutputConstructor]
        private GetLoadbalancerIpResult(
            string id,

            string? ipAddress,

            string? ipId,

            string lbId,

            string organizationId,

            string projectId,

            string region,

            string reverse,

            string zone)
        {
            Id = id;
            IpAddress = ipAddress;
            IpId = ipId;
            LbId = lbId;
            OrganizationId = organizationId;
            ProjectId = projectId;
            Region = region;
            Reverse = reverse;
            Zone = zone;
        }
    }
}
