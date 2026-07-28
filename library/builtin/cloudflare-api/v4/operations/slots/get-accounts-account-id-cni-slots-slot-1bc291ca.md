---
title: Get information about the specified slot
page_id: operation-get-accounts-account-id-cni-slots-slot-4bb720a7
path: operations/slots
description: Get information about the specified slot
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cni/slots/{slot}
operation_ids:
    - get_slot
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get information about the specified slot

`GET /accounts/{account_id}/cni/slots/{slot}`

Operation ID: `get_slot`

## Definition

```yaml
{"operationId": "get_slot", "summary": "Get information about the specified slot", "parameters": [{"name": "slot", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/nsc_AccountTag"}}], "responses": {"200": {"description": "Information about the specified slot", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/nsc_SlotInfo"}}}}, "400": {"description": "Bad request"}, "404": {"description": "Slot not found"}, "500": {"description": "Internal server error"}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Slots"], "x-api-token-group": ["Magic WAN Write", "Magic WAN Read", "Magic Transit Read", "Magic Transit Write"]}
```
