// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetVpcGatewayNetworkArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetVpcGatewayNetworkArgs Empty = new GetVpcGatewayNetworkArgs();

    /**
     * ID of the public gateway DHCP config
     * 
     */
    @Import(name="dhcpId")
    private @Nullable Output<String> dhcpId;

    /**
     * @return ID of the public gateway DHCP config
     * 
     */
    public Optional<Output<String>> dhcpId() {
        return Optional.ofNullable(this.dhcpId);
    }

    /**
     * If masquerade is enabled on requested network
     * 
     */
    @Import(name="enableMasquerade")
    private @Nullable Output<Boolean> enableMasquerade;

    /**
     * @return If masquerade is enabled on requested network
     * 
     */
    public Optional<Output<Boolean>> enableMasquerade() {
        return Optional.ofNullable(this.enableMasquerade);
    }

    /**
     * ID of the public gateway the gateway network is linked to
     * 
     */
    @Import(name="gatewayId")
    private @Nullable Output<String> gatewayId;

    /**
     * @return ID of the public gateway the gateway network is linked to
     * 
     */
    public Optional<Output<String>> gatewayId() {
        return Optional.ofNullable(this.gatewayId);
    }

    /**
     * ID of the gateway network.
     * 
     * &gt; Only one of `gateway_network_id` or filters should be specified. You can use all the filters you want.
     * 
     */
    @Import(name="gatewayNetworkId")
    private @Nullable Output<String> gatewayNetworkId;

    /**
     * @return ID of the gateway network.
     * 
     * &gt; Only one of `gateway_network_id` or filters should be specified. You can use all the filters you want.
     * 
     */
    public Optional<Output<String>> gatewayNetworkId() {
        return Optional.ofNullable(this.gatewayNetworkId);
    }

    /**
     * ID of the private network the gateway network is linked to
     * 
     */
    @Import(name="privateNetworkId")
    private @Nullable Output<String> privateNetworkId;

    /**
     * @return ID of the private network the gateway network is linked to
     * 
     */
    public Optional<Output<String>> privateNetworkId() {
        return Optional.ofNullable(this.privateNetworkId);
    }

    private GetVpcGatewayNetworkArgs() {}

    private GetVpcGatewayNetworkArgs(GetVpcGatewayNetworkArgs $) {
        this.dhcpId = $.dhcpId;
        this.enableMasquerade = $.enableMasquerade;
        this.gatewayId = $.gatewayId;
        this.gatewayNetworkId = $.gatewayNetworkId;
        this.privateNetworkId = $.privateNetworkId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetVpcGatewayNetworkArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetVpcGatewayNetworkArgs $;

        public Builder() {
            $ = new GetVpcGatewayNetworkArgs();
        }

        public Builder(GetVpcGatewayNetworkArgs defaults) {
            $ = new GetVpcGatewayNetworkArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param dhcpId ID of the public gateway DHCP config
         * 
         * @return builder
         * 
         */
        public Builder dhcpId(@Nullable Output<String> dhcpId) {
            $.dhcpId = dhcpId;
            return this;
        }

        /**
         * @param dhcpId ID of the public gateway DHCP config
         * 
         * @return builder
         * 
         */
        public Builder dhcpId(String dhcpId) {
            return dhcpId(Output.of(dhcpId));
        }

        /**
         * @param enableMasquerade If masquerade is enabled on requested network
         * 
         * @return builder
         * 
         */
        public Builder enableMasquerade(@Nullable Output<Boolean> enableMasquerade) {
            $.enableMasquerade = enableMasquerade;
            return this;
        }

        /**
         * @param enableMasquerade If masquerade is enabled on requested network
         * 
         * @return builder
         * 
         */
        public Builder enableMasquerade(Boolean enableMasquerade) {
            return enableMasquerade(Output.of(enableMasquerade));
        }

        /**
         * @param gatewayId ID of the public gateway the gateway network is linked to
         * 
         * @return builder
         * 
         */
        public Builder gatewayId(@Nullable Output<String> gatewayId) {
            $.gatewayId = gatewayId;
            return this;
        }

        /**
         * @param gatewayId ID of the public gateway the gateway network is linked to
         * 
         * @return builder
         * 
         */
        public Builder gatewayId(String gatewayId) {
            return gatewayId(Output.of(gatewayId));
        }

        /**
         * @param gatewayNetworkId ID of the gateway network.
         * 
         * &gt; Only one of `gateway_network_id` or filters should be specified. You can use all the filters you want.
         * 
         * @return builder
         * 
         */
        public Builder gatewayNetworkId(@Nullable Output<String> gatewayNetworkId) {
            $.gatewayNetworkId = gatewayNetworkId;
            return this;
        }

        /**
         * @param gatewayNetworkId ID of the gateway network.
         * 
         * &gt; Only one of `gateway_network_id` or filters should be specified. You can use all the filters you want.
         * 
         * @return builder
         * 
         */
        public Builder gatewayNetworkId(String gatewayNetworkId) {
            return gatewayNetworkId(Output.of(gatewayNetworkId));
        }

        /**
         * @param privateNetworkId ID of the private network the gateway network is linked to
         * 
         * @return builder
         * 
         */
        public Builder privateNetworkId(@Nullable Output<String> privateNetworkId) {
            $.privateNetworkId = privateNetworkId;
            return this;
        }

        /**
         * @param privateNetworkId ID of the private network the gateway network is linked to
         * 
         * @return builder
         * 
         */
        public Builder privateNetworkId(String privateNetworkId) {
            return privateNetworkId(Output.of(privateNetworkId));
        }

        public GetVpcGatewayNetworkArgs build() {
            return $;
        }
    }

}
