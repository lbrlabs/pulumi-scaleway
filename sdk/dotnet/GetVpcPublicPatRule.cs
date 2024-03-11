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
    public static class GetVpcPublicPatRule
    {
        /// <summary>
        /// Gets information about a public gateway PAT rule. For further information please check the
        /// API [documentation](https://developers.scaleway.com/en/products/vpc-gw/api/v1/#get-8faeea)
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// using Scaleway = Pulumiverse.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var sg01 = new Scaleway.InstanceSecurityGroup("sg01", new()
        ///     {
        ///         InboundDefaultPolicy = "drop",
        ///         OutboundDefaultPolicy = "accept",
        ///         InboundRules = new[]
        ///         {
        ///             new Scaleway.Inputs.InstanceSecurityGroupInboundRuleArgs
        ///             {
        ///                 Action = "accept",
        ///                 Port = 22,
        ///                 Protocol = "TCP",
        ///             },
        ///         },
        ///     });
        /// 
        ///     var srv01 = new Scaleway.InstanceServer("srv01", new()
        ///     {
        ///         Type = "PLAY2-NANO",
        ///         Image = "ubuntu_jammy",
        ///         SecurityGroupId = sg01.Id,
        ///     });
        /// 
        ///     var pn01 = new Scaleway.VpcPrivateNetwork("pn01");
        /// 
        ///     var pnic01 = new Scaleway.InstancePrivateNic("pnic01", new()
        ///     {
        ///         ServerId = srv01.Id,
        ///         PrivateNetworkId = pn01.Id,
        ///     });
        /// 
        ///     var dhcp01 = new Scaleway.VpcPublicGatewayDhcp("dhcp01", new()
        ///     {
        ///         Subnet = "192.168.0.0/24",
        ///     });
        /// 
        ///     var ip01 = new Scaleway.VpcPublicGatewayIp("ip01");
        /// 
        ///     var pg01 = new Scaleway.VpcPublicGateway("pg01", new()
        ///     {
        ///         Type = "VPC-GW-S",
        ///         IpId = ip01.Id,
        ///     });
        /// 
        ///     var gn01 = new Scaleway.VpcGatewayNetwork("gn01", new()
        ///     {
        ///         GatewayId = pg01.Id,
        ///         PrivateNetworkId = pn01.Id,
        ///         DhcpId = dhcp01.Id,
        ///         CleanupDhcp = true,
        ///         EnableMasquerade = true,
        ///     });
        /// 
        ///     var rsv01 = new Scaleway.VpcPublicGatewayDhcpReservation("rsv01", new()
        ///     {
        ///         GatewayNetworkId = gn01.Id,
        ///         MacAddress = pnic01.MacAddress,
        ///         IpAddress = "192.168.0.7",
        ///     });
        /// 
        ///     var pat01 = new Scaleway.VpcPublicGatewayPatRule("pat01", new()
        ///     {
        ///         GatewayId = pg01.Id,
        ///         PrivateIp = rsv01.IpAddress,
        ///         PrivatePort = 22,
        ///         PublicPort = 2202,
        ///         Protocol = "tcp",
        ///     });
        /// 
        ///     var main = Scaleway.GetVpcPublicPatRule.Invoke(new()
        ///     {
        ///         PatRuleId = pat01.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetVpcPublicPatRuleResult> InvokeAsync(GetVpcPublicPatRuleArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVpcPublicPatRuleResult>("scaleway:index/getVpcPublicPatRule:getVpcPublicPatRule", args ?? new GetVpcPublicPatRuleArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a public gateway PAT rule. For further information please check the
        /// API [documentation](https://developers.scaleway.com/en/products/vpc-gw/api/v1/#get-8faeea)
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// using Scaleway = Pulumiverse.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var sg01 = new Scaleway.InstanceSecurityGroup("sg01", new()
        ///     {
        ///         InboundDefaultPolicy = "drop",
        ///         OutboundDefaultPolicy = "accept",
        ///         InboundRules = new[]
        ///         {
        ///             new Scaleway.Inputs.InstanceSecurityGroupInboundRuleArgs
        ///             {
        ///                 Action = "accept",
        ///                 Port = 22,
        ///                 Protocol = "TCP",
        ///             },
        ///         },
        ///     });
        /// 
        ///     var srv01 = new Scaleway.InstanceServer("srv01", new()
        ///     {
        ///         Type = "PLAY2-NANO",
        ///         Image = "ubuntu_jammy",
        ///         SecurityGroupId = sg01.Id,
        ///     });
        /// 
        ///     var pn01 = new Scaleway.VpcPrivateNetwork("pn01");
        /// 
        ///     var pnic01 = new Scaleway.InstancePrivateNic("pnic01", new()
        ///     {
        ///         ServerId = srv01.Id,
        ///         PrivateNetworkId = pn01.Id,
        ///     });
        /// 
        ///     var dhcp01 = new Scaleway.VpcPublicGatewayDhcp("dhcp01", new()
        ///     {
        ///         Subnet = "192.168.0.0/24",
        ///     });
        /// 
        ///     var ip01 = new Scaleway.VpcPublicGatewayIp("ip01");
        /// 
        ///     var pg01 = new Scaleway.VpcPublicGateway("pg01", new()
        ///     {
        ///         Type = "VPC-GW-S",
        ///         IpId = ip01.Id,
        ///     });
        /// 
        ///     var gn01 = new Scaleway.VpcGatewayNetwork("gn01", new()
        ///     {
        ///         GatewayId = pg01.Id,
        ///         PrivateNetworkId = pn01.Id,
        ///         DhcpId = dhcp01.Id,
        ///         CleanupDhcp = true,
        ///         EnableMasquerade = true,
        ///     });
        /// 
        ///     var rsv01 = new Scaleway.VpcPublicGatewayDhcpReservation("rsv01", new()
        ///     {
        ///         GatewayNetworkId = gn01.Id,
        ///         MacAddress = pnic01.MacAddress,
        ///         IpAddress = "192.168.0.7",
        ///     });
        /// 
        ///     var pat01 = new Scaleway.VpcPublicGatewayPatRule("pat01", new()
        ///     {
        ///         GatewayId = pg01.Id,
        ///         PrivateIp = rsv01.IpAddress,
        ///         PrivatePort = 22,
        ///         PublicPort = 2202,
        ///         Protocol = "tcp",
        ///     });
        /// 
        ///     var main = Scaleway.GetVpcPublicPatRule.Invoke(new()
        ///     {
        ///         PatRuleId = pat01.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetVpcPublicPatRuleResult> Invoke(GetVpcPublicPatRuleInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVpcPublicPatRuleResult>("scaleway:index/getVpcPublicPatRule:getVpcPublicPatRule", args ?? new GetVpcPublicPatRuleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVpcPublicPatRuleArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the PAT rule to retrieve
        /// </summary>
        [Input("patRuleId", required: true)]
        public string PatRuleId { get; set; } = null!;

        /// <summary>
        /// `zone`) The zone in which
        /// the image exists.
        /// </summary>
        [Input("zone")]
        public string? Zone { get; set; }

        public GetVpcPublicPatRuleArgs()
        {
        }
        public static new GetVpcPublicPatRuleArgs Empty => new GetVpcPublicPatRuleArgs();
    }

    public sealed class GetVpcPublicPatRuleInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the PAT rule to retrieve
        /// </summary>
        [Input("patRuleId", required: true)]
        public Input<string> PatRuleId { get; set; } = null!;

        /// <summary>
        /// `zone`) The zone in which
        /// the image exists.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public GetVpcPublicPatRuleInvokeArgs()
        {
        }
        public static new GetVpcPublicPatRuleInvokeArgs Empty => new GetVpcPublicPatRuleInvokeArgs();
    }


    [OutputType]
    public sealed class GetVpcPublicPatRuleResult
    {
        public readonly string CreatedAt;
        /// <summary>
        /// The ID of the public gateway.
        /// </summary>
        public readonly string GatewayId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string OrganizationId;
        public readonly string PatRuleId;
        /// <summary>
        /// The Private IP to forward data to (IP address).
        /// </summary>
        public readonly string PrivateIp;
        /// <summary>
        /// The Private port to translate to.
        /// </summary>
        public readonly int PrivatePort;
        /// <summary>
        /// The Protocol the rule should apply to. Possible values are both, tcp and udp.
        /// </summary>
        public readonly string Protocol;
        /// <summary>
        /// The Public port to listen on.
        /// </summary>
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
