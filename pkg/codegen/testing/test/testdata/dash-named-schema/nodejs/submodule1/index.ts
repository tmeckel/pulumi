// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { FOOEncryptedBarClassArgs } from "./fooencryptedBarClass";
export type FOOEncryptedBarClass = import("./fooencryptedBarClass").FOOEncryptedBarClass;
export const FOOEncryptedBarClass: typeof import("./fooencryptedBarClass").FOOEncryptedBarClass = null as any;
utilities.lazyLoad(exports, ["FOOEncryptedBarClass"], () => require("./fooencryptedBarClass"));

export { ModuleResourceArgs } from "./moduleResource";
export type ModuleResource = import("./moduleResource").ModuleResource;
export const ModuleResource: typeof import("./moduleResource").ModuleResource = null as any;
utilities.lazyLoad(exports, ["ModuleResource"], () => require("./moduleResource"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "foo-bar:submodule1:FOOEncryptedBarClass":
                return new FOOEncryptedBarClass(name, <any>undefined, { urn })
            case "foo-bar:submodule1:ModuleResource":
                return new ModuleResource(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("foo-bar", "submodule1", _module)
