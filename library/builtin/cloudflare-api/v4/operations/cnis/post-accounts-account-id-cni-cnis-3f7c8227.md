---
title: Create a new CNI object
page_id: operation-post-accounts-account-id-cni-cnis-b5e446a2
path: operations/cnis
description: Create a new CNI object
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/cni/cnis
operation_ids:
    - create_cni
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a new CNI object

`POST /accounts/{account_id}/cni/cnis`

Operation ID: `create_cni`

## Definition

```yaml
{"operationId": "create_cni", "summary": "Create a new CNI object", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/nsc_AccountTag"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/nsc_CniCreate"}}}}, "responses": {"200": {"description": "CNI was successfully created", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/nsc_Cni"}}}}, "400": {"description": "Bad request"}, "409": {"description": "Name Conflict"}, "500": {"description": "Internal server error"}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["CNIs"], "x-api-token-group": ["Magic WAN Write", "Magic Transit Write"]}
```
