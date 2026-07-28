---
title: List existing CNI objects
page_id: operation-get-accounts-account-id-cni-cnis-2ca3d3d6
path: operations/cnis
description: List existing CNI objects
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cni/cnis
operation_ids:
    - list_cnis
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List existing CNI objects

`GET /accounts/{account_id}/cni/cnis`

Operation ID: `list_cnis`

## Definition

```yaml
{"operationId": "list_cnis", "summary": "List existing CNI objects", "parameters": [{"name": "slot", "in": "query", "description": "If specified, only show CNIs associated with the specified slot", "schema": {"type": "string", "nullable": true}}, {"name": "tunnel_id", "in": "query", "description": "If specified, only show cnis associated with the specified tunnel id", "schema": {"type": "string", "nullable": true}}, {"name": "cursor", "in": "query", "schema": {"type": "integer", "format": "int32", "nullable": true}}, {"name": "limit", "in": "query", "schema": {"type": "integer", "minimum": 0, "nullable": true}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/nsc_AccountTag"}}], "responses": {"200": {"description": "List of matching CNI objects", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/nsc_CniList"}}}}, "400": {"description": "Bad request"}, "500": {"description": "Internal server error"}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["CNIs"], "x-api-token-group": ["Magic WAN Write", "Magic WAN Read", "Magic Transit Read", "Magic Transit Write"]}
```
