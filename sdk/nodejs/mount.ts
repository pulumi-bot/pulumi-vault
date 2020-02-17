// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const example = new vault.Mount("example", {
 *     description: "This is an example mount",
 *     path: "dummy",
 *     type: "generic",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/mount.html.md.
 */
export class Mount extends pulumi.CustomResource {
    /**
     * Get an existing Mount resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MountState, opts?: pulumi.CustomResourceOptions): Mount {
        return new Mount(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:index/mount:Mount';

    /**
     * Returns true if the given object is an instance of Mount.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Mount {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Mount.__pulumiType;
    }

    /**
     * The accessor for this mount.
     */
    public /*out*/ readonly accessor!: pulumi.Output<string>;
    /**
     * Default lease duration for tokens and secrets in seconds
     */
    public readonly defaultLeaseTtlSeconds!: pulumi.Output<number>;
    /**
     * Human-friendly description of the mount
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Boolean flag that can be explicitly set to true to enforce local mount in HA environment
     */
    public readonly local!: pulumi.Output<boolean | undefined>;
    /**
     * Maximum possible lease duration for tokens and secrets in seconds
     */
    public readonly maxLeaseTtlSeconds!: pulumi.Output<number>;
    /**
     * Specifies mount type specific options that are passed to the backend
     */
    public readonly options!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Where the secret backend will be mounted
     */
    public readonly path!: pulumi.Output<string>;
    /**
     * Boolean flag that can be explicitly set to true to enable seal wrapping for the mount, causing values stored by the mount to be wrapped by the seal's encryption capability
     */
    public readonly sealWrap!: pulumi.Output<boolean>;
    /**
     * Type of the backend, such as "aws"
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a Mount resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MountArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MountArgs | MountState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as MountState | undefined;
            inputs["accessor"] = state ? state.accessor : undefined;
            inputs["defaultLeaseTtlSeconds"] = state ? state.defaultLeaseTtlSeconds : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["local"] = state ? state.local : undefined;
            inputs["maxLeaseTtlSeconds"] = state ? state.maxLeaseTtlSeconds : undefined;
            inputs["options"] = state ? state.options : undefined;
            inputs["path"] = state ? state.path : undefined;
            inputs["sealWrap"] = state ? state.sealWrap : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as MountArgs | undefined;
            if (!args || args.path === undefined) {
                throw new Error("Missing required property 'path'");
            }
            if (!args || args.type === undefined) {
                throw new Error("Missing required property 'type'");
            }
            inputs["defaultLeaseTtlSeconds"] = args ? args.defaultLeaseTtlSeconds : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["local"] = args ? args.local : undefined;
            inputs["maxLeaseTtlSeconds"] = args ? args.maxLeaseTtlSeconds : undefined;
            inputs["options"] = args ? args.options : undefined;
            inputs["path"] = args ? args.path : undefined;
            inputs["sealWrap"] = args ? args.sealWrap : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["accessor"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Mount.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Mount resources.
 */
export interface MountState {
    /**
     * The accessor for this mount.
     */
    readonly accessor?: pulumi.Input<string>;
    /**
     * Default lease duration for tokens and secrets in seconds
     */
    readonly defaultLeaseTtlSeconds?: pulumi.Input<number>;
    /**
     * Human-friendly description of the mount
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Boolean flag that can be explicitly set to true to enforce local mount in HA environment
     */
    readonly local?: pulumi.Input<boolean>;
    /**
     * Maximum possible lease duration for tokens and secrets in seconds
     */
    readonly maxLeaseTtlSeconds?: pulumi.Input<number>;
    /**
     * Specifies mount type specific options that are passed to the backend
     */
    readonly options?: pulumi.Input<{[key: string]: any}>;
    /**
     * Where the secret backend will be mounted
     */
    readonly path?: pulumi.Input<string>;
    /**
     * Boolean flag that can be explicitly set to true to enable seal wrapping for the mount, causing values stored by the mount to be wrapped by the seal's encryption capability
     */
    readonly sealWrap?: pulumi.Input<boolean>;
    /**
     * Type of the backend, such as "aws"
     */
    readonly type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Mount resource.
 */
export interface MountArgs {
    /**
     * Default lease duration for tokens and secrets in seconds
     */
    readonly defaultLeaseTtlSeconds?: pulumi.Input<number>;
    /**
     * Human-friendly description of the mount
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Boolean flag that can be explicitly set to true to enforce local mount in HA environment
     */
    readonly local?: pulumi.Input<boolean>;
    /**
     * Maximum possible lease duration for tokens and secrets in seconds
     */
    readonly maxLeaseTtlSeconds?: pulumi.Input<number>;
    /**
     * Specifies mount type specific options that are passed to the backend
     */
    readonly options?: pulumi.Input<{[key: string]: any}>;
    /**
     * Where the secret backend will be mounted
     */
    readonly path: pulumi.Input<string>;
    /**
     * Boolean flag that can be explicitly set to true to enable seal wrapping for the mount, causing values stored by the mount to be wrapped by the seal's encryption capability
     */
    readonly sealWrap?: pulumi.Input<boolean>;
    /**
     * Type of the backend, such as "aws"
     */
    readonly type: pulumi.Input<string>;
}
