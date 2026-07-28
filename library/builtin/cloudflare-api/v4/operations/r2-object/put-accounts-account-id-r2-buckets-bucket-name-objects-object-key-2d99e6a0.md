---
title: Upload Object
page_id: operation-put-accounts-account-id-r2-buckets-bucket-name-objects-object-key-4eec3bc3
path: operations/r2-object
description: |-
    Uploads an object to an R2 bucket. The object body is provided as the request body. Returns metadata about the uploaded object.

    The maximum upload size for this endpoint is 300 MB. For most workloads, we recommend using R2's [S3-compatible API](https://developers.cloudflare.com/r2/api/s3/api/) or a [Worker with an R2 binding](https://developers.cloudflare.com/r2/api/workers/workers-api-reference/) instead.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/r2/buckets/{bucket_name}/objects/{object_key}
operation_ids:
    - r2-put-object
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Upload Object

`PUT /accounts/{account_id}/r2/buckets/{bucket_name}/objects/{object_key}`

Operation ID: `r2-put-object`

Uploads an object to an R2 bucket. The object body is provided as the request body. Returns metadata about the uploaded object.

The maximum upload size for this endpoint is 300 MB. For most workloads, we recommend using R2's [S3-compatible API](https://developers.cloudflare.com/r2/api/s3/api/) or a [Worker with an R2 binding](https://developers.cloudflare.com/r2/api/workers/workers-api-reference/) instead.

## Definition

```yaml
{"operationId": "r2-put-object", "summary": "Upload Object", "description": "Uploads an object to an R2 bucket. The object body is provided as the request body. Returns metadata about the uploaded object.\n\nThe maximum upload size for this endpoint is 300 MB. For most workloads, we recommend using R2's [S3-compatible API](https://developers.cloudflare.com/r2/api/s3/api/) or a [Worker with an R2 binding](https://developers.cloudflare.com/r2/api/workers/workers-api-reference/) instead.\n", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_account_identifier"}}, {"name": "bucket_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_bucket_name"}}, {"name": "object_key", "in": "path", "required": true, "schema": {"description": "The key (name) to assign to the object. May contain slashes for path-like keys.\nSlashes (`/`) within the key MUST be sent literally and MUST NOT be percent-encoded\n(i.e. `%2F`); other reserved characters should be percent-encoded as usual.\n", "type": "string"}, "allowReserved": true, "example": "path/to/my-object.txt"}, {"name": "cf-r2-jurisdiction", "in": "header", "schema": {"$ref": "#/components/schemas/r2_jurisdiction"}}, {"name": "Content-Type", "in": "header", "schema": {"description": "The MIME type of the object being uploaded.", "type": "string"}}, {"name": "Content-Length", "in": "header", "schema": {"description": "The size of the object body in bytes.", "type": "integer"}}, {"name": "cf-r2-storage-class", "in": "header", "description": "Storage class for this object. Overrides the bucket default.", "schema": {"$ref": "#/components/schemas/r2_storage_class"}}], "requestBody": {"required": true, "content": {"application/octet-stream": {"schema": {"description": "The object body to upload.", "type": "string", "format": "binary"}}}}, "responses": {"200": {"description": "Upload Object response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/r2_v4_response"}, {"properties": {"result": {"$ref": "#/components/schemas/r2_r2_put_object_result"}}, "type": "object"}]}}}}, "4XX": {"description": "Upload Object response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2_v4_response_failure"}}}}}, "security": [{"api_token": []}], "tags": ["R2 Object"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.edge.r2.bucket.write"]}}
```
