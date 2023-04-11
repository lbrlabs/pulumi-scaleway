// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CockpitArgs extends com.pulumi.resources.ResourceArgs {

    public static final CockpitArgs Empty = new CockpitArgs();

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

    private CockpitArgs() {}

    private CockpitArgs(CockpitArgs $) {
        this.projectId = $.projectId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CockpitArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CockpitArgs $;

        public Builder() {
            $ = new CockpitArgs();
        }

        public Builder(CockpitArgs defaults) {
            $ = new CockpitArgs(Objects.requireNonNull(defaults));
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

        public CockpitArgs build() {
            return $;
        }
    }

}
