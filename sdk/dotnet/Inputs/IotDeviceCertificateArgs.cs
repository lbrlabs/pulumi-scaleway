// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Inputs
{

    public sealed class IotDeviceCertificateArgs : Pulumi.ResourceArgs
    {
        [Input("crt")]
        public Input<string>? Crt { get; set; }

        /// <summary>
        /// The private key of the device, in case it is generated by Scaleway.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        public IotDeviceCertificateArgs()
        {
        }
    }
}
