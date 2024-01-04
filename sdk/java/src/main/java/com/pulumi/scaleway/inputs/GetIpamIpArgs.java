// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.scaleway.inputs.GetIpamIpResourceArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetIpamIpArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetIpamIpArgs Empty = new GetIpamIpArgs();

    /**
     * The Mac Address linked to the IP.
     * 
     */
    @Import(name="macAddress")
    private @Nullable Output<String> macAddress;

    /**
     * @return The Mac Address linked to the IP.
     * 
     */
    public Optional<Output<String>> macAddress() {
        return Optional.ofNullable(this.macAddress);
    }

    /**
     * The ID of the private network the IP belong to.
     * 
     */
    @Import(name="privateNetworkId")
    private @Nullable Output<String> privateNetworkId;

    /**
     * @return The ID of the private network the IP belong to.
     * 
     */
    public Optional<Output<String>> privateNetworkId() {
        return Optional.ofNullable(this.privateNetworkId);
    }

    /**
     * `project_id`) The ID of the project the IP is associated with.
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return `project_id`) The ID of the project the IP is associated with.
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * `region`) The region in which the IP exists.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return `region`) The region in which the IP exists.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * Filter by resource ID, type or name. If specified, `type` is required, and at least one of `id` or `name` must be set.
     * 
     */
    @Import(name="resource")
    private @Nullable Output<GetIpamIpResourceArgs> resource;

    /**
     * @return Filter by resource ID, type or name. If specified, `type` is required, and at least one of `id` or `name` must be set.
     * 
     */
    public Optional<Output<GetIpamIpResourceArgs>> resource() {
        return Optional.ofNullable(this.resource);
    }

    /**
     * The tags associated with the IP.
     * As datasource only returns one IP, the search with given tags must return only one result.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<List<String>> tags;

    /**
     * @return The tags associated with the IP.
     * As datasource only returns one IP, the search with given tags must return only one result.
     * 
     */
    public Optional<Output<List<String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * The type of the resource to get the IP from. [Documentation](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/ipam/v1#pkg-constants) with type list.
     * 
     */
    @Import(name="type", required=true)
    private Output<String> type;

    /**
     * @return The type of the resource to get the IP from. [Documentation](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/ipam/v1#pkg-constants) with type list.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    /**
     * Only IPs that are zonal, and in this zone, will be returned.
     * 
     */
    @Import(name="zonal")
    private @Nullable Output<String> zonal;

    /**
     * @return Only IPs that are zonal, and in this zone, will be returned.
     * 
     */
    public Optional<Output<String>> zonal() {
        return Optional.ofNullable(this.zonal);
    }

    private GetIpamIpArgs() {}

    private GetIpamIpArgs(GetIpamIpArgs $) {
        this.macAddress = $.macAddress;
        this.privateNetworkId = $.privateNetworkId;
        this.projectId = $.projectId;
        this.region = $.region;
        this.resource = $.resource;
        this.tags = $.tags;
        this.type = $.type;
        this.zonal = $.zonal;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetIpamIpArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetIpamIpArgs $;

        public Builder() {
            $ = new GetIpamIpArgs();
        }

        public Builder(GetIpamIpArgs defaults) {
            $ = new GetIpamIpArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param macAddress The Mac Address linked to the IP.
         * 
         * @return builder
         * 
         */
        public Builder macAddress(@Nullable Output<String> macAddress) {
            $.macAddress = macAddress;
            return this;
        }

        /**
         * @param macAddress The Mac Address linked to the IP.
         * 
         * @return builder
         * 
         */
        public Builder macAddress(String macAddress) {
            return macAddress(Output.of(macAddress));
        }

        /**
         * @param privateNetworkId The ID of the private network the IP belong to.
         * 
         * @return builder
         * 
         */
        public Builder privateNetworkId(@Nullable Output<String> privateNetworkId) {
            $.privateNetworkId = privateNetworkId;
            return this;
        }

        /**
         * @param privateNetworkId The ID of the private network the IP belong to.
         * 
         * @return builder
         * 
         */
        public Builder privateNetworkId(String privateNetworkId) {
            return privateNetworkId(Output.of(privateNetworkId));
        }

        /**
         * @param projectId `project_id`) The ID of the project the IP is associated with.
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId `project_id`) The ID of the project the IP is associated with.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param region `region`) The region in which the IP exists.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region `region`) The region in which the IP exists.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param resource Filter by resource ID, type or name. If specified, `type` is required, and at least one of `id` or `name` must be set.
         * 
         * @return builder
         * 
         */
        public Builder resource(@Nullable Output<GetIpamIpResourceArgs> resource) {
            $.resource = resource;
            return this;
        }

        /**
         * @param resource Filter by resource ID, type or name. If specified, `type` is required, and at least one of `id` or `name` must be set.
         * 
         * @return builder
         * 
         */
        public Builder resource(GetIpamIpResourceArgs resource) {
            return resource(Output.of(resource));
        }

        /**
         * @param tags The tags associated with the IP.
         * As datasource only returns one IP, the search with given tags must return only one result.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<List<String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags The tags associated with the IP.
         * As datasource only returns one IP, the search with given tags must return only one result.
         * 
         * @return builder
         * 
         */
        public Builder tags(List<String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tags The tags associated with the IP.
         * As datasource only returns one IP, the search with given tags must return only one result.
         * 
         * @return builder
         * 
         */
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }

        /**
         * @param type The type of the resource to get the IP from. [Documentation](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/ipam/v1#pkg-constants) with type list.
         * 
         * @return builder
         * 
         */
        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type The type of the resource to get the IP from. [Documentation](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/ipam/v1#pkg-constants) with type list.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        /**
         * @param zonal Only IPs that are zonal, and in this zone, will be returned.
         * 
         * @return builder
         * 
         */
        public Builder zonal(@Nullable Output<String> zonal) {
            $.zonal = zonal;
            return this;
        }

        /**
         * @param zonal Only IPs that are zonal, and in this zone, will be returned.
         * 
         * @return builder
         * 
         */
        public Builder zonal(String zonal) {
            return zonal(Output.of(zonal));
        }

        public GetIpamIpArgs build() {
            $.type = Objects.requireNonNull($.type, "expected parameter 'type' to be non-null");
            return $;
        }
    }

}
