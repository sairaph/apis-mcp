---
title: Create Hyperdrive
page_id: operation-post-accounts-account-id-hyperdrive-configs-88ae6f82
path: operations/hyperdrive
description: Creates and returns a new Hyperdrive configuration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/hyperdrive/configs
operation_ids:
    - create-hyperdrive
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Hyperdrive

`POST /accounts/{account_id}/hyperdrive/configs`

Operation ID: `create-hyperdrive`

Creates and returns a new Hyperdrive configuration.

## Definition

```yaml
{"operationId": "create-hyperdrive", "summary": "Create Hyperdrive", "description": "Creates and returns a new Hyperdrive configuration.", "parameters": [{"name": "account_id", "in": "path", "description": "The Cloudflare account ID.", "required": true, "schema": {"$ref": "#/components/schemas/hyperdrive_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/hyperdrive_hyperdrive-config"}}}}, "responses": {"200": {"description": "Create Hyperdrive Response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/hyperdrive_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/hyperdrive_hyperdrive-config-response"}}, "type": "object"}]}}}}, "4XX": {"description": "Create Hyperdrive Failure Response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/hyperdrive_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Hyperdrive"], "x-api-token-group": ["Hyperdrive Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.edge.hyperdrive.database.create"]}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "hyperdrive", "x-fern-sdk-method-name": "create", "x-forge-hidden": true, "x-forge-params": {"mtls.mtls_certificate_id": {"flagName": "mtls-certificate-id"}, "mtls.sslmode": {"choices": ["require", "verify-ca", "verify-full"]}}}
```
