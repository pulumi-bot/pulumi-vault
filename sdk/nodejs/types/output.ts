// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as outputs from "../types/output";

export interface AuthBackendTune {
    /**
     * List of headers to whitelist and allowing
     * a plugin to include them in the response.
     */
    allowedResponseHeaders?: string[];
    /**
     * Specifies the list of keys that will
     * not be HMAC'd by audit devices in the request data object.
     */
    auditNonHmacRequestKeys?: string[];
    /**
     * Specifies the list of keys that will
     * not be HMAC'd by audit devices in the response data object.
     */
    auditNonHmacResponseKeys?: string[];
    /**
     * Specifies the default time-to-live.
     * If set, this overrides the global default.
     * Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)
     */
    defaultLeaseTtl?: string;
    /**
     * Specifies whether to show this mount in
     * the UI-specific listing endpoint. Valid values are "unauth" or "hidden".
     */
    listingVisibility?: string;
    /**
     * Specifies the maximum time-to-live.
     * If set, this overrides the global default.
     * Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)
     */
    maxLeaseTtl?: string;
    /**
     * List of headers to whitelist and
     * pass from the request to the backend.
     */
    passthroughRequestHeaders?: string[];
    /**
     * Specifies the type of tokens that should be returned by
     * the mount. Valid values are "default-service", "default-batch", "service", "batch".
     */
    tokenType?: string;
}

export interface GetPolicyDocumentRule {
    /**
     * Whitelists a list of keys and values that are permitted on the given path. See Parameters below.
     */
    allowedParameters?: outputs.GetPolicyDocumentRuleAllowedParameter[];
    /**
     * A list of capabilities that this rule apply to `path`. For example, ["read", "write"].
     */
    capabilities: string[];
    /**
     * Blacklists a list of parameter and values. Any values specified here take precedence over `allowedParameter`. See Parameters below.
     */
    deniedParameters?: outputs.GetPolicyDocumentRuleDeniedParameter[];
    /**
     * Description of the rule. Will be added as a commend to rendered rule.
     */
    description?: string;
    /**
     * The maximum allowed TTL that clients can specify for a wrapped response.
     */
    maxWrappingTtl?: string;
    /**
     * The minimum allowed TTL that clients can specify for a wrapped response.
     */
    minWrappingTtl?: string;
    /**
     * A path in Vault that this rule applies to.
     */
    path: string;
    /**
     * A list of parameters that must be specified.
     */
    requiredParameters?: string[];
}

export interface GetPolicyDocumentRuleAllowedParameter {
    /**
     * name of permitted or denied parameter.
     */
    key: string;
    /**
     * list of values what are permitted or denied by policy rule.
     */
    values: string[];
}

export interface GetPolicyDocumentRuleDeniedParameter {
    /**
     * name of permitted or denied parameter.
     */
    key: string;
    /**
     * list of values what are permitted or denied by policy rule.
     */
    values: string[];
}

export namespace azure {
    export interface BackendRoleAzureRole {
        roleId: string;
        roleName: string;
        scope: string;
    }
}

export namespace database {
    export interface SecretBackendConnectionCassandra {
        /**
         * The number of seconds to use as a connection
         * timeout.
         */
        connectTimeout?: number;
        /**
         * The hosts to connect to.
         */
        hosts?: string[];
        /**
         * Whether to skip verification of the server
         * certificate when using TLS.
         */
        insecureTls?: boolean;
        /**
         * The password to authenticate with.
         */
        password?: string;
        /**
         * Concatenated PEM blocks configuring the certificate
         * chain.
         */
        pemBundle?: string;
        /**
         * A JSON structure configuring the certificate chain.
         */
        pemJson?: string;
        /**
         * The default port to connect to if no port is specified as
         * part of the host.
         */
        port?: number;
        /**
         * The CQL protocol version to use.
         */
        protocolVersion?: number;
        /**
         * Whether to use TLS when connecting to Cassandra.
         */
        tls?: boolean;
        /**
         * The username to authenticate with.
         */
        username?: string;
    }

    export interface SecretBackendConnectionElasticsearch {
        /**
         * The password to authenticate with.
         */
        password: string;
        /**
         * The URL for Elasticsearch's API. https requires certificate
         * by trusted CA if used.
         */
        url: string;
        /**
         * The username to authenticate with.
         */
        username: string;
    }

    export interface SecretBackendConnectionHana {
        /**
         * A URL containing connection information. See
         * the [Vault
         * docs](https://www.vaultproject.io/api/secret/databases/mongodb.html#sample-payload)
         * for an example.
         */
        connectionUrl?: string;
        /**
         * The maximum number of seconds to keep
         * a connection alive for.
         */
        maxConnectionLifetime?: number;
        /**
         * The maximum number of idle connections to
         * maintain.
         */
        maxIdleConnections?: number;
        /**
         * The maximum number of open connections to
         * use.
         */
        maxOpenConnections?: number;
    }

    export interface SecretBackendConnectionMongodb {
        /**
         * A URL containing connection information. See
         * the [Vault
         * docs](https://www.vaultproject.io/api/secret/databases/mongodb.html#sample-payload)
         * for an example.
         */
        connectionUrl?: string;
        /**
         * The maximum number of seconds to keep
         * a connection alive for.
         */
        maxConnectionLifetime?: number;
        /**
         * The maximum number of idle connections to
         * maintain.
         */
        maxIdleConnections?: number;
        /**
         * The maximum number of open connections to
         * use.
         */
        maxOpenConnections?: number;
    }

    export interface SecretBackendConnectionMssql {
        /**
         * A URL containing connection information. See
         * the [Vault
         * docs](https://www.vaultproject.io/api/secret/databases/mongodb.html#sample-payload)
         * for an example.
         */
        connectionUrl?: string;
        /**
         * The maximum number of seconds to keep
         * a connection alive for.
         */
        maxConnectionLifetime?: number;
        /**
         * The maximum number of idle connections to
         * maintain.
         */
        maxIdleConnections?: number;
        /**
         * The maximum number of open connections to
         * use.
         */
        maxOpenConnections?: number;
    }

    export interface SecretBackendConnectionMysql {
        /**
         * A URL containing connection information. See
         * the [Vault
         * docs](https://www.vaultproject.io/api/secret/databases/mongodb.html#sample-payload)
         * for an example.
         */
        connectionUrl?: string;
        /**
         * The maximum number of seconds to keep
         * a connection alive for.
         */
        maxConnectionLifetime?: number;
        /**
         * The maximum number of idle connections to
         * maintain.
         */
        maxIdleConnections?: number;
        /**
         * The maximum number of open connections to
         * use.
         */
        maxOpenConnections?: number;
    }

    export interface SecretBackendConnectionMysqlAurora {
        /**
         * A URL containing connection information. See
         * the [Vault
         * docs](https://www.vaultproject.io/api/secret/databases/mongodb.html#sample-payload)
         * for an example.
         */
        connectionUrl?: string;
        /**
         * The maximum number of seconds to keep
         * a connection alive for.
         */
        maxConnectionLifetime?: number;
        /**
         * The maximum number of idle connections to
         * maintain.
         */
        maxIdleConnections?: number;
        /**
         * The maximum number of open connections to
         * use.
         */
        maxOpenConnections?: number;
    }

    export interface SecretBackendConnectionMysqlLegacy {
        /**
         * A URL containing connection information. See
         * the [Vault
         * docs](https://www.vaultproject.io/api/secret/databases/mongodb.html#sample-payload)
         * for an example.
         */
        connectionUrl?: string;
        /**
         * The maximum number of seconds to keep
         * a connection alive for.
         */
        maxConnectionLifetime?: number;
        /**
         * The maximum number of idle connections to
         * maintain.
         */
        maxIdleConnections?: number;
        /**
         * The maximum number of open connections to
         * use.
         */
        maxOpenConnections?: number;
    }

    export interface SecretBackendConnectionMysqlRds {
        /**
         * A URL containing connection information. See
         * the [Vault
         * docs](https://www.vaultproject.io/api/secret/databases/mongodb.html#sample-payload)
         * for an example.
         */
        connectionUrl?: string;
        /**
         * The maximum number of seconds to keep
         * a connection alive for.
         */
        maxConnectionLifetime?: number;
        /**
         * The maximum number of idle connections to
         * maintain.
         */
        maxIdleConnections?: number;
        /**
         * The maximum number of open connections to
         * use.
         */
        maxOpenConnections?: number;
    }

    export interface SecretBackendConnectionOracle {
        /**
         * A URL containing connection information. See
         * the [Vault
         * docs](https://www.vaultproject.io/api/secret/databases/mongodb.html#sample-payload)
         * for an example.
         */
        connectionUrl?: string;
        /**
         * The maximum number of seconds to keep
         * a connection alive for.
         */
        maxConnectionLifetime?: number;
        /**
         * The maximum number of idle connections to
         * maintain.
         */
        maxIdleConnections?: number;
        /**
         * The maximum number of open connections to
         * use.
         */
        maxOpenConnections?: number;
    }

    export interface SecretBackendConnectionPostgresql {
        /**
         * A URL containing connection information. See
         * the [Vault
         * docs](https://www.vaultproject.io/api/secret/databases/mongodb.html#sample-payload)
         * for an example.
         */
        connectionUrl?: string;
        /**
         * The maximum number of seconds to keep
         * a connection alive for.
         */
        maxConnectionLifetime?: number;
        /**
         * The maximum number of idle connections to
         * maintain.
         */
        maxIdleConnections?: number;
        /**
         * The maximum number of open connections to
         * use.
         */
        maxOpenConnections?: number;
    }
}

export namespace gcp {
    export interface SecretRolesetBinding {
        /**
         * Resource or resource path for which IAM policy information will be bound. The resource path may be specified in a few different [formats](https://www.vaultproject.io/docs/secrets/gcp/index.html#roleset-bindings).
         */
        resource: string;
        /**
         * List of [GCP IAM roles](https://cloud.google.com/iam/docs/understanding-roles) for the resource.
         */
        roles: string[];
    }
}

export namespace github {
    export interface AuthBackendTune {
        /**
         * List of headers to whitelist and allowing
         * a plugin to include them in the response.
         */
        allowedResponseHeaders?: string[];
        /**
         * Specifies the list of keys that will
         * not be HMAC'd by audit devices in the request data object.
         */
        auditNonHmacRequestKeys?: string[];
        /**
         * Specifies the list of keys that will
         * not be HMAC'd by audit devices in the response data object.
         */
        auditNonHmacResponseKeys?: string[];
        /**
         * Specifies the default time-to-live.
         * If set, this overrides the global default.
         * Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)
         */
        defaultLeaseTtl?: string;
        /**
         * Specifies whether to show this mount in
         * the UI-specific listing endpoint. Valid values are "unauth" or "hidden".
         */
        listingVisibility?: string;
        /**
         * Specifies the maximum time-to-live.
         * If set, this overrides the global default.
         * Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)
         */
        maxLeaseTtl?: string;
        /**
         * List of headers to whitelist and
         * pass from the request to the backend.
         */
        passthroughRequestHeaders?: string[];
        /**
         * Specifies the type of tokens that should be returned by
         * the mount. Valid values are "default-service", "default-batch", "service", "batch".
         */
        tokenType?: string;
    }
}

export namespace identity {
    export interface GetEntityAlias {
        /**
         * Canonical ID of the Alias
         */
        canonicalId: string;
        /**
         * Creation time of the Alias
         */
        creationTime: string;
        /**
         * ID of the alias
         */
        id: string;
        /**
         * Last update time of the alias
         */
        lastUpdateTime: string;
        /**
         * List of canonical IDs merged with this alias
         */
        mergedFromCanonicalIds: string[];
        /**
         * Arbitrary metadata
         */
        metadata: {[key: string]: any};
        /**
         * Authentication mount acccessor which this alias belongs to
         */
        mountAccessor: string;
        /**
         * Authentication mount path which this alias belongs to
         */
        mountPath: string;
        /**
         * Authentication mount type which this alias belongs to
         */
        mountType: string;
        /**
         * Name of the alias
         */
        name: string;
    }
}

export namespace jwt {
    export interface AuthBackendTune {
        /**
         * List of headers to whitelist and allowing
         * a plugin to include them in the response.
         */
        allowedResponseHeaders?: string[];
        /**
         * Specifies the list of keys that will
         * not be HMAC'd by audit devices in the request data object.
         */
        auditNonHmacRequestKeys?: string[];
        /**
         * Specifies the list of keys that will
         * not be HMAC'd by audit devices in the response data object.
         */
        auditNonHmacResponseKeys?: string[];
        /**
         * Specifies the default time-to-live.
         * If set, this overrides the global default.
         * Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)
         */
        defaultLeaseTtl?: string;
        /**
         * Specifies whether to show this mount in
         * the UI-specific listing endpoint. Valid values are "unauth" or "hidden".
         */
        listingVisibility?: string;
        /**
         * Specifies the maximum time-to-live.
         * If set, this overrides the global default.
         * Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)
         */
        maxLeaseTtl?: string;
        /**
         * List of headers to whitelist and
         * pass from the request to the backend.
         */
        passthroughRequestHeaders?: string[];
        /**
         * Specifies the type of tokens that should be returned by
         * the mount. Valid values are "default-service", "default-batch", "service", "batch".
         */
        tokenType?: string;
    }
}

export namespace okta {
    export interface AuthBackendGroup {
        /**
         * Name of the group within the Okta
         */
        groupName: string;
        /**
         * Vault policies to associate with this group
         */
        policies: string[];
    }

    export interface AuthBackendUser {
        /**
         * List of Okta groups to associate with this user
         */
        groups: string[];
        /**
         * Vault policies to associate with this group
         */
        policies?: string[];
        /**
         * Name of the user within Okta
         */
        username: string;
    }
}

export namespace rabbitMq {
    export interface SecretBackendRoleVhost {
        configure: string;
        host: string;
        read: string;
        write: string;
    }
}
