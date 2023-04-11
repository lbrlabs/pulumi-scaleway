// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.scaleway.InstanceUserDataArgs;
import com.pulumi.scaleway.Utilities;
import com.pulumi.scaleway.inputs.InstanceUserDataState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Creates and manages Scaleway Compute Instance User Data values.
 * 
 * User data is a key value store API you can use to provide data from and to your server without authentication. It is the mechanism by which a user can pass information contained in a local file to an Instance at launch time.
 * 
 * The typical use case is to pass something like a shell script or a configuration file as user data.
 * 
 * For more information about [user_data](https://developers.scaleway.com/en/products/instance/api/#patch-9ef3ec)  check our documentation guide [here](https://www.scaleway.com/en/docs/compute/instances/how-to/use-boot-modes/#how-to-use-cloud-init).
 * 
 * About cloud-init documentation please check this [link](https://cloudinit.readthedocs.io/en/latest/).
 * 
 * ## Examples
 * 
 * ## Import
 * 
 * User data can be imported using the `{zone}/{key}/{server_id}`, e.g. bash
 * 
 * ```sh
 *  $ pulumi import scaleway:index/instanceUserData:InstanceUserData main fr-par-1/cloud-init/11111111-1111-1111-1111-111111111111
 * ```
 * 
 */
@ResourceType(type="scaleway:index/instanceUserData:InstanceUserData")
public class InstanceUserData extends com.pulumi.resources.CustomResource {
    /**
     * Key of the user data.
     * 
     */
    @Export(name="key", refs={String.class}, tree="[0]")
    private Output<String> key;

    /**
     * @return Key of the user data.
     * 
     */
    public Output<String> key() {
        return this.key;
    }
    /**
     * The ID of the server associated with.
     * 
     */
    @Export(name="serverId", refs={String.class}, tree="[0]")
    private Output<String> serverId;

    /**
     * @return The ID of the server associated with.
     * 
     */
    public Output<String> serverId() {
        return this.serverId;
    }
    /**
     * Value associated with your key
     * 
     */
    @Export(name="value", refs={String.class}, tree="[0]")
    private Output<String> value;

    /**
     * @return Value associated with your key
     * 
     */
    public Output<String> value() {
        return this.value;
    }
    /**
     * `zone`) The zone in which the server should be created.
     * 
     */
    @Export(name="zone", refs={String.class}, tree="[0]")
    private Output<String> zone;

    /**
     * @return `zone`) The zone in which the server should be created.
     * 
     */
    public Output<String> zone() {
        return this.zone;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public InstanceUserData(String name) {
        this(name, InstanceUserDataArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public InstanceUserData(String name, InstanceUserDataArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public InstanceUserData(String name, InstanceUserDataArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("scaleway:index/instanceUserData:InstanceUserData", name, args == null ? InstanceUserDataArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private InstanceUserData(String name, Output<String> id, @Nullable InstanceUserDataState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("scaleway:index/instanceUserData:InstanceUserData", name, state, makeResourceOptions(options, id));
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
    public static InstanceUserData get(String name, Output<String> id, @Nullable InstanceUserDataState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new InstanceUserData(name, id, state, options);
    }
}
