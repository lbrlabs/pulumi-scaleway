// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.scaleway.AccountProjectArgs;
import com.pulumi.scaleway.Utilities;
import com.pulumi.scaleway.inputs.AccountProjectState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages organization&#39;s projects on Scaleway.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.scaleway.AccountProject;
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
 *         var project = new AccountProject(&#34;project&#34;);
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Projects can be imported using the `id`, e.g. bash
 * 
 * ```sh
 *  $ pulumi import scaleway:index/accountProject:AccountProject project 11111111-1111-1111-1111-111111111111
 * ```
 * 
 */
@ResourceType(type="scaleway:index/accountProject:AccountProject")
public class AccountProject extends com.pulumi.resources.CustomResource {
    /**
     * The Project creation time.
     * 
     */
    @Export(name="createdAt", type=String.class, parameters={})
    private Output<String> createdAt;

    /**
     * @return The Project creation time.
     * 
     */
    public Output<String> createdAt() {
        return this.createdAt;
    }
    /**
     * The description of the Project.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the Project.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The name of the Project.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The name of the Project.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * `organization_id`)The organization ID the Project is associated with. Please note that any change in `organization_id` will recreate the resource.
     * 
     */
    @Export(name="organizationId", type=String.class, parameters={})
    private Output<String> organizationId;

    /**
     * @return `organization_id`)The organization ID the Project is associated with. Please note that any change in `organization_id` will recreate the resource.
     * 
     */
    public Output<String> organizationId() {
        return this.organizationId;
    }
    /**
     * The Project last update time.
     * 
     */
    @Export(name="updatedAt", type=String.class, parameters={})
    private Output<String> updatedAt;

    /**
     * @return The Project last update time.
     * 
     */
    public Output<String> updatedAt() {
        return this.updatedAt;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AccountProject(String name) {
        this(name, AccountProjectArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AccountProject(String name, @Nullable AccountProjectArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AccountProject(String name, @Nullable AccountProjectArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("scaleway:index/accountProject:AccountProject", name, args == null ? AccountProjectArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AccountProject(String name, Output<String> id, @Nullable AccountProjectState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("scaleway:index/accountProject:AccountProject", name, state, makeResourceOptions(options, id));
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
    public static AccountProject get(String name, Output<String> id, @Nullable AccountProjectState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AccountProject(name, id, state, options);
    }
}