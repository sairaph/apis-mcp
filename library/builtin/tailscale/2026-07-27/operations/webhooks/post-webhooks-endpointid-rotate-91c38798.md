---
title: Rotate webhook secret
page_id: operation-post-webhooks-endpointid-rotate-7116d9f1
path: operations/webhooks
description: |-
    Rotate and generate a new secret for a specific webhook.

    This secret is used for generating the `Tailscale-Webhook-Signature` header in requests sent to the endpoint URL.
    Learn more about [verifying webhook event signatures](/kb/1213/webhooks#verifying-an-event-signature).

    OAuth Scope: `webhooks`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - POST
api_endpoints:
    - /webhooks/{endpointId}/rotate
operation_ids:
    - rotateWebhookSecret
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Rotate webhook secret

`POST /webhooks/{endpointId}/rotate`

Operation ID: `rotateWebhookSecret`

Rotate and generate a new secret for a specific webhook.

This secret is used for generating the `Tailscale-Webhook-Signature` header in requests sent to the endpoint URL.
Learn more about [verifying webhook event signatures](/kb/1213/webhooks#verifying-an-event-signature).

OAuth Scope: `webhooks`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/endpointId'
```

## Definition

```yaml
summary: Rotate webhook secret
description: |
    Rotate and generate a new secret for a specific webhook.

    This secret is used for generating the `Tailscale-Webhook-Signature` header in requests sent to the endpoint URL.
    Learn more about [verifying webhook event signatures](/kb/1213/webhooks#verifying-an-event-signature).

    OAuth Scope: `webhooks`.
operationId: rotateWebhookSecret
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
