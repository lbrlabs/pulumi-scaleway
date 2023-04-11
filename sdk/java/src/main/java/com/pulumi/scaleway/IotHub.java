// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.scaleway.IotHubArgs;
import com.pulumi.scaleway.Utilities;
import com.pulumi.scaleway.inputs.IotHubState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Import
 * 
 * IoT Hubs can be imported using the `{region}/{id}`, e.g. bash
 * 
 * ```sh
 *  $ pulumi import scaleway:index/iotHub:IotHub hub01 fr-par/11111111-1111-1111-1111-111111111111
 * ```
 * 
 */
@ResourceType(type="scaleway:index/iotHub:IotHub")
public class IotHub extends com.pulumi.resources.CustomResource {
    /**
     * The current number of connected devices in the Hub.
     * 
     */
    @Export(name="connectedDeviceCount", refs={Integer.class}, tree="[0]")
    private Output<Integer> connectedDeviceCount;

    /**
     * @return The current number of connected devices in the Hub.
     * 
     */
    public Output<Integer> connectedDeviceCount() {
        return this.connectedDeviceCount;
    }
    /**
     * The date and time the Hub was created.
     * 
     */
    @Export(name="createdAt", refs={String.class}, tree="[0]")
    private Output<String> createdAt;

    /**
     * @return The date and time the Hub was created.
     * 
     */
    public Output<String> createdAt() {
        return this.createdAt;
    }
    /**
     * Wether to enable the device auto provisioning or not
     * 
     */
    @Export(name="deviceAutoProvisioning", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> deviceAutoProvisioning;

    /**
     * @return Wether to enable the device auto provisioning or not
     * 
     */
    public Output<Optional<Boolean>> deviceAutoProvisioning() {
        return Codegen.optional(this.deviceAutoProvisioning);
    }
    /**
     * The number of registered devices in the Hub.
     * 
     */
    @Export(name="deviceCount", refs={Integer.class}, tree="[0]")
    private Output<Integer> deviceCount;

    /**
     * @return The number of registered devices in the Hub.
     * 
     */
    public Output<Integer> deviceCount() {
        return this.deviceCount;
    }
    /**
     * Whether to enable the hub events or not
     * 
     */
    @Export(name="disableEvents", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> disableEvents;

    /**
     * @return Whether to enable the hub events or not
     * 
     */
    public Output<Optional<Boolean>> disableEvents() {
        return Codegen.optional(this.disableEvents);
    }
    /**
     * Wether the IoT Hub instance should be enabled or not.
     * 
     */
    @Export(name="enabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enabled;

    /**
     * @return Wether the IoT Hub instance should be enabled or not.
     * 
     */
    public Output<Optional<Boolean>> enabled() {
        return Codegen.optional(this.enabled);
    }
    /**
     * The MQTT network endpoint to connect MQTT devices to.
     * 
     */
    @Export(name="endpoint", refs={String.class}, tree="[0]")
    private Output<String> endpoint;

    /**
     * @return The MQTT network endpoint to connect MQTT devices to.
     * 
     */
    public Output<String> endpoint() {
        return this.endpoint;
    }
    /**
     * Topic prefix for the hub events
     * 
     */
    @Export(name="eventsTopicPrefix", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> eventsTopicPrefix;

    /**
     * @return Topic prefix for the hub events
     * 
     */
    public Output<Optional<String>> eventsTopicPrefix() {
        return Codegen.optional(this.eventsTopicPrefix);
    }
    /**
     * Custom user provided certificate authority
     * 
     */
    @Export(name="hubCa", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> hubCa;

    /**
     * @return Custom user provided certificate authority
     * 
     */
    public Output<Optional<String>> hubCa() {
        return Codegen.optional(this.hubCa);
    }
    /**
     * Challenge certificate for the user provided hub CA
     * 
     */
    @Export(name="hubCaChallenge", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> hubCaChallenge;

    /**
     * @return Challenge certificate for the user provided hub CA
     * 
     */
    public Output<Optional<String>> hubCaChallenge() {
        return Codegen.optional(this.hubCaChallenge);
    }
    /**
     * The name of the IoT Hub instance you want to create (e.g. `my-hub`).
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the IoT Hub instance you want to create (e.g. `my-hub`).
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The organization_id you want to attach the resource to
     * 
     */
    @Export(name="organizationId", refs={String.class}, tree="[0]")
    private Output<String> organizationId;

    /**
     * @return The organization_id you want to attach the resource to
     * 
     */
    public Output<String> organizationId() {
        return this.organizationId;
    }
    /**
     * Product plan to create the hub, see documentation for available product plans (e.g. `plan_shared`)
     * 
     */
    @Export(name="productPlan", refs={String.class}, tree="[0]")
    private Output<String> productPlan;

    /**
     * @return Product plan to create the hub, see documentation for available product plans (e.g. `plan_shared`)
     * 
     */
    public Output<String> productPlan() {
        return this.productPlan;
    }
    /**
     * `project_id`) The ID of the project the IoT Hub Instance is associated with.
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output<String> projectId;

    /**
     * @return `project_id`) The ID of the project the IoT Hub Instance is associated with.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * `region`) The region in which the Database Instance should be created.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return `region`) The region in which the Database Instance should be created.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The current status of the Hub.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The current status of the Hub.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The date and time the Hub resource was updated.
     * 
     */
    @Export(name="updatedAt", refs={String.class}, tree="[0]")
    private Output<String> updatedAt;

    /**
     * @return The date and time the Hub resource was updated.
     * 
     */
    public Output<String> updatedAt() {
        return this.updatedAt;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public IotHub(String name) {
        this(name, IotHubArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public IotHub(String name, IotHubArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public IotHub(String name, IotHubArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("scaleway:index/iotHub:IotHub", name, args == null ? IotHubArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private IotHub(String name, Output<String> id, @Nullable IotHubState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("scaleway:index/iotHub:IotHub", name, state, makeResourceOptions(options, id));
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
    public static IotHub get(String name, Output<String> id, @Nullable IotHubState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new IotHub(name, id, state, options);
    }
}
