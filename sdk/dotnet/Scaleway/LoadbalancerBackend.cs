// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Scaleway
{
    /// <summary>
    /// Creates and manages Scaleway Load-Balancer Backends.
    /// For more information, see [the documentation](https://developers.scaleway.com/en/products/lb/zoned_api).
    /// 
    /// ## Examples
    /// 
    /// ### Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Scaleway = Lbrlabs.PulumiPackage.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var backend01 = new Scaleway.LoadbalancerBackend("backend01", new()
    ///     {
    ///         LbId = scaleway_lb.Lb01.Id,
    ///         ForwardProtocol = "http",
    ///         ForwardPort = 80,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### With HTTP Health Check
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Scaleway = Lbrlabs.PulumiPackage.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var backend01 = new Scaleway.LoadbalancerBackend("backend01", new()
    ///     {
    ///         LbId = scaleway_lb.Lb01.Id,
    ///         ForwardProtocol = "http",
    ///         ForwardPort = 80,
    ///         HealthCheckHttp = new Scaleway.Inputs.LoadbalancerBackendHealthCheckHttpArgs
    ///         {
    ///             Uri = "www.test.com/health",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Load-Balancer backend can be imported using the `{zone}/{id}`, e.g. bash
    /// 
    /// ```sh
    ///  $ pulumi import scaleway:index/loadbalancerBackend:LoadbalancerBackend backend01 fr-par-1/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/loadbalancerBackend:LoadbalancerBackend")]
    public partial class LoadbalancerBackend : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Scaleway S3 bucket website to be served in case all backend servers are down.
        /// &gt; **Note:** Only the host part of the Scaleway S3 bucket website is expected:
        /// e.g. 'failover-website.s3-website.fr-par.scw.cloud' if your bucket website URL is 'https://failover-website.s3-website.fr-par.scw.cloud/'.
        /// </summary>
        [Output("failoverHost")]
        public Output<string?> FailoverHost { get; private set; } = null!;

        /// <summary>
        /// User sessions will be forwarded to this port of backend servers.
        /// </summary>
        [Output("forwardPort")]
        public Output<int> ForwardPort { get; private set; } = null!;

        /// <summary>
        /// Load balancing algorithm. Possible values are: `roundrobin`, `leastconn` and `first`.
        /// </summary>
        [Output("forwardPortAlgorithm")]
        public Output<string?> ForwardPortAlgorithm { get; private set; } = null!;

        /// <summary>
        /// Backend protocol. Possible values are: `tcp` or `http`.
        /// </summary>
        [Output("forwardProtocol")]
        public Output<string> ForwardProtocol { get; private set; } = null!;

        /// <summary>
        /// Interval between two HC requests.
        /// </summary>
        [Output("healthCheckDelay")]
        public Output<string?> HealthCheckDelay { get; private set; } = null!;

        /// <summary>
        /// This block enable HTTP health check. Only one of `health_check_tcp`, `health_check_http` and `health_check_https` should be specified.
        /// </summary>
        [Output("healthCheckHttp")]
        public Output<Outputs.LoadbalancerBackendHealthCheckHttp?> HealthCheckHttp { get; private set; } = null!;

        /// <summary>
        /// This block enable HTTPS health check. Only one of `health_check_tcp`, `health_check_http` and `health_check_https` should be specified.
        /// </summary>
        [Output("healthCheckHttps")]
        public Output<Outputs.LoadbalancerBackendHealthCheckHttps?> HealthCheckHttps { get; private set; } = null!;

        /// <summary>
        /// Number of allowed failed HC requests before the backend server is marked down.
        /// </summary>
        [Output("healthCheckMaxRetries")]
        public Output<int?> HealthCheckMaxRetries { get; private set; } = null!;

        /// <summary>
        /// Port the HC requests will be send to.
        /// </summary>
        [Output("healthCheckPort")]
        public Output<int> HealthCheckPort { get; private set; } = null!;

        /// <summary>
        /// This block enable TCP health check. Only one of `health_check_tcp`, `health_check_http` and `health_check_https` should be specified.
        /// </summary>
        [Output("healthCheckTcp")]
        public Output<Outputs.LoadbalancerBackendHealthCheckTcp> HealthCheckTcp { get; private set; } = null!;

        /// <summary>
        /// Timeout before we consider a HC request failed.
        /// </summary>
        [Output("healthCheckTimeout")]
        public Output<string?> HealthCheckTimeout { get; private set; } = null!;

        /// <summary>
        /// Specifies whether the Load Balancer should check the backend server’s certificate before initiating a connection.
        /// </summary>
        [Output("ignoreSslServerVerify")]
        public Output<bool?> IgnoreSslServerVerify { get; private set; } = null!;

        /// <summary>
        /// The load-balancer ID this backend is attached to.
        /// &gt; **Important:** Updates to `lb_id` will recreate the backend.
        /// </summary>
        [Output("lbId")]
        public Output<string> LbId { get; private set; } = null!;

        /// <summary>
        /// The name of the load-balancer backend.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Modify what occurs when a backend server is marked down. Possible values are: `none` and `shutdown_sessions`.
        /// </summary>
        [Output("onMarkedDownAction")]
        public Output<string?> OnMarkedDownAction { get; private set; } = null!;

        /// <summary>
        /// Choose the type of PROXY protocol to enable (`none`, `v1`, `v2`, `v2_ssl`, `v2_ssl_cn`)
        /// </summary>
        [Output("proxyProtocol")]
        public Output<string?> ProxyProtocol { get; private set; } = null!;

        /// <summary>
        /// DEPRECATED please use `proxy_protocol` instead - (Default: `false`) Enables PROXY protocol version 2.
        /// </summary>
        [Output("sendProxyV2")]
        public Output<bool?> SendProxyV2 { get; private set; } = null!;

        /// <summary>
        /// List of backend server IP addresses. Addresses can be either IPv4 or IPv6.
        /// </summary>
        [Output("serverIps")]
        public Output<ImmutableArray<string>> ServerIps { get; private set; } = null!;

        /// <summary>
        /// Enables SSL between load balancer and backend servers.
        /// </summary>
        [Output("sslBridging")]
        public Output<bool?> SslBridging { get; private set; } = null!;

        /// <summary>
        /// Load balancing algorithm. Possible values are: `none`, `cookie` and `table`.
        /// </summary>
        [Output("stickySessions")]
        public Output<string?> StickySessions { get; private set; } = null!;

        /// <summary>
        /// Cookie name for for sticky sessions. Only applicable when sticky_sessions is set to `cookie`.
        /// </summary>
        [Output("stickySessionsCookieName")]
        public Output<string?> StickySessionsCookieName { get; private set; } = null!;

        /// <summary>
        /// Maximum initial server connection establishment time. (e.g.: `1s`)
        /// </summary>
        [Output("timeoutConnect")]
        public Output<string?> TimeoutConnect { get; private set; } = null!;

        /// <summary>
        /// Maximum server connection inactivity time. (e.g.: `1s`)
        /// </summary>
        [Output("timeoutServer")]
        public Output<string?> TimeoutServer { get; private set; } = null!;

        /// <summary>
        /// Maximum tunnel inactivity time. (e.g.: `1s`)
        /// </summary>
        [Output("timeoutTunnel")]
        public Output<string?> TimeoutTunnel { get; private set; } = null!;


        /// <summary>
        /// Create a LoadbalancerBackend resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LoadbalancerBackend(string name, LoadbalancerBackendArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/loadbalancerBackend:LoadbalancerBackend", name, args ?? new LoadbalancerBackendArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LoadbalancerBackend(string name, Input<string> id, LoadbalancerBackendState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/loadbalancerBackend:LoadbalancerBackend", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/lbrlabs",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing LoadbalancerBackend resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LoadbalancerBackend Get(string name, Input<string> id, LoadbalancerBackendState? state = null, CustomResourceOptions? options = null)
        {
            return new LoadbalancerBackend(name, id, state, options);
        }
    }

    public sealed class LoadbalancerBackendArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Scaleway S3 bucket website to be served in case all backend servers are down.
        /// &gt; **Note:** Only the host part of the Scaleway S3 bucket website is expected:
        /// e.g. 'failover-website.s3-website.fr-par.scw.cloud' if your bucket website URL is 'https://failover-website.s3-website.fr-par.scw.cloud/'.
        /// </summary>
        [Input("failoverHost")]
        public Input<string>? FailoverHost { get; set; }

        /// <summary>
        /// User sessions will be forwarded to this port of backend servers.
        /// </summary>
        [Input("forwardPort", required: true)]
        public Input<int> ForwardPort { get; set; } = null!;

        /// <summary>
        /// Load balancing algorithm. Possible values are: `roundrobin`, `leastconn` and `first`.
        /// </summary>
        [Input("forwardPortAlgorithm")]
        public Input<string>? ForwardPortAlgorithm { get; set; }

        /// <summary>
        /// Backend protocol. Possible values are: `tcp` or `http`.
        /// </summary>
        [Input("forwardProtocol", required: true)]
        public Input<string> ForwardProtocol { get; set; } = null!;

        /// <summary>
        /// Interval between two HC requests.
        /// </summary>
        [Input("healthCheckDelay")]
        public Input<string>? HealthCheckDelay { get; set; }

        /// <summary>
        /// This block enable HTTP health check. Only one of `health_check_tcp`, `health_check_http` and `health_check_https` should be specified.
        /// </summary>
        [Input("healthCheckHttp")]
        public Input<Inputs.LoadbalancerBackendHealthCheckHttpArgs>? HealthCheckHttp { get; set; }

        /// <summary>
        /// This block enable HTTPS health check. Only one of `health_check_tcp`, `health_check_http` and `health_check_https` should be specified.
        /// </summary>
        [Input("healthCheckHttps")]
        public Input<Inputs.LoadbalancerBackendHealthCheckHttpsArgs>? HealthCheckHttps { get; set; }

        /// <summary>
        /// Number of allowed failed HC requests before the backend server is marked down.
        /// </summary>
        [Input("healthCheckMaxRetries")]
        public Input<int>? HealthCheckMaxRetries { get; set; }

        /// <summary>
        /// Port the HC requests will be send to.
        /// </summary>
        [Input("healthCheckPort")]
        public Input<int>? HealthCheckPort { get; set; }

        /// <summary>
        /// This block enable TCP health check. Only one of `health_check_tcp`, `health_check_http` and `health_check_https` should be specified.
        /// </summary>
        [Input("healthCheckTcp")]
        public Input<Inputs.LoadbalancerBackendHealthCheckTcpArgs>? HealthCheckTcp { get; set; }

        /// <summary>
        /// Timeout before we consider a HC request failed.
        /// </summary>
        [Input("healthCheckTimeout")]
        public Input<string>? HealthCheckTimeout { get; set; }

        /// <summary>
        /// Specifies whether the Load Balancer should check the backend server’s certificate before initiating a connection.
        /// </summary>
        [Input("ignoreSslServerVerify")]
        public Input<bool>? IgnoreSslServerVerify { get; set; }

        /// <summary>
        /// The load-balancer ID this backend is attached to.
        /// &gt; **Important:** Updates to `lb_id` will recreate the backend.
        /// </summary>
        [Input("lbId", required: true)]
        public Input<string> LbId { get; set; } = null!;

        /// <summary>
        /// The name of the load-balancer backend.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Modify what occurs when a backend server is marked down. Possible values are: `none` and `shutdown_sessions`.
        /// </summary>
        [Input("onMarkedDownAction")]
        public Input<string>? OnMarkedDownAction { get; set; }

        /// <summary>
        /// Choose the type of PROXY protocol to enable (`none`, `v1`, `v2`, `v2_ssl`, `v2_ssl_cn`)
        /// </summary>
        [Input("proxyProtocol")]
        public Input<string>? ProxyProtocol { get; set; }

        /// <summary>
        /// DEPRECATED please use `proxy_protocol` instead - (Default: `false`) Enables PROXY protocol version 2.
        /// </summary>
        [Input("sendProxyV2")]
        public Input<bool>? SendProxyV2 { get; set; }

        [Input("serverIps")]
        private InputList<string>? _serverIps;

        /// <summary>
        /// List of backend server IP addresses. Addresses can be either IPv4 or IPv6.
        /// </summary>
        public InputList<string> ServerIps
        {
            get => _serverIps ?? (_serverIps = new InputList<string>());
            set => _serverIps = value;
        }

        /// <summary>
        /// Enables SSL between load balancer and backend servers.
        /// </summary>
        [Input("sslBridging")]
        public Input<bool>? SslBridging { get; set; }

        /// <summary>
        /// Load balancing algorithm. Possible values are: `none`, `cookie` and `table`.
        /// </summary>
        [Input("stickySessions")]
        public Input<string>? StickySessions { get; set; }

        /// <summary>
        /// Cookie name for for sticky sessions. Only applicable when sticky_sessions is set to `cookie`.
        /// </summary>
        [Input("stickySessionsCookieName")]
        public Input<string>? StickySessionsCookieName { get; set; }

        /// <summary>
        /// Maximum initial server connection establishment time. (e.g.: `1s`)
        /// </summary>
        [Input("timeoutConnect")]
        public Input<string>? TimeoutConnect { get; set; }

        /// <summary>
        /// Maximum server connection inactivity time. (e.g.: `1s`)
        /// </summary>
        [Input("timeoutServer")]
        public Input<string>? TimeoutServer { get; set; }

        /// <summary>
        /// Maximum tunnel inactivity time. (e.g.: `1s`)
        /// </summary>
        [Input("timeoutTunnel")]
        public Input<string>? TimeoutTunnel { get; set; }

        public LoadbalancerBackendArgs()
        {
        }
        public static new LoadbalancerBackendArgs Empty => new LoadbalancerBackendArgs();
    }

    public sealed class LoadbalancerBackendState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Scaleway S3 bucket website to be served in case all backend servers are down.
        /// &gt; **Note:** Only the host part of the Scaleway S3 bucket website is expected:
        /// e.g. 'failover-website.s3-website.fr-par.scw.cloud' if your bucket website URL is 'https://failover-website.s3-website.fr-par.scw.cloud/'.
        /// </summary>
        [Input("failoverHost")]
        public Input<string>? FailoverHost { get; set; }

        /// <summary>
        /// User sessions will be forwarded to this port of backend servers.
        /// </summary>
        [Input("forwardPort")]
        public Input<int>? ForwardPort { get; set; }

        /// <summary>
        /// Load balancing algorithm. Possible values are: `roundrobin`, `leastconn` and `first`.
        /// </summary>
        [Input("forwardPortAlgorithm")]
        public Input<string>? ForwardPortAlgorithm { get; set; }

        /// <summary>
        /// Backend protocol. Possible values are: `tcp` or `http`.
        /// </summary>
        [Input("forwardProtocol")]
        public Input<string>? ForwardProtocol { get; set; }

        /// <summary>
        /// Interval between two HC requests.
        /// </summary>
        [Input("healthCheckDelay")]
        public Input<string>? HealthCheckDelay { get; set; }

        /// <summary>
        /// This block enable HTTP health check. Only one of `health_check_tcp`, `health_check_http` and `health_check_https` should be specified.
        /// </summary>
        [Input("healthCheckHttp")]
        public Input<Inputs.LoadbalancerBackendHealthCheckHttpGetArgs>? HealthCheckHttp { get; set; }

        /// <summary>
        /// This block enable HTTPS health check. Only one of `health_check_tcp`, `health_check_http` and `health_check_https` should be specified.
        /// </summary>
        [Input("healthCheckHttps")]
        public Input<Inputs.LoadbalancerBackendHealthCheckHttpsGetArgs>? HealthCheckHttps { get; set; }

        /// <summary>
        /// Number of allowed failed HC requests before the backend server is marked down.
        /// </summary>
        [Input("healthCheckMaxRetries")]
        public Input<int>? HealthCheckMaxRetries { get; set; }

        /// <summary>
        /// Port the HC requests will be send to.
        /// </summary>
        [Input("healthCheckPort")]
        public Input<int>? HealthCheckPort { get; set; }

        /// <summary>
        /// This block enable TCP health check. Only one of `health_check_tcp`, `health_check_http` and `health_check_https` should be specified.
        /// </summary>
        [Input("healthCheckTcp")]
        public Input<Inputs.LoadbalancerBackendHealthCheckTcpGetArgs>? HealthCheckTcp { get; set; }

        /// <summary>
        /// Timeout before we consider a HC request failed.
        /// </summary>
        [Input("healthCheckTimeout")]
        public Input<string>? HealthCheckTimeout { get; set; }

        /// <summary>
        /// Specifies whether the Load Balancer should check the backend server’s certificate before initiating a connection.
        /// </summary>
        [Input("ignoreSslServerVerify")]
        public Input<bool>? IgnoreSslServerVerify { get; set; }

        /// <summary>
        /// The load-balancer ID this backend is attached to.
        /// &gt; **Important:** Updates to `lb_id` will recreate the backend.
        /// </summary>
        [Input("lbId")]
        public Input<string>? LbId { get; set; }

        /// <summary>
        /// The name of the load-balancer backend.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Modify what occurs when a backend server is marked down. Possible values are: `none` and `shutdown_sessions`.
        /// </summary>
        [Input("onMarkedDownAction")]
        public Input<string>? OnMarkedDownAction { get; set; }

        /// <summary>
        /// Choose the type of PROXY protocol to enable (`none`, `v1`, `v2`, `v2_ssl`, `v2_ssl_cn`)
        /// </summary>
        [Input("proxyProtocol")]
        public Input<string>? ProxyProtocol { get; set; }

        /// <summary>
        /// DEPRECATED please use `proxy_protocol` instead - (Default: `false`) Enables PROXY protocol version 2.
        /// </summary>
        [Input("sendProxyV2")]
        public Input<bool>? SendProxyV2 { get; set; }

        [Input("serverIps")]
        private InputList<string>? _serverIps;

        /// <summary>
        /// List of backend server IP addresses. Addresses can be either IPv4 or IPv6.
        /// </summary>
        public InputList<string> ServerIps
        {
            get => _serverIps ?? (_serverIps = new InputList<string>());
            set => _serverIps = value;
        }

        /// <summary>
        /// Enables SSL between load balancer and backend servers.
        /// </summary>
        [Input("sslBridging")]
        public Input<bool>? SslBridging { get; set; }

        /// <summary>
        /// Load balancing algorithm. Possible values are: `none`, `cookie` and `table`.
        /// </summary>
        [Input("stickySessions")]
        public Input<string>? StickySessions { get; set; }

        /// <summary>
        /// Cookie name for for sticky sessions. Only applicable when sticky_sessions is set to `cookie`.
        /// </summary>
        [Input("stickySessionsCookieName")]
        public Input<string>? StickySessionsCookieName { get; set; }

        /// <summary>
        /// Maximum initial server connection establishment time. (e.g.: `1s`)
        /// </summary>
        [Input("timeoutConnect")]
        public Input<string>? TimeoutConnect { get; set; }

        /// <summary>
        /// Maximum server connection inactivity time. (e.g.: `1s`)
        /// </summary>
        [Input("timeoutServer")]
        public Input<string>? TimeoutServer { get; set; }

        /// <summary>
        /// Maximum tunnel inactivity time. (e.g.: `1s`)
        /// </summary>
        [Input("timeoutTunnel")]
        public Input<string>? TimeoutTunnel { get; set; }

        public LoadbalancerBackendState()
        {
        }
        public static new LoadbalancerBackendState Empty => new LoadbalancerBackendState();
    }
}
