# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ContainerNamespaceArgs', 'ContainerNamespace']

@pulumi.input_type
class ContainerNamespaceArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 destroy_registry: Optional[pulumi.Input[bool]] = None,
                 environment_variables: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 secret_environment_variables: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a ContainerNamespace resource.
        :param pulumi.Input[str] description: The description of the namespace.
        :param pulumi.Input[bool] destroy_registry: Destroy registry on deletion
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] environment_variables: The environment variables of the namespace.
        :param pulumi.Input[str] name: The unique name of the container namespace.
               
               > **Important** Updates to `name` will recreate the namespace.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the namespace is associated with.
        :param pulumi.Input[str] region: `region`). The region in which the namespace should be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] secret_environment_variables: The secret environment variables of the namespace.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if destroy_registry is not None:
            warnings.warn("""Registry namespace is automatically destroyed with namespace""", DeprecationWarning)
            pulumi.log.warn("""destroy_registry is deprecated: Registry namespace is automatically destroyed with namespace""")
        if destroy_registry is not None:
            pulumi.set(__self__, "destroy_registry", destroy_registry)
        if environment_variables is not None:
            pulumi.set(__self__, "environment_variables", environment_variables)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if secret_environment_variables is not None:
            pulumi.set(__self__, "secret_environment_variables", secret_environment_variables)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the namespace.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="destroyRegistry")
    def destroy_registry(self) -> Optional[pulumi.Input[bool]]:
        """
        Destroy registry on deletion
        """
        return pulumi.get(self, "destroy_registry")

    @destroy_registry.setter
    def destroy_registry(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "destroy_registry", value)

    @property
    @pulumi.getter(name="environmentVariables")
    def environment_variables(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        The environment variables of the namespace.
        """
        return pulumi.get(self, "environment_variables")

    @environment_variables.setter
    def environment_variables(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "environment_variables", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The unique name of the container namespace.

        > **Important** Updates to `name` will recreate the namespace.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        `project_id`) The ID of the project the namespace is associated with.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`). The region in which the namespace should be created.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="secretEnvironmentVariables")
    def secret_environment_variables(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        The secret environment variables of the namespace.
        """
        return pulumi.get(self, "secret_environment_variables")

    @secret_environment_variables.setter
    def secret_environment_variables(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "secret_environment_variables", value)


@pulumi.input_type
class _ContainerNamespaceState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 destroy_registry: Optional[pulumi.Input[bool]] = None,
                 environment_variables: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 registry_endpoint: Optional[pulumi.Input[str]] = None,
                 registry_namespace_id: Optional[pulumi.Input[str]] = None,
                 secret_environment_variables: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering ContainerNamespace resources.
        :param pulumi.Input[str] description: The description of the namespace.
        :param pulumi.Input[bool] destroy_registry: Destroy registry on deletion
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] environment_variables: The environment variables of the namespace.
        :param pulumi.Input[str] name: The unique name of the container namespace.
               
               > **Important** Updates to `name` will recreate the namespace.
        :param pulumi.Input[str] organization_id: The organization ID the namespace is associated with.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the namespace is associated with.
        :param pulumi.Input[str] region: `region`). The region in which the namespace should be created.
        :param pulumi.Input[str] registry_endpoint: The registry endpoint of the namespace.
        :param pulumi.Input[str] registry_namespace_id: The registry namespace ID of the namespace.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] secret_environment_variables: The secret environment variables of the namespace.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if destroy_registry is not None:
            warnings.warn("""Registry namespace is automatically destroyed with namespace""", DeprecationWarning)
            pulumi.log.warn("""destroy_registry is deprecated: Registry namespace is automatically destroyed with namespace""")
        if destroy_registry is not None:
            pulumi.set(__self__, "destroy_registry", destroy_registry)
        if environment_variables is not None:
            pulumi.set(__self__, "environment_variables", environment_variables)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if organization_id is not None:
            pulumi.set(__self__, "organization_id", organization_id)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if registry_endpoint is not None:
            pulumi.set(__self__, "registry_endpoint", registry_endpoint)
        if registry_namespace_id is not None:
            pulumi.set(__self__, "registry_namespace_id", registry_namespace_id)
        if secret_environment_variables is not None:
            pulumi.set(__self__, "secret_environment_variables", secret_environment_variables)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the namespace.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="destroyRegistry")
    def destroy_registry(self) -> Optional[pulumi.Input[bool]]:
        """
        Destroy registry on deletion
        """
        return pulumi.get(self, "destroy_registry")

    @destroy_registry.setter
    def destroy_registry(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "destroy_registry", value)

    @property
    @pulumi.getter(name="environmentVariables")
    def environment_variables(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        The environment variables of the namespace.
        """
        return pulumi.get(self, "environment_variables")

    @environment_variables.setter
    def environment_variables(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "environment_variables", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The unique name of the container namespace.

        > **Important** Updates to `name` will recreate the namespace.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> Optional[pulumi.Input[str]]:
        """
        The organization ID the namespace is associated with.
        """
        return pulumi.get(self, "organization_id")

    @organization_id.setter
    def organization_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "organization_id", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        `project_id`) The ID of the project the namespace is associated with.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`). The region in which the namespace should be created.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="registryEndpoint")
    def registry_endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        The registry endpoint of the namespace.
        """
        return pulumi.get(self, "registry_endpoint")

    @registry_endpoint.setter
    def registry_endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "registry_endpoint", value)

    @property
    @pulumi.getter(name="registryNamespaceId")
    def registry_namespace_id(self) -> Optional[pulumi.Input[str]]:
        """
        The registry namespace ID of the namespace.
        """
        return pulumi.get(self, "registry_namespace_id")

    @registry_namespace_id.setter
    def registry_namespace_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "registry_namespace_id", value)

    @property
    @pulumi.getter(name="secretEnvironmentVariables")
    def secret_environment_variables(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        The secret environment variables of the namespace.
        """
        return pulumi.get(self, "secret_environment_variables")

    @secret_environment_variables.setter
    def secret_environment_variables(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "secret_environment_variables", value)


class ContainerNamespace(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 destroy_registry: Optional[pulumi.Input[bool]] = None,
                 environment_variables: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 secret_environment_variables: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Creates and manages Scaleway Serverless Container Namespace.
        For more information see [the documentation](https://developers.scaleway.com/en/products/containers/api/#namespaces-cdce79).

        ## Examples

        ### Basic

        ```python
        import pulumi
        import lbrlabs_pulumi_scaleway as scaleway

        main = scaleway.ContainerNamespace("main", description="Main container namespace")
        ```

        ## Import

        Namespaces can be imported using the `{region}/{id}`, e.g. bash

        ```sh
         $ pulumi import scaleway:index/containerNamespace:ContainerNamespace main fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the namespace.
        :param pulumi.Input[bool] destroy_registry: Destroy registry on deletion
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] environment_variables: The environment variables of the namespace.
        :param pulumi.Input[str] name: The unique name of the container namespace.
               
               > **Important** Updates to `name` will recreate the namespace.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the namespace is associated with.
        :param pulumi.Input[str] region: `region`). The region in which the namespace should be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] secret_environment_variables: The secret environment variables of the namespace.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ContainerNamespaceArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages Scaleway Serverless Container Namespace.
        For more information see [the documentation](https://developers.scaleway.com/en/products/containers/api/#namespaces-cdce79).

        ## Examples

        ### Basic

        ```python
        import pulumi
        import lbrlabs_pulumi_scaleway as scaleway

        main = scaleway.ContainerNamespace("main", description="Main container namespace")
        ```

        ## Import

        Namespaces can be imported using the `{region}/{id}`, e.g. bash

        ```sh
         $ pulumi import scaleway:index/containerNamespace:ContainerNamespace main fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param ContainerNamespaceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ContainerNamespaceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 destroy_registry: Optional[pulumi.Input[bool]] = None,
                 environment_variables: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 secret_environment_variables: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ContainerNamespaceArgs.__new__(ContainerNamespaceArgs)

            __props__.__dict__["description"] = description
            if destroy_registry is not None and not opts.urn:
                warnings.warn("""Registry namespace is automatically destroyed with namespace""", DeprecationWarning)
                pulumi.log.warn("""destroy_registry is deprecated: Registry namespace is automatically destroyed with namespace""")
            __props__.__dict__["destroy_registry"] = destroy_registry
            __props__.__dict__["environment_variables"] = environment_variables
            __props__.__dict__["name"] = name
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["region"] = region
            __props__.__dict__["secret_environment_variables"] = None if secret_environment_variables is None else pulumi.Output.secret(secret_environment_variables)
            __props__.__dict__["organization_id"] = None
            __props__.__dict__["registry_endpoint"] = None
            __props__.__dict__["registry_namespace_id"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["secretEnvironmentVariables"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(ContainerNamespace, __self__).__init__(
            'scaleway:index/containerNamespace:ContainerNamespace',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            destroy_registry: Optional[pulumi.Input[bool]] = None,
            environment_variables: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            organization_id: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            registry_endpoint: Optional[pulumi.Input[str]] = None,
            registry_namespace_id: Optional[pulumi.Input[str]] = None,
            secret_environment_variables: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'ContainerNamespace':
        """
        Get an existing ContainerNamespace resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the namespace.
        :param pulumi.Input[bool] destroy_registry: Destroy registry on deletion
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] environment_variables: The environment variables of the namespace.
        :param pulumi.Input[str] name: The unique name of the container namespace.
               
               > **Important** Updates to `name` will recreate the namespace.
        :param pulumi.Input[str] organization_id: The organization ID the namespace is associated with.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the namespace is associated with.
        :param pulumi.Input[str] region: `region`). The region in which the namespace should be created.
        :param pulumi.Input[str] registry_endpoint: The registry endpoint of the namespace.
        :param pulumi.Input[str] registry_namespace_id: The registry namespace ID of the namespace.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] secret_environment_variables: The secret environment variables of the namespace.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ContainerNamespaceState.__new__(_ContainerNamespaceState)

        __props__.__dict__["description"] = description
        __props__.__dict__["destroy_registry"] = destroy_registry
        __props__.__dict__["environment_variables"] = environment_variables
        __props__.__dict__["name"] = name
        __props__.__dict__["organization_id"] = organization_id
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["region"] = region
        __props__.__dict__["registry_endpoint"] = registry_endpoint
        __props__.__dict__["registry_namespace_id"] = registry_namespace_id
        __props__.__dict__["secret_environment_variables"] = secret_environment_variables
        return ContainerNamespace(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the namespace.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="destroyRegistry")
    def destroy_registry(self) -> pulumi.Output[Optional[bool]]:
        """
        Destroy registry on deletion
        """
        return pulumi.get(self, "destroy_registry")

    @property
    @pulumi.getter(name="environmentVariables")
    def environment_variables(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        The environment variables of the namespace.
        """
        return pulumi.get(self, "environment_variables")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The unique name of the container namespace.

        > **Important** Updates to `name` will recreate the namespace.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> pulumi.Output[str]:
        """
        The organization ID the namespace is associated with.
        """
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        `project_id`) The ID of the project the namespace is associated with.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        `region`). The region in which the namespace should be created.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="registryEndpoint")
    def registry_endpoint(self) -> pulumi.Output[str]:
        """
        The registry endpoint of the namespace.
        """
        return pulumi.get(self, "registry_endpoint")

    @property
    @pulumi.getter(name="registryNamespaceId")
    def registry_namespace_id(self) -> pulumi.Output[str]:
        """
        The registry namespace ID of the namespace.
        """
        return pulumi.get(self, "registry_namespace_id")

    @property
    @pulumi.getter(name="secretEnvironmentVariables")
    def secret_environment_variables(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        The secret environment variables of the namespace.
        """
        return pulumi.get(self, "secret_environment_variables")

