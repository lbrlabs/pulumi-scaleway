// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetSecretResult {
    private String createdAt;
    private String description;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private @Nullable String name;
    private String organizationId;
    private @Nullable String projectId;
    private @Nullable String region;
    private @Nullable String secretId;
    private String status;
    private List<String> tags;
    private String updatedAt;
    private Integer versionCount;

    private GetSecretResult() {}
    public String createdAt() {
        return this.createdAt;
    }
    public String description() {
        return this.description;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    public String organizationId() {
        return this.organizationId;
    }
    public Optional<String> projectId() {
        return Optional.ofNullable(this.projectId);
    }
    public Optional<String> region() {
        return Optional.ofNullable(this.region);
    }
    public Optional<String> secretId() {
        return Optional.ofNullable(this.secretId);
    }
    public String status() {
        return this.status;
    }
    public List<String> tags() {
        return this.tags;
    }
    public String updatedAt() {
        return this.updatedAt;
    }
    public Integer versionCount() {
        return this.versionCount;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetSecretResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String createdAt;
        private String description;
        private String id;
        private @Nullable String name;
        private String organizationId;
        private @Nullable String projectId;
        private @Nullable String region;
        private @Nullable String secretId;
        private String status;
        private List<String> tags;
        private String updatedAt;
        private Integer versionCount;
        public Builder() {}
        public Builder(GetSecretResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.createdAt = defaults.createdAt;
    	      this.description = defaults.description;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.organizationId = defaults.organizationId;
    	      this.projectId = defaults.projectId;
    	      this.region = defaults.region;
    	      this.secretId = defaults.secretId;
    	      this.status = defaults.status;
    	      this.tags = defaults.tags;
    	      this.updatedAt = defaults.updatedAt;
    	      this.versionCount = defaults.versionCount;
        }

        @CustomType.Setter
        public Builder createdAt(String createdAt) {
            this.createdAt = Objects.requireNonNull(createdAt);
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder name(@Nullable String name) {
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder organizationId(String organizationId) {
            this.organizationId = Objects.requireNonNull(organizationId);
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
        @CustomType.Setter
        public Builder secretId(@Nullable String secretId) {
            this.secretId = secretId;
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        @CustomType.Setter
        public Builder tags(List<String> tags) {
            this.tags = Objects.requireNonNull(tags);
            return this;
        }
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }
        @CustomType.Setter
        public Builder updatedAt(String updatedAt) {
            this.updatedAt = Objects.requireNonNull(updatedAt);
            return this;
        }
        @CustomType.Setter
        public Builder versionCount(Integer versionCount) {
            this.versionCount = Objects.requireNonNull(versionCount);
            return this;
        }
        public GetSecretResult build() {
            final var o = new GetSecretResult();
            o.createdAt = createdAt;
            o.description = description;
            o.id = id;
            o.name = name;
            o.organizationId = organizationId;
            o.projectId = projectId;
            o.region = region;
            o.secretId = secretId;
            o.status = status;
            o.tags = tags;
            o.updatedAt = updatedAt;
            o.versionCount = versionCount;
            return o;
        }
    }
}
