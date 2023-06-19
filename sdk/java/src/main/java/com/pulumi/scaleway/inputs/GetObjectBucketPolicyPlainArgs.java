// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetObjectBucketPolicyPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetObjectBucketPolicyPlainArgs Empty = new GetObjectBucketPolicyPlainArgs();

    /**
     * The bucket name.
     * 
     */
    @Import(name="bucket", required=true)
    private String bucket;

    /**
     * @return The bucket name.
     * 
     */
    public String bucket() {
        return this.bucket;
    }

    /**
     * `project_id`) The ID of the project the bucket is associated with.
     * 
     */
    @Import(name="projectId")
    private @Nullable String projectId;

    /**
     * @return `project_id`) The ID of the project the bucket is associated with.
     * 
     */
    public Optional<String> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * `region`) The region in which the Object Storage exists.
     * 
     */
    @Import(name="region")
    private @Nullable String region;

    /**
     * @return `region`) The region in which the Object Storage exists.
     * 
     */
    public Optional<String> region() {
        return Optional.ofNullable(this.region);
    }

    private GetObjectBucketPolicyPlainArgs() {}

    private GetObjectBucketPolicyPlainArgs(GetObjectBucketPolicyPlainArgs $) {
        this.bucket = $.bucket;
        this.projectId = $.projectId;
        this.region = $.region;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetObjectBucketPolicyPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetObjectBucketPolicyPlainArgs $;

        public Builder() {
            $ = new GetObjectBucketPolicyPlainArgs();
        }

        public Builder(GetObjectBucketPolicyPlainArgs defaults) {
            $ = new GetObjectBucketPolicyPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param bucket The bucket name.
         * 
         * @return builder
         * 
         */
        public Builder bucket(String bucket) {
            $.bucket = bucket;
            return this;
        }

        /**
         * @param projectId `project_id`) The ID of the project the bucket is associated with.
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable String projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param region `region`) The region in which the Object Storage exists.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable String region) {
            $.region = region;
            return this;
        }

        public GetObjectBucketPolicyPlainArgs build() {
            $.bucket = Objects.requireNonNull($.bucket, "expected parameter 'bucket' to be non-null");
            return $;
        }
    }

}