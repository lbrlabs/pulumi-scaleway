// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetLbAclsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetLbAclsPlainArgs Empty = new GetLbAclsPlainArgs();

    /**
     * The frontend ID this ACL is attached to. ACLs with a frontend ID like it are listed.
     * &gt; **Important:** LB Frontends&#39; IDs are zoned, which means they are of the form `{zone}/{id}`, e.g. `fr-par-1/11111111-1111-1111-1111-111111111111`
     * 
     */
    @Import(name="frontendId", required=true)
    private String frontendId;

    /**
     * @return The frontend ID this ACL is attached to. ACLs with a frontend ID like it are listed.
     * &gt; **Important:** LB Frontends&#39; IDs are zoned, which means they are of the form `{zone}/{id}`, e.g. `fr-par-1/11111111-1111-1111-1111-111111111111`
     * 
     */
    public String frontendId() {
        return this.frontendId;
    }

    /**
     * The ACL name used as filter. ACLs with a name like it are listed.
     * 
     */
    @Import(name="name")
    private @Nullable String name;

    /**
     * @return The ACL name used as filter. ACLs with a name like it are listed.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }

    @Import(name="projectId")
    private @Nullable String projectId;

    public Optional<String> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * `zone`) The zone in which ACLs exist.
     * 
     */
    @Import(name="zone")
    private @Nullable String zone;

    /**
     * @return `zone`) The zone in which ACLs exist.
     * 
     */
    public Optional<String> zone() {
        return Optional.ofNullable(this.zone);
    }

    private GetLbAclsPlainArgs() {}

    private GetLbAclsPlainArgs(GetLbAclsPlainArgs $) {
        this.frontendId = $.frontendId;
        this.name = $.name;
        this.projectId = $.projectId;
        this.zone = $.zone;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetLbAclsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetLbAclsPlainArgs $;

        public Builder() {
            $ = new GetLbAclsPlainArgs();
        }

        public Builder(GetLbAclsPlainArgs defaults) {
            $ = new GetLbAclsPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param frontendId The frontend ID this ACL is attached to. ACLs with a frontend ID like it are listed.
         * &gt; **Important:** LB Frontends&#39; IDs are zoned, which means they are of the form `{zone}/{id}`, e.g. `fr-par-1/11111111-1111-1111-1111-111111111111`
         * 
         * @return builder
         * 
         */
        public Builder frontendId(String frontendId) {
            $.frontendId = frontendId;
            return this;
        }

        /**
         * @param name The ACL name used as filter. ACLs with a name like it are listed.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable String name) {
            $.name = name;
            return this;
        }

        public Builder projectId(@Nullable String projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param zone `zone`) The zone in which ACLs exist.
         * 
         * @return builder
         * 
         */
        public Builder zone(@Nullable String zone) {
            $.zone = zone;
            return this;
        }

        public GetLbAclsPlainArgs build() {
            $.frontendId = Objects.requireNonNull($.frontendId, "expected parameter 'frontendId' to be non-null");
            return $;
        }
    }

}
