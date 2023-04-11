// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.scaleway.inputs.CockpitTokenScopesArgs;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CockpitTokenArgs extends com.pulumi.resources.ResourceArgs {

    public static final CockpitTokenArgs Empty = new CockpitTokenArgs();

    /**
     * The name of the token
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the token
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * `project_id`) The ID of the project the cockpit is associated with.
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return `project_id`) The ID of the project the cockpit is associated with.
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * Allowed scopes
     * 
     */
    @Import(name="scopes")
    private @Nullable Output<CockpitTokenScopesArgs> scopes;

    /**
     * @return Allowed scopes
     * 
     */
    public Optional<Output<CockpitTokenScopesArgs>> scopes() {
        return Optional.ofNullable(this.scopes);
    }

    private CockpitTokenArgs() {}

    private CockpitTokenArgs(CockpitTokenArgs $) {
        this.name = $.name;
        this.projectId = $.projectId;
        this.scopes = $.scopes;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CockpitTokenArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CockpitTokenArgs $;

        public Builder() {
            $ = new CockpitTokenArgs();
        }

        public Builder(CockpitTokenArgs defaults) {
            $ = new CockpitTokenArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name The name of the token
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the token
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param projectId `project_id`) The ID of the project the cockpit is associated with.
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId `project_id`) The ID of the project the cockpit is associated with.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param scopes Allowed scopes
         * 
         * @return builder
         * 
         */
        public Builder scopes(@Nullable Output<CockpitTokenScopesArgs> scopes) {
            $.scopes = scopes;
            return this;
        }

        /**
         * @param scopes Allowed scopes
         * 
         * @return builder
         * 
         */
        public Builder scopes(CockpitTokenScopesArgs scopes) {
            return scopes(Output.of(scopes));
        }

        public CockpitTokenArgs build() {
            return $;
        }
    }

}