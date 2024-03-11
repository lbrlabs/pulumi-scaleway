// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Gets information about a private network.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const myName = scaleway.getVpcPrivateNetwork({
 *     name: "foobar",
 * });
 * const myNameAndVpcId = scaleway.getVpcPrivateNetwork({
 *     name: "foobar",
 *     vpcId: "11111111-1111-1111-1111-111111111111",
 * });
 * const myId = scaleway.getVpcPrivateNetwork({
 *     privateNetworkId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getVpcPrivateNetwork(args?: GetVpcPrivateNetworkArgs, opts?: pulumi.InvokeOptions): Promise<GetVpcPrivateNetworkResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:index/getVpcPrivateNetwork:getVpcPrivateNetwork", {
        "name": args.name,
        "privateNetworkId": args.privateNetworkId,
        "projectId": args.projectId,
        "vpcId": args.vpcId,
    }, opts);
}

/**
 * A collection of arguments for invoking getVpcPrivateNetwork.
 */
export interface GetVpcPrivateNetworkArgs {
    /**
     * Name of the private network. Cannot be used with `privateNetworkId`.
     */
    name?: string;
    /**
     * ID of the private network. Cannot be used with `name` and `vpcId`.
     */
    privateNetworkId?: string;
    /**
     * The ID of the project the private network is associated with.
     */
    projectId?: string;
    /**
     * ID of the VPC in which the private network is. Cannot be used with `privateNetworkId`.
     */
    vpcId?: string;
}

/**
 * A collection of values returned by getVpcPrivateNetwork.
 */
export interface GetVpcPrivateNetworkResult {
    readonly createdAt: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The IPv4 subnet associated with the private network.
     */
    readonly ipv4Subnets: outputs.GetVpcPrivateNetworkIpv4Subnet[];
    /**
     * The IPv6 subnets associated with the private network.
     */
    readonly ipv6Subnets: outputs.GetVpcPrivateNetworkIpv6Subnet[];
    readonly isRegional: boolean;
    readonly name?: string;
    readonly organizationId: string;
    readonly privateNetworkId?: string;
    readonly projectId?: string;
    readonly region: string;
    readonly tags: string[];
    readonly updatedAt: string;
    readonly vpcId?: string;
    readonly zone: string;
}
/**
 * Gets information about a private network.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const myName = scaleway.getVpcPrivateNetwork({
 *     name: "foobar",
 * });
 * const myNameAndVpcId = scaleway.getVpcPrivateNetwork({
 *     name: "foobar",
 *     vpcId: "11111111-1111-1111-1111-111111111111",
 * });
 * const myId = scaleway.getVpcPrivateNetwork({
 *     privateNetworkId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getVpcPrivateNetworkOutput(args?: GetVpcPrivateNetworkOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVpcPrivateNetworkResult> {
    return pulumi.output(args).apply((a: any) => getVpcPrivateNetwork(a, opts))
}

/**
 * A collection of arguments for invoking getVpcPrivateNetwork.
 */
export interface GetVpcPrivateNetworkOutputArgs {
    /**
     * Name of the private network. Cannot be used with `privateNetworkId`.
     */
    name?: pulumi.Input<string>;
    /**
     * ID of the private network. Cannot be used with `name` and `vpcId`.
     */
    privateNetworkId?: pulumi.Input<string>;
    /**
     * The ID of the project the private network is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * ID of the VPC in which the private network is. Cannot be used with `privateNetworkId`.
     */
    vpcId?: pulumi.Input<string>;
}
