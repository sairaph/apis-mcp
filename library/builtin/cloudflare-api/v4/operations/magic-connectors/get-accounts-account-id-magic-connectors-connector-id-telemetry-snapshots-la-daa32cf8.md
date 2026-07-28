---
title: Get latest Snapshots
page_id: operation-get-accounts-account-id-magic-connectors-connector-id-telemetry-snapshot-cc00a7be
path: operations/magic-connectors
description: Gets latest Magic WAN Connector Telemetry Snapshots
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/magic/connectors/{connector_id}/telemetry/snapshots/latest
operation_ids:
    - mconn-connector-telemetry-snapshots-latest-get
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get latest Snapshots

`GET /accounts/{account_id}/magic/connectors/{connector_id}/telemetry/snapshots/latest`

Operation ID: `mconn-connector-telemetry-snapshots-latest-get`

Gets latest Magic WAN Connector Telemetry Snapshots

## Definition

```yaml
{"operationId": "mconn-connector-telemetry-snapshots-latest-get", "summary": "Get latest Snapshots", "description": "Gets latest Magic WAN Connector Telemetry Snapshots", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mconn_account_id"}}, {"name": "connector_id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "OK", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mconn_customer_snapshots_latest_get_success"}}}}, "400": {"description": "Bad Request", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mconn_envelope"}}}}, "401": {"description": "Unauthorized", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mconn_envelope"}}}}, "403": {"description": "Forbidden", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mconn_envelope"}}}}, "404": {"description": "Not Found", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mconn_envelope"}}}}, "429": {"description": "Too Many Requests", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mconn_envelope"}}}}, "500": {"description": "Internal Server Error", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mconn_envelope"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic Connectors"], "x-api-token-group": ["Magic WAN Write", "Magic WAN Read"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-ignore": true, "x-forge-hidden": true}
```
