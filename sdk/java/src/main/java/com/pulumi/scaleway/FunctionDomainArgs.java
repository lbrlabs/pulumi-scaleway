// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class FunctionDomainArgs extends com.pulumi.resources.ResourceArgs {

    public static final FunctionDomainArgs Empty = new FunctionDomainArgs();

    /**
     * The ID of the function you want to create a domain with.
     * 
     */
    @Import(name="functionId", required=true)
    private Output<String> functionId;

    /**
     * @return The ID of the function you want to create a domain with.
     * 
     */
    public Output<String> functionId() {
        return this.functionId;
    }

    /**
     * The hostname that should resolve to your function id native domain.
     * You should use a CNAME domain record that point to your native function `domain_name` for it.
     * 
     */
    @Import(name="hostname", required=true)
    private Output<String> hostname;

    /**
     * @return The hostname that should resolve to your function id native domain.
     * You should use a CNAME domain record that point to your native function `domain_name` for it.
     * 
     */
    public Output<String> hostname() {
        return this.hostname;
    }

    /**
     * (Defaults to provider `region`) The region in where the domain was created.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return (Defaults to provider `region`) The region in where the domain was created.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    private FunctionDomainArgs() {}

    private FunctionDomainArgs(FunctionDomainArgs $) {
        this.functionId = $.functionId;
        this.hostname = $.hostname;
        this.region = $.region;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FunctionDomainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FunctionDomainArgs $;

        public Builder() {
            $ = new FunctionDomainArgs();
        }

        public Builder(FunctionDomainArgs defaults) {
            $ = new FunctionDomainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param functionId The ID of the function you want to create a domain with.
         * 
         * @return builder
         * 
         */
        public Builder functionId(Output<String> functionId) {
            $.functionId = functionId;
            return this;
        }

        /**
         * @param functionId The ID of the function you want to create a domain with.
         * 
         * @return builder
         * 
         */
        public Builder functionId(String functionId) {
            return functionId(Output.of(functionId));
        }

        /**
         * @param hostname The hostname that should resolve to your function id native domain.
         * You should use a CNAME domain record that point to your native function `domain_name` for it.
         * 
         * @return builder
         * 
         */
        public Builder hostname(Output<String> hostname) {
            $.hostname = hostname;
            return this;
        }

        /**
         * @param hostname The hostname that should resolve to your function id native domain.
         * You should use a CNAME domain record that point to your native function `domain_name` for it.
         * 
         * @return builder
         * 
         */
        public Builder hostname(String hostname) {
            return hostname(Output.of(hostname));
        }

        /**
         * @param region (Defaults to provider `region`) The region in where the domain was created.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region (Defaults to provider `region`) The region in where the domain was created.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        public FunctionDomainArgs build() {
            $.functionId = Objects.requireNonNull($.functionId, "expected parameter 'functionId' to be non-null");
            $.hostname = Objects.requireNonNull($.hostname, "expected parameter 'hostname' to be non-null");
            return $;
        }
    }

}
