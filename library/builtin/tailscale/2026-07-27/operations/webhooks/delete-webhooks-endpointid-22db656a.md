---
title: Delete webhook
page_id: operation-delete-webhooks-endpointid-33fe4f12
path: operations/webhooks
description: |-
    Delete a specific webhook.

    OAuth Scope: `webhooks`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - DELETE
api_endpoints:
    - /webhooks/{endpointId}
operation_ids:
    - deleteWebhook
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Delete webhook

`DELETE /webhooks/{endpointId}`

Operation ID: `deleteWebhook`

Delete a specific webhook.

OAuth Scope: `webhooks`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/endpointId'
```

## Definition

```yaml
summary: Delete webhook
description: |
    Delete a specific webhook.

    OAuth Scope: `webhooks`.
operationId: deleteWebhook
tags:
    - Webhooks
responses:
    '200':
        description: Successful operation.
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
