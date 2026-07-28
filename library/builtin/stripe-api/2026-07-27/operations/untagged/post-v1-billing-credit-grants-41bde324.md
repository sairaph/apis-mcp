---
title: Create a credit grant
page_id: operation-post-v1-billing-credit-grants-f618bdb3
path: operations/untagged
description: <p>Creates a credit grant.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/billing/credit_grants
operation_ids:
    - PostBillingCreditGrants
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Create a credit grant

`POST /v1/billing/credit_grants`

Operation ID: `PostBillingCreditGrants`

<p>Creates a credit grant.</p>

## Definition

```yaml
{"summary": "Create a credit grant", "description": "<p>Creates a credit grant.</p>", "operationId": "PostBillingCreditGrants", "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["amount", "applicability_config"], "type": "object", "properties": {"amount": {"title": "amount_param", "required": ["type"], "type": "object", "properties": {"monetary": {"title": "monetary_amount_param", "required": ["currency", "value"], "type": "object", "properties": {"currency": {"type": "string", "format": "currency"}, "value": {"type": "integer"}}}, "type": {"type": "string", "enum": ["monetary"], "x-stripeBypassValidation": true}}, "description": "Amount of this credit grant."}, "applicability_config": {"title": "applicability_config_param", "required": ["scope"], "type": "object", "properties": {"scope": {"title": "scope_param", "type": "object", "properties": {"price_type": {"type": "string", "enum": ["metered"]}, "prices": {"type": "array", "items": {"title": "applicable_price_param", "required": ["id"], "type": "object", "properties": {"id": {"maxLength": 5000, "type": "string"}}}}}}}, "description": "Configuration specifying what this credit grant applies to. We currently only support `metered` prices that have a [Billing Meter](https://docs.stripe.com/api/billing/meter) attached to them."}, "category": {"type": "string", "description": "The category of this credit grant. It defaults to `paid` if not specified.", "enum": ["paid", "promotional"]}, "customer": {"maxLength": 5000, "type": "string", "description": "ID of the customer receiving the billing credits."}, "customer_account": {"maxLength": 5000, "type": "string", "description": "ID of the account representing the customer receiving the billing credits."}, "effective_at": {"type": "integer", "description": "The time when the billing credits become effective-when they're eligible for use. It defaults to the current timestamp if not specified.", "format": "unix-time"}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "expires_at": {"type": "integer", "description": "The time when the billing credits expire. If not specified, the billing credits don't expire.", "format": "unix-time"}, "metadata": {"type": "object", "additionalProperties": {"type": "string"}, "description": "Set of key-value pairs that you can attach to an object. You can use this to store additional information about the object (for example, cost basis) in a structured format."}, "name": {"maxLength": 100, "type": "string", "description": "A descriptive name shown in the Dashboard."}, "priority": {"type": "integer", "description": "The desired priority for applying this credit grant. If not specified, it will be set to the default value of 50. The highest priority is 0 and the lowest is 100."}}, "additionalProperties": false}, "encoding": {"amount": {"style": "deepObject", "explode": true}, "applicability_config": {"style": "deepObject", "explode": true}, "expand": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/billing.credit_grant"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
