---
title: Get information about an interconnect object
page_id: operation-get-accounts-account-id-cni-interconnects-icon-6267fc7b
path: operations/interconnects
description: Get information about an interconnect object
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cni/interconnects/{icon}
operation_ids:
    - get_interconnect
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get information about an interconnect object

`GET /accounts/{account_id}/cni/interconnects/{icon}`

Operation ID: `get_interconnect`

## Definition

```yaml
{"operationId": "get_interconnect", "summary": "Get information about an interconnect object", "parameters": [{"name": "icon", "in": "path", "description": "Interconnect name to retrieve information about", "required": true, "schema": {"type": "string"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/nsc_AccountTag"}}], "responses": {"200": {"description": "Information about the specified interconnect", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/nsc_Interconnect"}}}}, "400": {"description": "Bad request"}, "404": {"description": "Interconnect not found"}, "500": {"description": "Internal server error"}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Interconnects"], "x-api-token-group": ["Magic WAN Write", "Magic WAN Read", "Magic Transit Read", "Magic Transit Write"]}
```
