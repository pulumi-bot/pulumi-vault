// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Creates a Roleset in the [GCP Secrets Engine](https://www.vaultproject.io/docs/secrets/gcp/index.html) for Vault.
 * 
 * Each Roleset is [tied](https://www.vaultproject.io/docs/secrets/gcp/index.html#service-accounts-are-tied-to-rolesets) to a Service Account, and can have one or more [bindings](https://www.vaultproject.io/docs/secrets/gcp/index.html#roleset-bindings) associated with it.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fs from "fs";
 * import * as vault from "@pulumi/vault";
 * 
 * const project = "my-awesome-project";
 * const gcp = new vault.gcp.SecretBackend("gcp", {
 *     credentials: fs.readFileSync("credentials.json", "utf-8"),
 *     path: "gcp",
 * });
 * const roleset = new vault.gcp.SecretRoleset("roleset", {
 *     backend: gcp.path,
 *     bindings: [{
 *         resource: `//cloudresourcemanager.googleapis.com/projects/${project}`,
 *         roles: ["roles/viewer"],
 *     }],
 *     project: project,
 *     roleset: "projectViewer",
 *     secretType: "accessToken",
 *     tokenScopes: ["https://www.googleapis.com/auth/cloud-platform"],
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/gcp_secret_roleset.html.markdown.
 */
export class SecretRoleset extends pulumi.CustomResource {
    /**
     * Get an existing SecretRoleset resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretRolesetState, opts?: pulumi.CustomResourceOptions): SecretRoleset {
        return new SecretRoleset(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:gcp/secretRoleset:SecretRoleset';

    /**
     * Returns true if the given object is an instance of SecretRoleset.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecretRoleset {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecretRoleset.__pulumiType;
    }

    /**
     * Path where the GCP Secrets Engine is mounted
     */
    public readonly backend!: pulumi.Output<string>;
    /**
     * Bindings to create for this roleset. This can be specified multiple times for multiple bindings. Structure is documented below.
     */
    public readonly bindings!: pulumi.Output<outputs.gcp.SecretRolesetBinding[]>;
    /**
     * Name of the GCP project that this roleset's service account will belong to.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Name of the Roleset to create
     */
    public readonly roleset!: pulumi.Output<string>;
    /**
     * Type of secret generated for this role set. Accepted values: `accessToken`, `serviceAccountKey`. Defaults to `accessToken`.
     */
    public readonly secretType!: pulumi.Output<string>;
    /**
     * Email of the service account created by Vault for this Roleset
     */
    public /*out*/ readonly serviceAccountEmail!: pulumi.Output<string>;
    /**
     * List of OAuth scopes to assign to `accessToken` secrets generated under this role set (`accessToken` role sets only).
     */
    public readonly tokenScopes!: pulumi.Output<string[] | undefined>;

    /**
     * Create a SecretRoleset resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecretRolesetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretRolesetArgs | SecretRolesetState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SecretRolesetState | undefined;
            inputs["backend"] = state ? state.backend : undefined;
            inputs["bindings"] = state ? state.bindings : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["roleset"] = state ? state.roleset : undefined;
            inputs["secretType"] = state ? state.secretType : undefined;
            inputs["serviceAccountEmail"] = state ? state.serviceAccountEmail : undefined;
            inputs["tokenScopes"] = state ? state.tokenScopes : undefined;
        } else {
            const args = argsOrState as SecretRolesetArgs | undefined;
            if (!args || args.backend === undefined) {
                throw new Error("Missing required property 'backend'");
            }
            if (!args || args.bindings === undefined) {
                throw new Error("Missing required property 'bindings'");
            }
            if (!args || args.project === undefined) {
                throw new Error("Missing required property 'project'");
            }
            if (!args || args.roleset === undefined) {
                throw new Error("Missing required property 'roleset'");
            }
            inputs["backend"] = args ? args.backend : undefined;
            inputs["bindings"] = args ? args.bindings : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["roleset"] = args ? args.roleset : undefined;
            inputs["secretType"] = args ? args.secretType : undefined;
            inputs["tokenScopes"] = args ? args.tokenScopes : undefined;
            inputs["serviceAccountEmail"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(SecretRoleset.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecretRoleset resources.
 */
export interface SecretRolesetState {
    /**
     * Path where the GCP Secrets Engine is mounted
     */
    readonly backend?: pulumi.Input<string>;
    /**
     * Bindings to create for this roleset. This can be specified multiple times for multiple bindings. Structure is documented below.
     */
    readonly bindings?: pulumi.Input<pulumi.Input<inputs.gcp.SecretRolesetBinding>[]>;
    /**
     * Name of the GCP project that this roleset's service account will belong to.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Name of the Roleset to create
     */
    readonly roleset?: pulumi.Input<string>;
    /**
     * Type of secret generated for this role set. Accepted values: `accessToken`, `serviceAccountKey`. Defaults to `accessToken`.
     */
    readonly secretType?: pulumi.Input<string>;
    /**
     * Email of the service account created by Vault for this Roleset
     */
    readonly serviceAccountEmail?: pulumi.Input<string>;
    /**
     * List of OAuth scopes to assign to `accessToken` secrets generated under this role set (`accessToken` role sets only).
     */
    readonly tokenScopes?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a SecretRoleset resource.
 */
export interface SecretRolesetArgs {
    /**
     * Path where the GCP Secrets Engine is mounted
     */
    readonly backend: pulumi.Input<string>;
    /**
     * Bindings to create for this roleset. This can be specified multiple times for multiple bindings. Structure is documented below.
     */
    readonly bindings: pulumi.Input<pulumi.Input<inputs.gcp.SecretRolesetBinding>[]>;
    /**
     * Name of the GCP project that this roleset's service account will belong to.
     */
    readonly project: pulumi.Input<string>;
    /**
     * Name of the Roleset to create
     */
    readonly roleset: pulumi.Input<string>;
    /**
     * Type of secret generated for this role set. Accepted values: `accessToken`, `serviceAccountKey`. Defaults to `accessToken`.
     */
    readonly secretType?: pulumi.Input<string>;
    /**
     * List of OAuth scopes to assign to `accessToken` secrets generated under this role set (`accessToken` role sets only).
     */
    readonly tokenScopes?: pulumi.Input<pulumi.Input<string>[]>;
}
