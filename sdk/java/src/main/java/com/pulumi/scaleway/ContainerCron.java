// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.scaleway.ContainerCronArgs;
import com.pulumi.scaleway.Utilities;
import com.pulumi.scaleway.inputs.ContainerCronState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Creates and manages Scaleway Container Triggers. For the moment, the feature is limited to CRON Schedule (time-based).
 * 
 * For more information consult
 * the [documentation](https://www.scaleway.com/en/docs/compute/containers/api-cli/cont-uploading-with-serverless-framework/#configuring-events)
 * .
 * 
 * For more details about the limitation
 * check [containers-limitations](https://www.scaleway.com/en/docs/compute/containers/reference-content/containers-limitations/)
 * .
 * 
 * You can check also
 * our [containers cron api documentation](https://developers.scaleway.com/en/products/containers/api/#crons-942bf4).
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.scaleway.ContainerNamespace;
 * import com.pulumi.scaleway.Container;
 * import com.pulumi.scaleway.ContainerArgs;
 * import com.pulumi.scaleway.ContainerCron;
 * import com.pulumi.scaleway.ContainerCronArgs;
 * import static com.pulumi.codegen.internal.Serialization.*;
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
 *         var mainContainerNamespace = new ContainerNamespace(&#34;mainContainerNamespace&#34;);
 * 
 *         var mainContainer = new Container(&#34;mainContainer&#34;, ContainerArgs.builder()        
 *             .namespaceId(mainContainerNamespace.id())
 *             .build());
 * 
 *         var mainContainerCron = new ContainerCron(&#34;mainContainerCron&#34;, ContainerCronArgs.builder()        
 *             .containerId(mainContainer.id())
 *             .schedule(&#34;5 4 1 * *&#34;)
 *             .args(serializeJson(
 *                 jsonObject(
 *                     jsonProperty(&#34;address&#34;, jsonObject(
 *                         jsonProperty(&#34;city&#34;, &#34;Paris&#34;),
 *                         jsonProperty(&#34;country&#34;, &#34;FR&#34;)
 *                     )),
 *                     jsonProperty(&#34;age&#34;, 23),
 *                     jsonProperty(&#34;firstName&#34;, &#34;John&#34;),
 *                     jsonProperty(&#34;isAlive&#34;, true),
 *                     jsonProperty(&#34;lastName&#34;, &#34;Smith&#34;)
 *                 )))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Container Cron can be imported using the `{region}/{id}`, e.g. bash
 * 
 * ```sh
 *  $ pulumi import scaleway:index/containerCron:ContainerCron main fr-par/11111111-1111-1111-1111-111111111111
 * ```
 * 
 */
@ResourceType(type="scaleway:index/containerCron:ContainerCron")
public class ContainerCron extends com.pulumi.resources.CustomResource {
    /**
     * The key-value mapping to define arguments that will be passed to your container’s event object
     * during
     * 
     */
    @Export(name="args", type=String.class, parameters={})
    private Output<String> args;

    /**
     * @return The key-value mapping to define arguments that will be passed to your container’s event object
     * during
     * 
     */
    public Output<String> args() {
        return this.args;
    }
    /**
     * The container ID to link with your cron.
     * 
     */
    @Export(name="containerId", type=String.class, parameters={})
    private Output<String> containerId;

    /**
     * @return The container ID to link with your cron.
     * 
     */
    public Output<String> containerId() {
        return this.containerId;
    }
    /**
     * (Defaults to provider `region`) The region
     * in where the job was created.
     * 
     */
    @Export(name="region", type=String.class, parameters={})
    private Output<String> region;

    /**
     * @return (Defaults to provider `region`) The region
     * in where the job was created.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * Cron format string, e.g. @hourly, as schedule time of its jobs to be created and
     * executed.
     * 
     */
    @Export(name="schedule", type=String.class, parameters={})
    private Output<String> schedule;

    /**
     * @return Cron format string, e.g. @hourly, as schedule time of its jobs to be created and
     * executed.
     * 
     */
    public Output<String> schedule() {
        return this.schedule;
    }
    /**
     * The cron status.
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return The cron status.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ContainerCron(String name) {
        this(name, ContainerCronArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ContainerCron(String name, ContainerCronArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ContainerCron(String name, ContainerCronArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("scaleway:index/containerCron:ContainerCron", name, args == null ? ContainerCronArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ContainerCron(String name, Output<String> id, @Nullable ContainerCronState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("scaleway:index/containerCron:ContainerCron", name, state, makeResourceOptions(options, id));
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
    public static ContainerCron get(String name, Output<String> id, @Nullable ContainerCronState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ContainerCron(name, id, state, options);
    }
}