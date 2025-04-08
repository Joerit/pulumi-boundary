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

__all__ = ['HostCatalogArgs', 'HostCatalog']

@pulumi.input_type
class HostCatalogArgs:
    def __init__(__self__, *,
                 scope_id: pulumi.Input[str],
                 type: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a HostCatalog resource.
        :param pulumi.Input[str] scope_id: The scope ID in which the resource is created.
        :param pulumi.Input[str] type: The host catalog type. Only `static` is supported.
        :param pulumi.Input[str] description: The host catalog description.
        :param pulumi.Input[str] name: The host catalog name. Defaults to the resource name.
        """
        pulumi.set(__self__, "scope_id", scope_id)
        pulumi.set(__self__, "type", type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="scopeId")
    def scope_id(self) -> pulumi.Input[str]:
        """
        The scope ID in which the resource is created.
        """
        return pulumi.get(self, "scope_id")

    @scope_id.setter
    def scope_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "scope_id", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The host catalog type. Only `static` is supported.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The host catalog description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The host catalog name. Defaults to the resource name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _HostCatalogState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 scope_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering HostCatalog resources.
        :param pulumi.Input[str] description: The host catalog description.
        :param pulumi.Input[str] name: The host catalog name. Defaults to the resource name.
        :param pulumi.Input[str] scope_id: The scope ID in which the resource is created.
        :param pulumi.Input[str] type: The host catalog type. Only `static` is supported.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if scope_id is not None:
            pulumi.set(__self__, "scope_id", scope_id)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The host catalog description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The host catalog name. Defaults to the resource name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="scopeId")
    def scope_id(self) -> Optional[pulumi.Input[str]]:
        """
        The scope ID in which the resource is created.
        """
        return pulumi.get(self, "scope_id")

    @scope_id.setter
    def scope_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scope_id", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The host catalog type. Only `static` is supported.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class HostCatalog(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 scope_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a HostCatalog resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The host catalog description.
        :param pulumi.Input[str] name: The host catalog name. Defaults to the resource name.
        :param pulumi.Input[str] scope_id: The scope ID in which the resource is created.
        :param pulumi.Input[str] type: The host catalog type. Only `static` is supported.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: HostCatalogArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a HostCatalog resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param HostCatalogArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(HostCatalogArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 scope_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = HostCatalogArgs.__new__(HostCatalogArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            if scope_id is None and not opts.urn:
                raise TypeError("Missing required property 'scope_id'")
            __props__.__dict__["scope_id"] = scope_id
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
        super(HostCatalog, __self__).__init__(
            'boundary:index/hostCatalog:HostCatalog',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            scope_id: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'HostCatalog':
        """
        Get an existing HostCatalog resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The host catalog description.
        :param pulumi.Input[str] name: The host catalog name. Defaults to the resource name.
        :param pulumi.Input[str] scope_id: The scope ID in which the resource is created.
        :param pulumi.Input[str] type: The host catalog type. Only `static` is supported.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _HostCatalogState.__new__(_HostCatalogState)

        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["scope_id"] = scope_id
        __props__.__dict__["type"] = type
        return HostCatalog(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The host catalog description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The host catalog name. Defaults to the resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="scopeId")
    def scope_id(self) -> pulumi.Output[str]:
        """
        The scope ID in which the resource is created.
        """
        return pulumi.get(self, "scope_id")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The host catalog type. Only `static` is supported.
        """
        return pulumi.get(self, "type")

