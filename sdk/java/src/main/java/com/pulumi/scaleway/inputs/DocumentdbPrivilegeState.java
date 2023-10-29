// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DocumentdbPrivilegeState extends com.pulumi.resources.ResourceArgs {

    public static final DocumentdbPrivilegeState Empty = new DocumentdbPrivilegeState();

    /**
     * Name of the database (e.g. `my-db-name`).
     * 
     */
    @Import(name="databaseName")
    private @Nullable Output<String> databaseName;

    /**
     * @return Name of the database (e.g. `my-db-name`).
     * 
     */
    public Optional<Output<String>> databaseName() {
        return Optional.ofNullable(this.databaseName);
    }

    /**
     * UUID of the rdb instance.
     * 
     */
    @Import(name="instanceId")
    private @Nullable Output<String> instanceId;

    /**
     * @return UUID of the rdb instance.
     * 
     */
    public Optional<Output<String>> instanceId() {
        return Optional.ofNullable(this.instanceId);
    }

    /**
     * Permission to set. Valid values are `readonly`, `readwrite`, `all`, `custom` and `none`.
     * 
     */
    @Import(name="permission")
    private @Nullable Output<String> permission;

    /**
     * @return Permission to set. Valid values are `readonly`, `readwrite`, `all`, `custom` and `none`.
     * 
     */
    public Optional<Output<String>> permission() {
        return Optional.ofNullable(this.permission);
    }

    /**
     * `region`) The region in which the resource exists.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return `region`) The region in which the resource exists.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * Name of the user (e.g. `my-db-user`).
     * 
     */
    @Import(name="userName")
    private @Nullable Output<String> userName;

    /**
     * @return Name of the user (e.g. `my-db-user`).
     * 
     */
    public Optional<Output<String>> userName() {
        return Optional.ofNullable(this.userName);
    }

    private DocumentdbPrivilegeState() {}

    private DocumentdbPrivilegeState(DocumentdbPrivilegeState $) {
        this.databaseName = $.databaseName;
        this.instanceId = $.instanceId;
        this.permission = $.permission;
        this.region = $.region;
        this.userName = $.userName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DocumentdbPrivilegeState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DocumentdbPrivilegeState $;

        public Builder() {
            $ = new DocumentdbPrivilegeState();
        }

        public Builder(DocumentdbPrivilegeState defaults) {
            $ = new DocumentdbPrivilegeState(Objects.requireNonNull(defaults));
        }

        /**
         * @param databaseName Name of the database (e.g. `my-db-name`).
         * 
         * @return builder
         * 
         */
        public Builder databaseName(@Nullable Output<String> databaseName) {
            $.databaseName = databaseName;
            return this;
        }

        /**
         * @param databaseName Name of the database (e.g. `my-db-name`).
         * 
         * @return builder
         * 
         */
        public Builder databaseName(String databaseName) {
            return databaseName(Output.of(databaseName));
        }

        /**
         * @param instanceId UUID of the rdb instance.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(@Nullable Output<String> instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        /**
         * @param instanceId UUID of the rdb instance.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(String instanceId) {
            return instanceId(Output.of(instanceId));
        }

        /**
         * @param permission Permission to set. Valid values are `readonly`, `readwrite`, `all`, `custom` and `none`.
         * 
         * @return builder
         * 
         */
        public Builder permission(@Nullable Output<String> permission) {
            $.permission = permission;
            return this;
        }

        /**
         * @param permission Permission to set. Valid values are `readonly`, `readwrite`, `all`, `custom` and `none`.
         * 
         * @return builder
         * 
         */
        public Builder permission(String permission) {
            return permission(Output.of(permission));
        }

        /**
         * @param region `region`) The region in which the resource exists.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region `region`) The region in which the resource exists.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param userName Name of the user (e.g. `my-db-user`).
         * 
         * @return builder
         * 
         */
        public Builder userName(@Nullable Output<String> userName) {
            $.userName = userName;
            return this;
        }

        /**
         * @param userName Name of the user (e.g. `my-db-user`).
         * 
         * @return builder
         * 
         */
        public Builder userName(String userName) {
            return userName(Output.of(userName));
        }

        public DocumentdbPrivilegeState build() {
            return $;
        }
    }

}
