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
    /// Creates and manages Scaleway Container Token.
    /// For more information see [the documentation](https://developers.scaleway.com/en/products/containers/api/#tokens-26b085).
    /// 
    /// ## Example Usage
    /// 
    /// ### Basic
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumiverse.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var mainContainerNamespace = new Scaleway.ContainerNamespace("mainContainerNamespace");
    /// 
    ///     var mainContainer = new Scaleway.Container("mainContainer", new()
    ///     {
    ///         NamespaceId = mainContainerNamespace.Id,
    ///     });
    /// 
    ///     // Namespace Token
    ///     var @namespace = new Scaleway.ContainerToken("namespace", new()
    ///     {
    ///         NamespaceId = mainContainerNamespace.Id,
    ///         ExpiresAt = "2022-10-18T11:35:15+02:00",
    ///     });
    /// 
    ///     // Container Token
    ///     var container = new Scaleway.ContainerToken("container", new()
    ///     {
    ///         ContainerId = mainContainer.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Tokens can be imported using the `{region}/{id}`, e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:index/containerToken:ContainerToken main fr-par/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/containerToken:ContainerToken")]
    public partial class ContainerToken : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the container.
        /// 
        /// &gt; Only one of `namespace_id` or `container_id` must be set.
        /// </summary>
        [Output("containerId")]
        public Output<string?> ContainerId { get; private set; } = null!;

        /// <summary>
        /// The description of the token.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The expiration date of the token.
        /// </summary>
        [Output("expiresAt")]
        public Output<string?> ExpiresAt { get; private set; } = null!;

        /// <summary>
        /// The ID of the container namespace.
        /// </summary>
        [Output("namespaceId")]
        public Output<string?> NamespaceId { get; private set; } = null!;

        /// <summary>
        /// `region`). The region in which the namespace should be created.
        /// 
        /// &gt; **Important** Updates to any fields will recreate the token.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The token.
        /// </summary>
        [Output("token")]
        public Output<string> Token { get; private set; } = null!;


        /// <summary>
        /// Create a ContainerToken resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ContainerToken(string name, ContainerTokenArgs? args = null, CustomResourceOptions? options = null)
            : base("scaleway:index/containerToken:ContainerToken", name, args ?? new ContainerTokenArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ContainerToken(string name, Input<string> id, ContainerTokenState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/containerToken:ContainerToken", name, state, MakeResourceOptions(options, id))
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
                    "token",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ContainerToken resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ContainerToken Get(string name, Input<string> id, ContainerTokenState? state = null, CustomResourceOptions? options = null)
        {
            return new ContainerToken(name, id, state, options);
        }
    }

    public sealed class ContainerTokenArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the container.
        /// 
        /// &gt; Only one of `namespace_id` or `container_id` must be set.
        /// </summary>
        [Input("containerId")]
        public Input<string>? ContainerId { get; set; }

        /// <summary>
        /// The description of the token.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The expiration date of the token.
        /// </summary>
        [Input("expiresAt")]
        public Input<string>? ExpiresAt { get; set; }

        /// <summary>
        /// The ID of the container namespace.
        /// </summary>
        [Input("namespaceId")]
        public Input<string>? NamespaceId { get; set; }

        /// <summary>
        /// `region`). The region in which the namespace should be created.
        /// 
        /// &gt; **Important** Updates to any fields will recreate the token.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public ContainerTokenArgs()
        {
        }
        public static new ContainerTokenArgs Empty => new ContainerTokenArgs();
    }

    public sealed class ContainerTokenState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the container.
        /// 
        /// &gt; Only one of `namespace_id` or `container_id` must be set.
        /// </summary>
        [Input("containerId")]
        public Input<string>? ContainerId { get; set; }

        /// <summary>
        /// The description of the token.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The expiration date of the token.
        /// </summary>
        [Input("expiresAt")]
        public Input<string>? ExpiresAt { get; set; }

        /// <summary>
        /// The ID of the container namespace.
        /// </summary>
        [Input("namespaceId")]
        public Input<string>? NamespaceId { get; set; }

        /// <summary>
        /// `region`). The region in which the namespace should be created.
        /// 
        /// &gt; **Important** Updates to any fields will recreate the token.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("token")]
        private Input<string>? _token;

        /// <summary>
        /// The token.
        /// </summary>
        public Input<string>? Token
        {
            get => _token;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _token = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public ContainerTokenState()
        {
        }
        public static new ContainerTokenState Empty => new ContainerTokenState();
    }
}
