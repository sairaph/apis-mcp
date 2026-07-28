---
title: Retrieve a VerificationReport
page_id: operation-get-v1-identity-verification-reports-report-037f05af
path: operations/untagged
description: <p>Retrieves an existing VerificationReport</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/identity/verification_reports/{report}
operation_ids:
    - GetIdentityVerificationReportsReport
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve a VerificationReport

`GET /v1/identity/verification_reports/{report}`

Operation ID: `GetIdentityVerificationReportsReport`

<p>Retrieves an existing VerificationReport</p>

## Definition

```yaml
{"summary": "Retrieve a VerificationReport", "description": "<p>Retrieves an existing VerificationReport</p>", "operationId": "GetIdentityVerificationReportsReport", "parameters": [{"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "report", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/identity.verification_report"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
