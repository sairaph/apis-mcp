---
title: 'Test mode: Return an InboundTransfer'
page_id: operation-post-v1-test-helpers-treasury-inbound-transfers-id-return-7f2be6ff
path: operations/untagged
description: <p>Marks the test mode InboundTransfer object as returned and links the InboundTransfer to a ReceivedDebit. The InboundTransfer must already be in the <code>succeeded</code> state.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/test_helpers/treasury/inbound_transfers/{id}/return
operation_ids:
    - PostTestHelpersTreasuryInboundTransfersIdReturn
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Test mode: Return an InboundTransfer

`POST /v1/test_helpers/treasury/inbound_transfers/{id}/return`

Operation ID: `PostTestHelpersTreasuryInboundTransfersIdReturn`

<p>Marks the test mode InboundTransfer object as returned and links the InboundTransfer to a ReceivedDebit. The InboundTransfer must already be in the <code>succeeded</code> state.</p>

## Definition

```yaml
{"summary": "Test mode: Return an InboundTransfer", "description": "<p>Marks the test mode InboundTransfer object as returned and links the InboundTransfer to a ReceivedDebit. The InboundTransfer must already be in the <code>succeeded</code> state.</p>", "operationId": "PostTestHelpersTreasuryInboundTransfersIdReturn", "parameters": [{"name": "id", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/treasury.inbound_transfer"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
