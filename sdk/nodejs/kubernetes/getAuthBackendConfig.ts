// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getAuthBackendConfig(args?: GetAuthBackendConfigArgs, opts?: pulumi.InvokeOptions): Promise<GetAuthBackendConfigResult> & GetAuthBackendConfigResult {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetAuthBackendConfigResult> = pulumi.runtime.invoke("vault:kubernetes/getAuthBackendConfig:getAuthBackendConfig", {
        "backend": args.backend,
        "kubernetesCaCert": args.kubernetesCaCert,
        "kubernetesHost": args.kubernetesHost,
        "pemKeys": args.pemKeys,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getAuthBackendConfig.
 */
export interface GetAuthBackendConfigArgs {
    readonly backend?: string;
    readonly kubernetesCaCert?: string;
    readonly kubernetesHost?: string;
    readonly pemKeys?: string[];
}

/**
 * A collection of values returned by getAuthBackendConfig.
 */
export interface GetAuthBackendConfigResult {
    readonly backend?: string;
    readonly kubernetesCaCert: string;
    readonly kubernetesHost: string;
    readonly pemKeys: string[];
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
