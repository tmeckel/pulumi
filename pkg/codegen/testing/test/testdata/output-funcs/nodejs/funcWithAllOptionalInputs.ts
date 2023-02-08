// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Check codegen of functions with all optional inputs.
 */
export function funcWithAllOptionalInputs(args?: FuncWithAllOptionalInputsArgs, opts?: pulumi.InvokeOptions): Promise<FuncWithAllOptionalInputsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("mypkg::funcWithAllOptionalInputs", {
        "a": args.a,
        "b": args.b,
    }, opts);
}

export interface FuncWithAllOptionalInputsArgs {
    /**
     * Property A
     */
    a?: string;
    /**
     * Property B
     */
    b?: string;
}

export interface FuncWithAllOptionalInputsResult {
    readonly r: string;
}
/**
 * Check codegen of functions with all optional inputs.
 */
export function funcWithAllOptionalInputsOutput(args?: FuncWithAllOptionalInputsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<FuncWithAllOptionalInputsResult> {
    return pulumi.output(args).apply((a: any) => funcWithAllOptionalInputs(a, opts))
}

export interface FuncWithAllOptionalInputsOutputArgs {
    /**
     * Property A
     */
    a?: pulumi.Input<string | undefined>;
    /**
     * Property B
     */
    b?: pulumi.Input<string | undefined>;
}
