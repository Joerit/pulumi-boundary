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

__all__ = ['ManagedGroupLdapArgs', 'ManagedGroupLdap']

@pulumi.input_type
class ManagedGroupLdapArgs:
    def __init__(__self__, *,
                 auth_method_id: pulumi.Input[str],
                 group_names: pulumi.Input[Sequence[pulumi.Input[str]]],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ManagedGroupLdap resource.
        :param pulumi.Input[str] auth_method_id: The resource ID for the auth method.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] group_names: The list of groups that make up the managed group.
        :param pulumi.Input[str] description: The managed group description.
        :param pulumi.Input[str] name: The managed group name. Defaults to the resource name.
        """
        pulumi.set(__self__, "auth_method_id", auth_method_id)
        pulumi.set(__self__, "group_names", group_names)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="authMethodId")
    def auth_method_id(self) -> pulumi.Input[str]:
        """
        The resource ID for the auth method.
        """
        return pulumi.get(self, "auth_method_id")

    @auth_method_id.setter
    def auth_method_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "auth_method_id", value)

    @property
    @pulumi.getter(name="groupNames")
    def group_names(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        The list of groups that make up the managed group.
        """
        return pulumi.get(self, "group_names")

    @group_names.setter
    def group_names(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "group_names", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The managed group description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The managed group name. Defaults to the resource name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _ManagedGroupLdapState:
    def __init__(__self__, *,
                 auth_method_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 group_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ManagedGroupLdap resources.
        :param pulumi.Input[str] auth_method_id: The resource ID for the auth method.
        :param pulumi.Input[str] description: The managed group description.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] group_names: The list of groups that make up the managed group.
        :param pulumi.Input[str] name: The managed group name. Defaults to the resource name.
        """
        if auth_method_id is not None:
            pulumi.set(__self__, "auth_method_id", auth_method_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if group_names is not None:
            pulumi.set(__self__, "group_names", group_names)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="authMethodId")
    def auth_method_id(self) -> Optional[pulumi.Input[str]]:
        """
        The resource ID for the auth method.
        """
        return pulumi.get(self, "auth_method_id")

    @auth_method_id.setter
    def auth_method_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_method_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The managed group description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="groupNames")
    def group_names(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The list of groups that make up the managed group.
        """
        return pulumi.get(self, "group_names")

    @group_names.setter
    def group_names(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "group_names", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The managed group name. Defaults to the resource name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class ManagedGroupLdap(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_method_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 group_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a ManagedGroupLdap resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auth_method_id: The resource ID for the auth method.
        :param pulumi.Input[str] description: The managed group description.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] group_names: The list of groups that make up the managed group.
        :param pulumi.Input[str] name: The managed group name. Defaults to the resource name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ManagedGroupLdapArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a ManagedGroupLdap resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ManagedGroupLdapArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ManagedGroupLdapArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_method_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 group_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ManagedGroupLdapArgs.__new__(ManagedGroupLdapArgs)

            if auth_method_id is None and not opts.urn:
                raise TypeError("Missing required property 'auth_method_id'")
            __props__.__dict__["auth_method_id"] = auth_method_id
            __props__.__dict__["description"] = description
            if group_names is None and not opts.urn:
                raise TypeError("Missing required property 'group_names'")
            __props__.__dict__["group_names"] = group_names
            __props__.__dict__["name"] = name
        super(ManagedGroupLdap, __self__).__init__(
            'boundary:index/managedGroupLdap:ManagedGroupLdap',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auth_method_id: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            group_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'ManagedGroupLdap':
        """
        Get an existing ManagedGroupLdap resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auth_method_id: The resource ID for the auth method.
        :param pulumi.Input[str] description: The managed group description.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] group_names: The list of groups that make up the managed group.
        :param pulumi.Input[str] name: The managed group name. Defaults to the resource name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ManagedGroupLdapState.__new__(_ManagedGroupLdapState)

        __props__.__dict__["auth_method_id"] = auth_method_id
        __props__.__dict__["description"] = description
        __props__.__dict__["group_names"] = group_names
        __props__.__dict__["name"] = name
        return ManagedGroupLdap(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="authMethodId")
    def auth_method_id(self) -> pulumi.Output[str]:
        """
        The resource ID for the auth method.
        """
        return pulumi.get(self, "auth_method_id")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The managed group description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="groupNames")
    def group_names(self) -> pulumi.Output[Sequence[str]]:
        """
        The list of groups that make up the managed group.
        """
        return pulumi.get(self, "group_names")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The managed group name. Defaults to the resource name.
        """
        return pulumi.get(self, "name")

