---
title: List Events
page_id: operation-get-accounts-account-id-magic-connectors-connector-id-telemetry-events-2f3a0e90
path: operations/magic-connectors
description: Lists Magic WAN Connector Telemetry Events
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/magic/connectors/{connector_id}/telemetry/events
operation_ids:
    - mconn-connector-telemetry-events-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Events

`GET /accounts/{account_id}/magic/connectors/{connector_id}/telemetry/events`

Operation ID: `mconn-connector-telemetry-events-list`

Lists Magic WAN Connector Telemetry Events

## Definition

```yaml
{"operationId": "mconn-connector-telemetry-events-list", "summary": "List Events", "description": "Lists Magic WAN Connector Telemetry Events", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mconn_account_id"}}, {"name": "connector_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "from", "in": "query", "required": true, "schema": {"type": "number"}}, {"name": "to", "in": "query", "required": true, "schema": {"type": "number"}}, {"name": "limit", "in": "query", "schema": {"type": "number"}}, {"name": "cursor", "in": "query", "schema": {"type": "string"}}, {"name": "k", "in": "query", "description": "Filter by event kind", "schema": {"type": "string"}}], "responses": {"200": {"description": "OK", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mconn_customer_events_list_success"}}}}, "400": {"description": "Bad Request", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mconn_envelope"}}}}, "401": {"description": "Unauthorized", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mconn_envelope"}}}}, "403": {"description": "Forbidden", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mconn_envelope"}}}}, "429": {"description": "Too Many Requests", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mconn_envelope"}}}}, "500": {"description": "Internal Server Error", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mconn_envelope"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic Connectors"], "x-api-token-group": ["Magic WAN Write", "Magic WAN Read"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "connector-telemetry-events", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
