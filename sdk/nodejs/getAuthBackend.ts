// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 * 
 * const example = vault.getAuthBackend({
 *     path: "userpass",
 * });
 * ```
 * 
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/d/auth_backend.html.markdown.
 */
export function getAuthBackend(args: GetAuthBackendArgs, opts?: pulumi.InvokeOptions): Promise<GetAuthBackendResult> & GetAuthBackendResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetAuthBackendResult> = pulumi.runtime.invoke("vault:index/getAuthBackend:getAuthBackend", {
        "path": args.path,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getAuthBackend.
 */
export interface GetAuthBackendArgs {
    /**
     * The auth backend mount point.
     */
    readonly path: string;
}

/**
 * A collection of values returned by getAuthBackend.
 */
export interface GetAuthBackendResult {
    /**
     * The accessor for this auth method
     */
    readonly accessor: string;
    /**
     * The default lease duration in seconds.
     */
    readonly defaultLeaseTtlSeconds: number;
    /**
     * A description of the auth method.
     */
    readonly description: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Speficies whether to show this mount in the UI-specific listing endpoint.
     */
    readonly listingVisibility: string;
    /**
     * Specifies if the auth method is local only.
     */
    readonly local: boolean;
    /**
     * The maximum lease duration in seconds.
     */
    readonly maxLeaseTtlSeconds: number;
    readonly path: string;
    /**
     * The name of the auth method type.
     */
    readonly type: string;
}
