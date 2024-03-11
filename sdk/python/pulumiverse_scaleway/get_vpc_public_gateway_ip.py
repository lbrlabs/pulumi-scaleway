# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetVpcPublicGatewayIpResult',
    'AwaitableGetVpcPublicGatewayIpResult',
    'get_vpc_public_gateway_ip',
    'get_vpc_public_gateway_ip_output',
]

@pulumi.output_type
class GetVpcPublicGatewayIpResult:
    """
    A collection of values returned by getVpcPublicGatewayIp.
    """
    def __init__(__self__, address=None, created_at=None, id=None, ip_id=None, organization_id=None, project_id=None, reverse=None, tags=None, updated_at=None, zone=None):
        if address and not isinstance(address, str):
            raise TypeError("Expected argument 'address' to be a str")
        pulumi.set(__self__, "address", address)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip_id and not isinstance(ip_id, str):
            raise TypeError("Expected argument 'ip_id' to be a str")
        pulumi.set(__self__, "ip_id", ip_id)
        if organization_id and not isinstance(organization_id, str):
            raise TypeError("Expected argument 'organization_id' to be a str")
        pulumi.set(__self__, "organization_id", organization_id)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if reverse and not isinstance(reverse, str):
            raise TypeError("Expected argument 'reverse' to be a str")
        pulumi.set(__self__, "reverse", reverse)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)
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
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ipId")
    def ip_id(self) -> Optional[str]:
        return pulumi.get(self, "ip_id")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> str:
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def reverse(self) -> str:
        return pulumi.get(self, "reverse")

    @property
    @pulumi.getter
    def tags(self) -> Sequence[str]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter
    def zone(self) -> str:
        return pulumi.get(self, "zone")


class AwaitableGetVpcPublicGatewayIpResult(GetVpcPublicGatewayIpResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVpcPublicGatewayIpResult(
            address=self.address,
            created_at=self.created_at,
            id=self.id,
            ip_id=self.ip_id,
            organization_id=self.organization_id,
            project_id=self.project_id,
            reverse=self.reverse,
            tags=self.tags,
            updated_at=self.updated_at,
            zone=self.zone)


def get_vpc_public_gateway_ip(ip_id: Optional[str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVpcPublicGatewayIpResult:
    """
    Gets information about a public gateway IP.

    For further information please check the API [documentation](https://developers.scaleway.com/en/products/vpc-gw/api/v1/#get-66f0c0)

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_scaleway as scaleway
    import pulumiverse_scaleway as scaleway

    main = scaleway.VpcPublicGatewayIp("main")
    ip_by_id = scaleway.get_vpc_public_gateway_ip_output(ip_id=main.id)
    ```
    <!--End PulumiCodeChooser -->
    """
    __args__ = dict()
    __args__['ipId'] = ip_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:index/getVpcPublicGatewayIp:getVpcPublicGatewayIp', __args__, opts=opts, typ=GetVpcPublicGatewayIpResult).value

    return AwaitableGetVpcPublicGatewayIpResult(
        address=pulumi.get(__ret__, 'address'),
        created_at=pulumi.get(__ret__, 'created_at'),
        id=pulumi.get(__ret__, 'id'),
        ip_id=pulumi.get(__ret__, 'ip_id'),
        organization_id=pulumi.get(__ret__, 'organization_id'),
        project_id=pulumi.get(__ret__, 'project_id'),
        reverse=pulumi.get(__ret__, 'reverse'),
        tags=pulumi.get(__ret__, 'tags'),
        updated_at=pulumi.get(__ret__, 'updated_at'),
        zone=pulumi.get(__ret__, 'zone'))


@_utilities.lift_output_func(get_vpc_public_gateway_ip)
def get_vpc_public_gateway_ip_output(ip_id: Optional[pulumi.Input[Optional[str]]] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVpcPublicGatewayIpResult]:
    """
    Gets information about a public gateway IP.

    For further information please check the API [documentation](https://developers.scaleway.com/en/products/vpc-gw/api/v1/#get-66f0c0)

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_scaleway as scaleway
    import pulumiverse_scaleway as scaleway

    main = scaleway.VpcPublicGatewayIp("main")
    ip_by_id = scaleway.get_vpc_public_gateway_ip_output(ip_id=main.id)
    ```
    <!--End PulumiCodeChooser -->
    """
    ...
