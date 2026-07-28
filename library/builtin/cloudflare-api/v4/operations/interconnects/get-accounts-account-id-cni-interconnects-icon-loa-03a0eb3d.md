---
title: Generate the Letter of Authorization (LOA) for a given interconnect
page_id: operation-get-accounts-account-id-cni-interconnects-icon-loa-1e999514
path: operations/interconnects
description: Generate the Letter of Authorization (LOA) for a given interconnect
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cni/interconnects/{icon}/loa
operation_ids:
    - get_interconnect_loa
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Generate the Letter of Authorization (LOA) for a given interconnect

`GET /accounts/{account_id}/cni/interconnects/{icon}/loa`

Operation ID: `get_interconnect_loa`

## Definition

```yaml
{"operationId": "get_interconnect_loa", "summary": "Generate the Letter of Authorization (LOA) for a given interconnect", "parameters": [{"name": "icon", "in": "path", "description": "Interconnect name to retrieve information about", "required": true, "schema": {"type": "string"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/nsc_AccountTag"}}], "responses": {"200": {"description": "Generated LOA in PDF format"}, "400": {"description": "Bad request"}, "404": {"description": "Interconnect not found"}, "500": {"description": "Internal server error"}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Interconnects"], "x-api-token-group": ["Magic WAN Write", "Magic WAN Read", "Magic Transit Read", "Magic Transit Write"]}
```
