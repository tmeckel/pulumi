// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Example
{
    [ExampleResourceType("example::Component")]
    public partial class Component : global::Pulumi.ComponentResource
    {
        [Output("a")]
        public Output<bool> A { get; private set; } = null!;

        [Output("b")]
        public Output<bool?> B { get; private set; } = null!;

        [Output("bar")]
        public Output<Outputs.Foo?> Bar { get; private set; } = null!;

        [Output("baz")]
        public Output<ImmutableArray<Outputs.Foo>> Baz { get; private set; } = null!;

        [Output("c")]
        public Output<int> C { get; private set; } = null!;

        [Output("d")]
        public Output<int?> D { get; private set; } = null!;

        [Output("e")]
        public Output<string> E { get; private set; } = null!;

        [Output("f")]
        public Output<string?> F { get; private set; } = null!;

        [Output("foo")]
        public Output<Outputs.Foo?> Foo { get; private set; } = null!;


        /// <summary>
        /// Create a Component resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Component(string name, ComponentArgs args, ComponentResourceOptions? options = null)
            : base("example::Component", name, args ?? new ComponentArgs(), MakeResourceOptions(options, ""), remote: true)
        {
        }

        private static ComponentResourceOptions MakeResourceOptions(ComponentResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new ComponentResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = ComponentResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
    }

    public sealed class ComponentArgs : global::Pulumi.ResourceArgs
    {
        [Input("a", required: true)]
        public bool A { get; set; }

        [Input("b")]
        public bool? B { get; set; }

        [Input("bar")]
        public Inputs.FooArgs? Bar { get; set; }

        [Input("baz")]
        private List<Input<Inputs.FooArgs>>? _baz;
        public List<Input<Inputs.FooArgs>> Baz
        {
            get => _baz ?? (_baz = new List<Input<Inputs.FooArgs>>());
            set => _baz = value;
        }

        [Input("bazMap")]
        private Dictionary<string, Input<Inputs.FooArgs>>? _bazMap;
        public Dictionary<string, Input<Inputs.FooArgs>> BazMap
        {
            get => _bazMap ?? (_bazMap = new Dictionary<string, Input<Inputs.FooArgs>>());
            set => _bazMap = value;
        }

        [Input("c", required: true)]
        public int C { get; set; }

        [Input("d")]
        public int? D { get; set; }

        [Input("e", required: true)]
        public string E { get; set; } = null!;

        [Input("f")]
        public string? F { get; set; }

        [Input("foo")]
        public Input<Inputs.FooArgs?>? Foo { get; set; }

        public ComponentArgs()
        {
        }
        public static new ComponentArgs Empty => new ComponentArgs();
    }
}
