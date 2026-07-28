---
title: Get Object
page_id: operation-get-accounts-account-id-r2-buckets-bucket-name-objects-object-key-31376f22
path: operations/r2-object
description: |-
    Retrieves an object from an R2 bucket. Returns the object body along with metadata headers.

    For most workloads, we recommend using R2's [S3-compatible API](https://developers.cloudflare.com/r2/api/s3/api/) or a [Worker with an R2 binding](https://developers.cloudflare.com/r2/api/workers/workers-api-reference/) instead.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/r2/buckets/{bucket_name}/objects/{object_key}
operation_ids:
    - r2-get-object
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Object

`GET /accounts/{account_id}/r2/buckets/{bucket_name}/objects/{object_key}`

Operation ID: `r2-get-object`

Retrieves an object from an R2 bucket. Returns the object body along with metadata headers.

For most workloads, we recommend using R2's [S3-compatible API](https://developers.cloudflare.com/r2/api/s3/api/) or a [Worker with an R2 binding](https://developers.cloudflare.com/r2/api/workers/workers-api-reference/) instead.

## Definition

```yaml
{"operationId": "r2-get-object", "summary": "Get Object", "description": "Retrieves an object from an R2 bucket. Returns the object body along with metadata headers.\n\nFor most workloads, we recommend using R2's [S3-compatible API](https://developers.cloudflare.com/r2/api/s3/api/) or a [Worker with an R2 binding](https://developers.cloudflare.com/r2/api/workers/workers-api-reference/) instead.\n", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_account_identifier"}}, {"name": "bucket_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_bucket_name"}}, {"name": "object_key", "in": "path", "required": true, "schema": {"description": "The key (name) of the object to retrieve. May contain slashes for path-like keys.\nSlashes (`/`) within the key MUST be sent literally and MUST NOT be percent-encoded\n(i.e. `%2F`); other reserved characters should be percent-encoded as usual.\n", "type": "string"}, "allowReserved": true, "example": "path/to/my-object.txt"}, {"name": "cf-r2-jurisdiction", "in": "header", "schema": {"$ref": "#/components/schemas/r2_jurisdiction"}}, {"name": "If-None-Match", "in": "header", "schema": {"description": "Returns the object only if its ETag does not match the given value.", "type": "string"}}, {"name": "If-Modified-Since", "in": "header", "schema": {"description": "Returns the object only if it has been modified since the specified time.\nMust be formatted as an HTTP-date (RFC 7231), e.g. `Tue, 15 Jan 2024 10:30:00 GMT`.\n", "type": "string"}}], "responses": {"200": {"description": "Get Object response. Returns the object body with metadata headers.", "headers": {"Content-Length": {"description": "The size of the object in bytes.", "schema": {"type": "integer"}}, "Content-Type": {"description": "The MIME type of the object.", "schema": {"type": "string"}}, "ETag": {"description": "The entity tag for the object, wrapped in double-quotes per RFC 7232,\ne.g. `\"d41d8cd98f00b204e9800998ecf8427e\"`.\n", "schema": {"type": "string"}}, "Last-Modified": {"description": "The date and time the object was last modified, formatted as an HTTP-date (RFC 7231),\ne.g. `Tue, 15 Jan 2024 10:30:00 GMT`.\n", "schema": {"type": "string"}}, "cf-r2-storage-class": {"description": "The storage class of the object.", "schema": {"$ref": "#/components/schemas/r2_storage_class"}}}, "content": {"application/octet-stream": {"schema": {"description": "The object body.", "type": "string", "format": "binary"}}}}, "304": {"description": "Not Modified. Returned when conditional request headers indicate the object has not changed."}, "4XX": {"description": "Get Object response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2_v4_response_failure"}}}}}, "security": [{"api_token": []}], "tags": ["R2 Object"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.edge.r2.bucket.read"]}}
```
