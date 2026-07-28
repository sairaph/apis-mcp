---
title: Update an mTLS certificate's hostname settings
page_id: operation-put-zones-zone-id-access-certificates-settings-897870da
path: operations/zone-level-access-mtls-authentication
description: Updates an mTLS certificate's hostname settings.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/access/certificates/settings
operation_ids:
    - zone-level-access-mtls-authentication-update-an-mtls-certificate-settings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update an mTLS certificate's hostname settings

`PUT /zones/{zone_id}/access/certificates/settings`

Operation ID: `zone-level-access-mtls-authentication-update-an-mtls-certificate-settings`

Updates an mTLS certificate's hostname settings.

## Definition

```yaml
{"operationId": "zone-level-access-mtls-authentication-update-an-mtls-certificate-settings", "summary": "Update an mTLS certificate's hostname settings", "description": "Updates an mTLS certificate's hostname settings.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"settings": {"type": "array", "items": {"$ref": "#/components/schemas/access_settings-2"}}}, "required": ["settings"]}}}}, "responses": {"202": {"description": "Update an mTLS certificates hostname settings response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_response_collection_hostnames-2"}}}}, "4XX": {"description": "Update an mTLS certificates hostname settings failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Zone-Level Access mTLS authentication"], "x-api-token-group": ["Access: Mutual TLS Certificates Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.certificates.settings", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
