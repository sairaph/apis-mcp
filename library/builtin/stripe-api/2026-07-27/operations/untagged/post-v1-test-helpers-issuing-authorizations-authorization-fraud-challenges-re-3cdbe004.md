---
title: Respond to fraud challenge
page_id: operation-post-v1-test-helpers-issuing-authorizations-authorization-fraud-challeng-b294cfc2
path: operations/untagged
description: <p>Respond to a fraud challenge on a testmode Issuing authorization, simulating either a confirmation of fraud or a correction of legitimacy.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/test_helpers/issuing/authorizations/{authorization}/fraud_challenges/respond
operation_ids:
    - PostTestHelpersIssuingAuthorizationsAuthorizationFraudChallengesRespond
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Respond to fraud challenge

`POST /v1/test_helpers/issuing/authorizations/{authorization}/fraud_challenges/respond`

Operation ID: `PostTestHelpersIssuingAuthorizationsAuthorizationFraudChallengesRespond`

<p>Respond to a fraud challenge on a testmode Issuing authorization, simulating either a confirmation of fraud or a correction of legitimacy.</p>

## Definition

```yaml
{"summary": "Respond to fraud challenge", "description": "<p>Respond to a fraud challenge on a testmode Issuing authorization, simulating either a confirmation of fraud or a correction of legitimacy.</p>", "operationId": "PostTestHelpersIssuingAuthorizationsAuthorizationFraudChallengesRespond", "parameters": [{"name": "authorization", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["confirmed"], "type": "object", "properties": {"confirmed": {"type": "boolean", "description": "Whether to simulate the user confirming that the transaction was legitimate (true) or telling Stripe that it was fraudulent (false)."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/issuing.authorization"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
