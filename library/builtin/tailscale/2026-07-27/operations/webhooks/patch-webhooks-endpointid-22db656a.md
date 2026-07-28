---
title: Update webhook
page_id: operation-patch-webhooks-endpointid-99ad5015
path: operations/webhooks
description: |-
    Update a specific webhook.

    OAuth Scope: `webhooks`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - PATCH
api_endpoints:
    - /webhooks/{endpointId}
operation_ids:
    - updateWebhook
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Update webhook

`PATCH /webhooks/{endpointId}`

Operation ID: `updateWebhook`

Update a specific webhook.

OAuth Scope: `webhooks`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/endpointId'
```

## Definition

```yaml
summary: Update webhook
description: |
    Update a specific webhook.

    OAuth Scope: `webhooks`.
operationId: updateWebhook
tags:
    - Webhooks
requestBody:
    content:
        application/json:
            schema:
                type: object
                properties:
                    subscriptions:
                        $ref: '#/components/schemas/subscriptions'
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
        description: Tailnet not found.
        $ref: '#/components/responses/404'
    '500':
        $ref: '#/components/responses/500'
```
