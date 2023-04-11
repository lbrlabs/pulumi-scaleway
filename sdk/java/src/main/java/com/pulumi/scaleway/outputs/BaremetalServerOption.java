// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class BaremetalServerOption {
    /**
     * @return The auto expiration date for compatible options
     * 
     */
    private @Nullable String expiresAt;
    /**
     * @return The id of the private network to attach.
     * 
     */
    private String id;
    /**
     * @return The name of the server.
     * 
     */
    private @Nullable String name;

    private BaremetalServerOption() {}
    /**
     * @return The auto expiration date for compatible options
     * 
     */
    public Optional<String> expiresAt() {
        return Optional.ofNullable(this.expiresAt);
    }
    /**
     * @return The id of the private network to attach.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The name of the server.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(BaremetalServerOption defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String expiresAt;
        private String id;
        private @Nullable String name;
        public Builder() {}
        public Builder(BaremetalServerOption defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.expiresAt = defaults.expiresAt;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
        }

        @CustomType.Setter
        public Builder expiresAt(@Nullable String expiresAt) {
            this.expiresAt = expiresAt;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder name(@Nullable String name) {
            this.name = name;
            return this;
        }
        public BaremetalServerOption build() {
            final var o = new BaremetalServerOption();
            o.expiresAt = expiresAt;
            o.id = id;
            o.name = name;
            return o;
        }
    }
}