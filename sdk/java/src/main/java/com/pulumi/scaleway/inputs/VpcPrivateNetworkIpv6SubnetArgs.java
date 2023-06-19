// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VpcPrivateNetworkIpv6SubnetArgs extends com.pulumi.resources.ResourceArgs {

    public static final VpcPrivateNetworkIpv6SubnetArgs Empty = new VpcPrivateNetworkIpv6SubnetArgs();

    @Import(name="createdAt")
    private @Nullable Output<String> createdAt;

    public Optional<Output<String>> createdAt() {
        return Optional.ofNullable(this.createdAt);
    }

    /**
     * The ID of the private network.
     * 
     */
    @Import(name="id")
    private @Nullable Output<String> id;

    /**
     * @return The ID of the private network.
     * 
     */
    public Optional<Output<String>> id() {
        return Optional.ofNullable(this.id);
    }

    @Import(name="subnet")
    private @Nullable Output<String> subnet;

    public Optional<Output<String>> subnet() {
        return Optional.ofNullable(this.subnet);
    }

    @Import(name="updatedAt")
    private @Nullable Output<String> updatedAt;

    public Optional<Output<String>> updatedAt() {
        return Optional.ofNullable(this.updatedAt);
    }

    private VpcPrivateNetworkIpv6SubnetArgs() {}

    private VpcPrivateNetworkIpv6SubnetArgs(VpcPrivateNetworkIpv6SubnetArgs $) {
        this.createdAt = $.createdAt;
        this.id = $.id;
        this.subnet = $.subnet;
        this.updatedAt = $.updatedAt;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VpcPrivateNetworkIpv6SubnetArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VpcPrivateNetworkIpv6SubnetArgs $;

        public Builder() {
            $ = new VpcPrivateNetworkIpv6SubnetArgs();
        }

        public Builder(VpcPrivateNetworkIpv6SubnetArgs defaults) {
            $ = new VpcPrivateNetworkIpv6SubnetArgs(Objects.requireNonNull(defaults));
        }

        public Builder createdAt(@Nullable Output<String> createdAt) {
            $.createdAt = createdAt;
            return this;
        }

        public Builder createdAt(String createdAt) {
            return createdAt(Output.of(createdAt));
        }

        /**
         * @param id The ID of the private network.
         * 
         * @return builder
         * 
         */
        public Builder id(@Nullable Output<String> id) {
            $.id = id;
            return this;
        }

        /**
         * @param id The ID of the private network.
         * 
         * @return builder
         * 
         */
        public Builder id(String id) {
            return id(Output.of(id));
        }

        public Builder subnet(@Nullable Output<String> subnet) {
            $.subnet = subnet;
            return this;
        }

        public Builder subnet(String subnet) {
            return subnet(Output.of(subnet));
        }

        public Builder updatedAt(@Nullable Output<String> updatedAt) {
            $.updatedAt = updatedAt;
            return this;
        }

        public Builder updatedAt(String updatedAt) {
            return updatedAt(Output.of(updatedAt));
        }

        public VpcPrivateNetworkIpv6SubnetArgs build() {
            return $;
        }
    }

}