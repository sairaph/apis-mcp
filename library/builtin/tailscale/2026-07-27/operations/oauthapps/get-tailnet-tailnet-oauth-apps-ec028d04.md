---
title: List OAuth apps
page_id: operation-get-tailnet-tailnet-oauth-apps-249350c8
path: operations/oauthapps
description: |-
    List all OAuth apps for a tailnet.

    OAuth Scope: `oauth_apps:read`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - GET
api_endpoints:
    - /tailnet/{tailnet}/oauth-apps
operation_ids:
    - listOAuthApps
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# List OAuth apps

`GET /tailnet/{tailnet}/oauth-apps`

Operation ID: `listOAuthApps`

List all OAuth apps for a tailnet.

OAuth Scope: `oauth_apps:read`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
```

## Definition

```yaml
summary: List OAuth apps
description: |
    List all OAuth apps for a tailnet.

    OAuth Scope: `oauth_apps:read`.
operationId: listOAuthApps
tags:
    - OAuthApps
responses:
    '200':
        description: Successful operation.
        content:
            application/json:
                schema:
                    type: object
                    properties:
                        oauthApps:
                            type: array
                            items:
                                $ref: '#/components/schemas/OAuthApp'
    '400':
        $ref: '#/components/responses/400'
    '403':
        $ref: '#/components/responses/403'
    '404':
        description: Tailnet not found.
        $ref: '#/components/responses/404'
    '500':
        $ref: '#/components/responses/500'
```
