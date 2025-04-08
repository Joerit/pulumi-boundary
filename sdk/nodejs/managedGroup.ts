// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The managed group resource allows you to configure a Boundary group.
 */
export class ManagedGroup extends pulumi.CustomResource {
    /**
     * Get an existing ManagedGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ManagedGroupState, opts?: pulumi.CustomResourceOptions): ManagedGroup {
        return new ManagedGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'boundary:index/managedGroup:ManagedGroup';

    /**
     * Returns true if the given object is an instance of ManagedGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ManagedGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ManagedGroup.__pulumiType;
    }

    /**
     * The resource ID for the auth method.
     */
    public readonly authMethodId!: pulumi.Output<string>;
    /**
     * The managed group description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Boolean expression to filter the workers for this managed group.
     */
    public readonly filter!: pulumi.Output<string>;
    /**
     * The managed group name. Defaults to the resource name.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a ManagedGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ManagedGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ManagedGroupArgs | ManagedGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ManagedGroupState | undefined;
            resourceInputs["authMethodId"] = state ? state.authMethodId : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["filter"] = state ? state.filter : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as ManagedGroupArgs | undefined;
            if ((!args || args.authMethodId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'authMethodId'");
            }
            if ((!args || args.filter === undefined) && !opts.urn) {
                throw new Error("Missing required property 'filter'");
            }
            resourceInputs["authMethodId"] = args ? args.authMethodId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["filter"] = args ? args.filter : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ManagedGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ManagedGroup resources.
 */
export interface ManagedGroupState {
    /**
     * The resource ID for the auth method.
     */
    authMethodId?: pulumi.Input<string>;
    /**
     * The managed group description.
     */
    description?: pulumi.Input<string>;
    /**
     * Boolean expression to filter the workers for this managed group.
     */
    filter?: pulumi.Input<string>;
    /**
     * The managed group name. Defaults to the resource name.
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ManagedGroup resource.
 */
export interface ManagedGroupArgs {
    /**
     * The resource ID for the auth method.
     */
    authMethodId: pulumi.Input<string>;
    /**
     * The managed group description.
     */
    description?: pulumi.Input<string>;
    /**
     * Boolean expression to filter the workers for this managed group.
     */
    filter: pulumi.Input<string>;
    /**
     * The managed group name. Defaults to the resource name.
     */
    name?: pulumi.Input<string>;
}
