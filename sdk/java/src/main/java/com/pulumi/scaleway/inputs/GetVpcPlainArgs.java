// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetVpcPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetVpcPlainArgs Empty = new GetVpcPlainArgs();

    /**
     * To get default VPC&#39;s information.
     * 
     */
    @Import(name="isDefault")
    private @Nullable Boolean isDefault;

    /**
     * @return To get default VPC&#39;s information.
     * 
     */
    public Optional<Boolean> isDefault() {
        return Optional.ofNullable(this.isDefault);
    }

    /**
     * Name of the VPC. One of `name` and `vpc_id` should be specified.
     * 
     */
    @Import(name="name")
    private @Nullable String name;

    /**
     * @return Name of the VPC. One of `name` and `vpc_id` should be specified.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The ID of the organization the VPC is associated with.
     * 
     */
    @Import(name="organizationId")
    private @Nullable String organizationId;

    /**
     * @return The ID of the organization the VPC is associated with.
     * 
     */
    public Optional<String> organizationId() {
        return Optional.ofNullable(this.organizationId);
    }

    /**
     * `project_id`) The ID of the project the VPC is associated with.
     * 
     */
    @Import(name="projectId")
    private @Nullable String projectId;

    /**
     * @return `project_id`) The ID of the project the VPC is associated with.
     * 
     */
    public Optional<String> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    @Import(name="region")
    private @Nullable String region;

    public Optional<String> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * ID of the VPC. One of `name` and `vpc_id` should be specified.
     * 
     */
    @Import(name="vpcId")
    private @Nullable String vpcId;

    /**
     * @return ID of the VPC. One of `name` and `vpc_id` should be specified.
     * 
     */
    public Optional<String> vpcId() {
        return Optional.ofNullable(this.vpcId);
    }

    private GetVpcPlainArgs() {}

    private GetVpcPlainArgs(GetVpcPlainArgs $) {
        this.isDefault = $.isDefault;
        this.name = $.name;
        this.organizationId = $.organizationId;
        this.projectId = $.projectId;
        this.region = $.region;
        this.vpcId = $.vpcId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetVpcPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetVpcPlainArgs $;

        public Builder() {
            $ = new GetVpcPlainArgs();
        }

        public Builder(GetVpcPlainArgs defaults) {
            $ = new GetVpcPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param isDefault To get default VPC&#39;s information.
         * 
         * @return builder
         * 
         */
        public Builder isDefault(@Nullable Boolean isDefault) {
            $.isDefault = isDefault;
            return this;
        }

        /**
         * @param name Name of the VPC. One of `name` and `vpc_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable String name) {
            $.name = name;
            return this;
        }

        /**
         * @param organizationId The ID of the organization the VPC is associated with.
         * 
         * @return builder
         * 
         */
        public Builder organizationId(@Nullable String organizationId) {
            $.organizationId = organizationId;
            return this;
        }

        /**
         * @param projectId `project_id`) The ID of the project the VPC is associated with.
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable String projectId) {
            $.projectId = projectId;
            return this;
        }

        public Builder region(@Nullable String region) {
            $.region = region;
            return this;
        }

        /**
         * @param vpcId ID of the VPC. One of `name` and `vpc_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(@Nullable String vpcId) {
            $.vpcId = vpcId;
            return this;
        }

        public GetVpcPlainArgs build() {
            return $;
        }
    }

}
