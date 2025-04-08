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

__all__ = ['AccountArgs', 'Account']

@pulumi.input_type
class AccountArgs:
    def __init__(__self__, *,
                 auth_method_id: pulumi.Input[str],
                 type: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 login_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Account resource.
        :param pulumi.Input[str] auth_method_id: The resource ID for the auth method.
        :param pulumi.Input[str] type: The resource type.
        :param pulumi.Input[str] description: The account description.
        :param pulumi.Input[str] login_name: The login name for this account.
        :param pulumi.Input[str] name: The account name. Defaults to the resource name.
        :param pulumi.Input[str] password: The account password. Only set on create, changes will not be reflected when updating account.
        """
        pulumi.set(__self__, "auth_method_id", auth_method_id)
        pulumi.set(__self__, "type", type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if login_name is not None:
            pulumi.set(__self__, "login_name", login_name)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if password is not None:
            pulumi.set(__self__, "password", password)

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
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The resource type.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The account description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="loginName")
    def login_name(self) -> Optional[pulumi.Input[str]]:
        """
        The login name for this account.
        """
        return pulumi.get(self, "login_name")

    @login_name.setter
    def login_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "login_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The account name. Defaults to the resource name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        The account password. Only set on create, changes will not be reflected when updating account.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)


@pulumi.input_type
class _AccountState:
    def __init__(__self__, *,
                 auth_method_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 login_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Account resources.
        :param pulumi.Input[str] auth_method_id: The resource ID for the auth method.
        :param pulumi.Input[str] description: The account description.
        :param pulumi.Input[str] login_name: The login name for this account.
        :param pulumi.Input[str] name: The account name. Defaults to the resource name.
        :param pulumi.Input[str] password: The account password. Only set on create, changes will not be reflected when updating account.
        :param pulumi.Input[str] type: The resource type.
        """
        if auth_method_id is not None:
            pulumi.set(__self__, "auth_method_id", auth_method_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if login_name is not None:
            pulumi.set(__self__, "login_name", login_name)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if type is not None:
            pulumi.set(__self__, "type", type)

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
        The account description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="loginName")
    def login_name(self) -> Optional[pulumi.Input[str]]:
        """
        The login name for this account.
        """
        return pulumi.get(self, "login_name")

    @login_name.setter
    def login_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "login_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The account name. Defaults to the resource name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        The account password. Only set on create, changes will not be reflected when updating account.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The resource type.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class Account(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_method_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 login_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a Account resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auth_method_id: The resource ID for the auth method.
        :param pulumi.Input[str] description: The account description.
        :param pulumi.Input[str] login_name: The login name for this account.
        :param pulumi.Input[str] name: The account name. Defaults to the resource name.
        :param pulumi.Input[str] password: The account password. Only set on create, changes will not be reflected when updating account.
        :param pulumi.Input[str] type: The resource type.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AccountArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Account resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param AccountArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AccountArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_method_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 login_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AccountArgs.__new__(AccountArgs)

            if auth_method_id is None and not opts.urn:
                raise TypeError("Missing required property 'auth_method_id'")
            __props__.__dict__["auth_method_id"] = auth_method_id
            __props__.__dict__["description"] = description
            __props__.__dict__["login_name"] = login_name
            __props__.__dict__["name"] = name
            __props__.__dict__["password"] = None if password is None else pulumi.Output.secret(password)
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["password"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(Account, __self__).__init__(
            'boundary:index/account:Account',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auth_method_id: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            login_name: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'Account':
        """
        Get an existing Account resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auth_method_id: The resource ID for the auth method.
        :param pulumi.Input[str] description: The account description.
        :param pulumi.Input[str] login_name: The login name for this account.
        :param pulumi.Input[str] name: The account name. Defaults to the resource name.
        :param pulumi.Input[str] password: The account password. Only set on create, changes will not be reflected when updating account.
        :param pulumi.Input[str] type: The resource type.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AccountState.__new__(_AccountState)

        __props__.__dict__["auth_method_id"] = auth_method_id
        __props__.__dict__["description"] = description
        __props__.__dict__["login_name"] = login_name
        __props__.__dict__["name"] = name
        __props__.__dict__["password"] = password
        __props__.__dict__["type"] = type
        return Account(resource_name, opts=opts, __props__=__props__)

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
        The account description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="loginName")
    def login_name(self) -> pulumi.Output[Optional[str]]:
        """
        The login name for this account.
        """
        return pulumi.get(self, "login_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The account name. Defaults to the resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        """
        The account password. Only set on create, changes will not be reflected when updating account.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The resource type.
        """
        return pulumi.get(self, "type")

