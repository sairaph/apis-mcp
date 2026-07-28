---
title: Verify microdeposits on a SetupIntent
page_id: operation-post-v1-setup-intents-intent-verify-microdeposits-882ddaab
path: operations/untagged
description: <p>Verifies microdeposits on a SetupIntent object.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/setup_intents/{intent}/verify_microdeposits
operation_ids:
    - PostSetupIntentsIntentVerifyMicrodeposits
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Verify microdeposits on a SetupIntent

`POST /v1/setup_intents/{intent}/verify_microdeposits`

Operation ID: `PostSetupIntentsIntentVerifyMicrodeposits`

<p>Verifies microdeposits on a SetupIntent object.</p>

## Definition

```yaml
{"summary": "Verify microdeposits on a SetupIntent", "description": "<p>Verifies microdeposits on a SetupIntent object.</p>", "operationId": "PostSetupIntentsIntentVerifyMicrodeposits", "parameters": [{"name": "intent", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"amounts": {"type": "array", "description": "Two positive integers, in *cents*, equal to the values of the microdeposits sent to the bank account.", "items": {"type": "integer"}}, "client_secret": {"maxLength": 5000, "type": "string", "description": "The client secret of the SetupIntent."}, "descriptor_code": {"maxLength": 5000, "type": "string", "description": "A six-character code starting with SM present in the microdeposit sent to the bank account."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}}, "additionalProperties": false}, "encoding": {"amounts": {"style": "deepObject", "explode": true}, "expand": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/setup_intent"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
