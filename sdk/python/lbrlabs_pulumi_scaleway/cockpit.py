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
from ._inputs import *

__all__ = ['CockpitArgs', 'Cockpit']

@pulumi.input_type
class CockpitArgs:
    def __init__(__self__, *,
                 project_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Cockpit resource.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the cockpit is associated with.
        """
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        `project_id`) The ID of the project the cockpit is associated with.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)


@pulumi.input_type
class _CockpitState:
    def __init__(__self__, *,
                 endpoints: Optional[pulumi.Input[Sequence[pulumi.Input['CockpitEndpointArgs']]]] = None,
                 project_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Cockpit resources.
        :param pulumi.Input[Sequence[pulumi.Input['CockpitEndpointArgs']]] endpoints: Endpoints
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the cockpit is associated with.
        """
        if endpoints is not None:
            pulumi.set(__self__, "endpoints", endpoints)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)

    @property
    @pulumi.getter
    def endpoints(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CockpitEndpointArgs']]]]:
        """
        Endpoints
        """
        return pulumi.get(self, "endpoints")

    @endpoints.setter
    def endpoints(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CockpitEndpointArgs']]]]):
        pulumi.set(self, "endpoints", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        `project_id`) The ID of the project the cockpit is associated with.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)


class Cockpit(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Import

        Cockpits can be imported using the `{project_id}`, e.g. bash

        ```sh
         $ pulumi import scaleway:index/cockpit:Cockpit main 11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the cockpit is associated with.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[CockpitArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        Cockpits can be imported using the `{project_id}`, e.g. bash

        ```sh
         $ pulumi import scaleway:index/cockpit:Cockpit main 11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param CockpitArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CockpitArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CockpitArgs.__new__(CockpitArgs)

            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["endpoints"] = None
        super(Cockpit, __self__).__init__(
            'scaleway:index/cockpit:Cockpit',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            endpoints: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CockpitEndpointArgs']]]]] = None,
            project_id: Optional[pulumi.Input[str]] = None) -> 'Cockpit':
        """
        Get an existing Cockpit resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CockpitEndpointArgs']]]] endpoints: Endpoints
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the cockpit is associated with.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CockpitState.__new__(_CockpitState)

        __props__.__dict__["endpoints"] = endpoints
        __props__.__dict__["project_id"] = project_id
        return Cockpit(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def endpoints(self) -> pulumi.Output[Sequence['outputs.CockpitEndpoint']]:
        """
        Endpoints
        """
        return pulumi.get(self, "endpoints")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        `project_id`) The ID of the project the cockpit is associated with.
        """
        return pulumi.get(self, "project_id")

