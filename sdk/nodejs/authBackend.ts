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
 * const example = new vault.AuthBackend("example", {
 *     type: "github",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/auth_backend.html.markdown.
 */
export class AuthBackend extends pulumi.CustomResource {
    /**
     * Get an existing AuthBackend resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AuthBackendState, opts?: pulumi.CustomResourceOptions): AuthBackend {
        return new AuthBackend(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:index/authBackend:AuthBackend';

    /**
     * Returns true if the given object is an instance of AuthBackend.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AuthBackend {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AuthBackend.__pulumiType;
    }

    /**
     * The accessor for this auth method
     */
    public /*out*/ readonly accessor!: pulumi.Output<string>;
    /**
     * The default lease duration in seconds.
     */
    public readonly defaultLeaseTtlSeconds!: pulumi.Output<number>;
    /**
     * A description of the auth method
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Speficies whether to show this mount in the UI-specific listing endpoint.
     */
    public readonly listingVisibility!: pulumi.Output<string>;
    /**
     * Specifies if the auth method is local only.
     */
    public readonly local!: pulumi.Output<boolean | undefined>;
    /**
     * The maximum lease duration in seconds.
     */
    public readonly maxLeaseTtlSeconds!: pulumi.Output<number>;
    /**
     * The path to mount the auth method — this defaults to the name of the type
     */
    public readonly path!: pulumi.Output<string>;
    public readonly tune!: pulumi.Output<outputs.AuthBackendTune>;
    /**
     * The name of the auth method type
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a AuthBackend resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AuthBackendArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuthBackendArgs | AuthBackendState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AuthBackendState | undefined;
            inputs["accessor"] = state ? state.accessor : undefined;
            inputs["defaultLeaseTtlSeconds"] = state ? state.defaultLeaseTtlSeconds : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["listingVisibility"] = state ? state.listingVisibility : undefined;
            inputs["local"] = state ? state.local : undefined;
            inputs["maxLeaseTtlSeconds"] = state ? state.maxLeaseTtlSeconds : undefined;
            inputs["path"] = state ? state.path : undefined;
            inputs["tune"] = state ? state.tune : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as AuthBackendArgs | undefined;
            if (!args || args.type === undefined) {
                throw new Error("Missing required property 'type'");
            }
            inputs["defaultLeaseTtlSeconds"] = args ? args.defaultLeaseTtlSeconds : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["listingVisibility"] = args ? args.listingVisibility : undefined;
            inputs["local"] = args ? args.local : undefined;
            inputs["maxLeaseTtlSeconds"] = args ? args.maxLeaseTtlSeconds : undefined;
            inputs["path"] = args ? args.path : undefined;
            inputs["tune"] = args ? args.tune : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["accessor"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(AuthBackend.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AuthBackend resources.
 */
export interface AuthBackendState {
    /**
     * The accessor for this auth method
     */
    readonly accessor?: pulumi.Input<string>;
    /**
     * The default lease duration in seconds.
     */
    readonly defaultLeaseTtlSeconds?: pulumi.Input<number>;
    /**
     * A description of the auth method
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Speficies whether to show this mount in the UI-specific listing endpoint.
     */
    readonly listingVisibility?: pulumi.Input<string>;
    /**
     * Specifies if the auth method is local only.
     */
    readonly local?: pulumi.Input<boolean>;
    /**
     * The maximum lease duration in seconds.
     */
    readonly maxLeaseTtlSeconds?: pulumi.Input<number>;
    /**
     * The path to mount the auth method — this defaults to the name of the type
     */
    readonly path?: pulumi.Input<string>;
    readonly tune?: pulumi.Input<inputs.AuthBackendTune>;
    /**
     * The name of the auth method type
     */
    readonly type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AuthBackend resource.
 */
export interface AuthBackendArgs {
    /**
     * The default lease duration in seconds.
     */
    readonly defaultLeaseTtlSeconds?: pulumi.Input<number>;
    /**
     * A description of the auth method
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Speficies whether to show this mount in the UI-specific listing endpoint.
     */
    readonly listingVisibility?: pulumi.Input<string>;
    /**
     * Specifies if the auth method is local only.
     */
    readonly local?: pulumi.Input<boolean>;
    /**
     * The maximum lease duration in seconds.
     */
    readonly maxLeaseTtlSeconds?: pulumi.Input<number>;
    /**
     * The path to mount the auth method — this defaults to the name of the type
     */
    readonly path?: pulumi.Input<string>;
    readonly tune?: pulumi.Input<inputs.AuthBackendTune>;
    /**
     * The name of the auth method type
     */
    readonly type: pulumi.Input<string>;
}
