// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Creates and manages Scaleway IAM Policies. For more information, see [the documentation](https://developers.scaleway.com/en/products/iam/api/v1alpha1/#policies-54b8a7).
 *
 * ## Example Usage
 *
 * ### Create a policy for an organization's project
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const default = scaleway.getAccountProject({
 *     name: "default",
 * });
 * const app = new scaleway.IamApplication("app", {});
 * const objectReadOnly = new scaleway.IamPolicy("objectReadOnly", {
 *     description: "gives app readonly access to object storage in project",
 *     applicationId: app.id,
 *     rules: [{
 *         projectIds: [_default.then(_default => _default.id)],
 *         permissionSetNames: ["ObjectStorageReadOnly"],
 *     }],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Policies can be imported using the `{id}`, e.g.
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:index/iamPolicy:IamPolicy main 11111111-1111-1111-1111-111111111111
 * ```
 */
export class IamPolicy extends pulumi.CustomResource {
    /**
     * Get an existing IamPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IamPolicyState, opts?: pulumi.CustomResourceOptions): IamPolicy {
        return new IamPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/iamPolicy:IamPolicy';

    /**
     * Returns true if the given object is an instance of IamPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IamPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IamPolicy.__pulumiType;
    }

    /**
     * ID of the Application the policy will be linked to
     */
    public readonly applicationId!: pulumi.Output<string | undefined>;
    /**
     * The date and time of the creation of the policy.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The description of the iam policy.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Whether the policy is editable.
     */
    public /*out*/ readonly editable!: pulumi.Output<boolean>;
    /**
     * ID of the Group the policy will be linked to
     */
    public readonly groupId!: pulumi.Output<string | undefined>;
    /**
     * .The name of the iam policy.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * If the policy doesn't apply to a principal.
     *
     * > **Important** Only one of `userId`, `groupId`, `applicationId` and `noPrincipal`  may be set.
     */
    public readonly noPrincipal!: pulumi.Output<boolean | undefined>;
    /**
     * ID of organization scoped to the rule.
     */
    public readonly organizationId!: pulumi.Output<string>;
    /**
     * List of rules in the policy.
     */
    public readonly rules!: pulumi.Output<outputs.IamPolicyRule[]>;
    /**
     * The date and time of the last update of the policy.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * ID of the User the policy will be linked to
     */
    public readonly userId!: pulumi.Output<string | undefined>;

    /**
     * Create a IamPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IamPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IamPolicyArgs | IamPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IamPolicyState | undefined;
            resourceInputs["applicationId"] = state ? state.applicationId : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["editable"] = state ? state.editable : undefined;
            resourceInputs["groupId"] = state ? state.groupId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["noPrincipal"] = state ? state.noPrincipal : undefined;
            resourceInputs["organizationId"] = state ? state.organizationId : undefined;
            resourceInputs["rules"] = state ? state.rules : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
            resourceInputs["userId"] = state ? state.userId : undefined;
        } else {
            const args = argsOrState as IamPolicyArgs | undefined;
            if ((!args || args.rules === undefined) && !opts.urn) {
                throw new Error("Missing required property 'rules'");
            }
            resourceInputs["applicationId"] = args ? args.applicationId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["groupId"] = args ? args.groupId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["noPrincipal"] = args ? args.noPrincipal : undefined;
            resourceInputs["organizationId"] = args ? args.organizationId : undefined;
            resourceInputs["rules"] = args ? args.rules : undefined;
            resourceInputs["userId"] = args ? args.userId : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["editable"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IamPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IamPolicy resources.
 */
export interface IamPolicyState {
    /**
     * ID of the Application the policy will be linked to
     */
    applicationId?: pulumi.Input<string>;
    /**
     * The date and time of the creation of the policy.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The description of the iam policy.
     */
    description?: pulumi.Input<string>;
    /**
     * Whether the policy is editable.
     */
    editable?: pulumi.Input<boolean>;
    /**
     * ID of the Group the policy will be linked to
     */
    groupId?: pulumi.Input<string>;
    /**
     * .The name of the iam policy.
     */
    name?: pulumi.Input<string>;
    /**
     * If the policy doesn't apply to a principal.
     *
     * > **Important** Only one of `userId`, `groupId`, `applicationId` and `noPrincipal`  may be set.
     */
    noPrincipal?: pulumi.Input<boolean>;
    /**
     * ID of organization scoped to the rule.
     */
    organizationId?: pulumi.Input<string>;
    /**
     * List of rules in the policy.
     */
    rules?: pulumi.Input<pulumi.Input<inputs.IamPolicyRule>[]>;
    /**
     * The date and time of the last update of the policy.
     */
    updatedAt?: pulumi.Input<string>;
    /**
     * ID of the User the policy will be linked to
     */
    userId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IamPolicy resource.
 */
export interface IamPolicyArgs {
    /**
     * ID of the Application the policy will be linked to
     */
    applicationId?: pulumi.Input<string>;
    /**
     * The description of the iam policy.
     */
    description?: pulumi.Input<string>;
    /**
     * ID of the Group the policy will be linked to
     */
    groupId?: pulumi.Input<string>;
    /**
     * .The name of the iam policy.
     */
    name?: pulumi.Input<string>;
    /**
     * If the policy doesn't apply to a principal.
     *
     * > **Important** Only one of `userId`, `groupId`, `applicationId` and `noPrincipal`  may be set.
     */
    noPrincipal?: pulumi.Input<boolean>;
    /**
     * ID of organization scoped to the rule.
     */
    organizationId?: pulumi.Input<string>;
    /**
     * List of rules in the policy.
     */
    rules: pulumi.Input<pulumi.Input<inputs.IamPolicyRule>[]>;
    /**
     * ID of the User the policy will be linked to
     */
    userId?: pulumi.Input<string>;
}
