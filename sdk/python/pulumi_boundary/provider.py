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

__all__ = ['ProviderArgs', 'Provider']

@pulumi.input_type
class ProviderArgs:
    def __init__(__self__, *,
                 addr: pulumi.Input[str],
                 auth_method_id: Optional[pulumi.Input[str]] = None,
                 auth_method_login_name: Optional[pulumi.Input[str]] = None,
                 auth_method_password: Optional[pulumi.Input[str]] = None,
                 password_auth_method_login_name: Optional[pulumi.Input[str]] = None,
                 password_auth_method_password: Optional[pulumi.Input[str]] = None,
                 plugin_execution_dir: Optional[pulumi.Input[str]] = None,
                 recovery_kms_hcl: Optional[pulumi.Input[str]] = None,
                 scope_id: Optional[pulumi.Input[str]] = None,
                 tls_insecure: Optional[pulumi.Input[bool]] = None,
                 token: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Provider resource.
        :param pulumi.Input[str] addr: The base url of the Boundary API, e.g. "http://127.0.0.1:9200". If not set, it will be read from the "BOUNDARY_ADDR" env
               var.
        :param pulumi.Input[str] auth_method_id: The auth method ID e.g. ampw_1234567890. If not set, the default auth method for the given scope ID will be used.
        :param pulumi.Input[str] auth_method_login_name: The auth method login name for password-style or ldap-style auth methods
        :param pulumi.Input[str] auth_method_password: The auth method password for password-style or ldap-style auth methods
        :param pulumi.Input[str] password_auth_method_login_name: The auth method login name for password-style auth methods
        :param pulumi.Input[str] password_auth_method_password: The auth method password for password-style auth methods
        :param pulumi.Input[str] plugin_execution_dir: Specifies a directory that the Boundary provider can use to write and execute its built-in plugins.
        :param pulumi.Input[str] recovery_kms_hcl: Can be a heredoc string or a path on disk. If set, the string/file will be parsed as HCL and used with the recovery KMS
               mechanism. While this is set, it will override any other authentication information; the KMS mechanism will always be
               used. See Boundary's KMS docs for examples: https://boundaryproject.io/docs/configuration/kms
        :param pulumi.Input[str] scope_id: The scope ID for the default auth method.
        :param pulumi.Input[bool] tls_insecure: When set to true, does not validate the Boundary API endpoint certificate
        :param pulumi.Input[str] token: The Boundary token to use, as a string or path on disk containing just the string. If set, the token read here will be
               used in place of authenticating with the auth method specified in "auth_method_id", although the recovery KMS mechanism
               will still override this. Can also be set with the BOUNDARY_TOKEN environment variable.
        """
        pulumi.set(__self__, "addr", addr)
        if auth_method_id is not None:
            pulumi.set(__self__, "auth_method_id", auth_method_id)
        if auth_method_login_name is not None:
            pulumi.set(__self__, "auth_method_login_name", auth_method_login_name)
        if auth_method_password is not None:
            pulumi.set(__self__, "auth_method_password", auth_method_password)
        if password_auth_method_login_name is not None:
            warnings.warn("""Use auth_method_login_name instead""", DeprecationWarning)
            pulumi.log.warn("""password_auth_method_login_name is deprecated: Use auth_method_login_name instead""")
        if password_auth_method_login_name is not None:
            pulumi.set(__self__, "password_auth_method_login_name", password_auth_method_login_name)
        if password_auth_method_password is not None:
            warnings.warn("""Use auth_method_password instead""", DeprecationWarning)
            pulumi.log.warn("""password_auth_method_password is deprecated: Use auth_method_password instead""")
        if password_auth_method_password is not None:
            pulumi.set(__self__, "password_auth_method_password", password_auth_method_password)
        if plugin_execution_dir is not None:
            pulumi.set(__self__, "plugin_execution_dir", plugin_execution_dir)
        if recovery_kms_hcl is not None:
            pulumi.set(__self__, "recovery_kms_hcl", recovery_kms_hcl)
        if scope_id is not None:
            pulumi.set(__self__, "scope_id", scope_id)
        if tls_insecure is not None:
            pulumi.set(__self__, "tls_insecure", tls_insecure)
        if token is not None:
            pulumi.set(__self__, "token", token)

    @property
    @pulumi.getter
    def addr(self) -> pulumi.Input[str]:
        """
        The base url of the Boundary API, e.g. "http://127.0.0.1:9200". If not set, it will be read from the "BOUNDARY_ADDR" env
        var.
        """
        return pulumi.get(self, "addr")

    @addr.setter
    def addr(self, value: pulumi.Input[str]):
        pulumi.set(self, "addr", value)

    @property
    @pulumi.getter(name="authMethodId")
    def auth_method_id(self) -> Optional[pulumi.Input[str]]:
        """
        The auth method ID e.g. ampw_1234567890. If not set, the default auth method for the given scope ID will be used.
        """
        return pulumi.get(self, "auth_method_id")

    @auth_method_id.setter
    def auth_method_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_method_id", value)

    @property
    @pulumi.getter(name="authMethodLoginName")
    def auth_method_login_name(self) -> Optional[pulumi.Input[str]]:
        """
        The auth method login name for password-style or ldap-style auth methods
        """
        return pulumi.get(self, "auth_method_login_name")

    @auth_method_login_name.setter
    def auth_method_login_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_method_login_name", value)

    @property
    @pulumi.getter(name="authMethodPassword")
    def auth_method_password(self) -> Optional[pulumi.Input[str]]:
        """
        The auth method password for password-style or ldap-style auth methods
        """
        return pulumi.get(self, "auth_method_password")

    @auth_method_password.setter
    def auth_method_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_method_password", value)

    @property
    @pulumi.getter(name="passwordAuthMethodLoginName")
    @_utilities.deprecated("""Use auth_method_login_name instead""")
    def password_auth_method_login_name(self) -> Optional[pulumi.Input[str]]:
        """
        The auth method login name for password-style auth methods
        """
        return pulumi.get(self, "password_auth_method_login_name")

    @password_auth_method_login_name.setter
    def password_auth_method_login_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password_auth_method_login_name", value)

    @property
    @pulumi.getter(name="passwordAuthMethodPassword")
    @_utilities.deprecated("""Use auth_method_password instead""")
    def password_auth_method_password(self) -> Optional[pulumi.Input[str]]:
        """
        The auth method password for password-style auth methods
        """
        return pulumi.get(self, "password_auth_method_password")

    @password_auth_method_password.setter
    def password_auth_method_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password_auth_method_password", value)

    @property
    @pulumi.getter(name="pluginExecutionDir")
    def plugin_execution_dir(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies a directory that the Boundary provider can use to write and execute its built-in plugins.
        """
        return pulumi.get(self, "plugin_execution_dir")

    @plugin_execution_dir.setter
    def plugin_execution_dir(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "plugin_execution_dir", value)

    @property
    @pulumi.getter(name="recoveryKmsHcl")
    def recovery_kms_hcl(self) -> Optional[pulumi.Input[str]]:
        """
        Can be a heredoc string or a path on disk. If set, the string/file will be parsed as HCL and used with the recovery KMS
        mechanism. While this is set, it will override any other authentication information; the KMS mechanism will always be
        used. See Boundary's KMS docs for examples: https://boundaryproject.io/docs/configuration/kms
        """
        return pulumi.get(self, "recovery_kms_hcl")

    @recovery_kms_hcl.setter
    def recovery_kms_hcl(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "recovery_kms_hcl", value)

    @property
    @pulumi.getter(name="scopeId")
    def scope_id(self) -> Optional[pulumi.Input[str]]:
        """
        The scope ID for the default auth method.
        """
        return pulumi.get(self, "scope_id")

    @scope_id.setter
    def scope_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scope_id", value)

    @property
    @pulumi.getter(name="tlsInsecure")
    def tls_insecure(self) -> Optional[pulumi.Input[bool]]:
        """
        When set to true, does not validate the Boundary API endpoint certificate
        """
        return pulumi.get(self, "tls_insecure")

    @tls_insecure.setter
    def tls_insecure(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "tls_insecure", value)

    @property
    @pulumi.getter
    def token(self) -> Optional[pulumi.Input[str]]:
        """
        The Boundary token to use, as a string or path on disk containing just the string. If set, the token read here will be
        used in place of authenticating with the auth method specified in "auth_method_id", although the recovery KMS mechanism
        will still override this. Can also be set with the BOUNDARY_TOKEN environment variable.
        """
        return pulumi.get(self, "token")

    @token.setter
    def token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "token", value)


class Provider(pulumi.ProviderResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 addr: Optional[pulumi.Input[str]] = None,
                 auth_method_id: Optional[pulumi.Input[str]] = None,
                 auth_method_login_name: Optional[pulumi.Input[str]] = None,
                 auth_method_password: Optional[pulumi.Input[str]] = None,
                 password_auth_method_login_name: Optional[pulumi.Input[str]] = None,
                 password_auth_method_password: Optional[pulumi.Input[str]] = None,
                 plugin_execution_dir: Optional[pulumi.Input[str]] = None,
                 recovery_kms_hcl: Optional[pulumi.Input[str]] = None,
                 scope_id: Optional[pulumi.Input[str]] = None,
                 tls_insecure: Optional[pulumi.Input[bool]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The provider type for the boundary package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] addr: The base url of the Boundary API, e.g. "http://127.0.0.1:9200". If not set, it will be read from the "BOUNDARY_ADDR" env
               var.
        :param pulumi.Input[str] auth_method_id: The auth method ID e.g. ampw_1234567890. If not set, the default auth method for the given scope ID will be used.
        :param pulumi.Input[str] auth_method_login_name: The auth method login name for password-style or ldap-style auth methods
        :param pulumi.Input[str] auth_method_password: The auth method password for password-style or ldap-style auth methods
        :param pulumi.Input[str] password_auth_method_login_name: The auth method login name for password-style auth methods
        :param pulumi.Input[str] password_auth_method_password: The auth method password for password-style auth methods
        :param pulumi.Input[str] plugin_execution_dir: Specifies a directory that the Boundary provider can use to write and execute its built-in plugins.
        :param pulumi.Input[str] recovery_kms_hcl: Can be a heredoc string or a path on disk. If set, the string/file will be parsed as HCL and used with the recovery KMS
               mechanism. While this is set, it will override any other authentication information; the KMS mechanism will always be
               used. See Boundary's KMS docs for examples: https://boundaryproject.io/docs/configuration/kms
        :param pulumi.Input[str] scope_id: The scope ID for the default auth method.
        :param pulumi.Input[bool] tls_insecure: When set to true, does not validate the Boundary API endpoint certificate
        :param pulumi.Input[str] token: The Boundary token to use, as a string or path on disk containing just the string. If set, the token read here will be
               used in place of authenticating with the auth method specified in "auth_method_id", although the recovery KMS mechanism
               will still override this. Can also be set with the BOUNDARY_TOKEN environment variable.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProviderArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The provider type for the boundary package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param ProviderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProviderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 addr: Optional[pulumi.Input[str]] = None,
                 auth_method_id: Optional[pulumi.Input[str]] = None,
                 auth_method_login_name: Optional[pulumi.Input[str]] = None,
                 auth_method_password: Optional[pulumi.Input[str]] = None,
                 password_auth_method_login_name: Optional[pulumi.Input[str]] = None,
                 password_auth_method_password: Optional[pulumi.Input[str]] = None,
                 plugin_execution_dir: Optional[pulumi.Input[str]] = None,
                 recovery_kms_hcl: Optional[pulumi.Input[str]] = None,
                 scope_id: Optional[pulumi.Input[str]] = None,
                 tls_insecure: Optional[pulumi.Input[bool]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProviderArgs.__new__(ProviderArgs)

            if addr is None and not opts.urn:
                raise TypeError("Missing required property 'addr'")
            __props__.__dict__["addr"] = addr
            __props__.__dict__["auth_method_id"] = auth_method_id
            __props__.__dict__["auth_method_login_name"] = auth_method_login_name
            __props__.__dict__["auth_method_password"] = auth_method_password
            __props__.__dict__["password_auth_method_login_name"] = password_auth_method_login_name
            __props__.__dict__["password_auth_method_password"] = password_auth_method_password
            __props__.__dict__["plugin_execution_dir"] = plugin_execution_dir
            __props__.__dict__["recovery_kms_hcl"] = recovery_kms_hcl
            __props__.__dict__["scope_id"] = scope_id
            __props__.__dict__["tls_insecure"] = pulumi.Output.from_input(tls_insecure).apply(pulumi.runtime.to_json) if tls_insecure is not None else None
            __props__.__dict__["token"] = token
        super(Provider, __self__).__init__(
            'boundary',
            resource_name,
            __props__,
            opts)

    @property
    @pulumi.getter
    def addr(self) -> pulumi.Output[str]:
        """
        The base url of the Boundary API, e.g. "http://127.0.0.1:9200". If not set, it will be read from the "BOUNDARY_ADDR" env
        var.
        """
        return pulumi.get(self, "addr")

    @property
    @pulumi.getter(name="authMethodId")
    def auth_method_id(self) -> pulumi.Output[Optional[str]]:
        """
        The auth method ID e.g. ampw_1234567890. If not set, the default auth method for the given scope ID will be used.
        """
        return pulumi.get(self, "auth_method_id")

    @property
    @pulumi.getter(name="authMethodLoginName")
    def auth_method_login_name(self) -> pulumi.Output[Optional[str]]:
        """
        The auth method login name for password-style or ldap-style auth methods
        """
        return pulumi.get(self, "auth_method_login_name")

    @property
    @pulumi.getter(name="authMethodPassword")
    def auth_method_password(self) -> pulumi.Output[Optional[str]]:
        """
        The auth method password for password-style or ldap-style auth methods
        """
        return pulumi.get(self, "auth_method_password")

    @property
    @pulumi.getter(name="passwordAuthMethodLoginName")
    @_utilities.deprecated("""Use auth_method_login_name instead""")
    def password_auth_method_login_name(self) -> pulumi.Output[Optional[str]]:
        """
        The auth method login name for password-style auth methods
        """
        return pulumi.get(self, "password_auth_method_login_name")

    @property
    @pulumi.getter(name="passwordAuthMethodPassword")
    @_utilities.deprecated("""Use auth_method_password instead""")
    def password_auth_method_password(self) -> pulumi.Output[Optional[str]]:
        """
        The auth method password for password-style auth methods
        """
        return pulumi.get(self, "password_auth_method_password")

    @property
    @pulumi.getter(name="pluginExecutionDir")
    def plugin_execution_dir(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies a directory that the Boundary provider can use to write and execute its built-in plugins.
        """
        return pulumi.get(self, "plugin_execution_dir")

    @property
    @pulumi.getter(name="recoveryKmsHcl")
    def recovery_kms_hcl(self) -> pulumi.Output[Optional[str]]:
        """
        Can be a heredoc string or a path on disk. If set, the string/file will be parsed as HCL and used with the recovery KMS
        mechanism. While this is set, it will override any other authentication information; the KMS mechanism will always be
        used. See Boundary's KMS docs for examples: https://boundaryproject.io/docs/configuration/kms
        """
        return pulumi.get(self, "recovery_kms_hcl")

    @property
    @pulumi.getter(name="scopeId")
    def scope_id(self) -> pulumi.Output[Optional[str]]:
        """
        The scope ID for the default auth method.
        """
        return pulumi.get(self, "scope_id")

    @property
    @pulumi.getter
    def token(self) -> pulumi.Output[Optional[str]]:
        """
        The Boundary token to use, as a string or path on disk containing just the string. If set, the token read here will be
        used in place of authenticating with the auth method specified in "auth_method_id", although the recovery KMS mechanism
        will still override this. Can also be set with the BOUNDARY_TOKEN environment variable.
        """
        return pulumi.get(self, "token")

