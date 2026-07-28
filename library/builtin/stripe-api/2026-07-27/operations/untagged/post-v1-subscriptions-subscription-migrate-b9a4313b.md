---
title: Migrate a subscription
page_id: operation-post-v1-subscriptions-subscription-migrate-6e565bc6
path: operations/untagged
description: <p>Upgrade the billing_mode of an existing subscription.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/subscriptions/{subscription}/migrate
operation_ids:
    - PostSubscriptionsSubscriptionMigrate
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Migrate a subscription

`POST /v1/subscriptions/{subscription}/migrate`

Operation ID: `PostSubscriptionsSubscriptionMigrate`

<p>Upgrade the billing_mode of an existing subscription.</p>

## Definition

```yaml
{"summary": "Migrate a subscription", "description": "<p>Upgrade the billing_mode of an existing subscription.</p>", "operationId": "PostSubscriptionsSubscriptionMigrate", "parameters": [{"name": "subscription", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["billing_mode"], "type": "object", "properties": {"billing_mode": {"title": "billing_mode_migrate", "required": ["type"], "type": "object", "properties": {"flexible": {"title": "flexible_params", "type": "object", "properties": {"proration_discounts": {"type": "string", "enum": ["included", "itemized"]}}}, "type": {"type": "string", "enum": ["flexible"]}}, "description": "Controls how prorations and invoices for subscriptions are calculated and orchestrated."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}}, "additionalProperties": false}, "encoding": {"billing_mode": {"style": "deepObject", "explode": true}, "expand": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/subscription"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
