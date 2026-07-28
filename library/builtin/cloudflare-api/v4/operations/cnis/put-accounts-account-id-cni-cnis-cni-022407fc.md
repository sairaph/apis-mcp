---
title: Modify stored information about a CNI object
page_id: operation-put-accounts-account-id-cni-cnis-cni-ef2d09de
path: operations/cnis
description: Modify stored information about a CNI object
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/cni/cnis/{cni}
operation_ids:
    - update_cni
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Modify stored information about a CNI object

`PUT /accounts/{account_id}/cni/cnis/{cni}`

Operation ID: `update_cni`

## Definition

```yaml
{"operationId": "update_cni", "summary": "Modify stored information about a CNI object", "parameters": [{"name": "cni", "in": "path", "description": "CNI ID to retrieve information about", "required": true, "schema": {"type": "string", "format": "uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/nsc_AccountTag"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/nsc_Cni"}}}}, "responses": {"200": {"description": "CNI has been successfully modified", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/nsc_Cni"}}}}, "400": {"description": "Bad request"}, "404": {"description": "CNI not found"}, "500": {"description": "Internal server error"}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["CNIs"], "x-api-token-group": ["Magic WAN Write", "Magic Transit Write"]}
```
