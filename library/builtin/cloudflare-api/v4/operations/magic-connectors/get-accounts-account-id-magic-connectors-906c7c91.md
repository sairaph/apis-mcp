---
title: List Connectors
page_id: operation-get-accounts-account-id-magic-connectors-dc1578cd
path: operations/magic-connectors
description: Lists Magic WAN Connectors.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/magic/connectors
operation_ids:
    - mconn-connectors-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Connectors

`GET /accounts/{account_id}/magic/connectors`

Operation ID: `mconn-connectors-list`

Lists Magic WAN Connectors.

## Definition

```yaml
{"operationId": "mconn-connectors-list", "summary": "List Connectors", "description": "Lists Magic WAN Connectors.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mconn_account_id"}}, {"name": "device_type", "in": "query", "description": "Filter connectors by device type.", "schema": {"type": "string", "enum": ["MANAGED", "LICENSED"]}}], "responses": {"200": {"description": "OK", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mconn_customer_connectors_list_response"}}}}, "400": {"description": "Bad Request", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mconn_bad_response"}}}}, "401": {"description": "Unauthorized", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mconn_bad_response"}}}}, "403": {"description": "Forbidden", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mconn_bad_response"}}}}, "500": {"description": "Internal Server Error", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mconn_bad_response"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic Connectors"], "x-api-token-group": ["Magic WAN Write", "Magic WAN Read"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "connectors", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
