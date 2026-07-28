---
title: Create a Customer Session
page_id: operation-post-v1-customer-sessions-0ff2d1e6
path: operations/untagged
description: <p>Creates a Customer Session object that includes a single-use client secret that you can use on your front-end to grant client-side API access for certain customer resources.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/customer_sessions
operation_ids:
    - PostCustomerSessions
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Create a Customer Session

`POST /v1/customer_sessions`

Operation ID: `PostCustomerSessions`

<p>Creates a Customer Session object that includes a single-use client secret that you can use on your front-end to grant client-side API access for certain customer resources.</p>

## Definition

```yaml
{"summary": "Create a Customer Session", "description": "<p>Creates a Customer Session object that includes a single-use client secret that you can use on your front-end to grant client-side API access for certain customer resources.</p>", "operationId": "PostCustomerSessions", "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["components"], "type": "object", "properties": {"components": {"title": "components", "type": "object", "properties": {"buy_button": {"title": "buy_button_param", "required": ["enabled"], "type": "object", "properties": {"enabled": {"type": "boolean"}}}, "customer_sheet": {"title": "customer_sheet_param", "required": ["enabled"], "type": "object", "properties": {"enabled": {"type": "boolean"}, "features": {"title": "features_param", "type": "object", "properties": {"payment_method_allow_redisplay_filters": {"type": "array", "items": {"type": "string", "enum": ["always", "limited", "unspecified"]}}, "payment_method_remove": {"type": "string", "enum": ["disabled", "enabled"], "x-stripeBypassValidation": true}}}}}, "mobile_payment_element": {"title": "mobile_payment_element_param", "required": ["enabled"], "type": "object", "properties": {"enabled": {"type": "boolean"}, "features": {"title": "features_param", "type": "object", "properties": {"payment_method_allow_redisplay_filters": {"type": "array", "items": {"type": "string", "enum": ["always", "limited", "unspecified"]}}, "payment_method_redisplay": {"type": "string", "enum": ["disabled", "enabled"], "x-stripeBypassValidation": true}, "payment_method_remove": {"type": "string", "enum": ["disabled", "enabled"], "x-stripeBypassValidation": true}, "payment_method_save": {"type": "string", "enum": ["disabled", "enabled"], "x-stripeBypassValidation": true}, "payment_method_save_allow_redisplay_override": {"type": "string", "enum": ["always", "limited", "unspecified"]}}}}}, "payment_element": {"title": "payment_element_param", "required": ["enabled"], "type": "object", "properties": {"enabled": {"type": "boolean"}, "features": {"title": "features_param", "type": "object", "properties": {"payment_method_allow_redisplay_filters": {"type": "array", "items": {"type": "string", "enum": ["always", "limited", "unspecified"]}}, "payment_method_redisplay": {"type": "string", "enum": ["disabled", "enabled"], "x-stripeBypassValidation": true}, "payment_method_redisplay_limit": {"type": "integer"}, "payment_method_remove": {"type": "string", "enum": ["disabled", "enabled"], "x-stripeBypassValidation": true}, "payment_method_save": {"type": "string", "enum": ["disabled", "enabled"], "x-stripeBypassValidation": true}, "payment_method_save_usage": {"type": "string", "enum": ["off_session", "on_session"]}}}}}, "pricing_table": {"title": "pricing_table_param", "required": ["enabled"], "type": "object", "properties": {"enabled": {"type": "boolean"}}}}, "description": "Configuration for each component. At least 1 component must be enabled."}, "customer": {"maxLength": 5000, "type": "string", "description": "The ID of an existing customer for which to create the Customer Session."}, "customer_account": {"maxLength": 5000, "type": "string", "description": "The ID of an existing Account for which to create the Customer Session."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}}, "additionalProperties": false}, "encoding": {"components": {"style": "deepObject", "explode": true}, "expand": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/customer_session"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
