---
title: Retrieves a file from Binary Storage
page_id: operation-get-accounts-account-id-cloudforce-one-binary-hash-d2e79747
path: operations/bindb
description: Retrieves a binary file from the Cloudforce One binary storage for analysis.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/binary/{hash}
operation_ids:
    - get_BinDBGetBinary
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Retrieves a file from Binary Storage

`GET /accounts/{account_id}/cloudforce-one/binary/{hash}`

Operation ID: `get_BinDBGetBinary`

Retrieves a binary file from the Cloudforce One binary storage for analysis.

## Definition

```yaml
{"operationId": "get_BinDBGetBinary", "summary": "Retrieves a file from Binary Storage", "description": "Retrieves a binary file from the Cloudforce One binary storage for analysis.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "number"}}, {"name": "hash", "in": "path", "description": "hash of the binary", "required": true, "schema": {"description": "hash of the binary", "type": "string"}}], "responses": {"200": {"description": "Returns file information"}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["BinDB"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
