---
title: Restart Hyperdrive
page_id: operation-post-accounts-account-id-hyperdrive-configs-hyperdrive-id-restart-5d07683f
path: operations/hyperdrive
description: Restarts the connection pool for the specified Hyperdrive configuration without changing its configuration. Existing connections are drained and a new pool is established at the edge.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/hyperdrive/configs/{hyperdrive_id}/restart
operation_ids:
    - restart-hyperdrive
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Restart Hyperdrive

`POST /accounts/{account_id}/hyperdrive/configs/{hyperdrive_id}/restart`

Operation ID: `restart-hyperdrive`

Restarts the connection pool for the specified Hyperdrive configuration without changing its configuration. Existing connections are drained and a new pool is established at the edge.

## Definition

```yaml
{"operationId": "restart-hyperdrive", "summary": "Restart Hyperdrive", "description": "Restarts the connection pool for the specified Hyperdrive configuration without changing its configuration. Existing connections are drained and a new pool is established at the edge.", "parameters": [{"name": "account_id", "in": "path", "description": "The Cloudflare account ID.", "required": true, "schema": {"$ref": "#/components/schemas/hyperdrive_identifier"}}, {"name": "hyperdrive_id", "in": "path", "description": "The unique identifier of the Hyperdrive configuration.", "required": true, "schema": {"$ref": "#/components/schemas/hyperdrive_identifier"}}], "responses": {"200": {"description": "Restart Hyperdrive Response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/hyperdrive_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/hyperdrive_hyperdrive-config-response"}}, "type": "object"}]}}}}, "4XX": {"description": "Restart Hyperdrive Failure Response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/hyperdrive_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Hyperdrive"], "x-api-token-group": ["Hyperdrive Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "hyperdrive", "x-fern-sdk-method-name": "restart", "x-forge-hidden": true}
```
