// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway
{
    public static class GetVpcPublicPatRule
    {
        public static Task<GetVpcPublicPatRuleResult> InvokeAsync(GetVpcPublicPatRuleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVpcPublicPatRuleResult>("scaleway:index/getVpcPublicPatRule:getVpcPublicPatRule", args ?? new GetVpcPublicPatRuleArgs(), options.WithDefaults());

        public static Output<GetVpcPublicPatRuleResult> Invoke(GetVpcPublicPatRuleInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetVpcPublicPatRuleResult>("scaleway:index/getVpcPublicPatRule:getVpcPublicPatRule", args ?? new GetVpcPublicPatRuleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVpcPublicPatRuleArgs : Pulumi.InvokeArgs
    {
        [Input("patRuleId", required: true)]
        public string PatRuleId { get; set; } = null!;

        [Input("zone")]
        public string? Zone { get; set; }

        public GetVpcPublicPatRuleArgs()
        {
        }
    }

    public sealed class GetVpcPublicPatRuleInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("patRuleId", required: true)]
        public Input<string> PatRuleId { get; set; } = null!;

        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public GetVpcPublicPatRuleInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetVpcPublicPatRuleResult
    {
        public readonly string CreatedAt;
        public readonly string GatewayId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string OrganizationId;
        public readonly string PatRuleId;
        public readonly string PrivateIp;
        public readonly int PrivatePort;
        public readonly string Protocol;
        public readonly int PublicPort;
        public readonly string UpdatedAt;
        public readonly string? Zone;

        [OutputConstructor]
        private GetVpcPublicPatRuleResult(
            string createdAt,

            string gatewayId,

            string id,

            string organizationId,

            string patRuleId,

            string privateIp,

            int privatePort,

            string protocol,

            int publicPort,

            string updatedAt,

            string? zone)
        {
            CreatedAt = createdAt;
            GatewayId = gatewayId;
            Id = id;
            OrganizationId = organizationId;
            PatRuleId = patRuleId;
            PrivateIp = privateIp;
            PrivatePort = privatePort;
            Protocol = protocol;
            PublicPort = publicPort;
            UpdatedAt = updatedAt;
            Zone = zone;
        }
    }
}
