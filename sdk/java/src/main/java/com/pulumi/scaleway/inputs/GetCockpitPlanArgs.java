// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetCockpitPlanArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetCockpitPlanArgs Empty = new GetCockpitPlanArgs();

    /**
     * The name of the plan.
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return The name of the plan.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    private GetCockpitPlanArgs() {}

    private GetCockpitPlanArgs(GetCockpitPlanArgs $) {
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetCockpitPlanArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetCockpitPlanArgs $;

        public Builder() {
            $ = new GetCockpitPlanArgs();
        }

        public Builder(GetCockpitPlanArgs defaults) {
            $ = new GetCockpitPlanArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name The name of the plan.
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the plan.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public GetCockpitPlanArgs build() {
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            return $;
        }
    }

}
