---
title: Delete Object
page_id: operation-delete-accounts-account-id-r2-buckets-bucket-name-objects-object-key-35d071ce
path: operations/r2-object
description: |-
    Deletes an object from an R2 bucket.

    For most workloads, we recommend using R2's [S3-compatible API](https://developers.cloudflare.com/r2/api/s3/api/) or a [Worker with an R2 binding](https://developers.cloudflare.com/r2/api/workers/workers-api-reference/) instead.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/r2/buckets/{bucket_name}/objects/{object_key}
operation_ids:
    - r2-delete-object
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Object

`DELETE /accounts/{account_id}/r2/buckets/{bucket_name}/objects/{object_key}`

Operation ID: `r2-delete-object`

Deletes an object from an R2 bucket.

For most workloads, we recommend using R2's [S3-compatible API](https://developers.cloudflare.com/r2/api/s3/api/) or a [Worker with an R2 binding](https://developers.cloudflare.com/r2/api/workers/workers-api-reference/) instead.

## Definition

```yaml
{"operationId": "r2-delete-object", "summary": "Delete Object", "description": "Deletes an object from an R2 bucket.\n\nFor most workloads, we recommend using R2's [S3-compatible API](https://developers.cloudflare.com/r2/api/s3/api/) or a [Worker with an R2 binding](https://developers.cloudflare.com/r2/api/workers/workers-api-reference/) instead.\n", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_account_identifier"}}, {"name": "bucket_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_bucket_name"}}, {"name": "object_key", "in": "path", "required": true, "schema": {"description": "The key (name) of the object to delete. May contain slashes for path-like keys.\nSlashes (`/`) within the key MUST be sent literally and MUST NOT be percent-encoded\n(i.e. `%2F`); other reserved characters should be percent-encoded as usual.\n", "type": "string"}, "allowReserved": true, "example": "path/to/my-object.txt"}, {"name": "cf-r2-jurisdiction", "in": "header", "schema": {"$ref": "#/components/schemas/r2_jurisdiction"}}], "responses": {"200": {"description": "Delete Object response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/r2_v4_response"}, {"properties": {"result": {"$ref": "#/components/schemas/r2_r2_delete_object_result"}}, "type": "object"}]}}}}, "4XX": {"description": "Delete Object response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2_v4_response_failure"}}}}}, "security": [{"api_token": []}], "tags": ["R2 Object"], "x-api-token-group": ["Workers R2 Storage Write"]}
```
