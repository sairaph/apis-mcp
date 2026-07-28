---
title: 'Test mode: Return an OutboundPayment'
page_id: operation-post-v1-test-helpers-treasury-outbound-payments-id-return-b5d60302
path: operations/untagged
description: <p>Transitions a test mode created OutboundPayment to the <code>returned</code> status. The OutboundPayment must already be in the <code>processing</code> state.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/test_helpers/treasury/outbound_payments/{id}/return
operation_ids:
    - PostTestHelpersTreasuryOutboundPaymentsIdReturn
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Test mode: Return an OutboundPayment

`POST /v1/test_helpers/treasury/outbound_payments/{id}/return`

Operation ID: `PostTestHelpersTreasuryOutboundPaymentsIdReturn`

<p>Transitions a test mode created OutboundPayment to the <code>returned</code> status. The OutboundPayment must already be in the <code>processing</code> state.</p>

## Definition

```yaml
{"summary": "Test mode: Return an OutboundPayment", "description": "<p>Transitions a test mode created OutboundPayment to the <code>returned</code> status. The OutboundPayment must already be in the <code>processing</code> state.</p>", "operationId": "PostTestHelpersTreasuryOutboundPaymentsIdReturn", "parameters": [{"name": "id", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "returned_details": {"title": "returned_details_params", "type": "object", "properties": {"code": {"type": "string", "enum": ["account_closed", "account_frozen", "bank_account_restricted", "bank_ownership_changed", "declined", "incorrect_account_holder_name", "invalid_account_number", "invalid_currency", "no_account", "other"]}}, "description": "Optional hash to set the return code."}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "returned_details": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/treasury.outbound_payment"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
