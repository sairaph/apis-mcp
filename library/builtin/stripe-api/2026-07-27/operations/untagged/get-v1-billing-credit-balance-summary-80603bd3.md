---
title: Retrieve the credit balance summary for a customer
page_id: operation-get-v1-billing-credit-balance-summary-0fe783ad
path: operations/untagged
description: <p>Retrieves the credit balance summary for a customer.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/billing/credit_balance_summary
operation_ids:
    - GetBillingCreditBalanceSummary
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve the credit balance summary for a customer

`GET /v1/billing/credit_balance_summary`

Operation ID: `GetBillingCreditBalanceSummary`

<p>Retrieves the credit balance summary for a customer.</p>

## Definition

```yaml
{"summary": "Retrieve the credit balance summary for a customer", "description": "<p>Retrieves the credit balance summary for a customer.</p>", "operationId": "GetBillingCreditBalanceSummary", "parameters": [{"name": "customer", "in": "query", "description": "The customer whose credit balance summary you're retrieving.", "required": false, "style": "form", "explode": true, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "customer_account", "in": "query", "description": "The account representing the customer whose credit balance summary you're retrieving.", "required": false, "style": "form", "explode": true, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "filter", "in": "query", "description": "The filter criteria for the credit balance summary.", "required": true, "style": "deepObject", "explode": true, "schema": {"title": "balance_summary_filter_param", "required": ["type"], "type": "object", "properties": {"applicability_scope": {"title": "scope_param", "type": "object", "properties": {"price_type": {"type": "string", "enum": ["metered"]}, "prices": {"type": "array", "items": {"title": "applicable_price_param", "required": ["id"], "type": "object", "properties": {"id": {"maxLength": 5000, "type": "string"}}}}}}, "credit_grant": {"maxLength": 5000, "type": "string"}, "type": {"type": "string", "enum": ["applicability_scope", "credit_grant"]}}}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/billing.credit_balance_summary"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
