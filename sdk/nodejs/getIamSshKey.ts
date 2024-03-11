// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to get SSH key information based on its ID or name.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const myKey = scaleway.getIamSshKey({
 *     sshKeyId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getIamSshKey(args?: GetIamSshKeyArgs, opts?: pulumi.InvokeOptions): Promise<GetIamSshKeyResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:index/getIamSshKey:getIamSshKey", {
        "name": args.name,
        "projectId": args.projectId,
        "sshKeyId": args.sshKeyId,
    }, opts);
}

/**
 * A collection of arguments for invoking getIamSshKey.
 */
export interface GetIamSshKeyArgs {
    /**
     * The SSH key name. Only one of `name` and `sshKeyId` should be specified.
     */
    name?: string;
    /**
     * `projectId`) The ID of the project the SSH
     * key is associated with.
     */
    projectId?: string;
    /**
     * The SSH key id. Only one of `name` and `sshKeyId` should be specified.
     */
    sshKeyId?: string;
}

/**
 * A collection of values returned by getIamSshKey.
 */
export interface GetIamSshKeyResult {
    /**
     * The date and time of the creation of the SSH key.
     */
    readonly createdAt: string;
    /**
     * The SSH key status.
     */
    readonly disabled: boolean;
    readonly fingerprint: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name?: string;
    /**
     * The ID of the organization the SSH key is associated with.
     */
    readonly organizationId: string;
    readonly projectId?: string;
    /**
     * The SSH public key string
     */
    readonly publicKey: string;
    readonly sshKeyId?: string;
    /**
     * The date and time of the last update of the SSH key.
     */
    readonly updatedAt: string;
}
/**
 * Use this data source to get SSH key information based on its ID or name.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const myKey = scaleway.getIamSshKey({
 *     sshKeyId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getIamSshKeyOutput(args?: GetIamSshKeyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIamSshKeyResult> {
    return pulumi.output(args).apply((a: any) => getIamSshKey(a, opts))
}

/**
 * A collection of arguments for invoking getIamSshKey.
 */
export interface GetIamSshKeyOutputArgs {
    /**
     * The SSH key name. Only one of `name` and `sshKeyId` should be specified.
     */
    name?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the SSH
     * key is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The SSH key id. Only one of `name` and `sshKeyId` should be specified.
     */
    sshKeyId?: pulumi.Input<string>;
}
