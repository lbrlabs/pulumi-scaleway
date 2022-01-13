// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Gets information about a Load Balancer.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * // Get info by name
 * const byName = pulumi.output(scaleway.getLoadbalancer({
 *     name: "foobar",
 * }));
 * // Get info by ID
 * const byId = pulumi.output(scaleway.getLoadbalancer({
 *     lbId: "11111111-1111-1111-1111-111111111111",
 * }));
 * ```
 */
export function getLoadbalancer(args?: GetLoadbalancerArgs, opts?: pulumi.InvokeOptions): Promise<GetLoadbalancerResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("scaleway:index/getLoadbalancer:getLoadbalancer", {
        "lbId": args.lbId,
        "name": args.name,
        "releaseIp": args.releaseIp,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getLoadbalancer.
 */
export interface GetLoadbalancerArgs {
    /**
     * The ID.
     * Only one of `ipAddress` and `lbId` should be specified.
     */
    lbId?: string;
    /**
     * The IP address.
     * Only one of `name` and `lbId` should be specified.
     */
    name?: string;
    releaseIp?: boolean;
    /**
     * `region`) The region in which the LB exists.
     */
    zone?: string;
}

/**
 * A collection of values returned by getLoadbalancer.
 */
export interface GetLoadbalancerResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The load-balancer public IP Address
     */
    readonly ipAddress: string;
    readonly ipId: string;
    readonly lbId?: string;
    readonly name?: string;
    /**
     * The organization ID the load-balancer is associated with.
     */
    readonly organizationId: string;
    readonly privateNetworks: outputs.GetLoadbalancerPrivateNetwork[];
    readonly projectId: string;
    readonly region: string;
    readonly releaseIp?: boolean;
    /**
     * (Optional) The tags associated with the load-balancers.
     */
    readonly tags: string[];
    /**
     * (Required) The type of the load-balancer.
     */
    readonly type: string;
    readonly zone?: string;
}

export function getLoadbalancerOutput(args?: GetLoadbalancerOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLoadbalancerResult> {
    return pulumi.output(args).apply(a => getLoadbalancer(a, opts))
}

/**
 * A collection of arguments for invoking getLoadbalancer.
 */
export interface GetLoadbalancerOutputArgs {
    /**
     * The ID.
     * Only one of `ipAddress` and `lbId` should be specified.
     */
    lbId?: pulumi.Input<string>;
    /**
     * The IP address.
     * Only one of `name` and `lbId` should be specified.
     */
    name?: pulumi.Input<string>;
    releaseIp?: pulumi.Input<boolean>;
    /**
     * `region`) The region in which the LB exists.
     */
    zone?: pulumi.Input<string>;
}
