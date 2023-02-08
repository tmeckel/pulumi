// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class Component extends pulumi.ComponentResource {
    /** @internal */
    public static readonly __pulumiType = 'example::Component';

    /**
     * Returns true if the given object is an instance of Component.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Component {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Component.__pulumiType;
    }

    public readonly a!: pulumi.Output<boolean>;
    public readonly b!: pulumi.Output<boolean | undefined>;
    public readonly bar!: pulumi.Output<outputs.Foo | undefined>;
    public readonly baz!: pulumi.Output<outputs.Foo[] | undefined>;
    public readonly c!: pulumi.Output<number>;
    public readonly d!: pulumi.Output<number | undefined>;
    public readonly e!: pulumi.Output<string>;
    public readonly f!: pulumi.Output<string | undefined>;
    public readonly foo!: pulumi.Output<outputs.Foo | undefined>;

    /**
     * Create a Component resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ComponentArgs, opts?: pulumi.ComponentResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.a === undefined) && !opts.urn) {
                throw new Error("Missing required property 'a'");
            }
            if ((!args || args.c === undefined) && !opts.urn) {
                throw new Error("Missing required property 'c'");
            }
            if ((!args || args.e === undefined) && !opts.urn) {
                throw new Error("Missing required property 'e'");
            }
            resourceInputs["a"] = args ? args.a : undefined;
            resourceInputs["b"] = args ? args.b : undefined;
            resourceInputs["bar"] = args ? args.bar : undefined;
            resourceInputs["baz"] = args ? args.baz : undefined;
            resourceInputs["bazMap"] = args ? args.bazMap : undefined;
            resourceInputs["c"] = args ? args.c : undefined;
            resourceInputs["d"] = args ? args.d : undefined;
            resourceInputs["e"] = args ? args.e : undefined;
            resourceInputs["f"] = args ? args.f : undefined;
            resourceInputs["foo"] = args ? args.foo : undefined;
        } else {
            resourceInputs["a"] = undefined /*out*/;
            resourceInputs["b"] = undefined /*out*/;
            resourceInputs["bar"] = undefined /*out*/;
            resourceInputs["baz"] = undefined /*out*/;
            resourceInputs["c"] = undefined /*out*/;
            resourceInputs["d"] = undefined /*out*/;
            resourceInputs["e"] = undefined /*out*/;
            resourceInputs["f"] = undefined /*out*/;
            resourceInputs["foo"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Component.__pulumiType, name, resourceInputs, opts, true /*remote*/);
    }
}

/**
 * The set of arguments for constructing a Component resource.
 */
export interface ComponentArgs {
    a: boolean;
    b?: boolean;
    bar?: inputs.FooArgs;
    baz?: pulumi.Input<inputs.FooArgs>[];
    bazMap?: {[key: string]: pulumi.Input<inputs.FooArgs>};
    c: number;
    d?: number;
    e: string;
    f?: string;
    foo?: pulumi.Input<inputs.FooArgs | undefined>;
}
