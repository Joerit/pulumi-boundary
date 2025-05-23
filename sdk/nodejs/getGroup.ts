// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The boundary.Group data source allows you to find a Boundary group.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as boundary from "@pulumi/boundary";
 *
 * // Retrieve a user from the global scope
 * const globalGroup = boundary.getGroup({
 *     name: "admin",
 * });
 * // User from an org scope
 * const org = boundary.getScope({
 *     name: "org",
 *     scopeId: "global",
 * });
 * const orgGroup = org.then(org => boundary.getGroup({
 *     name: "username",
 *     scopeId: org.id,
 * }));
 * ```
 */
export function getGroup(args: GetGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetGroupResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("boundary:index/getGroup:getGroup", {
        "name": args.name,
        "scopeId": args.scopeId,
    }, opts);
}

/**
 * A collection of arguments for invoking getGroup.
 */
export interface GetGroupArgs {
    /**
     * The name of the group to retrieve.
     */
    name: string;
    /**
     * The scope ID in which the resource is created. Defaults `global` if unset.
     */
    scopeId?: string;
}

/**
 * A collection of values returned by getGroup.
 */
export interface GetGroupResult {
    /**
     * The description of the retrieved group.
     */
    readonly description: string;
    /**
     * The ID of the retrieved group.
     */
    readonly id: string;
    /**
     * Resource IDs for group members, these are most likely boundary users.
     */
    readonly memberIds: string[];
    /**
     * The name of the group to retrieve.
     */
    readonly name: string;
    /**
     * The scope ID in which the resource is created. Defaults `global` if unset.
     */
    readonly scopeId?: string;
    readonly scopes: outputs.GetGroupScope[];
}
/**
 * The boundary.Group data source allows you to find a Boundary group.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as boundary from "@pulumi/boundary";
 *
 * // Retrieve a user from the global scope
 * const globalGroup = boundary.getGroup({
 *     name: "admin",
 * });
 * // User from an org scope
 * const org = boundary.getScope({
 *     name: "org",
 *     scopeId: "global",
 * });
 * const orgGroup = org.then(org => boundary.getGroup({
 *     name: "username",
 *     scopeId: org.id,
 * }));
 * ```
 */
export function getGroupOutput(args: GetGroupOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetGroupResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("boundary:index/getGroup:getGroup", {
        "name": args.name,
        "scopeId": args.scopeId,
    }, opts);
}

/**
 * A collection of arguments for invoking getGroup.
 */
export interface GetGroupOutputArgs {
    /**
     * The name of the group to retrieve.
     */
    name: pulumi.Input<string>;
    /**
     * The scope ID in which the resource is created. Defaults `global` if unset.
     */
    scopeId?: pulumi.Input<string>;
}
