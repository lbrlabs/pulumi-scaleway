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
    public static class GetBaremetalOffer
    {
        public static Task<GetBaremetalOfferResult> InvokeAsync(GetBaremetalOfferArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBaremetalOfferResult>("scaleway:index/getBaremetalOffer:getBaremetalOffer", args ?? new GetBaremetalOfferArgs(), options.WithDefaults());

        public static Output<GetBaremetalOfferResult> Invoke(GetBaremetalOfferInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetBaremetalOfferResult>("scaleway:index/getBaremetalOffer:getBaremetalOffer", args ?? new GetBaremetalOfferInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBaremetalOfferArgs : Pulumi.InvokeArgs
    {
        [Input("includeDisabled")]
        public bool? IncludeDisabled { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        [Input("offerId")]
        public string? OfferId { get; set; }

        [Input("zone")]
        public string? Zone { get; set; }

        public GetBaremetalOfferArgs()
        {
        }
    }

    public sealed class GetBaremetalOfferInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("includeDisabled")]
        public Input<bool>? IncludeDisabled { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("offerId")]
        public Input<string>? OfferId { get; set; }

        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public GetBaremetalOfferInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetBaremetalOfferResult
    {
        public readonly int Bandwidth;
        public readonly string CommercialRange;
        public readonly Outputs.GetBaremetalOfferCpuResult Cpu;
        public readonly ImmutableArray<Outputs.GetBaremetalOfferDiskResult> Disks;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool? IncludeDisabled;
        public readonly ImmutableArray<Outputs.GetBaremetalOfferMemoryResult> Memories;
        public readonly string? Name;
        public readonly string? OfferId;
        public readonly string Stock;
        public readonly string Zone;

        [OutputConstructor]
        private GetBaremetalOfferResult(
            int bandwidth,

            string commercialRange,

            Outputs.GetBaremetalOfferCpuResult cpu,

            ImmutableArray<Outputs.GetBaremetalOfferDiskResult> disks,

            string id,

            bool? includeDisabled,

            ImmutableArray<Outputs.GetBaremetalOfferMemoryResult> memories,

            string? name,

            string? offerId,

            string stock,

            string zone)
        {
            Bandwidth = bandwidth;
            CommercialRange = commercialRange;
            Cpu = cpu;
            Disks = disks;
            Id = id;
            IncludeDisabled = includeDisabled;
            Memories = memories;
            Name = name;
            OfferId = offerId;
            Stock = stock;
            Zone = zone;
        }
    }
}
