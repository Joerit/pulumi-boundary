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

__all__ = ['PolicyStorageArgs', 'PolicyStorage']

@pulumi.input_type
class PolicyStorageArgs:
    def __init__(__self__, *,
                 scope_id: pulumi.Input[str],
                 delete_after_days: Optional[pulumi.Input[int]] = None,
                 delete_after_overridable: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 retain_for_days: Optional[pulumi.Input[int]] = None,
                 retain_for_overridable: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a PolicyStorage resource.
        :param pulumi.Input[str] scope_id: The scope for this policy.
        :param pulumi.Input[int] delete_after_days: The number of days after which a session recording will be automatically deleted. Defaults to 0: never automatically
               delete. However, delete_after_days and retain_for_days cannot both be 0.
        :param pulumi.Input[bool] delete_after_overridable: Whether or not the associated delete_after_days value can be overridden by org scopes. Note: if the associated
               delete_after_days value is 0, overridable is ignored
        :param pulumi.Input[str] description: The policy description.
        :param pulumi.Input[str] name: The policy name. Defaults to the resource name.
        :param pulumi.Input[int] retain_for_days: The number of days a session recording is required to be stored. Defaults to 0: allow deletions at any time. However,
               retain_for_days and delete_after_days cannot both be 0.
        :param pulumi.Input[bool] retain_for_overridable: Whether or not the associated retain_for_days value can be overridden by org scopes. Note: if the associated
               retain_for_days value is 0, overridable is ignored.
        """
        pulumi.set(__self__, "scope_id", scope_id)
        if delete_after_days is not None:
            pulumi.set(__self__, "delete_after_days", delete_after_days)
        if delete_after_overridable is not None:
            pulumi.set(__self__, "delete_after_overridable", delete_after_overridable)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if retain_for_days is not None:
            pulumi.set(__self__, "retain_for_days", retain_for_days)
        if retain_for_overridable is not None:
            pulumi.set(__self__, "retain_for_overridable", retain_for_overridable)

    @property
    @pulumi.getter(name="scopeId")
    def scope_id(self) -> pulumi.Input[str]:
        """
        The scope for this policy.
        """
        return pulumi.get(self, "scope_id")

    @scope_id.setter
    def scope_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "scope_id", value)

    @property
    @pulumi.getter(name="deleteAfterDays")
    def delete_after_days(self) -> Optional[pulumi.Input[int]]:
        """
        The number of days after which a session recording will be automatically deleted. Defaults to 0: never automatically
        delete. However, delete_after_days and retain_for_days cannot both be 0.
        """
        return pulumi.get(self, "delete_after_days")

    @delete_after_days.setter
    def delete_after_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "delete_after_days", value)

    @property
    @pulumi.getter(name="deleteAfterOverridable")
    def delete_after_overridable(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not the associated delete_after_days value can be overridden by org scopes. Note: if the associated
        delete_after_days value is 0, overridable is ignored
        """
        return pulumi.get(self, "delete_after_overridable")

    @delete_after_overridable.setter
    def delete_after_overridable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "delete_after_overridable", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The policy description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The policy name. Defaults to the resource name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="retainForDays")
    def retain_for_days(self) -> Optional[pulumi.Input[int]]:
        """
        The number of days a session recording is required to be stored. Defaults to 0: allow deletions at any time. However,
        retain_for_days and delete_after_days cannot both be 0.
        """
        return pulumi.get(self, "retain_for_days")

    @retain_for_days.setter
    def retain_for_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retain_for_days", value)

    @property
    @pulumi.getter(name="retainForOverridable")
    def retain_for_overridable(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not the associated retain_for_days value can be overridden by org scopes. Note: if the associated
        retain_for_days value is 0, overridable is ignored.
        """
        return pulumi.get(self, "retain_for_overridable")

    @retain_for_overridable.setter
    def retain_for_overridable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "retain_for_overridable", value)


@pulumi.input_type
class _PolicyStorageState:
    def __init__(__self__, *,
                 delete_after_days: Optional[pulumi.Input[int]] = None,
                 delete_after_overridable: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 retain_for_days: Optional[pulumi.Input[int]] = None,
                 retain_for_overridable: Optional[pulumi.Input[bool]] = None,
                 scope_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering PolicyStorage resources.
        :param pulumi.Input[int] delete_after_days: The number of days after which a session recording will be automatically deleted. Defaults to 0: never automatically
               delete. However, delete_after_days and retain_for_days cannot both be 0.
        :param pulumi.Input[bool] delete_after_overridable: Whether or not the associated delete_after_days value can be overridden by org scopes. Note: if the associated
               delete_after_days value is 0, overridable is ignored
        :param pulumi.Input[str] description: The policy description.
        :param pulumi.Input[str] name: The policy name. Defaults to the resource name.
        :param pulumi.Input[int] retain_for_days: The number of days a session recording is required to be stored. Defaults to 0: allow deletions at any time. However,
               retain_for_days and delete_after_days cannot both be 0.
        :param pulumi.Input[bool] retain_for_overridable: Whether or not the associated retain_for_days value can be overridden by org scopes. Note: if the associated
               retain_for_days value is 0, overridable is ignored.
        :param pulumi.Input[str] scope_id: The scope for this policy.
        """
        if delete_after_days is not None:
            pulumi.set(__self__, "delete_after_days", delete_after_days)
        if delete_after_overridable is not None:
            pulumi.set(__self__, "delete_after_overridable", delete_after_overridable)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if retain_for_days is not None:
            pulumi.set(__self__, "retain_for_days", retain_for_days)
        if retain_for_overridable is not None:
            pulumi.set(__self__, "retain_for_overridable", retain_for_overridable)
        if scope_id is not None:
            pulumi.set(__self__, "scope_id", scope_id)

    @property
    @pulumi.getter(name="deleteAfterDays")
    def delete_after_days(self) -> Optional[pulumi.Input[int]]:
        """
        The number of days after which a session recording will be automatically deleted. Defaults to 0: never automatically
        delete. However, delete_after_days and retain_for_days cannot both be 0.
        """
        return pulumi.get(self, "delete_after_days")

    @delete_after_days.setter
    def delete_after_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "delete_after_days", value)

    @property
    @pulumi.getter(name="deleteAfterOverridable")
    def delete_after_overridable(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not the associated delete_after_days value can be overridden by org scopes. Note: if the associated
        delete_after_days value is 0, overridable is ignored
        """
        return pulumi.get(self, "delete_after_overridable")

    @delete_after_overridable.setter
    def delete_after_overridable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "delete_after_overridable", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The policy description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The policy name. Defaults to the resource name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="retainForDays")
    def retain_for_days(self) -> Optional[pulumi.Input[int]]:
        """
        The number of days a session recording is required to be stored. Defaults to 0: allow deletions at any time. However,
        retain_for_days and delete_after_days cannot both be 0.
        """
        return pulumi.get(self, "retain_for_days")

    @retain_for_days.setter
    def retain_for_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retain_for_days", value)

    @property
    @pulumi.getter(name="retainForOverridable")
    def retain_for_overridable(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not the associated retain_for_days value can be overridden by org scopes. Note: if the associated
        retain_for_days value is 0, overridable is ignored.
        """
        return pulumi.get(self, "retain_for_overridable")

    @retain_for_overridable.setter
    def retain_for_overridable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "retain_for_overridable", value)

    @property
    @pulumi.getter(name="scopeId")
    def scope_id(self) -> Optional[pulumi.Input[str]]:
        """
        The scope for this policy.
        """
        return pulumi.get(self, "scope_id")

    @scope_id.setter
    def scope_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scope_id", value)


class PolicyStorage(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 delete_after_days: Optional[pulumi.Input[int]] = None,
                 delete_after_overridable: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 retain_for_days: Optional[pulumi.Input[int]] = None,
                 retain_for_overridable: Optional[pulumi.Input[bool]] = None,
                 scope_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a PolicyStorage resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] delete_after_days: The number of days after which a session recording will be automatically deleted. Defaults to 0: never automatically
               delete. However, delete_after_days and retain_for_days cannot both be 0.
        :param pulumi.Input[bool] delete_after_overridable: Whether or not the associated delete_after_days value can be overridden by org scopes. Note: if the associated
               delete_after_days value is 0, overridable is ignored
        :param pulumi.Input[str] description: The policy description.
        :param pulumi.Input[str] name: The policy name. Defaults to the resource name.
        :param pulumi.Input[int] retain_for_days: The number of days a session recording is required to be stored. Defaults to 0: allow deletions at any time. However,
               retain_for_days and delete_after_days cannot both be 0.
        :param pulumi.Input[bool] retain_for_overridable: Whether or not the associated retain_for_days value can be overridden by org scopes. Note: if the associated
               retain_for_days value is 0, overridable is ignored.
        :param pulumi.Input[str] scope_id: The scope for this policy.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PolicyStorageArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a PolicyStorage resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param PolicyStorageArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PolicyStorageArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 delete_after_days: Optional[pulumi.Input[int]] = None,
                 delete_after_overridable: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 retain_for_days: Optional[pulumi.Input[int]] = None,
                 retain_for_overridable: Optional[pulumi.Input[bool]] = None,
                 scope_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PolicyStorageArgs.__new__(PolicyStorageArgs)

            __props__.__dict__["delete_after_days"] = delete_after_days
            __props__.__dict__["delete_after_overridable"] = delete_after_overridable
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            __props__.__dict__["retain_for_days"] = retain_for_days
            __props__.__dict__["retain_for_overridable"] = retain_for_overridable
            if scope_id is None and not opts.urn:
                raise TypeError("Missing required property 'scope_id'")
            __props__.__dict__["scope_id"] = scope_id
        super(PolicyStorage, __self__).__init__(
            'boundary:index/policyStorage:PolicyStorage',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            delete_after_days: Optional[pulumi.Input[int]] = None,
            delete_after_overridable: Optional[pulumi.Input[bool]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            retain_for_days: Optional[pulumi.Input[int]] = None,
            retain_for_overridable: Optional[pulumi.Input[bool]] = None,
            scope_id: Optional[pulumi.Input[str]] = None) -> 'PolicyStorage':
        """
        Get an existing PolicyStorage resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] delete_after_days: The number of days after which a session recording will be automatically deleted. Defaults to 0: never automatically
               delete. However, delete_after_days and retain_for_days cannot both be 0.
        :param pulumi.Input[bool] delete_after_overridable: Whether or not the associated delete_after_days value can be overridden by org scopes. Note: if the associated
               delete_after_days value is 0, overridable is ignored
        :param pulumi.Input[str] description: The policy description.
        :param pulumi.Input[str] name: The policy name. Defaults to the resource name.
        :param pulumi.Input[int] retain_for_days: The number of days a session recording is required to be stored. Defaults to 0: allow deletions at any time. However,
               retain_for_days and delete_after_days cannot both be 0.
        :param pulumi.Input[bool] retain_for_overridable: Whether or not the associated retain_for_days value can be overridden by org scopes. Note: if the associated
               retain_for_days value is 0, overridable is ignored.
        :param pulumi.Input[str] scope_id: The scope for this policy.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PolicyStorageState.__new__(_PolicyStorageState)

        __props__.__dict__["delete_after_days"] = delete_after_days
        __props__.__dict__["delete_after_overridable"] = delete_after_overridable
        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["retain_for_days"] = retain_for_days
        __props__.__dict__["retain_for_overridable"] = retain_for_overridable
        __props__.__dict__["scope_id"] = scope_id
        return PolicyStorage(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="deleteAfterDays")
    def delete_after_days(self) -> pulumi.Output[Optional[int]]:
        """
        The number of days after which a session recording will be automatically deleted. Defaults to 0: never automatically
        delete. However, delete_after_days and retain_for_days cannot both be 0.
        """
        return pulumi.get(self, "delete_after_days")

    @property
    @pulumi.getter(name="deleteAfterOverridable")
    def delete_after_overridable(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether or not the associated delete_after_days value can be overridden by org scopes. Note: if the associated
        delete_after_days value is 0, overridable is ignored
        """
        return pulumi.get(self, "delete_after_overridable")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The policy description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The policy name. Defaults to the resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="retainForDays")
    def retain_for_days(self) -> pulumi.Output[Optional[int]]:
        """
        The number of days a session recording is required to be stored. Defaults to 0: allow deletions at any time. However,
        retain_for_days and delete_after_days cannot both be 0.
        """
        return pulumi.get(self, "retain_for_days")

    @property
    @pulumi.getter(name="retainForOverridable")
    def retain_for_overridable(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether or not the associated retain_for_days value can be overridden by org scopes. Note: if the associated
        retain_for_days value is 0, overridable is ignored.
        """
        return pulumi.get(self, "retain_for_overridable")

    @property
    @pulumi.getter(name="scopeId")
    def scope_id(self) -> pulumi.Output[str]:
        """
        The scope for this policy.
        """
        return pulumi.get(self, "scope_id")

