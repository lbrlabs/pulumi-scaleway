// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class VpcGatewayNetworkIpamConfig {
    /**
     * @return Defines whether the default route is enabled on that Gateway Network. Only one of `dhcp_id`, `static_address` and `ipam_config` should be specified.
     * 
     */
    private @Nullable Boolean pushDefaultRoute;

    private VpcGatewayNetworkIpamConfig() {}
    /**
     * @return Defines whether the default route is enabled on that Gateway Network. Only one of `dhcp_id`, `static_address` and `ipam_config` should be specified.
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
        private @Nullable Boolean pushDefaultRoute;
        public Builder() {}
        public Builder(VpcGatewayNetworkIpamConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.pushDefaultRoute = defaults.pushDefaultRoute;
        }

        @CustomType.Setter
        public Builder pushDefaultRoute(@Nullable Boolean pushDefaultRoute) {
            this.pushDefaultRoute = pushDefaultRoute;
            return this;
        }
        public VpcGatewayNetworkIpamConfig build() {
            final var o = new VpcGatewayNetworkIpamConfig();
            o.pushDefaultRoute = pushDefaultRoute;
            return o;
        }
    }
}
