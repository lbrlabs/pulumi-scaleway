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
    /// Creates and manages Scaleway IAM Groups.
    /// For more information, see [the documentation](https://developers.scaleway.com/en/products/iam/api/v1alpha1/#groups-f592eb).
    /// 
    /// ## Examples
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
    ///     var basic = new Scaleway.IamGroup("basic", new()
    ///     {
    ///         ApplicationIds = new[] {},
    ///         Description = "basic description",
    ///         UserIds = new[] {},
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ### With applications
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
    ///     var app = new Scaleway.IamApplication("app");
    /// 
    ///     var withApp = new Scaleway.IamGroup("withApp", new()
    ///     {
    ///         ApplicationIds = new[]
    ///         {
    ///             app.Id,
    ///         },
    ///         UserIds = new[] {},
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// IAM groups can be imported using the `{id}`, e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:index/iamGroup:IamGroup basic 11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/iamGroup:IamGroup")]
    public partial class IamGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The list of IDs of the applications attached to the group.
        /// </summary>
        [Output("applicationIds")]
        public Output<ImmutableArray<string>> ApplicationIds { get; private set; } = null!;

        /// <summary>
        /// The date and time of the creation of the group
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The description of the IAM group.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Manage membership externally. This make the resource ignore user_ids and application_ids. Should be used when using iam_group_membership
        /// </summary>
        [Output("externalMembership")]
        public Output<bool?> ExternalMembership { get; private set; } = null!;

        /// <summary>
        /// The name of the IAM group.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// `organization_id`) The ID of the organization the group is associated with.
        /// </summary>
        [Output("organizationId")]
        public Output<string> OrganizationId { get; private set; } = null!;

        /// <summary>
        /// The date and time of the last update of the group
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// The list of IDs of the users attached to the group.
        /// </summary>
        [Output("userIds")]
        public Output<ImmutableArray<string>> UserIds { get; private set; } = null!;


        /// <summary>
        /// Create a IamGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IamGroup(string name, IamGroupArgs? args = null, CustomResourceOptions? options = null)
            : base("scaleway:index/iamGroup:IamGroup", name, args ?? new IamGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IamGroup(string name, Input<string> id, IamGroupState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/iamGroup:IamGroup", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing IamGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IamGroup Get(string name, Input<string> id, IamGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new IamGroup(name, id, state, options);
        }
    }

    public sealed class IamGroupArgs : global::Pulumi.ResourceArgs
    {
        [Input("applicationIds")]
        private InputList<string>? _applicationIds;

        /// <summary>
        /// The list of IDs of the applications attached to the group.
        /// </summary>
        public InputList<string> ApplicationIds
        {
            get => _applicationIds ?? (_applicationIds = new InputList<string>());
            set => _applicationIds = value;
        }

        /// <summary>
        /// The description of the IAM group.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Manage membership externally. This make the resource ignore user_ids and application_ids. Should be used when using iam_group_membership
        /// </summary>
        [Input("externalMembership")]
        public Input<bool>? ExternalMembership { get; set; }

        /// <summary>
        /// The name of the IAM group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// `organization_id`) The ID of the organization the group is associated with.
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        [Input("userIds")]
        private InputList<string>? _userIds;

        /// <summary>
        /// The list of IDs of the users attached to the group.
        /// </summary>
        public InputList<string> UserIds
        {
            get => _userIds ?? (_userIds = new InputList<string>());
            set => _userIds = value;
        }

        public IamGroupArgs()
        {
        }
        public static new IamGroupArgs Empty => new IamGroupArgs();
    }

    public sealed class IamGroupState : global::Pulumi.ResourceArgs
    {
        [Input("applicationIds")]
        private InputList<string>? _applicationIds;

        /// <summary>
        /// The list of IDs of the applications attached to the group.
        /// </summary>
        public InputList<string> ApplicationIds
        {
            get => _applicationIds ?? (_applicationIds = new InputList<string>());
            set => _applicationIds = value;
        }

        /// <summary>
        /// The date and time of the creation of the group
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// The description of the IAM group.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Manage membership externally. This make the resource ignore user_ids and application_ids. Should be used when using iam_group_membership
        /// </summary>
        [Input("externalMembership")]
        public Input<bool>? ExternalMembership { get; set; }

        /// <summary>
        /// The name of the IAM group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// `organization_id`) The ID of the organization the group is associated with.
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        /// <summary>
        /// The date and time of the last update of the group
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        [Input("userIds")]
        private InputList<string>? _userIds;

        /// <summary>
        /// The list of IDs of the users attached to the group.
        /// </summary>
        public InputList<string> UserIds
        {
            get => _userIds ?? (_userIds = new InputList<string>());
            set => _userIds = value;
        }

        public IamGroupState()
        {
        }
        public static new IamGroupState Empty => new IamGroupState();
    }
}
