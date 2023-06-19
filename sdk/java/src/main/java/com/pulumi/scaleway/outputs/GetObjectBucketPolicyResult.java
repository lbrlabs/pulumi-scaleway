// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetObjectBucketPolicyResult {
    private String bucket;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The bucket&#39;s policy in JSON format.
     * 
     */
    private String policy;
    private @Nullable String projectId;
    private @Nullable String region;

    private GetObjectBucketPolicyResult() {}
    public String bucket() {
        return this.bucket;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The bucket&#39;s policy in JSON format.
     * 
     */
    public String policy() {
        return this.policy;
    }
    public Optional<String> projectId() {
        return Optional.ofNullable(this.projectId);
    }
    public Optional<String> region() {
        return Optional.ofNullable(this.region);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetObjectBucketPolicyResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String bucket;
        private String id;
        private String policy;
        private @Nullable String projectId;
        private @Nullable String region;
        public Builder() {}
        public Builder(GetObjectBucketPolicyResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.bucket = defaults.bucket;
    	      this.id = defaults.id;
    	      this.policy = defaults.policy;
    	      this.projectId = defaults.projectId;
    	      this.region = defaults.region;
        }

        @CustomType.Setter
        public Builder bucket(String bucket) {
            this.bucket = Objects.requireNonNull(bucket);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder policy(String policy) {
            this.policy = Objects.requireNonNull(policy);
            return this;
        }
        @CustomType.Setter
        public Builder projectId(@Nullable String projectId) {
            this.projectId = projectId;
            return this;
        }
        @CustomType.Setter
        public Builder region(@Nullable String region) {
            this.region = region;
            return this;
        }
        public GetObjectBucketPolicyResult build() {
            final var o = new GetObjectBucketPolicyResult();
            o.bucket = bucket;
            o.id = id;
            o.policy = policy;
            o.projectId = projectId;
            o.region = region;
            return o;
        }
    }
}
