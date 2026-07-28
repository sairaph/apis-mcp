---
title: Delete OAuth app
page_id: operation-delete-tailnet-tailnet-oauth-apps-appid-e7d4d813
path: operations/oauthapps
description: |-
    Delete a specific OAuth app.

    OAuth Scope: `oauth_apps`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - DELETE
api_endpoints:
    - /tailnet/{tailnet}/oauth-apps/{appId}
operation_ids:
    - deleteOAuthApp
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Delete OAuth app

`DELETE /tailnet/{tailnet}/oauth-apps/{appId}`

Operation ID: `deleteOAuthApp`

Delete a specific OAuth app.

OAuth Scope: `oauth_apps`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
- $ref: '#/components/parameters/appId'
```

## Definition

```yaml
summary: Delete OAuth app
description: |
    Delete a specific OAuth app.

    OAuth Scope: `oauth_apps`.
operationId: deleteOAuthApp
tags:
    - OAuthApps
responses:
    '200':
        description: Successful operation.
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
