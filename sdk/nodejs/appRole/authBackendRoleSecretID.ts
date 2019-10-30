// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class AuthBackendRoleSecretID extends pulumi.CustomResource {
    /**
     * Get an existing AuthBackendRoleSecretID resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AuthBackendRoleSecretIDState, opts?: pulumi.CustomResourceOptions): AuthBackendRoleSecretID {
        return new AuthBackendRoleSecretID(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:appRole/authBackendRoleSecretID:AuthBackendRoleSecretID';

    /**
     * Returns true if the given object is an instance of AuthBackendRoleSecretID.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AuthBackendRoleSecretID {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AuthBackendRoleSecretID.__pulumiType;
    }

    /**
     * The unique ID used to access this SecretID.
     */
    public /*out*/ readonly accessor!: pulumi.Output<string>;
    /**
     * Unique name of the auth backend to configure.
     */
    public readonly backend!: pulumi.Output<string | undefined>;
    /**
     * List of CIDR blocks that can log in using the SecretID.
     */
    public readonly cidrLists!: pulumi.Output<string[] | undefined>;
    /**
     * JSON-encoded secret data to write.
     */
    public readonly metadata!: pulumi.Output<string | undefined>;
    /**
     * Name of the role.
     */
    public readonly roleName!: pulumi.Output<string>;
    /**
     * The SecretID to be managed. If not specified, Vault auto-generates one.
     */
    public readonly secretId!: pulumi.Output<string>;
    /**
     * The wrapped SecretID accessor.
     */
    public /*out*/ readonly wrappingAccessor!: pulumi.Output<string>;
    /**
     * The wrapped SecretID token.
     */
    public /*out*/ readonly wrappingToken!: pulumi.Output<string>;
    /**
     * The TTL duration of the wrapped SecretID.
     */
    public readonly wrappingTtl!: pulumi.Output<string | undefined>;

    /**
     * Create a AuthBackendRoleSecretID resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AuthBackendRoleSecretIDArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuthBackendRoleSecretIDArgs | AuthBackendRoleSecretIDState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AuthBackendRoleSecretIDState | undefined;
            inputs["accessor"] = state ? state.accessor : undefined;
            inputs["backend"] = state ? state.backend : undefined;
            inputs["cidrLists"] = state ? state.cidrLists : undefined;
            inputs["metadata"] = state ? state.metadata : undefined;
            inputs["roleName"] = state ? state.roleName : undefined;
            inputs["secretId"] = state ? state.secretId : undefined;
            inputs["wrappingAccessor"] = state ? state.wrappingAccessor : undefined;
            inputs["wrappingToken"] = state ? state.wrappingToken : undefined;
            inputs["wrappingTtl"] = state ? state.wrappingTtl : undefined;
        } else {
            const args = argsOrState as AuthBackendRoleSecretIDArgs | undefined;
            if (!args || args.roleName === undefined) {
                throw new Error("Missing required property 'roleName'");
            }
            inputs["backend"] = args ? args.backend : undefined;
            inputs["cidrLists"] = args ? args.cidrLists : undefined;
            inputs["metadata"] = args ? args.metadata : undefined;
            inputs["roleName"] = args ? args.roleName : undefined;
            inputs["secretId"] = args ? args.secretId : undefined;
            inputs["wrappingTtl"] = args ? args.wrappingTtl : undefined;
            inputs["accessor"] = undefined /*out*/;
            inputs["wrappingAccessor"] = undefined /*out*/;
            inputs["wrappingToken"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(AuthBackendRoleSecretID.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AuthBackendRoleSecretID resources.
 */
export interface AuthBackendRoleSecretIDState {
    /**
     * The unique ID used to access this SecretID.
     */
    readonly accessor?: pulumi.Input<string>;
    /**
     * Unique name of the auth backend to configure.
     */
    readonly backend?: pulumi.Input<string>;
    /**
     * List of CIDR blocks that can log in using the SecretID.
     */
    readonly cidrLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * JSON-encoded secret data to write.
     */
    readonly metadata?: pulumi.Input<string>;
    /**
     * Name of the role.
     */
    readonly roleName?: pulumi.Input<string>;
    /**
     * The SecretID to be managed. If not specified, Vault auto-generates one.
     */
    readonly secretId?: pulumi.Input<string>;
    /**
     * The wrapped SecretID accessor.
     */
    readonly wrappingAccessor?: pulumi.Input<string>;
    /**
     * The wrapped SecretID token.
     */
    readonly wrappingToken?: pulumi.Input<string>;
    /**
     * The TTL duration of the wrapped SecretID.
     */
    readonly wrappingTtl?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AuthBackendRoleSecretID resource.
 */
export interface AuthBackendRoleSecretIDArgs {
    /**
     * Unique name of the auth backend to configure.
     */
    readonly backend?: pulumi.Input<string>;
    /**
     * List of CIDR blocks that can log in using the SecretID.
     */
    readonly cidrLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * JSON-encoded secret data to write.
     */
    readonly metadata?: pulumi.Input<string>;
    /**
     * Name of the role.
     */
    readonly roleName: pulumi.Input<string>;
    /**
     * The SecretID to be managed. If not specified, Vault auto-generates one.
     */
    readonly secretId?: pulumi.Input<string>;
    /**
     * The TTL duration of the wrapped SecretID.
     */
    readonly wrappingTtl?: pulumi.Input<string>;
}
