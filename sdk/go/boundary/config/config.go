// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/joerit/pulumi-boundary/sdk/go/boundary/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

var _ = internal.GetEnvOrDefault

// The base url of the Boundary API, e.g. "http://127.0.0.1:9200". If not set, it will be read from the "BOUNDARY_ADDR" env
// var.
func GetAddr(ctx *pulumi.Context) string {
	return config.Get(ctx, "boundary:addr")
}

// The auth method ID e.g. ampw_1234567890. If not set, the default auth method for the given scope ID will be used.
func GetAuthMethodId(ctx *pulumi.Context) string {
	return config.Get(ctx, "boundary:authMethodId")
}

// The auth method login name for password-style or ldap-style auth methods
func GetAuthMethodLoginName(ctx *pulumi.Context) string {
	return config.Get(ctx, "boundary:authMethodLoginName")
}

// The auth method password for password-style or ldap-style auth methods
func GetAuthMethodPassword(ctx *pulumi.Context) string {
	return config.Get(ctx, "boundary:authMethodPassword")
}

// The auth method login name for password-style auth methods
//
// Deprecated: Use authMethodLoginName instead
func GetPasswordAuthMethodLoginName(ctx *pulumi.Context) string {
	return config.Get(ctx, "boundary:passwordAuthMethodLoginName")
}

// The auth method password for password-style auth methods
//
// Deprecated: Use authMethodPassword instead
func GetPasswordAuthMethodPassword(ctx *pulumi.Context) string {
	return config.Get(ctx, "boundary:passwordAuthMethodPassword")
}

// Specifies a directory that the Boundary provider can use to write and execute its built-in plugins.
func GetPluginExecutionDir(ctx *pulumi.Context) string {
	return config.Get(ctx, "boundary:pluginExecutionDir")
}

// Can be a heredoc string or a path on disk. If set, the string/file will be parsed as HCL and used with the recovery KMS
// mechanism. While this is set, it will override any other authentication information; the KMS mechanism will always be
// used. See Boundary's KMS docs for examples: https://boundaryproject.io/docs/configuration/kms
func GetRecoveryKmsHcl(ctx *pulumi.Context) string {
	return config.Get(ctx, "boundary:recoveryKmsHcl")
}

// The scope ID for the default auth method.
func GetScopeId(ctx *pulumi.Context) string {
	return config.Get(ctx, "boundary:scopeId")
}

// When set to true, does not validate the Boundary API endpoint certificate
func GetTlsInsecure(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "boundary:tlsInsecure")
}

// The Boundary token to use, as a string or path on disk containing just the string. If set, the token read here will be
// used in place of authenticating with the auth method specified in "authMethodId", although the recovery KMS mechanism
// will still override this. Can also be set with the BOUNDARY_TOKEN environment variable.
func GetToken(ctx *pulumi.Context) string {
	return config.Get(ctx, "boundary:token")
}
