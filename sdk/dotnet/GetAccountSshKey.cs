// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway
{
    public static class GetAccountSshKey
    {
        public static Task<GetAccountSshKeyResult> InvokeAsync(GetAccountSshKeyArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAccountSshKeyResult>("scaleway:index/getAccountSshKey:getAccountSshKey", args ?? new GetAccountSshKeyArgs(), options.WithDefaults());

        public static Output<GetAccountSshKeyResult> Invoke(GetAccountSshKeyInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAccountSshKeyResult>("scaleway:index/getAccountSshKey:getAccountSshKey", args ?? new GetAccountSshKeyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAccountSshKeyArgs : Pulumi.InvokeArgs
    {
        [Input("name")]
        public string? Name { get; set; }

        [Input("sshKeyId")]
        public string? SshKeyId { get; set; }

        public GetAccountSshKeyArgs()
        {
        }
    }

    public sealed class GetAccountSshKeyInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("sshKeyId")]
        public Input<string>? SshKeyId { get; set; }

        public GetAccountSshKeyInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAccountSshKeyResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Name;
        public readonly string OrganizationId;
        public readonly string ProjectId;
        public readonly string PublicKey;
        public readonly string? SshKeyId;

        [OutputConstructor]
        private GetAccountSshKeyResult(
            string id,

            string? name,

            string organizationId,

            string projectId,

            string publicKey,

            string? sshKeyId)
        {
            Id = id;
            Name = name;
            OrganizationId = organizationId;
            ProjectId = projectId;
            PublicKey = publicKey;
            SshKeyId = sshKeyId;
        }
    }
}
