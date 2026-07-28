---
title: Set key
page_id: operation-put-tailnet-tailnet-keys-keyid-f91da9a4
path: operations/keys
description: |-
    Set the configuration for an existing OAuth client or federated identity.

    OAuth Scope: `oauth_keys` grants access to OAuth clients.

    OAuth Scope: `federated_keys` grants access to federated identities.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - PUT
api_endpoints:
    - /tailnet/{tailnet}/keys/{keyId}
operation_ids:
    - setKey
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Set key

`PUT /tailnet/{tailnet}/keys/{keyId}`

Operation ID: `setKey`

Set the configuration for an existing OAuth client or federated identity.

OAuth Scope: `oauth_keys` grants access to OAuth clients.

OAuth Scope: `federated_keys` grants access to federated identities.

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
- $ref: '#/components/parameters/keyId'
```

## Definition

```yaml
summary: Set key
description: |
    Set the configuration for an existing OAuth client or federated identity.

    OAuth Scope: `oauth_keys` grants access to OAuth clients.

    OAuth Scope: `federated_keys` grants access to federated identities.
operationId: setKey
tags:
    - Keys
requestBody:
    description: |
        The supported fields vary depending on the value of the `keyType` field.
    content:
        application/json:
            schema:
                type: object
                properties:
                    keyType:
                        type: string
                        enum:
                            - client
                            - federated
                        description: |
                            The type of the key being updated.
                    description:
                        type: string
                        description: |
                            A short string specifying the purpose of the key. Can be a maximum of 50 alphanumeric characters. Hyphens and spaces are also allowed.
                        example: dev access
                    scopes:
                        type: array
                        description: |
                            A list of scopes to grant to the key. At least one scope is required.
                            See [trust credentials scopes](https://tailscale.com/kb/1623/trust-credentials#scopes) for a list of available scopes.
                        items:
                            type: string
                        example:
                            - all:read
                    tags:
                        type: array
                        description: |
                            A list of tags associated to the trust credential. Auth keys created with this credential must have these exact tags, or tags owned by the credential's tags.
                            Mandatory if the scopes include "devices:core" or "auth_keys".
                        items:
                            type: string
                        example:
                            - tag:example
                    issuer:
                        type: string
                        format: uri
                        description: |
                            The issuer of the OIDC identity token used in the token exchange. Must be a valid https:// URL.

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
