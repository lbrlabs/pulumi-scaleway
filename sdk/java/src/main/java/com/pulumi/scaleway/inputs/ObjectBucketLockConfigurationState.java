// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.scaleway.inputs.ObjectBucketLockConfigurationRuleArgs;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ObjectBucketLockConfigurationState extends com.pulumi.resources.ResourceArgs {

    public static final ObjectBucketLockConfigurationState Empty = new ObjectBucketLockConfigurationState();

    /**
     * The bucket&#39;s name or regional ID.
     * 
     */
    @Import(name="bucket")
    private @Nullable Output<String> bucket;

    /**
     * @return The bucket&#39;s name or regional ID.
     * 
     */
    public Optional<Output<String>> bucket() {
        return Optional.ofNullable(this.bucket);
    }

    /**
     * The project_id you want to attach the resource to
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return The project_id you want to attach the resource to
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * The region you want to attach the resource to
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region you want to attach the resource to
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * Specifies the Object Lock rule for the specified object.
     * 
     */
    @Import(name="rule")
    private @Nullable Output<ObjectBucketLockConfigurationRuleArgs> rule;

    /**
     * @return Specifies the Object Lock rule for the specified object.
     * 
     */
    public Optional<Output<ObjectBucketLockConfigurationRuleArgs>> rule() {
        return Optional.ofNullable(this.rule);
    }

    private ObjectBucketLockConfigurationState() {}

    private ObjectBucketLockConfigurationState(ObjectBucketLockConfigurationState $) {
        this.bucket = $.bucket;
        this.projectId = $.projectId;
        this.region = $.region;
        this.rule = $.rule;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ObjectBucketLockConfigurationState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ObjectBucketLockConfigurationState $;

        public Builder() {
            $ = new ObjectBucketLockConfigurationState();
        }

        public Builder(ObjectBucketLockConfigurationState defaults) {
            $ = new ObjectBucketLockConfigurationState(Objects.requireNonNull(defaults));
        }

        /**
         * @param bucket The bucket&#39;s name or regional ID.
         * 
         * @return builder
         * 
         */
        public Builder bucket(@Nullable Output<String> bucket) {
            $.bucket = bucket;
            return this;
        }

        /**
         * @param bucket The bucket&#39;s name or regional ID.
         * 
         * @return builder
         * 
         */
        public Builder bucket(String bucket) {
            return bucket(Output.of(bucket));
        }

        /**
         * @param projectId The project_id you want to attach the resource to
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The project_id you want to attach the resource to
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param region The region you want to attach the resource to
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The region you want to attach the resource to
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param rule Specifies the Object Lock rule for the specified object.
         * 
         * @return builder
         * 
         */
        public Builder rule(@Nullable Output<ObjectBucketLockConfigurationRuleArgs> rule) {
            $.rule = rule;
            return this;
        }

        /**
         * @param rule Specifies the Object Lock rule for the specified object.
         * 
         * @return builder
         * 
         */
        public Builder rule(ObjectBucketLockConfigurationRuleArgs rule) {
            return rule(Output.of(rule));
        }

        public ObjectBucketLockConfigurationState build() {
            return $;
        }
    }

}
