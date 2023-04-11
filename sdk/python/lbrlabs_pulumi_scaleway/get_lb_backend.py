# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetLbBackendResult',
    'AwaitableGetLbBackendResult',
    'get_lb_backend',
    'get_lb_backend_output',
]

@pulumi.output_type
class GetLbBackendResult:
    """
    A collection of values returned by getLbBackend.
    """
    def __init__(__self__, backend_id=None, failover_host=None, forward_port=None, forward_port_algorithm=None, forward_protocol=None, health_check_delay=None, health_check_http=None, health_check_https=None, health_check_max_retries=None, health_check_port=None, health_check_tcps=None, health_check_timeout=None, id=None, ignore_ssl_server_verify=None, lb_id=None, name=None, on_marked_down_action=None, proxy_protocol=None, send_proxy_v2=None, server_ips=None, ssl_bridging=None, sticky_sessions=None, sticky_sessions_cookie_name=None, timeout_connect=None, timeout_server=None, timeout_tunnel=None):
        if backend_id and not isinstance(backend_id, str):
            raise TypeError("Expected argument 'backend_id' to be a str")
        pulumi.set(__self__, "backend_id", backend_id)
        if failover_host and not isinstance(failover_host, str):
            raise TypeError("Expected argument 'failover_host' to be a str")
        pulumi.set(__self__, "failover_host", failover_host)
        if forward_port and not isinstance(forward_port, int):
            raise TypeError("Expected argument 'forward_port' to be a int")
        pulumi.set(__self__, "forward_port", forward_port)
        if forward_port_algorithm and not isinstance(forward_port_algorithm, str):
            raise TypeError("Expected argument 'forward_port_algorithm' to be a str")
        pulumi.set(__self__, "forward_port_algorithm", forward_port_algorithm)
        if forward_protocol and not isinstance(forward_protocol, str):
            raise TypeError("Expected argument 'forward_protocol' to be a str")
        pulumi.set(__self__, "forward_protocol", forward_protocol)
        if health_check_delay and not isinstance(health_check_delay, str):
            raise TypeError("Expected argument 'health_check_delay' to be a str")
        pulumi.set(__self__, "health_check_delay", health_check_delay)
        if health_check_http and not isinstance(health_check_http, list):
            raise TypeError("Expected argument 'health_check_http' to be a list")
        pulumi.set(__self__, "health_check_http", health_check_http)
        if health_check_https and not isinstance(health_check_https, list):
            raise TypeError("Expected argument 'health_check_https' to be a list")
        pulumi.set(__self__, "health_check_https", health_check_https)
        if health_check_max_retries and not isinstance(health_check_max_retries, int):
            raise TypeError("Expected argument 'health_check_max_retries' to be a int")
        pulumi.set(__self__, "health_check_max_retries", health_check_max_retries)
        if health_check_port and not isinstance(health_check_port, int):
            raise TypeError("Expected argument 'health_check_port' to be a int")
        pulumi.set(__self__, "health_check_port", health_check_port)
        if health_check_tcps and not isinstance(health_check_tcps, list):
            raise TypeError("Expected argument 'health_check_tcps' to be a list")
        pulumi.set(__self__, "health_check_tcps", health_check_tcps)
        if health_check_timeout and not isinstance(health_check_timeout, str):
            raise TypeError("Expected argument 'health_check_timeout' to be a str")
        pulumi.set(__self__, "health_check_timeout", health_check_timeout)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ignore_ssl_server_verify and not isinstance(ignore_ssl_server_verify, bool):
            raise TypeError("Expected argument 'ignore_ssl_server_verify' to be a bool")
        pulumi.set(__self__, "ignore_ssl_server_verify", ignore_ssl_server_verify)
        if lb_id and not isinstance(lb_id, str):
            raise TypeError("Expected argument 'lb_id' to be a str")
        pulumi.set(__self__, "lb_id", lb_id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if on_marked_down_action and not isinstance(on_marked_down_action, str):
            raise TypeError("Expected argument 'on_marked_down_action' to be a str")
        pulumi.set(__self__, "on_marked_down_action", on_marked_down_action)
        if proxy_protocol and not isinstance(proxy_protocol, str):
            raise TypeError("Expected argument 'proxy_protocol' to be a str")
        pulumi.set(__self__, "proxy_protocol", proxy_protocol)
        if send_proxy_v2 and not isinstance(send_proxy_v2, bool):
            raise TypeError("Expected argument 'send_proxy_v2' to be a bool")
        pulumi.set(__self__, "send_proxy_v2", send_proxy_v2)
        if server_ips and not isinstance(server_ips, list):
            raise TypeError("Expected argument 'server_ips' to be a list")
        pulumi.set(__self__, "server_ips", server_ips)
        if ssl_bridging and not isinstance(ssl_bridging, bool):
            raise TypeError("Expected argument 'ssl_bridging' to be a bool")
        pulumi.set(__self__, "ssl_bridging", ssl_bridging)
        if sticky_sessions and not isinstance(sticky_sessions, str):
            raise TypeError("Expected argument 'sticky_sessions' to be a str")
        pulumi.set(__self__, "sticky_sessions", sticky_sessions)
        if sticky_sessions_cookie_name and not isinstance(sticky_sessions_cookie_name, str):
            raise TypeError("Expected argument 'sticky_sessions_cookie_name' to be a str")
        pulumi.set(__self__, "sticky_sessions_cookie_name", sticky_sessions_cookie_name)
        if timeout_connect and not isinstance(timeout_connect, str):
            raise TypeError("Expected argument 'timeout_connect' to be a str")
        pulumi.set(__self__, "timeout_connect", timeout_connect)
        if timeout_server and not isinstance(timeout_server, str):
            raise TypeError("Expected argument 'timeout_server' to be a str")
        pulumi.set(__self__, "timeout_server", timeout_server)
        if timeout_tunnel and not isinstance(timeout_tunnel, str):
            raise TypeError("Expected argument 'timeout_tunnel' to be a str")
        pulumi.set(__self__, "timeout_tunnel", timeout_tunnel)

    @property
    @pulumi.getter(name="backendId")
    def backend_id(self) -> Optional[str]:
        return pulumi.get(self, "backend_id")

    @property
    @pulumi.getter(name="failoverHost")
    def failover_host(self) -> str:
        return pulumi.get(self, "failover_host")

    @property
    @pulumi.getter(name="forwardPort")
    def forward_port(self) -> int:
        return pulumi.get(self, "forward_port")

    @property
    @pulumi.getter(name="forwardPortAlgorithm")
    def forward_port_algorithm(self) -> str:
        return pulumi.get(self, "forward_port_algorithm")

    @property
    @pulumi.getter(name="forwardProtocol")
    def forward_protocol(self) -> str:
        return pulumi.get(self, "forward_protocol")

    @property
    @pulumi.getter(name="healthCheckDelay")
    def health_check_delay(self) -> str:
        return pulumi.get(self, "health_check_delay")

    @property
    @pulumi.getter(name="healthCheckHttp")
    def health_check_http(self) -> Sequence['outputs.GetLbBackendHealthCheckHttpResult']:
        return pulumi.get(self, "health_check_http")

    @property
    @pulumi.getter(name="healthCheckHttps")
    def health_check_https(self) -> Sequence['outputs.GetLbBackendHealthCheckHttpResult']:
        return pulumi.get(self, "health_check_https")

    @property
    @pulumi.getter(name="healthCheckMaxRetries")
    def health_check_max_retries(self) -> int:
        return pulumi.get(self, "health_check_max_retries")

    @property
    @pulumi.getter(name="healthCheckPort")
    def health_check_port(self) -> int:
        return pulumi.get(self, "health_check_port")

    @property
    @pulumi.getter(name="healthCheckTcps")
    def health_check_tcps(self) -> Sequence['outputs.GetLbBackendHealthCheckTcpResult']:
        return pulumi.get(self, "health_check_tcps")

    @property
    @pulumi.getter(name="healthCheckTimeout")
    def health_check_timeout(self) -> str:
        return pulumi.get(self, "health_check_timeout")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ignoreSslServerVerify")
    def ignore_ssl_server_verify(self) -> bool:
        return pulumi.get(self, "ignore_ssl_server_verify")

    @property
    @pulumi.getter(name="lbId")
    def lb_id(self) -> Optional[str]:
        return pulumi.get(self, "lb_id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="onMarkedDownAction")
    def on_marked_down_action(self) -> str:
        return pulumi.get(self, "on_marked_down_action")

    @property
    @pulumi.getter(name="proxyProtocol")
    def proxy_protocol(self) -> str:
        return pulumi.get(self, "proxy_protocol")

    @property
    @pulumi.getter(name="sendProxyV2")
    def send_proxy_v2(self) -> bool:
        return pulumi.get(self, "send_proxy_v2")

    @property
    @pulumi.getter(name="serverIps")
    def server_ips(self) -> Sequence[str]:
        return pulumi.get(self, "server_ips")

    @property
    @pulumi.getter(name="sslBridging")
    def ssl_bridging(self) -> bool:
        return pulumi.get(self, "ssl_bridging")

    @property
    @pulumi.getter(name="stickySessions")
    def sticky_sessions(self) -> str:
        return pulumi.get(self, "sticky_sessions")

    @property
    @pulumi.getter(name="stickySessionsCookieName")
    def sticky_sessions_cookie_name(self) -> str:
        return pulumi.get(self, "sticky_sessions_cookie_name")

    @property
    @pulumi.getter(name="timeoutConnect")
    def timeout_connect(self) -> str:
        return pulumi.get(self, "timeout_connect")

    @property
    @pulumi.getter(name="timeoutServer")
    def timeout_server(self) -> str:
        return pulumi.get(self, "timeout_server")

    @property
    @pulumi.getter(name="timeoutTunnel")
    def timeout_tunnel(self) -> str:
        return pulumi.get(self, "timeout_tunnel")


class AwaitableGetLbBackendResult(GetLbBackendResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLbBackendResult(
            backend_id=self.backend_id,
            failover_host=self.failover_host,
            forward_port=self.forward_port,
            forward_port_algorithm=self.forward_port_algorithm,
            forward_protocol=self.forward_protocol,
            health_check_delay=self.health_check_delay,
            health_check_http=self.health_check_http,
            health_check_https=self.health_check_https,
            health_check_max_retries=self.health_check_max_retries,
            health_check_port=self.health_check_port,
            health_check_tcps=self.health_check_tcps,
            health_check_timeout=self.health_check_timeout,
            id=self.id,
            ignore_ssl_server_verify=self.ignore_ssl_server_verify,
            lb_id=self.lb_id,
            name=self.name,
            on_marked_down_action=self.on_marked_down_action,
            proxy_protocol=self.proxy_protocol,
            send_proxy_v2=self.send_proxy_v2,
            server_ips=self.server_ips,
            ssl_bridging=self.ssl_bridging,
            sticky_sessions=self.sticky_sessions,
            sticky_sessions_cookie_name=self.sticky_sessions_cookie_name,
            timeout_connect=self.timeout_connect,
            timeout_server=self.timeout_server,
            timeout_tunnel=self.timeout_tunnel)


def get_lb_backend(backend_id: Optional[str] = None,
                   lb_id: Optional[str] = None,
                   name: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLbBackendResult:
    """
    Get information about Scaleway Load-Balancer Backends.
    For more information, see [the documentation](https://developers.scaleway.com/en/products/lb/zoned_api/#backends-cbf4eb).

    ## Example Usage

    ```python
    import pulumi
    import lbrlabs_pulumi_scaleway as scaleway
    import pulumi_scaleway as scaleway

    main_loadbalancer_ip = scaleway.LoadbalancerIp("mainLoadbalancerIp")
    main_loadbalancer = scaleway.Loadbalancer("mainLoadbalancer",
        ip_id=main_loadbalancer_ip.id,
        type="LB-S")
    main_loadbalancer_backend = scaleway.LoadbalancerBackend("mainLoadbalancerBackend",
        lb_id=main_loadbalancer.id,
        forward_protocol="http",
        forward_port=80)
    by_id = scaleway.get_lb_backend_output(backend_id=main_loadbalancer_backend.id)
    by_name = scaleway.get_lb_backend_output(name=main_loadbalancer_backend.name,
        lb_id=main_loadbalancer.id)
    ```


    :param str backend_id: The backend id.
           - Only one of `name` and `backend_id` should be specified.
    :param str lb_id: The load-balancer ID this backend is attached to.
    :param str name: The name of the backend.
           - When using the `name` you should specify the `lb-id`
    """
    __args__ = dict()
    __args__['backendId'] = backend_id
    __args__['lbId'] = lb_id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:index/getLbBackend:getLbBackend', __args__, opts=opts, typ=GetLbBackendResult).value

    return AwaitableGetLbBackendResult(
        backend_id=__ret__.backend_id,
        failover_host=__ret__.failover_host,
        forward_port=__ret__.forward_port,
        forward_port_algorithm=__ret__.forward_port_algorithm,
        forward_protocol=__ret__.forward_protocol,
        health_check_delay=__ret__.health_check_delay,
        health_check_http=__ret__.health_check_http,
        health_check_https=__ret__.health_check_https,
        health_check_max_retries=__ret__.health_check_max_retries,
        health_check_port=__ret__.health_check_port,
        health_check_tcps=__ret__.health_check_tcps,
        health_check_timeout=__ret__.health_check_timeout,
        id=__ret__.id,
        ignore_ssl_server_verify=__ret__.ignore_ssl_server_verify,
        lb_id=__ret__.lb_id,
        name=__ret__.name,
        on_marked_down_action=__ret__.on_marked_down_action,
        proxy_protocol=__ret__.proxy_protocol,
        send_proxy_v2=__ret__.send_proxy_v2,
        server_ips=__ret__.server_ips,
        ssl_bridging=__ret__.ssl_bridging,
        sticky_sessions=__ret__.sticky_sessions,
        sticky_sessions_cookie_name=__ret__.sticky_sessions_cookie_name,
        timeout_connect=__ret__.timeout_connect,
        timeout_server=__ret__.timeout_server,
        timeout_tunnel=__ret__.timeout_tunnel)


@_utilities.lift_output_func(get_lb_backend)
def get_lb_backend_output(backend_id: Optional[pulumi.Input[Optional[str]]] = None,
                          lb_id: Optional[pulumi.Input[Optional[str]]] = None,
                          name: Optional[pulumi.Input[Optional[str]]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetLbBackendResult]:
    """
    Get information about Scaleway Load-Balancer Backends.
    For more information, see [the documentation](https://developers.scaleway.com/en/products/lb/zoned_api/#backends-cbf4eb).

    ## Example Usage

    ```python
    import pulumi
    import lbrlabs_pulumi_scaleway as scaleway
    import pulumi_scaleway as scaleway

    main_loadbalancer_ip = scaleway.LoadbalancerIp("mainLoadbalancerIp")
    main_loadbalancer = scaleway.Loadbalancer("mainLoadbalancer",
        ip_id=main_loadbalancer_ip.id,
        type="LB-S")
    main_loadbalancer_backend = scaleway.LoadbalancerBackend("mainLoadbalancerBackend",
        lb_id=main_loadbalancer.id,
        forward_protocol="http",
        forward_port=80)
    by_id = scaleway.get_lb_backend_output(backend_id=main_loadbalancer_backend.id)
    by_name = scaleway.get_lb_backend_output(name=main_loadbalancer_backend.name,
        lb_id=main_loadbalancer.id)
    ```


    :param str backend_id: The backend id.
           - Only one of `name` and `backend_id` should be specified.
    :param str lb_id: The load-balancer ID this backend is attached to.
    :param str name: The name of the backend.
           - When using the `name` you should specify the `lb-id`
    """
    ...
