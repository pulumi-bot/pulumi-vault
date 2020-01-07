// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Reads the Role of an Kubernetes from a Vault server. See the [Vault
 * documentation](https://www.vaultproject.io/api/auth/kubernetes/index.html#read-config) for more
 * information.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 * 
 * const config = vault.kubernetes.getAuthBackendConfig({
 *     backend: "my-kubernetes-backend",
 * });
 * 
 * export const tokenReviewerJwt = config.tokenReviewerJwt;
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/d/kubernetes_auth_backend_config.html.markdown.
 */
export function getAuthBackendConfig(args?: GetAuthBackendConfigArgs, opts?: pulumi.InvokeOptions): Promise<GetAuthBackendConfigResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("vault:kubernetes/getAuthBackendConfig:getAuthBackendConfig", {
        "backend": args.backend,
        "issuer": args.issuer,
        "kubernetesCaCert": args.kubernetesCaCert,
        "kubernetesHost": args.kubernetesHost,
        "pemKeys": args.pemKeys,
    }, opts);
}

/**
 * A collection of arguments for invoking getAuthBackendConfig.
 */
export interface GetAuthBackendConfigArgs {
    /**
     * The unique name for the Kubernetes backend the config to
     * retrieve Role attributes for resides in. Defaults to "kubernetes".
     */
    readonly backend?: string;
    readonly issuer?: string;
    readonly kubernetesCaCert?: string;
    readonly kubernetesHost?: string;
    readonly pemKeys?: string[];
}

/**
 * A collection of values returned by getAuthBackendConfig.
 */
export interface GetAuthBackendConfigResult {
    readonly backend?: string;
    /**
     * Optional JWT issuer. If no issuer is specified, `kubernetes.io/serviceaccount` will be used as the default issuer.
     */
    readonly issuer: string;
    /**
     * PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
     */
    readonly kubernetesCaCert: string;
    /**
     * Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
     */
    readonly kubernetesHost: string;
    /**
     * Optional list of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
     */
    readonly pemKeys: string[];
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
