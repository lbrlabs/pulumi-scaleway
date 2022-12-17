// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Scaleway.Inputs
{

    public sealed class IamPolicyRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of organization scoped to the rule.
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        [Input("permissionSetNames", required: true)]
        private InputList<string>? _permissionSetNames;

        /// <summary>
        /// Names of permission sets bound to the rule.
        /// </summary>
        public InputList<string> PermissionSetNames
        {
            get => _permissionSetNames ?? (_permissionSetNames = new InputList<string>());
            set => _permissionSetNames = value;
        }

        [Input("projectIds")]
        private InputList<string>? _projectIds;

        /// <summary>
        /// List of project IDs scoped to the rule.
        /// </summary>
        public InputList<string> ProjectIds
        {
            get => _projectIds ?? (_projectIds = new InputList<string>());
            set => _projectIds = value;
        }

        public IamPolicyRuleArgs()
        {
        }
        public static new IamPolicyRuleArgs Empty => new IamPolicyRuleArgs();
    }
}
