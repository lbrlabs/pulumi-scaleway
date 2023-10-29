// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetDocumentdbInstanceArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetDocumentdbInstanceArgs Empty = new GetDocumentdbInstanceArgs();

    /**
     * The DocumentDB instance ID.
     * Only one of `name` and `instance_id` should be specified.
     * 
     */
    @Import(name="instanceId")
    private @Nullable Output<String> instanceId;

    /**
     * @return The DocumentDB instance ID.
     * Only one of `name` and `instance_id` should be specified.
     * 
     */
    public Optional<Output<String>> instanceId() {
        return Optional.ofNullable(this.instanceId);
    }

    /**
     * The name of the DocumentDB instance.
     * Only one of `name` and `instance_id` should be specified.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the DocumentDB instance.
     * Only one of `name` and `instance_id` should be specified.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * `region`) The region in which the DocumentDB instance exists.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return `region`) The region in which the DocumentDB instance exists.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    private GetDocumentdbInstanceArgs() {}

    private GetDocumentdbInstanceArgs(GetDocumentdbInstanceArgs $) {
        this.instanceId = $.instanceId;
        this.name = $.name;
        this.region = $.region;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetDocumentdbInstanceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetDocumentdbInstanceArgs $;

        public Builder() {
            $ = new GetDocumentdbInstanceArgs();
        }

        public Builder(GetDocumentdbInstanceArgs defaults) {
            $ = new GetDocumentdbInstanceArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param instanceId The DocumentDB instance ID.
         * Only one of `name` and `instance_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(@Nullable Output<String> instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        /**
         * @param instanceId The DocumentDB instance ID.
         * Only one of `name` and `instance_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(String instanceId) {
            return instanceId(Output.of(instanceId));
        }

        /**
         * @param name The name of the DocumentDB instance.
         * Only one of `name` and `instance_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the DocumentDB instance.
         * Only one of `name` and `instance_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param region `region`) The region in which the DocumentDB instance exists.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region `region`) The region in which the DocumentDB instance exists.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        public GetDocumentdbInstanceArgs build() {
            return $;
        }
    }

}
