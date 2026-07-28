---
title: Create a feature
page_id: operation-post-v1-entitlements-features-a6dc9474
path: operations/untagged
description: <p>Creates a feature</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/entitlements/features
operation_ids:
    - PostEntitlementsFeatures
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Create a feature

`POST /v1/entitlements/features`

Operation ID: `PostEntitlementsFeatures`

<p>Creates a feature</p>

## Definition

```yaml
{"summary": "Create a feature", "description": "<p>Creates a feature</p>", "operationId": "PostEntitlementsFeatures", "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["lookup_key", "name"], "type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "lookup_key": {"maxLength": 80, "type": "string", "description": "A unique key you provide as your own system identifier. This may be up to 80 characters."}, "metadata": {"type": "object", "additionalProperties": {"type": "string"}, "description": "Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format."}, "name": {"maxLength": 80, "type": "string", "description": "The feature's name, for your own purpose, not meant to be displayable to the customer."}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/entitlements.feature"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
