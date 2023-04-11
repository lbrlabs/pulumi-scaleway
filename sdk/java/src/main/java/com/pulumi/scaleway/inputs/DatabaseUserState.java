// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DatabaseUserState extends com.pulumi.resources.ResourceArgs {

    public static final DatabaseUserState Empty = new DatabaseUserState();

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
     * Grant admin permissions to the Database User.
     * 
     */
    @Import(name="isAdmin")
    private @Nullable Output<Boolean> isAdmin;

    /**
     * @return Grant admin permissions to the Database User.
     * 
     */
    public Optional<Output<Boolean>> isAdmin() {
        return Optional.ofNullable(this.isAdmin);
    }

    /**
     * Database User name.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Database User name.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Database User password.
     * 
     */
    @Import(name="password")
    private @Nullable Output<String> password;

    /**
     * @return Database User password.
     * 
     */
    public Optional<Output<String>> password() {
        return Optional.ofNullable(this.password);
    }

    /**
     * The Scaleway region this resource resides in.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The Scaleway region this resource resides in.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    private DatabaseUserState() {}

    private DatabaseUserState(DatabaseUserState $) {
        this.instanceId = $.instanceId;
        this.isAdmin = $.isAdmin;
        this.name = $.name;
        this.password = $.password;
        this.region = $.region;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DatabaseUserState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DatabaseUserState $;

        public Builder() {
            $ = new DatabaseUserState();
        }

        public Builder(DatabaseUserState defaults) {
            $ = new DatabaseUserState(Objects.requireNonNull(defaults));
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
         * @param isAdmin Grant admin permissions to the Database User.
         * 
         * @return builder
         * 
         */
        public Builder isAdmin(@Nullable Output<Boolean> isAdmin) {
            $.isAdmin = isAdmin;
            return this;
        }

        /**
         * @param isAdmin Grant admin permissions to the Database User.
         * 
         * @return builder
         * 
         */
        public Builder isAdmin(Boolean isAdmin) {
            return isAdmin(Output.of(isAdmin));
        }

        /**
         * @param name Database User name.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Database User name.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param password Database User password.
         * 
         * @return builder
         * 
         */
        public Builder password(@Nullable Output<String> password) {
            $.password = password;
            return this;
        }

        /**
         * @param password Database User password.
         * 
         * @return builder
         * 
         */
        public Builder password(String password) {
            return password(Output.of(password));
        }

        /**
         * @param region The Scaleway region this resource resides in.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The Scaleway region this resource resides in.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        public DatabaseUserState build() {
            return $;
        }
    }

}
