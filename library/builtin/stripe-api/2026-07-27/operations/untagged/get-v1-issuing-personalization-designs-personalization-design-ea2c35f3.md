---
title: Retrieve a personalization design
page_id: operation-get-v1-issuing-personalization-designs-personalization-design-92e31b5e
path: operations/untagged
description: <p>Retrieves a personalization design object.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/issuing/personalization_designs/{personalization_design}
operation_ids:
    - GetIssuingPersonalizationDesignsPersonalizationDesign
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve a personalization design

`GET /v1/issuing/personalization_designs/{personalization_design}`

Operation ID: `GetIssuingPersonalizationDesignsPersonalizationDesign`

<p>Retrieves a personalization design object.</p>

## Definition

```yaml
{"summary": "Retrieve a personalization design", "description": "<p>Retrieves a personalization design object.</p>", "operationId": "GetIssuingPersonalizationDesignsPersonalizationDesign", "parameters": [{"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "personalization_design", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/issuing.personalization_design"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
