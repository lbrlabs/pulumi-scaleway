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


public final class LoadbalancerPrivateNetworkArgs extends com.pulumi.resources.ResourceArgs {

    public static final LoadbalancerPrivateNetworkArgs Empty = new LoadbalancerPrivateNetworkArgs();

    /**
     * (Optional) Set to true if you want to let DHCP assign IP addresses. See below.
     * 
     */
    @Import(name="dhcpConfig")
    private @Nullable Output<Boolean> dhcpConfig;

    /**
     * @return (Optional) Set to true if you want to let DHCP assign IP addresses. See below.
     * 
     */
    public Optional<Output<Boolean>> dhcpConfig() {
        return Optional.ofNullable(this.dhcpConfig);
    }

    /**
     * (Required) The ID of the Private Network to associate.
     * 
     */
    @Import(name="privateNetworkId", required=true)
    private Output<String> privateNetworkId;

    /**
     * @return (Required) The ID of the Private Network to associate.
     * 
     */
    public Output<String> privateNetworkId() {
        return this.privateNetworkId;
    }

    /**
     * (Optional) Define a local ip address of your choice for the load balancer instance. See below.
     * 
     */
    @Import(name="staticConfig")
    private @Nullable Output<String> staticConfig;

    /**
     * @return (Optional) Define a local ip address of your choice for the load balancer instance. See below.
     * 
     */
    public Optional<Output<String>> staticConfig() {
        return Optional.ofNullable(this.staticConfig);
    }

    @Import(name="status")
    private @Nullable Output<String> status;

    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * `zone`) The zone of the load-balancer.
     * 
     */
    @Import(name="zone")
    private @Nullable Output<String> zone;

    /**
     * @return `zone`) The zone of the load-balancer.
     * 
     */
    public Optional<Output<String>> zone() {
        return Optional.ofNullable(this.zone);
    }

    private LoadbalancerPrivateNetworkArgs() {}

    private LoadbalancerPrivateNetworkArgs(LoadbalancerPrivateNetworkArgs $) {
        this.dhcpConfig = $.dhcpConfig;
        this.privateNetworkId = $.privateNetworkId;
        this.staticConfig = $.staticConfig;
        this.status = $.status;
        this.zone = $.zone;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LoadbalancerPrivateNetworkArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LoadbalancerPrivateNetworkArgs $;

        public Builder() {
            $ = new LoadbalancerPrivateNetworkArgs();
        }

        public Builder(LoadbalancerPrivateNetworkArgs defaults) {
            $ = new LoadbalancerPrivateNetworkArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param dhcpConfig (Optional) Set to true if you want to let DHCP assign IP addresses. See below.
         * 
         * @return builder
         * 
         */
        public Builder dhcpConfig(@Nullable Output<Boolean> dhcpConfig) {
            $.dhcpConfig = dhcpConfig;
            return this;
        }

        /**
         * @param dhcpConfig (Optional) Set to true if you want to let DHCP assign IP addresses. See below.
         * 
         * @return builder
         * 
         */
        public Builder dhcpConfig(Boolean dhcpConfig) {
            return dhcpConfig(Output.of(dhcpConfig));
        }

        /**
         * @param privateNetworkId (Required) The ID of the Private Network to associate.
         * 
         * @return builder
         * 
         */
        public Builder privateNetworkId(Output<String> privateNetworkId) {
            $.privateNetworkId = privateNetworkId;
            return this;
        }

        /**
         * @param privateNetworkId (Required) The ID of the Private Network to associate.
         * 
         * @return builder
         * 
         */
        public Builder privateNetworkId(String privateNetworkId) {
            return privateNetworkId(Output.of(privateNetworkId));
        }

        /**
         * @param staticConfig (Optional) Define a local ip address of your choice for the load balancer instance. See below.
         * 
         * @return builder
         * 
         */
        public Builder staticConfig(@Nullable Output<String> staticConfig) {
            $.staticConfig = staticConfig;
            return this;
        }

        /**
         * @param staticConfig (Optional) Define a local ip address of your choice for the load balancer instance. See below.
         * 
         * @return builder
         * 
         */
        public Builder staticConfig(String staticConfig) {
            return staticConfig(Output.of(staticConfig));
        }

        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param zone `zone`) The zone of the load-balancer.
         * 
         * @return builder
         * 
         */
        public Builder zone(@Nullable Output<String> zone) {
            $.zone = zone;
            return this;
        }

        /**
         * @param zone `zone`) The zone of the load-balancer.
         * 
         * @return builder
         * 
         */
        public Builder zone(String zone) {
            return zone(Output.of(zone));
        }

        public LoadbalancerPrivateNetworkArgs build() {
            $.privateNetworkId = Objects.requireNonNull($.privateNetworkId, "expected parameter 'privateNetworkId' to be non-null");
            return $;
        }
    }

}
