// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Http.Inputs
{

    public sealed class GetHttpRetryArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The number of times the request is to be retried. For example, if 2 is specified, the request will be tried a maximum of 3 times.
        /// </summary>
        [Input("attempts")]
        public int? Attempts { get; set; }

        /// <summary>
        /// The maximum delay between retry requests in milliseconds.
        /// </summary>
        [Input("maxDelayMs")]
        public int? MaxDelayMs { get; set; }

        /// <summary>
        /// The minimum delay between retry requests in milliseconds.
        /// </summary>
        [Input("minDelayMs")]
        public int? MinDelayMs { get; set; }

        public GetHttpRetryArgs()
        {
        }
        public static new GetHttpRetryArgs Empty => new GetHttpRetryArgs();
    }
}
