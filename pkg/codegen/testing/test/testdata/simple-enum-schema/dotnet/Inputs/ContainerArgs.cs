// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Plant.Inputs
{

    public sealed class ContainerArgs : global::Pulumi.ResourceArgs
    {
        [Input("brightness")]
        public Input<Pulumi.Plant.ContainerBrightness?>? Brightness { get; set; }

        [Input("color")]
        public Input<Union<Input<Pulumi.Plant.ContainerColor>, Input<string>>?>? Color { get; set; }

        [Input("material")]
        public Input<string?>? Material { get; set; }

        [Input("size", required: true)]
        public Input<Pulumi.Plant.ContainerSize> Size { get; set; } = null!;

        public ContainerArgs()
        {
            Brightness = 1;
        }
        public static new ContainerArgs Empty => new ContainerArgs();
    }
}
