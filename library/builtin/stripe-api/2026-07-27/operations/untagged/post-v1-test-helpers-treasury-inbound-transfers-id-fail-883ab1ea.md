---
title: 'Test mode: Fail an InboundTransfer'
page_id: operation-post-v1-test-helpers-treasury-inbound-transfers-id-fail-22871857
path: operations/untagged
description: <p>Transitions a test mode created InboundTransfer to the <code>failed</code> status. The InboundTransfer must already be in the <code>processing</code> state.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/test_helpers/treasury/inbound_transfers/{id}/fail
operation_ids:
    - PostTestHelpersTreasuryInboundTransfersIdFail
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Test mode: Fail an InboundTransfer

`POST /v1/test_helpers/treasury/inbound_transfers/{id}/fail`

Operation ID: `PostTestHelpersTreasuryInboundTransfersIdFail`

<p>Transitions a test mode created InboundTransfer to the <code>failed</code> status. The InboundTransfer must already be in the <code>processing</code> state.</p>

## Definition

```yaml
{"summary": "Test mode: Fail an InboundTransfer", "description": "<p>Transitions a test mode created InboundTransfer to the <code>failed</code> status. The InboundTransfer must already be in the <code>processing</code> state.</p>", "operationId": "PostTestHelpersTreasuryInboundTransfersIdFail", "parameters": [{"name": "id", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "failure_details": {"title": "failure_details_param", "type": "object", "properties": {"code": {"type": "string", "enum": ["account_closed", "account_frozen", "bank_account_restricted", "bank_ownership_changed", "debit_not_authorized", "incorrect_account_holder_address", "incorrect_account_holder_name", "incorrect_account_holder_tax_id", "insufficient_funds", "invalid_account_number", "invalid_currency", "no_account", "other"]}}, "description": "Details about a failed InboundTransfer."}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "failure_details": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/treasury.inbound_transfer"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
