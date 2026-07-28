---
title: Delete a subscription discount
page_id: operation-delete-v1-subscriptions-subscription-exposed-id-discount-b5609098
path: operations/untagged
description: <p>Removes the currently applied discount on a subscription.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - DELETE
api_endpoints:
    - /v1/subscriptions/{subscription_exposed_id}/discount
operation_ids:
    - DeleteSubscriptionsSubscriptionExposedIdDiscount
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Delete a subscription discount

`DELETE /v1/subscriptions/{subscription_exposed_id}/discount`

Operation ID: `DeleteSubscriptionsSubscriptionExposedIdDiscount`

<p>Removes the currently applied discount on a subscription.</p>

## Definition

```yaml
{"summary": "Delete a subscription discount", "description": "<p>Removes the currently applied discount on a subscription.</p>", "operationId": "DeleteSubscriptionsSubscriptionExposedIdDiscount", "parameters": [{"name": "subscription_exposed_id", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/deleted_discount"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
