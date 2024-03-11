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
    'GetFunctionNamespaceResult',
    'AwaitableGetFunctionNamespaceResult',
    'get_function_namespace',
    'get_function_namespace_output',
]

@pulumi.output_type
class GetFunctionNamespaceResult:
    """
    A collection of values returned by getFunctionNamespace.
    """
    def __init__(__self__, description=None, environment_variables=None, id=None, name=None, namespace_id=None, organization_id=None, project_id=None, region=None, registry_endpoint=None, registry_namespace_id=None, secret_environment_variables=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if environment_variables and not isinstance(environment_variables, dict):
            raise TypeError("Expected argument 'environment_variables' to be a dict")
        pulumi.set(__self__, "environment_variables", environment_variables)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if namespace_id and not isinstance(namespace_id, str):
            raise TypeError("Expected argument 'namespace_id' to be a str")
        pulumi.set(__self__, "namespace_id", namespace_id)
        if organization_id and not isinstance(organization_id, str):
            raise TypeError("Expected argument 'organization_id' to be a str")
        pulumi.set(__self__, "organization_id", organization_id)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if registry_endpoint and not isinstance(registry_endpoint, str):
            raise TypeError("Expected argument 'registry_endpoint' to be a str")
        pulumi.set(__self__, "registry_endpoint", registry_endpoint)
        if registry_namespace_id and not isinstance(registry_namespace_id, str):
            raise TypeError("Expected argument 'registry_namespace_id' to be a str")
        pulumi.set(__self__, "registry_namespace_id", registry_namespace_id)
        if secret_environment_variables and not isinstance(secret_environment_variables, dict):
            raise TypeError("Expected argument 'secret_environment_variables' to be a dict")
        pulumi.set(__self__, "secret_environment_variables", secret_environment_variables)

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of the namespace.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="environmentVariables")
    def environment_variables(self) -> Mapping[str, str]:
        """
        The environment variables of the namespace.
        """
        return pulumi.get(self, "environment_variables")

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
    @pulumi.getter(name="namespaceId")
    def namespace_id(self) -> Optional[str]:
        return pulumi.get(self, "namespace_id")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> str:
        """
        The organization ID the namespace is associated with.
        """
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> Optional[str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="registryEndpoint")
    def registry_endpoint(self) -> str:
        """
        The registry endpoint of the namespace.
        """
        return pulumi.get(self, "registry_endpoint")

    @property
    @pulumi.getter(name="registryNamespaceId")
    def registry_namespace_id(self) -> str:
        """
        The registry namespace ID of the namespace.
        """
        return pulumi.get(self, "registry_namespace_id")

    @property
    @pulumi.getter(name="secretEnvironmentVariables")
    def secret_environment_variables(self) -> Mapping[str, str]:
        return pulumi.get(self, "secret_environment_variables")


class AwaitableGetFunctionNamespaceResult(GetFunctionNamespaceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFunctionNamespaceResult(
            description=self.description,
            environment_variables=self.environment_variables,
            id=self.id,
            name=self.name,
            namespace_id=self.namespace_id,
            organization_id=self.organization_id,
            project_id=self.project_id,
            region=self.region,
            registry_endpoint=self.registry_endpoint,
            registry_namespace_id=self.registry_namespace_id,
            secret_environment_variables=self.secret_environment_variables)


def get_function_namespace(name: Optional[str] = None,
                           namespace_id: Optional[str] = None,
                           region: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFunctionNamespaceResult:
    """
    Gets information about a function namespace.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    my_namespace = scaleway.get_function_namespace(namespace_id="11111111-1111-1111-1111-111111111111")
    ```
    <!--End PulumiCodeChooser -->


    :param str name: The namespace name.
           Only one of `name` and `namespace_id` should be specified.
    :param str namespace_id: The namespace id.
           Only one of `name` and `namespace_id` should be specified.
    :param str region: `region`) The region in which the namespace exists.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['namespaceId'] = namespace_id
    __args__['region'] = region
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:index/getFunctionNamespace:getFunctionNamespace', __args__, opts=opts, typ=GetFunctionNamespaceResult).value

    return AwaitableGetFunctionNamespaceResult(
        description=pulumi.get(__ret__, 'description'),
        environment_variables=pulumi.get(__ret__, 'environment_variables'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        namespace_id=pulumi.get(__ret__, 'namespace_id'),
        organization_id=pulumi.get(__ret__, 'organization_id'),
        project_id=pulumi.get(__ret__, 'project_id'),
        region=pulumi.get(__ret__, 'region'),
        registry_endpoint=pulumi.get(__ret__, 'registry_endpoint'),
        registry_namespace_id=pulumi.get(__ret__, 'registry_namespace_id'),
        secret_environment_variables=pulumi.get(__ret__, 'secret_environment_variables'))


@_utilities.lift_output_func(get_function_namespace)
def get_function_namespace_output(name: Optional[pulumi.Input[Optional[str]]] = None,
                                  namespace_id: Optional[pulumi.Input[Optional[str]]] = None,
                                  region: Optional[pulumi.Input[Optional[str]]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFunctionNamespaceResult]:
    """
    Gets information about a function namespace.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    my_namespace = scaleway.get_function_namespace(namespace_id="11111111-1111-1111-1111-111111111111")
    ```
    <!--End PulumiCodeChooser -->


    :param str name: The namespace name.
           Only one of `name` and `namespace_id` should be specified.
    :param str namespace_id: The namespace id.
           Only one of `name` and `namespace_id` should be specified.
    :param str region: `region`) The region in which the namespace exists.
    """
    ...
