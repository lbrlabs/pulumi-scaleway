// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.scaleway.IamApplicationArgs;
import com.pulumi.scaleway.Utilities;
import com.pulumi.scaleway.inputs.IamApplicationState;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Creates and manages Scaleway IAM Applications. For more information, see [the documentation](https://developers.scaleway.com/en/products/iam/api/v1alpha1/#applications-83ce5e).
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.scaleway.IamApplication;
 * import com.pulumi.scaleway.IamApplicationArgs;
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
 *         var main = new IamApplication(&#34;main&#34;, IamApplicationArgs.builder()        
 *             .description(&#34;a description&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Applications can be imported using the `{id}`, e.g. bash
 * 
 * ```sh
 *  $ pulumi import scaleway:index/iamApplication:IamApplication main 11111111-1111-1111-1111-111111111111
 * ```
 * 
 */
@ResourceType(type="scaleway:index/iamApplication:IamApplication")
public class IamApplication extends com.pulumi.resources.CustomResource {
    /**
     * The date and time of the creation of the application.
     * 
     */
    @Export(name="createdAt", refs={String.class}, tree="[0]")
    private Output<String> createdAt;

    /**
     * @return The date and time of the creation of the application.
     * 
     */
    public Output<String> createdAt() {
        return this.createdAt;
    }
    /**
     * The description of the iam application.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the iam application.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Whether the application is editable.
     * 
     */
    @Export(name="editable", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> editable;

    /**
     * @return Whether the application is editable.
     * 
     */
    public Output<Boolean> editable() {
        return this.editable;
    }
    /**
     * .The name of the iam application.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return .The name of the iam application.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * `organization_id`) The ID of the organization the application is associated with.
     * 
     */
    @Export(name="organizationId", refs={String.class}, tree="[0]")
    private Output<String> organizationId;

    /**
     * @return `organization_id`) The ID of the organization the application is associated with.
     * 
     */
    public Output<String> organizationId() {
        return this.organizationId;
    }
    /**
     * The date and time of the last update of the application.
     * 
     */
    @Export(name="updatedAt", refs={String.class}, tree="[0]")
    private Output<String> updatedAt;

    /**
     * @return The date and time of the last update of the application.
     * 
     */
    public Output<String> updatedAt() {
        return this.updatedAt;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public IamApplication(String name) {
        this(name, IamApplicationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public IamApplication(String name, @Nullable IamApplicationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public IamApplication(String name, @Nullable IamApplicationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("scaleway:index/iamApplication:IamApplication", name, args == null ? IamApplicationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private IamApplication(String name, Output<String> id, @Nullable IamApplicationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("scaleway:index/iamApplication:IamApplication", name, state, makeResourceOptions(options, id));
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
    public static IamApplication get(String name, Output<String> id, @Nullable IamApplicationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new IamApplication(name, id, state, options);
    }
}
