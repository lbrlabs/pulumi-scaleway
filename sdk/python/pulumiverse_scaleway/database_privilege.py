# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['DatabasePrivilegeArgs', 'DatabasePrivilege']

@pulumi.input_type
class DatabasePrivilegeArgs:
    def __init__(__self__, *,
                 database_name: pulumi.Input[str],
                 instance_id: pulumi.Input[str],
                 permission: pulumi.Input[str],
                 user_name: pulumi.Input[str],
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DatabasePrivilege resource.
        :param pulumi.Input[str] database_name: Name of the database (e.g. `my-db-name`).
        :param pulumi.Input[str] instance_id: UUID of the rdb instance.
        :param pulumi.Input[str] permission: Permission to set. Valid values are `readonly`, `readwrite`, `all`, `custom` and `none`.
        :param pulumi.Input[str] user_name: Name of the user (e.g. `my-db-user`).
        :param pulumi.Input[str] region: `region`) The region in which the resource exists.
        """
        pulumi.set(__self__, "database_name", database_name)
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "permission", permission)
        pulumi.set(__self__, "user_name", user_name)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> pulumi.Input[str]:
        """
        Name of the database (e.g. `my-db-name`).
        """
        return pulumi.get(self, "database_name")

    @database_name.setter
    def database_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "database_name", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        UUID of the rdb instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def permission(self) -> pulumi.Input[str]:
        """
        Permission to set. Valid values are `readonly`, `readwrite`, `all`, `custom` and `none`.
        """
        return pulumi.get(self, "permission")

    @permission.setter
    def permission(self, value: pulumi.Input[str]):
        pulumi.set(self, "permission", value)

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> pulumi.Input[str]:
        """
        Name of the user (e.g. `my-db-user`).
        """
        return pulumi.get(self, "user_name")

    @user_name.setter
    def user_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "user_name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`) The region in which the resource exists.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _DatabasePrivilegeState:
    def __init__(__self__, *,
                 database_name: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 permission: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 user_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DatabasePrivilege resources.
        :param pulumi.Input[str] database_name: Name of the database (e.g. `my-db-name`).
        :param pulumi.Input[str] instance_id: UUID of the rdb instance.
        :param pulumi.Input[str] permission: Permission to set. Valid values are `readonly`, `readwrite`, `all`, `custom` and `none`.
        :param pulumi.Input[str] region: `region`) The region in which the resource exists.
        :param pulumi.Input[str] user_name: Name of the user (e.g. `my-db-user`).
        """
        if database_name is not None:
            pulumi.set(__self__, "database_name", database_name)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if permission is not None:
            pulumi.set(__self__, "permission", permission)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if user_name is not None:
            pulumi.set(__self__, "user_name", user_name)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the database (e.g. `my-db-name`).
        """
        return pulumi.get(self, "database_name")

    @database_name.setter
    def database_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "database_name", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        UUID of the rdb instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def permission(self) -> Optional[pulumi.Input[str]]:
        """
        Permission to set. Valid values are `readonly`, `readwrite`, `all`, `custom` and `none`.
        """
        return pulumi.get(self, "permission")

    @permission.setter
    def permission(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "permission", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`) The region in which the resource exists.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the user (e.g. `my-db-user`).
        """
        return pulumi.get(self, "user_name")

    @user_name.setter
    def user_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_name", value)


class DatabasePrivilege(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 permission: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 user_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create and manage Scaleway RDB database privilege.
        For more information, see [the documentation](https://developers.scaleway.com/en/products/rdb/api/#user-and-permissions).

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        main_database_instance = scaleway.DatabaseInstance("mainDatabaseInstance",
            node_type="DB-DEV-S",
            engine="PostgreSQL-11",
            is_ha_cluster=True,
            disable_backup=True,
            user_name="my_initial_user",
            password="thiZ_is_v&ry_s3cret")
        main_database = scaleway.Database("mainDatabase", instance_id=main_database_instance.id)
        main_database_user = scaleway.DatabaseUser("mainDatabaseUser",
            instance_id=main_database_instance.id,
            password="thiZ_is_v&ry_s3cret",
            is_admin=False)
        main_database_privilege = scaleway.DatabasePrivilege("mainDatabasePrivilege",
            instance_id=main_database_instance.id,
            user_name=main_database_user.name,
            database_name=main_database.name,
            permission="all")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        The user privileges can be imported using the `{region}/{instance_id}/{database_name}/{user_name}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/databasePrivilege:DatabasePrivilege o fr-par/11111111-1111-1111-1111-111111111111/database_name/foo
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] database_name: Name of the database (e.g. `my-db-name`).
        :param pulumi.Input[str] instance_id: UUID of the rdb instance.
        :param pulumi.Input[str] permission: Permission to set. Valid values are `readonly`, `readwrite`, `all`, `custom` and `none`.
        :param pulumi.Input[str] region: `region`) The region in which the resource exists.
        :param pulumi.Input[str] user_name: Name of the user (e.g. `my-db-user`).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DatabasePrivilegeArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create and manage Scaleway RDB database privilege.
        For more information, see [the documentation](https://developers.scaleway.com/en/products/rdb/api/#user-and-permissions).

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        main_database_instance = scaleway.DatabaseInstance("mainDatabaseInstance",
            node_type="DB-DEV-S",
            engine="PostgreSQL-11",
            is_ha_cluster=True,
            disable_backup=True,
            user_name="my_initial_user",
            password="thiZ_is_v&ry_s3cret")
        main_database = scaleway.Database("mainDatabase", instance_id=main_database_instance.id)
        main_database_user = scaleway.DatabaseUser("mainDatabaseUser",
            instance_id=main_database_instance.id,
            password="thiZ_is_v&ry_s3cret",
            is_admin=False)
        main_database_privilege = scaleway.DatabasePrivilege("mainDatabasePrivilege",
            instance_id=main_database_instance.id,
            user_name=main_database_user.name,
            database_name=main_database.name,
            permission="all")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        The user privileges can be imported using the `{region}/{instance_id}/{database_name}/{user_name}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/databasePrivilege:DatabasePrivilege o fr-par/11111111-1111-1111-1111-111111111111/database_name/foo
        ```

        :param str resource_name: The name of the resource.
        :param DatabasePrivilegeArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DatabasePrivilegeArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 permission: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 user_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DatabasePrivilegeArgs.__new__(DatabasePrivilegeArgs)

            if database_name is None and not opts.urn:
                raise TypeError("Missing required property 'database_name'")
            __props__.__dict__["database_name"] = database_name
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            if permission is None and not opts.urn:
                raise TypeError("Missing required property 'permission'")
            __props__.__dict__["permission"] = permission
            __props__.__dict__["region"] = region
            if user_name is None and not opts.urn:
                raise TypeError("Missing required property 'user_name'")
            __props__.__dict__["user_name"] = user_name
        super(DatabasePrivilege, __self__).__init__(
            'scaleway:index/databasePrivilege:DatabasePrivilege',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            database_name: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            permission: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            user_name: Optional[pulumi.Input[str]] = None) -> 'DatabasePrivilege':
        """
        Get an existing DatabasePrivilege resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] database_name: Name of the database (e.g. `my-db-name`).
        :param pulumi.Input[str] instance_id: UUID of the rdb instance.
        :param pulumi.Input[str] permission: Permission to set. Valid values are `readonly`, `readwrite`, `all`, `custom` and `none`.
        :param pulumi.Input[str] region: `region`) The region in which the resource exists.
        :param pulumi.Input[str] user_name: Name of the user (e.g. `my-db-user`).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DatabasePrivilegeState.__new__(_DatabasePrivilegeState)

        __props__.__dict__["database_name"] = database_name
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["permission"] = permission
        __props__.__dict__["region"] = region
        __props__.__dict__["user_name"] = user_name
        return DatabasePrivilege(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> pulumi.Output[str]:
        """
        Name of the database (e.g. `my-db-name`).
        """
        return pulumi.get(self, "database_name")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        UUID of the rdb instance.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def permission(self) -> pulumi.Output[str]:
        """
        Permission to set. Valid values are `readonly`, `readwrite`, `all`, `custom` and `none`.
        """
        return pulumi.get(self, "permission")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        `region`) The region in which the resource exists.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> pulumi.Output[str]:
        """
        Name of the user (e.g. `my-db-user`).
        """
        return pulumi.get(self, "user_name")

