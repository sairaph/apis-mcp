---
title: Patch Hyperdrive
page_id: operation-patch-accounts-account-id-hyperdrive-configs-hyperdrive-id-8aa7a40a
path: operations/hyperdrive
description: Patches and returns the specified Hyperdrive configuration. Custom caching settings are not kept if caching is disabled.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/hyperdrive/configs/{hyperdrive_id}
operation_ids:
    - patch-hyperdrive
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Patch Hyperdrive

`PATCH /accounts/{account_id}/hyperdrive/configs/{hyperdrive_id}`

Operation ID: `patch-hyperdrive`

Patches and returns the specified Hyperdrive configuration. Custom caching settings are not kept if caching is disabled.

## Definition

```yaml
{"operationId": "patch-hyperdrive", "summary": "Patch Hyperdrive", "description": "Patches and returns the specified Hyperdrive configuration. Custom caching settings are not kept if caching is disabled.", "parameters": [{"name": "account_id", "in": "path", "description": "The Cloudflare account ID.", "required": true, "schema": {"$ref": "#/components/schemas/hyperdrive_identifier"}}, {"name": "hyperdrive_id", "in": "path", "description": "The unique identifier of the Hyperdrive configuration.", "required": true, "schema": {"$ref": "#/components/schemas/hyperdrive_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/hyperdrive_hyperdrive-config-patch"}}}}, "responses": {"200": {"description": "Patch Hyperdrive Response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/hyperdrive_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/hyperdrive_hyperdrive-config-response"}}, "type": "object"}]}}}}, "4XX": {"description": "Patch Hyperdrive Failure Response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/hyperdrive_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Hyperdrive"], "x-api-token-group": ["Hyperdrive Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.edge.hyperdrive.database.update"]}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "hyperdrive", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true, "x-forge-params": {"mtls.mtls_certificate_id": {"flagName": "mtls-certificate-id"}, "mtls.sslmode": {"choices": ["require", "verify-ca", "verify-full"]}}}
```
