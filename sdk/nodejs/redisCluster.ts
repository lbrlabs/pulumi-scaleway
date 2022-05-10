// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export class RedisCluster extends pulumi.CustomResource {
    /**
     * Get an existing RedisCluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RedisClusterState, opts?: pulumi.CustomResourceOptions): RedisCluster {
        return new RedisCluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/redisCluster:RedisCluster';

    /**
     * Returns true if the given object is an instance of RedisCluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RedisCluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RedisCluster.__pulumiType;
    }

    /**
     * List of acl rules.
     */
    public readonly acls!: pulumi.Output<outputs.RedisClusterAcl[] | undefined>;
    /**
     * Number of nodes for the cluster.
     */
    public readonly clusterSize!: pulumi.Output<number>;
    /**
     * The date and time of the creation of the Redis cluster
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Name of the redis cluster
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Type of node to use for the cluster
     */
    public readonly nodeType!: pulumi.Output<string>;
    /**
     * Password of the user
     */
    public readonly password!: pulumi.Output<string>;
    /**
     * The project_id you want to attach the resource to
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * Map of settings to define for the cluster.
     */
    public readonly settings!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * List of tags ["tag1", "tag2", ...] attached to a redis cluster
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * Whether or not TLS is enabled.
     */
    public readonly tlsEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The date and time of the last update of the Redis cluster
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * Name of the user created when the cluster is created
     */
    public readonly userName!: pulumi.Output<string>;
    /**
     * Redis version of the cluster
     */
    public readonly version!: pulumi.Output<string>;
    /**
     * The zone you want to attach the resource to
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a RedisCluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RedisClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RedisClusterArgs | RedisClusterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RedisClusterState | undefined;
            resourceInputs["acls"] = state ? state.acls : undefined;
            resourceInputs["clusterSize"] = state ? state.clusterSize : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nodeType"] = state ? state.nodeType : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["settings"] = state ? state.settings : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tlsEnabled"] = state ? state.tlsEnabled : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
            resourceInputs["userName"] = state ? state.userName : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as RedisClusterArgs | undefined;
            if ((!args || args.nodeType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nodeType'");
            }
            if ((!args || args.password === undefined) && !opts.urn) {
                throw new Error("Missing required property 'password'");
            }
            if ((!args || args.userName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userName'");
            }
            if ((!args || args.version === undefined) && !opts.urn) {
                throw new Error("Missing required property 'version'");
            }
            resourceInputs["acls"] = args ? args.acls : undefined;
            resourceInputs["clusterSize"] = args ? args.clusterSize : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["nodeType"] = args ? args.nodeType : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["settings"] = args ? args.settings : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["tlsEnabled"] = args ? args.tlsEnabled : undefined;
            resourceInputs["userName"] = args ? args.userName : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RedisCluster.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RedisCluster resources.
 */
export interface RedisClusterState {
    /**
     * List of acl rules.
     */
    acls?: pulumi.Input<pulumi.Input<inputs.RedisClusterAcl>[]>;
    /**
     * Number of nodes for the cluster.
     */
    clusterSize?: pulumi.Input<number>;
    /**
     * The date and time of the creation of the Redis cluster
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Name of the redis cluster
     */
    name?: pulumi.Input<string>;
    /**
     * Type of node to use for the cluster
     */
    nodeType?: pulumi.Input<string>;
    /**
     * Password of the user
     */
    password?: pulumi.Input<string>;
    /**
     * The project_id you want to attach the resource to
     */
    projectId?: pulumi.Input<string>;
    /**
     * Map of settings to define for the cluster.
     */
    settings?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * List of tags ["tag1", "tag2", ...] attached to a redis cluster
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether or not TLS is enabled.
     */
    tlsEnabled?: pulumi.Input<boolean>;
    /**
     * The date and time of the last update of the Redis cluster
     */
    updatedAt?: pulumi.Input<string>;
    /**
     * Name of the user created when the cluster is created
     */
    userName?: pulumi.Input<string>;
    /**
     * Redis version of the cluster
     */
    version?: pulumi.Input<string>;
    /**
     * The zone you want to attach the resource to
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RedisCluster resource.
 */
export interface RedisClusterArgs {
    /**
     * List of acl rules.
     */
    acls?: pulumi.Input<pulumi.Input<inputs.RedisClusterAcl>[]>;
    /**
     * Number of nodes for the cluster.
     */
    clusterSize?: pulumi.Input<number>;
    /**
     * Name of the redis cluster
     */
    name?: pulumi.Input<string>;
    /**
     * Type of node to use for the cluster
     */
    nodeType: pulumi.Input<string>;
    /**
     * Password of the user
     */
    password: pulumi.Input<string>;
    /**
     * The project_id you want to attach the resource to
     */
    projectId?: pulumi.Input<string>;
    /**
     * Map of settings to define for the cluster.
     */
    settings?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * List of tags ["tag1", "tag2", ...] attached to a redis cluster
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether or not TLS is enabled.
     */
    tlsEnabled?: pulumi.Input<boolean>;
    /**
     * Name of the user created when the cluster is created
     */
    userName: pulumi.Input<string>;
    /**
     * Redis version of the cluster
     */
    version: pulumi.Input<string>;
    /**
     * The zone you want to attach the resource to
     */
    zone?: pulumi.Input<string>;
}
