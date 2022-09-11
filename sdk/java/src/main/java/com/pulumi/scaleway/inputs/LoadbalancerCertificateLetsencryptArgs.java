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


public final class LoadbalancerCertificateLetsencryptArgs extends com.pulumi.resources.ResourceArgs {

    public static final LoadbalancerCertificateLetsencryptArgs Empty = new LoadbalancerCertificateLetsencryptArgs();

    /**
     * Main domain of the certificate. A new certificate will be created if this field is changed.
     * 
     */
    @Import(name="commonName", required=true)
    private Output<String> commonName;

    /**
     * @return Main domain of the certificate. A new certificate will be created if this field is changed.
     * 
     */
    public Output<String> commonName() {
        return this.commonName;
    }

    /**
     * Array of alternative domain names.  A new certificate will be created if this field is changed.
     * 
     */
    @Import(name="subjectAlternativeNames")
    private @Nullable Output<List<String>> subjectAlternativeNames;

    /**
     * @return Array of alternative domain names.  A new certificate will be created if this field is changed.
     * 
     */
    public Optional<Output<List<String>>> subjectAlternativeNames() {
        return Optional.ofNullable(this.subjectAlternativeNames);
    }

    private LoadbalancerCertificateLetsencryptArgs() {}

    private LoadbalancerCertificateLetsencryptArgs(LoadbalancerCertificateLetsencryptArgs $) {
        this.commonName = $.commonName;
        this.subjectAlternativeNames = $.subjectAlternativeNames;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LoadbalancerCertificateLetsencryptArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LoadbalancerCertificateLetsencryptArgs $;

        public Builder() {
            $ = new LoadbalancerCertificateLetsencryptArgs();
        }

        public Builder(LoadbalancerCertificateLetsencryptArgs defaults) {
            $ = new LoadbalancerCertificateLetsencryptArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param commonName Main domain of the certificate. A new certificate will be created if this field is changed.
         * 
         * @return builder
         * 
         */
        public Builder commonName(Output<String> commonName) {
            $.commonName = commonName;
            return this;
        }

        /**
         * @param commonName Main domain of the certificate. A new certificate will be created if this field is changed.
         * 
         * @return builder
         * 
         */
        public Builder commonName(String commonName) {
            return commonName(Output.of(commonName));
        }

        /**
         * @param subjectAlternativeNames Array of alternative domain names.  A new certificate will be created if this field is changed.
         * 
         * @return builder
         * 
         */
        public Builder subjectAlternativeNames(@Nullable Output<List<String>> subjectAlternativeNames) {
            $.subjectAlternativeNames = subjectAlternativeNames;
            return this;
        }

        /**
         * @param subjectAlternativeNames Array of alternative domain names.  A new certificate will be created if this field is changed.
         * 
         * @return builder
         * 
         */
        public Builder subjectAlternativeNames(List<String> subjectAlternativeNames) {
            return subjectAlternativeNames(Output.of(subjectAlternativeNames));
        }

        /**
         * @param subjectAlternativeNames Array of alternative domain names.  A new certificate will be created if this field is changed.
         * 
         * @return builder
         * 
         */
        public Builder subjectAlternativeNames(String... subjectAlternativeNames) {
            return subjectAlternativeNames(List.of(subjectAlternativeNames));
        }

        public LoadbalancerCertificateLetsencryptArgs build() {
            $.commonName = Objects.requireNonNull($.commonName, "expected parameter 'commonName' to be non-null");
            return $;
        }
    }

}