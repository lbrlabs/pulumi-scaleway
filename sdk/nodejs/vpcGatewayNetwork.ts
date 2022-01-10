// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Creates and manages Scaleway VPC Public Gateway Network.
 * It allows attaching Private Networks to the VPC Public Gateway and your DHCP config
 * For more information, see [the documentation](https://developers.scaleway.com/en/products/vpc-gw/api/v1/#step-3-attach-private-networks-to-the-vpc-public-gateway).
 *
 * ## Example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const pn01 = new scaleway.VpcPrivateNetwork("pn01", {});
 * const gw01 = new scaleway.VpcPublicGatewayIp("gw01", {});
 * const dhcp01 = new scaleway.VpcPublicGatewayDhcp("dhcp01", {subnet: "192.168.1.0/24"});
 * const pg01 = new scaleway.VpcPublicGateway("pg01", {
 *     type: "VPC-GW-S",
 *     ipId: gw01.id,
 * });
 * const main = new scaleway.VpcGatewayNetwork("main", {
 *     gatewayId: pg01.id,
 *     privateNetworkId: pn01.id,
 *     dhcpId: dhcp01.id,
 * });
 * ```
 *
 * ## Import
 *
 * Gateway network can be imported using the `{zone}/{id}`, e.g. bash
 *
 * ```sh
 *  $ pulumi import scaleway:index/vpcGatewayNetwork:VpcGatewayNetwork main fr-par-1/11111111-1111-1111-1111-111111111111
 * ```
 */
export class VpcGatewayNetwork extends pulumi.CustomResource {
    /**
     * Get an existing VpcGatewayNetwork resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpcGatewayNetworkState, opts?: pulumi.CustomResourceOptions): VpcGatewayNetwork {
        return new VpcGatewayNetwork(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/vpcGatewayNetwork:VpcGatewayNetwork';

    /**
     * Returns true if the given object is an instance of VpcGatewayNetwork.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpcGatewayNetwork {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpcGatewayNetwork.__pulumiType;
    }

    /**
     * Remove DHCP config on this network on destroy. It requires DHCP id.
     */
    public readonly cleanupDhcp!: pulumi.Output<boolean | undefined>;
    /**
     * The date and time of the creation of the gateway network.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The ID of the public gateway DHCP config.
     */
    public readonly dhcpId!: pulumi.Output<string | undefined>;
    /**
     * Enable DHCP config on this network. It requires DHCP id.
     */
    public readonly enableDhcp!: pulumi.Output<boolean | undefined>;
    /**
     * Enable masquerade on this network
     */
    public readonly enableMasquerade!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of the public gateway.
     */
    public readonly gatewayId!: pulumi.Output<string>;
    /**
     * The mac address of the creation of the gateway network.
     */
    public /*out*/ readonly macAddress!: pulumi.Output<string>;
    /**
     * The ID of the private network.
     */
    public readonly privateNetworkId!: pulumi.Output<string>;
    /**
     * Enable DHCP config on this network
     */
    public readonly staticAddress!: pulumi.Output<string | undefined>;
    /**
     * The date and time of the last update of the gateway network.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * `zone`) The zone in which the gateway network should be created.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a VpcGatewayNetwork resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpcGatewayNetworkArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpcGatewayNetworkArgs | VpcGatewayNetworkState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VpcGatewayNetworkState | undefined;
            inputs["cleanupDhcp"] = state ? state.cleanupDhcp : undefined;
            inputs["createdAt"] = state ? state.createdAt : undefined;
            inputs["dhcpId"] = state ? state.dhcpId : undefined;
            inputs["enableDhcp"] = state ? state.enableDhcp : undefined;
            inputs["enableMasquerade"] = state ? state.enableMasquerade : undefined;
            inputs["gatewayId"] = state ? state.gatewayId : undefined;
            inputs["macAddress"] = state ? state.macAddress : undefined;
            inputs["privateNetworkId"] = state ? state.privateNetworkId : undefined;
            inputs["staticAddress"] = state ? state.staticAddress : undefined;
            inputs["updatedAt"] = state ? state.updatedAt : undefined;
            inputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as VpcGatewayNetworkArgs | undefined;
            if ((!args || args.gatewayId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'gatewayId'");
            }
            if ((!args || args.privateNetworkId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'privateNetworkId'");
            }
            inputs["cleanupDhcp"] = args ? args.cleanupDhcp : undefined;
            inputs["dhcpId"] = args ? args.dhcpId : undefined;
            inputs["enableDhcp"] = args ? args.enableDhcp : undefined;
            inputs["enableMasquerade"] = args ? args.enableMasquerade : undefined;
            inputs["gatewayId"] = args ? args.gatewayId : undefined;
            inputs["privateNetworkId"] = args ? args.privateNetworkId : undefined;
            inputs["staticAddress"] = args ? args.staticAddress : undefined;
            inputs["zone"] = args ? args.zone : undefined;
            inputs["createdAt"] = undefined /*out*/;
            inputs["macAddress"] = undefined /*out*/;
            inputs["updatedAt"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(VpcGatewayNetwork.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpcGatewayNetwork resources.
 */
export interface VpcGatewayNetworkState {
    /**
     * Remove DHCP config on this network on destroy. It requires DHCP id.
     */
    cleanupDhcp?: pulumi.Input<boolean>;
    /**
     * The date and time of the creation of the gateway network.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The ID of the public gateway DHCP config.
     */
    dhcpId?: pulumi.Input<string>;
    /**
     * Enable DHCP config on this network. It requires DHCP id.
     */
    enableDhcp?: pulumi.Input<boolean>;
    /**
     * Enable masquerade on this network
     */
    enableMasquerade?: pulumi.Input<boolean>;
    /**
     * The ID of the public gateway.
     */
    gatewayId?: pulumi.Input<string>;
    /**
     * The mac address of the creation of the gateway network.
     */
    macAddress?: pulumi.Input<string>;
    /**
     * The ID of the private network.
     */
    privateNetworkId?: pulumi.Input<string>;
    /**
     * Enable DHCP config on this network
     */
    staticAddress?: pulumi.Input<string>;
    /**
     * The date and time of the last update of the gateway network.
     */
    updatedAt?: pulumi.Input<string>;
    /**
     * `zone`) The zone in which the gateway network should be created.
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VpcGatewayNetwork resource.
 */
export interface VpcGatewayNetworkArgs {
    /**
     * Remove DHCP config on this network on destroy. It requires DHCP id.
     */
    cleanupDhcp?: pulumi.Input<boolean>;
    /**
     * The ID of the public gateway DHCP config.
     */
    dhcpId?: pulumi.Input<string>;
    /**
     * Enable DHCP config on this network. It requires DHCP id.
     */
    enableDhcp?: pulumi.Input<boolean>;
    /**
     * Enable masquerade on this network
     */
    enableMasquerade?: pulumi.Input<boolean>;
    /**
     * The ID of the public gateway.
     */
    gatewayId: pulumi.Input<string>;
    /**
     * The ID of the private network.
     */
    privateNetworkId: pulumi.Input<string>;
    /**
     * Enable DHCP config on this network
     */
    staticAddress?: pulumi.Input<string>;
    /**
     * `zone`) The zone in which the gateway network should be created.
     */
    zone?: pulumi.Input<string>;
}
