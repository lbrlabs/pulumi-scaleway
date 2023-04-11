// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetIamUserArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetIamUserArgs Empty = new GetIamUserArgs();

    /**
     * The email address of the IAM user. Only one of the `email` and `user_id` should be specified.
     * 
     */
    @Import(name="email")
    private @Nullable Output<String> email;

    /**
     * @return The email address of the IAM user. Only one of the `email` and `user_id` should be specified.
     * 
     */
    public Optional<Output<String>> email() {
        return Optional.ofNullable(this.email);
    }

    /**
     * `organization_id`) The ID of the
     * organization the user is associated with.
     * 
     */
    @Import(name="organizationId")
    private @Nullable Output<String> organizationId;

    /**
     * @return `organization_id`) The ID of the
     * organization the user is associated with.
     * 
     */
    public Optional<Output<String>> organizationId() {
        return Optional.ofNullable(this.organizationId);
    }

    /**
     * The ID of the IAM user. Only one of the `email` and `user_id` should be specified.
     * 
     */
    @Import(name="userId")
    private @Nullable Output<String> userId;

    /**
     * @return The ID of the IAM user. Only one of the `email` and `user_id` should be specified.
     * 
     */
    public Optional<Output<String>> userId() {
        return Optional.ofNullable(this.userId);
    }

    private GetIamUserArgs() {}

    private GetIamUserArgs(GetIamUserArgs $) {
        this.email = $.email;
        this.organizationId = $.organizationId;
        this.userId = $.userId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetIamUserArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetIamUserArgs $;

        public Builder() {
            $ = new GetIamUserArgs();
        }

        public Builder(GetIamUserArgs defaults) {
            $ = new GetIamUserArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param email The email address of the IAM user. Only one of the `email` and `user_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder email(@Nullable Output<String> email) {
            $.email = email;
            return this;
        }

        /**
         * @param email The email address of the IAM user. Only one of the `email` and `user_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder email(String email) {
            return email(Output.of(email));
        }

        /**
         * @param organizationId `organization_id`) The ID of the
         * organization the user is associated with.
         * 
         * @return builder
         * 
         */
        public Builder organizationId(@Nullable Output<String> organizationId) {
            $.organizationId = organizationId;
            return this;
        }

        /**
         * @param organizationId `organization_id`) The ID of the
         * organization the user is associated with.
         * 
         * @return builder
         * 
         */
        public Builder organizationId(String organizationId) {
            return organizationId(Output.of(organizationId));
        }

        /**
         * @param userId The ID of the IAM user. Only one of the `email` and `user_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder userId(@Nullable Output<String> userId) {
            $.userId = userId;
            return this;
        }

        /**
         * @param userId The ID of the IAM user. Only one of the `email` and `user_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder userId(String userId) {
            return userId(Output.of(userId));
        }

        public GetIamUserArgs build() {
            return $;
        }
    }

}