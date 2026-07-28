---
title: Posts a file to Binary Storage
page_id: operation-post-accounts-account-id-cloudforce-one-binary-71b9dfef
path: operations/bindb
description: Uploads a binary file to Cloudforce One's binary database for malware analysis and threat intelligence correlation.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/binary
operation_ids:
    - post_BinDBPost
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Posts a file to Binary Storage

`POST /accounts/{account_id}/cloudforce-one/binary`

Operation ID: `post_BinDBPost`

Uploads a binary file to Cloudforce One's binary database for malware analysis and threat intelligence correlation.

## Definition

```yaml
{"operationId": "post_BinDBPost", "summary": "Posts a file to Binary Storage", "description": "Uploads a binary file to Cloudforce One's binary database for malware analysis and threat intelligence correlation.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "number"}}], "requestBody": {"description": "Binary file to be uploaded to the database.", "required": true, "content": {"multipart/form-data": {"schema": {"type": "object", "properties": {"file": {"description": "The binary file content to upload.", "type": "string", "format": "binary"}}, "required": ["file"]}}}}, "responses": {"200": {"description": "Returns file information", "content": {"application/json": {"schema": {"type": "object", "properties": {"content_type": {"type": "string", "example": "text/plain"}, "md5": {"type": "string", "example": "5d84ade76d2a8387c81175bb0cbe6492"}, "sha1": {"type": "string", "example": "9aff6879626d957eafadda044e4f879aae1e7278"}, "sha256": {"type": "string", "example": "0000a7f2692ef479e2e3d02661568882cadec451cc8a64d4e7faca29810cd626"}}, "required": ["md5", "sha1", "sha256", "content_type"]}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["BinDB"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
