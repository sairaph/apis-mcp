---
title: Update OAuth app
page_id: operation-put-tailnet-tailnet-oauth-apps-appid-a1617707
path: operations/oauthapps
description: |-
    Update a specific OAuth app.

    The client secret is not regenerated and is not returned in this response.

    OAuth Scope: `oauth_apps`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - PUT
api_endpoints:
    - /tailnet/{tailnet}/oauth-apps/{appId}
operation_ids:
    - updateOAuthApp
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Update OAuth app

`PUT /tailnet/{tailnet}/oauth-apps/{appId}`

Operation ID: `updateOAuthApp`

Update a specific OAuth app.

The client secret is not regenerated and is not returned in this response.

OAuth Scope: `oauth_apps`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
- $ref: '#/components/parameters/appId'
```

## Definition

```yaml
summary: Update OAuth app
description: |
    Update a specific OAuth app.

    The client secret is not regenerated and is not returned in this response.

    OAuth Scope: `oauth_apps`.
operationId: updateOAuthApp
tags:
    - OAuthApps
requestBody:
    content:
        application/json:
            schema:
                type: object
                properties:
                    name:
                        $ref: '#/components/schemas/name'
                    description:
                        $ref: '#/components/schemas/description'
                    redirectURIs:
                        $ref: '#/components/schemas/redirectURIs'
                    scopes:
                        $ref: '#/components/schemas/scopes'
                    allowedNodeAttributes:
                        $ref: '#/components/schemas/allowedNodeAttributes'
                required:
                    - name
                    - redirectURIs
                    - scopes
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
