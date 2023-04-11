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
    'GetBaremetalServerResult',
    'AwaitableGetBaremetalServerResult',
    'get_baremetal_server',
    'get_baremetal_server_output',
]

@pulumi.output_type
class GetBaremetalServerResult:
    """
    A collection of values returned by getBaremetalServer.
    """
    def __init__(__self__, description=None, domain=None, hostname=None, id=None, ips=None, name=None, offer=None, offer_id=None, offer_name=None, options=None, organization_id=None, os=None, os_name=None, password=None, private_networks=None, project_id=None, reinstall_on_config_changes=None, server_id=None, service_password=None, service_user=None, ssh_key_ids=None, tags=None, user=None, zone=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if domain and not isinstance(domain, str):
            raise TypeError("Expected argument 'domain' to be a str")
        pulumi.set(__self__, "domain", domain)
        if hostname and not isinstance(hostname, str):
            raise TypeError("Expected argument 'hostname' to be a str")
        pulumi.set(__self__, "hostname", hostname)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ips and not isinstance(ips, list):
            raise TypeError("Expected argument 'ips' to be a list")
        pulumi.set(__self__, "ips", ips)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if offer and not isinstance(offer, str):
            raise TypeError("Expected argument 'offer' to be a str")
        pulumi.set(__self__, "offer", offer)
        if offer_id and not isinstance(offer_id, str):
            raise TypeError("Expected argument 'offer_id' to be a str")
        pulumi.set(__self__, "offer_id", offer_id)
        if offer_name and not isinstance(offer_name, str):
            raise TypeError("Expected argument 'offer_name' to be a str")
        pulumi.set(__self__, "offer_name", offer_name)
        if options and not isinstance(options, list):
            raise TypeError("Expected argument 'options' to be a list")
        pulumi.set(__self__, "options", options)
        if organization_id and not isinstance(organization_id, str):
            raise TypeError("Expected argument 'organization_id' to be a str")
        pulumi.set(__self__, "organization_id", organization_id)
        if os and not isinstance(os, str):
            raise TypeError("Expected argument 'os' to be a str")
        pulumi.set(__self__, "os", os)
        if os_name and not isinstance(os_name, str):
            raise TypeError("Expected argument 'os_name' to be a str")
        pulumi.set(__self__, "os_name", os_name)
        if password and not isinstance(password, str):
            raise TypeError("Expected argument 'password' to be a str")
        pulumi.set(__self__, "password", password)
        if private_networks and not isinstance(private_networks, list):
            raise TypeError("Expected argument 'private_networks' to be a list")
        pulumi.set(__self__, "private_networks", private_networks)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if reinstall_on_config_changes and not isinstance(reinstall_on_config_changes, bool):
            raise TypeError("Expected argument 'reinstall_on_config_changes' to be a bool")
        pulumi.set(__self__, "reinstall_on_config_changes", reinstall_on_config_changes)
        if server_id and not isinstance(server_id, str):
            raise TypeError("Expected argument 'server_id' to be a str")
        pulumi.set(__self__, "server_id", server_id)
        if service_password and not isinstance(service_password, str):
            raise TypeError("Expected argument 'service_password' to be a str")
        pulumi.set(__self__, "service_password", service_password)
        if service_user and not isinstance(service_user, str):
            raise TypeError("Expected argument 'service_user' to be a str")
        pulumi.set(__self__, "service_user", service_user)
        if ssh_key_ids and not isinstance(ssh_key_ids, list):
            raise TypeError("Expected argument 'ssh_key_ids' to be a list")
        pulumi.set(__self__, "ssh_key_ids", ssh_key_ids)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if user and not isinstance(user, str):
            raise TypeError("Expected argument 'user' to be a str")
        pulumi.set(__self__, "user", user)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def domain(self) -> str:
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter
    def hostname(self) -> str:
        return pulumi.get(self, "hostname")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ips(self) -> Sequence['outputs.GetBaremetalServerIpResult']:
        return pulumi.get(self, "ips")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def offer(self) -> str:
        return pulumi.get(self, "offer")

    @property
    @pulumi.getter(name="offerId")
    def offer_id(self) -> str:
        return pulumi.get(self, "offer_id")

    @property
    @pulumi.getter(name="offerName")
    def offer_name(self) -> str:
        return pulumi.get(self, "offer_name")

    @property
    @pulumi.getter
    def options(self) -> Sequence['outputs.GetBaremetalServerOptionResult']:
        return pulumi.get(self, "options")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> str:
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter
    def os(self) -> str:
        return pulumi.get(self, "os")

    @property
    @pulumi.getter(name="osName")
    def os_name(self) -> str:
        return pulumi.get(self, "os_name")

    @property
    @pulumi.getter
    def password(self) -> str:
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="privateNetworks")
    def private_networks(self) -> Sequence['outputs.GetBaremetalServerPrivateNetworkResult']:
        return pulumi.get(self, "private_networks")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="reinstallOnConfigChanges")
    def reinstall_on_config_changes(self) -> bool:
        return pulumi.get(self, "reinstall_on_config_changes")

    @property
    @pulumi.getter(name="serverId")
    def server_id(self) -> Optional[str]:
        return pulumi.get(self, "server_id")

    @property
    @pulumi.getter(name="servicePassword")
    def service_password(self) -> str:
        return pulumi.get(self, "service_password")

    @property
    @pulumi.getter(name="serviceUser")
    def service_user(self) -> str:
        return pulumi.get(self, "service_user")

    @property
    @pulumi.getter(name="sshKeyIds")
    def ssh_key_ids(self) -> Sequence[str]:
        return pulumi.get(self, "ssh_key_ids")

    @property
    @pulumi.getter
    def tags(self) -> Sequence[str]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def user(self) -> str:
        return pulumi.get(self, "user")

    @property
    @pulumi.getter
    def zone(self) -> Optional[str]:
        return pulumi.get(self, "zone")


class AwaitableGetBaremetalServerResult(GetBaremetalServerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBaremetalServerResult(
            description=self.description,
            domain=self.domain,
            hostname=self.hostname,
            id=self.id,
            ips=self.ips,
            name=self.name,
            offer=self.offer,
            offer_id=self.offer_id,
            offer_name=self.offer_name,
            options=self.options,
            organization_id=self.organization_id,
            os=self.os,
            os_name=self.os_name,
            password=self.password,
            private_networks=self.private_networks,
            project_id=self.project_id,
            reinstall_on_config_changes=self.reinstall_on_config_changes,
            server_id=self.server_id,
            service_password=self.service_password,
            service_user=self.service_user,
            ssh_key_ids=self.ssh_key_ids,
            tags=self.tags,
            user=self.user,
            zone=self.zone)


def get_baremetal_server(name: Optional[str] = None,
                         server_id: Optional[str] = None,
                         zone: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBaremetalServerResult:
    """
    Gets information about a baremetal server.
    For more information, see [the documentation](https://developers.scaleway.com/en/products/baremetal/api).

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    by_name = scaleway.get_baremetal_server(name="foobar",
        zone="fr-par-2")
    by_id = scaleway.get_baremetal_server(server_id="11111111-1111-1111-1111-111111111111")
    ```


    :param str name: The server name. Only one of `name` and `server_id` should be specified.
    :param str zone: `zone`) The zone in which the server exists.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['serverId'] = server_id
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:index/getBaremetalServer:getBaremetalServer', __args__, opts=opts, typ=GetBaremetalServerResult).value

    return AwaitableGetBaremetalServerResult(
        description=__ret__.description,
        domain=__ret__.domain,
        hostname=__ret__.hostname,
        id=__ret__.id,
        ips=__ret__.ips,
        name=__ret__.name,
        offer=__ret__.offer,
        offer_id=__ret__.offer_id,
        offer_name=__ret__.offer_name,
        options=__ret__.options,
        organization_id=__ret__.organization_id,
        os=__ret__.os,
        os_name=__ret__.os_name,
        password=__ret__.password,
        private_networks=__ret__.private_networks,
        project_id=__ret__.project_id,
        reinstall_on_config_changes=__ret__.reinstall_on_config_changes,
        server_id=__ret__.server_id,
        service_password=__ret__.service_password,
        service_user=__ret__.service_user,
        ssh_key_ids=__ret__.ssh_key_ids,
        tags=__ret__.tags,
        user=__ret__.user,
        zone=__ret__.zone)


@_utilities.lift_output_func(get_baremetal_server)
def get_baremetal_server_output(name: Optional[pulumi.Input[Optional[str]]] = None,
                                server_id: Optional[pulumi.Input[Optional[str]]] = None,
                                zone: Optional[pulumi.Input[Optional[str]]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetBaremetalServerResult]:
    """
    Gets information about a baremetal server.
    For more information, see [the documentation](https://developers.scaleway.com/en/products/baremetal/api).

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    by_name = scaleway.get_baremetal_server(name="foobar",
        zone="fr-par-2")
    by_id = scaleway.get_baremetal_server(server_id="11111111-1111-1111-1111-111111111111")
    ```


    :param str name: The server name. Only one of `name` and `server_id` should be specified.
    :param str zone: `zone`) The zone in which the server exists.
    """
    ...
