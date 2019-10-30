// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class AuthBackendIdentityWhitelist extends pulumi.CustomResource {
    /**
     * Get an existing AuthBackendIdentityWhitelist resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AuthBackendIdentityWhitelistState, opts?: pulumi.CustomResourceOptions): AuthBackendIdentityWhitelist {
        return new AuthBackendIdentityWhitelist(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:aws/authBackendIdentityWhitelist:AuthBackendIdentityWhitelist';

    /**
     * Returns true if the given object is an instance of AuthBackendIdentityWhitelist.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AuthBackendIdentityWhitelist {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AuthBackendIdentityWhitelist.__pulumiType;
    }

    /**
     * Unique name of the auth backend to configure.
     */
    public readonly backend!: pulumi.Output<string | undefined>;
    /**
     * If true, disables the periodic tidying of the identiy whitelist entries.
     */
    public readonly disablePeriodicTidy!: pulumi.Output<boolean | undefined>;
    /**
     * The amount of extra time that must have passed beyond the roletag expiration, before it's removed from backend
     * storage.
     */
    public readonly safetyBuffer!: pulumi.Output<number | undefined>;

    /**
     * Create a AuthBackendIdentityWhitelist resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AuthBackendIdentityWhitelistArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuthBackendIdentityWhitelistArgs | AuthBackendIdentityWhitelistState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AuthBackendIdentityWhitelistState | undefined;
            inputs["backend"] = state ? state.backend : undefined;
            inputs["disablePeriodicTidy"] = state ? state.disablePeriodicTidy : undefined;
            inputs["safetyBuffer"] = state ? state.safetyBuffer : undefined;
        } else {
            const args = argsOrState as AuthBackendIdentityWhitelistArgs | undefined;
            inputs["backend"] = args ? args.backend : undefined;
            inputs["disablePeriodicTidy"] = args ? args.disablePeriodicTidy : undefined;
            inputs["safetyBuffer"] = args ? args.safetyBuffer : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(AuthBackendIdentityWhitelist.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AuthBackendIdentityWhitelist resources.
 */
export interface AuthBackendIdentityWhitelistState {
    /**
     * Unique name of the auth backend to configure.
     */
    readonly backend?: pulumi.Input<string>;
    /**
     * If true, disables the periodic tidying of the identiy whitelist entries.
     */
    readonly disablePeriodicTidy?: pulumi.Input<boolean>;
    /**
     * The amount of extra time that must have passed beyond the roletag expiration, before it's removed from backend
     * storage.
     */
    readonly safetyBuffer?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a AuthBackendIdentityWhitelist resource.
 */
export interface AuthBackendIdentityWhitelistArgs {
    /**
     * Unique name of the auth backend to configure.
     */
    readonly backend?: pulumi.Input<string>;
    /**
     * If true, disables the periodic tidying of the identiy whitelist entries.
     */
    readonly disablePeriodicTidy?: pulumi.Input<boolean>;
    /**
     * The amount of extra time that must have passed beyond the roletag expiration, before it's removed from backend
     * storage.
     */
    readonly safetyBuffer?: pulumi.Input<number>;
}
