---
title: List Hyperdrives
page_id: operation-get-accounts-account-id-hyperdrive-configs-d5d5c54f
path: operations/hyperdrive
description: Returns a list of Hyperdrives.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/hyperdrive/configs
operation_ids:
    - list-hyperdrive
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Hyperdrives

`GET /accounts/{account_id}/hyperdrive/configs`

Operation ID: `list-hyperdrive`

Returns a list of Hyperdrives.

## Definition

```yaml
{"operationId": "list-hyperdrive", "summary": "List Hyperdrives", "description": "Returns a list of Hyperdrives.", "parameters": [{"name": "account_id", "in": "path", "description": "The Cloudflare account ID.", "required": true, "schema": {"$ref": "#/components/schemas/hyperdrive_identifier"}}, {"name": "page", "in": "query", "description": "Page number of paginated results.", "schema": {"type": "integer", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "description": "Maximum number of results per page.", "schema": {"type": "integer", "default": 20, "maximum": 100, "minimum": 1}}], "responses": {"200": {"description": "List Hyperdrives Response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/hyperdrive_api-response-collection"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/hyperdrive_hyperdrive-config-response"}}}, "type": "object"}]}}}}, "4XX": {"description": "List Hyperdrives Failure Response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/hyperdrive_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Hyperdrive"], "x-api-token-group": ["Hyperdrive Write", "Hyperdrive Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.edge.hyperdrive.database.list"]}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "hyperdrive", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
