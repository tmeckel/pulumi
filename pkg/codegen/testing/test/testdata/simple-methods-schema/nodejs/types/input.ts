// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export namespace nested {
    export interface Baz {
        hello?: string;
        world?: string;
    }

    export interface BazArgs {
        hello?: pulumi.Input<string | undefined>;
        world?: pulumi.Input<string | undefined>;
    }
}
