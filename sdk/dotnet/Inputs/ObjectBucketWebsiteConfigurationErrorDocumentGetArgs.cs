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

    public sealed class ObjectBucketWebsiteConfigurationErrorDocumentGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The object key name to use when a 4XX class error occurs.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        public ObjectBucketWebsiteConfigurationErrorDocumentGetArgs()
        {
        }
        public static new ObjectBucketWebsiteConfigurationErrorDocumentGetArgs Empty => new ObjectBucketWebsiteConfigurationErrorDocumentGetArgs();
    }
}
