---
title: Test a webhook
page_id: operation-post-webhooks-endpointid-test-0fe3693a
path: operations/webhooks
description: |-
    Test a specific webhook by sending out a test event to the endpoint URL.
    This endpoint queues the event which is sent out asynchronously.

    If your webhook is configured correctly, within a few seconds your webhook endpoint should receive an event with type of "test".

    OAuth Scope: `webhooks`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - POST
api_endpoints:
    - /webhooks/{endpointId}/test
operation_ids:
    - testWebhook
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Test a webhook

`POST /webhooks/{endpointId}/test`

Operation ID: `testWebhook`

Test a specific webhook by sending out a test event to the endpoint URL.
This endpoint queues the event which is sent out asynchronously.

If your webhook is configured correctly, within a few seconds your webhook endpoint should receive an event with type of "test".

OAuth Scope: `webhooks`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/endpointId'
```

## Definition

```yaml
summary: Test a webhook
description: |
    Test a specific webhook by sending out a test event to the endpoint URL.
    This endpoint queues the event which is sent out asynchronously.

    If your webhook is configured correctly, within a few seconds your webhook endpoint should receive an event with type of "test".

    OAuth Scope: `webhooks`.
operationId: testWebhook
tags:
    - Webhooks
responses:
    '202':
        description: Successfully queued test event.
    '400':
        $ref: '#/components/responses/400'
    '403':
        $ref: '#/components/responses/403'
    '404':
        description: User not found.
        $ref: '#/components/responses/404'
    '500':
        $ref: '#/components/responses/500'
```
