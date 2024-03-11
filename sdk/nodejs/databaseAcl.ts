// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Creates and manages Scaleway Database instance authorized IPs.
 * For more information, see [the documentation](https://developers.scaleway.com/en/products/rdb/api/#acl-rules-allowed-ips).
 *
 * ## Example Usage
 *
 * ### Basic
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const main = new scaleway.DatabaseAcl("main", {
 *     instanceId: scaleway_rdb_instance.main.id,
 *     aclRules: [{
 *         ip: "1.2.3.4/32",
 *         description: "foo",
 *     }],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Database Instance can be imported using the `{region}/{id}`, e.g.
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:index/databaseAcl:DatabaseAcl acl01 fr-par/11111111-1111-1111-1111-111111111111
 * ```
 */
export class DatabaseAcl extends pulumi.CustomResource {
    /**
     * Get an existing DatabaseAcl resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatabaseAclState, opts?: pulumi.CustomResourceOptions): DatabaseAcl {
        return new DatabaseAcl(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/databaseAcl:DatabaseAcl';

    /**
     * Returns true if the given object is an instance of DatabaseAcl.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DatabaseAcl {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DatabaseAcl.__pulumiType;
    }

    /**
     * A list of ACLs (structure is described below)
     */
    public readonly aclRules!: pulumi.Output<outputs.DatabaseAclAclRule[]>;
    /**
     * UUID of the rdb instance.
     *
     * > **Important:** Updates to `instanceId` will recreate the Database ACL.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * `region`) The region in which the Database Instance should be created.
     */
    public readonly region!: pulumi.Output<string>;

    /**
     * Create a DatabaseAcl resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatabaseAclArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatabaseAclArgs | DatabaseAclState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DatabaseAclState | undefined;
            resourceInputs["aclRules"] = state ? state.aclRules : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as DatabaseAclArgs | undefined;
            if ((!args || args.aclRules === undefined) && !opts.urn) {
                throw new Error("Missing required property 'aclRules'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["aclRules"] = args ? args.aclRules : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DatabaseAcl.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DatabaseAcl resources.
 */
export interface DatabaseAclState {
    /**
     * A list of ACLs (structure is described below)
     */
    aclRules?: pulumi.Input<pulumi.Input<inputs.DatabaseAclAclRule>[]>;
    /**
     * UUID of the rdb instance.
     *
     * > **Important:** Updates to `instanceId` will recreate the Database ACL.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * `region`) The region in which the Database Instance should be created.
     */
    region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DatabaseAcl resource.
 */
export interface DatabaseAclArgs {
    /**
     * A list of ACLs (structure is described below)
     */
    aclRules: pulumi.Input<pulumi.Input<inputs.DatabaseAclAclRule>[]>;
    /**
     * UUID of the rdb instance.
     *
     * > **Important:** Updates to `instanceId` will recreate the Database ACL.
     */
    instanceId: pulumi.Input<string>;
    /**
     * `region`) The region in which the Database Instance should be created.
     */
    region?: pulumi.Input<string>;
}
