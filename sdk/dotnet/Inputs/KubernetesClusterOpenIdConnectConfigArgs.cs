// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway.Inputs
{

    public sealed class KubernetesClusterOpenIdConnectConfigArgs : Pulumi.ResourceArgs
    {
        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        [Input("groupsClaims")]
        private InputList<string>? _groupsClaims;
        public InputList<string> GroupsClaims
        {
            get => _groupsClaims ?? (_groupsClaims = new InputList<string>());
            set => _groupsClaims = value;
        }

        [Input("groupsPrefix")]
        public Input<string>? GroupsPrefix { get; set; }

        [Input("issuerUrl", required: true)]
        public Input<string> IssuerUrl { get; set; } = null!;

        [Input("requiredClaims")]
        private InputList<string>? _requiredClaims;
        public InputList<string> RequiredClaims
        {
            get => _requiredClaims ?? (_requiredClaims = new InputList<string>());
            set => _requiredClaims = value;
        }

        [Input("usernameClaim")]
        public Input<string>? UsernameClaim { get; set; }

        [Input("usernamePrefix")]
        public Input<string>? UsernamePrefix { get; set; }

        public KubernetesClusterOpenIdConnectConfigArgs()
        {
        }
    }
}
