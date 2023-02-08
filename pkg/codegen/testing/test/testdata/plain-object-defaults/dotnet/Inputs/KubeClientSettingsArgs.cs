// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Example.Inputs
{

    /// <summary>
    /// Options for tuning the Kubernetes client used by a Provider.
    /// </summary>
    public sealed class KubeClientSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Maximum burst for throttle. Default value is 10.
        /// </summary>
        [Input("burst")]
        public Input<int?>? Burst { get; set; }

        /// <summary>
        /// Maximum queries per second (QPS) to the API server from this client. Default value is 5.
        /// </summary>
        [Input("qps")]
        public Input<double?>? Qps { get; set; }

        [Input("recTest")]
        public Input<Inputs.KubeClientSettingsArgs?>? RecTest { get; set; }

        public KubeClientSettingsArgs()
        {
            Burst = Utilities.GetEnv("PULUMI_K8S_CLIENT_BURST");
            Qps = Utilities.GetEnv("PULUMI_K8S_CLIENT_QPS");
        }
        public static new KubeClientSettingsArgs Empty => new KubeClientSettingsArgs();
    }
}
