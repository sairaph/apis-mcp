---
title: Key
page_id: schema-key-99a52df3
path: schemas
description: An API access token or Auth Key.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Key

An API access token or Auth Key.

```yaml
type: object
description: |
    An API access token or Auth Key.
properties:
    id:
        type: string
        example: k123456CNTRL
    key:
        type: string
        description: The secret key material (only populated at creation time).
        example: tskey-auth-xxxxxxxxxxxx-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
    keyType:
        type: string
        enum:
            - auth
            - client
            - api
            - federated
        description: |
            The type of key. Can be one of "auth", "client", "federated", or "api".
            - "auth" refers to machine auth keys.
            - "client" refers to OAuth clients.
            - "federated" refers to federated identities.
            - "api" refers to personal API access tokens or tokens generated using an OAuth client or federated identity.
        example: auth
    expirySeconds:
        type: integer
        format: int64
        description: |
            Duration in seconds until the key expires.
            Only applies to auth keys.
        example: 7776000
    created:
        type: string
        format: date-time
        example: '2021-12-09T23:22:39Z'
    updated:
        type: string
        format: date-time
        example: '2021-12-09T23:22:39Z'
    expires:
        type: string
        format: date-time
        example: '2022-03-09T23:22:39Z'
    revoked:
        type: string
        format: date-time
        example: '2022-03-12T23:22:39Z'
    capabilities:
        $ref: '#/components/schemas/KeyCapabilities'
    scopes:
        type: array
        description: |
            A list of scopes granted to the key.

            Only applies to OAuth clients, API access tokens, and federated identities.
        items:
            type: string
        example:
            - all:read
    tags:
        type: array
        description: "A list of tags associated to the trust credential. \nAuth keys created with this client must have these exact tags, or tags owned by the client's tags. \nMandatory if the scopes include \"devices:core\" or \"auth_keys\". \n\nOnly applies to OAuth clients and federated identities.\n"
        items:
            type: string
        example:
            - tag:example
    description:
        type: string
        example: dev access
    invalid:
        type: boolean
        description: |
            Response for a revoked (deleted) or expired key will have an `invalid` field set to true.
        example: false
    userId:
        type: string
        description: |
            ID of the user who created this key, empty for keys created by trust credentials.
        example: uscwcTtzzo11DEVEL
    audience:
        type: string
        description: |
            The value used when matching against the `aud` claim from an OIDC identity token.

            Specifying the audience is optional as Tailscale will generate a secure audience at creation time by default.
            It is recommended to let Tailscale generate the audience unless the identity provider you are integrating with
            requires a specific audience format.

            Only applies to federated identities.
        example: api.tailscale.com/Tz8TefihCR11DEVEL-kqc11MVpwu11DEVEL
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
            This pattern can include `*` characters to match against any character.

            Only applies to federated identities.
        example: my-example-subject-*
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
```
