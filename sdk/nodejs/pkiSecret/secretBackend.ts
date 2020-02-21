// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates an PKI Secret Backend for Vault. PKI secret backends can then issue certificates, once a role has been added to
 * the backend.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const pki = new vault.pkiSecret.SecretBackend("pki", {
 *     defaultLeaseTtlSeconds: 3600,
 *     maxLeaseTtlSeconds: 86400,
 *     path: "pki",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/pki_secret_backend.html.md.
 */
export class SecretBackend extends pulumi.CustomResource {
    /**
     * Get an existing SecretBackend resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretBackendState, opts?: pulumi.CustomResourceOptions): SecretBackend {
        return new SecretBackend(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:pkiSecret/secretBackend:SecretBackend';

    /**
     * Returns true if the given object is an instance of SecretBackend.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecretBackend {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecretBackend.__pulumiType;
    }

    /**
     * The default TTL for credentials issued by this backend.
     */
    public readonly defaultLeaseTtlSeconds!: pulumi.Output<number>;
    /**
     * A human-friendly description for this backend.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The maximum TTL that can be requested for credentials issued by this backend.
     */
    public readonly maxLeaseTtlSeconds!: pulumi.Output<number>;
    /**
     * The unique path this backend should be mounted at. Must not begin or end with a `/`.
     */
    public readonly path!: pulumi.Output<string>;

    /**
     * Create a SecretBackend resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecretBackendArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretBackendArgs | SecretBackendState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SecretBackendState | undefined;
            inputs["defaultLeaseTtlSeconds"] = state ? state.defaultLeaseTtlSeconds : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["maxLeaseTtlSeconds"] = state ? state.maxLeaseTtlSeconds : undefined;
            inputs["path"] = state ? state.path : undefined;
        } else {
            const args = argsOrState as SecretBackendArgs | undefined;
            if (!args || args.path === undefined) {
                throw new Error("Missing required property 'path'");
            }
            inputs["defaultLeaseTtlSeconds"] = args ? args.defaultLeaseTtlSeconds : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["maxLeaseTtlSeconds"] = args ? args.maxLeaseTtlSeconds : undefined;
            inputs["path"] = args ? args.path : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(SecretBackend.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecretBackend resources.
 */
export interface SecretBackendState {
    /**
     * The default TTL for credentials issued by this backend.
     */
    readonly defaultLeaseTtlSeconds?: pulumi.Input<number>;
    /**
     * A human-friendly description for this backend.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The maximum TTL that can be requested for credentials issued by this backend.
     */
    readonly maxLeaseTtlSeconds?: pulumi.Input<number>;
    /**
     * The unique path this backend should be mounted at. Must not begin or end with a `/`.
     */
    readonly path?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecretBackend resource.
 */
export interface SecretBackendArgs {
    /**
     * The default TTL for credentials issued by this backend.
     */
    readonly defaultLeaseTtlSeconds?: pulumi.Input<number>;
    /**
     * A human-friendly description for this backend.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The maximum TTL that can be requested for credentials issued by this backend.
     */
    readonly maxLeaseTtlSeconds?: pulumi.Input<number>;
    /**
     * The unique path this backend should be mounted at. Must not begin or end with a `/`.
     */
    readonly path: pulumi.Input<string>;
}
