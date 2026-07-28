---
title: Create Interrupt
page_id: operation-post-accounts-account-id-magic-connectors-connector-id-interrupts-37cf5085
path: operations/magic-connectors
description: Creates an interrupt for a Magic WAN Connector.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/magic/connectors/{connector_id}/interrupts
operation_ids:
    - mconn-connector-interrupts-create
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Interrupt

`POST /accounts/{account_id}/magic/connectors/{connector_id}/interrupts`

Operation ID: `mconn-connector-interrupts-create`

Creates an interrupt for a Magic WAN Connector.

## Definition

```yaml
{"operationId": "mconn-connector-interrupts-create", "summary": "Create Interrupt", "description": "Creates an interrupt for a Magic WAN Connector.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mconn_account_id"}}, {"name": "connector_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mconn_uuid"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mconn_interrupt"}}}}, "responses": {"200": {"description": "OK", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mconn_connector_interrupts_create_response"}}}}, "400": {"description": "Bad Request", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mconn_bad_response"}}}}, "401": {"description": "Unauthorized", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mconn_bad_response"}}}}, "403": {"description": "Forbidden", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mconn_bad_response"}}}}, "404": {"description": "Not Found", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mconn_bad_response"}}}}, "409": {"description": "Conflict", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mconn_bad_response"}}}}, "500": {"description": "Internal Server Error", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mconn_bad_response"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic Connectors"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "connector-interrupts", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
