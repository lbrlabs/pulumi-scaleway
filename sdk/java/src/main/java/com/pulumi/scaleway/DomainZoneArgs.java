// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DomainZoneArgs extends com.pulumi.resources.ResourceArgs {

    public static final DomainZoneArgs Empty = new DomainZoneArgs();

    /**
     * The domain where the DNS zone will be created.
     * 
     */
    @Import(name="domain", required=true)
    private Output<String> domain;

    /**
     * @return The domain where the DNS zone will be created.
     * 
     */
    public Output<String> domain() {
        return this.domain;
    }

    /**
     * `project_id`) The ID of the project the domain is associated with.
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return `project_id`) The ID of the project the domain is associated with.
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * The subdomain(zone name) to create in the domain.
     * 
     */
    @Import(name="subdomain", required=true)
    private Output<String> subdomain;

    /**
     * @return The subdomain(zone name) to create in the domain.
     * 
     */
    public Output<String> subdomain() {
        return this.subdomain;
    }

    private DomainZoneArgs() {}

    private DomainZoneArgs(DomainZoneArgs $) {
        this.domain = $.domain;
        this.projectId = $.projectId;
        this.subdomain = $.subdomain;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DomainZoneArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DomainZoneArgs $;

        public Builder() {
            $ = new DomainZoneArgs();
        }

        public Builder(DomainZoneArgs defaults) {
            $ = new DomainZoneArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param domain The domain where the DNS zone will be created.
         * 
         * @return builder
         * 
         */
        public Builder domain(Output<String> domain) {
            $.domain = domain;
            return this;
        }

        /**
         * @param domain The domain where the DNS zone will be created.
         * 
         * @return builder
         * 
         */
        public Builder domain(String domain) {
            return domain(Output.of(domain));
        }

        /**
         * @param projectId `project_id`) The ID of the project the domain is associated with.
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId `project_id`) The ID of the project the domain is associated with.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param subdomain The subdomain(zone name) to create in the domain.
         * 
         * @return builder
         * 
         */
        public Builder subdomain(Output<String> subdomain) {
            $.subdomain = subdomain;
            return this;
        }

        /**
         * @param subdomain The subdomain(zone name) to create in the domain.
         * 
         * @return builder
         * 
         */
        public Builder subdomain(String subdomain) {
            return subdomain(Output.of(subdomain));
        }

        public DomainZoneArgs build() {
            $.domain = Objects.requireNonNull($.domain, "expected parameter 'domain' to be non-null");
            $.subdomain = Objects.requireNonNull($.subdomain, "expected parameter 'subdomain' to be non-null");
            return $;
        }
    }

}