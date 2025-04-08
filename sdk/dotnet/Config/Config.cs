// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Immutable;

namespace Pulumi.Boundary
{
    public static class Config
    {
        [global::System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Design", "IDE1006", Justification = 
        "Double underscore prefix used to avoid conflicts with variable names.")]
        private sealed class __Value<T>
        {
            private readonly Func<T> _getter;
            private T _value = default!;
            private bool _set;

            public __Value(Func<T> getter)
            {
                _getter = getter;
            }

            public T Get() => _set ? _value : _getter();

            public void Set(T value)
            {
                _value = value;
                _set = true;
            }
        }

        private static readonly global::Pulumi.Config __config = new global::Pulumi.Config("boundary");

        private static readonly __Value<string?> _addr = new __Value<string?>(() => __config.Get("addr"));
        /// <summary>
        /// The base url of the Boundary API, e.g. "http://127.0.0.1:9200". If not set, it will be read from the "BOUNDARY_ADDR" env
        /// var.
        /// </summary>
        public static string? Addr
        {
            get => _addr.Get();
            set => _addr.Set(value);
        }

        private static readonly __Value<string?> _authMethodId = new __Value<string?>(() => __config.Get("authMethodId"));
        /// <summary>
        /// The auth method ID e.g. ampw_1234567890. If not set, the default auth method for the given scope ID will be used.
        /// </summary>
        public static string? AuthMethodId
        {
            get => _authMethodId.Get();
            set => _authMethodId.Set(value);
        }

        private static readonly __Value<string?> _authMethodLoginName = new __Value<string?>(() => __config.Get("authMethodLoginName"));
        /// <summary>
        /// The auth method login name for password-style or ldap-style auth methods
        /// </summary>
        public static string? AuthMethodLoginName
        {
            get => _authMethodLoginName.Get();
            set => _authMethodLoginName.Set(value);
        }

        private static readonly __Value<string?> _authMethodPassword = new __Value<string?>(() => __config.Get("authMethodPassword"));
        /// <summary>
        /// The auth method password for password-style or ldap-style auth methods
        /// </summary>
        public static string? AuthMethodPassword
        {
            get => _authMethodPassword.Get();
            set => _authMethodPassword.Set(value);
        }

        private static readonly __Value<string?> _passwordAuthMethodLoginName = new __Value<string?>(() => __config.Get("passwordAuthMethodLoginName"));
        /// <summary>
        /// The auth method login name for password-style auth methods
        /// </summary>
        public static string? PasswordAuthMethodLoginName
        {
            get => _passwordAuthMethodLoginName.Get();
            set => _passwordAuthMethodLoginName.Set(value);
        }

        private static readonly __Value<string?> _passwordAuthMethodPassword = new __Value<string?>(() => __config.Get("passwordAuthMethodPassword"));
        /// <summary>
        /// The auth method password for password-style auth methods
        /// </summary>
        public static string? PasswordAuthMethodPassword
        {
            get => _passwordAuthMethodPassword.Get();
            set => _passwordAuthMethodPassword.Set(value);
        }

        private static readonly __Value<string?> _pluginExecutionDir = new __Value<string?>(() => __config.Get("pluginExecutionDir"));
        /// <summary>
        /// Specifies a directory that the Boundary provider can use to write and execute its built-in plugins.
        /// </summary>
        public static string? PluginExecutionDir
        {
            get => _pluginExecutionDir.Get();
            set => _pluginExecutionDir.Set(value);
        }

        private static readonly __Value<string?> _recoveryKmsHcl = new __Value<string?>(() => __config.Get("recoveryKmsHcl"));
        /// <summary>
        /// Can be a heredoc string or a path on disk. If set, the string/file will be parsed as HCL and used with the recovery KMS
        /// mechanism. While this is set, it will override any other authentication information; the KMS mechanism will always be
        /// used. See Boundary's KMS docs for examples: https://boundaryproject.io/docs/configuration/kms
        /// </summary>
        public static string? RecoveryKmsHcl
        {
            get => _recoveryKmsHcl.Get();
            set => _recoveryKmsHcl.Set(value);
        }

        private static readonly __Value<string?> _scopeId = new __Value<string?>(() => __config.Get("scopeId"));
        /// <summary>
        /// The scope ID for the default auth method.
        /// </summary>
        public static string? ScopeId
        {
            get => _scopeId.Get();
            set => _scopeId.Set(value);
        }

        private static readonly __Value<bool?> _tlsInsecure = new __Value<bool?>(() => __config.GetBoolean("tlsInsecure"));
        /// <summary>
        /// When set to true, does not validate the Boundary API endpoint certificate
        /// </summary>
        public static bool? TlsInsecure
        {
            get => _tlsInsecure.Get();
            set => _tlsInsecure.Set(value);
        }

        private static readonly __Value<string?> _token = new __Value<string?>(() => __config.Get("token"));
        /// <summary>
        /// The Boundary token to use, as a string or path on disk containing just the string. If set, the token read here will be
        /// used in place of authenticating with the auth method specified in "auth_method_id", although the recovery KMS mechanism
        /// will still override this. Can also be set with the BOUNDARY_TOKEN environment variable.
        /// </summary>
        public static string? Token
        {
            get => _token.Get();
            set => _token.Set(value);
        }

    }
}
