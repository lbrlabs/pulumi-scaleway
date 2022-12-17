// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class IamApiKeyArgs extends com.pulumi.resources.ResourceArgs {

    public static final IamApiKeyArgs Empty = new IamApiKeyArgs();

    /**
     * ID of the application attached to the api key
     * 
     */
    @Import(name="applicationId")
    private @Nullable Output<String> applicationId;

    /**
     * @return ID of the application attached to the api key
     * 
     */
    public Optional<Output<String>> applicationId() {
        return Optional.ofNullable(this.applicationId);
    }

    /**
     * The default project ID to use with object storage.
     * 
     */
    @Import(name="defaultProjectId")
    private @Nullable Output<String> defaultProjectId;

    /**
     * @return The default project ID to use with object storage.
     * 
     */
    public Optional<Output<String>> defaultProjectId() {
        return Optional.ofNullable(this.defaultProjectId);
    }

    /**
     * The description of the iam api key
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the iam api key
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The date and time of the expiration of the iam api key. Please note that in case of change,
     * the resource will be recreated.
     * 
     */
    @Import(name="expiresAt")
    private @Nullable Output<String> expiresAt;

    /**
     * @return The date and time of the expiration of the iam api key. Please note that in case of change,
     * the resource will be recreated.
     * 
     */
    public Optional<Output<String>> expiresAt() {
        return Optional.ofNullable(this.expiresAt);
    }

    /**
     * ID of the user attached to the api key.
     * Only one of the `application_id` and `user_id` should be specified.
     * 
     */
    @Import(name="userId")
    private @Nullable Output<String> userId;

    /**
     * @return ID of the user attached to the api key.
     * Only one of the `application_id` and `user_id` should be specified.
     * 
     */
    public Optional<Output<String>> userId() {
        return Optional.ofNullable(this.userId);
    }

    private IamApiKeyArgs() {}

    private IamApiKeyArgs(IamApiKeyArgs $) {
        this.applicationId = $.applicationId;
        this.defaultProjectId = $.defaultProjectId;
        this.description = $.description;
        this.expiresAt = $.expiresAt;
        this.userId = $.userId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IamApiKeyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IamApiKeyArgs $;

        public Builder() {
            $ = new IamApiKeyArgs();
        }

        public Builder(IamApiKeyArgs defaults) {
            $ = new IamApiKeyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param applicationId ID of the application attached to the api key
         * 
         * @return builder
         * 
         */
        public Builder applicationId(@Nullable Output<String> applicationId) {
            $.applicationId = applicationId;
            return this;
        }

        /**
         * @param applicationId ID of the application attached to the api key
         * 
         * @return builder
         * 
         */
        public Builder applicationId(String applicationId) {
            return applicationId(Output.of(applicationId));
        }

        /**
         * @param defaultProjectId The default project ID to use with object storage.
         * 
         * @return builder
         * 
         */
        public Builder defaultProjectId(@Nullable Output<String> defaultProjectId) {
            $.defaultProjectId = defaultProjectId;
            return this;
        }

        /**
         * @param defaultProjectId The default project ID to use with object storage.
         * 
         * @return builder
         * 
         */
        public Builder defaultProjectId(String defaultProjectId) {
            return defaultProjectId(Output.of(defaultProjectId));
        }

        /**
         * @param description The description of the iam api key
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the iam api key
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param expiresAt The date and time of the expiration of the iam api key. Please note that in case of change,
         * the resource will be recreated.
         * 
         * @return builder
         * 
         */
        public Builder expiresAt(@Nullable Output<String> expiresAt) {
            $.expiresAt = expiresAt;
            return this;
        }

        /**
         * @param expiresAt The date and time of the expiration of the iam api key. Please note that in case of change,
         * the resource will be recreated.
         * 
         * @return builder
         * 
         */
        public Builder expiresAt(String expiresAt) {
            return expiresAt(Output.of(expiresAt));
        }

        /**
         * @param userId ID of the user attached to the api key.
         * Only one of the `application_id` and `user_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder userId(@Nullable Output<String> userId) {
            $.userId = userId;
            return this;
        }

        /**
         * @param userId ID of the user attached to the api key.
         * Only one of the `application_id` and `user_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder userId(String userId) {
            return userId(Output.of(userId));
        }

        public IamApiKeyArgs build() {
            return $;
        }
    }

}
