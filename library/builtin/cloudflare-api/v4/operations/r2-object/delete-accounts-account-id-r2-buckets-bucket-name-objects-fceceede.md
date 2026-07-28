---
title: Delete Objects
page_id: operation-delete-accounts-account-id-r2-buckets-bucket-name-objects-da114de7
path: operations/r2-object
description: |-
    Deletes multiple objects from an R2 bucket. Two modes are supported:

    1. **Delete by list** (default): Provide a JSON array of object keys in the request body.
       All listed objects are deleted; per-key errors are reported in the response.
    2. **Delete by prefix**: Provide the `prefix` query parameter (and an empty/no body)
       to schedule deletion of every object whose key begins with the given prefix.
       This kicks off an asynchronous prefix-delete job and returns the job descriptor.

    For most workloads, we recommend using R2's [S3-compatible API](https://developers.cloudflare.com/r2/api/s3/api/) or a [Worker with an R2 binding](https://developers.cloudflare.com/r2/api/workers/workers-api-reference/) instead.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/r2/buckets/{bucket_name}/objects
operation_ids:
    - r2-delete-objects
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Objects

`DELETE /accounts/{account_id}/r2/buckets/{bucket_name}/objects`

Operation ID: `r2-delete-objects`

Deletes multiple objects from an R2 bucket. Two modes are supported:

1. **Delete by list** (default): Provide a JSON array of object keys in the request body.
   All listed objects are deleted; per-key errors are reported in the response.
2. **Delete by prefix**: Provide the `prefix` query parameter (and an empty/no body)
   to schedule deletion of every object whose key begins with the given prefix.
   This kicks off an asynchronous prefix-delete job and returns the job descriptor.

For most workloads, we recommend using R2's [S3-compatible API](https://developers.cloudflare.com/r2/api/s3/api/) or a [Worker with an R2 binding](https://developers.cloudflare.com/r2/api/workers/workers-api-reference/) instead.

## Definition

```yaml
{"operationId": "r2-delete-objects", "summary": "Delete Objects", "description": "Deletes multiple objects from an R2 bucket. Two modes are supported:\n\n1. **Delete by list** (default): Provide a JSON array of object keys in the request body.\n   All listed objects are deleted; per-key errors are reported in the response.\n2. **Delete by prefix**: Provide the `prefix` query parameter (and an empty/no body)\n   to schedule deletion of every object whose key begins with the given prefix.\n   This kicks off an asynchronous prefix-delete job and returns the job descriptor.\n\nFor most workloads, we recommend using R2's [S3-compatible API](https://developers.cloudflare.com/r2/api/s3/api/) or a [Worker with an R2 binding](https://developers.cloudflare.com/r2/api/workers/workers-api-reference/) instead.\n", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_account_identifier"}}, {"name": "bucket_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_bucket_name"}}, {"name": "cf-r2-jurisdiction", "in": "header", "schema": {"$ref": "#/components/schemas/r2_jurisdiction"}}, {"name": "prefix", "in": "query", "schema": {"description": "When set, switches the operation to \"delete by prefix\" mode and asynchronously\ndeletes every object whose key begins with the given prefix. When omitted, the\nrequest body is interpreted as a JSON array of object keys to delete.\n", "type": "string"}}], "requestBody": {"description": "Required for \"delete by list\" mode (when `prefix` query parameter is omitted).\nA JSON array of object keys to delete. Ignored when `prefix` is provided.\n", "content": {"application/json": {"schema": {"type": "array", "items": {"type": "string"}, "example": ["path/to/object-a.txt", "path/to/object-b.txt"]}}}}, "responses": {"200": {"description": "Delete Objects response. Body shape depends on the operation mode:\n- \"delete by list\" returns an array of `{ key }` entries for successfully deleted objects;\n- \"delete by prefix\" returns the prefix-delete job descriptor.\n", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"$ref": "#/components/schemas/r2_errors"}, "messages": {"$ref": "#/components/schemas/r2_messages"}, "result": {"oneOf": [{"description": "Per-key delete results returned in \"delete by list\" mode.", "items": {"$ref": "#/components/schemas/r2_r2_delete_object_result"}, "type": "array"}, {"$ref": "#/components/schemas/r2_r2_delete_objects_by_prefix_result"}]}, "success": {"description": "Whether the API call was successful.", "type": "boolean", "enum": [true]}}, "required": ["success", "errors", "messages", "result"]}}}}, "4XX": {"description": "Delete Objects response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2_v4_response_failure"}}}}}, "security": [{"api_token": []}], "tags": ["R2 Object"], "x-api-token-group": ["Workers R2 Storage Write"]}
```
