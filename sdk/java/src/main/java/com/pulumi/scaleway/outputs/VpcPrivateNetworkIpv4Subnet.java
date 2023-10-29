// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class VpcPrivateNetworkIpv4Subnet {
    /**
     * @return The network address of the subnet in dotted decimal notation, e.g., &#39;192.168.0.0&#39; for a &#39;192.168.0.0/24&#39; subnet.
     * 
     */
    private @Nullable String address;
    /**
     * @return The date and time of the creation of the subnet.
     * 
     */
    private @Nullable String createdAt;
    /**
     * @return The subnet ID.
     * 
     */
    private @Nullable String id;
    /**
     * @return The length of the network prefix, e.g., 24 for a 255.255.255.0 mask.
     * 
     */
    private @Nullable Integer prefixLength;
    /**
     * @return The subnet CIDR.
     * 
     */
    private @Nullable String subnet;
    /**
     * @return The subnet mask expressed in dotted decimal notation, e.g., &#39;255.255.255.0&#39; for a /24 subnet
     * 
     */
    private @Nullable String subnetMask;
    /**
     * @return The date and time of the last update of the subnet.
     * 
     */
    private @Nullable String updatedAt;

    private VpcPrivateNetworkIpv4Subnet() {}
    /**
     * @return The network address of the subnet in dotted decimal notation, e.g., &#39;192.168.0.0&#39; for a &#39;192.168.0.0/24&#39; subnet.
     * 
     */
    public Optional<String> address() {
        return Optional.ofNullable(this.address);
    }
    /**
     * @return The date and time of the creation of the subnet.
     * 
     */
    public Optional<String> createdAt() {
        return Optional.ofNullable(this.createdAt);
    }
    /**
     * @return The subnet ID.
     * 
     */
    public Optional<String> id() {
        return Optional.ofNullable(this.id);
    }
    /**
     * @return The length of the network prefix, e.g., 24 for a 255.255.255.0 mask.
     * 
     */
    public Optional<Integer> prefixLength() {
        return Optional.ofNullable(this.prefixLength);
    }
    /**
     * @return The subnet CIDR.
     * 
     */
    public Optional<String> subnet() {
        return Optional.ofNullable(this.subnet);
    }
    /**
     * @return The subnet mask expressed in dotted decimal notation, e.g., &#39;255.255.255.0&#39; for a /24 subnet
     * 
     */
    public Optional<String> subnetMask() {
        return Optional.ofNullable(this.subnetMask);
    }
    /**
     * @return The date and time of the last update of the subnet.
     * 
     */
    public Optional<String> updatedAt() {
        return Optional.ofNullable(this.updatedAt);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(VpcPrivateNetworkIpv4Subnet defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String address;
        private @Nullable String createdAt;
        private @Nullable String id;
        private @Nullable Integer prefixLength;
        private @Nullable String subnet;
        private @Nullable String subnetMask;
        private @Nullable String updatedAt;
        public Builder() {}
        public Builder(VpcPrivateNetworkIpv4Subnet defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.address = defaults.address;
    	      this.createdAt = defaults.createdAt;
    	      this.id = defaults.id;
    	      this.prefixLength = defaults.prefixLength;
    	      this.subnet = defaults.subnet;
    	      this.subnetMask = defaults.subnetMask;
    	      this.updatedAt = defaults.updatedAt;
        }

        @CustomType.Setter
        public Builder address(@Nullable String address) {
            this.address = address;
            return this;
        }
        @CustomType.Setter
        public Builder createdAt(@Nullable String createdAt) {
            this.createdAt = createdAt;
            return this;
        }
        @CustomType.Setter
        public Builder id(@Nullable String id) {
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder prefixLength(@Nullable Integer prefixLength) {
            this.prefixLength = prefixLength;
            return this;
        }
        @CustomType.Setter
        public Builder subnet(@Nullable String subnet) {
            this.subnet = subnet;
            return this;
        }
        @CustomType.Setter
        public Builder subnetMask(@Nullable String subnetMask) {
            this.subnetMask = subnetMask;
            return this;
        }
        @CustomType.Setter
        public Builder updatedAt(@Nullable String updatedAt) {
            this.updatedAt = updatedAt;
            return this;
        }
        public VpcPrivateNetworkIpv4Subnet build() {
            final var o = new VpcPrivateNetworkIpv4Subnet();
            o.address = address;
            o.createdAt = createdAt;
            o.id = id;
            o.prefixLength = prefixLength;
            o.subnet = subnet;
            o.subnetMask = subnetMask;
            o.updatedAt = updatedAt;
            return o;
        }
    }
}
