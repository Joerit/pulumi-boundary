// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The boundary.AuthMethod data source allows you to find a Boundary auth method.
 */
export function getAuthMethod(args: GetAuthMethodArgs, opts?: pulumi.InvokeOptions): Promise<GetAuthMethodResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("boundary:index/getAuthMethod:getAuthMethod", {
        "name": args.name,
        "scopeId": args.scopeId,
    }, opts);
}

/**
 * A collection of arguments for invoking getAuthMethod.
 */
export interface GetAuthMethodArgs {
    /**
     * The name of the auth method to retrieve.
     */
    name: string;
    /**
     * The scope ID in which the resource is created. Defaults `global` if unset.
     */
    scopeId?: string;
}

/**
 * A collection of values returned by getAuthMethod.
 */
export interface GetAuthMethodResult {
    /**
     * The description of the retrieved auth method.
     */
    readonly description: string;
    /**
     * The ID of the retrieved auth method.
     */
    readonly id: string;
    /**
     * The name of the auth method to retrieve.
     */
    readonly name: string;
    /**
     * The scope ID in which the resource is created. Defaults `global` if unset.
     */
    readonly scopeId?: string;
    readonly scopes: outputs.GetAuthMethodScope[];
    /**
     * The type of the auth method
     */
    readonly type: string;
}
/**
 * The boundary.AuthMethod data source allows you to find a Boundary auth method.
 */
export function getAuthMethodOutput(args: GetAuthMethodOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetAuthMethodResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("boundary:index/getAuthMethod:getAuthMethod", {
        "name": args.name,
        "scopeId": args.scopeId,
    }, opts);
}

/**
 * A collection of arguments for invoking getAuthMethod.
 */
export interface GetAuthMethodOutputArgs {
    /**
     * The name of the auth method to retrieve.
     */
    name: pulumi.Input<string>;
    /**
     * The scope ID in which the resource is created. Defaults `global` if unset.
     */
    scopeId?: pulumi.Input<string>;
}
