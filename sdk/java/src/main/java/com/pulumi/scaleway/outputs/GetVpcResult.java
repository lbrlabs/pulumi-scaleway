// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetVpcResult {
    private String createdAt;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private @Nullable Boolean isDefault;
    private @Nullable String name;
    private String organizationId;
    private @Nullable String projectId;
    private @Nullable String region;
    private List<String> tags;
    private String updatedAt;
    private @Nullable String vpcId;

    private GetVpcResult() {}
    public String createdAt() {
        return this.createdAt;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public Optional<Boolean> isDefault() {
        return Optional.ofNullable(this.isDefault);
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
    public List<String> tags() {
        return this.tags;
    }
    public String updatedAt() {
        return this.updatedAt;
    }
    public Optional<String> vpcId() {
        return Optional.ofNullable(this.vpcId);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetVpcResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String createdAt;
        private String id;
        private @Nullable Boolean isDefault;
        private @Nullable String name;
        private String organizationId;
        private @Nullable String projectId;
        private @Nullable String region;
        private List<String> tags;
        private String updatedAt;
        private @Nullable String vpcId;
        public Builder() {}
        public Builder(GetVpcResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.createdAt = defaults.createdAt;
    	      this.id = defaults.id;
    	      this.isDefault = defaults.isDefault;
    	      this.name = defaults.name;
    	      this.organizationId = defaults.organizationId;
    	      this.projectId = defaults.projectId;
    	      this.region = defaults.region;
    	      this.tags = defaults.tags;
    	      this.updatedAt = defaults.updatedAt;
    	      this.vpcId = defaults.vpcId;
        }

        @CustomType.Setter
        public Builder createdAt(String createdAt) {
            this.createdAt = Objects.requireNonNull(createdAt);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder isDefault(@Nullable Boolean isDefault) {
            this.isDefault = isDefault;
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
        public Builder vpcId(@Nullable String vpcId) {
            this.vpcId = vpcId;
            return this;
        }
        public GetVpcResult build() {
            final var o = new GetVpcResult();
            o.createdAt = createdAt;
            o.id = id;
            o.isDefault = isDefault;
            o.name = name;
            o.organizationId = organizationId;
            o.projectId = projectId;
            o.region = region;
            o.tags = tags;
            o.updatedAt = updatedAt;
            o.vpcId = vpcId;
            return o;
        }
    }
}
