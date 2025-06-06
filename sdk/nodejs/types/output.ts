// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface GetHttpRetry {
    /**
     * The number of times the request is to be retried. For example, if 2 is specified, the request will be tried a maximum of 3 times.
     */
    attempts?: number;
    /**
     * The maximum delay between retry requests in milliseconds.
     */
    maxDelayMs?: number;
    /**
     * The minimum delay between retry requests in milliseconds.
     */
    minDelayMs?: number;
}

