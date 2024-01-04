// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.scaleway.InstanceVolumeArgs;
import com.pulumi.scaleway.Utilities;
import com.pulumi.scaleway.inputs.InstanceVolumeState;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Creates and manages Scaleway Compute Instance Volumes.
 * For more information, see [the documentation](https://developers.scaleway.com/en/products/instance/api/#volumes-7e8a39).
 * 
 * ## Example
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.scaleway.InstanceVolume;
 * import com.pulumi.scaleway.InstanceVolumeArgs;
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
 *         var serverVolume = new InstanceVolume(&#34;serverVolume&#34;, InstanceVolumeArgs.builder()        
 *             .sizeInGb(20)
 *             .type(&#34;l_ssd&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * volumes can be imported using the `{zone}/{id}`, e.g. bash
 * 
 * ```sh
 *  $ pulumi import scaleway:index/instanceVolume:InstanceVolume server_volume fr-par-1/11111111-1111-1111-1111-111111111111
 * ```
 * 
 */
@ResourceType(type="scaleway:index/instanceVolume:InstanceVolume")
public class InstanceVolume extends com.pulumi.resources.CustomResource {
    /**
     * If set, the new volume will be created from this snapshot. Only one of `size_in_gb` and `from_snapshot_id` should be specified.
     * 
     */
    @Export(name="fromSnapshotId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> fromSnapshotId;

    /**
     * @return If set, the new volume will be created from this snapshot. Only one of `size_in_gb` and `from_snapshot_id` should be specified.
     * 
     */
    public Output<Optional<String>> fromSnapshotId() {
        return Codegen.optional(this.fromSnapshotId);
    }
    /**
     * The name of the volume. If not provided it will be randomly generated.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the volume. If not provided it will be randomly generated.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The organization ID the volume is associated with.
     * 
     */
    @Export(name="organizationId", refs={String.class}, tree="[0]")
    private Output<String> organizationId;

    /**
     * @return The organization ID the volume is associated with.
     * 
     */
    public Output<String> organizationId() {
        return this.organizationId;
    }
    /**
     * `project_id`) The ID of the project the volume is associated with.
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output<String> projectId;

    /**
     * @return `project_id`) The ID of the project the volume is associated with.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * The id of the associated server.
     * 
     */
    @Export(name="serverId", refs={String.class}, tree="[0]")
    private Output<String> serverId;

    /**
     * @return The id of the associated server.
     * 
     */
    public Output<String> serverId() {
        return this.serverId;
    }
    /**
     * The size of the volume. Only one of `size_in_gb` and `from_snapshot_id` should be specified.
     * 
     */
    @Export(name="sizeInGb", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> sizeInGb;

    /**
     * @return The size of the volume. Only one of `size_in_gb` and `from_snapshot_id` should be specified.
     * 
     */
    public Output<Optional<Integer>> sizeInGb() {
        return Codegen.optional(this.sizeInGb);
    }
    /**
     * A list of tags to apply to the volume.
     * 
     */
    @Export(name="tags", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> tags;

    /**
     * @return A list of tags to apply to the volume.
     * 
     */
    public Output<Optional<List<String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * The type of the volume. The possible values are: `b_ssd` (Block SSD), `l_ssd` (Local SSD), `scratch` (Local Scratch SSD).
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return The type of the volume. The possible values are: `b_ssd` (Block SSD), `l_ssd` (Local SSD), `scratch` (Local Scratch SSD).
     * 
     */
    public Output<String> type() {
        return this.type;
    }
    /**
     * `zone`) The zone in which the volume should be created.
     * 
     */
    @Export(name="zone", refs={String.class}, tree="[0]")
    private Output<String> zone;

    /**
     * @return `zone`) The zone in which the volume should be created.
     * 
     */
    public Output<String> zone() {
        return this.zone;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public InstanceVolume(String name) {
        this(name, InstanceVolumeArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public InstanceVolume(String name, InstanceVolumeArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public InstanceVolume(String name, InstanceVolumeArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("scaleway:index/instanceVolume:InstanceVolume", name, args == null ? InstanceVolumeArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private InstanceVolume(String name, Output<String> id, @Nullable InstanceVolumeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("scaleway:index/instanceVolume:InstanceVolume", name, state, makeResourceOptions(options, id));
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
    public static InstanceVolume get(String name, Output<String> id, @Nullable InstanceVolumeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new InstanceVolume(name, id, state, options);
    }
}
