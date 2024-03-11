# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['InstancePrivateNicArgs', 'InstancePrivateNic']

@pulumi.input_type
class InstancePrivateNicArgs:
    def __init__(__self__, *,
                 private_network_id: pulumi.Input[str],
                 server_id: pulumi.Input[str],
                 ip_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a InstancePrivateNic resource.
        :param pulumi.Input[str] private_network_id: The ID of the private network attached to.
        :param pulumi.Input[str] server_id: The ID of the server associated with.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_ids: IPAM ip list, should be for internal use only
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: The tags associated with the private NIC.
        :param pulumi.Input[str] zone: `zone`) The zone in which the server must be created.
        """
        pulumi.set(__self__, "private_network_id", private_network_id)
        pulumi.set(__self__, "server_id", server_id)
        if ip_ids is not None:
            pulumi.set(__self__, "ip_ids", ip_ids)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="privateNetworkId")
    def private_network_id(self) -> pulumi.Input[str]:
        """
        The ID of the private network attached to.
        """
        return pulumi.get(self, "private_network_id")

    @private_network_id.setter
    def private_network_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "private_network_id", value)

    @property
    @pulumi.getter(name="serverId")
    def server_id(self) -> pulumi.Input[str]:
        """
        The ID of the server associated with.
        """
        return pulumi.get(self, "server_id")

    @server_id.setter
    def server_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "server_id", value)

    @property
    @pulumi.getter(name="ipIds")
    def ip_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        IPAM ip list, should be for internal use only
        """
        return pulumi.get(self, "ip_ids")

    @ip_ids.setter
    def ip_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ip_ids", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The tags associated with the private NIC.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        `zone`) The zone in which the server must be created.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


@pulumi.input_type
class _InstancePrivateNicState:
    def __init__(__self__, *,
                 ip_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 mac_address: Optional[pulumi.Input[str]] = None,
                 private_network_id: Optional[pulumi.Input[str]] = None,
                 server_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering InstancePrivateNic resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_ids: IPAM ip list, should be for internal use only
        :param pulumi.Input[str] mac_address: The MAC address of the private NIC.
        :param pulumi.Input[str] private_network_id: The ID of the private network attached to.
        :param pulumi.Input[str] server_id: The ID of the server associated with.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: The tags associated with the private NIC.
        :param pulumi.Input[str] zone: `zone`) The zone in which the server must be created.
        """
        if ip_ids is not None:
            pulumi.set(__self__, "ip_ids", ip_ids)
        if mac_address is not None:
            pulumi.set(__self__, "mac_address", mac_address)
        if private_network_id is not None:
            pulumi.set(__self__, "private_network_id", private_network_id)
        if server_id is not None:
            pulumi.set(__self__, "server_id", server_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="ipIds")
    def ip_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        IPAM ip list, should be for internal use only
        """
        return pulumi.get(self, "ip_ids")

    @ip_ids.setter
    def ip_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ip_ids", value)

    @property
    @pulumi.getter(name="macAddress")
    def mac_address(self) -> Optional[pulumi.Input[str]]:
        """
        The MAC address of the private NIC.
        """
        return pulumi.get(self, "mac_address")

    @mac_address.setter
    def mac_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mac_address", value)

    @property
    @pulumi.getter(name="privateNetworkId")
    def private_network_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the private network attached to.
        """
        return pulumi.get(self, "private_network_id")

    @private_network_id.setter
    def private_network_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_network_id", value)

    @property
    @pulumi.getter(name="serverId")
    def server_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the server associated with.
        """
        return pulumi.get(self, "server_id")

    @server_id.setter
    def server_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The tags associated with the private NIC.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        `zone`) The zone in which the server must be created.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


class InstancePrivateNic(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ip_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 private_network_id: Optional[pulumi.Input[str]] = None,
                 server_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates and manages Scaleway Instance Private NICs. For more information, see
        [the documentation](https://developers.scaleway.com/en/products/instance/api/#private-nics-a42eea).

        ## Examples

        ### Basic

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        pnic01 = scaleway.InstancePrivateNic("pnic01",
            private_network_id="fr-par-1/aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
            server_id="fr-par-1/11111111-1111-1111-1111-111111111111")
        ```
        <!--End PulumiCodeChooser -->

        ### With zone

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        pn01 = scaleway.VpcPrivateNetwork("pn01", zone="fr-par-2")
        base = scaleway.InstanceServer("base",
            image="ubuntu_jammy",
            type="DEV1-S",
            zone=pn01.zone)
        pnic01 = scaleway.InstancePrivateNic("pnic01",
            server_id=base.id,
            private_network_id=pn01.id,
            zone=pn01.zone)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Private NICs can be imported using the `{zone}/{server_id}/{private_nic_id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/instancePrivateNic:InstancePrivateNic pnic01 fr-par-1/11111111-1111-1111-1111-111111111111/22222222-2222-2222-2222-222222222222
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_ids: IPAM ip list, should be for internal use only
        :param pulumi.Input[str] private_network_id: The ID of the private network attached to.
        :param pulumi.Input[str] server_id: The ID of the server associated with.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: The tags associated with the private NIC.
        :param pulumi.Input[str] zone: `zone`) The zone in which the server must be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InstancePrivateNicArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages Scaleway Instance Private NICs. For more information, see
        [the documentation](https://developers.scaleway.com/en/products/instance/api/#private-nics-a42eea).

        ## Examples

        ### Basic

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        pnic01 = scaleway.InstancePrivateNic("pnic01",
            private_network_id="fr-par-1/aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
            server_id="fr-par-1/11111111-1111-1111-1111-111111111111")
        ```
        <!--End PulumiCodeChooser -->

        ### With zone

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        pn01 = scaleway.VpcPrivateNetwork("pn01", zone="fr-par-2")
        base = scaleway.InstanceServer("base",
            image="ubuntu_jammy",
            type="DEV1-S",
            zone=pn01.zone)
        pnic01 = scaleway.InstancePrivateNic("pnic01",
            server_id=base.id,
            private_network_id=pn01.id,
            zone=pn01.zone)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Private NICs can be imported using the `{zone}/{server_id}/{private_nic_id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/instancePrivateNic:InstancePrivateNic pnic01 fr-par-1/11111111-1111-1111-1111-111111111111/22222222-2222-2222-2222-222222222222
        ```

        :param str resource_name: The name of the resource.
        :param InstancePrivateNicArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InstancePrivateNicArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ip_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 private_network_id: Optional[pulumi.Input[str]] = None,
                 server_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = InstancePrivateNicArgs.__new__(InstancePrivateNicArgs)

            __props__.__dict__["ip_ids"] = ip_ids
            if private_network_id is None and not opts.urn:
                raise TypeError("Missing required property 'private_network_id'")
            __props__.__dict__["private_network_id"] = private_network_id
            if server_id is None and not opts.urn:
                raise TypeError("Missing required property 'server_id'")
            __props__.__dict__["server_id"] = server_id
            __props__.__dict__["tags"] = tags
            __props__.__dict__["zone"] = zone
            __props__.__dict__["mac_address"] = None
        super(InstancePrivateNic, __self__).__init__(
            'scaleway:index/instancePrivateNic:InstancePrivateNic',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            ip_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            mac_address: Optional[pulumi.Input[str]] = None,
            private_network_id: Optional[pulumi.Input[str]] = None,
            server_id: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            zone: Optional[pulumi.Input[str]] = None) -> 'InstancePrivateNic':
        """
        Get an existing InstancePrivateNic resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_ids: IPAM ip list, should be for internal use only
        :param pulumi.Input[str] mac_address: The MAC address of the private NIC.
        :param pulumi.Input[str] private_network_id: The ID of the private network attached to.
        :param pulumi.Input[str] server_id: The ID of the server associated with.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: The tags associated with the private NIC.
        :param pulumi.Input[str] zone: `zone`) The zone in which the server must be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InstancePrivateNicState.__new__(_InstancePrivateNicState)

        __props__.__dict__["ip_ids"] = ip_ids
        __props__.__dict__["mac_address"] = mac_address
        __props__.__dict__["private_network_id"] = private_network_id
        __props__.__dict__["server_id"] = server_id
        __props__.__dict__["tags"] = tags
        __props__.__dict__["zone"] = zone
        return InstancePrivateNic(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="ipIds")
    def ip_ids(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        IPAM ip list, should be for internal use only
        """
        return pulumi.get(self, "ip_ids")

    @property
    @pulumi.getter(name="macAddress")
    def mac_address(self) -> pulumi.Output[str]:
        """
        The MAC address of the private NIC.
        """
        return pulumi.get(self, "mac_address")

    @property
    @pulumi.getter(name="privateNetworkId")
    def private_network_id(self) -> pulumi.Output[str]:
        """
        The ID of the private network attached to.
        """
        return pulumi.get(self, "private_network_id")

    @property
    @pulumi.getter(name="serverId")
    def server_id(self) -> pulumi.Output[str]:
        """
        The ID of the server associated with.
        """
        return pulumi.get(self, "server_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The tags associated with the private NIC.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[str]:
        """
        `zone`) The zone in which the server must be created.
        """
        return pulumi.get(self, "zone")

