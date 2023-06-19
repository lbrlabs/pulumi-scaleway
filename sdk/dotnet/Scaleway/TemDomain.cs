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
    /// <summary>
    /// Creates and manages Scaleway Transactional Email Domains.
    /// For more information see [the documentation](https://developers.scaleway.com/en/products/transactional_email/api/).
    /// 
    /// ## Examples
    /// 
    /// ### Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Lbrlabs.PulumiPackage.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var main = new Scaleway.TemDomain("main", new()
    ///     {
    ///         AcceptTos = true,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Add the required records to your DNS zone
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Lbrlabs.PulumiPackage.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Config();
    ///     var domainName = config.Require("domainName");
    ///     var main = new Scaleway.TemDomain("main", new()
    ///     {
    ///         AcceptTos = true,
    ///     });
    /// 
    ///     var spf = new Scaleway.DomainRecord("spf", new()
    ///     {
    ///         DnsZone = domainName,
    ///         Type = "TXT",
    ///         Data = main.SpfConfig.Apply(spfConfig =&gt; $"v=spf1 {spfConfig} -all"),
    ///     });
    /// 
    ///     var dkim = new Scaleway.DomainRecord("dkim", new()
    ///     {
    ///         DnsZone = domainName,
    ///         Type = "TXT",
    ///         Data = main.DkimConfig,
    ///     });
    /// 
    ///     var mx = new Scaleway.DomainRecord("mx", new()
    ///     {
    ///         DnsZone = domainName,
    ///         Type = "MX",
    ///         Data = ".",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Domains can be imported using the `{region}/{id}`, e.g. bash
    /// 
    /// ```sh
    ///  $ pulumi import scaleway:index/temDomain:TemDomain main fr-par/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/temDomain:TemDomain")]
    public partial class TemDomain : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Acceptation of the [Term of Service](https://tem.s3.fr-par.scw.cloud/antispam_policy.pdf).
        /// &gt; **Important:**  This attribute must be set to `true`.
        /// </summary>
        [Output("acceptTos")]
        public Output<bool> AcceptTos { get; private set; } = null!;

        /// <summary>
        /// The date and time of the Transaction Email Domain's creation (RFC 3339 format).
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The DKIM public key, as should be recorded in the DNS zone.
        /// </summary>
        [Output("dkimConfig")]
        public Output<string> DkimConfig { get; private set; } = null!;

        /// <summary>
        /// The error message if the last check failed.
        /// </summary>
        [Output("lastError")]
        public Output<string> LastError { get; private set; } = null!;

        /// <summary>
        /// The date and time the domain was last found to be valid (RFC 3339 format).
        /// </summary>
        [Output("lastValidAt")]
        public Output<string> LastValidAt { get; private set; } = null!;

        /// <summary>
        /// The domain name, must not be used in another Transactional Email Domain.
        /// &gt; **Important:** Updates to `name` will recreate the domain.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The date and time of the next scheduled check (RFC 3339 format).
        /// </summary>
        [Output("nextCheckAt")]
        public Output<string> NextCheckAt { get; private set; } = null!;

        /// <summary>
        /// `project_id`) The ID of the project the domain is associated with.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// `region`). The region in which the domain should be created.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The date and time of the revocation of the domain (RFC 3339 format).
        /// </summary>
        [Output("revokedAt")]
        public Output<string> RevokedAt { get; private set; } = null!;

        /// <summary>
        /// SMTP host to use to send emails
        /// </summary>
        [Output("smtpHost")]
        public Output<string> SmtpHost { get; private set; } = null!;

        /// <summary>
        /// SMTP port to use to send emails over TLS. (Port 587)
        /// </summary>
        [Output("smtpPort")]
        public Output<int> SmtpPort { get; private set; } = null!;

        /// <summary>
        /// SMTP port to use to send emails over TLS. (Port 2587)
        /// </summary>
        [Output("smtpPortAlternative")]
        public Output<int> SmtpPortAlternative { get; private set; } = null!;

        /// <summary>
        /// SMTP port to use to send emails. (Port 25)
        /// </summary>
        [Output("smtpPortUnsecure")]
        public Output<int> SmtpPortUnsecure { get; private set; } = null!;

        /// <summary>
        /// SMTPS port to use to send emails over TLS Wrapper. (Port 465)
        /// </summary>
        [Output("smtpsPort")]
        public Output<int> SmtpsPort { get; private set; } = null!;

        /// <summary>
        /// SMTPS port to use to send emails over TLS Wrapper. (Port 2465)
        /// </summary>
        [Output("smtpsPortAlternative")]
        public Output<int> SmtpsPortAlternative { get; private set; } = null!;

        /// <summary>
        /// The snippet of the SPF record that should be registered in the DNS zone.
        /// </summary>
        [Output("spfConfig")]
        public Output<string> SpfConfig { get; private set; } = null!;

        /// <summary>
        /// The status of the Transaction Email Domain.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a TemDomain resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TemDomain(string name, TemDomainArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/temDomain:TemDomain", name, args ?? new TemDomainArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TemDomain(string name, Input<string> id, TemDomainState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/temDomain:TemDomain", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/lbrlabs",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing TemDomain resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TemDomain Get(string name, Input<string> id, TemDomainState? state = null, CustomResourceOptions? options = null)
        {
            return new TemDomain(name, id, state, options);
        }
    }

    public sealed class TemDomainArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Acceptation of the [Term of Service](https://tem.s3.fr-par.scw.cloud/antispam_policy.pdf).
        /// &gt; **Important:**  This attribute must be set to `true`.
        /// </summary>
        [Input("acceptTos", required: true)]
        public Input<bool> AcceptTos { get; set; } = null!;

        /// <summary>
        /// The domain name, must not be used in another Transactional Email Domain.
        /// &gt; **Important:** Updates to `name` will recreate the domain.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the domain is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`). The region in which the domain should be created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public TemDomainArgs()
        {
        }
        public static new TemDomainArgs Empty => new TemDomainArgs();
    }

    public sealed class TemDomainState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Acceptation of the [Term of Service](https://tem.s3.fr-par.scw.cloud/antispam_policy.pdf).
        /// &gt; **Important:**  This attribute must be set to `true`.
        /// </summary>
        [Input("acceptTos")]
        public Input<bool>? AcceptTos { get; set; }

        /// <summary>
        /// The date and time of the Transaction Email Domain's creation (RFC 3339 format).
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// The DKIM public key, as should be recorded in the DNS zone.
        /// </summary>
        [Input("dkimConfig")]
        public Input<string>? DkimConfig { get; set; }

        /// <summary>
        /// The error message if the last check failed.
        /// </summary>
        [Input("lastError")]
        public Input<string>? LastError { get; set; }

        /// <summary>
        /// The date and time the domain was last found to be valid (RFC 3339 format).
        /// </summary>
        [Input("lastValidAt")]
        public Input<string>? LastValidAt { get; set; }

        /// <summary>
        /// The domain name, must not be used in another Transactional Email Domain.
        /// &gt; **Important:** Updates to `name` will recreate the domain.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The date and time of the next scheduled check (RFC 3339 format).
        /// </summary>
        [Input("nextCheckAt")]
        public Input<string>? NextCheckAt { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the domain is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`). The region in which the domain should be created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The date and time of the revocation of the domain (RFC 3339 format).
        /// </summary>
        [Input("revokedAt")]
        public Input<string>? RevokedAt { get; set; }

        /// <summary>
        /// SMTP host to use to send emails
        /// </summary>
        [Input("smtpHost")]
        public Input<string>? SmtpHost { get; set; }

        /// <summary>
        /// SMTP port to use to send emails over TLS. (Port 587)
        /// </summary>
        [Input("smtpPort")]
        public Input<int>? SmtpPort { get; set; }

        /// <summary>
        /// SMTP port to use to send emails over TLS. (Port 2587)
        /// </summary>
        [Input("smtpPortAlternative")]
        public Input<int>? SmtpPortAlternative { get; set; }

        /// <summary>
        /// SMTP port to use to send emails. (Port 25)
        /// </summary>
        [Input("smtpPortUnsecure")]
        public Input<int>? SmtpPortUnsecure { get; set; }

        /// <summary>
        /// SMTPS port to use to send emails over TLS Wrapper. (Port 465)
        /// </summary>
        [Input("smtpsPort")]
        public Input<int>? SmtpsPort { get; set; }

        /// <summary>
        /// SMTPS port to use to send emails over TLS Wrapper. (Port 2465)
        /// </summary>
        [Input("smtpsPortAlternative")]
        public Input<int>? SmtpsPortAlternative { get; set; }

        /// <summary>
        /// The snippet of the SPF record that should be registered in the DNS zone.
        /// </summary>
        [Input("spfConfig")]
        public Input<string>? SpfConfig { get; set; }

        /// <summary>
        /// The status of the Transaction Email Domain.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public TemDomainState()
        {
        }
        public static new TemDomainState Empty => new TemDomainState();
    }
}
