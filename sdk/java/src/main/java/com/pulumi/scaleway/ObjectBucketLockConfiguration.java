// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.scaleway.ObjectBucketLockConfigurationArgs;
import com.pulumi.scaleway.Utilities;
import com.pulumi.scaleway.inputs.ObjectBucketLockConfigurationState;
import com.pulumi.scaleway.outputs.ObjectBucketLockConfigurationRule;
import java.lang.String;
import javax.annotation.Nullable;

@ResourceType(type="scaleway:index/objectBucketLockConfiguration:ObjectBucketLockConfiguration")
public class ObjectBucketLockConfiguration extends com.pulumi.resources.CustomResource {
    /**
     * The bucket&#39;s name or regional ID.
     * 
     */
    @Export(name="bucket", refs={String.class}, tree="[0]")
    private Output<String> bucket;

    /**
     * @return The bucket&#39;s name or regional ID.
     * 
     */
    public Output<String> bucket() {
        return this.bucket;
    }
    /**
     * The project_id you want to attach the resource to
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output<String> projectId;

    /**
     * @return The project_id you want to attach the resource to
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * The region you want to attach the resource to
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region you want to attach the resource to
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * Specifies the Object Lock rule for the specified object.
     * 
     */
    @Export(name="rule", refs={ObjectBucketLockConfigurationRule.class}, tree="[0]")
    private Output<ObjectBucketLockConfigurationRule> rule;

    /**
     * @return Specifies the Object Lock rule for the specified object.
     * 
     */
    public Output<ObjectBucketLockConfigurationRule> rule() {
        return this.rule;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ObjectBucketLockConfiguration(String name) {
        this(name, ObjectBucketLockConfigurationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ObjectBucketLockConfiguration(String name, ObjectBucketLockConfigurationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ObjectBucketLockConfiguration(String name, ObjectBucketLockConfigurationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("scaleway:index/objectBucketLockConfiguration:ObjectBucketLockConfiguration", name, args == null ? ObjectBucketLockConfigurationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ObjectBucketLockConfiguration(String name, Output<String> id, @Nullable ObjectBucketLockConfigurationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("scaleway:index/objectBucketLockConfiguration:ObjectBucketLockConfiguration", name, state, makeResourceOptions(options, id));
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
    public static ObjectBucketLockConfiguration get(String name, Output<String> id, @Nullable ObjectBucketLockConfigurationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ObjectBucketLockConfiguration(name, id, state, options);
    }
}
