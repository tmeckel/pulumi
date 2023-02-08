// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

import * as pulumiRandom from "@pulumi/random";

export interface PetArgs {
    age?: pulumi.Input<number | undefined>;
    name?: pulumi.Input<pulumiRandom.RandomPet | undefined>;
    nameArray?: pulumi.Input<pulumi.Input<pulumiRandom.RandomPet>[] | undefined>;
    nameMap?: pulumi.Input<{[key: string]: pulumi.Input<pulumiRandom.RandomPet>} | undefined>;
    requiredName: pulumi.Input<pulumiRandom.RandomPet>;
    requiredNameArray: pulumi.Input<pulumi.Input<pulumiRandom.RandomPet>[]>;
    requiredNameMap: pulumi.Input<{[key: string]: pulumi.Input<pulumiRandom.RandomPet>}>;
}
