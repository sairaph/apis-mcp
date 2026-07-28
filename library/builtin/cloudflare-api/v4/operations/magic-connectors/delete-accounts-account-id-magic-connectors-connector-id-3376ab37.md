---
title: Delete Connector
page_id: operation-delete-accounts-account-id-magic-connectors-connector-id-9ec54e67
path: operations/magic-connectors
description: Deletes a Magic WAN Connector.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/magic/connectors/{connector_id}
operation_ids:
    - mconn-connectors-delete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Connector

`DELETE /accounts/{account_id}/magic/connectors/{connector_id}`

Operation ID: `mconn-connectors-delete`

Deletes a Magic WAN Connector.

## Definition

```yaml
{"operationId": "mconn-connectors-delete", "summary": "Delete Connector", "description": "Deletes a Magic WAN Connector.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mconn_account_id"}}, {"name": "connector_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mconn_uuid"}}], "responses": {"200": {"description": "OK", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mconn_customer_connectors_delete_response"}}}}, "400": {"description": "Bad Request", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mconn_bad_response"}}}}, "401": {"description": "Unauthorized", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mconn_bad_response"}}}}, "403": {"description": "Forbidden", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mconn_bad_response"}}}}, "404": {"description": "Not Found", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mconn_bad_response"}}}}, "500": {"description": "Internal Server Error", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mconn_bad_response"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic Connectors"], "x-api-token-group": ["Magic WAN Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "connectors", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
