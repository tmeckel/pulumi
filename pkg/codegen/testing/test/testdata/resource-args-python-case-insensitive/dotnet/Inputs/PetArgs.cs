// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Example.Inputs
{

    public sealed class PetArgs : global::Pulumi.ResourceArgs
    {
        [Input("name")]
        public Input<string?>? Name { get; set; }

        public PetArgs()
        {
        }
        public static new PetArgs Empty => new PetArgs();
    }
}
