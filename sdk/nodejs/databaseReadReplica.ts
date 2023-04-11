// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Creates and manages Scaleway Database read replicas.
 * For more information, see [the documentation](https://developers.scaleway.com/en/products/rdb/api).
 *
 * ## Examples
 *
 * ### Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@lbrlabs/pulumi-scaleway";
 *
 * const instance = new scaleway.DatabaseInstance("instance", {
 *     nodeType: "db-dev-s",
 *     engine: "PostgreSQL-14",
 *     isHaCluster: false,
 *     disableBackup: true,
 *     userName: "my_initial_user",
 *     password: "thiZ_is_v&ry_s3cret",
 *     tags: [
 *         "terraform-test",
 *         "scaleway_rdb_read_replica",
 *         "minimal",
 *     ],
 * });
 * const replica = new scaleway.DatabaseReadReplica("replica", {
 *     instanceId: instance.id,
 *     directAccess: {},
 * });
 * ```
 *
 * ### Private network
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@lbrlabs/pulumi-scaleway";
 *
 * const instance = new scaleway.DatabaseInstance("instance", {
 *     nodeType: "db-dev-s",
 *     engine: "PostgreSQL-14",
 *     isHaCluster: false,
 *     disableBackup: true,
 *     userName: "my_initial_user",
 *     password: "thiZ_is_v&ry_s3cret",
 * });
 * const pn = new scaleway.VpcPrivateNetwork("pn", {});
 * const replica = new scaleway.DatabaseReadReplica("replica", {
 *     instanceId: instance.id,
 *     privateNetwork: {
 *         privateNetworkId: pn.id,
 *         serviceIp: "192.168.1.254/24",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Database Read replica can be imported using the `{region}/{id}`, e.g. bash
 *
 * ```sh
 *  $ pulumi import scaleway:index/databaseReadReplica:DatabaseReadReplica rr fr-par/11111111-1111-1111-1111-111111111111
 * ```
 */
export class DatabaseReadReplica extends pulumi.CustomResource {
    /**
     * Get an existing DatabaseReadReplica resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatabaseReadReplicaState, opts?: pulumi.CustomResourceOptions): DatabaseReadReplica {
        return new DatabaseReadReplica(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/databaseReadReplica:DatabaseReadReplica';

    /**
     * Returns true if the given object is an instance of DatabaseReadReplica.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DatabaseReadReplica {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DatabaseReadReplica.__pulumiType;
    }

    /**
     * Creates a direct access endpoint to rdb replica.
     */
    public readonly directAccess!: pulumi.Output<outputs.DatabaseReadReplicaDirectAccess | undefined>;
    /**
     * UUID of the rdb instance.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Create an endpoint in a private network.
     */
    public readonly privateNetwork!: pulumi.Output<outputs.DatabaseReadReplicaPrivateNetwork | undefined>;
    /**
     * `region`) The region
     * in which the Database read replica should be created.
     */
    public readonly region!: pulumi.Output<string>;

    /**
     * Create a DatabaseReadReplica resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatabaseReadReplicaArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatabaseReadReplicaArgs | DatabaseReadReplicaState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DatabaseReadReplicaState | undefined;
            resourceInputs["directAccess"] = state ? state.directAccess : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["privateNetwork"] = state ? state.privateNetwork : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as DatabaseReadReplicaArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["directAccess"] = args ? args.directAccess : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["privateNetwork"] = args ? args.privateNetwork : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DatabaseReadReplica.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DatabaseReadReplica resources.
 */
export interface DatabaseReadReplicaState {
    /**
     * Creates a direct access endpoint to rdb replica.
     */
    directAccess?: pulumi.Input<inputs.DatabaseReadReplicaDirectAccess>;
    /**
     * UUID of the rdb instance.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Create an endpoint in a private network.
     */
    privateNetwork?: pulumi.Input<inputs.DatabaseReadReplicaPrivateNetwork>;
    /**
     * `region`) The region
     * in which the Database read replica should be created.
     */
    region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DatabaseReadReplica resource.
 */
export interface DatabaseReadReplicaArgs {
    /**
     * Creates a direct access endpoint to rdb replica.
     */
    directAccess?: pulumi.Input<inputs.DatabaseReadReplicaDirectAccess>;
    /**
     * UUID of the rdb instance.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Create an endpoint in a private network.
     */
    privateNetwork?: pulumi.Input<inputs.DatabaseReadReplicaPrivateNetwork>;
    /**
     * `region`) The region
     * in which the Database read replica should be created.
     */
    region?: pulumi.Input<string>;
}
