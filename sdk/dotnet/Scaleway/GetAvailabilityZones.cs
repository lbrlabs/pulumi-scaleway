// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Scaleway
{
    public static class GetAvailabilityZones
    {
        /// <summary>
        /// Use this data source to get the available zones information based on its Region.
        /// 
        /// For technical and legal reasons, some products are split by Region or by Availability Zones. When using such product,
        /// you can choose the location that better fits your need (country, latency, …).
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var main = Scaleway.GetAvailabilityZones.Invoke(new()
        ///     {
        ///         Region = "nl-ams",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAvailabilityZonesResult> InvokeAsync(GetAvailabilityZonesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAvailabilityZonesResult>("scaleway:index/getAvailabilityZones:getAvailabilityZones", args ?? new GetAvailabilityZonesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get the available zones information based on its Region.
        /// 
        /// For technical and legal reasons, some products are split by Region or by Availability Zones. When using such product,
        /// you can choose the location that better fits your need (country, latency, …).
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var main = Scaleway.GetAvailabilityZones.Invoke(new()
        ///     {
        ///         Region = "nl-ams",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetAvailabilityZonesResult> Invoke(GetAvailabilityZonesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAvailabilityZonesResult>("scaleway:index/getAvailabilityZones:getAvailabilityZones", args ?? new GetAvailabilityZonesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAvailabilityZonesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Region is represented as a Geographical area such as France. Defaults: `fr-par`.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        public GetAvailabilityZonesArgs()
        {
        }
        public static new GetAvailabilityZonesArgs Empty => new GetAvailabilityZonesArgs();
    }

    public sealed class GetAvailabilityZonesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Region is represented as a Geographical area such as France. Defaults: `fr-par`.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public GetAvailabilityZonesInvokeArgs()
        {
        }
        public static new GetAvailabilityZonesInvokeArgs Empty => new GetAvailabilityZonesInvokeArgs();
    }


    [OutputType]
    public sealed class GetAvailabilityZonesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Region;
        /// <summary>
        /// List of availability zones by regions
        /// </summary>
        public readonly ImmutableArray<string> Zones;

        [OutputConstructor]
        private GetAvailabilityZonesResult(
            string id,

            string? region,

            ImmutableArray<string> zones)
        {
            Id = id;
            Region = region;
            Zones = zones;
        }
    }
}
