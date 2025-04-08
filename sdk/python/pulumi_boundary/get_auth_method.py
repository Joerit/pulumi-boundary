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
from . import outputs

__all__ = [
    'GetAuthMethodResult',
    'AwaitableGetAuthMethodResult',
    'get_auth_method',
    'get_auth_method_output',
]

@pulumi.output_type
class GetAuthMethodResult:
    """
    A collection of values returned by getAuthMethod.
    """
    def __init__(__self__, description=None, id=None, name=None, scope_id=None, scopes=None, type=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if scope_id and not isinstance(scope_id, str):
            raise TypeError("Expected argument 'scope_id' to be a str")
        pulumi.set(__self__, "scope_id", scope_id)
        if scopes and not isinstance(scopes, list):
            raise TypeError("Expected argument 'scopes' to be a list")
        pulumi.set(__self__, "scopes", scopes)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="scopeId")
    def scope_id(self) -> Optional[str]:
        return pulumi.get(self, "scope_id")

    @property
    @pulumi.getter
    def scopes(self) -> Sequence['outputs.GetAuthMethodScopeResult']:
        return pulumi.get(self, "scopes")

    @property
    @pulumi.getter
    def type(self) -> str:
        return pulumi.get(self, "type")


class AwaitableGetAuthMethodResult(GetAuthMethodResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAuthMethodResult(
            description=self.description,
            id=self.id,
            name=self.name,
            scope_id=self.scope_id,
            scopes=self.scopes,
            type=self.type)


def get_auth_method(name: Optional[str] = None,
                    scope_id: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAuthMethodResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['scopeId'] = scope_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('boundary:index/getAuthMethod:getAuthMethod', __args__, opts=opts, typ=GetAuthMethodResult).value

    return AwaitableGetAuthMethodResult(
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        scope_id=pulumi.get(__ret__, 'scope_id'),
        scopes=pulumi.get(__ret__, 'scopes'),
        type=pulumi.get(__ret__, 'type'))
def get_auth_method_output(name: Optional[pulumi.Input[str]] = None,
                           scope_id: Optional[pulumi.Input[Optional[str]]] = None,
                           opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetAuthMethodResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['scopeId'] = scope_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('boundary:index/getAuthMethod:getAuthMethod', __args__, opts=opts, typ=GetAuthMethodResult)
    return __ret__.apply(lambda __response__: GetAuthMethodResult(
        description=pulumi.get(__response__, 'description'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        scope_id=pulumi.get(__response__, 'scope_id'),
        scopes=pulumi.get(__response__, 'scopes'),
        type=pulumi.get(__response__, 'type')))
