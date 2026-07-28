---
title: Create a CreditReversal
page_id: operation-post-v1-treasury-credit-reversals-d184890a
path: operations/untagged
description: <p>Reverses a ReceivedCredit and creates a CreditReversal object.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/treasury/credit_reversals
operation_ids:
    - PostTreasuryCreditReversals
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Create a CreditReversal

`POST /v1/treasury/credit_reversals`

Operation ID: `PostTreasuryCreditReversals`

<p>Reverses a ReceivedCredit and creates a CreditReversal object.</p>

## Definition

```yaml
{"summary": "Create a CreditReversal", "description": "<p>Reverses a ReceivedCredit and creates a CreditReversal object.</p>", "operationId": "PostTreasuryCreditReversals", "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["received_credit"], "type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "metadata": {"type": "object", "additionalProperties": {"type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`."}, "received_credit": {"maxLength": 5000, "type": "string", "description": "The ReceivedCredit to reverse."}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/treasury.credit_reversal"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
