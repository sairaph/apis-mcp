---
title: Create an auth key or trust credential
page_id: operation-post-tailnet-tailnet-keys-f770a204
path: operations/keys
description: |-
    Creates a new [auth key](https://tailscale.com/kb/1085/), or [trust credential](https://tailscale.com/kb/1623/) in the specified tailnet.
    Trust credentials include [OAuth clients](https://tailscale.com/kb/1215/) and [federated identities](https://tailscale.com/kb/1581/).
    The key will be associated with the user who owns the API access token used to make this call,
    or, if the call is made with an access token derived from an OAuth client or federated identity, the key will be owned by the tailnet.

    Returns a JSON object with the generated key.
    The key should be recorded and kept safe and secure because it wields the capabilities or scopes specified in the request.
    The identity of the key is embedded in the key itself and can be used to perform operations on the key (e.g., revoking it or retrieving information about it).
    The full key can no longer be retrieved after the initial response.

    OAuth Scope: `auth_keys` grants access to create machine auth keys.

    OAuth Scope: `oauth_keys` grants access to create OAuth clients.

    OAuth Scope: `federated_keys` grants access to created federated identities.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - POST
api_endpoints:
    - /tailnet/{tailnet}/keys
operation_ids:
    - createKey
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Create an auth key or trust credential

`POST /tailnet/{tailnet}/keys`

Operation ID: `createKey`

Creates a new [auth key](https://tailscale.com/kb/1085/), or [trust credential](https://tailscale.com/kb/1623/) in the specified tailnet.
Trust credentials include [OAuth clients](https://tailscale.com/kb/1215/) and [federated identities](https://tailscale.com/kb/1581/).
The key will be associated with the user who owns the API access token used to make this call,
or, if the call is made with an access token derived from an OAuth client or federated identity, the key will be owned by the tailnet.

Returns a JSON object with the generated key.
The key should be recorded and kept safe and secure because it wields the capabilities or scopes specified in the request.
The identity of the key is embedded in the key itself and can be used to perform operations on the key (e.g., revoking it or retrieving information about it).
The full key can no longer be retrieved after the initial response.

OAuth Scope: `auth_keys` grants access to create machine auth keys.

OAuth Scope: `oauth_keys` grants access to create OAuth clients.

OAuth Scope: `federated_keys` grants access to created federated identities.

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
```

## Definition

```yaml
summary: Create an auth key or trust credential
description: |
    Creates a new [auth key](https://tailscale.com/kb/1085/), or [trust credential](https://tailscale.com/kb/1623/) in the specified tailnet.
    Trust credentials include [OAuth clients](https://tailscale.com/kb/1215/) and [federated identities](https://tailscale.com/kb/1581/).
    The key will be associated with the user who owns the API access token used to make this call,
    or, if the call is made with an access token derived from an OAuth client or federated identity, the key will be owned by the tailnet.

    Returns a JSON object with the generated key.
    The key should be recorded and kept safe and secure because it wields the capabilities or scopes specified in the request.
    The identity of the key is embedded in the key itself and can be used to perform operations on the key (e.g., revoking it or retrieving information about it).
    The full key can no longer be retrieved after the initial response.

    OAuth Scope: `auth_keys` grants access to create machine auth keys.

    OAuth Scope: `oauth_keys` grants access to create OAuth clients.

    OAuth Scope: `federated_keys` grants access to created federated identities.
operationId: createKey
tags:
    - Keys
requestBody:
    description: |
        The supported fields vary depending on the value of the `keyType` field.

        For auth keys, at a minimum, the request POST body must have a `capabilities` object with a `devices` object,
        though it can be an empty JSON object.
        With nothing else supplied, such a request generates a single-use key with no tags.

        For OAuth clients, at a minimum the request POST body must have at least one scope.

        For federated identities, at a minimum the request POST body must have at least one scope, a valid issuer, and a subject.
    content:
        application/json:
            schema:
                type: object
                properties:
                    keyType:
                        type: string
                        enum:
                            - auth
                            - client
                            - federated
                        description: |
                            The type of key to create. Defaults to "auth" if omitted.
                    description:
                        type: string
                        description: |
                            A short string specifying the purpose of the key. Can be a maximum of 50 alphanumeric characters. Hyphens and spaces are also allowed.
                        example: dev access
                    capabilities:
                        $ref: '#/components/schemas/KeyCapabilities'
                    expirySeconds:
                        type: integer
                        format: int64
                        description: |
                            Specifies the duration in seconds until the key expires. Defaults to 90 days if not supplied.

                            Only applies to auth keys.
                        example: 86400
                    scopes:
                        type: array
                        description: |
                            A list of scopes to grant to the key. At least one scope is required for OAuth clients and federated identities.
                            See [trust credentials scopes](https://tailscale.com/kb/1623/trust-credentials#scopes) for a list of available scopes.

                            Only applies to OAuth clients and federated identities.
                        items:
                            type: string
                        example:
                            - all:read
                    tags:
                        type: array
                        description: |
                            A list of tags associated to the trust credential. Auth keys created with this credential must have these exact tags, or tags owned by the credential's tags.
                            Mandatory if the scopes include "devices:core" or "auth_keys".

                            Only applies to OAuth clients and federated identities.
                        items:
                            type: string
                        example:
                            - tag:example
                    issuer:
                        type: string
                        format: uri
                        description: |
                            The issuer of the OIDC identity token used in the token exchange. Must be a valid and publicly reachable https:// URL.

                            Only applies to federated identities.
                        example: https://example.com
                    subject:
                        type: string
                        description: |
                            The pattern used when matching against the `sub` claim from an OIDC identity token.
                            Patterns can include `*` characters to match against any character.

                            Only applies to federated identities.
                        example: my-example-subject-*
                    audience:
                        type: string
                        description: |
                            The value used when matching against the `aud` claim from an OIDC identity token.

                            Specifying the audience is optional as Tailscale will generate a secure audience at creation time by default.
                            It is recommended to let Tailscale generate the audience unless the identity provider you are integrating with
                            requires a specific audience format.

                            Only applies to federated identities.
                        example: api.tailscale.com/Tz8TefihCR11DEVEL-kqc11MVpwu11DEVEL
                    customClaimRules:
                        type: object
                        additionalProperties:
                            x-additionalPropertiesName: Custom claim rule
                            type: string
                        description: |
                            A map of claim names to pattern strings used to match against arbitrary claims in the OIDC identity token.
                            Patterns can include `*` characters to match against any character.

                            Only applies to federated identities.
                        example:
                            exampleAdditionalClaim: valueToMatch
                            otherAdditionalClaim: valueWithWildcard*
responses:
    '200':
        description: Successful operation.
        content:
            application/json:
                schema:
                    $ref: '#/components/schemas/Key'
    '404':
        description: Tailnet not found.
        $ref: '#/components/responses/404'
    '500':
        $ref: '#/components/responses/500'
```
