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
from .. import _utilities

addr: Optional[str]
"""
The base url of the Boundary API, e.g. "http://127.0.0.1:9200". If not set, it will be read from the "BOUNDARY_ADDR" env
var.
"""

authMethodId: Optional[str]
"""
The auth method ID e.g. ampw_1234567890. If not set, the default auth method for the given scope ID will be used.
"""

authMethodLoginName: Optional[str]
"""
The auth method login name for password-style or ldap-style auth methods
"""

authMethodPassword: Optional[str]
"""
The auth method password for password-style or ldap-style auth methods
"""

passwordAuthMethodLoginName: Optional[str]
"""
The auth method login name for password-style auth methods
"""

passwordAuthMethodPassword: Optional[str]
"""
The auth method password for password-style auth methods
"""

pluginExecutionDir: Optional[str]
"""
Specifies a directory that the Boundary provider can use to write and execute its built-in plugins.
"""

recoveryKmsHcl: Optional[str]
"""
Can be a heredoc string or a path on disk. If set, the string/file will be parsed as HCL and used with the recovery KMS
mechanism. While this is set, it will override any other authentication information; the KMS mechanism will always be
used. See Boundary's KMS docs for examples: https://boundaryproject.io/docs/configuration/kms
"""

scopeId: Optional[str]
"""
The scope ID for the default auth method.
"""

tlsInsecure: Optional[bool]
"""
When set to true, does not validate the Boundary API endpoint certificate
"""

token: Optional[str]
"""
The Boundary token to use, as a string or path on disk containing just the string. If set, the token read here will be
used in place of authenticating with the auth method specified in "auth_method_id", although the recovery KMS mechanism
will still override this. Can also be set with the BOUNDARY_TOKEN environment variable.
"""

