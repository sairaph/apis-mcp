---
title: Get Hyperdrive
page_id: operation-get-accounts-account-id-hyperdrive-configs-hyperdrive-id-67a122af
path: operations/hyperdrive
description: Returns the specified Hyperdrive configuration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/hyperdrive/configs/{hyperdrive_id}
operation_ids:
    - get-hyperdrive
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Hyperdrive

`GET /accounts/{account_id}/hyperdrive/configs/{hyperdrive_id}`

Operation ID: `get-hyperdrive`

Returns the specified Hyperdrive configuration.

## Definition

```yaml
{"operationId": "get-hyperdrive", "summary": "Get Hyperdrive", "description": "Returns the specified Hyperdrive configuration.", "parameters": [{"name": "account_id", "in": "path", "description": "The Cloudflare account ID.", "required": true, "schema": {"$ref": "#/components/schemas/hyperdrive_identifier"}}, {"name": "hyperdrive_id", "in": "path", "description": "The unique identifier of the Hyperdrive configuration.", "required": true, "schema": {"$ref": "#/components/schemas/hyperdrive_identifier"}}], "responses": {"200": {"description": "Get Hyperdrive Response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/hyperdrive_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/hyperdrive_hyperdrive-config-response"}}, "type": "object"}]}}}}, "4XX": {"description": "Get Hyperdrive Failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/hyperdrive_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Hyperdrive"], "x-api-token-group": ["Hyperdrive Write", "Hyperdrive Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.edge.hyperdrive.database.read"]}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "hyperdrive", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
