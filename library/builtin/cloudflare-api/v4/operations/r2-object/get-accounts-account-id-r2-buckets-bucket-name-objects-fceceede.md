---
title: List Objects
page_id: operation-get-accounts-account-id-r2-buckets-bucket-name-objects-94554c72
path: operations/r2-object
description: |-
    Lists objects in an R2 bucket. Returns object metadata including key, size, etag, last modified date, HTTP metadata, and custom metadata.

    For most workloads, we recommend using R2's [S3-compatible API](https://developers.cloudflare.com/r2/api/s3/api/) or a [Worker with an R2 binding](https://developers.cloudflare.com/r2/api/workers/workers-api-reference/) instead.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/r2/buckets/{bucket_name}/objects
operation_ids:
    - r2-list-objects
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Objects

`GET /accounts/{account_id}/r2/buckets/{bucket_name}/objects`

Operation ID: `r2-list-objects`

Lists objects in an R2 bucket. Returns object metadata including key, size, etag, last modified date, HTTP metadata, and custom metadata.

For most workloads, we recommend using R2's [S3-compatible API](https://developers.cloudflare.com/r2/api/s3/api/) or a [Worker with an R2 binding](https://developers.cloudflare.com/r2/api/workers/workers-api-reference/) instead.

## Definition

```yaml
{"operationId": "r2-list-objects", "summary": "List Objects", "description": "Lists objects in an R2 bucket. Returns object metadata including key, size, etag, last modified date, HTTP metadata, and custom metadata.\n\nFor most workloads, we recommend using R2's [S3-compatible API](https://developers.cloudflare.com/r2/api/s3/api/) or a [Worker with an R2 binding](https://developers.cloudflare.com/r2/api/workers/workers-api-reference/) instead.\n", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_account_identifier"}}, {"name": "bucket_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_bucket_name"}}, {"name": "cf-r2-jurisdiction", "in": "header", "schema": {"$ref": "#/components/schemas/r2_jurisdiction"}}, {"name": "per_page", "in": "query", "schema": {"description": "Maximum number of objects to return per page.", "type": "integer", "default": 20, "maximum": 1000, "minimum": 1}}, {"name": "prefix", "in": "query", "schema": {"description": "Restricts results to only those objects whose keys begin with the specified prefix.", "type": "string"}}, {"name": "delimiter", "in": "query", "schema": {"description": "A single character used to group keys. All keys that contain the delimiter between the prefix and the first occurrence of the delimiter after the prefix are grouped under a single result element.", "type": "string"}}, {"name": "cursor", "in": "query", "schema": {"description": "Pagination cursor received from a previous List Objects call. Used to retrieve the next page of results.", "type": "string"}}, {"name": "start_after", "in": "query", "schema": {"description": "Returns objects with keys that come after the specified key in lexicographic order.", "type": "string"}}], "responses": {"200": {"description": "List Objects response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"$ref": "#/components/schemas/r2_errors"}, "messages": {"$ref": "#/components/schemas/r2_messages"}, "result": {"type": "array", "items": {"$ref": "#/components/schemas/r2_r2_object"}}, "result_info": {"$ref": "#/components/schemas/r2_r2_list_objects_result_info"}, "success": {"description": "Whether the API call was successful.", "type": "boolean", "enum": [true]}}, "required": ["success", "errors", "messages", "result"]}}}}, "4XX": {"description": "List Objects response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2_v4_response_failure"}}}}}, "security": [{"api_token": []}], "tags": ["R2 Object"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.edge.r2.bucket.read"]}}
```
