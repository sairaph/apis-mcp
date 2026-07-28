---
title: Retrieve a subscription
page_id: operation-get-v1-customers-customer-subscriptions-subscription-exposed-id-411e914a
path: operations/untagged
description: <p>Retrieves the subscription with the given ID.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/customers/{customer}/subscriptions/{subscription_exposed_id}
operation_ids:
    - GetCustomersCustomerSubscriptionsSubscriptionExposedId
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve a subscription

`GET /v1/customers/{customer}/subscriptions/{subscription_exposed_id}`

Operation ID: `GetCustomersCustomerSubscriptionsSubscriptionExposedId`

<p>Retrieves the subscription with the given ID.</p>

## Definition

```yaml
{"summary": "Retrieve a subscription", "description": "<p>Retrieves the subscription with the given ID.</p>", "operationId": "GetCustomersCustomerSubscriptionsSubscriptionExposedId", "parameters": [{"name": "customer", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "subscription_exposed_id", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/subscription"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
