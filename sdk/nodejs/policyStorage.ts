// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class PolicyStorage extends pulumi.CustomResource {
    /**
     * Get an existing PolicyStorage resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PolicyStorageState, opts?: pulumi.CustomResourceOptions): PolicyStorage {
        return new PolicyStorage(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'boundary:index/policyStorage:PolicyStorage';

    /**
     * Returns true if the given object is an instance of PolicyStorage.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PolicyStorage {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PolicyStorage.__pulumiType;
    }

    /**
     * The number of days after which a session recording will be automatically deleted. Defaults to 0: never automatically
     * delete. However, deleteAfterDays and retainForDays cannot both be 0.
     */
    public readonly deleteAfterDays!: pulumi.Output<number | undefined>;
    /**
     * Whether or not the associated deleteAfterDays value can be overridden by org scopes. Note: if the associated
     * deleteAfterDays value is 0, overridable is ignored
     */
    public readonly deleteAfterOverridable!: pulumi.Output<boolean | undefined>;
    /**
     * The policy description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The policy name. Defaults to the resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The number of days a session recording is required to be stored. Defaults to 0: allow deletions at any time. However,
     * retainForDays and deleteAfterDays cannot both be 0.
     */
    public readonly retainForDays!: pulumi.Output<number | undefined>;
    /**
     * Whether or not the associated retainForDays value can be overridden by org scopes. Note: if the associated retainForDays
     * value is 0, overridable is ignored.
     */
    public readonly retainForOverridable!: pulumi.Output<boolean | undefined>;
    /**
     * The scope for this policy.
     */
    public readonly scopeId!: pulumi.Output<string>;

    /**
     * Create a PolicyStorage resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PolicyStorageArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PolicyStorageArgs | PolicyStorageState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PolicyStorageState | undefined;
            resourceInputs["deleteAfterDays"] = state ? state.deleteAfterDays : undefined;
            resourceInputs["deleteAfterOverridable"] = state ? state.deleteAfterOverridable : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["retainForDays"] = state ? state.retainForDays : undefined;
            resourceInputs["retainForOverridable"] = state ? state.retainForOverridable : undefined;
            resourceInputs["scopeId"] = state ? state.scopeId : undefined;
        } else {
            const args = argsOrState as PolicyStorageArgs | undefined;
            if ((!args || args.scopeId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scopeId'");
            }
            resourceInputs["deleteAfterDays"] = args ? args.deleteAfterDays : undefined;
            resourceInputs["deleteAfterOverridable"] = args ? args.deleteAfterOverridable : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["retainForDays"] = args ? args.retainForDays : undefined;
            resourceInputs["retainForOverridable"] = args ? args.retainForOverridable : undefined;
            resourceInputs["scopeId"] = args ? args.scopeId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PolicyStorage.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PolicyStorage resources.
 */
export interface PolicyStorageState {
    /**
     * The number of days after which a session recording will be automatically deleted. Defaults to 0: never automatically
     * delete. However, deleteAfterDays and retainForDays cannot both be 0.
     */
    deleteAfterDays?: pulumi.Input<number>;
    /**
     * Whether or not the associated deleteAfterDays value can be overridden by org scopes. Note: if the associated
     * deleteAfterDays value is 0, overridable is ignored
     */
    deleteAfterOverridable?: pulumi.Input<boolean>;
    /**
     * The policy description.
     */
    description?: pulumi.Input<string>;
    /**
     * The policy name. Defaults to the resource name.
     */
    name?: pulumi.Input<string>;
    /**
     * The number of days a session recording is required to be stored. Defaults to 0: allow deletions at any time. However,
     * retainForDays and deleteAfterDays cannot both be 0.
     */
    retainForDays?: pulumi.Input<number>;
    /**
     * Whether or not the associated retainForDays value can be overridden by org scopes. Note: if the associated retainForDays
     * value is 0, overridable is ignored.
     */
    retainForOverridable?: pulumi.Input<boolean>;
    /**
     * The scope for this policy.
     */
    scopeId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PolicyStorage resource.
 */
export interface PolicyStorageArgs {
    /**
     * The number of days after which a session recording will be automatically deleted. Defaults to 0: never automatically
     * delete. However, deleteAfterDays and retainForDays cannot both be 0.
     */
    deleteAfterDays?: pulumi.Input<number>;
    /**
     * Whether or not the associated deleteAfterDays value can be overridden by org scopes. Note: if the associated
     * deleteAfterDays value is 0, overridable is ignored
     */
    deleteAfterOverridable?: pulumi.Input<boolean>;
    /**
     * The policy description.
     */
    description?: pulumi.Input<string>;
    /**
     * The policy name. Defaults to the resource name.
     */
    name?: pulumi.Input<string>;
    /**
     * The number of days a session recording is required to be stored. Defaults to 0: allow deletions at any time. However,
     * retainForDays and deleteAfterDays cannot both be 0.
     */
    retainForDays?: pulumi.Input<number>;
    /**
     * Whether or not the associated retainForDays value can be overridden by org scopes. Note: if the associated retainForDays
     * value is 0, overridable is ignored.
     */
    retainForOverridable?: pulumi.Input<boolean>;
    /**
     * The scope for this policy.
     */
    scopeId: pulumi.Input<string>;
}
