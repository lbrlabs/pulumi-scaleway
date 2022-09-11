// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.scaleway.InstancePlacementGroupArgs;
import com.pulumi.scaleway.Utilities;
import com.pulumi.scaleway.inputs.InstancePlacementGroupState;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Creates and manages Compute Instance Placement Groups. For more information, see [the documentation](https://developers.scaleway.com/en/products/instance/api/#placement-groups-d8f653).
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.scaleway.InstancePlacementGroup;
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
 *         var availabilityGroup = new InstancePlacementGroup(&#34;availabilityGroup&#34;);
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Placement groups can be imported using the `{zone}/{id}`, e.g. bash
 * 
 * ```sh
 *  $ pulumi import scaleway:index/instancePlacementGroup:InstancePlacementGroup availability_group fr-par-1/11111111-1111-1111-1111-111111111111
 * ```
 * 
 */
@ResourceType(type="scaleway:index/instancePlacementGroup:InstancePlacementGroup")
public class InstancePlacementGroup extends com.pulumi.resources.CustomResource {
    /**
     * The name of the placement group.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The name of the placement group.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The organization ID the placement group is associated with.
     * 
     */
    @Export(name="organizationId", type=String.class, parameters={})
    private Output<String> organizationId;

    /**
     * @return The organization ID the placement group is associated with.
     * 
     */
    public Output<String> organizationId() {
        return this.organizationId;
    }
    /**
     * The [policy mode](https://developers.scaleway.com/en/products/instance/api/#placement-groups-d8f653) of the placement group. Possible values are: `optional` or `enforced`.
     * 
     */
    @Export(name="policyMode", type=String.class, parameters={})
    private Output</* @Nullable */ String> policyMode;

    /**
     * @return The [policy mode](https://developers.scaleway.com/en/products/instance/api/#placement-groups-d8f653) of the placement group. Possible values are: `optional` or `enforced`.
     * 
     */
    public Output<Optional<String>> policyMode() {
        return Codegen.optional(this.policyMode);
    }
    /**
     * Is true when the policy is respected.
     * 
     */
    @Export(name="policyRespected", type=Boolean.class, parameters={})
    private Output<Boolean> policyRespected;

    /**
     * @return Is true when the policy is respected.
     * 
     */
    public Output<Boolean> policyRespected() {
        return this.policyRespected;
    }
    /**
     * The [policy type](https://developers.scaleway.com/en/products/instance/api/#placement-groups-d8f653) of the placement group. Possible values are: `low_latency` or `max_availability`.
     * 
     */
    @Export(name="policyType", type=String.class, parameters={})
    private Output</* @Nullable */ String> policyType;

    /**
     * @return The [policy type](https://developers.scaleway.com/en/products/instance/api/#placement-groups-d8f653) of the placement group. Possible values are: `low_latency` or `max_availability`.
     * 
     */
    public Output<Optional<String>> policyType() {
        return Codegen.optional(this.policyType);
    }
    /**
     * `project_id`) The ID of the project the placement group is associated with.
     * 
     */
    @Export(name="projectId", type=String.class, parameters={})
    private Output<String> projectId;

    /**
     * @return `project_id`) The ID of the project the placement group is associated with.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * A list of tags to apply to the placement group.
     * 
     */
    @Export(name="tags", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> tags;

    /**
     * @return A list of tags to apply to the placement group.
     * 
     */
    public Output<Optional<List<String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * `zone`) The zone in which the placement group should be created.
     * 
     */
    @Export(name="zone", type=String.class, parameters={})
    private Output<String> zone;

    /**
     * @return `zone`) The zone in which the placement group should be created.
     * 
     */
    public Output<String> zone() {
        return this.zone;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public InstancePlacementGroup(String name) {
        this(name, InstancePlacementGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public InstancePlacementGroup(String name, @Nullable InstancePlacementGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public InstancePlacementGroup(String name, @Nullable InstancePlacementGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("scaleway:index/instancePlacementGroup:InstancePlacementGroup", name, args == null ? InstancePlacementGroupArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private InstancePlacementGroup(String name, Output<String> id, @Nullable InstancePlacementGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("scaleway:index/instancePlacementGroup:InstancePlacementGroup", name, state, makeResourceOptions(options, id));
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
    public static InstancePlacementGroup get(String name, Output<String> id, @Nullable InstancePlacementGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new InstancePlacementGroup(name, id, state, options);
    }
}