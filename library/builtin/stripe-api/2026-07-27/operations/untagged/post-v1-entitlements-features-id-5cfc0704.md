---
title: Updates a feature
page_id: operation-post-v1-entitlements-features-id-e8415ea3
path: operations/untagged
description: <p>Update a feature’s metadata or permanently deactivate it.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/entitlements/features/{id}
operation_ids:
    - PostEntitlementsFeaturesId
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Updates a feature

`POST /v1/entitlements/features/{id}`

Operation ID: `PostEntitlementsFeaturesId`

<p>Update a feature’s metadata or permanently deactivate it.</p>

## Definition

```yaml
{"summary": "Updates a feature", "description": "<p>Update a feature’s metadata or permanently deactivate it.</p>", "operationId": "PostEntitlementsFeaturesId", "parameters": [{"name": "id", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"active": {"type": "boolean", "description": "Inactive features cannot be attached to new products and will not be returned from the features list endpoint."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "metadata": {"description": "Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.", "anyOf": [{"type": "object", "additionalProperties": {"type": "string"}}, {"type": "string", "enum": [""]}]}, "name": {"maxLength": 80, "type": "string", "description": "The feature's name, for your own purpose, not meant to be displayable to the customer."}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/entitlements.feature"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
