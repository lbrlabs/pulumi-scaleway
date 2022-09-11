// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetKubernetesClusterOpenIdConnectConfig {
    private String clientId;
    private List<String> groupsClaims;
    private String groupsPrefix;
    private String issuerUrl;
    private List<String> requiredClaims;
    private String usernameClaim;
    private String usernamePrefix;

    private GetKubernetesClusterOpenIdConnectConfig() {}
    public String clientId() {
        return this.clientId;
    }
    public List<String> groupsClaims() {
        return this.groupsClaims;
    }
    public String groupsPrefix() {
        return this.groupsPrefix;
    }
    public String issuerUrl() {
        return this.issuerUrl;
    }
    public List<String> requiredClaims() {
        return this.requiredClaims;
    }
    public String usernameClaim() {
        return this.usernameClaim;
    }
    public String usernamePrefix() {
        return this.usernamePrefix;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetKubernetesClusterOpenIdConnectConfig defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String clientId;
        private List<String> groupsClaims;
        private String groupsPrefix;
        private String issuerUrl;
        private List<String> requiredClaims;
        private String usernameClaim;
        private String usernamePrefix;
        public Builder() {}
        public Builder(GetKubernetesClusterOpenIdConnectConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.clientId = defaults.clientId;
    	      this.groupsClaims = defaults.groupsClaims;
    	      this.groupsPrefix = defaults.groupsPrefix;
    	      this.issuerUrl = defaults.issuerUrl;
    	      this.requiredClaims = defaults.requiredClaims;
    	      this.usernameClaim = defaults.usernameClaim;
    	      this.usernamePrefix = defaults.usernamePrefix;
        }

        @CustomType.Setter
        public Builder clientId(String clientId) {
            this.clientId = Objects.requireNonNull(clientId);
            return this;
        }
        @CustomType.Setter
        public Builder groupsClaims(List<String> groupsClaims) {
            this.groupsClaims = Objects.requireNonNull(groupsClaims);
            return this;
        }
        public Builder groupsClaims(String... groupsClaims) {
            return groupsClaims(List.of(groupsClaims));
        }
        @CustomType.Setter
        public Builder groupsPrefix(String groupsPrefix) {
            this.groupsPrefix = Objects.requireNonNull(groupsPrefix);
            return this;
        }
        @CustomType.Setter
        public Builder issuerUrl(String issuerUrl) {
            this.issuerUrl = Objects.requireNonNull(issuerUrl);
            return this;
        }
        @CustomType.Setter
        public Builder requiredClaims(List<String> requiredClaims) {
            this.requiredClaims = Objects.requireNonNull(requiredClaims);
            return this;
        }
        public Builder requiredClaims(String... requiredClaims) {
            return requiredClaims(List.of(requiredClaims));
        }
        @CustomType.Setter
        public Builder usernameClaim(String usernameClaim) {
            this.usernameClaim = Objects.requireNonNull(usernameClaim);
            return this;
        }
        @CustomType.Setter
        public Builder usernamePrefix(String usernamePrefix) {
            this.usernamePrefix = Objects.requireNonNull(usernamePrefix);
            return this;
        }
        public GetKubernetesClusterOpenIdConnectConfig build() {
            final var o = new GetKubernetesClusterOpenIdConnectConfig();
            o.clientId = clientId;
            o.groupsClaims = groupsClaims;
            o.groupsPrefix = groupsPrefix;
            o.issuerUrl = issuerUrl;
            o.requiredClaims = requiredClaims;
            o.usernameClaim = usernameClaim;
            o.usernamePrefix = usernamePrefix;
            return o;
        }
    }
}