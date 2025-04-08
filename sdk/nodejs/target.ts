// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The target resource allows you to configure a Boundary target.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as boundary from "@pulumi/boundary";
 *
 * const global = new boundary.Scope("global", {
 *     globalScope: true,
 *     scopeId: "global",
 * });
 * const org = new boundary.Scope("org", {
 *     name: "organization_one",
 *     description: "My first scope!",
 *     scopeId: global.id,
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
 *     name: "vault_store",
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
 *     credentialType: "username_password",
 * });
 * const fooHostCatalog = new boundary.HostCatalog("foo", {
 *     name: "test",
 *     description: "test catalog",
 *     scopeId: project.id,
 *     type: "static",
 * });
 * const fooHost = new boundary.Host("foo", {
 *     type: "static",
 *     name: "foo",
 *     hostCatalogId: fooHostCatalog.id,
 *     address: "10.0.0.1",
 * });
 * const bar = new boundary.Host("bar", {
 *     type: "static",
 *     name: "bar",
 *     hostCatalogId: fooHostCatalog.id,
 *     address: "10.0.0.1",
 * });
 * const fooHostSet = new boundary.HostSet("foo", {
 *     type: "static",
 *     name: "foo",
 *     hostCatalogId: fooHostCatalog.id,
 *     hostIds: [
 *         fooHost.id,
 *         bar.id,
 *     ],
 * });
 * const awsExample = new boundary.StorageBucket("aws_example", {
 *     name: "My aws storage bucket",
 *     description: "My first storage bucket!",
 *     scopeId: org.id,
 *     pluginName: "aws",
 *     bucketName: "mybucket",
 *     attributesJson: JSON.stringify({
 *         region: "us-east-1",
 *     }),
 *     secretsJson: JSON.stringify({
 *         access_key_id: "aws_access_key_id_value",
 *         secret_access_key: "aws_secret_access_key_value",
 *     }),
 *     workerFilter: "\"pki\" in \"/tags/type\"",
 * });
 * const fooTarget = new boundary.Target("foo", {
 *     name: "foo",
 *     description: "Foo target",
 *     type: "tcp",
 *     defaultPort: 22,
 *     scopeId: project.id,
 *     hostSourceIds: [fooHostSet.id],
 *     brokeredCredentialSourceIds: [fooCredentialLibraryVault.id],
 * });
 * const sshFoo = new boundary.Target("ssh_foo", {
 *     name: "ssh_foo",
 *     description: "Ssh target",
 *     type: "ssh",
 *     defaultPort: 22,
 *     scopeId: project.id,
 *     hostSourceIds: [fooHostSet.id],
 *     injectedApplicationCredentialSourceIds: [fooCredentialLibraryVault.id],
 * });
 * const sshSessionRecordingFoo = new boundary.Target("ssh_session_recording_foo", {
 *     name: "ssh_foo",
 *     description: "Ssh target",
 *     type: "ssh",
 *     defaultPort: 22,
 *     scopeId: project.id,
 *     hostSourceIds: [fooHostSet.id],
 *     injectedApplicationCredentialSourceIds: [fooCredentialLibraryVault.id],
 *     enableSessionRecording: true,
 *     storageBucketId: awsExample,
 * });
 * const addressFoo = new boundary.Target("address_foo", {
 *     name: "address_foo",
 *     description: "Foo target with an address",
 *     type: "tcp",
 *     defaultPort: 22,
 *     scopeId: project.id,
 *     address: "127.0.0.1",
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import boundary:index/target:Target foo <my-id>
 * ```
 */
export class Target extends pulumi.CustomResource {
    /**
     * Get an existing Target resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TargetState, opts?: pulumi.CustomResourceOptions): Target {
        return new Target(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'boundary:index/target:Target';

    /**
     * Returns true if the given object is an instance of Target.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Target {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Target.__pulumiType;
    }

    /**
     * Optionally, a valid network address to connect to for this target. Cannot be used alongside host*source*ids.
     */
    public readonly address!: pulumi.Output<string | undefined>;
    /**
     * A list of brokered credential source ID's.
     */
    public readonly brokeredCredentialSourceIds!: pulumi.Output<string[] | undefined>;
    /**
     * The default client port for this target.
     */
    public readonly defaultClientPort!: pulumi.Output<number | undefined>;
    /**
     * The default port for this target.
     */
    public readonly defaultPort!: pulumi.Output<number | undefined>;
    /**
     * The target description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Boolean expression to filter the workers used to access this target
     */
    public readonly egressWorkerFilter!: pulumi.Output<string | undefined>;
    /**
     * HCP/Ent Only. Enable sessions recording for this target. Only applicable for SSH targets.
     */
    public readonly enableSessionRecording!: pulumi.Output<boolean | undefined>;
    /**
     * A list of host source ID's. Cannot be used alongside address.
     */
    public readonly hostSourceIds!: pulumi.Output<string[] | undefined>;
    /**
     * HCP Only. Boolean expression to filter the workers a user will connect to when initiating a session against this target
     */
    public readonly ingressWorkerFilter!: pulumi.Output<string | undefined>;
    /**
     * A list of injected application credential source ID's.
     */
    public readonly injectedApplicationCredentialSourceIds!: pulumi.Output<string[] | undefined>;
    /**
     * The target name. Defaults to the resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The scope ID in which the resource is created. Defaults to the provider's `defaultScope` if unset.
     */
    public readonly scopeId!: pulumi.Output<string>;
    public readonly sessionConnectionLimit!: pulumi.Output<number>;
    public readonly sessionMaxSeconds!: pulumi.Output<number>;
    /**
     * HCP/Ent Only. Storage bucket for this target. Only applicable for SSH targets.
     */
    public readonly storageBucketId!: pulumi.Output<string | undefined>;
    /**
     * The target resource type.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * Boolean expression to filter the workers for this target
     *
     * @deprecated Deprecated. Use `egressWorkerFilter` and `ingressWorkerFilter` instead
     */
    public readonly workerFilter!: pulumi.Output<string | undefined>;

    /**
     * Create a Target resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TargetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TargetArgs | TargetState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TargetState | undefined;
            resourceInputs["address"] = state ? state.address : undefined;
            resourceInputs["brokeredCredentialSourceIds"] = state ? state.brokeredCredentialSourceIds : undefined;
            resourceInputs["defaultClientPort"] = state ? state.defaultClientPort : undefined;
            resourceInputs["defaultPort"] = state ? state.defaultPort : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["egressWorkerFilter"] = state ? state.egressWorkerFilter : undefined;
            resourceInputs["enableSessionRecording"] = state ? state.enableSessionRecording : undefined;
            resourceInputs["hostSourceIds"] = state ? state.hostSourceIds : undefined;
            resourceInputs["ingressWorkerFilter"] = state ? state.ingressWorkerFilter : undefined;
            resourceInputs["injectedApplicationCredentialSourceIds"] = state ? state.injectedApplicationCredentialSourceIds : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["scopeId"] = state ? state.scopeId : undefined;
            resourceInputs["sessionConnectionLimit"] = state ? state.sessionConnectionLimit : undefined;
            resourceInputs["sessionMaxSeconds"] = state ? state.sessionMaxSeconds : undefined;
            resourceInputs["storageBucketId"] = state ? state.storageBucketId : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["workerFilter"] = state ? state.workerFilter : undefined;
        } else {
            const args = argsOrState as TargetArgs | undefined;
            if ((!args || args.scopeId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scopeId'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["address"] = args ? args.address : undefined;
            resourceInputs["brokeredCredentialSourceIds"] = args ? args.brokeredCredentialSourceIds : undefined;
            resourceInputs["defaultClientPort"] = args ? args.defaultClientPort : undefined;
            resourceInputs["defaultPort"] = args ? args.defaultPort : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["egressWorkerFilter"] = args ? args.egressWorkerFilter : undefined;
            resourceInputs["enableSessionRecording"] = args ? args.enableSessionRecording : undefined;
            resourceInputs["hostSourceIds"] = args ? args.hostSourceIds : undefined;
            resourceInputs["ingressWorkerFilter"] = args ? args.ingressWorkerFilter : undefined;
            resourceInputs["injectedApplicationCredentialSourceIds"] = args ? args.injectedApplicationCredentialSourceIds : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["scopeId"] = args ? args.scopeId : undefined;
            resourceInputs["sessionConnectionLimit"] = args ? args.sessionConnectionLimit : undefined;
            resourceInputs["sessionMaxSeconds"] = args ? args.sessionMaxSeconds : undefined;
            resourceInputs["storageBucketId"] = args ? args.storageBucketId : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["workerFilter"] = args ? args.workerFilter : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Target.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Target resources.
 */
export interface TargetState {
    /**
     * Optionally, a valid network address to connect to for this target. Cannot be used alongside host*source*ids.
     */
    address?: pulumi.Input<string>;
    /**
     * A list of brokered credential source ID's.
     */
    brokeredCredentialSourceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The default client port for this target.
     */
    defaultClientPort?: pulumi.Input<number>;
    /**
     * The default port for this target.
     */
    defaultPort?: pulumi.Input<number>;
    /**
     * The target description.
     */
    description?: pulumi.Input<string>;
    /**
     * Boolean expression to filter the workers used to access this target
     */
    egressWorkerFilter?: pulumi.Input<string>;
    /**
     * HCP/Ent Only. Enable sessions recording for this target. Only applicable for SSH targets.
     */
    enableSessionRecording?: pulumi.Input<boolean>;
    /**
     * A list of host source ID's. Cannot be used alongside address.
     */
    hostSourceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * HCP Only. Boolean expression to filter the workers a user will connect to when initiating a session against this target
     */
    ingressWorkerFilter?: pulumi.Input<string>;
    /**
     * A list of injected application credential source ID's.
     */
    injectedApplicationCredentialSourceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The target name. Defaults to the resource name.
     */
    name?: pulumi.Input<string>;
    /**
     * The scope ID in which the resource is created. Defaults to the provider's `defaultScope` if unset.
     */
    scopeId?: pulumi.Input<string>;
    sessionConnectionLimit?: pulumi.Input<number>;
    sessionMaxSeconds?: pulumi.Input<number>;
    /**
     * HCP/Ent Only. Storage bucket for this target. Only applicable for SSH targets.
     */
    storageBucketId?: pulumi.Input<string>;
    /**
     * The target resource type.
     */
    type?: pulumi.Input<string>;
    /**
     * Boolean expression to filter the workers for this target
     *
     * @deprecated Deprecated. Use `egressWorkerFilter` and `ingressWorkerFilter` instead
     */
    workerFilter?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Target resource.
 */
export interface TargetArgs {
    /**
     * Optionally, a valid network address to connect to for this target. Cannot be used alongside host*source*ids.
     */
    address?: pulumi.Input<string>;
    /**
     * A list of brokered credential source ID's.
     */
    brokeredCredentialSourceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The default client port for this target.
     */
    defaultClientPort?: pulumi.Input<number>;
    /**
     * The default port for this target.
     */
    defaultPort?: pulumi.Input<number>;
    /**
     * The target description.
     */
    description?: pulumi.Input<string>;
    /**
     * Boolean expression to filter the workers used to access this target
     */
    egressWorkerFilter?: pulumi.Input<string>;
    /**
     * HCP/Ent Only. Enable sessions recording for this target. Only applicable for SSH targets.
     */
    enableSessionRecording?: pulumi.Input<boolean>;
    /**
     * A list of host source ID's. Cannot be used alongside address.
     */
    hostSourceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * HCP Only. Boolean expression to filter the workers a user will connect to when initiating a session against this target
     */
    ingressWorkerFilter?: pulumi.Input<string>;
    /**
     * A list of injected application credential source ID's.
     */
    injectedApplicationCredentialSourceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The target name. Defaults to the resource name.
     */
    name?: pulumi.Input<string>;
    /**
     * The scope ID in which the resource is created. Defaults to the provider's `defaultScope` if unset.
     */
    scopeId: pulumi.Input<string>;
    sessionConnectionLimit?: pulumi.Input<number>;
    sessionMaxSeconds?: pulumi.Input<number>;
    /**
     * HCP/Ent Only. Storage bucket for this target. Only applicable for SSH targets.
     */
    storageBucketId?: pulumi.Input<string>;
    /**
     * The target resource type.
     */
    type: pulumi.Input<string>;
    /**
     * Boolean expression to filter the workers for this target
     *
     * @deprecated Deprecated. Use `egressWorkerFilter` and `ingressWorkerFilter` instead
     */
    workerFilter?: pulumi.Input<string>;
}
