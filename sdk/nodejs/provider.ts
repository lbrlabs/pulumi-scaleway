// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The provider type for the scaleway package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 */
export class Provider extends pulumi.ProviderResource {
    /** @internal */
    public static readonly __pulumiType = 'scaleway';

    /**
     * Returns true if the given object is an instance of Provider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Provider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Provider.__pulumiType;
    }

    /**
     * The Scaleway access key.
     */
    public readonly accessKey!: pulumi.Output<string | undefined>;
    /**
     * The Scaleway API URL to use.
     */
    public readonly apiUrl!: pulumi.Output<string | undefined>;
    /**
     * The Scaleway profile to use.
     */
    public readonly profile!: pulumi.Output<string | undefined>;
    /**
     * The Scaleway project ID.
     */
    public readonly projectId!: pulumi.Output<string | undefined>;
    /**
     * The region you want to attach the resource to
     */
    public readonly region!: pulumi.Output<string | undefined>;
    /**
     * The Scaleway secret Key.
     */
    public readonly secretKey!: pulumi.Output<string | undefined>;
    /**
     * The zone you want to attach the resource to
     */
    public readonly zone!: pulumi.Output<string | undefined>;

    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        {
            resourceInputs["accessKey"] = (args ? args.accessKey : undefined) ?? utilities.getEnv("SCW_ACCESS_KEY");
            resourceInputs["apiUrl"] = args ? args.apiUrl : undefined;
            resourceInputs["profile"] = args ? args.profile : undefined;
            resourceInputs["projectId"] = (args ? args.projectId : undefined) ?? utilities.getEnv("SCW_DEFAULT_PROJECT_ID");
            resourceInputs["region"] = (args ? args.region : undefined) ?? utilities.getEnv("SCW_DEFAULT_REGION");
            resourceInputs["secretKey"] = (args ? args.secretKey : undefined) ?? utilities.getEnv("SCW_SECRET_KEY");
            resourceInputs["zone"] = (args ? args.zone : undefined) ?? utilities.getEnv("SCW_DEFAULT_ZONE");
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Provider.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    /**
     * The Scaleway access key.
     */
    accessKey?: pulumi.Input<string>;
    /**
     * The Scaleway API URL to use.
     */
    apiUrl?: pulumi.Input<string>;
    /**
     * The Scaleway profile to use.
     */
    profile?: pulumi.Input<string>;
    /**
     * The Scaleway project ID.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The region you want to attach the resource to
     */
    region?: pulumi.Input<string>;
    /**
     * The Scaleway secret Key.
     */
    secretKey?: pulumi.Input<string>;
    /**
     * The zone you want to attach the resource to
     */
    zone?: pulumi.Input<string>;
}
