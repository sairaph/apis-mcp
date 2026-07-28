---
title: Update Connector
page_id: operation-put-accounts-account-id-magic-connectors-connector-id-116a54b1
path: operations/magic-connectors
description: Updates properties of a Magic WAN Connector. May be used to re-provision a license key.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/magic/connectors/{connector_id}
operation_ids:
    - mconn-connectors-update
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Connector

`PUT /accounts/{account_id}/magic/connectors/{connector_id}`

Operation ID: `mconn-connectors-update`

Updates properties of a Magic WAN Connector. May be used to re-provision a license key.

## Definition

```yaml
{"operationId": "mconn-connectors-update", "summary": "Update Connector", "description": "Updates properties of a Magic WAN Connector. May be used to re-provision a license key.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mconn_account_id"}}, {"name": "connector_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mconn_uuid"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mconn_customer_connectors_update_request"}}}}, "responses": {"200": {"description": "OK", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mconn_customer_connectors_update_response"}}}}, "400": {"description": "Bad Request", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mconn_bad_response"}}}}, "401": {"description": "Unauthorized", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mconn_bad_response"}}}}, "403": {"description": "Forbidden", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mconn_bad_response"}}}}, "404": {"description": "Not Found", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mconn_bad_response"}}}}, "500": {"description": "Internal Server Error", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mconn_bad_response"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic Connectors"], "x-api-token-group": ["Magic WAN Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "connectors", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
