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
    /// Creates and manages Scaleway IAM Policies. For more information, see [the documentation](https://developers.scaleway.com/en/products/iam/api/v1alpha1/#policies-54b8a7).
    /// 
    /// &gt; You can find a detailed list of all permission sets available at Scaleway in the permission sets [reference page](https://www.scaleway.com/en/docs/identity-and-access-management/iam/reference-content/permission-sets/).
    /// 
    /// ## Example Usage
    /// ### Create a policy for an organization's project
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Lbrlabs.PulumiPackage.Scaleway;
    /// using Scaleway = Pulumi.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = Scaleway.GetAccountProject.Invoke(new()
    ///     {
    ///         Name = "default",
    ///     });
    /// 
    ///     var app = new Scaleway.IamApplication("app");
    /// 
    ///     var objectReadOnly = new Scaleway.IamPolicy("objectReadOnly", new()
    ///     {
    ///         Description = "gives app readonly access to object storage in project",
    ///         ApplicationId = app.Id,
    ///         Rules = new[]
    ///         {
    ///             new Scaleway.Inputs.IamPolicyRuleArgs
    ///             {
    ///                 ProjectIds = new[]
    ///                 {
    ///                     @default.Apply(@default =&gt; @default.Apply(getAccountProjectResult =&gt; getAccountProjectResult.Id)),
    ///                 },
    ///                 PermissionSetNames = new[]
    ///                 {
    ///                     "ObjectStorageReadOnly",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Policies can be imported using the `{id}`, e.g. bash
    /// 
    /// ```sh
    ///  $ pulumi import scaleway:index/iamPolicy:IamPolicy main 11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/iamPolicy:IamPolicy")]
    public partial class IamPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ID of the Application the policy will be linked to
        /// </summary>
        [Output("applicationId")]
        public Output<string?> ApplicationId { get; private set; } = null!;

        /// <summary>
        /// The date and time of the creation of the policy.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The description of the iam policy.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Whether the policy is editable.
        /// </summary>
        [Output("editable")]
        public Output<bool> Editable { get; private set; } = null!;

        /// <summary>
        /// ID of the Group the policy will be linked to
        /// </summary>
        [Output("groupId")]
        public Output<string?> GroupId { get; private set; } = null!;

        /// <summary>
        /// The name of the iam policy.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// If the policy doesn't apply to a principal.
        /// 
        /// &gt; **Important** Only one of `user_id`, `group_id`, `application_id` and `no_principal`  may be set.
        /// </summary>
        [Output("noPrincipal")]
        public Output<bool?> NoPrincipal { get; private set; } = null!;

        /// <summary>
        /// ID of organization scoped to the rule.
        /// </summary>
        [Output("organizationId")]
        public Output<string> OrganizationId { get; private set; } = null!;

        /// <summary>
        /// List of rules in the policy.
        /// </summary>
        [Output("rules")]
        public Output<ImmutableArray<Outputs.IamPolicyRule>> Rules { get; private set; } = null!;

        /// <summary>
        /// The tags associated with the iam policy.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// The date and time of the last update of the policy.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// ID of the User the policy will be linked to
        /// </summary>
        [Output("userId")]
        public Output<string?> UserId { get; private set; } = null!;


        /// <summary>
        /// Create a IamPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IamPolicy(string name, IamPolicyArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/iamPolicy:IamPolicy", name, args ?? new IamPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IamPolicy(string name, Input<string> id, IamPolicyState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/iamPolicy:IamPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IamPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IamPolicy Get(string name, Input<string> id, IamPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new IamPolicy(name, id, state, options);
        }
    }

    public sealed class IamPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the Application the policy will be linked to
        /// </summary>
        [Input("applicationId")]
        public Input<string>? ApplicationId { get; set; }

        /// <summary>
        /// The description of the iam policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// ID of the Group the policy will be linked to
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        /// <summary>
        /// The name of the iam policy.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// If the policy doesn't apply to a principal.
        /// 
        /// &gt; **Important** Only one of `user_id`, `group_id`, `application_id` and `no_principal`  may be set.
        /// </summary>
        [Input("noPrincipal")]
        public Input<bool>? NoPrincipal { get; set; }

        /// <summary>
        /// ID of organization scoped to the rule.
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        [Input("rules", required: true)]
        private InputList<Inputs.IamPolicyRuleArgs>? _rules;

        /// <summary>
        /// List of rules in the policy.
        /// </summary>
        public InputList<Inputs.IamPolicyRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.IamPolicyRuleArgs>());
            set => _rules = value;
        }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The tags associated with the iam policy.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// ID of the User the policy will be linked to
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        public IamPolicyArgs()
        {
        }
        public static new IamPolicyArgs Empty => new IamPolicyArgs();
    }

    public sealed class IamPolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the Application the policy will be linked to
        /// </summary>
        [Input("applicationId")]
        public Input<string>? ApplicationId { get; set; }

        /// <summary>
        /// The date and time of the creation of the policy.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// The description of the iam policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether the policy is editable.
        /// </summary>
        [Input("editable")]
        public Input<bool>? Editable { get; set; }

        /// <summary>
        /// ID of the Group the policy will be linked to
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        /// <summary>
        /// The name of the iam policy.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// If the policy doesn't apply to a principal.
        /// 
        /// &gt; **Important** Only one of `user_id`, `group_id`, `application_id` and `no_principal`  may be set.
        /// </summary>
        [Input("noPrincipal")]
        public Input<bool>? NoPrincipal { get; set; }

        /// <summary>
        /// ID of organization scoped to the rule.
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        [Input("rules")]
        private InputList<Inputs.IamPolicyRuleGetArgs>? _rules;

        /// <summary>
        /// List of rules in the policy.
        /// </summary>
        public InputList<Inputs.IamPolicyRuleGetArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.IamPolicyRuleGetArgs>());
            set => _rules = value;
        }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The tags associated with the iam policy.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The date and time of the last update of the policy.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        /// <summary>
        /// ID of the User the policy will be linked to
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        public IamPolicyState()
        {
        }
        public static new IamPolicyState Empty => new IamPolicyState();
    }
}
