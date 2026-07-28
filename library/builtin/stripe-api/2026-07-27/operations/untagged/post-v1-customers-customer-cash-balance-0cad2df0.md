---
title: Update a cash balance's settings
page_id: operation-post-v1-customers-customer-cash-balance-8fd0bf81
path: operations/untagged
description: <p>Changes the settings on a customer’s cash balance.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/customers/{customer}/cash_balance
operation_ids:
    - PostCustomersCustomerCashBalance
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Update a cash balance's settings

`POST /v1/customers/{customer}/cash_balance`

Operation ID: `PostCustomersCustomerCashBalance`

<p>Changes the settings on a customer’s cash balance.</p>

## Definition

```yaml
{"summary": "Update a cash balance's settings", "description": "<p>Changes the settings on a customer’s cash balance.</p>", "operationId": "PostCustomersCustomerCashBalance", "parameters": [{"name": "customer", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "settings": {"title": "balance_settings_param", "type": "object", "properties": {"reconciliation_mode": {"type": "string", "enum": ["automatic", "manual", "merchant_default"]}}, "description": "A hash of settings for this cash balance."}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "settings": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cash_balance"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
