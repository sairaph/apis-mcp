---
title: Reject a testmode personalization design
page_id: operation-post-v1-test-helpers-issuing-personalization-designs-personalization-des-f363483a
path: operations/untagged
description: <p>Updates the <code>status</code> of the specified testmode personalization design object to <code>rejected</code>.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/test_helpers/issuing/personalization_designs/{personalization_design}/reject
operation_ids:
    - PostTestHelpersIssuingPersonalizationDesignsPersonalizationDesignReject
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Reject a testmode personalization design

`POST /v1/test_helpers/issuing/personalization_designs/{personalization_design}/reject`

Operation ID: `PostTestHelpersIssuingPersonalizationDesignsPersonalizationDesignReject`

<p>Updates the <code>status</code> of the specified testmode personalization design object to <code>rejected</code>.</p>

## Definition

```yaml
{"summary": "Reject a testmode personalization design", "description": "<p>Updates the <code>status</code> of the specified testmode personalization design object to <code>rejected</code>.</p>", "operationId": "PostTestHelpersIssuingPersonalizationDesignsPersonalizationDesignReject", "parameters": [{"name": "personalization_design", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["rejection_reasons"], "type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "rejection_reasons": {"title": "rejection_reasons_param", "type": "object", "properties": {"card_logo": {"type": "array", "items": {"type": "string", "enum": ["geographic_location", "inappropriate", "network_name", "non_binary_image", "non_fiat_currency", "other", "other_entity", "promotional_material"]}}, "carrier_text": {"type": "array", "items": {"type": "string", "enum": ["geographic_location", "inappropriate", "network_name", "non_fiat_currency", "other", "other_entity", "promotional_material"]}}}, "description": "The reason(s) the personalization design was rejected."}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "rejection_reasons": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/issuing.personalization_design"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
