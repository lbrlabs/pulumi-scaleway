// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Creates and manages a Scaleway Serverless Job Definition. For more information, see [the documentation](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/jobs/v1alpha1).
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
 * const main = new scaleway.JobDefinition("main", {
 *     cpuLimit: 140,
 *     memoryLimit: 256,
 *     imageUri: "docker.io/alpine:latest",
 *     command: "ls",
 *     timeout: "10m",
 *     env: {
 *         foo: "bar",
 *     },
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Serverless Jobs can be imported using the `{region}/{id}`, e.g.
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:index/jobDefinition:JobDefinition job fr-par/11111111-1111-1111-1111-111111111111
 * ```
 */
export class JobDefinition extends pulumi.CustomResource {
    /**
     * Get an existing JobDefinition resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: JobDefinitionState, opts?: pulumi.CustomResourceOptions): JobDefinition {
        return new JobDefinition(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/jobDefinition:JobDefinition';

    /**
     * Returns true if the given object is an instance of JobDefinition.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is JobDefinition {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === JobDefinition.__pulumiType;
    }

    /**
     * The command that will be run in the container if specified.
     */
    public readonly command!: pulumi.Output<string | undefined>;
    /**
     * The amount of vCPU computing resources to allocate to each container running the job.
     */
    public readonly cpuLimit!: pulumi.Output<number>;
    /**
     * The description of the job
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The environment variables of the container.
     */
    public readonly env!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The uri of the container image that will be used for the job run.
     */
    public readonly imageUri!: pulumi.Output<string | undefined>;
    /**
     * The memory computing resources in MB to allocate to each container running the job.
     */
    public readonly memoryLimit!: pulumi.Output<number>;
    /**
     * The name of the job.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * `projectId`) The ID of the project the Job is associated with.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * `region`) The region of the Job.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The job run timeout, in Go Time format (ex: `2h30m25s`)
     */
    public readonly timeout!: pulumi.Output<string>;

    /**
     * Create a JobDefinition resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: JobDefinitionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: JobDefinitionArgs | JobDefinitionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as JobDefinitionState | undefined;
            resourceInputs["command"] = state ? state.command : undefined;
            resourceInputs["cpuLimit"] = state ? state.cpuLimit : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["env"] = state ? state.env : undefined;
            resourceInputs["imageUri"] = state ? state.imageUri : undefined;
            resourceInputs["memoryLimit"] = state ? state.memoryLimit : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["timeout"] = state ? state.timeout : undefined;
        } else {
            const args = argsOrState as JobDefinitionArgs | undefined;
            if ((!args || args.cpuLimit === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cpuLimit'");
            }
            if ((!args || args.memoryLimit === undefined) && !opts.urn) {
                throw new Error("Missing required property 'memoryLimit'");
            }
            resourceInputs["command"] = args ? args.command : undefined;
            resourceInputs["cpuLimit"] = args ? args.cpuLimit : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["env"] = args ? args.env : undefined;
            resourceInputs["imageUri"] = args ? args.imageUri : undefined;
            resourceInputs["memoryLimit"] = args ? args.memoryLimit : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["timeout"] = args ? args.timeout : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(JobDefinition.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering JobDefinition resources.
 */
export interface JobDefinitionState {
    /**
     * The command that will be run in the container if specified.
     */
    command?: pulumi.Input<string>;
    /**
     * The amount of vCPU computing resources to allocate to each container running the job.
     */
    cpuLimit?: pulumi.Input<number>;
    /**
     * The description of the job
     */
    description?: pulumi.Input<string>;
    /**
     * The environment variables of the container.
     */
    env?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The uri of the container image that will be used for the job run.
     */
    imageUri?: pulumi.Input<string>;
    /**
     * The memory computing resources in MB to allocate to each container running the job.
     */
    memoryLimit?: pulumi.Input<number>;
    /**
     * The name of the job.
     */
    name?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the Job is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`) The region of the Job.
     */
    region?: pulumi.Input<string>;
    /**
     * The job run timeout, in Go Time format (ex: `2h30m25s`)
     */
    timeout?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a JobDefinition resource.
 */
export interface JobDefinitionArgs {
    /**
     * The command that will be run in the container if specified.
     */
    command?: pulumi.Input<string>;
    /**
     * The amount of vCPU computing resources to allocate to each container running the job.
     */
    cpuLimit: pulumi.Input<number>;
    /**
     * The description of the job
     */
    description?: pulumi.Input<string>;
    /**
     * The environment variables of the container.
     */
    env?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The uri of the container image that will be used for the job run.
     */
    imageUri?: pulumi.Input<string>;
    /**
     * The memory computing resources in MB to allocate to each container running the job.
     */
    memoryLimit: pulumi.Input<number>;
    /**
     * The name of the job.
     */
    name?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the Job is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`) The region of the Job.
     */
    region?: pulumi.Input<string>;
    /**
     * The job run timeout, in Go Time format (ex: `2h30m25s`)
     */
    timeout?: pulumi.Input<string>;
}
