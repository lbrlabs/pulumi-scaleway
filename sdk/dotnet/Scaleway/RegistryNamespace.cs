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
    /// Creates and manages Scaleway Container Registry.
    /// For more information see [the documentation](https://developers.scaleway.com/en/products/registry/api/).
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
    ///     var main = new Scaleway.RegistryNamespace("main", new()
    ///     {
    ///         Description = "Main container registry",
    ///         IsPublic = false,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Namespaces can be imported using the `{region}/{id}`, e.g. bash
    /// 
    /// ```sh
    ///  $ pulumi import scaleway:index/registryNamespace:RegistryNamespace main fr-par/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/registryNamespace:RegistryNamespace")]
    public partial class RegistryNamespace : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The description of the namespace.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Endpoint reachable by Docker.
        /// </summary>
        [Output("endpoint")]
        public Output<string> Endpoint { get; private set; } = null!;

        /// <summary>
        /// Whether the images stored in the namespace should be downloadable publicly (docker pull).
        /// </summary>
        [Output("isPublic")]
        public Output<bool?> IsPublic { get; private set; } = null!;

        /// <summary>
        /// The unique name of the namespace.
        /// 
        /// &gt; **Important** Updates to `name` will recreate the namespace.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The organization ID the namespace is associated with.
        /// </summary>
        [Output("organizationId")]
        public Output<string> OrganizationId { get; private set; } = null!;

        /// <summary>
        /// `project_id`) The ID of the project the namespace is associated with.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// `region`). The region in which the namespace should be created.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;


        /// <summary>
        /// Create a RegistryNamespace resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RegistryNamespace(string name, RegistryNamespaceArgs? args = null, CustomResourceOptions? options = null)
            : base("scaleway:index/registryNamespace:RegistryNamespace", name, args ?? new RegistryNamespaceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RegistryNamespace(string name, Input<string> id, RegistryNamespaceState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/registryNamespace:RegistryNamespace", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RegistryNamespace resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RegistryNamespace Get(string name, Input<string> id, RegistryNamespaceState? state = null, CustomResourceOptions? options = null)
        {
            return new RegistryNamespace(name, id, state, options);
        }
    }

    public sealed class RegistryNamespaceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the namespace.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether the images stored in the namespace should be downloadable publicly (docker pull).
        /// </summary>
        [Input("isPublic")]
        public Input<bool>? IsPublic { get; set; }

        /// <summary>
        /// The unique name of the namespace.
        /// 
        /// &gt; **Important** Updates to `name` will recreate the namespace.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the namespace is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`). The region in which the namespace should be created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public RegistryNamespaceArgs()
        {
        }
        public static new RegistryNamespaceArgs Empty => new RegistryNamespaceArgs();
    }

    public sealed class RegistryNamespaceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the namespace.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Endpoint reachable by Docker.
        /// </summary>
        [Input("endpoint")]
        public Input<string>? Endpoint { get; set; }

        /// <summary>
        /// Whether the images stored in the namespace should be downloadable publicly (docker pull).
        /// </summary>
        [Input("isPublic")]
        public Input<bool>? IsPublic { get; set; }

        /// <summary>
        /// The unique name of the namespace.
        /// 
        /// &gt; **Important** Updates to `name` will recreate the namespace.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The organization ID the namespace is associated with.
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the namespace is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`). The region in which the namespace should be created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public RegistryNamespaceState()
        {
        }
        public static new RegistryNamespaceState Empty => new RegistryNamespaceState();
    }
}
