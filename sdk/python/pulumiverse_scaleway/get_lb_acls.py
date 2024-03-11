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
    'GetLbAclsResult',
    'AwaitableGetLbAclsResult',
    'get_lb_acls',
    'get_lb_acls_output',
]

@pulumi.output_type
class GetLbAclsResult:
    """
    A collection of values returned by getLbAcls.
    """
    def __init__(__self__, acls=None, frontend_id=None, id=None, name=None, organization_id=None, project_id=None, zone=None):
        if acls and not isinstance(acls, list):
            raise TypeError("Expected argument 'acls' to be a list")
        pulumi.set(__self__, "acls", acls)
        if frontend_id and not isinstance(frontend_id, str):
            raise TypeError("Expected argument 'frontend_id' to be a str")
        pulumi.set(__self__, "frontend_id", frontend_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if organization_id and not isinstance(organization_id, str):
            raise TypeError("Expected argument 'organization_id' to be a str")
        pulumi.set(__self__, "organization_id", organization_id)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def acls(self) -> Sequence['outputs.GetLbAclsAclResult']:
        """
        List of found ACLs
        """
        return pulumi.get(self, "acls")

    @property
    @pulumi.getter(name="frontendId")
    def frontend_id(self) -> str:
        return pulumi.get(self, "frontend_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

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
    def zone(self) -> str:
        return pulumi.get(self, "zone")


class AwaitableGetLbAclsResult(GetLbAclsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLbAclsResult(
            acls=self.acls,
            frontend_id=self.frontend_id,
            id=self.id,
            name=self.name,
            organization_id=self.organization_id,
            project_id=self.project_id,
            zone=self.zone)


def get_lb_acls(frontend_id: Optional[str] = None,
                name: Optional[str] = None,
                project_id: Optional[str] = None,
                zone: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLbAclsResult:
    """
    Gets information about multiple Load Balancer ACLs.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    by_front_id = scaleway.get_lb_acls(frontend_id=scaleway_lb_frontend["frt01"]["id"])
    by_front_id_and_name = scaleway.get_lb_acls(frontend_id=scaleway_lb_frontend["frt01"]["id"],
        name="tf-acls-datasource")
    ```
    <!--End PulumiCodeChooser -->


    :param str frontend_id: The frontend ID this ACL is attached to. ACLs with a frontend ID like it are listed.
           > **Important:** LB Frontends' IDs are zoned, which means they are of the form `{zone}/{id}`, e.g. `fr-par-1/11111111-1111-1111-1111-111111111111`
    :param str name: The ACL name used as filter. ACLs with a name like it are listed.
    :param str zone: `zone`) The zone in which ACLs exist.
    """
    __args__ = dict()
    __args__['frontendId'] = frontend_id
    __args__['name'] = name
    __args__['projectId'] = project_id
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:index/getLbAcls:getLbAcls', __args__, opts=opts, typ=GetLbAclsResult).value

    return AwaitableGetLbAclsResult(
        acls=pulumi.get(__ret__, 'acls'),
        frontend_id=pulumi.get(__ret__, 'frontend_id'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        organization_id=pulumi.get(__ret__, 'organization_id'),
        project_id=pulumi.get(__ret__, 'project_id'),
        zone=pulumi.get(__ret__, 'zone'))


@_utilities.lift_output_func(get_lb_acls)
def get_lb_acls_output(frontend_id: Optional[pulumi.Input[str]] = None,
                       name: Optional[pulumi.Input[Optional[str]]] = None,
                       project_id: Optional[pulumi.Input[Optional[str]]] = None,
                       zone: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetLbAclsResult]:
    """
    Gets information about multiple Load Balancer ACLs.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    by_front_id = scaleway.get_lb_acls(frontend_id=scaleway_lb_frontend["frt01"]["id"])
    by_front_id_and_name = scaleway.get_lb_acls(frontend_id=scaleway_lb_frontend["frt01"]["id"],
        name="tf-acls-datasource")
    ```
    <!--End PulumiCodeChooser -->


    :param str frontend_id: The frontend ID this ACL is attached to. ACLs with a frontend ID like it are listed.
           > **Important:** LB Frontends' IDs are zoned, which means they are of the form `{zone}/{id}`, e.g. `fr-par-1/11111111-1111-1111-1111-111111111111`
    :param str name: The ACL name used as filter. ACLs with a name like it are listed.
    :param str zone: `zone`) The zone in which ACLs exist.
    """
    ...
