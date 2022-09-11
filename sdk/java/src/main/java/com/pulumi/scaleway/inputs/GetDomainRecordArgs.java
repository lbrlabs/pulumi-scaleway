// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetDomainRecordArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetDomainRecordArgs Empty = new GetDomainRecordArgs();

    /**
     * The content of the record (an IPv4 for an `A`, a string for a `TXT`...).
     * Cannot be used with `record_id`.
     * 
     */
    @Import(name="data")
    private @Nullable Output<String> data;

    /**
     * @return The content of the record (an IPv4 for an `A`, a string for a `TXT`...).
     * Cannot be used with `record_id`.
     * 
     */
    public Optional<Output<String>> data() {
        return Optional.ofNullable(this.data);
    }

    /**
     * The IP address.
     * 
     */
    @Import(name="dnsZone")
    private @Nullable Output<String> dnsZone;

    /**
     * @return The IP address.
     * 
     */
    public Optional<Output<String>> dnsZone() {
        return Optional.ofNullable(this.dnsZone);
    }

    /**
     * The name of the record (can be an empty string for a root record).
     * Cannot be used with `record_id`.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the record (can be an empty string for a root record).
     * Cannot be used with `record_id`.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The record ID.
     * Cannot be used with `name`, `type` and `data`.
     * 
     */
    @Import(name="recordId")
    private @Nullable Output<String> recordId;

    /**
     * @return The record ID.
     * Cannot be used with `name`, `type` and `data`.
     * 
     */
    public Optional<Output<String>> recordId() {
        return Optional.ofNullable(this.recordId);
    }

    /**
     * The type of the record (`A`, `AAAA`, `MX`, `CNAME`, `DNAME`, `ALIAS`, `NS`, `PTR`, `SRV`, `TXT`, `TLSA`, or `CAA`).
     * Cannot be used with `record_id`.
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return The type of the record (`A`, `AAAA`, `MX`, `CNAME`, `DNAME`, `ALIAS`, `NS`, `PTR`, `SRV`, `TXT`, `TLSA`, or `CAA`).
     * Cannot be used with `record_id`.
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    private GetDomainRecordArgs() {}

    private GetDomainRecordArgs(GetDomainRecordArgs $) {
        this.data = $.data;
        this.dnsZone = $.dnsZone;
        this.name = $.name;
        this.recordId = $.recordId;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetDomainRecordArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetDomainRecordArgs $;

        public Builder() {
            $ = new GetDomainRecordArgs();
        }

        public Builder(GetDomainRecordArgs defaults) {
            $ = new GetDomainRecordArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param data The content of the record (an IPv4 for an `A`, a string for a `TXT`...).
         * Cannot be used with `record_id`.
         * 
         * @return builder
         * 
         */
        public Builder data(@Nullable Output<String> data) {
            $.data = data;
            return this;
        }

        /**
         * @param data The content of the record (an IPv4 for an `A`, a string for a `TXT`...).
         * Cannot be used with `record_id`.
         * 
         * @return builder
         * 
         */
        public Builder data(String data) {
            return data(Output.of(data));
        }

        /**
         * @param dnsZone The IP address.
         * 
         * @return builder
         * 
         */
        public Builder dnsZone(@Nullable Output<String> dnsZone) {
            $.dnsZone = dnsZone;
            return this;
        }

        /**
         * @param dnsZone The IP address.
         * 
         * @return builder
         * 
         */
        public Builder dnsZone(String dnsZone) {
            return dnsZone(Output.of(dnsZone));
        }

        /**
         * @param name The name of the record (can be an empty string for a root record).
         * Cannot be used with `record_id`.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the record (can be an empty string for a root record).
         * Cannot be used with `record_id`.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param recordId The record ID.
         * Cannot be used with `name`, `type` and `data`.
         * 
         * @return builder
         * 
         */
        public Builder recordId(@Nullable Output<String> recordId) {
            $.recordId = recordId;
            return this;
        }

        /**
         * @param recordId The record ID.
         * Cannot be used with `name`, `type` and `data`.
         * 
         * @return builder
         * 
         */
        public Builder recordId(String recordId) {
            return recordId(Output.of(recordId));
        }

        /**
         * @param type The type of the record (`A`, `AAAA`, `MX`, `CNAME`, `DNAME`, `ALIAS`, `NS`, `PTR`, `SRV`, `TXT`, `TLSA`, or `CAA`).
         * Cannot be used with `record_id`.
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type The type of the record (`A`, `AAAA`, `MX`, `CNAME`, `DNAME`, `ALIAS`, `NS`, `PTR`, `SRV`, `TXT`, `TLSA`, or `CAA`).
         * Cannot be used with `record_id`.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public GetDomainRecordArgs build() {
            return $;
        }
    }

}