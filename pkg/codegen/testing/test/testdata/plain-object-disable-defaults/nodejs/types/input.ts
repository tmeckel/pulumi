// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

import * as utilities from "../utilities";

/**
 * BETA FEATURE - Options to configure the Helm Release resource.
 */
export interface HelmReleaseSettings {
    /**
     * The backend storage driver for Helm. Values are: configmap, secret, memory, sql.
     */
    driver?: string;
    /**
     * The path to the helm plugins directory.
     */
    pluginsPath?: string;
    /**
     * to test required args
     */
    requiredArg: string;
}
/**
 * helmReleaseSettingsProvideDefaults sets the appropriate defaults for HelmReleaseSettings
 */
export function helmReleaseSettingsProvideDefaults(val: HelmReleaseSettings): HelmReleaseSettings {
    return {
        ...val,
        driver: (val.driver) ?? (utilities.getEnv("PULUMI_K8S_HELM_DRIVER") || "secret"),
        pluginsPath: (val.pluginsPath) ?? utilities.getEnv("PULUMI_K8S_HELM_PLUGINS_PATH"),
    };
}

/**
 * BETA FEATURE - Options to configure the Helm Release resource.
 */
export interface HelmReleaseSettingsArgs {
    /**
     * The backend storage driver for Helm. Values are: configmap, secret, memory, sql.
     */
    driver?: pulumi.Input<string | undefined>;
    /**
     * The path to the helm plugins directory.
     */
    pluginsPath?: pulumi.Input<string | undefined>;
    /**
     * to test required args
     */
    requiredArg: pulumi.Input<string>;
}
/**
 * helmReleaseSettingsArgsProvideDefaults sets the appropriate defaults for HelmReleaseSettingsArgs
 */
export function helmReleaseSettingsArgsProvideDefaults(val: HelmReleaseSettingsArgs): HelmReleaseSettingsArgs {
    return {
        ...val,
        driver: (val.driver) ?? (utilities.getEnv("PULUMI_K8S_HELM_DRIVER") || "secret"),
        pluginsPath: (val.pluginsPath) ?? utilities.getEnv("PULUMI_K8S_HELM_PLUGINS_PATH"),
    };
}

/**
 * Options for tuning the Kubernetes client used by a Provider.
 */
export interface KubeClientSettingsArgs {
    /**
     * Maximum burst for throttle. Default value is 10.
     */
    burst?: pulumi.Input<number | undefined>;
    /**
     * Maximum queries per second (QPS) to the API server from this client. Default value is 5.
     */
    qps?: pulumi.Input<number | undefined>;
    recTest?: pulumi.Input<inputs.KubeClientSettingsArgs | undefined>;
}
/**
 * kubeClientSettingsArgsProvideDefaults sets the appropriate defaults for KubeClientSettingsArgs
 */
export function kubeClientSettingsArgsProvideDefaults(val: KubeClientSettingsArgs): KubeClientSettingsArgs {
    return {
        ...val,
        burst: (val.burst) ?? utilities.getEnvNumber("PULUMI_K8S_CLIENT_BURST"),
        qps: (val.qps) ?? utilities.getEnvNumber("PULUMI_K8S_CLIENT_QPS"),
        recTest: (val.recTest ? pulumi.output(val.recTest).apply(inputs.kubeClientSettingsArgsProvideDefaults) : undefined),
    };
}

/**
 * Make sure that defaults propagate through types
 */
export interface LayeredTypeArgs {
    /**
     * The answer to the question
     */
    answer?: pulumi.Input<number | undefined>;
    other: pulumi.Input<inputs.HelmReleaseSettingsArgs>;
    /**
     * Test how plain types interact
     */
    plainOther?: inputs.HelmReleaseSettingsArgs;
    /**
     * The question already answered
     */
    question?: pulumi.Input<string | undefined>;
    recursive?: pulumi.Input<inputs.LayeredTypeArgs | undefined>;
    /**
     * To ask and answer
     */
    thinker: pulumi.Input<string>;
}
/**
 * layeredTypeArgsProvideDefaults sets the appropriate defaults for LayeredTypeArgs
 */
export function layeredTypeArgsProvideDefaults(val: LayeredTypeArgs): LayeredTypeArgs {
    return {
        ...val,
        answer: (val.answer) ?? 42,
        other: pulumi.output(val.other).apply(inputs.helmReleaseSettingsArgsProvideDefaults),
        plainOther: (val.plainOther ? inputs.helmReleaseSettingsArgsProvideDefaults(val.plainOther) : undefined),
        question: (val.question) ?? (utilities.getEnv("PULUMI_THE_QUESTION") || "<unknown>"),
        recursive: (val.recursive ? pulumi.output(val.recursive).apply(inputs.layeredTypeArgsProvideDefaults) : undefined),
        thinker: (val.thinker) ?? "not a good interaction",
    };
}

/**
 * A test for namespaces (mod main)
 */
export interface TypArgs {
    mod1?: pulumi.Input<inputs.mod1.TypArgs | undefined>;
    mod2?: pulumi.Input<inputs.mod2.TypArgs | undefined>;
    val?: pulumi.Input<string | undefined>;
}
/**
 * typArgsProvideDefaults sets the appropriate defaults for TypArgs
 */
export function typArgsProvideDefaults(val: TypArgs): TypArgs {
    return {
        ...val,
        mod1: (val.mod1 ? pulumi.output(val.mod1).apply(inputs.mod1.typArgsProvideDefaults) : undefined),
        mod2: (val.mod2 ? pulumi.output(val.mod2).apply(inputs.mod2.typArgsProvideDefaults) : undefined),
        val: (val.val) ?? "mod main",
    };
}
export namespace mod1 {
    /**
     * A test for namespaces (mod 1)
     */
    export interface TypArgs {
        val?: pulumi.Input<string | undefined>;
    }
    /**
     * typArgsProvideDefaults sets the appropriate defaults for TypArgs
     */
    export function typArgsProvideDefaults(val: TypArgs): TypArgs {
        return {
            ...val,
            val: (val.val) ?? "mod1",
        };
    }
}

export namespace mod2 {
    /**
     * A test for namespaces (mod 2)
     */
    export interface TypArgs {
        mod1?: pulumi.Input<inputs.mod1.TypArgs | undefined>;
        val?: pulumi.Input<string | undefined>;
    }
    /**
     * typArgsProvideDefaults sets the appropriate defaults for TypArgs
     */
    export function typArgsProvideDefaults(val: TypArgs): TypArgs {
        return {
            ...val,
            mod1: (val.mod1 ? pulumi.output(val.mod1).apply(inputs.mod1.typArgsProvideDefaults) : undefined),
            val: (val.val) ?? "mod2",
        };
    }
}
