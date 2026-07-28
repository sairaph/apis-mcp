---
title: Delete Hyperdrive
page_id: operation-delete-accounts-account-id-hyperdrive-configs-hyperdrive-id-302c1ea8
path: operations/hyperdrive
description: Deletes the specified Hyperdrive.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/hyperdrive/configs/{hyperdrive_id}
operation_ids:
    - delete-hyperdrive
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Hyperdrive

`DELETE /accounts/{account_id}/hyperdrive/configs/{hyperdrive_id}`

Operation ID: `delete-hyperdrive`

Deletes the specified Hyperdrive.

## Definition

```yaml
{"operationId": "delete-hyperdrive", "summary": "Delete Hyperdrive", "description": "Deletes the specified Hyperdrive.", "parameters": [{"name": "account_id", "in": "path", "description": "The Cloudflare account ID.", "required": true, "schema": {"$ref": "#/components/schemas/hyperdrive_identifier"}}, {"name": "hyperdrive_id", "in": "path", "description": "The unique identifier of the Hyperdrive configuration.", "required": true, "schema": {"$ref": "#/components/schemas/hyperdrive_identifier"}}], "responses": {"200": {"description": "Delete Hyperdrive Response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/hyperdrive_api-response-single"}, {"properties": {"result": {"type": "object", "nullable": true}}, "type": "object"}]}}}}, "4XX": {"description": "Delete Hyperdrive Failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/hyperdrive_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Hyperdrive"], "x-api-token-group": ["Hyperdrive Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.edge.hyperdrive.database.delete"]}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "hyperdrive", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
