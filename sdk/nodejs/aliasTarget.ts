// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The target alias resource allows you to configure a Boundary target alias.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as boundary from "@pulumi/boundary";
 *
 * const org = new boundary.Scope("org", {
 *     name: "organization_one",
 *     description: "global scope",
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
 * const foo = new boundary.HostCatalogStatic("foo", {
 *     name: "test",
 *     description: "test catalog",
 *     scopeId: project.id,
 * });
 * const fooHostStatic = new boundary.HostStatic("foo", {
 *     name: "foo",
 *     hostCatalogId: foo.id,
 *     address: "10.0.0.1",
 * });
 * const bar = new boundary.HostStatic("bar", {
 *     name: "bar",
 *     hostCatalogId: foo.id,
 *     address: "127.0.0.1",
 * });
 * const fooHostSetStatic = new boundary.HostSetStatic("foo", {
 *     name: "foo",
 *     hostCatalogId: foo.id,
 *     hostIds: [
 *         fooHostStatic.id,
 *         bar.id,
 *     ],
 * });
 * const fooTarget = new boundary.Target("foo", {
 *     name: "foo",
 *     description: "Foo target",
 *     type: "tcp",
 *     defaultPort: 22,
 *     scopeId: project.id,
 *     hostSourceIds: [fooHostSetStatic.id],
 * });
 * const exampleAliasTarget = new boundary.AliasTarget("example_alias_target", {
 *     name: "example_alias_target",
 *     description: "Example alias to target foo using host boundary_host_static.bar",
 *     scopeId: "global",
 *     value: "example.bar.foo.boundary",
 *     destinationId: fooTarget.id,
 *     authorizeSessionHostId: bar.id,
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import boundary:index/aliasTarget:AliasTarget example_alias_target <my-id>
 * ```
 */
export class AliasTarget extends pulumi.CustomResource {
    /**
     * Get an existing AliasTarget resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AliasTargetState, opts?: pulumi.CustomResourceOptions): AliasTarget {
        return new AliasTarget(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'boundary:index/aliasTarget:AliasTarget';

    /**
     * Returns true if the given object is an instance of AliasTarget.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AliasTarget {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AliasTarget.__pulumiType;
    }

    /**
     * The host id to pass to Boundary when performing an authorize session action.
     */
    public readonly authorizeSessionHostId!: pulumi.Output<string | undefined>;
    /**
     * The alias description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The destination of the alias.
     */
    public readonly destinationId!: pulumi.Output<string | undefined>;
    /**
     * The alias name. Defaults to the resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The scope ID.
     */
    public readonly scopeId!: pulumi.Output<string>;
    /**
     * The type of alias; hardcoded.
     */
    public readonly type!: pulumi.Output<string | undefined>;
    /**
     * The value of the alias.
     */
    public readonly value!: pulumi.Output<string>;

    /**
     * Create a AliasTarget resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AliasTargetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AliasTargetArgs | AliasTargetState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AliasTargetState | undefined;
            resourceInputs["authorizeSessionHostId"] = state ? state.authorizeSessionHostId : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["destinationId"] = state ? state.destinationId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["scopeId"] = state ? state.scopeId : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["value"] = state ? state.value : undefined;
        } else {
            const args = argsOrState as AliasTargetArgs | undefined;
            if ((!args || args.scopeId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scopeId'");
            }
            if ((!args || args.value === undefined) && !opts.urn) {
                throw new Error("Missing required property 'value'");
            }
            resourceInputs["authorizeSessionHostId"] = args ? args.authorizeSessionHostId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["destinationId"] = args ? args.destinationId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["scopeId"] = args ? args.scopeId : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["value"] = args ? args.value : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AliasTarget.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AliasTarget resources.
 */
export interface AliasTargetState {
    /**
     * The host id to pass to Boundary when performing an authorize session action.
     */
    authorizeSessionHostId?: pulumi.Input<string>;
    /**
     * The alias description.
     */
    description?: pulumi.Input<string>;
    /**
     * The destination of the alias.
     */
    destinationId?: pulumi.Input<string>;
    /**
     * The alias name. Defaults to the resource name.
     */
    name?: pulumi.Input<string>;
    /**
     * The scope ID.
     */
    scopeId?: pulumi.Input<string>;
    /**
     * The type of alias; hardcoded.
     */
    type?: pulumi.Input<string>;
    /**
     * The value of the alias.
     */
    value?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AliasTarget resource.
 */
export interface AliasTargetArgs {
    /**
     * The host id to pass to Boundary when performing an authorize session action.
     */
    authorizeSessionHostId?: pulumi.Input<string>;
    /**
     * The alias description.
     */
    description?: pulumi.Input<string>;
    /**
     * The destination of the alias.
     */
    destinationId?: pulumi.Input<string>;
    /**
     * The alias name. Defaults to the resource name.
     */
    name?: pulumi.Input<string>;
    /**
     * The scope ID.
     */
    scopeId: pulumi.Input<string>;
    /**
     * The type of alias; hardcoded.
     */
    type?: pulumi.Input<string>;
    /**
     * The value of the alias.
     */
    value: pulumi.Input<string>;
}
