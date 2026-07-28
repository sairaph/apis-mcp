---
title: Create a new interconnect
page_id: operation-post-accounts-account-id-cni-interconnects-6eb711a3
path: operations/interconnects
description: Create a new interconnect
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/cni/interconnects
operation_ids:
    - create_interconnect
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a new interconnect

`POST /accounts/{account_id}/cni/interconnects`

Operation ID: `create_interconnect`

## Definition

```yaml
{"operationId": "create_interconnect", "summary": "Create a new interconnect", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/nsc_AccountTag"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/nsc_InterconnectCreate"}}}}, "responses": {"200": {"description": "Information about the new interconnect", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/nsc_Interconnect"}}}}, "400": {"description": "Bad request"}, "500": {"description": "Internal server error"}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Interconnects"], "x-api-token-group": ["Magic WAN Write", "Magic Transit Write"]}
```
