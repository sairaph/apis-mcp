---
title: List existing interconnects
page_id: operation-get-accounts-account-id-cni-interconnects-e5079965
path: operations/interconnects
description: List existing interconnects
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cni/interconnects
operation_ids:
    - list_interconnects
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List existing interconnects

`GET /accounts/{account_id}/cni/interconnects`

Operation ID: `list_interconnects`

## Definition

```yaml
{"operationId": "list_interconnects", "summary": "List existing interconnects", "parameters": [{"name": "site", "in": "query", "description": "If specified, only show interconnects located at the given site", "schema": {"type": "string", "nullable": true}}, {"name": "type", "in": "query", "description": "If specified, only show interconnects of the given type", "schema": {"type": "string", "nullable": true}}, {"name": "cursor", "in": "query", "schema": {"type": "integer", "format": "int32", "nullable": true}}, {"name": "limit", "in": "query", "schema": {"type": "integer", "minimum": 0, "nullable": true}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/nsc_AccountTag"}}], "responses": {"200": {"description": "List of matching interconnect objects", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/nsc_InterconnectList"}}}}, "400": {"description": "Bad request"}, "500": {"description": "Internal server error"}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Interconnects"], "x-api-token-group": ["Magic WAN Write", "Magic WAN Read", "Magic Transit Read", "Magic Transit Write"]}
```
