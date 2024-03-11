// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Gets information about the RDB instance network Access Control List.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const myAcl = scaleway.getDatabaseAcl({
 *     instanceId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getDatabaseAcl(args: GetDatabaseAclArgs, opts?: pulumi.InvokeOptions): Promise<GetDatabaseAclResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:index/getDatabaseAcl:getDatabaseAcl", {
        "instanceId": args.instanceId,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getDatabaseAcl.
 */
export interface GetDatabaseAclArgs {
    /**
     * The RDB instance ID.
     */
    instanceId: string;
    /**
     * `region`) The region in which the Database Instance should be created.
     */
    region?: string;
}

/**
 * A collection of values returned by getDatabaseAcl.
 */
export interface GetDatabaseAclResult {
    /**
     * A list of ACLs rules (structure is described below)
     */
    readonly aclRules: outputs.GetDatabaseAclAclRule[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    readonly region?: string;
}
/**
 * Gets information about the RDB instance network Access Control List.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const myAcl = scaleway.getDatabaseAcl({
 *     instanceId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getDatabaseAclOutput(args: GetDatabaseAclOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDatabaseAclResult> {
    return pulumi.output(args).apply((a: any) => getDatabaseAcl(a, opts))
}

/**
 * A collection of arguments for invoking getDatabaseAcl.
 */
export interface GetDatabaseAclOutputArgs {
    /**
     * The RDB instance ID.
     */
    instanceId: pulumi.Input<string>;
    /**
     * `region`) The region in which the Database Instance should be created.
     */
    region?: pulumi.Input<string>;
}
