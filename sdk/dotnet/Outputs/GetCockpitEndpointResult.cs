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
    public sealed class GetCockpitEndpointResult
    {
        /// <summary>
        /// The alertmanager URL
        /// </summary>
        public readonly string AlertmanagerUrl;
        /// <summary>
        /// The grafana URL
        /// </summary>
        public readonly string GrafanaUrl;
        /// <summary>
        /// The logs URL
        /// </summary>
        public readonly string LogsUrl;
        /// <summary>
        /// The metrics URL
        /// </summary>
        public readonly string MetricsUrl;
        /// <summary>
        /// The traces URL
        /// </summary>
        public readonly string TracesUrl;

        [OutputConstructor]
        private GetCockpitEndpointResult(
            string alertmanagerUrl,

            string grafanaUrl,

            string logsUrl,

            string metricsUrl,

            string tracesUrl)
        {
            AlertmanagerUrl = alertmanagerUrl;
            GrafanaUrl = grafanaUrl;
            LogsUrl = logsUrl;
            MetricsUrl = metricsUrl;
            TracesUrl = tracesUrl;
        }
    }
}
