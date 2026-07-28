---
title: Retrieve a list of all slots matching the specified parameters
page_id: operation-get-accounts-account-id-cni-slots-9463ca94
path: operations/slots
description: Retrieve a list of all slots matching the specified parameters
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cni/slots
operation_ids:
    - list_slots
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Retrieve a list of all slots matching the specified parameters

`GET /accounts/{account_id}/cni/slots`

Operation ID: `list_slots`

## Definition

```yaml
{"operationId": "list_slots", "summary": "Retrieve a list of all slots matching the specified parameters", "parameters": [{"name": "address_contains", "in": "query", "description": "If specified, only show slots with the given text in their address field", "schema": {"type": "string", "nullable": true}}, {"name": "site", "in": "query", "description": "If specified, only show slots located at the given site", "schema": {"type": "string", "nullable": true}}, {"name": "speed", "in": "query", "description": "If specified, only show slots that support the given speed", "schema": {"type": "string", "nullable": true}}, {"name": "occupied", "in": "query", "description": "If specified, only show slots with a specific occupied/unoccupied state", "schema": {"type": "boolean", "nullable": true}}, {"name": "cursor", "in": "query", "schema": {"type": "integer", "format": "int32", "nullable": true}}, {"name": "limit", "in": "query", "schema": {"type": "integer", "minimum": 0, "nullable": true}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/nsc_AccountTag"}}], "responses": {"200": {"description": "List of matching slots", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/nsc_SlotList"}}}}, "400": {"description": "Bad request"}, "500": {"description": "Internal server error"}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Slots"], "x-api-token-group": ["Magic WAN Write", "Magic WAN Read", "Magic Transit Read", "Magic Transit Write"]}
```
