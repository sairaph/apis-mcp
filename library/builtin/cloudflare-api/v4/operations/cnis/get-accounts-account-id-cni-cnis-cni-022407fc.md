---
title: Get information about a CNI object
page_id: operation-get-accounts-account-id-cni-cnis-cni-342a5282
path: operations/cnis
description: Get information about a CNI object
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cni/cnis/{cni}
operation_ids:
    - get_cni
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get information about a CNI object

`GET /accounts/{account_id}/cni/cnis/{cni}`

Operation ID: `get_cni`

## Definition

```yaml
{"operationId": "get_cni", "summary": "Get information about a CNI object", "parameters": [{"name": "cni", "in": "path", "description": "CNI ID to retrieve information about", "required": true, "schema": {"type": "string", "format": "uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/nsc_AccountTag"}}], "responses": {"200": {"description": "CNI's associated data", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/nsc_Cni"}}}}, "400": {"description": "Bad request"}, "404": {"description": "CNI not found"}, "500": {"description": "Internal server error"}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["CNIs"], "x-api-token-group": ["Magic WAN Write", "Magic WAN Read", "Magic Transit Read", "Magic Transit Write"]}
```
