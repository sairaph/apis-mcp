---
title: 'Test mode: Update an OutboundPayment'
page_id: operation-post-v1-test-helpers-treasury-outbound-payments-id-7af8d32a
path: operations/untagged
description: <p>Updates a test mode created OutboundPayment with tracking details. The OutboundPayment must not be cancelable, and cannot be in the <code>canceled</code> or <code>failed</code> states.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/test_helpers/treasury/outbound_payments/{id}
operation_ids:
    - PostTestHelpersTreasuryOutboundPaymentsId
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Test mode: Update an OutboundPayment

`POST /v1/test_helpers/treasury/outbound_payments/{id}`

Operation ID: `PostTestHelpersTreasuryOutboundPaymentsId`

<p>Updates a test mode created OutboundPayment with tracking details. The OutboundPayment must not be cancelable, and cannot be in the <code>canceled</code> or <code>failed</code> states.</p>

## Definition

```yaml
{"summary": "Test mode: Update an OutboundPayment", "description": "<p>Updates a test mode created OutboundPayment with tracking details. The OutboundPayment must not be cancelable, and cannot be in the <code>canceled</code> or <code>failed</code> states.</p>", "operationId": "PostTestHelpersTreasuryOutboundPaymentsId", "parameters": [{"name": "id", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["tracking_details"], "type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "tracking_details": {"title": "tracking_details_params", "required": ["type"], "type": "object", "properties": {"ach": {"title": "ach_tracking_details_params", "required": ["trace_id"], "type": "object", "properties": {"trace_id": {"maxLength": 5000, "type": "string"}}}, "type": {"type": "string", "enum": ["ach", "us_domestic_wire"]}, "us_domestic_wire": {"title": "us_domestic_wire_tracking_details_params", "type": "object", "properties": {"chips": {"maxLength": 5000, "type": "string"}, "imad": {"maxLength": 5000, "type": "string"}, "omad": {"maxLength": 5000, "type": "string"}}}}, "description": "Details about network-specific tracking information."}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "tracking_details": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/treasury.outbound_payment"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
