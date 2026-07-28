---
title: Cancel an InboundTransfer
page_id: operation-post-v1-treasury-inbound-transfers-inbound-transfer-cancel-aa3ebdd8
path: operations/untagged
description: <p>Cancels an InboundTransfer.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/treasury/inbound_transfers/{inbound_transfer}/cancel
operation_ids:
    - PostTreasuryInboundTransfersInboundTransferCancel
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Cancel an InboundTransfer

`POST /v1/treasury/inbound_transfers/{inbound_transfer}/cancel`

Operation ID: `PostTreasuryInboundTransfersInboundTransferCancel`

<p>Cancels an InboundTransfer.</p>

## Definition

```yaml
{"summary": "Cancel an InboundTransfer", "description": "<p>Cancels an InboundTransfer.</p>", "operationId": "PostTreasuryInboundTransfersInboundTransferCancel", "parameters": [{"name": "inbound_transfer", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/treasury.inbound_transfer"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
