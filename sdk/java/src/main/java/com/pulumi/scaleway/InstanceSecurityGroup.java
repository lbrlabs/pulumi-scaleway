// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.scaleway.InstanceSecurityGroupArgs;
import com.pulumi.scaleway.Utilities;
import com.pulumi.scaleway.inputs.InstanceSecurityGroupState;
import com.pulumi.scaleway.outputs.InstanceSecurityGroupInboundRule;
import com.pulumi.scaleway.outputs.InstanceSecurityGroupOutboundRule;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Import
 * 
 * Instance security group can be imported using the `{zone}/{id}`, e.g. bash
 * 
 * ```sh
 *  $ pulumi import scaleway:index/instanceSecurityGroup:InstanceSecurityGroup web fr-par-1/11111111-1111-1111-1111-111111111111
 * ```
 * 
 */
@ResourceType(type="scaleway:index/instanceSecurityGroup:InstanceSecurityGroup")
public class InstanceSecurityGroup extends com.pulumi.resources.CustomResource {
    /**
     * The description of the security group.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the security group.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Whether to block SMTP on IPv4/IPv6 (Port 25, 465, 587). Set to false will unblock SMTP if your account is authorized to. If your organization is not yet authorized to send SMTP traffic, [open a support ticket](https://console.scaleway.com/support/tickets).
     * 
     */
    @Export(name="enableDefaultSecurity", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enableDefaultSecurity;

    /**
     * @return Whether to block SMTP on IPv4/IPv6 (Port 25, 465, 587). Set to false will unblock SMTP if your account is authorized to. If your organization is not yet authorized to send SMTP traffic, [open a support ticket](https://console.scaleway.com/support/tickets).
     * 
     */
    public Output<Optional<Boolean>> enableDefaultSecurity() {
        return Codegen.optional(this.enableDefaultSecurity);
    }
    /**
     * A boolean to specify whether to use instance_security_group_rules.
     * If `external_rules` is set to `true`, `inbound_rule` and `outbound_rule` can not be set directly in the security group.
     * 
     */
    @Export(name="externalRules", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> externalRules;

    /**
     * @return A boolean to specify whether to use instance_security_group_rules.
     * If `external_rules` is set to `true`, `inbound_rule` and `outbound_rule` can not be set directly in the security group.
     * 
     */
    public Output<Optional<Boolean>> externalRules() {
        return Codegen.optional(this.externalRules);
    }
    /**
     * The default policy on incoming traffic. Possible values are: `accept` or `drop`.
     * 
     */
    @Export(name="inboundDefaultPolicy", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> inboundDefaultPolicy;

    /**
     * @return The default policy on incoming traffic. Possible values are: `accept` or `drop`.
     * 
     */
    public Output<Optional<String>> inboundDefaultPolicy() {
        return Codegen.optional(this.inboundDefaultPolicy);
    }
    /**
     * A list of inbound rule to add to the security group. (Structure is documented below.)
     * 
     */
    @Export(name="inboundRules", refs={List.class,InstanceSecurityGroupInboundRule.class}, tree="[0,1]")
    private Output</* @Nullable */ List<InstanceSecurityGroupInboundRule>> inboundRules;

    /**
     * @return A list of inbound rule to add to the security group. (Structure is documented below.)
     * 
     */
    public Output<Optional<List<InstanceSecurityGroupInboundRule>>> inboundRules() {
        return Codegen.optional(this.inboundRules);
    }
    /**
     * The name of the security group.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the security group.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The organization ID the security group is associated with.
     * 
     */
    @Export(name="organizationId", refs={String.class}, tree="[0]")
    private Output<String> organizationId;

    /**
     * @return The organization ID the security group is associated with.
     * 
     */
    public Output<String> organizationId() {
        return this.organizationId;
    }
    /**
     * The default policy on outgoing traffic. Possible values are: `accept` or `drop`.
     * 
     */
    @Export(name="outboundDefaultPolicy", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> outboundDefaultPolicy;

    /**
     * @return The default policy on outgoing traffic. Possible values are: `accept` or `drop`.
     * 
     */
    public Output<Optional<String>> outboundDefaultPolicy() {
        return Codegen.optional(this.outboundDefaultPolicy);
    }
    /**
     * A list of outbound rule to add to the security group. (Structure is documented below.)
     * 
     */
    @Export(name="outboundRules", refs={List.class,InstanceSecurityGroupOutboundRule.class}, tree="[0,1]")
    private Output</* @Nullable */ List<InstanceSecurityGroupOutboundRule>> outboundRules;

    /**
     * @return A list of outbound rule to add to the security group. (Structure is documented below.)
     * 
     */
    public Output<Optional<List<InstanceSecurityGroupOutboundRule>>> outboundRules() {
        return Codegen.optional(this.outboundRules);
    }
    /**
     * `project_id`) The ID of the project the security group is associated with.
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output<String> projectId;

    /**
     * @return `project_id`) The ID of the project the security group is associated with.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * A boolean to specify whether the security group should be stateful or not.
     * 
     */
    @Export(name="stateful", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> stateful;

    /**
     * @return A boolean to specify whether the security group should be stateful or not.
     * 
     */
    public Output<Optional<Boolean>> stateful() {
        return Codegen.optional(this.stateful);
    }
    /**
     * The tags of the security group.
     * 
     */
    @Export(name="tags", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> tags;

    /**
     * @return The tags of the security group.
     * 
     */
    public Output<Optional<List<String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * `zone`) The zone in which the security group should be created.
     * 
     */
    @Export(name="zone", refs={String.class}, tree="[0]")
    private Output<String> zone;

    /**
     * @return `zone`) The zone in which the security group should be created.
     * 
     */
    public Output<String> zone() {
        return this.zone;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public InstanceSecurityGroup(String name) {
        this(name, InstanceSecurityGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public InstanceSecurityGroup(String name, @Nullable InstanceSecurityGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public InstanceSecurityGroup(String name, @Nullable InstanceSecurityGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("scaleway:index/instanceSecurityGroup:InstanceSecurityGroup", name, args == null ? InstanceSecurityGroupArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private InstanceSecurityGroup(String name, Output<String> id, @Nullable InstanceSecurityGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("scaleway:index/instanceSecurityGroup:InstanceSecurityGroup", name, state, makeResourceOptions(options, id));
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
    public static InstanceSecurityGroup get(String name, Output<String> id, @Nullable InstanceSecurityGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new InstanceSecurityGroup(name, id, state, options);
    }
}
