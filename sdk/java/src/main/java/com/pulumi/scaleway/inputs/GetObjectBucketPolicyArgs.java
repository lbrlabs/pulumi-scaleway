// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetObjectBucketPolicyArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetObjectBucketPolicyArgs Empty = new GetObjectBucketPolicyArgs();

    /**
     * The bucket name.
     * 
     */
    @Import(name="bucket", required=true)
    private Output<String> bucket;

    /**
     * @return The bucket name.
     * 
     */
    public Output<String> bucket() {
        return this.bucket;
    }

    /**
     * `project_id`) The ID of the project the bucket is associated with.
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return `project_id`) The ID of the project the bucket is associated with.
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * `region`) The region in which the Object Storage exists.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return `region`) The region in which the Object Storage exists.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    private GetObjectBucketPolicyArgs() {}

    private GetObjectBucketPolicyArgs(GetObjectBucketPolicyArgs $) {
        this.bucket = $.bucket;
        this.projectId = $.projectId;
        this.region = $.region;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetObjectBucketPolicyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetObjectBucketPolicyArgs $;

        public Builder() {
            $ = new GetObjectBucketPolicyArgs();
        }

        public Builder(GetObjectBucketPolicyArgs defaults) {
            $ = new GetObjectBucketPolicyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param bucket The bucket name.
         * 
         * @return builder
         * 
         */
        public Builder bucket(Output<String> bucket) {
            $.bucket = bucket;
            return this;
        }

        /**
         * @param bucket The bucket name.
         * 
         * @return builder
         * 
         */
        public Builder bucket(String bucket) {
            return bucket(Output.of(bucket));
        }

        /**
         * @param projectId `project_id`) The ID of the project the bucket is associated with.
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId `project_id`) The ID of the project the bucket is associated with.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param region `region`) The region in which the Object Storage exists.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region `region`) The region in which the Object Storage exists.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        public GetObjectBucketPolicyArgs build() {
            $.bucket = Objects.requireNonNull($.bucket, "expected parameter 'bucket' to be non-null");
            return $;
        }
    }

}
