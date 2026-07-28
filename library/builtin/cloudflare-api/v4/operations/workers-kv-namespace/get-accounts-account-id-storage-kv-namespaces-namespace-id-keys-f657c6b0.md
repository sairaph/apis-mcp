---
title: List a Namespace's Keys
page_id: operation-get-accounts-account-id-storage-kv-namespaces-namespace-id-keys-ae582b19
path: operations/workers-kv-namespace
description: Lists a namespace's keys.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/storage/kv/namespaces/{namespace_id}/keys
operation_ids:
    - workers-kv-namespace-list-a-namespace'-s-keys
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List a Namespace's Keys

`GET /accounts/{account_id}/storage/kv/namespaces/{namespace_id}/keys`

Operation ID: `workers-kv-namespace-list-a-namespace'-s-keys`

Lists a namespace's keys.

## Definition

```yaml
{"operationId": "workers-kv-namespace-list-a-namespace'-s-keys", "summary": "List a Namespace's Keys", "description": "Lists a namespace's keys.", "parameters": [{"name": "namespace_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers-kv_namespace_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers-kv_identifier"}}, {"name": "limit", "in": "query", "schema": {"description": "Limits the number of keys returned in the response. The cursor attribute may be used to iterate over the next batch of keys if there are more than the limit.", "type": "number", "default": 1000, "maximum": 1000, "minimum": 10}}, {"name": "prefix", "in": "query", "schema": {"description": "Filters returned keys by a name prefix. Exact matches and any key names that begin with the prefix will be returned.", "type": "string", "example": "My-Prefix"}}, {"name": "cursor", "in": "query", "schema": {"description": "Opaque token indicating the position from which to continue when requesting the next set of records if the amount of list results was limited by the limit parameter. A valid value for the cursor can be obtained from the `cursors` object in the `result_info` structure.", "type": "string", "example": "6Ck1la0VxJ0djhidm1MdX2FyDGxLKVeeHZZmORS_8XeSuhz9SjIJRaSa2lnsF01tQOHrfTGAP3R5X1Kv5iVUuMbNKhWNAXHOl6ePB0TUL8nw", "x-stainless-pagination-property": {"purpose": "next_cursor_param"}}}], "responses": {"200": {"description": "List a Namespace's Keys response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers-kv_api-response-common"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/workers-kv_key"}, "x-stainless-pagination-property": {"purpose": "items"}}, "result_info": {"$ref": "#/components/schemas/workers-kv_cursor_result_info"}}, "type": "object"}]}}}}, "4XX": {"description": "List a Namespace's Keys response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers-kv_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Workers KV Namespace"], "x-api-token-group": ["Workers KV Storage Write", "Workers KV Storage Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "keys", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
