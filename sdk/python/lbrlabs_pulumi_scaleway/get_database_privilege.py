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
    'GetDatabasePrivilegeResult',
    'AwaitableGetDatabasePrivilegeResult',
    'get_database_privilege',
    'get_database_privilege_output',
]

@pulumi.output_type
class GetDatabasePrivilegeResult:
    """
    A collection of values returned by getDatabasePrivilege.
    """
    def __init__(__self__, database_name=None, id=None, instance_id=None, permission=None, region=None, user_name=None):
        if database_name and not isinstance(database_name, str):
            raise TypeError("Expected argument 'database_name' to be a str")
        pulumi.set(__self__, "database_name", database_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if permission and not isinstance(permission, str):
            raise TypeError("Expected argument 'permission' to be a str")
        pulumi.set(__self__, "permission", permission)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if user_name and not isinstance(user_name, str):
            raise TypeError("Expected argument 'user_name' to be a str")
        pulumi.set(__self__, "user_name", user_name)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> str:
        return pulumi.get(self, "database_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def permission(self) -> str:
        """
        The permission for this user on the database. Possible values are `readonly`, `readwrite`, `all`
        , `custom` and `none`.
        """
        return pulumi.get(self, "permission")

    @property
    @pulumi.getter
    def region(self) -> Optional[str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> str:
        return pulumi.get(self, "user_name")


class AwaitableGetDatabasePrivilegeResult(GetDatabasePrivilegeResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatabasePrivilegeResult(
            database_name=self.database_name,
            id=self.id,
            instance_id=self.instance_id,
            permission=self.permission,
            region=self.region,
            user_name=self.user_name)


def get_database_privilege(database_name: Optional[str] = None,
                           instance_id: Optional[str] = None,
                           region: Optional[str] = None,
                           user_name: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatabasePrivilegeResult:
    """
    Gets information about the privilege on RDB database.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    main = scaleway.get_database_privilege(database_name="my-database",
        instance_id="11111111-1111-111111111111",
        user_name="my-user")
    ```


    :param str database_name: The database name.
    :param str instance_id: The RDB instance ID.
    :param str region: `region`) The region in which the resource exists.
    :param str user_name: The user name.
    """
    __args__ = dict()
    __args__['databaseName'] = database_name
    __args__['instanceId'] = instance_id
    __args__['region'] = region
    __args__['userName'] = user_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:index/getDatabasePrivilege:getDatabasePrivilege', __args__, opts=opts, typ=GetDatabasePrivilegeResult).value

    return AwaitableGetDatabasePrivilegeResult(
        database_name=pulumi.get(__ret__, 'database_name'),
        id=pulumi.get(__ret__, 'id'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        permission=pulumi.get(__ret__, 'permission'),
        region=pulumi.get(__ret__, 'region'),
        user_name=pulumi.get(__ret__, 'user_name'))


@_utilities.lift_output_func(get_database_privilege)
def get_database_privilege_output(database_name: Optional[pulumi.Input[str]] = None,
                                  instance_id: Optional[pulumi.Input[str]] = None,
                                  region: Optional[pulumi.Input[Optional[str]]] = None,
                                  user_name: Optional[pulumi.Input[str]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDatabasePrivilegeResult]:
    """
    Gets information about the privilege on RDB database.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    main = scaleway.get_database_privilege(database_name="my-database",
        instance_id="11111111-1111-111111111111",
        user_name="my-user")
    ```


    :param str database_name: The database name.
    :param str instance_id: The RDB instance ID.
    :param str region: `region`) The region in which the resource exists.
    :param str user_name: The user name.
    """
    ...
