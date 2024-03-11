// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Gets information about a public gateway.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const main = new scaleway.VpcPublicGateway("main", {
 *     type: "VPC-GW-S",
 *     zone: "nl-ams-1",
 * });
 * const pgTestByName = scaleway.getVpcPublicGatewayOutput({
 *     name: main.name,
 *     zone: "nl-ams-1",
 * });
 * const pgTestById = scaleway.getVpcPublicGatewayOutput({
 *     publicGatewayId: main.id,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getVpcPublicGateway(args?: GetVpcPublicGatewayArgs, opts?: pulumi.InvokeOptions): Promise<GetVpcPublicGatewayResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:index/getVpcPublicGateway:getVpcPublicGateway", {
        "name": args.name,
        "projectId": args.projectId,
        "publicGatewayId": args.publicGatewayId,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getVpcPublicGateway.
 */
export interface GetVpcPublicGatewayArgs {
    /**
     * Exact name of the public gateway.
     */
    name?: string;
    /**
     * The ID of the project the public gateway is associated with.
     */
    projectId?: string;
    publicGatewayId?: string;
    /**
     * `zone`) The zone in which
     * the public gateway should be created.
     */
    zone?: string;
}

/**
 * A collection of values returned by getVpcPublicGateway.
 */
export interface GetVpcPublicGatewayResult {
    readonly bastionEnabled: boolean;
    readonly bastionPort: number;
    readonly createdAt: string;
    readonly enableSmtp: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ipId: string;
    readonly name?: string;
    readonly organizationId: string;
    readonly projectId?: string;
    readonly publicGatewayId?: string;
    readonly status: string;
    readonly tags: string[];
    readonly type: string;
    readonly updatedAt: string;
    readonly upstreamDnsServers: string[];
    readonly zone?: string;
}
/**
 * Gets information about a public gateway.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const main = new scaleway.VpcPublicGateway("main", {
 *     type: "VPC-GW-S",
 *     zone: "nl-ams-1",
 * });
 * const pgTestByName = scaleway.getVpcPublicGatewayOutput({
 *     name: main.name,
 *     zone: "nl-ams-1",
 * });
 * const pgTestById = scaleway.getVpcPublicGatewayOutput({
 *     publicGatewayId: main.id,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getVpcPublicGatewayOutput(args?: GetVpcPublicGatewayOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVpcPublicGatewayResult> {
    return pulumi.output(args).apply((a: any) => getVpcPublicGateway(a, opts))
}

/**
 * A collection of arguments for invoking getVpcPublicGateway.
 */
export interface GetVpcPublicGatewayOutputArgs {
    /**
     * Exact name of the public gateway.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project the public gateway is associated with.
     */
    projectId?: pulumi.Input<string>;
    publicGatewayId?: pulumi.Input<string>;
    /**
     * `zone`) The zone in which
     * the public gateway should be created.
     */
    zone?: pulumi.Input<string>;
}
