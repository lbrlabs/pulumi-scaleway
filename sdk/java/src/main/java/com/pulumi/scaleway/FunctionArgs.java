// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class FunctionArgs extends com.pulumi.resources.ResourceArgs {

    public static final FunctionArgs Empty = new FunctionArgs();

    /**
     * Define if the function should be deployed, terraform will wait for function to be deployed
     * 
     */
    @Import(name="deploy")
    private @Nullable Output<Boolean> deploy;

    /**
     * @return Define if the function should be deployed, terraform will wait for function to be deployed
     * 
     */
    public Optional<Output<Boolean>> deploy() {
        return Optional.ofNullable(this.deploy);
    }

    /**
     * The description of the function.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the function.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The environment variables of the function.
     * 
     */
    @Import(name="environmentVariables")
    private @Nullable Output<Map<String,String>> environmentVariables;

    /**
     * @return The environment variables of the function.
     * 
     */
    public Optional<Output<Map<String,String>>> environmentVariables() {
        return Optional.ofNullable(this.environmentVariables);
    }

    /**
     * Handler of the function. Depends on the runtime ([function guide](https://developers.scaleway.com/en/products/functions/api/#create-a-function))
     * 
     */
    @Import(name="handler", required=true)
    private Output<String> handler;

    /**
     * @return Handler of the function. Depends on the runtime ([function guide](https://developers.scaleway.com/en/products/functions/api/#create-a-function))
     * 
     */
    public Output<String> handler() {
        return this.handler;
    }

    /**
     * HTTP traffic configuration
     * 
     */
    @Import(name="httpOption")
    private @Nullable Output<String> httpOption;

    /**
     * @return HTTP traffic configuration
     * 
     */
    public Optional<Output<String>> httpOption() {
        return Optional.ofNullable(this.httpOption);
    }

    /**
     * Maximum replicas for your function (defaults to 20), our system will scale your functions automatically based on incoming workload, but will never scale the number of replicas above the configured max_scale.
     * 
     */
    @Import(name="maxScale")
    private @Nullable Output<Integer> maxScale;

    /**
     * @return Maximum replicas for your function (defaults to 20), our system will scale your functions automatically based on incoming workload, but will never scale the number of replicas above the configured max_scale.
     * 
     */
    public Optional<Output<Integer>> maxScale() {
        return Optional.ofNullable(this.maxScale);
    }

    /**
     * Memory limit in MB for your function, defaults to 128MB
     * 
     */
    @Import(name="memoryLimit")
    private @Nullable Output<Integer> memoryLimit;

    /**
     * @return Memory limit in MB for your function, defaults to 128MB
     * 
     */
    public Optional<Output<Integer>> memoryLimit() {
        return Optional.ofNullable(this.memoryLimit);
    }

    /**
     * Minimum replicas for your function, defaults to 0, Note that a function is billed when it gets executed, and using a min_scale greater than 0 will cause your function container to run constantly.
     * 
     */
    @Import(name="minScale")
    private @Nullable Output<Integer> minScale;

    /**
     * @return Minimum replicas for your function, defaults to 0, Note that a function is billed when it gets executed, and using a min_scale greater than 0 will cause your function container to run constantly.
     * 
     */
    public Optional<Output<Integer>> minScale() {
        return Optional.ofNullable(this.minScale);
    }

    /**
     * The unique name of the function.
     * 
     * &gt; **Important** Updates to `name` will recreate the function.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The unique name of the function.
     * 
     * &gt; **Important** Updates to `name` will recreate the function.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The namespace ID associated with this function
     * 
     */
    @Import(name="namespaceId", required=true)
    private Output<String> namespaceId;

    /**
     * @return The namespace ID associated with this function
     * 
     */
    public Output<String> namespaceId() {
        return this.namespaceId;
    }

    /**
     * Privacy of the function. Can be either `private` or `public`. Read more on [authentication](https://developers.scaleway.com/en/products/functions/api/#authentication)
     * 
     */
    @Import(name="privacy", required=true)
    private Output<String> privacy;

    /**
     * @return Privacy of the function. Can be either `private` or `public`. Read more on [authentication](https://developers.scaleway.com/en/products/functions/api/#authentication)
     * 
     */
    public Output<String> privacy() {
        return this.privacy;
    }

    /**
     * `project_id`) The ID of the project the namespace is associated with.
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return `project_id`) The ID of the project the namespace is associated with.
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * `region`). The region in which the namespace should be created.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return `region`). The region in which the namespace should be created.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * Runtime of the function. Runtimes can be fetched using [specific route](https://developers.scaleway.com/en/products/functions/api/#get-f7de6a)
     * 
     */
    @Import(name="runtime", required=true)
    private Output<String> runtime;

    /**
     * @return Runtime of the function. Runtimes can be fetched using [specific route](https://developers.scaleway.com/en/products/functions/api/#get-f7de6a)
     * 
     */
    public Output<String> runtime() {
        return this.runtime;
    }

    /**
     * The [secret environment](https://www.scaleway.com/en/docs/compute/functions/concepts/#secrets) variables of the function.
     * 
     */
    @Import(name="secretEnvironmentVariables")
    private @Nullable Output<Map<String,String>> secretEnvironmentVariables;

    /**
     * @return The [secret environment](https://www.scaleway.com/en/docs/compute/functions/concepts/#secrets) variables of the function.
     * 
     */
    public Optional<Output<Map<String,String>>> secretEnvironmentVariables() {
        return Optional.ofNullable(this.secretEnvironmentVariables);
    }

    /**
     * Holds the max duration (in seconds) the function is allowed for responding to a request
     * 
     */
    @Import(name="timeout")
    private @Nullable Output<Integer> timeout;

    /**
     * @return Holds the max duration (in seconds) the function is allowed for responding to a request
     * 
     */
    public Optional<Output<Integer>> timeout() {
        return Optional.ofNullable(this.timeout);
    }

    /**
     * Location of the zip file to upload containing your function sources
     * 
     */
    @Import(name="zipFile")
    private @Nullable Output<String> zipFile;

    /**
     * @return Location of the zip file to upload containing your function sources
     * 
     */
    public Optional<Output<String>> zipFile() {
        return Optional.ofNullable(this.zipFile);
    }

    /**
     * The hash of your source zip file, changing it will re-apply function. Can be any string
     * 
     */
    @Import(name="zipHash")
    private @Nullable Output<String> zipHash;

    /**
     * @return The hash of your source zip file, changing it will re-apply function. Can be any string
     * 
     */
    public Optional<Output<String>> zipHash() {
        return Optional.ofNullable(this.zipHash);
    }

    private FunctionArgs() {}

    private FunctionArgs(FunctionArgs $) {
        this.deploy = $.deploy;
        this.description = $.description;
        this.environmentVariables = $.environmentVariables;
        this.handler = $.handler;
        this.httpOption = $.httpOption;
        this.maxScale = $.maxScale;
        this.memoryLimit = $.memoryLimit;
        this.minScale = $.minScale;
        this.name = $.name;
        this.namespaceId = $.namespaceId;
        this.privacy = $.privacy;
        this.projectId = $.projectId;
        this.region = $.region;
        this.runtime = $.runtime;
        this.secretEnvironmentVariables = $.secretEnvironmentVariables;
        this.timeout = $.timeout;
        this.zipFile = $.zipFile;
        this.zipHash = $.zipHash;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FunctionArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FunctionArgs $;

        public Builder() {
            $ = new FunctionArgs();
        }

        public Builder(FunctionArgs defaults) {
            $ = new FunctionArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param deploy Define if the function should be deployed, terraform will wait for function to be deployed
         * 
         * @return builder
         * 
         */
        public Builder deploy(@Nullable Output<Boolean> deploy) {
            $.deploy = deploy;
            return this;
        }

        /**
         * @param deploy Define if the function should be deployed, terraform will wait for function to be deployed
         * 
         * @return builder
         * 
         */
        public Builder deploy(Boolean deploy) {
            return deploy(Output.of(deploy));
        }

        /**
         * @param description The description of the function.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the function.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param environmentVariables The environment variables of the function.
         * 
         * @return builder
         * 
         */
        public Builder environmentVariables(@Nullable Output<Map<String,String>> environmentVariables) {
            $.environmentVariables = environmentVariables;
            return this;
        }

        /**
         * @param environmentVariables The environment variables of the function.
         * 
         * @return builder
         * 
         */
        public Builder environmentVariables(Map<String,String> environmentVariables) {
            return environmentVariables(Output.of(environmentVariables));
        }

        /**
         * @param handler Handler of the function. Depends on the runtime ([function guide](https://developers.scaleway.com/en/products/functions/api/#create-a-function))
         * 
         * @return builder
         * 
         */
        public Builder handler(Output<String> handler) {
            $.handler = handler;
            return this;
        }

        /**
         * @param handler Handler of the function. Depends on the runtime ([function guide](https://developers.scaleway.com/en/products/functions/api/#create-a-function))
         * 
         * @return builder
         * 
         */
        public Builder handler(String handler) {
            return handler(Output.of(handler));
        }

        /**
         * @param httpOption HTTP traffic configuration
         * 
         * @return builder
         * 
         */
        public Builder httpOption(@Nullable Output<String> httpOption) {
            $.httpOption = httpOption;
            return this;
        }

        /**
         * @param httpOption HTTP traffic configuration
         * 
         * @return builder
         * 
         */
        public Builder httpOption(String httpOption) {
            return httpOption(Output.of(httpOption));
        }

        /**
         * @param maxScale Maximum replicas for your function (defaults to 20), our system will scale your functions automatically based on incoming workload, but will never scale the number of replicas above the configured max_scale.
         * 
         * @return builder
         * 
         */
        public Builder maxScale(@Nullable Output<Integer> maxScale) {
            $.maxScale = maxScale;
            return this;
        }

        /**
         * @param maxScale Maximum replicas for your function (defaults to 20), our system will scale your functions automatically based on incoming workload, but will never scale the number of replicas above the configured max_scale.
         * 
         * @return builder
         * 
         */
        public Builder maxScale(Integer maxScale) {
            return maxScale(Output.of(maxScale));
        }

        /**
         * @param memoryLimit Memory limit in MB for your function, defaults to 128MB
         * 
         * @return builder
         * 
         */
        public Builder memoryLimit(@Nullable Output<Integer> memoryLimit) {
            $.memoryLimit = memoryLimit;
            return this;
        }

        /**
         * @param memoryLimit Memory limit in MB for your function, defaults to 128MB
         * 
         * @return builder
         * 
         */
        public Builder memoryLimit(Integer memoryLimit) {
            return memoryLimit(Output.of(memoryLimit));
        }

        /**
         * @param minScale Minimum replicas for your function, defaults to 0, Note that a function is billed when it gets executed, and using a min_scale greater than 0 will cause your function container to run constantly.
         * 
         * @return builder
         * 
         */
        public Builder minScale(@Nullable Output<Integer> minScale) {
            $.minScale = minScale;
            return this;
        }

        /**
         * @param minScale Minimum replicas for your function, defaults to 0, Note that a function is billed when it gets executed, and using a min_scale greater than 0 will cause your function container to run constantly.
         * 
         * @return builder
         * 
         */
        public Builder minScale(Integer minScale) {
            return minScale(Output.of(minScale));
        }

        /**
         * @param name The unique name of the function.
         * 
         * &gt; **Important** Updates to `name` will recreate the function.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The unique name of the function.
         * 
         * &gt; **Important** Updates to `name` will recreate the function.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param namespaceId The namespace ID associated with this function
         * 
         * @return builder
         * 
         */
        public Builder namespaceId(Output<String> namespaceId) {
            $.namespaceId = namespaceId;
            return this;
        }

        /**
         * @param namespaceId The namespace ID associated with this function
         * 
         * @return builder
         * 
         */
        public Builder namespaceId(String namespaceId) {
            return namespaceId(Output.of(namespaceId));
        }

        /**
         * @param privacy Privacy of the function. Can be either `private` or `public`. Read more on [authentication](https://developers.scaleway.com/en/products/functions/api/#authentication)
         * 
         * @return builder
         * 
         */
        public Builder privacy(Output<String> privacy) {
            $.privacy = privacy;
            return this;
        }

        /**
         * @param privacy Privacy of the function. Can be either `private` or `public`. Read more on [authentication](https://developers.scaleway.com/en/products/functions/api/#authentication)
         * 
         * @return builder
         * 
         */
        public Builder privacy(String privacy) {
            return privacy(Output.of(privacy));
        }

        /**
         * @param projectId `project_id`) The ID of the project the namespace is associated with.
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId `project_id`) The ID of the project the namespace is associated with.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param region `region`). The region in which the namespace should be created.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region `region`). The region in which the namespace should be created.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param runtime Runtime of the function. Runtimes can be fetched using [specific route](https://developers.scaleway.com/en/products/functions/api/#get-f7de6a)
         * 
         * @return builder
         * 
         */
        public Builder runtime(Output<String> runtime) {
            $.runtime = runtime;
            return this;
        }

        /**
         * @param runtime Runtime of the function. Runtimes can be fetched using [specific route](https://developers.scaleway.com/en/products/functions/api/#get-f7de6a)
         * 
         * @return builder
         * 
         */
        public Builder runtime(String runtime) {
            return runtime(Output.of(runtime));
        }

        /**
         * @param secretEnvironmentVariables The [secret environment](https://www.scaleway.com/en/docs/compute/functions/concepts/#secrets) variables of the function.
         * 
         * @return builder
         * 
         */
        public Builder secretEnvironmentVariables(@Nullable Output<Map<String,String>> secretEnvironmentVariables) {
            $.secretEnvironmentVariables = secretEnvironmentVariables;
            return this;
        }

        /**
         * @param secretEnvironmentVariables The [secret environment](https://www.scaleway.com/en/docs/compute/functions/concepts/#secrets) variables of the function.
         * 
         * @return builder
         * 
         */
        public Builder secretEnvironmentVariables(Map<String,String> secretEnvironmentVariables) {
            return secretEnvironmentVariables(Output.of(secretEnvironmentVariables));
        }

        /**
         * @param timeout Holds the max duration (in seconds) the function is allowed for responding to a request
         * 
         * @return builder
         * 
         */
        public Builder timeout(@Nullable Output<Integer> timeout) {
            $.timeout = timeout;
            return this;
        }

        /**
         * @param timeout Holds the max duration (in seconds) the function is allowed for responding to a request
         * 
         * @return builder
         * 
         */
        public Builder timeout(Integer timeout) {
            return timeout(Output.of(timeout));
        }

        /**
         * @param zipFile Location of the zip file to upload containing your function sources
         * 
         * @return builder
         * 
         */
        public Builder zipFile(@Nullable Output<String> zipFile) {
            $.zipFile = zipFile;
            return this;
        }

        /**
         * @param zipFile Location of the zip file to upload containing your function sources
         * 
         * @return builder
         * 
         */
        public Builder zipFile(String zipFile) {
            return zipFile(Output.of(zipFile));
        }

        /**
         * @param zipHash The hash of your source zip file, changing it will re-apply function. Can be any string
         * 
         * @return builder
         * 
         */
        public Builder zipHash(@Nullable Output<String> zipHash) {
            $.zipHash = zipHash;
            return this;
        }

        /**
         * @param zipHash The hash of your source zip file, changing it will re-apply function. Can be any string
         * 
         * @return builder
         * 
         */
        public Builder zipHash(String zipHash) {
            return zipHash(Output.of(zipHash));
        }

        public FunctionArgs build() {
            $.handler = Objects.requireNonNull($.handler, "expected parameter 'handler' to be non-null");
            $.namespaceId = Objects.requireNonNull($.namespaceId, "expected parameter 'namespaceId' to be non-null");
            $.privacy = Objects.requireNonNull($.privacy, "expected parameter 'privacy' to be non-null");
            $.runtime = Objects.requireNonNull($.runtime, "expected parameter 'runtime' to be non-null");
            return $;
        }
    }

}
