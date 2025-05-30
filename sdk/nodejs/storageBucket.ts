// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The storage bucket resource allows you to configure a Boundary storage bucket. A storage bucket can only belong to the Global scope or an Org scope. At this time, the only supported storage for storage buckets is AWS S3. This feature requires Boundary Enterprise or Boundary HCP.
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
 *     scopeId: global.id,
 *     autoCreateAdminRole: true,
 *     autoCreateDefaultRole: true,
 * });
 * const awsStaticCredentialsExample = new boundary.StorageBucket("aws_static_credentials_example", {
 *     name: "My aws storage bucket with static credentials",
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
 * const awsDynamicCredentialsExample = new boundary.StorageBucket("aws_dynamic_credentials_example", {
 *     name: "My aws storage bucket with dynamic credentials",
 *     description: "My first storage bucket!",
 *     scopeId: org.id,
 *     pluginName: "aws",
 *     bucketName: "mybucket",
 *     attributesJson: JSON.stringify({
 *         region: "us-east-1",
 *         role_arn: "arn:aws:iam::123456789012:role/S3Access",
 *         disable_credential_rotation: "true",
 *     }),
 *     workerFilter: "\"pki\" in \"/tags/type\"",
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import boundary:index/storageBucket:StorageBucket foo <my-id>
 * ```
 */
export class StorageBucket extends pulumi.CustomResource {
    /**
     * Get an existing StorageBucket resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StorageBucketState, opts?: pulumi.CustomResourceOptions): StorageBucket {
        return new StorageBucket(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'boundary:index/storageBucket:StorageBucket';

    /**
     * Returns true if the given object is an instance of StorageBucket.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StorageBucket {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StorageBucket.__pulumiType;
    }

    /**
     * The attributes for the storage bucket. The "region" attribute field is required when creating an AWS storage bucket. Values are either encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" or remove the block to clear all attributes in the storage bucket.
     */
    public readonly attributesJson!: pulumi.Output<string | undefined>;
    /**
     * The name of the bucket within the external object store service.
     */
    public readonly bucketName!: pulumi.Output<string>;
    /**
     * The prefix used to organize the data held within the external object store.
     */
    public readonly bucketPrefix!: pulumi.Output<string | undefined>;
    /**
     * The storage bucket description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Internal only. Used to force update so that we can always check the value of secrets.
     */
    public /*out*/ readonly internalForceUpdate!: pulumi.Output<string>;
    /**
     * Internal only. The Boundary-provided HMAC used to calculate the current value of the HMAC'd config. Used for drift detection.
     */
    public /*out*/ readonly internalHmacUsedForSecretsConfigHmac!: pulumi.Output<string>;
    /**
     * Internal only. HMAC of (serverSecretsHmac + config secrets). Used for proper secrets handling.
     */
    public /*out*/ readonly internalSecretsConfigHmac!: pulumi.Output<string>;
    /**
     * The storage bucket name. Defaults to the resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the plugin that should back the resource. This or pluginName must be defined.
     */
    public readonly pluginId!: pulumi.Output<string>;
    /**
     * The name of the plugin that should back the resource. This or pluginId must be defined.
     */
    public readonly pluginName!: pulumi.Output<string | undefined>;
    /**
     * The scope for this storage bucket.
     */
    public readonly scopeId!: pulumi.Output<string>;
    /**
     * The HMAC'd secrets value returned from the server.
     */
    public /*out*/ readonly secretsHmac!: pulumi.Output<string>;
    /**
     * The secrets for the storage bucket. Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" to clear any existing values. NOTE: Unlike "attributesJson", removing this block will NOT clear secrets from the storage bucket; this allows injecting secrets for one call, then removing them for storage.
     */
    public readonly secretsJson!: pulumi.Output<string | undefined>;
    /**
     * Filters to the worker(s) that can handle requests for this storage bucket. The filter must match an existing worker in order to create a storage bucket.
     */
    public readonly workerFilter!: pulumi.Output<string>;

    /**
     * Create a StorageBucket resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StorageBucketArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StorageBucketArgs | StorageBucketState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as StorageBucketState | undefined;
            resourceInputs["attributesJson"] = state ? state.attributesJson : undefined;
            resourceInputs["bucketName"] = state ? state.bucketName : undefined;
            resourceInputs["bucketPrefix"] = state ? state.bucketPrefix : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["internalForceUpdate"] = state ? state.internalForceUpdate : undefined;
            resourceInputs["internalHmacUsedForSecretsConfigHmac"] = state ? state.internalHmacUsedForSecretsConfigHmac : undefined;
            resourceInputs["internalSecretsConfigHmac"] = state ? state.internalSecretsConfigHmac : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["pluginId"] = state ? state.pluginId : undefined;
            resourceInputs["pluginName"] = state ? state.pluginName : undefined;
            resourceInputs["scopeId"] = state ? state.scopeId : undefined;
            resourceInputs["secretsHmac"] = state ? state.secretsHmac : undefined;
            resourceInputs["secretsJson"] = state ? state.secretsJson : undefined;
            resourceInputs["workerFilter"] = state ? state.workerFilter : undefined;
        } else {
            const args = argsOrState as StorageBucketArgs | undefined;
            if ((!args || args.bucketName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bucketName'");
            }
            if ((!args || args.scopeId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scopeId'");
            }
            if ((!args || args.workerFilter === undefined) && !opts.urn) {
                throw new Error("Missing required property 'workerFilter'");
            }
            resourceInputs["attributesJson"] = args ? args.attributesJson : undefined;
            resourceInputs["bucketName"] = args ? args.bucketName : undefined;
            resourceInputs["bucketPrefix"] = args ? args.bucketPrefix : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["pluginId"] = args ? args.pluginId : undefined;
            resourceInputs["pluginName"] = args ? args.pluginName : undefined;
            resourceInputs["scopeId"] = args ? args.scopeId : undefined;
            resourceInputs["secretsJson"] = args?.secretsJson ? pulumi.secret(args.secretsJson) : undefined;
            resourceInputs["workerFilter"] = args ? args.workerFilter : undefined;
            resourceInputs["internalForceUpdate"] = undefined /*out*/;
            resourceInputs["internalHmacUsedForSecretsConfigHmac"] = undefined /*out*/;
            resourceInputs["internalSecretsConfigHmac"] = undefined /*out*/;
            resourceInputs["secretsHmac"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["secretsJson"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(StorageBucket.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering StorageBucket resources.
 */
export interface StorageBucketState {
    /**
     * The attributes for the storage bucket. The "region" attribute field is required when creating an AWS storage bucket. Values are either encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" or remove the block to clear all attributes in the storage bucket.
     */
    attributesJson?: pulumi.Input<string>;
    /**
     * The name of the bucket within the external object store service.
     */
    bucketName?: pulumi.Input<string>;
    /**
     * The prefix used to organize the data held within the external object store.
     */
    bucketPrefix?: pulumi.Input<string>;
    /**
     * The storage bucket description.
     */
    description?: pulumi.Input<string>;
    /**
     * Internal only. Used to force update so that we can always check the value of secrets.
     */
    internalForceUpdate?: pulumi.Input<string>;
    /**
     * Internal only. The Boundary-provided HMAC used to calculate the current value of the HMAC'd config. Used for drift detection.
     */
    internalHmacUsedForSecretsConfigHmac?: pulumi.Input<string>;
    /**
     * Internal only. HMAC of (serverSecretsHmac + config secrets). Used for proper secrets handling.
     */
    internalSecretsConfigHmac?: pulumi.Input<string>;
    /**
     * The storage bucket name. Defaults to the resource name.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the plugin that should back the resource. This or pluginName must be defined.
     */
    pluginId?: pulumi.Input<string>;
    /**
     * The name of the plugin that should back the resource. This or pluginId must be defined.
     */
    pluginName?: pulumi.Input<string>;
    /**
     * The scope for this storage bucket.
     */
    scopeId?: pulumi.Input<string>;
    /**
     * The HMAC'd secrets value returned from the server.
     */
    secretsHmac?: pulumi.Input<string>;
    /**
     * The secrets for the storage bucket. Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" to clear any existing values. NOTE: Unlike "attributesJson", removing this block will NOT clear secrets from the storage bucket; this allows injecting secrets for one call, then removing them for storage.
     */
    secretsJson?: pulumi.Input<string>;
    /**
     * Filters to the worker(s) that can handle requests for this storage bucket. The filter must match an existing worker in order to create a storage bucket.
     */
    workerFilter?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a StorageBucket resource.
 */
export interface StorageBucketArgs {
    /**
     * The attributes for the storage bucket. The "region" attribute field is required when creating an AWS storage bucket. Values are either encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" or remove the block to clear all attributes in the storage bucket.
     */
    attributesJson?: pulumi.Input<string>;
    /**
     * The name of the bucket within the external object store service.
     */
    bucketName: pulumi.Input<string>;
    /**
     * The prefix used to organize the data held within the external object store.
     */
    bucketPrefix?: pulumi.Input<string>;
    /**
     * The storage bucket description.
     */
    description?: pulumi.Input<string>;
    /**
     * The storage bucket name. Defaults to the resource name.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the plugin that should back the resource. This or pluginName must be defined.
     */
    pluginId?: pulumi.Input<string>;
    /**
     * The name of the plugin that should back the resource. This or pluginId must be defined.
     */
    pluginName?: pulumi.Input<string>;
    /**
     * The scope for this storage bucket.
     */
    scopeId: pulumi.Input<string>;
    /**
     * The secrets for the storage bucket. Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" to clear any existing values. NOTE: Unlike "attributesJson", removing this block will NOT clear secrets from the storage bucket; this allows injecting secrets for one call, then removing them for storage.
     */
    secretsJson?: pulumi.Input<string>;
    /**
     * Filters to the worker(s) that can handle requests for this storage bucket. The filter must match an existing worker in order to create a storage bucket.
     */
    workerFilter: pulumi.Input<string>;
}
