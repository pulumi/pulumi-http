// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.http.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetHttpRetry extends com.pulumi.resources.InvokeArgs {

    public static final GetHttpRetry Empty = new GetHttpRetry();

    /**
     * The number of times the request is to be retried. For example, if 2 is specified, the request will be tried a maximum of 3 times.
     * 
     */
    @Import(name="attempts")
    private @Nullable Integer attempts;

    /**
     * @return The number of times the request is to be retried. For example, if 2 is specified, the request will be tried a maximum of 3 times.
     * 
     */
    public Optional<Integer> attempts() {
        return Optional.ofNullable(this.attempts);
    }

    /**
     * The maximum delay between retry requests in milliseconds.
     * 
     */
    @Import(name="maxDelayMs")
    private @Nullable Integer maxDelayMs;

    /**
     * @return The maximum delay between retry requests in milliseconds.
     * 
     */
    public Optional<Integer> maxDelayMs() {
        return Optional.ofNullable(this.maxDelayMs);
    }

    /**
     * The minimum delay between retry requests in milliseconds.
     * 
     */
    @Import(name="minDelayMs")
    private @Nullable Integer minDelayMs;

    /**
     * @return The minimum delay between retry requests in milliseconds.
     * 
     */
    public Optional<Integer> minDelayMs() {
        return Optional.ofNullable(this.minDelayMs);
    }

    private GetHttpRetry() {}

    private GetHttpRetry(GetHttpRetry $) {
        this.attempts = $.attempts;
        this.maxDelayMs = $.maxDelayMs;
        this.minDelayMs = $.minDelayMs;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetHttpRetry defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetHttpRetry $;

        public Builder() {
            $ = new GetHttpRetry();
        }

        public Builder(GetHttpRetry defaults) {
            $ = new GetHttpRetry(Objects.requireNonNull(defaults));
        }

        /**
         * @param attempts The number of times the request is to be retried. For example, if 2 is specified, the request will be tried a maximum of 3 times.
         * 
         * @return builder
         * 
         */
        public Builder attempts(@Nullable Integer attempts) {
            $.attempts = attempts;
            return this;
        }

        /**
         * @param maxDelayMs The maximum delay between retry requests in milliseconds.
         * 
         * @return builder
         * 
         */
        public Builder maxDelayMs(@Nullable Integer maxDelayMs) {
            $.maxDelayMs = maxDelayMs;
            return this;
        }

        /**
         * @param minDelayMs The minimum delay between retry requests in milliseconds.
         * 
         * @return builder
         * 
         */
        public Builder minDelayMs(@Nullable Integer minDelayMs) {
            $.minDelayMs = minDelayMs;
            return this;
        }

        public GetHttpRetry build() {
            return $;
        }
    }

}
