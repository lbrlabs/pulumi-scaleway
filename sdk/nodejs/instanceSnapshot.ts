// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Creates and manages Scaleway Compute Snapshots.
 * For more information, see [the documentation](https://developers.scaleway.com/en/products/instance/api/#snapshots-756fae).
 *
 * ## Example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const main = new scaleway.InstanceSnapshot("main", {
 *     volumeId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 *
 * ## Import
 *
 * Snapshots can be imported using the `{zone}/{id}`, e.g. bash
 *
 * ```sh
 *  $ pulumi import scaleway:index/instanceSnapshot:InstanceSnapshot main fr-par-1/11111111-1111-1111-1111-111111111111
 * ```
 */
export class InstanceSnapshot extends pulumi.CustomResource {
    /**
     * Get an existing InstanceSnapshot resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceSnapshotState, opts?: pulumi.CustomResourceOptions): InstanceSnapshot {
        return new InstanceSnapshot(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/instanceSnapshot:InstanceSnapshot';

    /**
     * Returns true if the given object is an instance of InstanceSnapshot.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InstanceSnapshot {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InstanceSnapshot.__pulumiType;
    }

    /**
     * The snapshot creation time.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The name of the snapshot. If not provided it will be randomly generated.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The organization ID the snapshot is associated with.
     */
    public /*out*/ readonly organizationId!: pulumi.Output<string>;
    /**
     * `projectId`) The ID of the project the snapshot is associated with.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * (Optional) The size of the snapshot.
     */
    public /*out*/ readonly sizeInGb!: pulumi.Output<number>;
    /**
     * The type of the snapshot. The possible values are: `bSsd` (Block SSD), `lSsd` (Local SSD).
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The ID of the volume to take a snapshot from.
     */
    public readonly volumeId!: pulumi.Output<string>;
    /**
     * `zone`) The zone in which the snapshot should be created.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a InstanceSnapshot resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceSnapshotArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceSnapshotArgs | InstanceSnapshotState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstanceSnapshotState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["organizationId"] = state ? state.organizationId : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["sizeInGb"] = state ? state.sizeInGb : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["volumeId"] = state ? state.volumeId : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as InstanceSnapshotArgs | undefined;
            if ((!args || args.volumeId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'volumeId'");
            }
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["volumeId"] = args ? args.volumeId : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["organizationId"] = undefined /*out*/;
            resourceInputs["sizeInGb"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(InstanceSnapshot.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InstanceSnapshot resources.
 */
export interface InstanceSnapshotState {
    /**
     * The snapshot creation time.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The name of the snapshot. If not provided it will be randomly generated.
     */
    name?: pulumi.Input<string>;
    /**
     * The organization ID the snapshot is associated with.
     */
    organizationId?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the snapshot is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * (Optional) The size of the snapshot.
     */
    sizeInGb?: pulumi.Input<number>;
    /**
     * The type of the snapshot. The possible values are: `bSsd` (Block SSD), `lSsd` (Local SSD).
     */
    type?: pulumi.Input<string>;
    /**
     * The ID of the volume to take a snapshot from.
     */
    volumeId?: pulumi.Input<string>;
    /**
     * `zone`) The zone in which the snapshot should be created.
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a InstanceSnapshot resource.
 */
export interface InstanceSnapshotArgs {
    /**
     * The name of the snapshot. If not provided it will be randomly generated.
     */
    name?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the snapshot is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The ID of the volume to take a snapshot from.
     */
    volumeId: pulumi.Input<string>;
    /**
     * `zone`) The zone in which the snapshot should be created.
     */
    zone?: pulumi.Input<string>;
}
