// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { ComponentArgs } from "./component";
export type Component = import("./component").Component;
export const Component: typeof import("./component").Component = null as any;
utilities.lazyLoad(exports, ["Component"], () => require("./component"));

export { Component2Args } from "./component2";
export type Component2 = import("./component2").Component2;
export const Component2: typeof import("./component2").Component2 = null as any;
utilities.lazyLoad(exports, ["Component2"], () => require("./component2"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "example:mod:Component":
                return new Component(name, <any>undefined, { urn })
            case "example:mod:Component2":
                return new Component2(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("example", "mod", _module)
