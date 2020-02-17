// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a user in an
 * [Okta auth backend within Vault](https://www.vaultproject.io/docs/auth/okta.html).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const example = new vault.okta.AuthBackend("example", {
 *     organization: "dummy",
 *     path: "userOkta",
 * });
 * const foo = new vault.okta.AuthBackendUser("foo", {
 *     groups: [
 *         "one",
 *         "two",
 *     ],
 *     path: example.path,
 *     username: "foo",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/okta_auth_backend_user.html.markdown.
 */
export class AuthBackendUser extends pulumi.CustomResource {
    /**
     * Get an existing AuthBackendUser resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AuthBackendUserState, opts?: pulumi.CustomResourceOptions): AuthBackendUser {
        return new AuthBackendUser(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:okta/authBackendUser:AuthBackendUser';

    /**
     * Returns true if the given object is an instance of AuthBackendUser.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AuthBackendUser {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AuthBackendUser.__pulumiType;
    }

    /**
     * List of Okta groups to associate with this user
     */
    public readonly groups!: pulumi.Output<string[] | undefined>;
    /**
     * The path where the Okta auth backend is mounted
     */
    public readonly path!: pulumi.Output<string>;
    /**
     * List of Vault policies to associate with this user
     */
    public readonly policies!: pulumi.Output<string[] | undefined>;
    /**
     * Name of the user within Okta
     */
    public readonly username!: pulumi.Output<string>;

    /**
     * Create a AuthBackendUser resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AuthBackendUserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuthBackendUserArgs | AuthBackendUserState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AuthBackendUserState | undefined;
            inputs["groups"] = state ? state.groups : undefined;
            inputs["path"] = state ? state.path : undefined;
            inputs["policies"] = state ? state.policies : undefined;
            inputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as AuthBackendUserArgs | undefined;
            if (!args || args.path === undefined) {
                throw new Error("Missing required property 'path'");
            }
            if (!args || args.username === undefined) {
                throw new Error("Missing required property 'username'");
            }
            inputs["groups"] = args ? args.groups : undefined;
            inputs["path"] = args ? args.path : undefined;
            inputs["policies"] = args ? args.policies : undefined;
            inputs["username"] = args ? args.username : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(AuthBackendUser.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AuthBackendUser resources.
 */
export interface AuthBackendUserState {
    /**
     * List of Okta groups to associate with this user
     */
    readonly groups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The path where the Okta auth backend is mounted
     */
    readonly path?: pulumi.Input<string>;
    /**
     * List of Vault policies to associate with this user
     */
    readonly policies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the user within Okta
     */
    readonly username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AuthBackendUser resource.
 */
export interface AuthBackendUserArgs {
    /**
     * List of Okta groups to associate with this user
     */
    readonly groups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The path where the Okta auth backend is mounted
     */
    readonly path: pulumi.Input<string>;
    /**
     * List of Vault policies to associate with this user
     */
    readonly policies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the user within Okta
     */
    readonly username: pulumi.Input<string>;
}
