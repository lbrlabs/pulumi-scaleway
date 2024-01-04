// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.scaleway.JobDefinitionArgs;
import com.pulumi.scaleway.Utilities;
import com.pulumi.scaleway.inputs.JobDefinitionState;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Creates and manages a Scaleway Serverless Job Definition. For more information, see [the documentation](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/jobs/v1alpha1).
 * 
 * ## Example
 * 
 * ### Basic
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.scaleway.JobDefinition;
 * import com.pulumi.scaleway.JobDefinitionArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var main = new JobDefinition(&#34;main&#34;, JobDefinitionArgs.builder()        
 *             .cpuLimit(140)
 *             .memoryLimit(256)
 *             .imageUri(&#34;docker.io/alpine:latest&#34;)
 *             .command(&#34;ls&#34;)
 *             .timeout(&#34;10m&#34;)
 *             .env(Map.of(&#34;foo&#34;, &#34;bar&#34;))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Serverless Jobs can be imported using the `{region}/{id}`, e.g. bash
 * 
 * ```sh
 *  $ pulumi import scaleway:index/jobDefinition:JobDefinition job fr-par/11111111-1111-1111-1111-111111111111
 * ```
 * 
 */
@ResourceType(type="scaleway:index/jobDefinition:JobDefinition")
public class JobDefinition extends com.pulumi.resources.CustomResource {
    /**
     * The command that will be run in the container if specified.
     * 
     */
    @Export(name="command", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> command;

    /**
     * @return The command that will be run in the container if specified.
     * 
     */
    public Output<Optional<String>> command() {
        return Codegen.optional(this.command);
    }
    /**
     * The amount of vCPU computing resources to allocate to each container running the job.
     * 
     */
    @Export(name="cpuLimit", refs={Integer.class}, tree="[0]")
    private Output<Integer> cpuLimit;

    /**
     * @return The amount of vCPU computing resources to allocate to each container running the job.
     * 
     */
    public Output<Integer> cpuLimit() {
        return this.cpuLimit;
    }
    /**
     * The description of the job
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the job
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The environment variables of the container.
     * 
     */
    @Export(name="env", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> env;

    /**
     * @return The environment variables of the container.
     * 
     */
    public Output<Optional<Map<String,String>>> env() {
        return Codegen.optional(this.env);
    }
    /**
     * The uri of the container image that will be used for the job run.
     * 
     */
    @Export(name="imageUri", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> imageUri;

    /**
     * @return The uri of the container image that will be used for the job run.
     * 
     */
    public Output<Optional<String>> imageUri() {
        return Codegen.optional(this.imageUri);
    }
    /**
     * The memory computing resources in MB to allocate to each container running the job.
     * 
     */
    @Export(name="memoryLimit", refs={Integer.class}, tree="[0]")
    private Output<Integer> memoryLimit;

    /**
     * @return The memory computing resources in MB to allocate to each container running the job.
     * 
     */
    public Output<Integer> memoryLimit() {
        return this.memoryLimit;
    }
    /**
     * The name of the job.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the job.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * `project_id`) The ID of the project the Job is associated with.
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output<String> projectId;

    /**
     * @return `project_id`) The ID of the project the Job is associated with.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * `region`) The region of the Job.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return `region`) The region of the Job.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The job run timeout, in Go Time format (ex: `2h30m25s`)
     * 
     */
    @Export(name="timeout", refs={String.class}, tree="[0]")
    private Output<String> timeout;

    /**
     * @return The job run timeout, in Go Time format (ex: `2h30m25s`)
     * 
     */
    public Output<String> timeout() {
        return this.timeout;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public JobDefinition(String name) {
        this(name, JobDefinitionArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public JobDefinition(String name, JobDefinitionArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public JobDefinition(String name, JobDefinitionArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("scaleway:index/jobDefinition:JobDefinition", name, args == null ? JobDefinitionArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private JobDefinition(String name, Output<String> id, @Nullable JobDefinitionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("scaleway:index/jobDefinition:JobDefinition", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static JobDefinition get(String name, Output<String> id, @Nullable JobDefinitionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new JobDefinition(name, id, state, options);
    }
}
