# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetVpcPublicGatewayDhcpResult',
    'AwaitableGetVpcPublicGatewayDhcpResult',
    'get_vpc_public_gateway_dhcp',
    'get_vpc_public_gateway_dhcp_output',
]

@pulumi.output_type
class GetVpcPublicGatewayDhcpResult:
    """
    A collection of values returned by getVpcPublicGatewayDhcp.
    """
    def __init__(__self__, address=None, created_at=None, dhcp_id=None, dns_local_name=None, dns_searches=None, dns_server_overrides=None, enable_dynamic=None, id=None, organization_id=None, pool_high=None, pool_low=None, project_id=None, push_default_route=None, push_dns_server=None, rebind_timer=None, renew_timer=None, subnet=None, updated_at=None, valid_lifetime=None, zone=None):
        if address and not isinstance(address, str):
            raise TypeError("Expected argument 'address' to be a str")
        pulumi.set(__self__, "address", address)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if dhcp_id and not isinstance(dhcp_id, str):
            raise TypeError("Expected argument 'dhcp_id' to be a str")
        pulumi.set(__self__, "dhcp_id", dhcp_id)
        if dns_local_name and not isinstance(dns_local_name, str):
            raise TypeError("Expected argument 'dns_local_name' to be a str")
        pulumi.set(__self__, "dns_local_name", dns_local_name)
        if dns_searches and not isinstance(dns_searches, list):
            raise TypeError("Expected argument 'dns_searches' to be a list")
        pulumi.set(__self__, "dns_searches", dns_searches)
        if dns_server_overrides and not isinstance(dns_server_overrides, list):
            raise TypeError("Expected argument 'dns_server_overrides' to be a list")
        pulumi.set(__self__, "dns_server_overrides", dns_server_overrides)
        if enable_dynamic and not isinstance(enable_dynamic, bool):
            raise TypeError("Expected argument 'enable_dynamic' to be a bool")
        pulumi.set(__self__, "enable_dynamic", enable_dynamic)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if organization_id and not isinstance(organization_id, str):
            raise TypeError("Expected argument 'organization_id' to be a str")
        pulumi.set(__self__, "organization_id", organization_id)
        if pool_high and not isinstance(pool_high, str):
            raise TypeError("Expected argument 'pool_high' to be a str")
        pulumi.set(__self__, "pool_high", pool_high)
        if pool_low and not isinstance(pool_low, str):
            raise TypeError("Expected argument 'pool_low' to be a str")
        pulumi.set(__self__, "pool_low", pool_low)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if push_default_route and not isinstance(push_default_route, bool):
            raise TypeError("Expected argument 'push_default_route' to be a bool")
        pulumi.set(__self__, "push_default_route", push_default_route)
        if push_dns_server and not isinstance(push_dns_server, bool):
            raise TypeError("Expected argument 'push_dns_server' to be a bool")
        pulumi.set(__self__, "push_dns_server", push_dns_server)
        if rebind_timer and not isinstance(rebind_timer, int):
            raise TypeError("Expected argument 'rebind_timer' to be a int")
        pulumi.set(__self__, "rebind_timer", rebind_timer)
        if renew_timer and not isinstance(renew_timer, int):
            raise TypeError("Expected argument 'renew_timer' to be a int")
        pulumi.set(__self__, "renew_timer", renew_timer)
        if subnet and not isinstance(subnet, str):
            raise TypeError("Expected argument 'subnet' to be a str")
        pulumi.set(__self__, "subnet", subnet)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)
        if valid_lifetime and not isinstance(valid_lifetime, int):
            raise TypeError("Expected argument 'valid_lifetime' to be a int")
        pulumi.set(__self__, "valid_lifetime", valid_lifetime)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def address(self) -> str:
        return pulumi.get(self, "address")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="dhcpId")
    def dhcp_id(self) -> str:
        return pulumi.get(self, "dhcp_id")

    @property
    @pulumi.getter(name="dnsLocalName")
    def dns_local_name(self) -> str:
        return pulumi.get(self, "dns_local_name")

    @property
    @pulumi.getter(name="dnsSearches")
    def dns_searches(self) -> Sequence[str]:
        return pulumi.get(self, "dns_searches")

    @property
    @pulumi.getter(name="dnsServerOverrides")
    def dns_server_overrides(self) -> Sequence[str]:
        return pulumi.get(self, "dns_server_overrides")

    @property
    @pulumi.getter(name="enableDynamic")
    def enable_dynamic(self) -> bool:
        return pulumi.get(self, "enable_dynamic")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> str:
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter(name="poolHigh")
    def pool_high(self) -> str:
        return pulumi.get(self, "pool_high")

    @property
    @pulumi.getter(name="poolLow")
    def pool_low(self) -> str:
        return pulumi.get(self, "pool_low")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="pushDefaultRoute")
    def push_default_route(self) -> bool:
        return pulumi.get(self, "push_default_route")

    @property
    @pulumi.getter(name="pushDnsServer")
    def push_dns_server(self) -> bool:
        return pulumi.get(self, "push_dns_server")

    @property
    @pulumi.getter(name="rebindTimer")
    def rebind_timer(self) -> int:
        return pulumi.get(self, "rebind_timer")

    @property
    @pulumi.getter(name="renewTimer")
    def renew_timer(self) -> int:
        return pulumi.get(self, "renew_timer")

    @property
    @pulumi.getter
    def subnet(self) -> str:
        return pulumi.get(self, "subnet")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter(name="validLifetime")
    def valid_lifetime(self) -> int:
        return pulumi.get(self, "valid_lifetime")

    @property
    @pulumi.getter
    def zone(self) -> str:
        return pulumi.get(self, "zone")


class AwaitableGetVpcPublicGatewayDhcpResult(GetVpcPublicGatewayDhcpResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVpcPublicGatewayDhcpResult(
            address=self.address,
            created_at=self.created_at,
            dhcp_id=self.dhcp_id,
            dns_local_name=self.dns_local_name,
            dns_searches=self.dns_searches,
            dns_server_overrides=self.dns_server_overrides,
            enable_dynamic=self.enable_dynamic,
            id=self.id,
            organization_id=self.organization_id,
            pool_high=self.pool_high,
            pool_low=self.pool_low,
            project_id=self.project_id,
            push_default_route=self.push_default_route,
            push_dns_server=self.push_dns_server,
            rebind_timer=self.rebind_timer,
            renew_timer=self.renew_timer,
            subnet=self.subnet,
            updated_at=self.updated_at,
            valid_lifetime=self.valid_lifetime,
            zone=self.zone)


def get_vpc_public_gateway_dhcp(dhcp_id: Optional[str] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVpcPublicGatewayDhcpResult:
    """
    Gets information about a public gateway DHCP.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    main = scaleway.VpcPublicGatewayDhcp("main", subnet="192.168.0.0/24")
    dhcp_by_id = scaleway.get_vpc_public_gateway_dhcp_output(dhcp_id=main.id)
    ```
    """
    __args__ = dict()
    __args__['dhcpId'] = dhcp_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('scaleway:index/getVpcPublicGatewayDhcp:getVpcPublicGatewayDhcp', __args__, opts=opts, typ=GetVpcPublicGatewayDhcpResult).value

    return AwaitableGetVpcPublicGatewayDhcpResult(
        address=__ret__.address,
        created_at=__ret__.created_at,
        dhcp_id=__ret__.dhcp_id,
        dns_local_name=__ret__.dns_local_name,
        dns_searches=__ret__.dns_searches,
        dns_server_overrides=__ret__.dns_server_overrides,
        enable_dynamic=__ret__.enable_dynamic,
        id=__ret__.id,
        organization_id=__ret__.organization_id,
        pool_high=__ret__.pool_high,
        pool_low=__ret__.pool_low,
        project_id=__ret__.project_id,
        push_default_route=__ret__.push_default_route,
        push_dns_server=__ret__.push_dns_server,
        rebind_timer=__ret__.rebind_timer,
        renew_timer=__ret__.renew_timer,
        subnet=__ret__.subnet,
        updated_at=__ret__.updated_at,
        valid_lifetime=__ret__.valid_lifetime,
        zone=__ret__.zone)


@_utilities.lift_output_func(get_vpc_public_gateway_dhcp)
def get_vpc_public_gateway_dhcp_output(dhcp_id: Optional[pulumi.Input[str]] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVpcPublicGatewayDhcpResult]:
    """
    Gets information about a public gateway DHCP.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    main = scaleway.VpcPublicGatewayDhcp("main", subnet="192.168.0.0/24")
    dhcp_by_id = scaleway.get_vpc_public_gateway_dhcp_output(dhcp_id=main.id)
    ```
    """
    ...
