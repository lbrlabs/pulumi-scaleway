// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetVpcPrivateNetworkIpv6Subnet {
    private String createdAt;
    /**
     * @return The ID of the private network.
     * 
     */
    private String id;
    private String subnet;
    private String updatedAt;

    private GetVpcPrivateNetworkIpv6Subnet() {}
    public String createdAt() {
        return this.createdAt;
    }
    /**
     * @return The ID of the private network.
     * 
     */
    public String id() {
        return this.id;
    }
    public String subnet() {
        return this.subnet;
    }
    public String updatedAt() {
        return this.updatedAt;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetVpcPrivateNetworkIpv6Subnet defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String createdAt;
        private String id;
        private String subnet;
        private String updatedAt;
        public Builder() {}
        public Builder(GetVpcPrivateNetworkIpv6Subnet defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.createdAt = defaults.createdAt;
    	      this.id = defaults.id;
    	      this.subnet = defaults.subnet;
    	      this.updatedAt = defaults.updatedAt;
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
        public Builder subnet(String subnet) {
            this.subnet = Objects.requireNonNull(subnet);
            return this;
        }
        @CustomType.Setter
        public Builder updatedAt(String updatedAt) {
            this.updatedAt = Objects.requireNonNull(updatedAt);
            return this;
        }
        public GetVpcPrivateNetworkIpv6Subnet build() {
            final var o = new GetVpcPrivateNetworkIpv6Subnet();
            o.createdAt = createdAt;
            o.id = id;
            o.subnet = subnet;
            o.updatedAt = updatedAt;
            return o;
        }
    }
}