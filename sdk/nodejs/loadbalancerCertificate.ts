// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class LoadbalancerCertificate extends pulumi.CustomResource {
    /**
     * Get an existing LoadbalancerCertificate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LoadbalancerCertificateState, opts?: pulumi.CustomResourceOptions): LoadbalancerCertificate {
        return new LoadbalancerCertificate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/loadbalancerCertificate:LoadbalancerCertificate';

    /**
     * Returns true if the given object is an instance of LoadbalancerCertificate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LoadbalancerCertificate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LoadbalancerCertificate.__pulumiType;
    }

    /**
     * Main domain of the certificate. A new certificate will be created if this field is changed.
     */
    public /*out*/ readonly commonName!: pulumi.Output<string>;
    /**
     * Configuration block for custom certificate chain. Only one of `letsencrypt` and `customCertificate` should be specified.
     */
    public readonly customCertificate!: pulumi.Output<outputs.LoadbalancerCertificateCustomCertificate | undefined>;
    /**
     * The identifier (SHA-1) of the certificate
     */
    public /*out*/ readonly fingerprint!: pulumi.Output<string>;
    /**
     * The load-balancer ID this certificate is attached to.
     *
     * > **Important:** Updates to `lbId` will recreate the load-balancer certificate.
     */
    public readonly lbId!: pulumi.Output<string>;
    /**
     * Configuration block for Let's Encrypt configuration. Only one of `letsencrypt` and `customCertificate` should be specified.
     */
    public readonly letsencrypt!: pulumi.Output<outputs.LoadbalancerCertificateLetsencrypt | undefined>;
    /**
     * The name of the certificate backend.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The not valid after validity bound timestamp
     */
    public /*out*/ readonly notValidAfter!: pulumi.Output<string>;
    /**
     * The not valid before validity bound timestamp
     */
    public /*out*/ readonly notValidBefore!: pulumi.Output<string>;
    /**
     * Certificate status
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Array of alternative domain names.  A new certificate will be created if this field is changed.
     *
     * > **Important:** Updates to `letsencrypt` will recreate the load-balancer certificate.
     */
    public /*out*/ readonly subjectAlternativeNames!: pulumi.Output<string[]>;

    /**
     * Create a LoadbalancerCertificate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LoadbalancerCertificateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LoadbalancerCertificateArgs | LoadbalancerCertificateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LoadbalancerCertificateState | undefined;
            resourceInputs["commonName"] = state ? state.commonName : undefined;
            resourceInputs["customCertificate"] = state ? state.customCertificate : undefined;
            resourceInputs["fingerprint"] = state ? state.fingerprint : undefined;
            resourceInputs["lbId"] = state ? state.lbId : undefined;
            resourceInputs["letsencrypt"] = state ? state.letsencrypt : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["notValidAfter"] = state ? state.notValidAfter : undefined;
            resourceInputs["notValidBefore"] = state ? state.notValidBefore : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["subjectAlternativeNames"] = state ? state.subjectAlternativeNames : undefined;
        } else {
            const args = argsOrState as LoadbalancerCertificateArgs | undefined;
            if ((!args || args.lbId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'lbId'");
            }
            resourceInputs["customCertificate"] = args ? args.customCertificate : undefined;
            resourceInputs["lbId"] = args ? args.lbId : undefined;
            resourceInputs["letsencrypt"] = args ? args.letsencrypt : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["commonName"] = undefined /*out*/;
            resourceInputs["fingerprint"] = undefined /*out*/;
            resourceInputs["notValidAfter"] = undefined /*out*/;
            resourceInputs["notValidBefore"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["subjectAlternativeNames"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LoadbalancerCertificate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LoadbalancerCertificate resources.
 */
export interface LoadbalancerCertificateState {
    /**
     * Main domain of the certificate. A new certificate will be created if this field is changed.
     */
    commonName?: pulumi.Input<string>;
    /**
     * Configuration block for custom certificate chain. Only one of `letsencrypt` and `customCertificate` should be specified.
     */
    customCertificate?: pulumi.Input<inputs.LoadbalancerCertificateCustomCertificate>;
    /**
     * The identifier (SHA-1) of the certificate
     */
    fingerprint?: pulumi.Input<string>;
    /**
     * The load-balancer ID this certificate is attached to.
     *
     * > **Important:** Updates to `lbId` will recreate the load-balancer certificate.
     */
    lbId?: pulumi.Input<string>;
    /**
     * Configuration block for Let's Encrypt configuration. Only one of `letsencrypt` and `customCertificate` should be specified.
     */
    letsencrypt?: pulumi.Input<inputs.LoadbalancerCertificateLetsencrypt>;
    /**
     * The name of the certificate backend.
     */
    name?: pulumi.Input<string>;
    /**
     * The not valid after validity bound timestamp
     */
    notValidAfter?: pulumi.Input<string>;
    /**
     * The not valid before validity bound timestamp
     */
    notValidBefore?: pulumi.Input<string>;
    /**
     * Certificate status
     */
    status?: pulumi.Input<string>;
    /**
     * Array of alternative domain names.  A new certificate will be created if this field is changed.
     *
     * > **Important:** Updates to `letsencrypt` will recreate the load-balancer certificate.
     */
    subjectAlternativeNames?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a LoadbalancerCertificate resource.
 */
export interface LoadbalancerCertificateArgs {
    /**
     * Configuration block for custom certificate chain. Only one of `letsencrypt` and `customCertificate` should be specified.
     */
    customCertificate?: pulumi.Input<inputs.LoadbalancerCertificateCustomCertificate>;
    /**
     * The load-balancer ID this certificate is attached to.
     *
     * > **Important:** Updates to `lbId` will recreate the load-balancer certificate.
     */
    lbId: pulumi.Input<string>;
    /**
     * Configuration block for Let's Encrypt configuration. Only one of `letsencrypt` and `customCertificate` should be specified.
     */
    letsencrypt?: pulumi.Input<inputs.LoadbalancerCertificateLetsencrypt>;
    /**
     * The name of the certificate backend.
     */
    name?: pulumi.Input<string>;
}
