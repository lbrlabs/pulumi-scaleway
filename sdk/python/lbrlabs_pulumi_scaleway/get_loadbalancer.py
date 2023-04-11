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
    'GetLoadbalancerResult',
    'AwaitableGetLoadbalancerResult',
    'get_loadbalancer',
    'get_loadbalancer_output',
]

@pulumi.output_type
class GetLoadbalancerResult:
    """
    A collection of values returned by getLoadbalancer.
    """
    def __init__(__self__, description=None, id=None, ip_address=None, ip_id=None, lb_id=None, name=None, organization_id=None, private_networks=None, project_id=None, region=None, release_ip=None, ssl_compatibility_level=None, tags=None, type=None, zone=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip_address and not isinstance(ip_address, str):
            raise TypeError("Expected argument 'ip_address' to be a str")
        pulumi.set(__self__, "ip_address", ip_address)
        if ip_id and not isinstance(ip_id, str):
            raise TypeError("Expected argument 'ip_id' to be a str")
        pulumi.set(__self__, "ip_id", ip_id)
        if lb_id and not isinstance(lb_id, str):
            raise TypeError("Expected argument 'lb_id' to be a str")
        pulumi.set(__self__, "lb_id", lb_id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if organization_id and not isinstance(organization_id, str):
            raise TypeError("Expected argument 'organization_id' to be a str")
        pulumi.set(__self__, "organization_id", organization_id)
        if private_networks and not isinstance(private_networks, list):
            raise TypeError("Expected argument 'private_networks' to be a list")
        pulumi.set(__self__, "private_networks", private_networks)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if release_ip and not isinstance(release_ip, bool):
            raise TypeError("Expected argument 'release_ip' to be a bool")
        pulumi.set(__self__, "release_ip", release_ip)
        if ssl_compatibility_level and not isinstance(ssl_compatibility_level, str):
            raise TypeError("Expected argument 'ssl_compatibility_level' to be a str")
        pulumi.set(__self__, "ssl_compatibility_level", ssl_compatibility_level)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> str:
        """
        The load-balancer public IP Address.
        """
        return pulumi.get(self, "ip_address")

    @property
    @pulumi.getter(name="ipId")
    def ip_id(self) -> str:
        return pulumi.get(self, "ip_id")

    @property
    @pulumi.getter(name="lbId")
    def lb_id(self) -> Optional[str]:
        return pulumi.get(self, "lb_id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> str:
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter(name="privateNetworks")
    def private_networks(self) -> Sequence['outputs.GetLoadbalancerPrivateNetworkResult']:
        return pulumi.get(self, "private_networks")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        """
        (Defaults to provider `project_id`) The ID of the project the LB is associated with.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="releaseIp")
    def release_ip(self) -> Optional[bool]:
        return pulumi.get(self, "release_ip")

    @property
    @pulumi.getter(name="sslCompatibilityLevel")
    def ssl_compatibility_level(self) -> str:
        return pulumi.get(self, "ssl_compatibility_level")

    @property
    @pulumi.getter
    def tags(self) -> Sequence[str]:
        """
        The tags associated with the load-balancer.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the load-balancer.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def zone(self) -> Optional[str]:
        """
        (Defaults to provider `zone`) The zone in which the LB exists.
        """
        return pulumi.get(self, "zone")


class AwaitableGetLoadbalancerResult(GetLoadbalancerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLoadbalancerResult(
            description=self.description,
            id=self.id,
            ip_address=self.ip_address,
            ip_id=self.ip_id,
            lb_id=self.lb_id,
            name=self.name,
            organization_id=self.organization_id,
            private_networks=self.private_networks,
            project_id=self.project_id,
            region=self.region,
            release_ip=self.release_ip,
            ssl_compatibility_level=self.ssl_compatibility_level,
            tags=self.tags,
            type=self.type,
            zone=self.zone)


def get_loadbalancer(lb_id: Optional[str] = None,
                     name: Optional[str] = None,
                     release_ip: Optional[bool] = None,
                     zone: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLoadbalancerResult:
    """
    Gets information about a Load Balancer.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    by_name = scaleway.get_loadbalancer(name="foobar")
    by_id = scaleway.get_loadbalancer(lb_id="11111111-1111-1111-1111-111111111111")
    ```


    :param str name: The load balancer name.
    :param str zone: (Defaults to provider `zone`) The zone in which the LB exists.
    """
    __args__ = dict()
    __args__['lbId'] = lb_id
    __args__['name'] = name
    __args__['releaseIp'] = release_ip
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:index/getLoadbalancer:getLoadbalancer', __args__, opts=opts, typ=GetLoadbalancerResult).value

    return AwaitableGetLoadbalancerResult(
        description=__ret__.description,
        id=__ret__.id,
        ip_address=__ret__.ip_address,
        ip_id=__ret__.ip_id,
        lb_id=__ret__.lb_id,
        name=__ret__.name,
        organization_id=__ret__.organization_id,
        private_networks=__ret__.private_networks,
        project_id=__ret__.project_id,
        region=__ret__.region,
        release_ip=__ret__.release_ip,
        ssl_compatibility_level=__ret__.ssl_compatibility_level,
        tags=__ret__.tags,
        type=__ret__.type,
        zone=__ret__.zone)


@_utilities.lift_output_func(get_loadbalancer)
def get_loadbalancer_output(lb_id: Optional[pulumi.Input[Optional[str]]] = None,
                            name: Optional[pulumi.Input[Optional[str]]] = None,
                            release_ip: Optional[pulumi.Input[Optional[bool]]] = None,
                            zone: Optional[pulumi.Input[Optional[str]]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetLoadbalancerResult]:
    """
    Gets information about a Load Balancer.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    by_name = scaleway.get_loadbalancer(name="foobar")
    by_id = scaleway.get_loadbalancer(lb_id="11111111-1111-1111-1111-111111111111")
    ```


    :param str name: The load balancer name.
    :param str zone: (Defaults to provider `zone`) The zone in which the LB exists.
    """
    ...
