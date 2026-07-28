---
title: Get webhook
page_id: operation-get-webhooks-endpointid-e5b74130
path: operations/webhooks
description: |-
    Retrieve a specific webhook.

    OAuth Scope: `webhooks:read`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - GET
api_endpoints:
    - /webhooks/{endpointId}
operation_ids:
    - getWebhook
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Get webhook

`GET /webhooks/{endpointId}`

Operation ID: `getWebhook`

Retrieve a specific webhook.

OAuth Scope: `webhooks:read`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/endpointId'
```

## Definition

```yaml
summary: Get webhook
description: |
    Retrieve a specific webhook.

    OAuth Scope: `webhooks:read`.
operationId: getWebhook
tags:
    - Webhooks
responses:
    '200':
        description: Successful operation.
        content:
            application/json:
                schema:
                    $ref: '#/components/schemas/Webhook'
    '400':
        $ref: '#/components/responses/400'
    '403':
        $ref: '#/components/responses/403'
    '404':
        description: Webhook not found.
        $ref: '#/components/responses/404'
    '500':
        $ref: '#/components/responses/500'
```
