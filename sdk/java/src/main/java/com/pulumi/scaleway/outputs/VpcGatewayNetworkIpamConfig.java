// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class VpcGatewayNetworkIpamConfig {
    /**
     * @return Use this IPAM-booked IP ID as the Gateway&#39;s IP in this Private Network.
     * 
     */
    private @Nullable String ipamIpId;
    /**
     * @return Defines whether the default route is enabled on that Gateway Network.
     * 
     */
    private @Nullable Boolean pushDefaultRoute;

    private VpcGatewayNetworkIpamConfig() {}
    /**
     * @return Use this IPAM-booked IP ID as the Gateway&#39;s IP in this Private Network.
     * 
     */
    public Optional<String> ipamIpId() {
        return Optional.ofNullable(this.ipamIpId);
    }
    /**
     * @return Defines whether the default route is enabled on that Gateway Network.
     * 
     */
    public Optional<Boolean> pushDefaultRoute() {
        return Optional.ofNullable(this.pushDefaultRoute);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(VpcGatewayNetworkIpamConfig defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String ipamIpId;
        private @Nullable Boolean pushDefaultRoute;
        public Builder() {}
        public Builder(VpcGatewayNetworkIpamConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.ipamIpId = defaults.ipamIpId;
    	      this.pushDefaultRoute = defaults.pushDefaultRoute;
        }

        @CustomType.Setter
        public Builder ipamIpId(@Nullable String ipamIpId) {
            this.ipamIpId = ipamIpId;
            return this;
        }
        @CustomType.Setter
        public Builder pushDefaultRoute(@Nullable Boolean pushDefaultRoute) {
            this.pushDefaultRoute = pushDefaultRoute;
            return this;
        }
        public VpcGatewayNetworkIpamConfig build() {
            final var o = new VpcGatewayNetworkIpamConfig();
            o.ipamIpId = ipamIpId;
            o.pushDefaultRoute = pushDefaultRoute;
            return o;
        }
    }
}
