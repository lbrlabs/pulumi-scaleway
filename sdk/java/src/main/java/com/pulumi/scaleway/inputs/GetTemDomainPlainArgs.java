// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetTemDomainPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetTemDomainPlainArgs Empty = new GetTemDomainPlainArgs();

    @Import(name="domainId")
    private @Nullable String domainId;

    public Optional<String> domainId() {
        return Optional.ofNullable(this.domainId);
    }

    /**
     * The domain name.
     * Only one of `name` and `id` should be specified.
     * 
     */
    @Import(name="name")
    private @Nullable String name;

    /**
     * @return The domain name.
     * Only one of `name` and `id` should be specified.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * `region`) The region in which the domain exists.
     * 
     */
    @Import(name="region")
    private @Nullable String region;

    /**
     * @return `region`) The region in which the domain exists.
     * 
     */
    public Optional<String> region() {
        return Optional.ofNullable(this.region);
    }

    private GetTemDomainPlainArgs() {}

    private GetTemDomainPlainArgs(GetTemDomainPlainArgs $) {
        this.domainId = $.domainId;
        this.name = $.name;
        this.region = $.region;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetTemDomainPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetTemDomainPlainArgs $;

        public Builder() {
            $ = new GetTemDomainPlainArgs();
        }

        public Builder(GetTemDomainPlainArgs defaults) {
            $ = new GetTemDomainPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder domainId(@Nullable String domainId) {
            $.domainId = domainId;
            return this;
        }

        /**
         * @param name The domain name.
         * Only one of `name` and `id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable String name) {
            $.name = name;
            return this;
        }

        /**
         * @param region `region`) The region in which the domain exists.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable String region) {
            $.region = region;
            return this;
        }

        public GetTemDomainPlainArgs build() {
            return $;
        }
    }

}
