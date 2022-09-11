// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.scaleway.outputs.LoadbalancerFrontendAclAction;
import com.pulumi.scaleway.outputs.LoadbalancerFrontendAclMatch;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class LoadbalancerFrontendAcl {
    /**
     * @return Action to undertake when an ACL filter matches.
     * 
     */
    private LoadbalancerFrontendAclAction action;
    /**
     * @return The ACL match rule. At least `ip_subnet` or `http_filter` and `http_filter_value` are required.
     * 
     */
    private LoadbalancerFrontendAclMatch match;
    /**
     * @return The ACL name. If not provided it will be randomly generated.
     * 
     */
    private @Nullable String name;

    private LoadbalancerFrontendAcl() {}
    /**
     * @return Action to undertake when an ACL filter matches.
     * 
     */
    public LoadbalancerFrontendAclAction action() {
        return this.action;
    }
    /**
     * @return The ACL match rule. At least `ip_subnet` or `http_filter` and `http_filter_value` are required.
     * 
     */
    public LoadbalancerFrontendAclMatch match() {
        return this.match;
    }
    /**
     * @return The ACL name. If not provided it will be randomly generated.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(LoadbalancerFrontendAcl defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private LoadbalancerFrontendAclAction action;
        private LoadbalancerFrontendAclMatch match;
        private @Nullable String name;
        public Builder() {}
        public Builder(LoadbalancerFrontendAcl defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.action = defaults.action;
    	      this.match = defaults.match;
    	      this.name = defaults.name;
        }

        @CustomType.Setter
        public Builder action(LoadbalancerFrontendAclAction action) {
            this.action = Objects.requireNonNull(action);
            return this;
        }
        @CustomType.Setter
        public Builder match(LoadbalancerFrontendAclMatch match) {
            this.match = Objects.requireNonNull(match);
            return this;
        }
        @CustomType.Setter
        public Builder name(@Nullable String name) {
            this.name = name;
            return this;
        }
        public LoadbalancerFrontendAcl build() {
            final var o = new LoadbalancerFrontendAcl();
            o.action = action;
            o.match = match;
            o.name = name;
            return o;
        }
    }
}