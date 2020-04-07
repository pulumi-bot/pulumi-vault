// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * This is a data source which can be used to construct a HCL representation of an Vault policy document, for use with resources which expect policy documents, such as the `vault..Policy` resource.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 * 
 * const examplePolicyDocument = vault.getPolicyDocument({
 *     rules: [{
 *         capabilities: [
 *             "create",
 *             "read",
 *             "update",
 *             "delete",
 *             "list",
 *         ],
 *         description: "allow all on secrets",
 *         path: "secret/*",
 *     }],
 * });
 * const examplePolicy = new vault.Policy("example", {
 *     policy: examplePolicyDocument.hcl,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/d/policy_document.md.
 */
export function getPolicyDocument(args?: GetPolicyDocumentArgs, opts?: pulumi.InvokeOptions): Promise<GetPolicyDocumentResult> & GetPolicyDocumentResult {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetPolicyDocumentResult> = pulumi.runtime.invoke("vault:index/getPolicyDocument:getPolicyDocument", {
        "rules": args.rules,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getPolicyDocument.
 */
export interface GetPolicyDocumentArgs {
    readonly rules?: inputs.GetPolicyDocumentRule[];
}

/**
 * A collection of values returned by getPolicyDocument.
 */
export interface GetPolicyDocumentResult {
    /**
     * The above arguments serialized as a standard Vault HCL policy document.
     */
    readonly hcl: string;
    readonly rules: outputs.GetPolicyDocumentRule[];
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
