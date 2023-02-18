// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ObjectBucketLifecycleRuleTransitionArgs extends com.pulumi.resources.ResourceArgs {

    public static final ObjectBucketLifecycleRuleTransitionArgs Empty = new ObjectBucketLifecycleRuleTransitionArgs();

    /**
     * Specifies the number of days after object creation when the specific rule action takes effect.
     * 
     */
    @Import(name="days")
    private @Nullable Output<Integer> days;

    /**
     * @return Specifies the number of days after object creation when the specific rule action takes effect.
     * 
     */
    public Optional<Output<Integer>> days() {
        return Optional.ofNullable(this.days);
    }

    /**
     * Specifies the Scaleway [storage class](https://www.scaleway.com/en/docs/storage/object/concepts/#storage-class) `STANDARD`, `GLACIER`, `ONEZONE_IA`  to which you want the object to transition.
     * 
     */
    @Import(name="storageClass", required=true)
    private Output<String> storageClass;

    /**
     * @return Specifies the Scaleway [storage class](https://www.scaleway.com/en/docs/storage/object/concepts/#storage-class) `STANDARD`, `GLACIER`, `ONEZONE_IA`  to which you want the object to transition.
     * 
     */
    public Output<String> storageClass() {
        return this.storageClass;
    }

    private ObjectBucketLifecycleRuleTransitionArgs() {}

    private ObjectBucketLifecycleRuleTransitionArgs(ObjectBucketLifecycleRuleTransitionArgs $) {
        this.days = $.days;
        this.storageClass = $.storageClass;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ObjectBucketLifecycleRuleTransitionArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ObjectBucketLifecycleRuleTransitionArgs $;

        public Builder() {
            $ = new ObjectBucketLifecycleRuleTransitionArgs();
        }

        public Builder(ObjectBucketLifecycleRuleTransitionArgs defaults) {
            $ = new ObjectBucketLifecycleRuleTransitionArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param days Specifies the number of days after object creation when the specific rule action takes effect.
         * 
         * @return builder
         * 
         */
        public Builder days(@Nullable Output<Integer> days) {
            $.days = days;
            return this;
        }

        /**
         * @param days Specifies the number of days after object creation when the specific rule action takes effect.
         * 
         * @return builder
         * 
         */
        public Builder days(Integer days) {
            return days(Output.of(days));
        }

        /**
         * @param storageClass Specifies the Scaleway [storage class](https://www.scaleway.com/en/docs/storage/object/concepts/#storage-class) `STANDARD`, `GLACIER`, `ONEZONE_IA`  to which you want the object to transition.
         * 
         * @return builder
         * 
         */
        public Builder storageClass(Output<String> storageClass) {
            $.storageClass = storageClass;
            return this;
        }

        /**
         * @param storageClass Specifies the Scaleway [storage class](https://www.scaleway.com/en/docs/storage/object/concepts/#storage-class) `STANDARD`, `GLACIER`, `ONEZONE_IA`  to which you want the object to transition.
         * 
         * @return builder
         * 
         */
        public Builder storageClass(String storageClass) {
            return storageClass(Output.of(storageClass));
        }

        public ObjectBucketLifecycleRuleTransitionArgs build() {
            $.storageClass = Objects.requireNonNull($.storageClass, "expected parameter 'storageClass' to be non-null");
            return $;
        }
    }

}
