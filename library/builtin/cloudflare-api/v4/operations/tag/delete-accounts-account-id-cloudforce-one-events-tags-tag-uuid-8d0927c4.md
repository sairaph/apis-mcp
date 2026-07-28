---
title: Deletes a tag (SoT)
page_id: operation-delete-accounts-account-id-cloudforce-one-events-tags-tag-uuid-3ffeb246
path: operations/tag
description: Deletes a Source-of-Truth tag by UUID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/tags/{tag_uuid}
operation_ids:
    - delete_TagDelete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Deletes a tag (SoT)

`DELETE /accounts/{account_id}/cloudforce-one/events/tags/{tag_uuid}`

Operation ID: `delete_TagDelete`

Deletes a Source-of-Truth tag by UUID.

## Definition

```yaml
{"operationId": "delete_TagDelete", "summary": "Deletes a tag (SoT)", "description": "Deletes a Source-of-Truth tag by UUID.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "tag_uuid", "in": "path", "description": "Tag UUID.", "required": true, "schema": {"description": "Tag UUID.", "type": "string"}}], "responses": {"200": {"description": "Returns the uuid of the deleted tag.", "content": {"application/json": {"schema": {"type": "object", "properties": {"uuid": {"type": "string", "example": "12345678-1234-1234-1234-1234567890ab"}}, "required": ["uuid"]}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}, "404": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["Tag"], "x-api-token-group": ["Cloudforce One Write"]}
```
