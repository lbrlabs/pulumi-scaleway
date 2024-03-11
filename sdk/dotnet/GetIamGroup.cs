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
    public static class GetIamGroup
    {
        /// <summary>
        /// Gets information about an existing IAM group. For more information, please
        /// check [the documentation](https://developers.scaleway.com/en/products/iam/api/v1alpha1/#applications-83ce5e)
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var findByName = Scaleway.GetIamGroup.Invoke(new()
        ///     {
        ///         Name = "foobar",
        ///     });
        /// 
        ///     var findById = Scaleway.GetIamGroup.Invoke(new()
        ///     {
        ///         GroupId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetIamGroupResult> InvokeAsync(GetIamGroupArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIamGroupResult>("scaleway:index/getIamGroup:getIamGroup", args ?? new GetIamGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about an existing IAM group. For more information, please
        /// check [the documentation](https://developers.scaleway.com/en/products/iam/api/v1alpha1/#applications-83ce5e)
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var findByName = Scaleway.GetIamGroup.Invoke(new()
        ///     {
        ///         Name = "foobar",
        ///     });
        /// 
        ///     var findById = Scaleway.GetIamGroup.Invoke(new()
        ///     {
        ///         GroupId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetIamGroupResult> Invoke(GetIamGroupInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIamGroupResult>("scaleway:index/getIamGroup:getIamGroup", args ?? new GetIamGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIamGroupArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the IAM group.
        /// Only one of the `name` and `group_id` should be specified.
        /// </summary>
        [Input("groupId")]
        public string? GroupId { get; set; }

        /// <summary>
        /// The name of the IAM group.
        /// Only one of the `name` and `group_id` should be specified.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// `organization_id`) The ID of the
        /// organization the group is associated with.
        /// </summary>
        [Input("organizationId")]
        public string? OrganizationId { get; set; }

        public GetIamGroupArgs()
        {
        }
        public static new GetIamGroupArgs Empty => new GetIamGroupArgs();
    }

    public sealed class GetIamGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the IAM group.
        /// Only one of the `name` and `group_id` should be specified.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        /// <summary>
        /// The name of the IAM group.
        /// Only one of the `name` and `group_id` should be specified.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// `organization_id`) The ID of the
        /// organization the group is associated with.
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        public GetIamGroupInvokeArgs()
        {
        }
        public static new GetIamGroupInvokeArgs Empty => new GetIamGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetIamGroupResult
    {
        public readonly ImmutableArray<string> ApplicationIds;
        public readonly string CreatedAt;
        public readonly string Description;
        public readonly bool ExternalMembership;
        public readonly string? GroupId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Name;
        public readonly string? OrganizationId;
        public readonly string UpdatedAt;
        public readonly ImmutableArray<string> UserIds;

        [OutputConstructor]
        private GetIamGroupResult(
            ImmutableArray<string> applicationIds,

            string createdAt,

            string description,

            bool externalMembership,

            string? groupId,

            string id,

            string? name,

            string? organizationId,

            string updatedAt,

            ImmutableArray<string> userIds)
        {
            ApplicationIds = applicationIds;
            CreatedAt = createdAt;
            Description = description;
            ExternalMembership = externalMembership;
            GroupId = groupId;
            Id = id;
            Name = name;
            OrganizationId = organizationId;
            UpdatedAt = updatedAt;
            UserIds = userIds;
        }
    }
}
