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
    public sealed class LoadbalancerAclMatch
    {
        /// <summary>
        /// The HTTP filter to match. This filter is supported only if your backend protocol has an HTTP forward protocol.
        /// It extracts the request's URL path, which starts at the first slash and ends before the question mark (without the host part).
        /// Possible values are: `acl_http_filter_none`, `path_begin`, `path_end`, `http_header_match` or `regex`.
        /// </summary>
        public readonly string? HttpFilter;
        /// <summary>
        /// You can use this field with http_header_match acl type to set the header name to filter
        /// </summary>
        public readonly string? HttpFilterOption;
        /// <summary>
        /// A list of possible values to match for the given HTTP filter.
        /// Keep in mind that in the case of `http_header_match` the HTTP header field name is case-insensitive.
        /// </summary>
        public readonly ImmutableArray<string> HttpFilterValues;
        /// <summary>
        /// If set to `true`, the condition will be of type "unless".
        /// </summary>
        public readonly bool? Invert;
        /// <summary>
        /// A list of IPs or CIDR v4/v6 addresses of the client of the session to match.
        /// </summary>
        public readonly ImmutableArray<string> IpSubnets;

        [OutputConstructor]
        private LoadbalancerAclMatch(
            string? httpFilter,

            string? httpFilterOption,

            ImmutableArray<string> httpFilterValues,

            bool? invert,

            ImmutableArray<string> ipSubnets)
        {
            HttpFilter = httpFilter;
            HttpFilterOption = httpFilterOption;
            HttpFilterValues = httpFilterValues;
            Invert = invert;
            IpSubnets = ipSubnets;
        }
    }
}
