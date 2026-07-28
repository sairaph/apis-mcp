---
title: Get OAuth app
page_id: operation-get-tailnet-tailnet-oauth-apps-appid-1bacf708
path: operations/oauthapps
description: |-
    Retrieve a specific OAuth app.

    OAuth Scope: `oauth_apps:read`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - GET
api_endpoints:
    - /tailnet/{tailnet}/oauth-apps/{appId}
operation_ids:
    - getOAuthApp
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Get OAuth app

`GET /tailnet/{tailnet}/oauth-apps/{appId}`

Operation ID: `getOAuthApp`

Retrieve a specific OAuth app.

OAuth Scope: `oauth_apps:read`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
- $ref: '#/components/parameters/appId'
```

## Definition

```yaml
summary: Get OAuth app
description: |
    Retrieve a specific OAuth app.

    OAuth Scope: `oauth_apps:read`.
operationId: getOAuthApp
tags:
    - OAuthApps
responses:
    '200':
        description: Successful operation.
        content:
            application/json:
                schema:
                    $ref: '#/components/schemas/OAuthApp'
    '400':
        $ref: '#/components/responses/400'
    '403':
        $ref: '#/components/responses/403'
    '404':
        description: OAuth app not found.
        $ref: '#/components/responses/404'
    '500':
        $ref: '#/components/responses/500'
```
