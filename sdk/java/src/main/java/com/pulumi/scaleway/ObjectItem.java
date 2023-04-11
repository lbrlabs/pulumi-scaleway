// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.scaleway.ObjectItemArgs;
import com.pulumi.scaleway.Utilities;
import com.pulumi.scaleway.inputs.ObjectItemState;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Creates and manages Scaleway object storage objects.
 * For more information, see [the documentation](https://www.scaleway.com/en/docs/object-storage-feature/).
 * 
 * ## Import
 * 
 * Objects can be imported using the `{region}/{bucketName}/{objectKey}` identifier, e.g. bash
 * 
 * ```sh
 *  $ pulumi import scaleway:index/objectItem:ObjectItem some_object fr-par/some-bucket/some-file
 * ```
 * 
 */
@ResourceType(type="scaleway:index/objectItem:ObjectItem")
public class ObjectItem extends com.pulumi.resources.CustomResource {
    /**
     * The name of the bucket.
     * 
     */
    @Export(name="bucket", refs={String.class}, tree="[0]")
    private Output<String> bucket;

    /**
     * @return The name of the bucket.
     * 
     */
    public Output<String> bucket() {
        return this.bucket;
    }
    /**
     * The name of the file to upload, defaults to an empty file
     * 
     */
    @Export(name="file", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> file;

    /**
     * @return The name of the file to upload, defaults to an empty file
     * 
     */
    public Output<Optional<String>> file() {
        return Codegen.optional(this.file);
    }
    /**
     * Hash of the file, used to trigger upload on file change
     * 
     */
    @Export(name="hash", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> hash;

    /**
     * @return Hash of the file, used to trigger upload on file change
     * 
     */
    public Output<Optional<String>> hash() {
        return Codegen.optional(this.hash);
    }
    /**
     * The path of the object.
     * 
     */
    @Export(name="key", refs={String.class}, tree="[0]")
    private Output<String> key;

    /**
     * @return The path of the object.
     * 
     */
    public Output<String> key() {
        return this.key;
    }
    /**
     * Map of metadata used for the object, keys must be lowercase
     * 
     */
    @Export(name="metadata", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> metadata;

    /**
     * @return Map of metadata used for the object, keys must be lowercase
     * 
     */
    public Output<Optional<Map<String,String>>> metadata() {
        return Codegen.optional(this.metadata);
    }
    /**
     * `project_id`) The ID of the project the bucket is associated with.
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output<String> projectId;

    /**
     * @return `project_id`) The ID of the project the bucket is associated with.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * The Scaleway region this bucket resides in.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The Scaleway region this bucket resides in.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * Specifies the Scaleway [storage class](https://www.scaleway.com/en/docs/storage/object/concepts/#storage-class) `STANDARD`, `GLACIER`, `ONEZONE_IA` used to store the object.
     * 
     */
    @Export(name="storageClass", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> storageClass;

    /**
     * @return Specifies the Scaleway [storage class](https://www.scaleway.com/en/docs/storage/object/concepts/#storage-class) `STANDARD`, `GLACIER`, `ONEZONE_IA` used to store the object.
     * 
     */
    public Output<Optional<String>> storageClass() {
        return Codegen.optional(this.storageClass);
    }
    /**
     * Map of tags
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return Map of tags
     * 
     */
    public Output<Optional<Map<String,String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * Visibility of the object, `public-read` or `private`
     * 
     */
    @Export(name="visibility", refs={String.class}, tree="[0]")
    private Output<String> visibility;

    /**
     * @return Visibility of the object, `public-read` or `private`
     * 
     */
    public Output<String> visibility() {
        return this.visibility;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ObjectItem(String name) {
        this(name, ObjectItemArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ObjectItem(String name, ObjectItemArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ObjectItem(String name, ObjectItemArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("scaleway:index/objectItem:ObjectItem", name, args == null ? ObjectItemArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ObjectItem(String name, Output<String> id, @Nullable ObjectItemState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("scaleway:index/objectItem:ObjectItem", name, state, makeResourceOptions(options, id));
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
    public static ObjectItem get(String name, Output<String> id, @Nullable ObjectItemState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ObjectItem(name, id, state, options);
    }
}
