// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetInstancePrivateNicArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetInstancePrivateNicArgs Empty = new GetInstancePrivateNicArgs();

    /**
     * The ID of the private network
     * Only one of `private_nic_id` and `private_network_id` should be specified.
     * 
     */
    @Import(name="privateNetworkId")
    private @Nullable Output<String> privateNetworkId;

    /**
     * @return The ID of the private network
     * Only one of `private_nic_id` and `private_network_id` should be specified.
     * 
     */
    public Optional<Output<String>> privateNetworkId() {
        return Optional.ofNullable(this.privateNetworkId);
    }

    /**
     * The ID of the instance server private nic
     * Only one of `private_nic_id` and `private_network_id` should be specified.
     * 
     */
    @Import(name="privateNicId")
    private @Nullable Output<String> privateNicId;

    /**
     * @return The ID of the instance server private nic
     * Only one of `private_nic_id` and `private_network_id` should be specified.
     * 
     */
    public Optional<Output<String>> privateNicId() {
        return Optional.ofNullable(this.privateNicId);
    }

    /**
     * The server&#39;s id
     * 
     */
    @Import(name="serverId", required=true)
    private Output<String> serverId;

    /**
     * @return The server&#39;s id
     * 
     */
    public Output<String> serverId() {
        return this.serverId;
    }

    /**
     * The tags associated with the private NIC.
     * As datasource only returns one private NIC, the search with given tags must return only one result
     * 
     */
    @Import(name="tags")
    private @Nullable Output<List<String>> tags;

    /**
     * @return The tags associated with the private NIC.
     * As datasource only returns one private NIC, the search with given tags must return only one result
     * 
     */
    public Optional<Output<List<String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * `zone`) The zone in which the private nic exists.
     * 
     */
    @Import(name="zone")
    private @Nullable Output<String> zone;

    /**
     * @return `zone`) The zone in which the private nic exists.
     * 
     */
    public Optional<Output<String>> zone() {
        return Optional.ofNullable(this.zone);
    }

    private GetInstancePrivateNicArgs() {}

    private GetInstancePrivateNicArgs(GetInstancePrivateNicArgs $) {
        this.privateNetworkId = $.privateNetworkId;
        this.privateNicId = $.privateNicId;
        this.serverId = $.serverId;
        this.tags = $.tags;
        this.zone = $.zone;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetInstancePrivateNicArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetInstancePrivateNicArgs $;

        public Builder() {
            $ = new GetInstancePrivateNicArgs();
        }

        public Builder(GetInstancePrivateNicArgs defaults) {
            $ = new GetInstancePrivateNicArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param privateNetworkId The ID of the private network
         * Only one of `private_nic_id` and `private_network_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder privateNetworkId(@Nullable Output<String> privateNetworkId) {
            $.privateNetworkId = privateNetworkId;
            return this;
        }

        /**
         * @param privateNetworkId The ID of the private network
         * Only one of `private_nic_id` and `private_network_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder privateNetworkId(String privateNetworkId) {
            return privateNetworkId(Output.of(privateNetworkId));
        }

        /**
         * @param privateNicId The ID of the instance server private nic
         * Only one of `private_nic_id` and `private_network_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder privateNicId(@Nullable Output<String> privateNicId) {
            $.privateNicId = privateNicId;
            return this;
        }

        /**
         * @param privateNicId The ID of the instance server private nic
         * Only one of `private_nic_id` and `private_network_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder privateNicId(String privateNicId) {
            return privateNicId(Output.of(privateNicId));
        }

        /**
         * @param serverId The server&#39;s id
         * 
         * @return builder
         * 
         */
        public Builder serverId(Output<String> serverId) {
            $.serverId = serverId;
            return this;
        }

        /**
         * @param serverId The server&#39;s id
         * 
         * @return builder
         * 
         */
        public Builder serverId(String serverId) {
            return serverId(Output.of(serverId));
        }

        /**
         * @param tags The tags associated with the private NIC.
         * As datasource only returns one private NIC, the search with given tags must return only one result
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<List<String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags The tags associated with the private NIC.
         * As datasource only returns one private NIC, the search with given tags must return only one result
         * 
         * @return builder
         * 
         */
        public Builder tags(List<String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tags The tags associated with the private NIC.
         * As datasource only returns one private NIC, the search with given tags must return only one result
         * 
         * @return builder
         * 
         */
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }

        /**
         * @param zone `zone`) The zone in which the private nic exists.
         * 
         * @return builder
         * 
         */
        public Builder zone(@Nullable Output<String> zone) {
            $.zone = zone;
            return this;
        }

        /**
         * @param zone `zone`) The zone in which the private nic exists.
         * 
         * @return builder
         * 
         */
        public Builder zone(String zone) {
            return zone(Output.of(zone));
        }

        public GetInstancePrivateNicArgs build() {
            $.serverId = Objects.requireNonNull($.serverId, "expected parameter 'serverId' to be non-null");
            return $;
        }
    }

}