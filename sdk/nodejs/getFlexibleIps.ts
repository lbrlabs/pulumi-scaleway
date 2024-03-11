// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Gets information about multiple Flexible IPs.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const fipsByTags = scaleway.getFlexibleIps({
 *     tags: ["a tag"],
 * });
 * const myOffer = scaleway.getBaremetalOffer({
 *     name: "EM-B112X-SSD",
 * });
 * const base = new scaleway.BaremetalServer("base", {
 *     offer: myOffer.then(myOffer => myOffer.offerId),
 *     installConfigAfterward: true,
 * });
 * const first = new scaleway.FlexibleIp("first", {
 *     serverId: base.id,
 *     tags: [
 *         "foo",
 *         "first",
 *     ],
 * });
 * const second = new scaleway.FlexibleIp("second", {
 *     serverId: base.id,
 *     tags: [
 *         "foo",
 *         "second",
 *     ],
 * });
 * const fipsByServerId = scaleway.getFlexibleIpsOutput({
 *     serverIds: [base.id],
 * });
 * ```
 */
export function getFlexibleIps(args?: GetFlexibleIpsArgs, opts?: pulumi.InvokeOptions): Promise<GetFlexibleIpsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:index/getFlexibleIps:getFlexibleIps", {
        "projectId": args.projectId,
        "serverIds": args.serverIds,
        "tags": args.tags,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getFlexibleIps.
 */
export interface GetFlexibleIpsArgs {
    /**
     * (Defaults to provider `projectId`) The ID of the project the IP is in.
     */
    projectId?: string;
    /**
     * List of server IDs used as filter. IPs with these exact server IDs are listed.
     */
    serverIds?: string[];
    /**
     * List of tags used as filter. IPs with these exact tags are listed.
     */
    tags?: string[];
    /**
     * `zone`) The zone in which IPs exist.
     */
    zone?: string;
}

/**
 * A collection of values returned by getFlexibleIps.
 */
export interface GetFlexibleIpsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * List of found flexible IPS
     */
    readonly ips: outputs.GetFlexibleIpsIp[];
    /**
     * (Defaults to provider `organizationId`) The ID of the organization the IP is in.
     */
    readonly organizationId: string;
    /**
     * (Defaults to provider `projectId`) The ID of the project the IP is in.
     */
    readonly projectId: string;
    readonly serverIds?: string[];
    /**
     * The list of tags which are attached to the flexible IP.
     */
    readonly tags?: string[];
    /**
     * (Defaults to provider `zone`) The zone in which the MAC address exist.
     */
    readonly zone: string;
}
/**
 * Gets information about multiple Flexible IPs.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const fipsByTags = scaleway.getFlexibleIps({
 *     tags: ["a tag"],
 * });
 * const myOffer = scaleway.getBaremetalOffer({
 *     name: "EM-B112X-SSD",
 * });
 * const base = new scaleway.BaremetalServer("base", {
 *     offer: myOffer.then(myOffer => myOffer.offerId),
 *     installConfigAfterward: true,
 * });
 * const first = new scaleway.FlexibleIp("first", {
 *     serverId: base.id,
 *     tags: [
 *         "foo",
 *         "first",
 *     ],
 * });
 * const second = new scaleway.FlexibleIp("second", {
 *     serverId: base.id,
 *     tags: [
 *         "foo",
 *         "second",
 *     ],
 * });
 * const fipsByServerId = scaleway.getFlexibleIpsOutput({
 *     serverIds: [base.id],
 * });
 * ```
 */
export function getFlexibleIpsOutput(args?: GetFlexibleIpsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFlexibleIpsResult> {
    return pulumi.output(args).apply((a: any) => getFlexibleIps(a, opts))
}

/**
 * A collection of arguments for invoking getFlexibleIps.
 */
export interface GetFlexibleIpsOutputArgs {
    /**
     * (Defaults to provider `projectId`) The ID of the project the IP is in.
     */
    projectId?: pulumi.Input<string>;
    /**
     * List of server IDs used as filter. IPs with these exact server IDs are listed.
     */
    serverIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of tags used as filter. IPs with these exact tags are listed.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * `zone`) The zone in which IPs exist.
     */
    zone?: pulumi.Input<string>;
}