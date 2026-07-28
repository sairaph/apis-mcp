---
title: List webhooks
page_id: operation-get-tailnet-tailnet-webhooks-425a5e2b
path: operations/webhooks
description: |-
    List all webhooks for a tailnet.

    OAuth Scope: `webhooks:read`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - GET
api_endpoints:
    - /tailnet/{tailnet}/webhooks
operation_ids:
    - listWebhooks
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# List webhooks

`GET /tailnet/{tailnet}/webhooks`

Operation ID: `listWebhooks`

List all webhooks for a tailnet.

OAuth Scope: `webhooks:read`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
```

## Definition

```yaml
summary: List webhooks
description: |
    List all webhooks for a tailnet.

    OAuth Scope: `webhooks:read`.
operationId: listWebhooks
tags:
    - Webhooks
responses:
    '200':
        description: Successful operation.
        content:
            application/json:
                schema:
                    type: object
                    properties:
                        webhooks:
                            type: array
                            items:
                                $ref: '#/components/schemas/Webhook'
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
