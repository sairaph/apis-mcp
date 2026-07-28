---
title: GET /v1/customers/{customer}/subscriptions/{subscription_exposed_id}/discount
page_id: operation-get-v1-customers-customer-subscriptions-subscription-exposed-id-discount-be64feb8
path: operations/untagged
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/customers/{customer}/subscriptions/{subscription_exposed_id}/discount
operation_ids:
    - GetCustomersCustomerSubscriptionsSubscriptionExposedIdDiscount
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# GET /v1/customers/{customer}/subscriptions/{subscription_exposed_id}/discount

`GET /v1/customers/{customer}/subscriptions/{subscription_exposed_id}/discount`

Operation ID: `GetCustomersCustomerSubscriptionsSubscriptionExposedIdDiscount`

## Definition

```yaml
{"description": "", "operationId": "GetCustomersCustomerSubscriptionsSubscriptionExposedIdDiscount", "parameters": [{"name": "customer", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "subscription_exposed_id", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/discount"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
