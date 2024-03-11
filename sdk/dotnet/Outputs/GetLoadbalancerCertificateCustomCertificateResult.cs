// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway.Outputs
{

    [OutputType]
    public sealed class GetLoadbalancerCertificateCustomCertificateResult
    {
        /// <summary>
        /// The full PEM-formatted certificate chain
        /// </summary>
        public readonly string CertificateChain;

        [OutputConstructor]
        private GetLoadbalancerCertificateCustomCertificateResult(string certificateChain)
        {
            CertificateChain = certificateChain;
        }
    }
}
