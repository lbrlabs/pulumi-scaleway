// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.scaleway.inputs.ObjectBucketWebsiteConfigurationErrorDocumentArgs;
import com.pulumi.scaleway.inputs.ObjectBucketWebsiteConfigurationIndexDocumentArgs;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ObjectBucketWebsiteConfigurationState extends com.pulumi.resources.ResourceArgs {

    public static final ObjectBucketWebsiteConfigurationState Empty = new ObjectBucketWebsiteConfigurationState();

    /**
     * (Required, Forces new resource) The name of the bucket.
     * 
     */
    @Import(name="bucket")
    private @Nullable Output<String> bucket;

    /**
     * @return (Required, Forces new resource) The name of the bucket.
     * 
     */
    public Optional<Output<String>> bucket() {
        return Optional.ofNullable(this.bucket);
    }

    /**
     * (Optional) The name of the error document for the website detailed below.
     * 
     */
    @Import(name="errorDocument")
    private @Nullable Output<ObjectBucketWebsiteConfigurationErrorDocumentArgs> errorDocument;

    /**
     * @return (Optional) The name of the error document for the website detailed below.
     * 
     */
    public Optional<Output<ObjectBucketWebsiteConfigurationErrorDocumentArgs>> errorDocument() {
        return Optional.ofNullable(this.errorDocument);
    }

    /**
     * (Required) The name of the index document for the website detailed below.
     * 
     */
    @Import(name="indexDocument")
    private @Nullable Output<ObjectBucketWebsiteConfigurationIndexDocumentArgs> indexDocument;

    /**
     * @return (Required) The name of the index document for the website detailed below.
     * 
     */
    public Optional<Output<ObjectBucketWebsiteConfigurationIndexDocumentArgs>> indexDocument() {
        return Optional.ofNullable(this.indexDocument);
    }

    /**
     * (Defaults to provider `project_id`) The ID of the project the bucket is associated with.
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return (Defaults to provider `project_id`) The ID of the project the bucket is associated with.
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * The region you want to attach the resource to
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region you want to attach the resource to
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * The website endpoint.
     * 
     */
    @Import(name="websiteDomain")
    private @Nullable Output<String> websiteDomain;

    /**
     * @return The website endpoint.
     * 
     */
    public Optional<Output<String>> websiteDomain() {
        return Optional.ofNullable(this.websiteDomain);
    }

    /**
     * The domain of the website endpoint.
     * 
     */
    @Import(name="websiteEndpoint")
    private @Nullable Output<String> websiteEndpoint;

    /**
     * @return The domain of the website endpoint.
     * 
     */
    public Optional<Output<String>> websiteEndpoint() {
        return Optional.ofNullable(this.websiteEndpoint);
    }

    private ObjectBucketWebsiteConfigurationState() {}

    private ObjectBucketWebsiteConfigurationState(ObjectBucketWebsiteConfigurationState $) {
        this.bucket = $.bucket;
        this.errorDocument = $.errorDocument;
        this.indexDocument = $.indexDocument;
        this.projectId = $.projectId;
        this.region = $.region;
        this.websiteDomain = $.websiteDomain;
        this.websiteEndpoint = $.websiteEndpoint;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ObjectBucketWebsiteConfigurationState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ObjectBucketWebsiteConfigurationState $;

        public Builder() {
            $ = new ObjectBucketWebsiteConfigurationState();
        }

        public Builder(ObjectBucketWebsiteConfigurationState defaults) {
            $ = new ObjectBucketWebsiteConfigurationState(Objects.requireNonNull(defaults));
        }

        /**
         * @param bucket (Required, Forces new resource) The name of the bucket.
         * 
         * @return builder
         * 
         */
        public Builder bucket(@Nullable Output<String> bucket) {
            $.bucket = bucket;
            return this;
        }

        /**
         * @param bucket (Required, Forces new resource) The name of the bucket.
         * 
         * @return builder
         * 
         */
        public Builder bucket(String bucket) {
            return bucket(Output.of(bucket));
        }

        /**
         * @param errorDocument (Optional) The name of the error document for the website detailed below.
         * 
         * @return builder
         * 
         */
        public Builder errorDocument(@Nullable Output<ObjectBucketWebsiteConfigurationErrorDocumentArgs> errorDocument) {
            $.errorDocument = errorDocument;
            return this;
        }

        /**
         * @param errorDocument (Optional) The name of the error document for the website detailed below.
         * 
         * @return builder
         * 
         */
        public Builder errorDocument(ObjectBucketWebsiteConfigurationErrorDocumentArgs errorDocument) {
            return errorDocument(Output.of(errorDocument));
        }

        /**
         * @param indexDocument (Required) The name of the index document for the website detailed below.
         * 
         * @return builder
         * 
         */
        public Builder indexDocument(@Nullable Output<ObjectBucketWebsiteConfigurationIndexDocumentArgs> indexDocument) {
            $.indexDocument = indexDocument;
            return this;
        }

        /**
         * @param indexDocument (Required) The name of the index document for the website detailed below.
         * 
         * @return builder
         * 
         */
        public Builder indexDocument(ObjectBucketWebsiteConfigurationIndexDocumentArgs indexDocument) {
            return indexDocument(Output.of(indexDocument));
        }

        /**
         * @param projectId (Defaults to provider `project_id`) The ID of the project the bucket is associated with.
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId (Defaults to provider `project_id`) The ID of the project the bucket is associated with.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param region The region you want to attach the resource to
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The region you want to attach the resource to
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param websiteDomain The website endpoint.
         * 
         * @return builder
         * 
         */
        public Builder websiteDomain(@Nullable Output<String> websiteDomain) {
            $.websiteDomain = websiteDomain;
            return this;
        }

        /**
         * @param websiteDomain The website endpoint.
         * 
         * @return builder
         * 
         */
        public Builder websiteDomain(String websiteDomain) {
            return websiteDomain(Output.of(websiteDomain));
        }

        /**
         * @param websiteEndpoint The domain of the website endpoint.
         * 
         * @return builder
         * 
         */
        public Builder websiteEndpoint(@Nullable Output<String> websiteEndpoint) {
            $.websiteEndpoint = websiteEndpoint;
            return this;
        }

        /**
         * @param websiteEndpoint The domain of the website endpoint.
         * 
         * @return builder
         * 
         */
        public Builder websiteEndpoint(String websiteEndpoint) {
            return websiteEndpoint(Output.of(websiteEndpoint));
        }

        public ObjectBucketWebsiteConfigurationState build() {
            return $;
        }
    }

}
