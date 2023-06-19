// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class LoadbalancerCertificateCustomCertificate {
    /**
     * @return Full PEM-formatted certificate chain.
     * 
     * &gt; **Important:** Updates to `custom_certificate` will recreate the load-balancer certificate.
     * 
     */
    private String certificateChain;

    private LoadbalancerCertificateCustomCertificate() {}
    /**
     * @return Full PEM-formatted certificate chain.
     * 
     * &gt; **Important:** Updates to `custom_certificate` will recreate the load-balancer certificate.
     * 
     */
    public String certificateChain() {
        return this.certificateChain;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(LoadbalancerCertificateCustomCertificate defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String certificateChain;
        public Builder() {}
        public Builder(LoadbalancerCertificateCustomCertificate defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.certificateChain = defaults.certificateChain;
        }

        @CustomType.Setter
        public Builder certificateChain(String certificateChain) {
            this.certificateChain = Objects.requireNonNull(certificateChain);
            return this;
        }
        public LoadbalancerCertificateCustomCertificate build() {
            final var o = new LoadbalancerCertificateCustomCertificate();
            o.certificateChain = certificateChain;
            return o;
        }
    }
}
