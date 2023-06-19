// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ContainerNamespaceState extends com.pulumi.resources.ResourceArgs {

    public static final ContainerNamespaceState Empty = new ContainerNamespaceState();

    /**
     * The description of the namespace.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the namespace.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Destroy registry on deletion
     * 
     * @deprecated
     * Registry namespace is automatically destroyed with namespace
     * 
     */
    @Deprecated /* Registry namespace is automatically destroyed with namespace */
    @Import(name="destroyRegistry")
    private @Nullable Output<Boolean> destroyRegistry;

    /**
     * @return Destroy registry on deletion
     * 
     * @deprecated
     * Registry namespace is automatically destroyed with namespace
     * 
     */
    @Deprecated /* Registry namespace is automatically destroyed with namespace */
    public Optional<Output<Boolean>> destroyRegistry() {
        return Optional.ofNullable(this.destroyRegistry);
    }

    /**
     * The environment variables of the namespace.
     * 
     */
    @Import(name="environmentVariables")
    private @Nullable Output<Map<String,String>> environmentVariables;

    /**
     * @return The environment variables of the namespace.
     * 
     */
    public Optional<Output<Map<String,String>>> environmentVariables() {
        return Optional.ofNullable(this.environmentVariables);
    }

    /**
     * The unique name of the container namespace.
     * 
     * &gt; **Important** Updates to `name` will recreate the namespace.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The unique name of the container namespace.
     * 
     * &gt; **Important** Updates to `name` will recreate the namespace.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The organization ID the namespace is associated with.
     * 
     */
    @Import(name="organizationId")
    private @Nullable Output<String> organizationId;

    /**
     * @return The organization ID the namespace is associated with.
     * 
     */
    public Optional<Output<String>> organizationId() {
        return Optional.ofNullable(this.organizationId);
    }

    /**
     * `project_id`) The ID of the project the namespace is associated with.
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return `project_id`) The ID of the project the namespace is associated with.
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * `region`). The region in which the namespace should be created.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return `region`). The region in which the namespace should be created.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * The registry endpoint of the namespace.
     * 
     */
    @Import(name="registryEndpoint")
    private @Nullable Output<String> registryEndpoint;

    /**
     * @return The registry endpoint of the namespace.
     * 
     */
    public Optional<Output<String>> registryEndpoint() {
        return Optional.ofNullable(this.registryEndpoint);
    }

    /**
     * The registry namespace ID of the namespace.
     * 
     */
    @Import(name="registryNamespaceId")
    private @Nullable Output<String> registryNamespaceId;

    /**
     * @return The registry namespace ID of the namespace.
     * 
     */
    public Optional<Output<String>> registryNamespaceId() {
        return Optional.ofNullable(this.registryNamespaceId);
    }

    /**
     * The secret environment variables of the namespace.
     * 
     */
    @Import(name="secretEnvironmentVariables")
    private @Nullable Output<Map<String,String>> secretEnvironmentVariables;

    /**
     * @return The secret environment variables of the namespace.
     * 
     */
    public Optional<Output<Map<String,String>>> secretEnvironmentVariables() {
        return Optional.ofNullable(this.secretEnvironmentVariables);
    }

    private ContainerNamespaceState() {}

    private ContainerNamespaceState(ContainerNamespaceState $) {
        this.description = $.description;
        this.destroyRegistry = $.destroyRegistry;
        this.environmentVariables = $.environmentVariables;
        this.name = $.name;
        this.organizationId = $.organizationId;
        this.projectId = $.projectId;
        this.region = $.region;
        this.registryEndpoint = $.registryEndpoint;
        this.registryNamespaceId = $.registryNamespaceId;
        this.secretEnvironmentVariables = $.secretEnvironmentVariables;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ContainerNamespaceState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ContainerNamespaceState $;

        public Builder() {
            $ = new ContainerNamespaceState();
        }

        public Builder(ContainerNamespaceState defaults) {
            $ = new ContainerNamespaceState(Objects.requireNonNull(defaults));
        }

        /**
         * @param description The description of the namespace.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the namespace.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param destroyRegistry Destroy registry on deletion
         * 
         * @return builder
         * 
         * @deprecated
         * Registry namespace is automatically destroyed with namespace
         * 
         */
        @Deprecated /* Registry namespace is automatically destroyed with namespace */
        public Builder destroyRegistry(@Nullable Output<Boolean> destroyRegistry) {
            $.destroyRegistry = destroyRegistry;
            return this;
        }

        /**
         * @param destroyRegistry Destroy registry on deletion
         * 
         * @return builder
         * 
         * @deprecated
         * Registry namespace is automatically destroyed with namespace
         * 
         */
        @Deprecated /* Registry namespace is automatically destroyed with namespace */
        public Builder destroyRegistry(Boolean destroyRegistry) {
            return destroyRegistry(Output.of(destroyRegistry));
        }

        /**
         * @param environmentVariables The environment variables of the namespace.
         * 
         * @return builder
         * 
         */
        public Builder environmentVariables(@Nullable Output<Map<String,String>> environmentVariables) {
            $.environmentVariables = environmentVariables;
            return this;
        }

        /**
         * @param environmentVariables The environment variables of the namespace.
         * 
         * @return builder
         * 
         */
        public Builder environmentVariables(Map<String,String> environmentVariables) {
            return environmentVariables(Output.of(environmentVariables));
        }

        /**
         * @param name The unique name of the container namespace.
         * 
         * &gt; **Important** Updates to `name` will recreate the namespace.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The unique name of the container namespace.
         * 
         * &gt; **Important** Updates to `name` will recreate the namespace.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param organizationId The organization ID the namespace is associated with.
         * 
         * @return builder
         * 
         */
        public Builder organizationId(@Nullable Output<String> organizationId) {
            $.organizationId = organizationId;
            return this;
        }

        /**
         * @param organizationId The organization ID the namespace is associated with.
         * 
         * @return builder
         * 
         */
        public Builder organizationId(String organizationId) {
            return organizationId(Output.of(organizationId));
        }

        /**
         * @param projectId `project_id`) The ID of the project the namespace is associated with.
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId `project_id`) The ID of the project the namespace is associated with.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param region `region`). The region in which the namespace should be created.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region `region`). The region in which the namespace should be created.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param registryEndpoint The registry endpoint of the namespace.
         * 
         * @return builder
         * 
         */
        public Builder registryEndpoint(@Nullable Output<String> registryEndpoint) {
            $.registryEndpoint = registryEndpoint;
            return this;
        }

        /**
         * @param registryEndpoint The registry endpoint of the namespace.
         * 
         * @return builder
         * 
         */
        public Builder registryEndpoint(String registryEndpoint) {
            return registryEndpoint(Output.of(registryEndpoint));
        }

        /**
         * @param registryNamespaceId The registry namespace ID of the namespace.
         * 
         * @return builder
         * 
         */
        public Builder registryNamespaceId(@Nullable Output<String> registryNamespaceId) {
            $.registryNamespaceId = registryNamespaceId;
            return this;
        }

        /**
         * @param registryNamespaceId The registry namespace ID of the namespace.
         * 
         * @return builder
         * 
         */
        public Builder registryNamespaceId(String registryNamespaceId) {
            return registryNamespaceId(Output.of(registryNamespaceId));
        }

        /**
         * @param secretEnvironmentVariables The secret environment variables of the namespace.
         * 
         * @return builder
         * 
         */
        public Builder secretEnvironmentVariables(@Nullable Output<Map<String,String>> secretEnvironmentVariables) {
            $.secretEnvironmentVariables = secretEnvironmentVariables;
            return this;
        }

        /**
         * @param secretEnvironmentVariables The secret environment variables of the namespace.
         * 
         * @return builder
         * 
         */
        public Builder secretEnvironmentVariables(Map<String,String> secretEnvironmentVariables) {
            return secretEnvironmentVariables(Output.of(secretEnvironmentVariables));
        }

        public ContainerNamespaceState build() {
            return $;
        }
    }

}
