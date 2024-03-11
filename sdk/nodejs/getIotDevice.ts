// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Gets information about an IOT Device.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const myDevice = scaleway.getIotDevice({
 *     deviceId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getIotDevice(args?: GetIotDeviceArgs, opts?: pulumi.InvokeOptions): Promise<GetIotDeviceResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:index/getIotDevice:getIotDevice", {
        "deviceId": args.deviceId,
        "hubId": args.hubId,
        "name": args.name,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getIotDevice.
 */
export interface GetIotDeviceArgs {
    /**
     * The device ID.
     * Only one of the `name` and `deviceId` should be specified.
     */
    deviceId?: string;
    /**
     * The hub ID.
     */
    hubId?: string;
    /**
     * The name of the Hub.
     * Only one of the `name` and `deviceId` should be specified.
     */
    name?: string;
    /**
     * `region`) The region in which the hub exists.
     */
    region?: string;
}

/**
 * A collection of values returned by getIotDevice.
 */
export interface GetIotDeviceResult {
    readonly allowInsecure: boolean;
    readonly allowMultipleConnections: boolean;
    readonly certificates: outputs.GetIotDeviceCertificate[];
    readonly createdAt: string;
    readonly description: string;
    readonly deviceId?: string;
    readonly hubId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly isConnected: boolean;
    readonly lastActivityAt: string;
    readonly messageFilters: outputs.GetIotDeviceMessageFilter[];
    readonly name?: string;
    readonly region?: string;
    readonly status: string;
    readonly updatedAt: string;
}
/**
 * Gets information about an IOT Device.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const myDevice = scaleway.getIotDevice({
 *     deviceId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getIotDeviceOutput(args?: GetIotDeviceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIotDeviceResult> {
    return pulumi.output(args).apply((a: any) => getIotDevice(a, opts))
}

/**
 * A collection of arguments for invoking getIotDevice.
 */
export interface GetIotDeviceOutputArgs {
    /**
     * The device ID.
     * Only one of the `name` and `deviceId` should be specified.
     */
    deviceId?: pulumi.Input<string>;
    /**
     * The hub ID.
     */
    hubId?: pulumi.Input<string>;
    /**
     * The name of the Hub.
     * Only one of the `name` and `deviceId` should be specified.
     */
    name?: pulumi.Input<string>;
    /**
     * `region`) The region in which the hub exists.
     */
    region?: pulumi.Input<string>;
}
