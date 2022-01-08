// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Create and manage Scaleway RDB database privilege.
 * For more information, see [the documentation](https://developers.scaleway.com/en/products/rdb/api).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const priv = new scaleway.DatabasePrivilege("priv", {
 *     instanceId: scaleway_rdb_instance.rdb.id,
 *     userName: "my-db-user",
 *     databaseName: "my-db-name",
 *     permission: "all",
 * });
 * ```
 */
export class DatabasePrivilege extends pulumi.CustomResource {
    /**
     * Get an existing DatabasePrivilege resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatabasePrivilegeState, opts?: pulumi.CustomResourceOptions): DatabasePrivilege {
        return new DatabasePrivilege(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/databasePrivilege:DatabasePrivilege';

    /**
     * Returns true if the given object is an instance of DatabasePrivilege.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DatabasePrivilege {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DatabasePrivilege.__pulumiType;
    }

    /**
     * Name of the database (e.g. `my-db-name`).
     */
    public readonly databaseName!: pulumi.Output<string>;
    /**
     * UUID of the instance where to create the database.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Permission to set. Valid values are `readonly`, `readwrite`, `all`, `custom` and `none`.
     */
    public readonly permission!: pulumi.Output<string>;
    /**
     * Name of the user (e.g. `my-db-user`).
     */
    public readonly userName!: pulumi.Output<string>;

    /**
     * Create a DatabasePrivilege resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatabasePrivilegeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatabasePrivilegeArgs | DatabasePrivilegeState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DatabasePrivilegeState | undefined;
            resourceInputs["databaseName"] = state ? state.databaseName : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["permission"] = state ? state.permission : undefined;
            resourceInputs["userName"] = state ? state.userName : undefined;
        } else {
            const args = argsOrState as DatabasePrivilegeArgs | undefined;
            if ((!args || args.databaseName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'databaseName'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.permission === undefined) && !opts.urn) {
                throw new Error("Missing required property 'permission'");
            }
            if ((!args || args.userName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userName'");
            }
            resourceInputs["databaseName"] = args ? args.databaseName : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["permission"] = args ? args.permission : undefined;
            resourceInputs["userName"] = args ? args.userName : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(DatabasePrivilege.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DatabasePrivilege resources.
 */
export interface DatabasePrivilegeState {
    /**
     * Name of the database (e.g. `my-db-name`).
     */
    databaseName?: pulumi.Input<string>;
    /**
     * UUID of the instance where to create the database.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Permission to set. Valid values are `readonly`, `readwrite`, `all`, `custom` and `none`.
     */
    permission?: pulumi.Input<string>;
    /**
     * Name of the user (e.g. `my-db-user`).
     */
    userName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DatabasePrivilege resource.
 */
export interface DatabasePrivilegeArgs {
    /**
     * Name of the database (e.g. `my-db-name`).
     */
    databaseName: pulumi.Input<string>;
    /**
     * UUID of the instance where to create the database.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Permission to set. Valid values are `readonly`, `readwrite`, `all`, `custom` and `none`.
     */
    permission: pulumi.Input<string>;
    /**
     * Name of the user (e.g. `my-db-user`).
     */
    userName: pulumi.Input<string>;
}
