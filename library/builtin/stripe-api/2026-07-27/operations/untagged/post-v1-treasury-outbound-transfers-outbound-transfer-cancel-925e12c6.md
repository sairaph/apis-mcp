---
title: Cancel an OutboundTransfer
page_id: operation-post-v1-treasury-outbound-transfers-outbound-transfer-cancel-83224822
path: operations/untagged
description: <p>An OutboundTransfer can be canceled if the funds have not yet been paid out.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/treasury/outbound_transfers/{outbound_transfer}/cancel
operation_ids:
    - PostTreasuryOutboundTransfersOutboundTransferCancel
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Cancel an OutboundTransfer

`POST /v1/treasury/outbound_transfers/{outbound_transfer}/cancel`

Operation ID: `PostTreasuryOutboundTransfersOutboundTransferCancel`

<p>An OutboundTransfer can be canceled if the funds have not yet been paid out.</p>

## Definition

```yaml
{"summary": "Cancel an OutboundTransfer", "description": "<p>An OutboundTransfer can be canceled if the funds have not yet been paid out.</p>", "operationId": "PostTreasuryOutboundTransfersOutboundTransferCancel", "parameters": [{"name": "outbound_transfer", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/treasury.outbound_transfer"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
