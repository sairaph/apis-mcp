---
title: Get key
page_id: operation-get-tailnet-tailnet-keys-keyid-1ed9fc31
path: operations/keys
description: |-
    Returns a JSON object with information about a specific api access token, OAuth client, federated identity, or auth key, such as its creation and expiration dates and its capabilities.

    OAuth Scope: `api_access_tokens:read` grants access to personal API access tokens.

    OAuth Scope: `auth_keys:read` grants access to machine auth keys.

    OAuth Scope: `oauth_keys:read` grants access to OAuth clients and OAuth tokens.

    OAuth Scope: `federated_keys:read` grants access to federated identities.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - GET
api_endpoints:
    - /tailnet/{tailnet}/keys/{keyId}
operation_ids:
    - getKey
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Get key

`GET /tailnet/{tailnet}/keys/{keyId}`

Operation ID: `getKey`

Returns a JSON object with information about a specific api access token, OAuth client, federated identity, or auth key, such as its creation and expiration dates and its capabilities.

OAuth Scope: `api_access_tokens:read` grants access to personal API access tokens.

OAuth Scope: `auth_keys:read` grants access to machine auth keys.

OAuth Scope: `oauth_keys:read` grants access to OAuth clients and OAuth tokens.

OAuth Scope: `federated_keys:read` grants access to federated identities.

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
- $ref: '#/components/parameters/keyId'
```

## Definition

```yaml
summary: Get key
description: |
    Returns a JSON object with information about a specific api access token, OAuth client, federated identity, or auth key, such as its creation and expiration dates and its capabilities.

    OAuth Scope: `api_access_tokens:read` grants access to personal API access tokens.

    OAuth Scope: `auth_keys:read` grants access to machine auth keys.

    OAuth Scope: `oauth_keys:read` grants access to OAuth clients and OAuth tokens.

    OAuth Scope: `federated_keys:read` grants access to federated identities.
operationId: getKey
tags:
    - Keys
responses:
    '200':
        description: |
            Successful operation.

            Response for a revoked (deleted) or expired key will have an `invalid` field set to true.
        content:
            application/json:
                schema:
                    $ref: '#/components/schemas/Key'
    '404':
        description: Tailnet or key not found.
        $ref: '#/components/responses/404'
    '500':
        $ref: '#/components/responses/500'
```
