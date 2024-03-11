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
    'GetDocumentdbInstanceResult',
    'AwaitableGetDocumentdbInstanceResult',
    'get_documentdb_instance',
    'get_documentdb_instance_output',
]

@pulumi.output_type
class GetDocumentdbInstanceResult:
    """
    A collection of values returned by getDocumentdbInstance.
    """
    def __init__(__self__, engine=None, id=None, instance_id=None, is_ha_cluster=None, name=None, node_type=None, password=None, project_id=None, region=None, tags=None, telemetry_enabled=None, user_name=None, volume_size_in_gb=None, volume_type=None):
        if engine and not isinstance(engine, str):
            raise TypeError("Expected argument 'engine' to be a str")
        pulumi.set(__self__, "engine", engine)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if is_ha_cluster and not isinstance(is_ha_cluster, bool):
            raise TypeError("Expected argument 'is_ha_cluster' to be a bool")
        pulumi.set(__self__, "is_ha_cluster", is_ha_cluster)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if node_type and not isinstance(node_type, str):
            raise TypeError("Expected argument 'node_type' to be a str")
        pulumi.set(__self__, "node_type", node_type)
        if password and not isinstance(password, str):
            raise TypeError("Expected argument 'password' to be a str")
        pulumi.set(__self__, "password", password)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if telemetry_enabled and not isinstance(telemetry_enabled, bool):
            raise TypeError("Expected argument 'telemetry_enabled' to be a bool")
        pulumi.set(__self__, "telemetry_enabled", telemetry_enabled)
        if user_name and not isinstance(user_name, str):
            raise TypeError("Expected argument 'user_name' to be a str")
        pulumi.set(__self__, "user_name", user_name)
        if volume_size_in_gb and not isinstance(volume_size_in_gb, int):
            raise TypeError("Expected argument 'volume_size_in_gb' to be a int")
        pulumi.set(__self__, "volume_size_in_gb", volume_size_in_gb)
        if volume_type and not isinstance(volume_type, str):
            raise TypeError("Expected argument 'volume_type' to be a str")
        pulumi.set(__self__, "volume_type", volume_type)

    @property
    @pulumi.getter
    def engine(self) -> str:
        return pulumi.get(self, "engine")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[str]:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="isHaCluster")
    def is_ha_cluster(self) -> bool:
        return pulumi.get(self, "is_ha_cluster")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nodeType")
    def node_type(self) -> str:
        return pulumi.get(self, "node_type")

    @property
    @pulumi.getter
    def password(self) -> str:
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> Optional[str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def tags(self) -> Sequence[str]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="telemetryEnabled")
    def telemetry_enabled(self) -> bool:
        return pulumi.get(self, "telemetry_enabled")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> str:
        return pulumi.get(self, "user_name")

    @property
    @pulumi.getter(name="volumeSizeInGb")
    def volume_size_in_gb(self) -> int:
        return pulumi.get(self, "volume_size_in_gb")

    @property
    @pulumi.getter(name="volumeType")
    def volume_type(self) -> str:
        return pulumi.get(self, "volume_type")


class AwaitableGetDocumentdbInstanceResult(GetDocumentdbInstanceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDocumentdbInstanceResult(
            engine=self.engine,
            id=self.id,
            instance_id=self.instance_id,
            is_ha_cluster=self.is_ha_cluster,
            name=self.name,
            node_type=self.node_type,
            password=self.password,
            project_id=self.project_id,
            region=self.region,
            tags=self.tags,
            telemetry_enabled=self.telemetry_enabled,
            user_name=self.user_name,
            volume_size_in_gb=self.volume_size_in_gb,
            volume_type=self.volume_type)


def get_documentdb_instance(instance_id: Optional[str] = None,
                            name: Optional[str] = None,
                            region: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDocumentdbInstanceResult:
    """
    Gets information about an DocumentDB instance. For further information see our [developers website](https://www.scaleway.com/en/developers/api/document_db/)

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    db = scaleway.get_documentdb_instance(instance_id="11111111-1111-1111-1111-111111111111")
    ```
    <!--End PulumiCodeChooser -->


    :param str instance_id: The DocumentDB instance ID.
           Only one of `name` and `instance_id` should be specified.
    :param str name: The name of the DocumentDB instance.
           Only one of `name` and `instance_id` should be specified.
    :param str region: `region`) The region in which the DocumentDB instance exists.
    """
    __args__ = dict()
    __args__['instanceId'] = instance_id
    __args__['name'] = name
    __args__['region'] = region
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:index/getDocumentdbInstance:getDocumentdbInstance', __args__, opts=opts, typ=GetDocumentdbInstanceResult).value

    return AwaitableGetDocumentdbInstanceResult(
        engine=pulumi.get(__ret__, 'engine'),
        id=pulumi.get(__ret__, 'id'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        is_ha_cluster=pulumi.get(__ret__, 'is_ha_cluster'),
        name=pulumi.get(__ret__, 'name'),
        node_type=pulumi.get(__ret__, 'node_type'),
        password=pulumi.get(__ret__, 'password'),
        project_id=pulumi.get(__ret__, 'project_id'),
        region=pulumi.get(__ret__, 'region'),
        tags=pulumi.get(__ret__, 'tags'),
        telemetry_enabled=pulumi.get(__ret__, 'telemetry_enabled'),
        user_name=pulumi.get(__ret__, 'user_name'),
        volume_size_in_gb=pulumi.get(__ret__, 'volume_size_in_gb'),
        volume_type=pulumi.get(__ret__, 'volume_type'))


@_utilities.lift_output_func(get_documentdb_instance)
def get_documentdb_instance_output(instance_id: Optional[pulumi.Input[Optional[str]]] = None,
                                   name: Optional[pulumi.Input[Optional[str]]] = None,
                                   region: Optional[pulumi.Input[Optional[str]]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDocumentdbInstanceResult]:
    """
    Gets information about an DocumentDB instance. For further information see our [developers website](https://www.scaleway.com/en/developers/api/document_db/)

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    db = scaleway.get_documentdb_instance(instance_id="11111111-1111-1111-1111-111111111111")
    ```
    <!--End PulumiCodeChooser -->


    :param str instance_id: The DocumentDB instance ID.
           Only one of `name` and `instance_id` should be specified.
    :param str name: The name of the DocumentDB instance.
           Only one of `name` and `instance_id` should be specified.
    :param str region: `region`) The region in which the DocumentDB instance exists.
    """
    ...
