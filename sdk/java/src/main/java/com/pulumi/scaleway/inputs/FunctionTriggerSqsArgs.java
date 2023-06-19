// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class FunctionTriggerSqsArgs extends com.pulumi.resources.ResourceArgs {

    public static final FunctionTriggerSqsArgs Empty = new FunctionTriggerSqsArgs();

    /**
     * ID of the mnq namespace
     * 
     */
    @Import(name="namespaceId", required=true)
    private Output<String> namespaceId;

    /**
     * @return ID of the mnq namespace
     * 
     */
    public Output<String> namespaceId() {
        return this.namespaceId;
    }

    /**
     * ID of the project that contain the mnq namespace, defaults to provider&#39;s project
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return ID of the project that contain the mnq namespace, defaults to provider&#39;s project
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * Name of the queue
     * 
     */
    @Import(name="queue", required=true)
    private Output<String> queue;

    /**
     * @return Name of the queue
     * 
     */
    public Output<String> queue() {
        return this.queue;
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

    private FunctionTriggerSqsArgs() {}

    private FunctionTriggerSqsArgs(FunctionTriggerSqsArgs $) {
        this.namespaceId = $.namespaceId;
        this.projectId = $.projectId;
        this.queue = $.queue;
        this.region = $.region;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FunctionTriggerSqsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FunctionTriggerSqsArgs $;

        public Builder() {
            $ = new FunctionTriggerSqsArgs();
        }

        public Builder(FunctionTriggerSqsArgs defaults) {
            $ = new FunctionTriggerSqsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param namespaceId ID of the mnq namespace
         * 
         * @return builder
         * 
         */
        public Builder namespaceId(Output<String> namespaceId) {
            $.namespaceId = namespaceId;
            return this;
        }

        /**
         * @param namespaceId ID of the mnq namespace
         * 
         * @return builder
         * 
         */
        public Builder namespaceId(String namespaceId) {
            return namespaceId(Output.of(namespaceId));
        }

        /**
         * @param projectId ID of the project that contain the mnq namespace, defaults to provider&#39;s project
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId ID of the project that contain the mnq namespace, defaults to provider&#39;s project
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param queue Name of the queue
         * 
         * @return builder
         * 
         */
        public Builder queue(Output<String> queue) {
            $.queue = queue;
            return this;
        }

        /**
         * @param queue Name of the queue
         * 
         * @return builder
         * 
         */
        public Builder queue(String queue) {
            return queue(Output.of(queue));
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

        public FunctionTriggerSqsArgs build() {
            $.namespaceId = Objects.requireNonNull($.namespaceId, "expected parameter 'namespaceId' to be non-null");
            $.queue = Objects.requireNonNull($.queue, "expected parameter 'queue' to be non-null");
            return $;
        }
    }

}
