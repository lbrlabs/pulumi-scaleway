// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetIotHubResult {
    private Integer connectedDeviceCount;
    private String createdAt;
    private Boolean deviceAutoProvisioning;
    private Integer deviceCount;
    private Boolean disableEvents;
    private Boolean enabled;
    private String endpoint;
    private String eventsTopicPrefix;
    private String hubCa;
    private String hubCaChallenge;
    private @Nullable String hubId;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private @Nullable String name;
    private String organizationId;
    private String productPlan;
    private String projectId;
    private @Nullable String region;
    private String status;
    private String updatedAt;

    private GetIotHubResult() {}
    public Integer connectedDeviceCount() {
        return this.connectedDeviceCount;
    }
    public String createdAt() {
        return this.createdAt;
    }
    public Boolean deviceAutoProvisioning() {
        return this.deviceAutoProvisioning;
    }
    public Integer deviceCount() {
        return this.deviceCount;
    }
    public Boolean disableEvents() {
        return this.disableEvents;
    }
    public Boolean enabled() {
        return this.enabled;
    }
    public String endpoint() {
        return this.endpoint;
    }
    public String eventsTopicPrefix() {
        return this.eventsTopicPrefix;
    }
    public String hubCa() {
        return this.hubCa;
    }
    public String hubCaChallenge() {
        return this.hubCaChallenge;
    }
    public Optional<String> hubId() {
        return Optional.ofNullable(this.hubId);
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
    public String productPlan() {
        return this.productPlan;
    }
    public String projectId() {
        return this.projectId;
    }
    public Optional<String> region() {
        return Optional.ofNullable(this.region);
    }
    public String status() {
        return this.status;
    }
    public String updatedAt() {
        return this.updatedAt;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetIotHubResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer connectedDeviceCount;
        private String createdAt;
        private Boolean deviceAutoProvisioning;
        private Integer deviceCount;
        private Boolean disableEvents;
        private Boolean enabled;
        private String endpoint;
        private String eventsTopicPrefix;
        private String hubCa;
        private String hubCaChallenge;
        private @Nullable String hubId;
        private String id;
        private @Nullable String name;
        private String organizationId;
        private String productPlan;
        private String projectId;
        private @Nullable String region;
        private String status;
        private String updatedAt;
        public Builder() {}
        public Builder(GetIotHubResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.connectedDeviceCount = defaults.connectedDeviceCount;
    	      this.createdAt = defaults.createdAt;
    	      this.deviceAutoProvisioning = defaults.deviceAutoProvisioning;
    	      this.deviceCount = defaults.deviceCount;
    	      this.disableEvents = defaults.disableEvents;
    	      this.enabled = defaults.enabled;
    	      this.endpoint = defaults.endpoint;
    	      this.eventsTopicPrefix = defaults.eventsTopicPrefix;
    	      this.hubCa = defaults.hubCa;
    	      this.hubCaChallenge = defaults.hubCaChallenge;
    	      this.hubId = defaults.hubId;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.organizationId = defaults.organizationId;
    	      this.productPlan = defaults.productPlan;
    	      this.projectId = defaults.projectId;
    	      this.region = defaults.region;
    	      this.status = defaults.status;
    	      this.updatedAt = defaults.updatedAt;
        }

        @CustomType.Setter
        public Builder connectedDeviceCount(Integer connectedDeviceCount) {
            this.connectedDeviceCount = Objects.requireNonNull(connectedDeviceCount);
            return this;
        }
        @CustomType.Setter
        public Builder createdAt(String createdAt) {
            this.createdAt = Objects.requireNonNull(createdAt);
            return this;
        }
        @CustomType.Setter
        public Builder deviceAutoProvisioning(Boolean deviceAutoProvisioning) {
            this.deviceAutoProvisioning = Objects.requireNonNull(deviceAutoProvisioning);
            return this;
        }
        @CustomType.Setter
        public Builder deviceCount(Integer deviceCount) {
            this.deviceCount = Objects.requireNonNull(deviceCount);
            return this;
        }
        @CustomType.Setter
        public Builder disableEvents(Boolean disableEvents) {
            this.disableEvents = Objects.requireNonNull(disableEvents);
            return this;
        }
        @CustomType.Setter
        public Builder enabled(Boolean enabled) {
            this.enabled = Objects.requireNonNull(enabled);
            return this;
        }
        @CustomType.Setter
        public Builder endpoint(String endpoint) {
            this.endpoint = Objects.requireNonNull(endpoint);
            return this;
        }
        @CustomType.Setter
        public Builder eventsTopicPrefix(String eventsTopicPrefix) {
            this.eventsTopicPrefix = Objects.requireNonNull(eventsTopicPrefix);
            return this;
        }
        @CustomType.Setter
        public Builder hubCa(String hubCa) {
            this.hubCa = Objects.requireNonNull(hubCa);
            return this;
        }
        @CustomType.Setter
        public Builder hubCaChallenge(String hubCaChallenge) {
            this.hubCaChallenge = Objects.requireNonNull(hubCaChallenge);
            return this;
        }
        @CustomType.Setter
        public Builder hubId(@Nullable String hubId) {
            this.hubId = hubId;
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
        public Builder productPlan(String productPlan) {
            this.productPlan = Objects.requireNonNull(productPlan);
            return this;
        }
        @CustomType.Setter
        public Builder projectId(String projectId) {
            this.projectId = Objects.requireNonNull(projectId);
            return this;
        }
        @CustomType.Setter
        public Builder region(@Nullable String region) {
            this.region = region;
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        @CustomType.Setter
        public Builder updatedAt(String updatedAt) {
            this.updatedAt = Objects.requireNonNull(updatedAt);
            return this;
        }
        public GetIotHubResult build() {
            final var o = new GetIotHubResult();
            o.connectedDeviceCount = connectedDeviceCount;
            o.createdAt = createdAt;
            o.deviceAutoProvisioning = deviceAutoProvisioning;
            o.deviceCount = deviceCount;
            o.disableEvents = disableEvents;
            o.enabled = enabled;
            o.endpoint = endpoint;
            o.eventsTopicPrefix = eventsTopicPrefix;
            o.hubCa = hubCa;
            o.hubCaChallenge = hubCaChallenge;
            o.hubId = hubId;
            o.id = id;
            o.name = name;
            o.organizationId = organizationId;
            o.productPlan = productPlan;
            o.projectId = projectId;
            o.region = region;
            o.status = status;
            o.updatedAt = updatedAt;
            return o;
        }
    }
}