// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Inputs
{

    public sealed class LoadbalancerFrontendAclArgs : Pulumi.ResourceArgs
    {
        [Input("action", required: true)]
        public Input<Inputs.LoadbalancerFrontendAclActionArgs> Action { get; set; } = null!;

        [Input("match", required: true)]
        public Input<Inputs.LoadbalancerFrontendAclMatchArgs> Match { get; set; } = null!;

        [Input("name")]
        public Input<string>? Name { get; set; }

        public LoadbalancerFrontendAclArgs()
        {
        }
    }
}
