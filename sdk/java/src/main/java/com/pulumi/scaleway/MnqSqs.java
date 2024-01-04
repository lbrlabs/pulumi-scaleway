// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.scaleway.MnqSqsArgs;
import com.pulumi.scaleway.Utilities;
import com.pulumi.scaleway.inputs.MnqSqsState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Activate Scaleway Messaging and queuing SQS for a project.
 * For further information please check
 * our [documentation](https://www.scaleway.com/en/docs/serverless/messaging/reference-content/sqs-overview/)
 * 
 * ## Examples
 * 
 * ### Basic
 * 
 * Activate SQS for default project
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.scaleway.MnqSqs;
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
 *         var main = new MnqSqs(&#34;main&#34;);
 * 
 *     }
 * }
 * ```
 * 
 * Activate SQS for a specific project
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.scaleway.ScalewayFunctions;
 * import com.pulumi.scaleway.inputs.GetAccountProjectArgs;
 * import com.pulumi.scaleway.MnqSqs;
 * import com.pulumi.scaleway.MnqSqsArgs;
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
 *         final var project = ScalewayFunctions.getAccountProject(GetAccountProjectArgs.builder()
 *             .name(&#34;default&#34;)
 *             .build());
 * 
 *         var forProject = new MnqSqs(&#34;forProject&#34;, MnqSqsArgs.builder()        
 *             .projectId(project.applyValue(getAccountProjectResult -&gt; getAccountProjectResult.id()))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * SQS status can be imported using the `{region}/{project_id}`, e.g. bash
 * 
 * ```sh
 *  $ pulumi import scaleway:index/mnqSqs:MnqSqs main fr-par/11111111111111111111111111111111
 * ```
 * 
 */
@ResourceType(type="scaleway:index/mnqSqs:MnqSqs")
public class MnqSqs extends com.pulumi.resources.CustomResource {
    /**
     * The endpoint of the SQS service for this project.
     * 
     */
    @Export(name="endpoint", refs={String.class}, tree="[0]")
    private Output<String> endpoint;

    /**
     * @return The endpoint of the SQS service for this project.
     * 
     */
    public Output<String> endpoint() {
        return this.endpoint;
    }
    /**
     * `project_id`) The ID of the project the sqs will be enabled for.
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output<String> projectId;

    /**
     * @return `project_id`) The ID of the project the sqs will be enabled for.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * `region`). The region
     * in which sqs will be enabled.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return `region`). The region
     * in which sqs will be enabled.
     * 
     */
    public Output<String> region() {
        return this.region;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public MnqSqs(String name) {
        this(name, MnqSqsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public MnqSqs(String name, @Nullable MnqSqsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public MnqSqs(String name, @Nullable MnqSqsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("scaleway:index/mnqSqs:MnqSqs", name, args == null ? MnqSqsArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private MnqSqs(String name, Output<String> id, @Nullable MnqSqsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("scaleway:index/mnqSqs:MnqSqs", name, state, makeResourceOptions(options, id));
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
    public static MnqSqs get(String name, Output<String> id, @Nullable MnqSqsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new MnqSqs(name, id, state, options);
    }
}
