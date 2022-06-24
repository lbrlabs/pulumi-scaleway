// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export function getLoadbalancerCertificate(args?: GetLoadbalancerCertificateArgs, opts?: pulumi.InvokeOptions): Promise<GetLoadbalancerCertificateResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("scaleway:index/getLoadbalancerCertificate:getLoadbalancerCertificate", {
        "certificateId": args.certificateId,
        "lbId": args.lbId,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getLoadbalancerCertificate.
 */
export interface GetLoadbalancerCertificateArgs {
    certificateId?: string;
    lbId?: string;
    name?: string;
}

/**
 * A collection of values returned by getLoadbalancerCertificate.
 */
export interface GetLoadbalancerCertificateResult {
    readonly certificateId?: string;
    readonly commonName: string;
    readonly customCertificates: outputs.GetLoadbalancerCertificateCustomCertificate[];
    readonly fingerprint: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly lbId?: string;
    readonly letsencrypts: outputs.GetLoadbalancerCertificateLetsencrypt[];
    readonly name?: string;
    readonly notValidAfter: string;
    readonly notValidBefore: string;
    readonly status: string;
    readonly subjectAlternativeNames: string[];
}

export function getLoadbalancerCertificateOutput(args?: GetLoadbalancerCertificateOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLoadbalancerCertificateResult> {
    return pulumi.output(args).apply(a => getLoadbalancerCertificate(a, opts))
}

/**
 * A collection of arguments for invoking getLoadbalancerCertificate.
 */
export interface GetLoadbalancerCertificateOutputArgs {
    certificateId?: pulumi.Input<string>;
    lbId?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
}