---
title: Delete logo query
page_id: operation-delete-accounts-account-id-cloudforce-one-v2-brand-protection-logo-queri-236338f0
path: operations/brand-protection
description: Delete a saved brand protection logo query. Returns 404 if the query ID doesn't exist.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/v2/brand-protection/logo/queries/{query_id}
operation_ids:
    - delete_DeleteLogoQuery
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete logo query

`DELETE /accounts/{account_id}/cloudforce-one/v2/brand-protection/logo/queries/{query_id}`

Operation ID: `delete_DeleteLogoQuery`

Delete a saved brand protection logo query. Returns 404 if the query ID doesn't exist.

## Definition

```yaml
{"operationId": "delete_DeleteLogoQuery", "summary": "Delete logo query", "description": "Delete a saved brand protection logo query. Returns 404 if the query ID doesn't exist.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}, {"name": "query_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}], "responses": {"200": {"description": "Logo query deleted successfully", "content": {"application/json": {"schema": {"type": "object", "properties": {"message": {"type": "string"}, "success": {"type": "boolean"}}, "required": ["success", "message"]}}}}}, "security": [{"api_token": []}], "tags": ["Brand Protection"], "x-api-token-group": ["Cloudforce One Write"]}
```
