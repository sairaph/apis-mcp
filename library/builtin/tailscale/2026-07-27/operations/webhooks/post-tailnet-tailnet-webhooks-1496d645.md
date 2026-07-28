---
title: Create a webhook
page_id: operation-post-tailnet-tailnet-webhooks-649402a9
path: operations/webhooks
description: |-
    Create a webhook within a tailnet.

    OAuth Scope: `webhooks`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - POST
api_endpoints:
    - /tailnet/{tailnet}/webhooks
operation_ids:
    - createWebhook
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Create a webhook

`POST /tailnet/{tailnet}/webhooks`

Operation ID: `createWebhook`

Create a webhook within a tailnet.

OAuth Scope: `webhooks`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
```

## Definition

```yaml
summary: Create a webhook
description: |
    Create a webhook within a tailnet.

    OAuth Scope: `webhooks`.
operationId: createWebhook
tags:
    - Webhooks
requestBody:
    content:
        application/json:
            schema:
                type: object
                properties:
                    endpointUrl:
                        type: string
                        example: https://example.com/endpoint
                        description: |
                            The endpoint that events are sent to from Tailscale via POST requests.
                    providerType:
                        $ref: '#/components/schemas/providerType'
                    subscriptions:
                        $ref: '#/components/schemas/subscriptions'
                required:
                    - endpointUrl
                    - subscriptions
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
