// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway
{
    [ScalewayResourceType("scaleway:index/loadbalancerBackend:LoadbalancerBackend")]
    public partial class LoadbalancerBackend : Pulumi.CustomResource
    {
        /// <summary>
        /// User sessions will be forwarded to this port of backend servers
        /// </summary>
        [Output("forwardPort")]
        public Output<int> ForwardPort { get; private set; } = null!;

        /// <summary>
        /// Load balancing algorithm
        /// </summary>
        [Output("forwardPortAlgorithm")]
        public Output<string?> ForwardPortAlgorithm { get; private set; } = null!;

        /// <summary>
        /// Backend protocol
        /// </summary>
        [Output("forwardProtocol")]
        public Output<string> ForwardProtocol { get; private set; } = null!;

        /// <summary>
        /// Interval between two HC requests
        /// </summary>
        [Output("healthCheckDelay")]
        public Output<string?> HealthCheckDelay { get; private set; } = null!;

        [Output("healthCheckHttp")]
        public Output<Outputs.LoadbalancerBackendHealthCheckHttp?> HealthCheckHttp { get; private set; } = null!;

        [Output("healthCheckHttps")]
        public Output<Outputs.LoadbalancerBackendHealthCheckHttps?> HealthCheckHttps { get; private set; } = null!;

        /// <summary>
        /// Number of allowed failed HC requests before the backend server is marked down
        /// </summary>
        [Output("healthCheckMaxRetries")]
        public Output<int?> HealthCheckMaxRetries { get; private set; } = null!;

        /// <summary>
        /// Port the HC requests will be send to. Default to `forward_port`
        /// </summary>
        [Output("healthCheckPort")]
        public Output<int> HealthCheckPort { get; private set; } = null!;

        [Output("healthCheckTcp")]
        public Output<Outputs.LoadbalancerBackendHealthCheckTcp> HealthCheckTcp { get; private set; } = null!;

        /// <summary>
        /// Timeout before we consider a HC request failed
        /// </summary>
        [Output("healthCheckTimeout")]
        public Output<string?> HealthCheckTimeout { get; private set; } = null!;

        /// <summary>
        /// The load-balancer ID
        /// </summary>
        [Output("lbId")]
        public Output<string> LbId { get; private set; } = null!;

        /// <summary>
        /// The name of the backend
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Modify what occurs when a backend server is marked down
        /// </summary>
        [Output("onMarkedDownAction")]
        public Output<string?> OnMarkedDownAction { get; private set; } = null!;

        /// <summary>
        /// Type of PROXY protocol to enable
        /// </summary>
        [Output("proxyProtocol")]
        public Output<string?> ProxyProtocol { get; private set; } = null!;

        /// <summary>
        /// Enables PROXY protocol version 2
        /// </summary>
        [Output("sendProxyV2")]
        public Output<bool?> SendProxyV2 { get; private set; } = null!;

        /// <summary>
        /// Backend server IP addresses list (IPv4 or IPv6)
        /// </summary>
        [Output("serverIps")]
        public Output<ImmutableArray<string>> ServerIps { get; private set; } = null!;

        /// <summary>
        /// Load balancing algorithm
        /// </summary>
        [Output("stickySessions")]
        public Output<string?> StickySessions { get; private set; } = null!;

        /// <summary>
        /// Cookie name for for sticky sessions
        /// </summary>
        [Output("stickySessionsCookieName")]
        public Output<string?> StickySessionsCookieName { get; private set; } = null!;

        /// <summary>
        /// Maximum initial server connection establishment time
        /// </summary>
        [Output("timeoutConnect")]
        public Output<string?> TimeoutConnect { get; private set; } = null!;

        /// <summary>
        /// Maximum server connection inactivity time
        /// </summary>
        [Output("timeoutServer")]
        public Output<string?> TimeoutServer { get; private set; } = null!;

        /// <summary>
        /// Maximum tunnel inactivity time
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
                PluginDownloadURL = "https://github.com/pulumiverse/pulumi-scaleway/releases/download/${VERSION}",
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

    public sealed class LoadbalancerBackendArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// User sessions will be forwarded to this port of backend servers
        /// </summary>
        [Input("forwardPort", required: true)]
        public Input<int> ForwardPort { get; set; } = null!;

        /// <summary>
        /// Load balancing algorithm
        /// </summary>
        [Input("forwardPortAlgorithm")]
        public Input<string>? ForwardPortAlgorithm { get; set; }

        /// <summary>
        /// Backend protocol
        /// </summary>
        [Input("forwardProtocol", required: true)]
        public Input<string> ForwardProtocol { get; set; } = null!;

        /// <summary>
        /// Interval between two HC requests
        /// </summary>
        [Input("healthCheckDelay")]
        public Input<string>? HealthCheckDelay { get; set; }

        [Input("healthCheckHttp")]
        public Input<Inputs.LoadbalancerBackendHealthCheckHttpArgs>? HealthCheckHttp { get; set; }

        [Input("healthCheckHttps")]
        public Input<Inputs.LoadbalancerBackendHealthCheckHttpsArgs>? HealthCheckHttps { get; set; }

        /// <summary>
        /// Number of allowed failed HC requests before the backend server is marked down
        /// </summary>
        [Input("healthCheckMaxRetries")]
        public Input<int>? HealthCheckMaxRetries { get; set; }

        /// <summary>
        /// Port the HC requests will be send to. Default to `forward_port`
        /// </summary>
        [Input("healthCheckPort")]
        public Input<int>? HealthCheckPort { get; set; }

        [Input("healthCheckTcp")]
        public Input<Inputs.LoadbalancerBackendHealthCheckTcpArgs>? HealthCheckTcp { get; set; }

        /// <summary>
        /// Timeout before we consider a HC request failed
        /// </summary>
        [Input("healthCheckTimeout")]
        public Input<string>? HealthCheckTimeout { get; set; }

        /// <summary>
        /// The load-balancer ID
        /// </summary>
        [Input("lbId", required: true)]
        public Input<string> LbId { get; set; } = null!;

        /// <summary>
        /// The name of the backend
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Modify what occurs when a backend server is marked down
        /// </summary>
        [Input("onMarkedDownAction")]
        public Input<string>? OnMarkedDownAction { get; set; }

        /// <summary>
        /// Type of PROXY protocol to enable
        /// </summary>
        [Input("proxyProtocol")]
        public Input<string>? ProxyProtocol { get; set; }

        /// <summary>
        /// Enables PROXY protocol version 2
        /// </summary>
        [Input("sendProxyV2")]
        public Input<bool>? SendProxyV2 { get; set; }

        [Input("serverIps")]
        private InputList<string>? _serverIps;

        /// <summary>
        /// Backend server IP addresses list (IPv4 or IPv6)
        /// </summary>
        public InputList<string> ServerIps
        {
            get => _serverIps ?? (_serverIps = new InputList<string>());
            set => _serverIps = value;
        }

        /// <summary>
        /// Load balancing algorithm
        /// </summary>
        [Input("stickySessions")]
        public Input<string>? StickySessions { get; set; }

        /// <summary>
        /// Cookie name for for sticky sessions
        /// </summary>
        [Input("stickySessionsCookieName")]
        public Input<string>? StickySessionsCookieName { get; set; }

        /// <summary>
        /// Maximum initial server connection establishment time
        /// </summary>
        [Input("timeoutConnect")]
        public Input<string>? TimeoutConnect { get; set; }

        /// <summary>
        /// Maximum server connection inactivity time
        /// </summary>
        [Input("timeoutServer")]
        public Input<string>? TimeoutServer { get; set; }

        /// <summary>
        /// Maximum tunnel inactivity time
        /// </summary>
        [Input("timeoutTunnel")]
        public Input<string>? TimeoutTunnel { get; set; }

        public LoadbalancerBackendArgs()
        {
        }
    }

    public sealed class LoadbalancerBackendState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// User sessions will be forwarded to this port of backend servers
        /// </summary>
        [Input("forwardPort")]
        public Input<int>? ForwardPort { get; set; }

        /// <summary>
        /// Load balancing algorithm
        /// </summary>
        [Input("forwardPortAlgorithm")]
        public Input<string>? ForwardPortAlgorithm { get; set; }

        /// <summary>
        /// Backend protocol
        /// </summary>
        [Input("forwardProtocol")]
        public Input<string>? ForwardProtocol { get; set; }

        /// <summary>
        /// Interval between two HC requests
        /// </summary>
        [Input("healthCheckDelay")]
        public Input<string>? HealthCheckDelay { get; set; }

        [Input("healthCheckHttp")]
        public Input<Inputs.LoadbalancerBackendHealthCheckHttpGetArgs>? HealthCheckHttp { get; set; }

        [Input("healthCheckHttps")]
        public Input<Inputs.LoadbalancerBackendHealthCheckHttpsGetArgs>? HealthCheckHttps { get; set; }

        /// <summary>
        /// Number of allowed failed HC requests before the backend server is marked down
        /// </summary>
        [Input("healthCheckMaxRetries")]
        public Input<int>? HealthCheckMaxRetries { get; set; }

        /// <summary>
        /// Port the HC requests will be send to. Default to `forward_port`
        /// </summary>
        [Input("healthCheckPort")]
        public Input<int>? HealthCheckPort { get; set; }

        [Input("healthCheckTcp")]
        public Input<Inputs.LoadbalancerBackendHealthCheckTcpGetArgs>? HealthCheckTcp { get; set; }

        /// <summary>
        /// Timeout before we consider a HC request failed
        /// </summary>
        [Input("healthCheckTimeout")]
        public Input<string>? HealthCheckTimeout { get; set; }

        /// <summary>
        /// The load-balancer ID
        /// </summary>
        [Input("lbId")]
        public Input<string>? LbId { get; set; }

        /// <summary>
        /// The name of the backend
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Modify what occurs when a backend server is marked down
        /// </summary>
        [Input("onMarkedDownAction")]
        public Input<string>? OnMarkedDownAction { get; set; }

        /// <summary>
        /// Type of PROXY protocol to enable
        /// </summary>
        [Input("proxyProtocol")]
        public Input<string>? ProxyProtocol { get; set; }

        /// <summary>
        /// Enables PROXY protocol version 2
        /// </summary>
        [Input("sendProxyV2")]
        public Input<bool>? SendProxyV2 { get; set; }

        [Input("serverIps")]
        private InputList<string>? _serverIps;

        /// <summary>
        /// Backend server IP addresses list (IPv4 or IPv6)
        /// </summary>
        public InputList<string> ServerIps
        {
            get => _serverIps ?? (_serverIps = new InputList<string>());
            set => _serverIps = value;
        }

        /// <summary>
        /// Load balancing algorithm
        /// </summary>
        [Input("stickySessions")]
        public Input<string>? StickySessions { get; set; }

        /// <summary>
        /// Cookie name for for sticky sessions
        /// </summary>
        [Input("stickySessionsCookieName")]
        public Input<string>? StickySessionsCookieName { get; set; }

        /// <summary>
        /// Maximum initial server connection establishment time
        /// </summary>
        [Input("timeoutConnect")]
        public Input<string>? TimeoutConnect { get; set; }

        /// <summary>
        /// Maximum server connection inactivity time
        /// </summary>
        [Input("timeoutServer")]
        public Input<string>? TimeoutServer { get; set; }

        /// <summary>
        /// Maximum tunnel inactivity time
        /// </summary>
        [Input("timeoutTunnel")]
        public Input<string>? TimeoutTunnel { get; set; }

        public LoadbalancerBackendState()
        {
        }
    }
}
