---
title: Delete a specified CNI object
page_id: operation-delete-accounts-account-id-cni-cnis-cni-8f9096be
path: operations/cnis
description: Delete a specified CNI object
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/cni/cnis/{cni}
operation_ids:
    - delete_cni
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a specified CNI object

`DELETE /accounts/{account_id}/cni/cnis/{cni}`

Operation ID: `delete_cni`

## Definition

```yaml
{"operationId": "delete_cni", "summary": "Delete a specified CNI object", "parameters": [{"name": "cni", "in": "path", "description": "CNI ID to retrieve information about", "required": true, "schema": {"type": "string", "format": "uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/nsc_AccountTag"}}], "responses": {"200": {"description": "CNI has been successfully deleted"}, "400": {"description": "Bad request"}, "404": {"description": "CNI not found"}, "500": {"description": "Internal server error"}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["CNIs"], "x-api-token-group": ["Magic WAN Write", "Magic Transit Write"]}
```
