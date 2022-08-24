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
    [ScalewayResourceType("scaleway:index/accountSshKey:AccountSshKey")]
    public partial class AccountSshKey : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the SSH key
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The organization_id you want to attach the resource to
        /// </summary>
        [Output("organizationId")]
        public Output<string> OrganizationId { get; private set; } = null!;

        /// <summary>
        /// The project_id you want to attach the resource to
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The public SSH key
        /// </summary>
        [Output("publicKey")]
        public Output<string> PublicKey { get; private set; } = null!;


        /// <summary>
        /// Create a AccountSshKey resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AccountSshKey(string name, AccountSshKeyArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/accountSshKey:AccountSshKey", name, args ?? new AccountSshKeyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AccountSshKey(string name, Input<string> id, AccountSshKeyState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/accountSshKey:AccountSshKey", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AccountSshKey resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AccountSshKey Get(string name, Input<string> id, AccountSshKeyState? state = null, CustomResourceOptions? options = null)
        {
            return new AccountSshKey(name, id, state, options);
        }
    }

    public sealed class AccountSshKeyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the SSH key
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The project_id you want to attach the resource to
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The public SSH key
        /// </summary>
        [Input("publicKey", required: true)]
        public Input<string> PublicKey { get; set; } = null!;

        public AccountSshKeyArgs()
        {
        }
    }

    public sealed class AccountSshKeyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the SSH key
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The organization_id you want to attach the resource to
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        /// <summary>
        /// The project_id you want to attach the resource to
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The public SSH key
        /// </summary>
        [Input("publicKey")]
        public Input<string>? PublicKey { get; set; }

        public AccountSshKeyState()
        {
        }
    }
}
