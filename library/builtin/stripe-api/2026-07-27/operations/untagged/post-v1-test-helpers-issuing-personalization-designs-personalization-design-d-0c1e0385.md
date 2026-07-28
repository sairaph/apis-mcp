---
title: Deactivate a testmode personalization design
page_id: operation-post-v1-test-helpers-issuing-personalization-designs-personalization-des-ab9cdbf1
path: operations/untagged
description: <p>Updates the <code>status</code> of the specified testmode personalization design object to <code>inactive</code>.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/test_helpers/issuing/personalization_designs/{personalization_design}/deactivate
operation_ids:
    - PostTestHelpersIssuingPersonalizationDesignsPersonalizationDesignDeactivate
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Deactivate a testmode personalization design

`POST /v1/test_helpers/issuing/personalization_designs/{personalization_design}/deactivate`

Operation ID: `PostTestHelpersIssuingPersonalizationDesignsPersonalizationDesignDeactivate`

<p>Updates the <code>status</code> of the specified testmode personalization design object to <code>inactive</code>.</p>

## Definition

```yaml
{"summary": "Deactivate a testmode personalization design", "description": "<p>Updates the <code>status</code> of the specified testmode personalization design object to <code>inactive</code>.</p>", "operationId": "PostTestHelpersIssuingPersonalizationDesignsPersonalizationDesignDeactivate", "parameters": [{"name": "personalization_design", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/issuing.personalization_design"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
