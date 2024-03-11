// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Creates and manages Scaleway Function Domain bindings.
 * For more information see [the documentation](https://developers.scaleway.com/en/products/functions/api/).
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
 * const mainFunctionNamespace = new scaleway.FunctionNamespace("mainFunctionNamespace", {});
 * const mainFunction = new scaleway.Function("mainFunction", {
 *     namespaceId: mainFunctionNamespace.id,
 *     runtime: "go118",
 *     privacy: "private",
 *     handler: "Handle",
 *     zipFile: "testfixture/gofunction.zip",
 *     deploy: true,
 * });
 * const mainFunctionDomain = new scaleway.FunctionDomain("mainFunctionDomain", {
 *     functionId: mainFunction.id,
 *     hostname: "example.com",
 * }, {
 *     dependsOn: [mainFunction],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Domain can be imported using the `{region}/{id}`, e.g.
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:index/functionDomain:FunctionDomain main fr-par/11111111-1111-1111-1111-111111111111
 * ```
 */
export class FunctionDomain extends pulumi.CustomResource {
    /**
     * Get an existing FunctionDomain resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FunctionDomainState, opts?: pulumi.CustomResourceOptions): FunctionDomain {
        return new FunctionDomain(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/functionDomain:FunctionDomain';

    /**
     * Returns true if the given object is an instance of FunctionDomain.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FunctionDomain {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FunctionDomain.__pulumiType;
    }

    /**
     * The ID of the function you want to create a domain with.
     */
    public readonly functionId!: pulumi.Output<string>;
    /**
     * The hostname that should resolve to your function id native domain.
     * You should use a CNAME domain record that point to your native function `domainName` for it.
     *
     * > **Important** Updates to `functionId` or `hostname` will recreate the domain.
     */
    public readonly hostname!: pulumi.Output<string>;
    /**
     * (Defaults to provider `region`) The region in where the domain was created.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The URL that triggers the function
     */
    public /*out*/ readonly url!: pulumi.Output<string>;

    /**
     * Create a FunctionDomain resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FunctionDomainArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FunctionDomainArgs | FunctionDomainState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FunctionDomainState | undefined;
            resourceInputs["functionId"] = state ? state.functionId : undefined;
            resourceInputs["hostname"] = state ? state.hostname : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as FunctionDomainArgs | undefined;
            if ((!args || args.functionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'functionId'");
            }
            if ((!args || args.hostname === undefined) && !opts.urn) {
                throw new Error("Missing required property 'hostname'");
            }
            resourceInputs["functionId"] = args ? args.functionId : undefined;
            resourceInputs["hostname"] = args ? args.hostname : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["url"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FunctionDomain.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FunctionDomain resources.
 */
export interface FunctionDomainState {
    /**
     * The ID of the function you want to create a domain with.
     */
    functionId?: pulumi.Input<string>;
    /**
     * The hostname that should resolve to your function id native domain.
     * You should use a CNAME domain record that point to your native function `domainName` for it.
     *
     * > **Important** Updates to `functionId` or `hostname` will recreate the domain.
     */
    hostname?: pulumi.Input<string>;
    /**
     * (Defaults to provider `region`) The region in where the domain was created.
     */
    region?: pulumi.Input<string>;
    /**
     * The URL that triggers the function
     */
    url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FunctionDomain resource.
 */
export interface FunctionDomainArgs {
    /**
     * The ID of the function you want to create a domain with.
     */
    functionId: pulumi.Input<string>;
    /**
     * The hostname that should resolve to your function id native domain.
     * You should use a CNAME domain record that point to your native function `domainName` for it.
     *
     * > **Important** Updates to `functionId` or `hostname` will recreate the domain.
     */
    hostname: pulumi.Input<string>;
    /**
     * (Defaults to provider `region`) The region in where the domain was created.
     */
    region?: pulumi.Input<string>;
}
