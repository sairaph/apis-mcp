---
title: Update a personalization design
page_id: operation-post-v1-issuing-personalization-designs-personalization-design-049a7ee8
path: operations/untagged
description: <p>Updates a card personalization object.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/issuing/personalization_designs/{personalization_design}
operation_ids:
    - PostIssuingPersonalizationDesignsPersonalizationDesign
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Update a personalization design

`POST /v1/issuing/personalization_designs/{personalization_design}`

Operation ID: `PostIssuingPersonalizationDesignsPersonalizationDesign`

<p>Updates a card personalization object.</p>

## Definition

```yaml
{"summary": "Update a personalization design", "description": "<p>Updates a card personalization object.</p>", "operationId": "PostIssuingPersonalizationDesignsPersonalizationDesign", "parameters": [{"name": "personalization_design", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"card_logo": {"description": "The file for the card logo, for use with physical bundles that support card logos. Must have a `purpose` value of `issuing_logo`.", "anyOf": [{"type": "string"}, {"type": "string", "enum": [""]}]}, "carrier_text": {"description": "Hash containing carrier text, for use with physical bundles that support carrier text.", "anyOf": [{"title": "carrier_text_param", "type": "object", "properties": {"footer_body": {"anyOf": [{"maxLength": 200, "type": "string"}, {"type": "string", "enum": [""]}]}, "footer_title": {"anyOf": [{"maxLength": 30, "type": "string"}, {"type": "string", "enum": [""]}]}, "header_body": {"anyOf": [{"maxLength": 200, "type": "string"}, {"type": "string", "enum": [""]}]}, "header_title": {"anyOf": [{"maxLength": 30, "type": "string"}, {"type": "string", "enum": [""]}]}}}, {"type": "string", "enum": [""]}]}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "lookup_key": {"description": "A lookup key used to retrieve personalization designs dynamically from a static string. This may be up to 200 characters.", "anyOf": [{"maxLength": 200, "type": "string"}, {"type": "string", "enum": [""]}]}, "metadata": {"type": "object", "additionalProperties": {"type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`."}, "name": {"description": "Friendly display name. Providing an empty string will set the field to null.", "anyOf": [{"maxLength": 200, "type": "string"}, {"type": "string", "enum": [""]}]}, "physical_bundle": {"maxLength": 5000, "type": "string", "description": "The physical bundle object belonging to this personalization design."}, "preferences": {"title": "preferences_param", "required": ["is_default"], "type": "object", "properties": {"is_default": {"type": "boolean"}}, "description": "Information on whether this personalization design is used to create cards when one is not specified."}, "transfer_lookup_key": {"type": "boolean", "description": "If set to true, will atomically remove the lookup key from the existing personalization design, and assign it to this personalization design."}}, "additionalProperties": false}, "encoding": {"card_logo": {"style": "deepObject", "explode": true}, "carrier_text": {"style": "deepObject", "explode": true}, "expand": {"style": "deepObject", "explode": true}, "lookup_key": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}, "name": {"style": "deepObject", "explode": true}, "preferences": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/issuing.personalization_design"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
