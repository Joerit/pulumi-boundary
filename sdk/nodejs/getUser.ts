// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

export function getUser(args: GetUserArgs, opts?: pulumi.InvokeOptions): Promise<GetUserResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("boundary:index/getUser:getUser", {
        "name": args.name,
        "scopeId": args.scopeId,
    }, opts);
}

/**
 * A collection of arguments for invoking getUser.
 */
export interface GetUserArgs {
    name: string;
    scopeId?: string;
}

/**
 * A collection of values returned by getUser.
 */
export interface GetUserResult {
    readonly accountIds: string[];
    readonly authorizedActions: string[];
    readonly description: string;
    readonly id: string;
    readonly loginName: string;
    readonly name: string;
    readonly primaryAccountId: string;
    readonly scopeId?: string;
    readonly scopes: outputs.GetUserScope[];
}
export function getUserOutput(args: GetUserOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetUserResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("boundary:index/getUser:getUser", {
        "name": args.name,
        "scopeId": args.scopeId,
    }, opts);
}

/**
 * A collection of arguments for invoking getUser.
 */
export interface GetUserOutputArgs {
    name: pulumi.Input<string>;
    scopeId?: pulumi.Input<string>;
}
