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
    /// <summary>
    /// Creates and manages Scaleway Cockpit Tokens.
    /// 
    /// For more information consult the [documentation](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#tokens).
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumi.Scaleway;
    /// using Scaleway = Pulumiverse.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var mainCockpit = Scaleway.GetCockpit.Invoke();
    /// 
    ///     // Create a token for the cockpit that can write metrics and logs
    ///     var mainCockpitToken = new Scaleway.CockpitToken("mainCockpitToken", new()
    ///     {
    ///         ProjectId = mainCockpit.Apply(getCockpitResult =&gt; getCockpitResult.ProjectId),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumi.Scaleway;
    /// using Scaleway = Pulumiverse.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var mainCockpit = Scaleway.GetCockpit.Invoke();
    /// 
    ///     // Create a token for the cockpit that can read metrics and logs but not write
    ///     var mainCockpitToken = new Scaleway.CockpitToken("mainCockpitToken", new()
    ///     {
    ///         ProjectId = mainCockpit.Apply(getCockpitResult =&gt; getCockpitResult.ProjectId),
    ///         Scopes = new Scaleway.Inputs.CockpitTokenScopesArgs
    ///         {
    ///             QueryMetrics = true,
    ///             WriteMetrics = false,
    ///             QueryLogs = true,
    ///             WriteLogs = false,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Cockpits can be imported using the token ID, e.g. bash
    /// 
    /// ```sh
    ///  $ pulumi import scaleway:index/cockpitToken:CockpitToken main 11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/cockpitToken:CockpitToken")]
    public partial class CockpitToken : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the token
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// `project_id`) The ID of the project the cockpit is associated with.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Allowed scopes
        /// </summary>
        [Output("scopes")]
        public Output<Outputs.CockpitTokenScopes> Scopes { get; private set; } = null!;

        /// <summary>
        /// The secret key of the token
        /// </summary>
        [Output("secretKey")]
        public Output<string> SecretKey { get; private set; } = null!;


        /// <summary>
        /// Create a CockpitToken resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CockpitToken(string name, CockpitTokenArgs? args = null, CustomResourceOptions? options = null)
            : base("scaleway:index/cockpitToken:CockpitToken", name, args ?? new CockpitTokenArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CockpitToken(string name, Input<string> id, CockpitTokenState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/cockpitToken:CockpitToken", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
                AdditionalSecretOutputs =
                {
                    "secretKey",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing CockpitToken resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CockpitToken Get(string name, Input<string> id, CockpitTokenState? state = null, CustomResourceOptions? options = null)
        {
            return new CockpitToken(name, id, state, options);
        }
    }

    public sealed class CockpitTokenArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the token
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the cockpit is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Allowed scopes
        /// </summary>
        [Input("scopes")]
        public Input<Inputs.CockpitTokenScopesArgs>? Scopes { get; set; }

        public CockpitTokenArgs()
        {
        }
        public static new CockpitTokenArgs Empty => new CockpitTokenArgs();
    }

    public sealed class CockpitTokenState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the token
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the cockpit is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Allowed scopes
        /// </summary>
        [Input("scopes")]
        public Input<Inputs.CockpitTokenScopesGetArgs>? Scopes { get; set; }

        [Input("secretKey")]
        private Input<string>? _secretKey;

        /// <summary>
        /// The secret key of the token
        /// </summary>
        public Input<string>? SecretKey
        {
            get => _secretKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secretKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public CockpitTokenState()
        {
        }
        public static new CockpitTokenState Empty => new CockpitTokenState();
    }
}