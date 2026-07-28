---
title: Retrieve an early fraud warning
page_id: operation-get-v1-radar-early-fraud-warnings-early-fraud-warning-290c63c0
path: operations/untagged
description: |-
    <p>Retrieves the details of an early fraud warning that has previously been created. </p>

    <p>Please refer to the <a href="#early_fraud_warning_object">early fraud warning</a> object reference for more details.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/radar/early_fraud_warnings/{early_fraud_warning}
operation_ids:
    - GetRadarEarlyFraudWarningsEarlyFraudWarning
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve an early fraud warning

`GET /v1/radar/early_fraud_warnings/{early_fraud_warning}`

Operation ID: `GetRadarEarlyFraudWarningsEarlyFraudWarning`

<p>Retrieves the details of an early fraud warning that has previously been created. </p>

<p>Please refer to the <a href="#early_fraud_warning_object">early fraud warning</a> object reference for more details.</p>

## Definition

```yaml
{"summary": "Retrieve an early fraud warning", "description": "<p>Retrieves the details of an early fraud warning that has previously been created. </p>\n\n<p>Please refer to the <a href=\"#early_fraud_warning_object\">early fraud warning</a> object reference for more details.</p>", "operationId": "GetRadarEarlyFraudWarningsEarlyFraudWarning", "parameters": [{"name": "early_fraud_warning", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/radar.early_fraud_warning"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
