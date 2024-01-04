// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class IamApplicationState extends com.pulumi.resources.ResourceArgs {

    public static final IamApplicationState Empty = new IamApplicationState();

    /**
     * The date and time of the creation of the application.
     * 
     */
    @Import(name="createdAt")
    private @Nullable Output<String> createdAt;

    /**
     * @return The date and time of the creation of the application.
     * 
     */
    public Optional<Output<String>> createdAt() {
        return Optional.ofNullable(this.createdAt);
    }

    /**
     * The description of the iam application.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the iam application.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Whether the application is editable.
     * 
     */
    @Import(name="editable")
    private @Nullable Output<Boolean> editable;

    /**
     * @return Whether the application is editable.
     * 
     */
    public Optional<Output<Boolean>> editable() {
        return Optional.ofNullable(this.editable);
    }

    /**
     * The name of the iam application.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the iam application.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * `organization_id`) The ID of the organization the application is associated with.
     * 
     */
    @Import(name="organizationId")
    private @Nullable Output<String> organizationId;

    /**
     * @return `organization_id`) The ID of the organization the application is associated with.
     * 
     */
    public Optional<Output<String>> organizationId() {
        return Optional.ofNullable(this.organizationId);
    }

    /**
     * The tags associated with the application.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<List<String>> tags;

    /**
     * @return The tags associated with the application.
     * 
     */
    public Optional<Output<List<String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * The date and time of the last update of the application.
     * 
     */
    @Import(name="updatedAt")
    private @Nullable Output<String> updatedAt;

    /**
     * @return The date and time of the last update of the application.
     * 
     */
    public Optional<Output<String>> updatedAt() {
        return Optional.ofNullable(this.updatedAt);
    }

    private IamApplicationState() {}

    private IamApplicationState(IamApplicationState $) {
        this.createdAt = $.createdAt;
        this.description = $.description;
        this.editable = $.editable;
        this.name = $.name;
        this.organizationId = $.organizationId;
        this.tags = $.tags;
        this.updatedAt = $.updatedAt;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IamApplicationState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IamApplicationState $;

        public Builder() {
            $ = new IamApplicationState();
        }

        public Builder(IamApplicationState defaults) {
            $ = new IamApplicationState(Objects.requireNonNull(defaults));
        }

        /**
         * @param createdAt The date and time of the creation of the application.
         * 
         * @return builder
         * 
         */
        public Builder createdAt(@Nullable Output<String> createdAt) {
            $.createdAt = createdAt;
            return this;
        }

        /**
         * @param createdAt The date and time of the creation of the application.
         * 
         * @return builder
         * 
         */
        public Builder createdAt(String createdAt) {
            return createdAt(Output.of(createdAt));
        }

        /**
         * @param description The description of the iam application.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the iam application.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param editable Whether the application is editable.
         * 
         * @return builder
         * 
         */
        public Builder editable(@Nullable Output<Boolean> editable) {
            $.editable = editable;
            return this;
        }

        /**
         * @param editable Whether the application is editable.
         * 
         * @return builder
         * 
         */
        public Builder editable(Boolean editable) {
            return editable(Output.of(editable));
        }

        /**
         * @param name The name of the iam application.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the iam application.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param organizationId `organization_id`) The ID of the organization the application is associated with.
         * 
         * @return builder
         * 
         */
        public Builder organizationId(@Nullable Output<String> organizationId) {
            $.organizationId = organizationId;
            return this;
        }

        /**
         * @param organizationId `organization_id`) The ID of the organization the application is associated with.
         * 
         * @return builder
         * 
         */
        public Builder organizationId(String organizationId) {
            return organizationId(Output.of(organizationId));
        }

        /**
         * @param tags The tags associated with the application.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<List<String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags The tags associated with the application.
         * 
         * @return builder
         * 
         */
        public Builder tags(List<String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tags The tags associated with the application.
         * 
         * @return builder
         * 
         */
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }

        /**
         * @param updatedAt The date and time of the last update of the application.
         * 
         * @return builder
         * 
         */
        public Builder updatedAt(@Nullable Output<String> updatedAt) {
            $.updatedAt = updatedAt;
            return this;
        }

        /**
         * @param updatedAt The date and time of the last update of the application.
         * 
         * @return builder
         * 
         */
        public Builder updatedAt(String updatedAt) {
            return updatedAt(Output.of(updatedAt));
        }

        public IamApplicationState build() {
            return $;
        }
    }

}
