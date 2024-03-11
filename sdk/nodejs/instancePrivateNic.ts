// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Creates and manages Scaleway Instance Private NICs. For more information, see
 * [the documentation](https://developers.scaleway.com/en/products/instance/api/#private-nics-a42eea).
 *
 * ## Examples
 *
 * ### Basic
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const pnic01 = new scaleway.InstancePrivateNic("pnic01", {
 *     privateNetworkId: "fr-par-1/aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
 *     serverId: "fr-par-1/11111111-1111-1111-1111-111111111111",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ### With zone
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const pn01 = new scaleway.VpcPrivateNetwork("pn01", {zone: "fr-par-2"});
 * const base = new scaleway.InstanceServer("base", {
 *     image: "ubuntu_jammy",
 *     type: "DEV1-S",
 *     zone: pn01.zone,
 * });
 * const pnic01 = new scaleway.InstancePrivateNic("pnic01", {
 *     serverId: base.id,
 *     privateNetworkId: pn01.id,
 *     zone: pn01.zone,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Private NICs can be imported using the `{zone}/{server_id}/{private_nic_id}`, e.g.
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:index/instancePrivateNic:InstancePrivateNic pnic01 fr-par-1/11111111-1111-1111-1111-111111111111/22222222-2222-2222-2222-222222222222
 * ```
 */
export class InstancePrivateNic extends pulumi.CustomResource {
    /**
     * Get an existing InstancePrivateNic resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstancePrivateNicState, opts?: pulumi.CustomResourceOptions): InstancePrivateNic {
        return new InstancePrivateNic(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/instancePrivateNic:InstancePrivateNic';

    /**
     * Returns true if the given object is an instance of InstancePrivateNic.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InstancePrivateNic {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InstancePrivateNic.__pulumiType;
    }

    /**
     * IPAM ip list, should be for internal use only
     */
    public readonly ipIds!: pulumi.Output<string[] | undefined>;
    /**
     * The MAC address of the private NIC.
     */
    public /*out*/ readonly macAddress!: pulumi.Output<string>;
    /**
     * The ID of the private network attached to.
     */
    public readonly privateNetworkId!: pulumi.Output<string>;
    /**
     * The ID of the server associated with.
     */
    public readonly serverId!: pulumi.Output<string>;
    /**
     * The tags associated with the private NIC.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * `zone`) The zone in which the server must be created.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a InstancePrivateNic resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstancePrivateNicArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstancePrivateNicArgs | InstancePrivateNicState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstancePrivateNicState | undefined;
            resourceInputs["ipIds"] = state ? state.ipIds : undefined;
            resourceInputs["macAddress"] = state ? state.macAddress : undefined;
            resourceInputs["privateNetworkId"] = state ? state.privateNetworkId : undefined;
            resourceInputs["serverId"] = state ? state.serverId : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as InstancePrivateNicArgs | undefined;
            if ((!args || args.privateNetworkId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'privateNetworkId'");
            }
            if ((!args || args.serverId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serverId'");
            }
            resourceInputs["ipIds"] = args ? args.ipIds : undefined;
            resourceInputs["privateNetworkId"] = args ? args.privateNetworkId : undefined;
            resourceInputs["serverId"] = args ? args.serverId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
            resourceInputs["macAddress"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(InstancePrivateNic.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InstancePrivateNic resources.
 */
export interface InstancePrivateNicState {
    /**
     * IPAM ip list, should be for internal use only
     */
    ipIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The MAC address of the private NIC.
     */
    macAddress?: pulumi.Input<string>;
    /**
     * The ID of the private network attached to.
     */
    privateNetworkId?: pulumi.Input<string>;
    /**
     * The ID of the server associated with.
     */
    serverId?: pulumi.Input<string>;
    /**
     * The tags associated with the private NIC.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * `zone`) The zone in which the server must be created.
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a InstancePrivateNic resource.
 */
export interface InstancePrivateNicArgs {
    /**
     * IPAM ip list, should be for internal use only
     */
    ipIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the private network attached to.
     */
    privateNetworkId: pulumi.Input<string>;
    /**
     * The ID of the server associated with.
     */
    serverId: pulumi.Input<string>;
    /**
     * The tags associated with the private NIC.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * `zone`) The zone in which the server must be created.
     */
    zone?: pulumi.Input<string>;
}
