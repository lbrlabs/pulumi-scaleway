# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['DomainRecordArgs', 'DomainRecord']

@pulumi.input_type
class DomainRecordArgs:
    def __init__(__self__, *,
                 data: pulumi.Input[str],
                 dns_zone: pulumi.Input[str],
                 type: pulumi.Input[str],
                 geo_ip: Optional[pulumi.Input['DomainRecordGeoIpArgs']] = None,
                 http_service: Optional[pulumi.Input['DomainRecordHttpServiceArgs']] = None,
                 keep_empty_zone: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 ttl: Optional[pulumi.Input[int]] = None,
                 views: Optional[pulumi.Input[Sequence[pulumi.Input['DomainRecordViewArgs']]]] = None,
                 weighteds: Optional[pulumi.Input[Sequence[pulumi.Input['DomainRecordWeightedArgs']]]] = None):
        """
        The set of arguments for constructing a DomainRecord resource.
        :param pulumi.Input[str] data: The data of the record
        :param pulumi.Input[str] dns_zone: The zone you want to add the record in
        :param pulumi.Input[str] type: The type of the record
        :param pulumi.Input['DomainRecordGeoIpArgs'] geo_ip: Return record based on client localisation
        :param pulumi.Input['DomainRecordHttpServiceArgs'] http_service: Return record based on client localisation
        :param pulumi.Input[bool] keep_empty_zone: When destroy a resource record, if a zone have only NS, delete the zone
        :param pulumi.Input[str] name: The name of the record
        :param pulumi.Input[int] priority: The priority of the record
        :param pulumi.Input[str] project_id: The project_id you want to attach the resource to
        :param pulumi.Input[int] ttl: The ttl of the record
        :param pulumi.Input[Sequence[pulumi.Input['DomainRecordViewArgs']]] views: Return record based on client subnet
        :param pulumi.Input[Sequence[pulumi.Input['DomainRecordWeightedArgs']]] weighteds: Return record based on weight
        """
        pulumi.set(__self__, "data", data)
        pulumi.set(__self__, "dns_zone", dns_zone)
        pulumi.set(__self__, "type", type)
        if geo_ip is not None:
            pulumi.set(__self__, "geo_ip", geo_ip)
        if http_service is not None:
            pulumi.set(__self__, "http_service", http_service)
        if keep_empty_zone is not None:
            pulumi.set(__self__, "keep_empty_zone", keep_empty_zone)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if priority is not None:
            pulumi.set(__self__, "priority", priority)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if ttl is not None:
            pulumi.set(__self__, "ttl", ttl)
        if views is not None:
            pulumi.set(__self__, "views", views)
        if weighteds is not None:
            pulumi.set(__self__, "weighteds", weighteds)

    @property
    @pulumi.getter
    def data(self) -> pulumi.Input[str]:
        """
        The data of the record
        """
        return pulumi.get(self, "data")

    @data.setter
    def data(self, value: pulumi.Input[str]):
        pulumi.set(self, "data", value)

    @property
    @pulumi.getter(name="dnsZone")
    def dns_zone(self) -> pulumi.Input[str]:
        """
        The zone you want to add the record in
        """
        return pulumi.get(self, "dns_zone")

    @dns_zone.setter
    def dns_zone(self, value: pulumi.Input[str]):
        pulumi.set(self, "dns_zone", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The type of the record
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="geoIp")
    def geo_ip(self) -> Optional[pulumi.Input['DomainRecordGeoIpArgs']]:
        """
        Return record based on client localisation
        """
        return pulumi.get(self, "geo_ip")

    @geo_ip.setter
    def geo_ip(self, value: Optional[pulumi.Input['DomainRecordGeoIpArgs']]):
        pulumi.set(self, "geo_ip", value)

    @property
    @pulumi.getter(name="httpService")
    def http_service(self) -> Optional[pulumi.Input['DomainRecordHttpServiceArgs']]:
        """
        Return record based on client localisation
        """
        return pulumi.get(self, "http_service")

    @http_service.setter
    def http_service(self, value: Optional[pulumi.Input['DomainRecordHttpServiceArgs']]):
        pulumi.set(self, "http_service", value)

    @property
    @pulumi.getter(name="keepEmptyZone")
    def keep_empty_zone(self) -> Optional[pulumi.Input[bool]]:
        """
        When destroy a resource record, if a zone have only NS, delete the zone
        """
        return pulumi.get(self, "keep_empty_zone")

    @keep_empty_zone.setter
    def keep_empty_zone(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "keep_empty_zone", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the record
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def priority(self) -> Optional[pulumi.Input[int]]:
        """
        The priority of the record
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The project_id you want to attach the resource to
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def ttl(self) -> Optional[pulumi.Input[int]]:
        """
        The ttl of the record
        """
        return pulumi.get(self, "ttl")

    @ttl.setter
    def ttl(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ttl", value)

    @property
    @pulumi.getter
    def views(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DomainRecordViewArgs']]]]:
        """
        Return record based on client subnet
        """
        return pulumi.get(self, "views")

    @views.setter
    def views(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DomainRecordViewArgs']]]]):
        pulumi.set(self, "views", value)

    @property
    @pulumi.getter
    def weighteds(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DomainRecordWeightedArgs']]]]:
        """
        Return record based on weight
        """
        return pulumi.get(self, "weighteds")

    @weighteds.setter
    def weighteds(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DomainRecordWeightedArgs']]]]):
        pulumi.set(self, "weighteds", value)


@pulumi.input_type
class _DomainRecordState:
    def __init__(__self__, *,
                 data: Optional[pulumi.Input[str]] = None,
                 dns_zone: Optional[pulumi.Input[str]] = None,
                 geo_ip: Optional[pulumi.Input['DomainRecordGeoIpArgs']] = None,
                 http_service: Optional[pulumi.Input['DomainRecordHttpServiceArgs']] = None,
                 keep_empty_zone: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 root_zone: Optional[pulumi.Input[bool]] = None,
                 ttl: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 views: Optional[pulumi.Input[Sequence[pulumi.Input['DomainRecordViewArgs']]]] = None,
                 weighteds: Optional[pulumi.Input[Sequence[pulumi.Input['DomainRecordWeightedArgs']]]] = None):
        """
        Input properties used for looking up and filtering DomainRecord resources.
        :param pulumi.Input[str] data: The data of the record
        :param pulumi.Input[str] dns_zone: The zone you want to add the record in
        :param pulumi.Input['DomainRecordGeoIpArgs'] geo_ip: Return record based on client localisation
        :param pulumi.Input['DomainRecordHttpServiceArgs'] http_service: Return record based on client localisation
        :param pulumi.Input[bool] keep_empty_zone: When destroy a resource record, if a zone have only NS, delete the zone
        :param pulumi.Input[str] name: The name of the record
        :param pulumi.Input[int] priority: The priority of the record
        :param pulumi.Input[str] project_id: The project_id you want to attach the resource to
        :param pulumi.Input[bool] root_zone: Does the DNS zone is the root zone or not
        :param pulumi.Input[int] ttl: The ttl of the record
        :param pulumi.Input[str] type: The type of the record
        :param pulumi.Input[Sequence[pulumi.Input['DomainRecordViewArgs']]] views: Return record based on client subnet
        :param pulumi.Input[Sequence[pulumi.Input['DomainRecordWeightedArgs']]] weighteds: Return record based on weight
        """
        if data is not None:
            pulumi.set(__self__, "data", data)
        if dns_zone is not None:
            pulumi.set(__self__, "dns_zone", dns_zone)
        if geo_ip is not None:
            pulumi.set(__self__, "geo_ip", geo_ip)
        if http_service is not None:
            pulumi.set(__self__, "http_service", http_service)
        if keep_empty_zone is not None:
            pulumi.set(__self__, "keep_empty_zone", keep_empty_zone)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if priority is not None:
            pulumi.set(__self__, "priority", priority)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if root_zone is not None:
            pulumi.set(__self__, "root_zone", root_zone)
        if ttl is not None:
            pulumi.set(__self__, "ttl", ttl)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if views is not None:
            pulumi.set(__self__, "views", views)
        if weighteds is not None:
            pulumi.set(__self__, "weighteds", weighteds)

    @property
    @pulumi.getter
    def data(self) -> Optional[pulumi.Input[str]]:
        """
        The data of the record
        """
        return pulumi.get(self, "data")

    @data.setter
    def data(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data", value)

    @property
    @pulumi.getter(name="dnsZone")
    def dns_zone(self) -> Optional[pulumi.Input[str]]:
        """
        The zone you want to add the record in
        """
        return pulumi.get(self, "dns_zone")

    @dns_zone.setter
    def dns_zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dns_zone", value)

    @property
    @pulumi.getter(name="geoIp")
    def geo_ip(self) -> Optional[pulumi.Input['DomainRecordGeoIpArgs']]:
        """
        Return record based on client localisation
        """
        return pulumi.get(self, "geo_ip")

    @geo_ip.setter
    def geo_ip(self, value: Optional[pulumi.Input['DomainRecordGeoIpArgs']]):
        pulumi.set(self, "geo_ip", value)

    @property
    @pulumi.getter(name="httpService")
    def http_service(self) -> Optional[pulumi.Input['DomainRecordHttpServiceArgs']]:
        """
        Return record based on client localisation
        """
        return pulumi.get(self, "http_service")

    @http_service.setter
    def http_service(self, value: Optional[pulumi.Input['DomainRecordHttpServiceArgs']]):
        pulumi.set(self, "http_service", value)

    @property
    @pulumi.getter(name="keepEmptyZone")
    def keep_empty_zone(self) -> Optional[pulumi.Input[bool]]:
        """
        When destroy a resource record, if a zone have only NS, delete the zone
        """
        return pulumi.get(self, "keep_empty_zone")

    @keep_empty_zone.setter
    def keep_empty_zone(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "keep_empty_zone", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the record
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def priority(self) -> Optional[pulumi.Input[int]]:
        """
        The priority of the record
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The project_id you want to attach the resource to
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="rootZone")
    def root_zone(self) -> Optional[pulumi.Input[bool]]:
        """
        Does the DNS zone is the root zone or not
        """
        return pulumi.get(self, "root_zone")

    @root_zone.setter
    def root_zone(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "root_zone", value)

    @property
    @pulumi.getter
    def ttl(self) -> Optional[pulumi.Input[int]]:
        """
        The ttl of the record
        """
        return pulumi.get(self, "ttl")

    @ttl.setter
    def ttl(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ttl", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the record
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def views(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DomainRecordViewArgs']]]]:
        """
        Return record based on client subnet
        """
        return pulumi.get(self, "views")

    @views.setter
    def views(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DomainRecordViewArgs']]]]):
        pulumi.set(self, "views", value)

    @property
    @pulumi.getter
    def weighteds(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DomainRecordWeightedArgs']]]]:
        """
        Return record based on weight
        """
        return pulumi.get(self, "weighteds")

    @weighteds.setter
    def weighteds(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DomainRecordWeightedArgs']]]]):
        pulumi.set(self, "weighteds", value)


class DomainRecord(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data: Optional[pulumi.Input[str]] = None,
                 dns_zone: Optional[pulumi.Input[str]] = None,
                 geo_ip: Optional[pulumi.Input[pulumi.InputType['DomainRecordGeoIpArgs']]] = None,
                 http_service: Optional[pulumi.Input[pulumi.InputType['DomainRecordHttpServiceArgs']]] = None,
                 keep_empty_zone: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 ttl: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 views: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DomainRecordViewArgs']]]]] = None,
                 weighteds: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DomainRecordWeightedArgs']]]]] = None,
                 __props__=None):
        """
        Create a DomainRecord resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] data: The data of the record
        :param pulumi.Input[str] dns_zone: The zone you want to add the record in
        :param pulumi.Input[pulumi.InputType['DomainRecordGeoIpArgs']] geo_ip: Return record based on client localisation
        :param pulumi.Input[pulumi.InputType['DomainRecordHttpServiceArgs']] http_service: Return record based on client localisation
        :param pulumi.Input[bool] keep_empty_zone: When destroy a resource record, if a zone have only NS, delete the zone
        :param pulumi.Input[str] name: The name of the record
        :param pulumi.Input[int] priority: The priority of the record
        :param pulumi.Input[str] project_id: The project_id you want to attach the resource to
        :param pulumi.Input[int] ttl: The ttl of the record
        :param pulumi.Input[str] type: The type of the record
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DomainRecordViewArgs']]]] views: Return record based on client subnet
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DomainRecordWeightedArgs']]]] weighteds: Return record based on weight
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DomainRecordArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a DomainRecord resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param DomainRecordArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DomainRecordArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data: Optional[pulumi.Input[str]] = None,
                 dns_zone: Optional[pulumi.Input[str]] = None,
                 geo_ip: Optional[pulumi.Input[pulumi.InputType['DomainRecordGeoIpArgs']]] = None,
                 http_service: Optional[pulumi.Input[pulumi.InputType['DomainRecordHttpServiceArgs']]] = None,
                 keep_empty_zone: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 ttl: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 views: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DomainRecordViewArgs']]]]] = None,
                 weighteds: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DomainRecordWeightedArgs']]]]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DomainRecordArgs.__new__(DomainRecordArgs)

            if data is None and not opts.urn:
                raise TypeError("Missing required property 'data'")
            __props__.__dict__["data"] = data
            if dns_zone is None and not opts.urn:
                raise TypeError("Missing required property 'dns_zone'")
            __props__.__dict__["dns_zone"] = dns_zone
            __props__.__dict__["geo_ip"] = geo_ip
            __props__.__dict__["http_service"] = http_service
            __props__.__dict__["keep_empty_zone"] = keep_empty_zone
            __props__.__dict__["name"] = name
            __props__.__dict__["priority"] = priority
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["ttl"] = ttl
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["views"] = views
            __props__.__dict__["weighteds"] = weighteds
            __props__.__dict__["root_zone"] = None
        super(DomainRecord, __self__).__init__(
            'scaleway:index/domainRecord:DomainRecord',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            data: Optional[pulumi.Input[str]] = None,
            dns_zone: Optional[pulumi.Input[str]] = None,
            geo_ip: Optional[pulumi.Input[pulumi.InputType['DomainRecordGeoIpArgs']]] = None,
            http_service: Optional[pulumi.Input[pulumi.InputType['DomainRecordHttpServiceArgs']]] = None,
            keep_empty_zone: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            priority: Optional[pulumi.Input[int]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            root_zone: Optional[pulumi.Input[bool]] = None,
            ttl: Optional[pulumi.Input[int]] = None,
            type: Optional[pulumi.Input[str]] = None,
            views: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DomainRecordViewArgs']]]]] = None,
            weighteds: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DomainRecordWeightedArgs']]]]] = None) -> 'DomainRecord':
        """
        Get an existing DomainRecord resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] data: The data of the record
        :param pulumi.Input[str] dns_zone: The zone you want to add the record in
        :param pulumi.Input[pulumi.InputType['DomainRecordGeoIpArgs']] geo_ip: Return record based on client localisation
        :param pulumi.Input[pulumi.InputType['DomainRecordHttpServiceArgs']] http_service: Return record based on client localisation
        :param pulumi.Input[bool] keep_empty_zone: When destroy a resource record, if a zone have only NS, delete the zone
        :param pulumi.Input[str] name: The name of the record
        :param pulumi.Input[int] priority: The priority of the record
        :param pulumi.Input[str] project_id: The project_id you want to attach the resource to
        :param pulumi.Input[bool] root_zone: Does the DNS zone is the root zone or not
        :param pulumi.Input[int] ttl: The ttl of the record
        :param pulumi.Input[str] type: The type of the record
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DomainRecordViewArgs']]]] views: Return record based on client subnet
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DomainRecordWeightedArgs']]]] weighteds: Return record based on weight
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DomainRecordState.__new__(_DomainRecordState)

        __props__.__dict__["data"] = data
        __props__.__dict__["dns_zone"] = dns_zone
        __props__.__dict__["geo_ip"] = geo_ip
        __props__.__dict__["http_service"] = http_service
        __props__.__dict__["keep_empty_zone"] = keep_empty_zone
        __props__.__dict__["name"] = name
        __props__.__dict__["priority"] = priority
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["root_zone"] = root_zone
        __props__.__dict__["ttl"] = ttl
        __props__.__dict__["type"] = type
        __props__.__dict__["views"] = views
        __props__.__dict__["weighteds"] = weighteds
        return DomainRecord(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def data(self) -> pulumi.Output[str]:
        """
        The data of the record
        """
        return pulumi.get(self, "data")

    @property
    @pulumi.getter(name="dnsZone")
    def dns_zone(self) -> pulumi.Output[str]:
        """
        The zone you want to add the record in
        """
        return pulumi.get(self, "dns_zone")

    @property
    @pulumi.getter(name="geoIp")
    def geo_ip(self) -> pulumi.Output[Optional['outputs.DomainRecordGeoIp']]:
        """
        Return record based on client localisation
        """
        return pulumi.get(self, "geo_ip")

    @property
    @pulumi.getter(name="httpService")
    def http_service(self) -> pulumi.Output[Optional['outputs.DomainRecordHttpService']]:
        """
        Return record based on client localisation
        """
        return pulumi.get(self, "http_service")

    @property
    @pulumi.getter(name="keepEmptyZone")
    def keep_empty_zone(self) -> pulumi.Output[Optional[bool]]:
        """
        When destroy a resource record, if a zone have only NS, delete the zone
        """
        return pulumi.get(self, "keep_empty_zone")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the record
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def priority(self) -> pulumi.Output[int]:
        """
        The priority of the record
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        The project_id you want to attach the resource to
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="rootZone")
    def root_zone(self) -> pulumi.Output[bool]:
        """
        Does the DNS zone is the root zone or not
        """
        return pulumi.get(self, "root_zone")

    @property
    @pulumi.getter
    def ttl(self) -> pulumi.Output[Optional[int]]:
        """
        The ttl of the record
        """
        return pulumi.get(self, "ttl")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the record
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def views(self) -> pulumi.Output[Optional[Sequence['outputs.DomainRecordView']]]:
        """
        Return record based on client subnet
        """
        return pulumi.get(self, "views")

    @property
    @pulumi.getter
    def weighteds(self) -> pulumi.Output[Optional[Sequence['outputs.DomainRecordWeighted']]]:
        """
        Return record based on weight
        """
        return pulumi.get(self, "weighteds")

