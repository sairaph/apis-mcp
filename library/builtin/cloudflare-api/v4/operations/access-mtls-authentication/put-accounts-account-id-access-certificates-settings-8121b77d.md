---
title: Update an mTLS certificate's hostname settings
page_id: operation-put-accounts-account-id-access-certificates-settings-793bfc25
path: operations/access-mtls-authentication
description: Updates an mTLS certificate's hostname settings.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/access/certificates/settings
operation_ids:
    - access-mtls-authentication-update-an-mtls-certificate-settings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update an mTLS certificate's hostname settings

`PUT /accounts/{account_id}/access/certificates/settings`

Operation ID: `access-mtls-authentication-update-an-mtls-certificate-settings`

Updates an mTLS certificate's hostname settings.

## Definition

```yaml
{"operationId": "access-mtls-authentication-update-an-mtls-certificate-settings", "summary": "Update an mTLS certificate's hostname settings", "description": "Updates an mTLS certificate's hostname settings.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"settings": {"type": "array", "items": {"$ref": "#/components/schemas/access_settings"}}}, "required": ["settings"]}}}}, "responses": {"202": {"description": "Update an mTLS certificates hostname settings response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_response_collection_hostnames"}}}}, "4XX": {"description": "Update an mTLS certificates hostname settings failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access mTLS authentication"], "x-api-token-group": ["Access: Mutual TLS Certificates Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.certificates.settings", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
