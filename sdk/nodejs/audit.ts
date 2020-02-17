// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage (file audit device)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const test = new vault.Audit("test", {
 *     options: {
 *         file_path: "C:/temp/audit.txt",
 *     },
 *     type: "file",
 * });
 * ```
 *
 * ## Example Usage (socket audit device)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const test = new vault.Audit("test", {
 *     options: {
 *         address: "127.0.0.1:8000",
 *         description: "application x socket",
 *         socket_type: "tcp",
 *     },
 *     path: "appSocket",
 *     type: "socket",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/audit.html.markdown.
 */
export class Audit extends pulumi.CustomResource {
    /**
     * Get an existing Audit resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AuditState, opts?: pulumi.CustomResourceOptions): Audit {
        return new Audit(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:index/audit:Audit';

    /**
     * Returns true if the given object is an instance of Audit.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Audit {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Audit.__pulumiType;
    }

    /**
     * Human-friendly description of the audit device.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Configuration options to pass to the audit device itself.
     */
    public readonly options!: pulumi.Output<{[key: string]: string}>;
    /**
     * The path to mount the audit device. This defaults to the type.
     */
    public readonly path!: pulumi.Output<string>;
    /**
     * Type of the audit device, such as 'file'.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a Audit resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AuditArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuditArgs | AuditState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AuditState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["options"] = state ? state.options : undefined;
            inputs["path"] = state ? state.path : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as AuditArgs | undefined;
            if (!args || args.options === undefined) {
                throw new Error("Missing required property 'options'");
            }
            if (!args || args.type === undefined) {
                throw new Error("Missing required property 'type'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["options"] = args ? args.options : undefined;
            inputs["path"] = args ? args.path : undefined;
            inputs["type"] = args ? args.type : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Audit.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Audit resources.
 */
export interface AuditState {
    /**
     * Human-friendly description of the audit device.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Configuration options to pass to the audit device itself.
     */
    readonly options?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The path to mount the audit device. This defaults to the type.
     */
    readonly path?: pulumi.Input<string>;
    /**
     * Type of the audit device, such as 'file'.
     */
    readonly type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Audit resource.
 */
export interface AuditArgs {
    /**
     * Human-friendly description of the audit device.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Configuration options to pass to the audit device itself.
     */
    readonly options: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The path to mount the audit device. This defaults to the type.
     */
    readonly path?: pulumi.Input<string>;
    /**
     * Type of the audit device, such as 'file'.
     */
    readonly type: pulumi.Input<string>;
}
