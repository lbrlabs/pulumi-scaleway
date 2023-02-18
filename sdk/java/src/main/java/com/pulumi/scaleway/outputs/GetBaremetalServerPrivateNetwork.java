// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetBaremetalServerPrivateNetwork {
    private String createdAt;
    private String id;
    private String status;
    private String updatedAt;
    private Integer vlan;

    private GetBaremetalServerPrivateNetwork() {}
    public String createdAt() {
        return this.createdAt;
    }
    public String id() {
        return this.id;
    }
    public String status() {
        return this.status;
    }
    public String updatedAt() {
        return this.updatedAt;
    }
    public Integer vlan() {
        return this.vlan;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetBaremetalServerPrivateNetwork defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String createdAt;
        private String id;
        private String status;
        private String updatedAt;
        private Integer vlan;
        public Builder() {}
        public Builder(GetBaremetalServerPrivateNetwork defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.createdAt = defaults.createdAt;
    	      this.id = defaults.id;
    	      this.status = defaults.status;
    	      this.updatedAt = defaults.updatedAt;
    	      this.vlan = defaults.vlan;
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
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        @CustomType.Setter
        public Builder updatedAt(String updatedAt) {
            this.updatedAt = Objects.requireNonNull(updatedAt);
            return this;
        }
        @CustomType.Setter
        public Builder vlan(Integer vlan) {
            this.vlan = Objects.requireNonNull(vlan);
            return this;
        }
        public GetBaremetalServerPrivateNetwork build() {
            final var o = new GetBaremetalServerPrivateNetwork();
            o.createdAt = createdAt;
            o.id = id;
            o.status = status;
            o.updatedAt = updatedAt;
            o.vlan = vlan;
            return o;
        }
    }
}
