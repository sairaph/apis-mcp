---
title: Update balance settings
page_id: operation-post-v1-balance-settings-48e4e384
path: operations/untagged
description: |-
    <p>Updates balance settings for a given connected account.
     Related guide: <a href="/connect/authentication">Making API calls for connected accounts</a></p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/balance_settings
operation_ids:
    - PostBalanceSettings
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Update balance settings

`POST /v1/balance_settings`

Operation ID: `PostBalanceSettings`

<p>Updates balance settings for a given connected account.
 Related guide: <a href="/connect/authentication">Making API calls for connected accounts</a></p>

## Definition

```yaml
{"summary": "Update balance settings", "description": "<p>Updates balance settings for a given connected account.\n Related guide: <a href=\"/connect/authentication\">Making API calls for connected accounts</a></p>", "operationId": "PostBalanceSettings", "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "payments": {"title": "payments", "type": "object", "properties": {"debit_negative_balances": {"type": "boolean"}, "payouts": {"title": "payouts", "type": "object", "properties": {"automatic_transfer_rules_by_currency": {"anyOf": [{"type": "object", "additionalProperties": {"anyOf": [{"type": "array", "items": {"title": "automatic_transfer_rule", "required": ["payout_method", "type"], "type": "object", "properties": {"payout_method": {"type": "string"}, "transfer_up_to_amount": {"type": "integer"}, "type": {"type": "string", "enum": ["transfer_all", "transfer_up_to_amount"]}}}}, {"type": "string", "enum": [""]}]}}, {"type": "string", "enum": [""]}]}, "minimum_balance_by_currency": {"anyOf": [{"type": "object", "additionalProperties": {"anyOf": [{"type": "integer"}, {"type": "string", "enum": [""]}]}}, {"type": "string", "enum": [""]}]}, "schedule": {"title": "payout_schedule", "type": "object", "properties": {"interval": {"type": "string", "enum": ["daily", "manual", "monthly", "weekly"]}, "monthly_payout_days": {"type": "array", "items": {"type": "integer"}}, "weekly_payout_days": {"type": "array", "items": {"type": "string", "enum": ["friday", "monday", "thursday", "tuesday", "wednesday"], "x-stripeBypassValidation": true}}}}, "statement_descriptor": {"maxLength": 22, "type": "string"}}}, "settlement_timing": {"title": "settlement_timing", "type": "object", "properties": {"delay_days_override": {"anyOf": [{"type": "integer"}, {"type": "string", "enum": [""]}]}, "start_of_day": {"anyOf": [{"title": "start_of_day", "type": "object", "properties": {"hour": {"type": "integer"}, "minutes": {"type": "integer"}, "timezone": {"maxLength": 5000, "type": "string"}}}, {"type": "string", "enum": [""]}]}}}}, "description": "Settings that apply to the [Payments Balance](https://docs.stripe.com/api/balance)."}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "payments": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/balance_settings"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
