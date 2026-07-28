---
title: 'Test mode: Fail an OutboundTransfer'
page_id: operation-post-v1-test-helpers-treasury-outbound-transfers-outbound-transfer-fail-2895e682
path: operations/untagged
description: <p>Transitions a test mode created OutboundTransfer to the <code>failed</code> status. The OutboundTransfer must already be in the <code>processing</code> state.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/test_helpers/treasury/outbound_transfers/{outbound_transfer}/fail
operation_ids:
    - PostTestHelpersTreasuryOutboundTransfersOutboundTransferFail
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Test mode: Fail an OutboundTransfer

`POST /v1/test_helpers/treasury/outbound_transfers/{outbound_transfer}/fail`

Operation ID: `PostTestHelpersTreasuryOutboundTransfersOutboundTransferFail`

<p>Transitions a test mode created OutboundTransfer to the <code>failed</code> status. The OutboundTransfer must already be in the <code>processing</code> state.</p>

## Definition

```yaml
{"summary": "Test mode: Fail an OutboundTransfer", "description": "<p>Transitions a test mode created OutboundTransfer to the <code>failed</code> status. The OutboundTransfer must already be in the <code>processing</code> state.</p>", "operationId": "PostTestHelpersTreasuryOutboundTransfersOutboundTransferFail", "parameters": [{"name": "outbound_transfer", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/treasury.outbound_transfer"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
