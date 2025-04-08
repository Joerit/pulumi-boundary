# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .account import *
from .account_ldap import *
from .account_oidc import *
from .account_password import *
from .alias_target import *
from .auth_method import *
from .auth_method_ldap import *
from .auth_method_oidc import *
from .auth_method_password import *
from .credential_json import *
from .credential_library_vault import *
from .credential_library_vault_ssh_certificate import *
from .credential_ssh_private_key import *
from .credential_store_static import *
from .credential_store_vault import *
from .credential_username_password import *
from .get_account import *
from .get_auth_method import *
from .get_group import *
from .get_scope import *
from .get_user import *
from .group import *
from .host import *
from .host_catalog import *
from .host_catalog_plugin import *
from .host_catalog_static import *
from .host_set import *
from .host_set_plugin import *
from .host_set_static import *
from .host_static import *
from .managed_group import *
from .managed_group_ldap import *
from .policy_storage import *
from .provider import *
from .role import *
from .scope import *
from .scope_policy_attachment import *
from .storage_bucket import *
from .target import *
from .user import *
from .worker import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_boundary.config as __config
    config = __config
    import pulumi_boundary.region as __region
    region = __region
else:
    config = _utilities.lazy_import('pulumi_boundary.config')
    region = _utilities.lazy_import('pulumi_boundary.region')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "boundary",
  "mod": "index/account",
  "fqn": "pulumi_boundary",
  "classes": {
   "boundary:index/account:Account": "Account"
  }
 },
 {
  "pkg": "boundary",
  "mod": "index/accountLdap",
  "fqn": "pulumi_boundary",
  "classes": {
   "boundary:index/accountLdap:AccountLdap": "AccountLdap"
  }
 },
 {
  "pkg": "boundary",
  "mod": "index/accountOidc",
  "fqn": "pulumi_boundary",
  "classes": {
   "boundary:index/accountOidc:AccountOidc": "AccountOidc"
  }
 },
 {
  "pkg": "boundary",
  "mod": "index/accountPassword",
  "fqn": "pulumi_boundary",
  "classes": {
   "boundary:index/accountPassword:AccountPassword": "AccountPassword"
  }
 },
 {
  "pkg": "boundary",
  "mod": "index/aliasTarget",
  "fqn": "pulumi_boundary",
  "classes": {
   "boundary:index/aliasTarget:AliasTarget": "AliasTarget"
  }
 },
 {
  "pkg": "boundary",
  "mod": "index/authMethod",
  "fqn": "pulumi_boundary",
  "classes": {
   "boundary:index/authMethod:AuthMethod": "AuthMethod"
  }
 },
 {
  "pkg": "boundary",
  "mod": "index/authMethodLdap",
  "fqn": "pulumi_boundary",
  "classes": {
   "boundary:index/authMethodLdap:AuthMethodLdap": "AuthMethodLdap"
  }
 },
 {
  "pkg": "boundary",
  "mod": "index/authMethodOidc",
  "fqn": "pulumi_boundary",
  "classes": {
   "boundary:index/authMethodOidc:AuthMethodOidc": "AuthMethodOidc"
  }
 },
 {
  "pkg": "boundary",
  "mod": "index/authMethodPassword",
  "fqn": "pulumi_boundary",
  "classes": {
   "boundary:index/authMethodPassword:AuthMethodPassword": "AuthMethodPassword"
  }
 },
 {
  "pkg": "boundary",
  "mod": "index/credentialJson",
  "fqn": "pulumi_boundary",
  "classes": {
   "boundary:index/credentialJson:CredentialJson": "CredentialJson"
  }
 },
 {
  "pkg": "boundary",
  "mod": "index/credentialLibraryVault",
  "fqn": "pulumi_boundary",
  "classes": {
   "boundary:index/credentialLibraryVault:CredentialLibraryVault": "CredentialLibraryVault"
  }
 },
 {
  "pkg": "boundary",
  "mod": "index/credentialLibraryVaultSshCertificate",
  "fqn": "pulumi_boundary",
  "classes": {
   "boundary:index/credentialLibraryVaultSshCertificate:CredentialLibraryVaultSshCertificate": "CredentialLibraryVaultSshCertificate"
  }
 },
 {
  "pkg": "boundary",
  "mod": "index/credentialSshPrivateKey",
  "fqn": "pulumi_boundary",
  "classes": {
   "boundary:index/credentialSshPrivateKey:CredentialSshPrivateKey": "CredentialSshPrivateKey"
  }
 },
 {
  "pkg": "boundary",
  "mod": "index/credentialStoreStatic",
  "fqn": "pulumi_boundary",
  "classes": {
   "boundary:index/credentialStoreStatic:CredentialStoreStatic": "CredentialStoreStatic"
  }
 },
 {
  "pkg": "boundary",
  "mod": "index/credentialStoreVault",
  "fqn": "pulumi_boundary",
  "classes": {
   "boundary:index/credentialStoreVault:CredentialStoreVault": "CredentialStoreVault"
  }
 },
 {
  "pkg": "boundary",
  "mod": "index/credentialUsernamePassword",
  "fqn": "pulumi_boundary",
  "classes": {
   "boundary:index/credentialUsernamePassword:CredentialUsernamePassword": "CredentialUsernamePassword"
  }
 },
 {
  "pkg": "boundary",
  "mod": "index/group",
  "fqn": "pulumi_boundary",
  "classes": {
   "boundary:index/group:Group": "Group"
  }
 },
 {
  "pkg": "boundary",
  "mod": "index/host",
  "fqn": "pulumi_boundary",
  "classes": {
   "boundary:index/host:Host": "Host"
  }
 },
 {
  "pkg": "boundary",
  "mod": "index/hostCatalog",
  "fqn": "pulumi_boundary",
  "classes": {
   "boundary:index/hostCatalog:HostCatalog": "HostCatalog"
  }
 },
 {
  "pkg": "boundary",
  "mod": "index/hostCatalogPlugin",
  "fqn": "pulumi_boundary",
  "classes": {
   "boundary:index/hostCatalogPlugin:HostCatalogPlugin": "HostCatalogPlugin"
  }
 },
 {
  "pkg": "boundary",
  "mod": "index/hostCatalogStatic",
  "fqn": "pulumi_boundary",
  "classes": {
   "boundary:index/hostCatalogStatic:HostCatalogStatic": "HostCatalogStatic"
  }
 },
 {
  "pkg": "boundary",
  "mod": "index/hostSet",
  "fqn": "pulumi_boundary",
  "classes": {
   "boundary:index/hostSet:HostSet": "HostSet"
  }
 },
 {
  "pkg": "boundary",
  "mod": "index/hostSetPlugin",
  "fqn": "pulumi_boundary",
  "classes": {
   "boundary:index/hostSetPlugin:HostSetPlugin": "HostSetPlugin"
  }
 },
 {
  "pkg": "boundary",
  "mod": "index/hostSetStatic",
  "fqn": "pulumi_boundary",
  "classes": {
   "boundary:index/hostSetStatic:HostSetStatic": "HostSetStatic"
  }
 },
 {
  "pkg": "boundary",
  "mod": "index/hostStatic",
  "fqn": "pulumi_boundary",
  "classes": {
   "boundary:index/hostStatic:HostStatic": "HostStatic"
  }
 },
 {
  "pkg": "boundary",
  "mod": "index/managedGroup",
  "fqn": "pulumi_boundary",
  "classes": {
   "boundary:index/managedGroup:ManagedGroup": "ManagedGroup"
  }
 },
 {
  "pkg": "boundary",
  "mod": "index/managedGroupLdap",
  "fqn": "pulumi_boundary",
  "classes": {
   "boundary:index/managedGroupLdap:ManagedGroupLdap": "ManagedGroupLdap"
  }
 },
 {
  "pkg": "boundary",
  "mod": "index/policyStorage",
  "fqn": "pulumi_boundary",
  "classes": {
   "boundary:index/policyStorage:PolicyStorage": "PolicyStorage"
  }
 },
 {
  "pkg": "boundary",
  "mod": "index/role",
  "fqn": "pulumi_boundary",
  "classes": {
   "boundary:index/role:Role": "Role"
  }
 },
 {
  "pkg": "boundary",
  "mod": "index/scope",
  "fqn": "pulumi_boundary",
  "classes": {
   "boundary:index/scope:Scope": "Scope"
  }
 },
 {
  "pkg": "boundary",
  "mod": "index/scopePolicyAttachment",
  "fqn": "pulumi_boundary",
  "classes": {
   "boundary:index/scopePolicyAttachment:ScopePolicyAttachment": "ScopePolicyAttachment"
  }
 },
 {
  "pkg": "boundary",
  "mod": "index/storageBucket",
  "fqn": "pulumi_boundary",
  "classes": {
   "boundary:index/storageBucket:StorageBucket": "StorageBucket"
  }
 },
 {
  "pkg": "boundary",
  "mod": "index/target",
  "fqn": "pulumi_boundary",
  "classes": {
   "boundary:index/target:Target": "Target"
  }
 },
 {
  "pkg": "boundary",
  "mod": "index/user",
  "fqn": "pulumi_boundary",
  "classes": {
   "boundary:index/user:User": "User"
  }
 },
 {
  "pkg": "boundary",
  "mod": "index/worker",
  "fqn": "pulumi_boundary",
  "classes": {
   "boundary:index/worker:Worker": "Worker"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "boundary",
  "token": "pulumi:providers:boundary",
  "fqn": "pulumi_boundary",
  "class": "Provider"
 }
]
"""
)
