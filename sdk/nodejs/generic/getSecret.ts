// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getSecret(args: GetSecretArgs, opts?: pulumi.InvokeOptions): Promise<GetSecretResult> & GetSecretResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetSecretResult> = pulumi.runtime.invoke("vault:generic/getSecret:getSecret", {
        "path": args.path,
        "version": args.version,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getSecret.
 */
export interface GetSecretArgs {
    readonly path: string;
    readonly version?: number;
}

/**
 * A collection of values returned by getSecret.
 */
export interface GetSecretResult {
    readonly data: {[key: string]: any};
    readonly dataJson: string;
    readonly leaseDuration: number;
    readonly leaseId: string;
    readonly leaseRenewable: boolean;
    readonly leaseStartTime: string;
    readonly path: string;
    readonly version?: number;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
