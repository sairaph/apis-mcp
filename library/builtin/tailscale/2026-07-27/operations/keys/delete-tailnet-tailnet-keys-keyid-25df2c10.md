---
title: Delete key
page_id: operation-delete-tailnet-tailnet-keys-keyid-96966c71
path: operations/keys
description: |-
    Deletes a specific api access token or auth key.

    OAuth Scope: `api_access_tokens` grants access to personal API access tokens.

    OAuth Scope: `auth_keys` grants access to machine auth keys.

    OAuth Scope: `oauth_keys` grants access to OAuth clients and OAuth tokens.

    OAuth Scope: `federated_keys` grants access to federated identities.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - DELETE
api_endpoints:
    - /tailnet/{tailnet}/keys/{keyId}
operation_ids:
    - deleteKey
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Delete key

`DELETE /tailnet/{tailnet}/keys/{keyId}`

Operation ID: `deleteKey`

Deletes a specific api access token or auth key.

OAuth Scope: `api_access_tokens` grants access to personal API access tokens.

OAuth Scope: `auth_keys` grants access to machine auth keys.

OAuth Scope: `oauth_keys` grants access to OAuth clients and OAuth tokens.

OAuth Scope: `federated_keys` grants access to federated identities.

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
- $ref: '#/components/parameters/keyId'
```

## Definition

```yaml
summary: Delete key
description: |
    Deletes a specific api access token or auth key.

    OAuth Scope: `api_access_tokens` grants access to personal API access tokens.

    OAuth Scope: `auth_keys` grants access to machine auth keys.

    OAuth Scope: `oauth_keys` grants access to OAuth clients and OAuth tokens.

    OAuth Scope: `federated_keys` grants access to federated identities.
operationId: deleteKey
tags:
    - Keys
responses:
    '200':
        description: Successful operation.
    '403':
        description: User does not have sufficient access to delete this key.
        $ref: '#/components/responses/403'
    '404':
        description: Tailnet not found.
        $ref: '#/components/responses/404'
    '500':
        $ref: '#/components/responses/500'
```
