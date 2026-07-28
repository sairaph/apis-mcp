---
title: Delete a customer discount
page_id: operation-delete-v1-customers-customer-subscriptions-subscription-exposed-id-disco-94414fb0
path: operations/untagged
description: <p>Removes the currently applied discount on a customer.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - DELETE
api_endpoints:
    - /v1/customers/{customer}/subscriptions/{subscription_exposed_id}/discount
operation_ids:
    - DeleteCustomersCustomerSubscriptionsSubscriptionExposedIdDiscount
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Delete a customer discount

`DELETE /v1/customers/{customer}/subscriptions/{subscription_exposed_id}/discount`

Operation ID: `DeleteCustomersCustomerSubscriptionsSubscriptionExposedIdDiscount`

<p>Removes the currently applied discount on a customer.</p>

## Definition

```yaml
{"summary": "Delete a customer discount", "description": "<p>Removes the currently applied discount on a customer.</p>", "operationId": "DeleteCustomersCustomerSubscriptionsSubscriptionExposedIdDiscount", "parameters": [{"name": "customer", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "subscription_exposed_id", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/deleted_discount"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
