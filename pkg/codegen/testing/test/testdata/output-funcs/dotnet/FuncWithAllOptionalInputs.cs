// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Mypkg
{
    public static class FuncWithAllOptionalInputs
    {
        /// <summary>
        /// Check codegen of functions with all optional inputs.
        /// </summary>
        public static Task<FuncWithAllOptionalInputsResult> InvokeAsync(FuncWithAllOptionalInputsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<FuncWithAllOptionalInputsResult>("mypkg::funcWithAllOptionalInputs", args ?? new FuncWithAllOptionalInputsArgs(), options.WithDefaults());

        /// <summary>
        /// Check codegen of functions with all optional inputs.
        /// </summary>
        public static Output<FuncWithAllOptionalInputsResult> Invoke(FuncWithAllOptionalInputsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<FuncWithAllOptionalInputsResult>("mypkg::funcWithAllOptionalInputs", args ?? new FuncWithAllOptionalInputsInvokeArgs(), options.WithDefaults());
    }


    public sealed class FuncWithAllOptionalInputsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Property A
        /// </summary>
        [Input("a")]
        public string? A { get; set; }

        /// <summary>
        /// Property B
        /// </summary>
        [Input("b")]
        public string? B { get; set; }

        public FuncWithAllOptionalInputsArgs()
        {
        }
        public static new FuncWithAllOptionalInputsArgs Empty => new FuncWithAllOptionalInputsArgs();
    }

    public sealed class FuncWithAllOptionalInputsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Property A
        /// </summary>
        [Input("a")]
        public Input<string?>? A { get; set; }

        /// <summary>
        /// Property B
        /// </summary>
        [Input("b")]
        public Input<string?>? B { get; set; }

        public FuncWithAllOptionalInputsInvokeArgs()
        {
        }
        public static new FuncWithAllOptionalInputsInvokeArgs Empty => new FuncWithAllOptionalInputsInvokeArgs();
    }


    [OutputType]
    public sealed class FuncWithAllOptionalInputsResult
    {
        public readonly string R;

        [OutputConstructor]
        private FuncWithAllOptionalInputsResult(string r)
        {
            R = r;
        }
    }
}
