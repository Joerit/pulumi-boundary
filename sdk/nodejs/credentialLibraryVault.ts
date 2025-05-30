// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The credential library for Vault resource allows you to configure a Boundary credential library for Vault.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as boundary from "@pulumi/boundary";
 *
 * const org = new boundary.Scope("org", {
 *     name: "organization_one",
 *     description: "My first scope!",
 *     scopeId: "global",
 *     autoCreateAdminRole: true,
 *     autoCreateDefaultRole: true,
 * });
 * const project = new boundary.Scope("project", {
 *     name: "project_one",
 *     description: "My first scope!",
 *     scopeId: org.id,
 *     autoCreateAdminRole: true,
 * });
 * const foo = new boundary.CredentialStoreVault("foo", {
 *     name: "foo",
 *     description: "My first Vault credential store!",
 *     address: "http://127.0.0.1:8200",
 *     token: "s.0ufRo6XEGU2jOqnIr7OlFYP5",
 *     scopeId: project.id,
 * });
 * const fooCredentialLibraryVault = new boundary.CredentialLibraryVault("foo", {
 *     name: "foo",
 *     description: "My first Vault credential library!",
 *     credentialStoreId: foo.id,
 *     path: "my/secret/foo",
 *     httpMethod: "GET",
 * });
 * const bar = new boundary.CredentialLibraryVault("bar", {
 *     name: "bar",
 *     description: "My second Vault credential library!",
 *     credentialStoreId: foo.id,
 *     path: "my/secret/bar",
 *     httpMethod: "POST",
 *     httpRequestBody: `{
 *   "key": "Value",
 * }
 * `,
 * });
 * const baz = new boundary.CredentialLibraryVault("baz", {
 *     name: "baz",
 *     description: "vault username password credential with mapping overrides",
 *     credentialStoreId: foo.id,
 *     path: "my/secret/baz",
 *     httpMethod: "GET",
 *     credentialType: "username_password",
 *     credentialMappingOverrides: {
 *         password_attribute: "alternative_password_label",
 *         username_attribute: "alternative_username_label",
 *     },
 * });
 * const quz = new boundary.CredentialLibraryVault("quz", {
 *     name: "quz",
 *     description: "vault ssh private key credential with mapping overrides",
 *     credentialStoreId: foo.id,
 *     path: "my/secret/quz",
 *     httpMethod: "GET",
 *     credentialType: "ssh_private_key",
 *     credentialMappingOverrides: {
 *         private_key_attribute: "alternative_key_label",
 *         private_key_passphrase_attribute: "alternative_passphrase_label",
 *         username_attribute: "alternative_username_label",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import boundary:index/credentialLibraryVault:CredentialLibraryVault foo <my-id>
 * ```
 */
export class CredentialLibraryVault extends pulumi.CustomResource {
    /**
     * Get an existing CredentialLibraryVault resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CredentialLibraryVaultState, opts?: pulumi.CustomResourceOptions): CredentialLibraryVault {
        return new CredentialLibraryVault(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'boundary:index/credentialLibraryVault:CredentialLibraryVault';

    /**
     * Returns true if the given object is an instance of CredentialLibraryVault.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CredentialLibraryVault {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CredentialLibraryVault.__pulumiType;
    }

    /**
     * The credential mapping override.
     */
    public readonly credentialMappingOverrides!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The ID of the credential store that this library belongs to.
     */
    public readonly credentialStoreId!: pulumi.Output<string>;
    /**
     * The type of credential the library generates. Cannot be updated on an existing resource.
     */
    public readonly credentialType!: pulumi.Output<string | undefined>;
    /**
     * The Vault credential library description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The HTTP method the library uses when requesting credentials from Vault. Defaults to 'GET'
     */
    public readonly httpMethod!: pulumi.Output<string | undefined>;
    /**
     * The body of the HTTP request the library sends to Vault when requesting credentials. Only valid if `httpMethod` is set to `POST`.
     */
    public readonly httpRequestBody!: pulumi.Output<string | undefined>;
    /**
     * The Vault credential library name. Defaults to the resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The path in Vault to request credentials from.
     */
    public readonly path!: pulumi.Output<string>;

    /**
     * Create a CredentialLibraryVault resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CredentialLibraryVaultArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CredentialLibraryVaultArgs | CredentialLibraryVaultState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CredentialLibraryVaultState | undefined;
            resourceInputs["credentialMappingOverrides"] = state ? state.credentialMappingOverrides : undefined;
            resourceInputs["credentialStoreId"] = state ? state.credentialStoreId : undefined;
            resourceInputs["credentialType"] = state ? state.credentialType : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["httpMethod"] = state ? state.httpMethod : undefined;
            resourceInputs["httpRequestBody"] = state ? state.httpRequestBody : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["path"] = state ? state.path : undefined;
        } else {
            const args = argsOrState as CredentialLibraryVaultArgs | undefined;
            if ((!args || args.credentialStoreId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'credentialStoreId'");
            }
            if ((!args || args.path === undefined) && !opts.urn) {
                throw new Error("Missing required property 'path'");
            }
            resourceInputs["credentialMappingOverrides"] = args ? args.credentialMappingOverrides : undefined;
            resourceInputs["credentialStoreId"] = args ? args.credentialStoreId : undefined;
            resourceInputs["credentialType"] = args ? args.credentialType : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["httpMethod"] = args ? args.httpMethod : undefined;
            resourceInputs["httpRequestBody"] = args ? args.httpRequestBody : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["path"] = args ? args.path : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CredentialLibraryVault.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CredentialLibraryVault resources.
 */
export interface CredentialLibraryVaultState {
    /**
     * The credential mapping override.
     */
    credentialMappingOverrides?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the credential store that this library belongs to.
     */
    credentialStoreId?: pulumi.Input<string>;
    /**
     * The type of credential the library generates. Cannot be updated on an existing resource.
     */
    credentialType?: pulumi.Input<string>;
    /**
     * The Vault credential library description.
     */
    description?: pulumi.Input<string>;
    /**
     * The HTTP method the library uses when requesting credentials from Vault. Defaults to 'GET'
     */
    httpMethod?: pulumi.Input<string>;
    /**
     * The body of the HTTP request the library sends to Vault when requesting credentials. Only valid if `httpMethod` is set to `POST`.
     */
    httpRequestBody?: pulumi.Input<string>;
    /**
     * The Vault credential library name. Defaults to the resource name.
     */
    name?: pulumi.Input<string>;
    /**
     * The path in Vault to request credentials from.
     */
    path?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CredentialLibraryVault resource.
 */
export interface CredentialLibraryVaultArgs {
    /**
     * The credential mapping override.
     */
    credentialMappingOverrides?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the credential store that this library belongs to.
     */
    credentialStoreId: pulumi.Input<string>;
    /**
     * The type of credential the library generates. Cannot be updated on an existing resource.
     */
    credentialType?: pulumi.Input<string>;
    /**
     * The Vault credential library description.
     */
    description?: pulumi.Input<string>;
    /**
     * The HTTP method the library uses when requesting credentials from Vault. Defaults to 'GET'
     */
    httpMethod?: pulumi.Input<string>;
    /**
     * The body of the HTTP request the library sends to Vault when requesting credentials. Only valid if `httpMethod` is set to `POST`.
     */
    httpRequestBody?: pulumi.Input<string>;
    /**
     * The Vault credential library name. Defaults to the resource name.
     */
    name?: pulumi.Input<string>;
    /**
     * The path in Vault to request credentials from.
     */
    path: pulumi.Input<string>;
}
