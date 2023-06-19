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
    'GetCockpitPlanResult',
    'AwaitableGetCockpitPlanResult',
    'get_cockpit_plan',
    'get_cockpit_plan_output',
]

@pulumi.output_type
class GetCockpitPlanResult:
    """
    A collection of values returned by getCockpitPlan.
    """
    def __init__(__self__, id=None, name=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")


class AwaitableGetCockpitPlanResult(GetCockpitPlanResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCockpitPlanResult(
            id=self.id,
            name=self.name)


def get_cockpit_plan(name: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCockpitPlanResult:
    """
    Gets information about a Scaleway Cockpit plan.

    ## Example Usage

    ```python
    import pulumi
    import lbrlabs_pulumi_scaleway as scaleway
    import pulumi_scaleway as scaleway

    premium = scaleway.get_cockpit_plan(name="premium")
    main = scaleway.Cockpit("main", plan=premium.id)
    ```


    :param str name: The name of the plan.
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:index/getCockpitPlan:getCockpitPlan', __args__, opts=opts, typ=GetCockpitPlanResult).value

    return AwaitableGetCockpitPlanResult(
        id=__ret__.id,
        name=__ret__.name)


@_utilities.lift_output_func(get_cockpit_plan)
def get_cockpit_plan_output(name: Optional[pulumi.Input[str]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetCockpitPlanResult]:
    """
    Gets information about a Scaleway Cockpit plan.

    ## Example Usage

    ```python
    import pulumi
    import lbrlabs_pulumi_scaleway as scaleway
    import pulumi_scaleway as scaleway

    premium = scaleway.get_cockpit_plan(name="premium")
    main = scaleway.Cockpit("main", plan=premium.id)
    ```


    :param str name: The name of the plan.
    """
    ...
