# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities

__all__ = ['HostStaticArgs', 'HostStatic']

@pulumi.input_type
class HostStaticArgs:
    def __init__(__self__, *,
                 host_catalog_id: pulumi.Input[str],
                 address: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a HostStatic resource.
        :param pulumi.Input[str] address: The static address of the host resource as `<IP>` (note: port assignment occurs in the target resource definition, do
               not add :port here) or a domain name.
        :param pulumi.Input[str] description: The host description.
        :param pulumi.Input[str] name: The host name. Defaults to the resource name.
        :param pulumi.Input[str] type: The type of host
        """
        pulumi.set(__self__, "host_catalog_id", host_catalog_id)
        if address is not None:
            pulumi.set(__self__, "address", address)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="hostCatalogId")
    def host_catalog_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "host_catalog_id")

    @host_catalog_id.setter
    def host_catalog_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "host_catalog_id", value)

    @property
    @pulumi.getter
    def address(self) -> Optional[pulumi.Input[str]]:
        """
        The static address of the host resource as `<IP>` (note: port assignment occurs in the target resource definition, do
        not add :port here) or a domain name.
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The host description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The host name. Defaults to the resource name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of host
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class _HostStaticState:
    def __init__(__self__, *,
                 address: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 host_catalog_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering HostStatic resources.
        :param pulumi.Input[str] address: The static address of the host resource as `<IP>` (note: port assignment occurs in the target resource definition, do
               not add :port here) or a domain name.
        :param pulumi.Input[str] description: The host description.
        :param pulumi.Input[str] name: The host name. Defaults to the resource name.
        :param pulumi.Input[str] type: The type of host
        """
        if address is not None:
            pulumi.set(__self__, "address", address)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if host_catalog_id is not None:
            pulumi.set(__self__, "host_catalog_id", host_catalog_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def address(self) -> Optional[pulumi.Input[str]]:
        """
        The static address of the host resource as `<IP>` (note: port assignment occurs in the target resource definition, do
        not add :port here) or a domain name.
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The host description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="hostCatalogId")
    def host_catalog_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "host_catalog_id")

    @host_catalog_id.setter
    def host_catalog_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host_catalog_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The host name. Defaults to the resource name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of host
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class HostStatic(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 host_catalog_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a HostStatic resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: The static address of the host resource as `<IP>` (note: port assignment occurs in the target resource definition, do
               not add :port here) or a domain name.
        :param pulumi.Input[str] description: The host description.
        :param pulumi.Input[str] name: The host name. Defaults to the resource name.
        :param pulumi.Input[str] type: The type of host
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: HostStaticArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a HostStatic resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param HostStaticArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(HostStaticArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 host_catalog_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = HostStaticArgs.__new__(HostStaticArgs)

            __props__.__dict__["address"] = address
            __props__.__dict__["description"] = description
            if host_catalog_id is None and not opts.urn:
                raise TypeError("Missing required property 'host_catalog_id'")
            __props__.__dict__["host_catalog_id"] = host_catalog_id
            __props__.__dict__["name"] = name
            __props__.__dict__["type"] = type
        super(HostStatic, __self__).__init__(
            'boundary:index/hostStatic:HostStatic',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            address: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            host_catalog_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'HostStatic':
        """
        Get an existing HostStatic resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: The static address of the host resource as `<IP>` (note: port assignment occurs in the target resource definition, do
               not add :port here) or a domain name.
        :param pulumi.Input[str] description: The host description.
        :param pulumi.Input[str] name: The host name. Defaults to the resource name.
        :param pulumi.Input[str] type: The type of host
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _HostStaticState.__new__(_HostStaticState)

        __props__.__dict__["address"] = address
        __props__.__dict__["description"] = description
        __props__.__dict__["host_catalog_id"] = host_catalog_id
        __props__.__dict__["name"] = name
        __props__.__dict__["type"] = type
        return HostStatic(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Output[Optional[str]]:
        """
        The static address of the host resource as `<IP>` (note: port assignment occurs in the target resource definition, do
        not add :port here) or a domain name.
        """
        return pulumi.get(self, "address")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The host description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="hostCatalogId")
    def host_catalog_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "host_catalog_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The host name. Defaults to the resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[str]]:
        """
        The type of host
        """
        return pulumi.get(self, "type")

