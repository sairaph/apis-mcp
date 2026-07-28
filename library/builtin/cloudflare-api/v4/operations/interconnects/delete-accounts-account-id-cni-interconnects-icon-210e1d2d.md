---
title: Delete an interconnect object
page_id: operation-delete-accounts-account-id-cni-interconnects-icon-494a8ad2
path: operations/interconnects
description: Delete an interconnect object
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/cni/interconnects/{icon}
operation_ids:
    - delete_interconnect
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete an interconnect object

`DELETE /accounts/{account_id}/cni/interconnects/{icon}`

Operation ID: `delete_interconnect`

## Definition

```yaml
{"operationId": "delete_interconnect", "summary": "Delete an interconnect object", "parameters": [{"name": "icon", "in": "path", "description": "Interconnect name to retrieve information about", "required": true, "schema": {"type": "string"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/nsc_AccountTag"}}], "responses": {"200": {"description": "Successfully deleted interconnect"}, "400": {"description": "Bad request"}, "404": {"description": "Interconnect not found"}, "500": {"description": "Internal server error"}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Interconnects"], "x-api-token-group": ["Magic WAN Write", "Magic Transit Write"]}
```
