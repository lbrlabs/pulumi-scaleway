// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class MnqCredentialSqsSnsCredentialsPermissionsArgs extends com.pulumi.resources.ResourceArgs {

    public static final MnqCredentialSqsSnsCredentialsPermissionsArgs Empty = new MnqCredentialSqsSnsCredentialsPermissionsArgs();

    /**
     * . Defines if user can manage the associated resource(s).
     * 
     */
    @Import(name="canManage")
    private @Nullable Output<Boolean> canManage;

    /**
     * @return . Defines if user can manage the associated resource(s).
     * 
     */
    public Optional<Output<Boolean>> canManage() {
        return Optional.ofNullable(this.canManage);
    }

    /**
     * . Defines if user can publish messages to the service.
     * 
     */
    @Import(name="canPublish")
    private @Nullable Output<Boolean> canPublish;

    /**
     * @return . Defines if user can publish messages to the service.
     * 
     */
    public Optional<Output<Boolean>> canPublish() {
        return Optional.ofNullable(this.canPublish);
    }

    /**
     * . Defines if user can receive messages from the service.
     * 
     */
    @Import(name="canReceive")
    private @Nullable Output<Boolean> canReceive;

    /**
     * @return . Defines if user can receive messages from the service.
     * 
     */
    public Optional<Output<Boolean>> canReceive() {
        return Optional.ofNullable(this.canReceive);
    }

    private MnqCredentialSqsSnsCredentialsPermissionsArgs() {}

    private MnqCredentialSqsSnsCredentialsPermissionsArgs(MnqCredentialSqsSnsCredentialsPermissionsArgs $) {
        this.canManage = $.canManage;
        this.canPublish = $.canPublish;
        this.canReceive = $.canReceive;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(MnqCredentialSqsSnsCredentialsPermissionsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private MnqCredentialSqsSnsCredentialsPermissionsArgs $;

        public Builder() {
            $ = new MnqCredentialSqsSnsCredentialsPermissionsArgs();
        }

        public Builder(MnqCredentialSqsSnsCredentialsPermissionsArgs defaults) {
            $ = new MnqCredentialSqsSnsCredentialsPermissionsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param canManage . Defines if user can manage the associated resource(s).
         * 
         * @return builder
         * 
         */
        public Builder canManage(@Nullable Output<Boolean> canManage) {
            $.canManage = canManage;
            return this;
        }

        /**
         * @param canManage . Defines if user can manage the associated resource(s).
         * 
         * @return builder
         * 
         */
        public Builder canManage(Boolean canManage) {
            return canManage(Output.of(canManage));
        }

        /**
         * @param canPublish . Defines if user can publish messages to the service.
         * 
         * @return builder
         * 
         */
        public Builder canPublish(@Nullable Output<Boolean> canPublish) {
            $.canPublish = canPublish;
            return this;
        }

        /**
         * @param canPublish . Defines if user can publish messages to the service.
         * 
         * @return builder
         * 
         */
        public Builder canPublish(Boolean canPublish) {
            return canPublish(Output.of(canPublish));
        }

        /**
         * @param canReceive . Defines if user can receive messages from the service.
         * 
         * @return builder
         * 
         */
        public Builder canReceive(@Nullable Output<Boolean> canReceive) {
            $.canReceive = canReceive;
            return this;
        }

        /**
         * @param canReceive . Defines if user can receive messages from the service.
         * 
         * @return builder
         * 
         */
        public Builder canReceive(Boolean canReceive) {
            return canReceive(Output.of(canReceive));
        }

        public MnqCredentialSqsSnsCredentialsPermissionsArgs build() {
            return $;
        }
    }

}
