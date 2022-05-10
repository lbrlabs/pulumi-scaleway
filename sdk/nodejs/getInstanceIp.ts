// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getInstanceIp(args?: GetInstanceIpArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceIpResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("scaleway:index/getInstanceIp:getInstanceIp", {
        "address": args.address,
        "id": args.id,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstanceIp.
 */
export interface GetInstanceIpArgs {
    address?: string;
    id?: string;
}

/**
 * A collection of values returned by getInstanceIp.
 */
export interface GetInstanceIpResult {
    readonly address?: string;
    readonly id?: string;
    readonly organizationId: string;
    readonly projectId: string;
    readonly reverse: string;
    readonly serverId: string;
    readonly tags: string[];
    readonly zone: string;
}

export function getInstanceIpOutput(args?: GetInstanceIpOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstanceIpResult> {
    return pulumi.output(args).apply(a => getInstanceIp(a, opts))
}

/**
 * A collection of arguments for invoking getInstanceIp.
 */
export interface GetInstanceIpOutputArgs {
    address?: pulumi.Input<string>;
    id?: pulumi.Input<string>;
}
