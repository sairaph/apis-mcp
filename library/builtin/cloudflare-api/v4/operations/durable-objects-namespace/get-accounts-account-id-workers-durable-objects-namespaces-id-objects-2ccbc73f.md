---
title: List Objects
page_id: operation-get-accounts-account-id-workers-durable-objects-namespaces-id-objects-fae2706f
path: operations/durable-objects-namespace
description: Returns the Durable Objects in a given namespace.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/workers/durable_objects/namespaces/{id}/objects
operation_ids:
    - durable-objects-namespace-list-objects
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Objects

`GET /accounts/{account_id}/workers/durable_objects/namespaces/{id}/objects`

Operation ID: `durable-objects-namespace-list-objects`

Returns the Durable Objects in a given namespace.

## Definition

```yaml
{"operationId": "durable-objects-namespace-list-objects", "summary": "List Objects", "description": "Returns the Durable Objects in a given namespace.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_schemas-id"}}, {"name": "limit", "in": "query", "schema": {"description": "The number of objects to return. The cursor attribute may be used to iterate over the next batch of objects if there are more than the limit.", "type": "number", "default": 1000, "maximum": 10000, "minimum": 10}}, {"name": "cursor", "in": "query", "schema": {"description": "Opaque token indicating the position from which to continue when requesting the next set of records. A valid value for the cursor can be obtained from the cursors object in the result_info structure.", "type": "string", "example": "AAAAANuhDN7SjacTnSVsDu3WW1Lvst6dxJGTjRY5BhxPXdf6L6uTcpd_NVtjhn11OUYRsVEykxoUwF-JQU4dn6QylZSKTOJuG0indrdn_MlHpMRtsxgXjs-RPdHYIVm3odE_uvEQ_dTQGFm8oikZMohns34DLBgrQpc"}}], "responses": {"200": {"description": "List Objects response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers_api-response-collection"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/workers_object"}}, "result_info": {"type": "object", "properties": {"count": {"description": "Total results returned based on your list parameters.", "type": "number", "example": 1}, "cursor": {"$ref": "#/components/schemas/workers_cursor"}}}}, "type": "object"}]}}}}, "4XX": {"description": "List Objects response failure.", "content": {"application/json": {"schema": {"allOf": [{"allOf": [{"$ref": "#/components/schemas/workers_api-response-collection"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/workers_object"}}, "result_info": {"type": "object", "properties": {"count": {"description": "Total results returned based on your list parameters.", "type": "number", "example": 1}, "cursor": {"$ref": "#/components/schemas/workers_cursor"}}}}, "type": "object"}]}, {"$ref": "#/components/schemas/workers_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Durable Objects Namespace"], "x-api-token-group": ["Workers Scripts Write", "Workers Scripts Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "durable-objects.namespaces.objects", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
