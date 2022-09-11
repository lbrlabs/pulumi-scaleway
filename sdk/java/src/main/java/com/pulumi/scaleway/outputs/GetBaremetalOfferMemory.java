// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetBaremetalOfferMemory {
    private Integer capacity;
    /**
     * @return Frequency of the memory in MHz.
     * 
     */
    private Integer frequency;
    private Boolean isEcc;
    /**
     * @return Type of memory.
     * 
     */
    private String type;

    private GetBaremetalOfferMemory() {}
    public Integer capacity() {
        return this.capacity;
    }
    /**
     * @return Frequency of the memory in MHz.
     * 
     */
    public Integer frequency() {
        return this.frequency;
    }
    public Boolean isEcc() {
        return this.isEcc;
    }
    /**
     * @return Type of memory.
     * 
     */
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetBaremetalOfferMemory defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer capacity;
        private Integer frequency;
        private Boolean isEcc;
        private String type;
        public Builder() {}
        public Builder(GetBaremetalOfferMemory defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.capacity = defaults.capacity;
    	      this.frequency = defaults.frequency;
    	      this.isEcc = defaults.isEcc;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder capacity(Integer capacity) {
            this.capacity = Objects.requireNonNull(capacity);
            return this;
        }
        @CustomType.Setter
        public Builder frequency(Integer frequency) {
            this.frequency = Objects.requireNonNull(frequency);
            return this;
        }
        @CustomType.Setter
        public Builder isEcc(Boolean isEcc) {
            this.isEcc = Objects.requireNonNull(isEcc);
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            this.type = Objects.requireNonNull(type);
            return this;
        }
        public GetBaremetalOfferMemory build() {
            final var o = new GetBaremetalOfferMemory();
            o.capacity = capacity;
            o.frequency = frequency;
            o.isEcc = isEcc;
            o.type = type;
            return o;
        }
    }
}