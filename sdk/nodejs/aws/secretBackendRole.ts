// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/aws_secret_backend_role.html.md.
 */
export class SecretBackendRole extends pulumi.CustomResource {
    /**
     * Get an existing SecretBackendRole resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretBackendRoleState, opts?: pulumi.CustomResourceOptions): SecretBackendRole {
        return new SecretBackendRole(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:aws/secretBackendRole:SecretBackendRole';

    /**
     * Returns true if the given object is an instance of SecretBackendRole.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecretBackendRole {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecretBackendRole.__pulumiType;
    }

    /**
     * The path the AWS secret backend is mounted at,
     * with no leading or trailing `/`s.
     */
    public readonly backend!: pulumi.Output<string>;
    /**
     * Specifies the type of credential to be used when
     * retrieving credentials from the role. Must be one of `iamUser`, `assumedRole`, or
     * `federationToken`.
     */
    public readonly credentialType!: pulumi.Output<string>;
    /**
     * The default TTL in seconds for STS credentials.
     * When a TTL is not specified when STS credentials are requested,
     * and a default TTL is specified on the role,
     * then this default TTL will be used. Valid only when `credentialType` is one of
     * `assumedRole` or `federationToken`.
     */
    public readonly defaultStsTtl!: pulumi.Output<number>;
    /**
     * The max allowed TTL in seconds for STS credentials
     * (credentials TTL are capped to `maxStsTtl`). Valid only when `credentialType` is
     * one of `assumedRole` or `federationToken`.
     */
    public readonly maxStsTtl!: pulumi.Output<number>;
    /**
     * The name to identify this role within the backend.
     * Must be unique within the backend.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ARN for a pre-existing policy to associate
     * with this role. Either `policyDocument` or `policyArns` must be specified.
     */
    public readonly policyArns!: pulumi.Output<string[] | undefined>;
    /**
     * The JSON-formatted policy to associate with this
     * role. Either `policyDocument` or `policyArns` must be specified.
     */
    public readonly policyDocument!: pulumi.Output<string | undefined>;
    /**
     * Specifies the ARNs of the AWS roles this Vault role
     * is allowed to assume. Required when `credentialType` is `assumedRole` and
     * prohibited otherwise.
     */
    public readonly roleArns!: pulumi.Output<string[] | undefined>;

    /**
     * Create a SecretBackendRole resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecretBackendRoleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretBackendRoleArgs | SecretBackendRoleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SecretBackendRoleState | undefined;
            inputs["backend"] = state ? state.backend : undefined;
            inputs["credentialType"] = state ? state.credentialType : undefined;
            inputs["defaultStsTtl"] = state ? state.defaultStsTtl : undefined;
            inputs["maxStsTtl"] = state ? state.maxStsTtl : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["policyArns"] = state ? state.policyArns : undefined;
            inputs["policyDocument"] = state ? state.policyDocument : undefined;
            inputs["roleArns"] = state ? state.roleArns : undefined;
        } else {
            const args = argsOrState as SecretBackendRoleArgs | undefined;
            if (!args || args.backend === undefined) {
                throw new Error("Missing required property 'backend'");
            }
            if (!args || args.credentialType === undefined) {
                throw new Error("Missing required property 'credentialType'");
            }
            inputs["backend"] = args ? args.backend : undefined;
            inputs["credentialType"] = args ? args.credentialType : undefined;
            inputs["defaultStsTtl"] = args ? args.defaultStsTtl : undefined;
            inputs["maxStsTtl"] = args ? args.maxStsTtl : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["policyArns"] = args ? args.policyArns : undefined;
            inputs["policyDocument"] = args ? args.policyDocument : undefined;
            inputs["roleArns"] = args ? args.roleArns : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(SecretBackendRole.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecretBackendRole resources.
 */
export interface SecretBackendRoleState {
    /**
     * The path the AWS secret backend is mounted at,
     * with no leading or trailing `/`s.
     */
    readonly backend?: pulumi.Input<string>;
    /**
     * Specifies the type of credential to be used when
     * retrieving credentials from the role. Must be one of `iamUser`, `assumedRole`, or
     * `federationToken`.
     */
    readonly credentialType?: pulumi.Input<string>;
    /**
     * The default TTL in seconds for STS credentials.
     * When a TTL is not specified when STS credentials are requested,
     * and a default TTL is specified on the role,
     * then this default TTL will be used. Valid only when `credentialType` is one of
     * `assumedRole` or `federationToken`.
     */
    readonly defaultStsTtl?: pulumi.Input<number>;
    /**
     * The max allowed TTL in seconds for STS credentials
     * (credentials TTL are capped to `maxStsTtl`). Valid only when `credentialType` is
     * one of `assumedRole` or `federationToken`.
     */
    readonly maxStsTtl?: pulumi.Input<number>;
    /**
     * The name to identify this role within the backend.
     * Must be unique within the backend.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ARN for a pre-existing policy to associate
     * with this role. Either `policyDocument` or `policyArns` must be specified.
     */
    readonly policyArns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The JSON-formatted policy to associate with this
     * role. Either `policyDocument` or `policyArns` must be specified.
     */
    readonly policyDocument?: pulumi.Input<string>;
    /**
     * Specifies the ARNs of the AWS roles this Vault role
     * is allowed to assume. Required when `credentialType` is `assumedRole` and
     * prohibited otherwise.
     */
    readonly roleArns?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a SecretBackendRole resource.
 */
export interface SecretBackendRoleArgs {
    /**
     * The path the AWS secret backend is mounted at,
     * with no leading or trailing `/`s.
     */
    readonly backend: pulumi.Input<string>;
    /**
     * Specifies the type of credential to be used when
     * retrieving credentials from the role. Must be one of `iamUser`, `assumedRole`, or
     * `federationToken`.
     */
    readonly credentialType: pulumi.Input<string>;
    /**
     * The default TTL in seconds for STS credentials.
     * When a TTL is not specified when STS credentials are requested,
     * and a default TTL is specified on the role,
     * then this default TTL will be used. Valid only when `credentialType` is one of
     * `assumedRole` or `federationToken`.
     */
    readonly defaultStsTtl?: pulumi.Input<number>;
    /**
     * The max allowed TTL in seconds for STS credentials
     * (credentials TTL are capped to `maxStsTtl`). Valid only when `credentialType` is
     * one of `assumedRole` or `federationToken`.
     */
    readonly maxStsTtl?: pulumi.Input<number>;
    /**
     * The name to identify this role within the backend.
     * Must be unique within the backend.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ARN for a pre-existing policy to associate
     * with this role. Either `policyDocument` or `policyArns` must be specified.
     */
    readonly policyArns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The JSON-formatted policy to associate with this
     * role. Either `policyDocument` or `policyArns` must be specified.
     */
    readonly policyDocument?: pulumi.Input<string>;
    /**
     * Specifies the ARNs of the AWS roles this Vault role
     * is allowed to assume. Required when `credentialType` is `assumedRole` and
     * prohibited otherwise.
     */
    readonly roleArns?: pulumi.Input<pulumi.Input<string>[]>;
}
