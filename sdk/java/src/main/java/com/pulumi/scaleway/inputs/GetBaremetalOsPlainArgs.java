// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetBaremetalOsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetBaremetalOsPlainArgs Empty = new GetBaremetalOsPlainArgs();

    /**
     * The os name. Only one of `name` and `os_id` should be specified.
     * 
     */
    @Import(name="name")
    private @Nullable String name;

    /**
     * @return The os name. Only one of `name` and `os_id` should be specified.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The operating system id. Only one of `name` and `os_id` should be specified.
     * 
     */
    @Import(name="osId")
    private @Nullable String osId;

    /**
     * @return The operating system id. Only one of `name` and `os_id` should be specified.
     * 
     */
    public Optional<String> osId() {
        return Optional.ofNullable(this.osId);
    }

    /**
     * The os version.
     * 
     */
    @Import(name="version")
    private @Nullable String version;

    /**
     * @return The os version.
     * 
     */
    public Optional<String> version() {
        return Optional.ofNullable(this.version);
    }

    /**
     * `zone`) The zone in which the os exists.
     * 
     */
    @Import(name="zone")
    private @Nullable String zone;

    /**
     * @return `zone`) The zone in which the os exists.
     * 
     */
    public Optional<String> zone() {
        return Optional.ofNullable(this.zone);
    }

    private GetBaremetalOsPlainArgs() {}

    private GetBaremetalOsPlainArgs(GetBaremetalOsPlainArgs $) {
        this.name = $.name;
        this.osId = $.osId;
        this.version = $.version;
        this.zone = $.zone;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetBaremetalOsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetBaremetalOsPlainArgs $;

        public Builder() {
            $ = new GetBaremetalOsPlainArgs();
        }

        public Builder(GetBaremetalOsPlainArgs defaults) {
            $ = new GetBaremetalOsPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name The os name. Only one of `name` and `os_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable String name) {
            $.name = name;
            return this;
        }

        /**
         * @param osId The operating system id. Only one of `name` and `os_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder osId(@Nullable String osId) {
            $.osId = osId;
            return this;
        }

        /**
         * @param version The os version.
         * 
         * @return builder
         * 
         */
        public Builder version(@Nullable String version) {
            $.version = version;
            return this;
        }

        /**
         * @param zone `zone`) The zone in which the os exists.
         * 
         * @return builder
         * 
         */
        public Builder zone(@Nullable String zone) {
            $.zone = zone;
            return this;
        }

        public GetBaremetalOsPlainArgs build() {
            return $;
        }
    }

}