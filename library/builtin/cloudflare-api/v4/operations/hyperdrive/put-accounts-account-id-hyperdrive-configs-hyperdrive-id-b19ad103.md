---
title: Update Hyperdrive
page_id: operation-put-accounts-account-id-hyperdrive-configs-hyperdrive-id-bc870fe9
path: operations/hyperdrive
description: Updates and returns the specified Hyperdrive configuration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/hyperdrive/configs/{hyperdrive_id}
operation_ids:
    - update-hyperdrive
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Hyperdrive

`PUT /accounts/{account_id}/hyperdrive/configs/{hyperdrive_id}`

Operation ID: `update-hyperdrive`

Updates and returns the specified Hyperdrive configuration.

## Definition

```yaml
{"operationId": "update-hyperdrive", "summary": "Update Hyperdrive", "description": "Updates and returns the specified Hyperdrive configuration.", "parameters": [{"name": "account_id", "in": "path", "description": "The Cloudflare account ID.", "required": true, "schema": {"$ref": "#/components/schemas/hyperdrive_identifier"}}, {"name": "hyperdrive_id", "in": "path", "description": "The unique identifier of the Hyperdrive configuration.", "required": true, "schema": {"$ref": "#/components/schemas/hyperdrive_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/hyperdrive_hyperdrive-config"}}}}, "responses": {"200": {"description": "Update Hyperdrive Response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/hyperdrive_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/hyperdrive_hyperdrive-config-response"}}, "type": "object"}]}}}}, "4XX": {"description": "Update Hyperdrive Failure Response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/hyperdrive_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Hyperdrive"], "x-api-token-group": ["Hyperdrive Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.edge.hyperdrive.database.update"]}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "hyperdrive", "x-fern-sdk-method-name": "update", "x-forge-hidden": true, "x-forge-params": {"mtls.mtls_certificate_id": {"flagName": "mtls-certificate-id"}, "mtls.sslmode": {"choices": ["require", "verify-ca", "verify-full"]}}}
```
