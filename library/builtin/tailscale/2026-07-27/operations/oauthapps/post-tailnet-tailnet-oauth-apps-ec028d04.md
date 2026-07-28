---
title: Create an OAuth app
page_id: operation-post-tailnet-tailnet-oauth-apps-1aefa49a
path: operations/oauthapps
description: |-
    Create an OAuth app within a tailnet.

    The generated client secret is returned only in this response and cannot be retrieved later, so be sure to record it.

    OAuth Scope: `oauth_apps`.  If you include `allowedNodeAttributes` in the request body, you must also have `devices:posture_attributes`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - POST
api_endpoints:
    - /tailnet/{tailnet}/oauth-apps
operation_ids:
    - createOAuthApp
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Create an OAuth app

`POST /tailnet/{tailnet}/oauth-apps`

Operation ID: `createOAuthApp`

Create an OAuth app within a tailnet.

The generated client secret is returned only in this response and cannot be retrieved later, so be sure to record it.

OAuth Scope: `oauth_apps`.  If you include `allowedNodeAttributes` in the request body, you must also have `devices:posture_attributes`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
```

## Definition

```yaml
summary: Create an OAuth app
description: |
    Create an OAuth app within a tailnet.

    The generated client secret is returned only in this response and cannot be retrieved later, so be sure to record it.

    OAuth Scope: `oauth_apps`.  If you include `allowedNodeAttributes` in the request body, you must also have `devices:posture_attributes`.
operationId: createOAuthApp
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
        description: Tailnet not found.
        $ref: '#/components/responses/404'
    '500':
        $ref: '#/components/responses/500'
```
