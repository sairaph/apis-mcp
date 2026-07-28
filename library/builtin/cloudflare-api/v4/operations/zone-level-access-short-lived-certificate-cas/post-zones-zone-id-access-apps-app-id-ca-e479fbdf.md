---
title: Create a short-lived certificate CA
page_id: operation-post-zones-zone-id-access-apps-app-id-ca-96fb2e74
path: operations/zone-level-access-short-lived-certificate-cas
description: Generates a new short-lived certificate CA and public key.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/access/apps/{app_id}/ca
operation_ids:
    - zone-level-access-short-lived-certificate-c-as-create-a-short-lived-certificate-ca
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a short-lived certificate CA

`POST /zones/{zone_id}/access/apps/{app_id}/ca`

Operation ID: `zone-level-access-short-lived-certificate-c-as-create-a-short-lived-certificate-ca`

Generates a new short-lived certificate CA and public key.

## Definition

```yaml
{"operationId": "zone-level-access-short-lived-certificate-c-as-create-a-short-lived-certificate-ca", "summary": "Create a short-lived certificate CA", "description": "Generates a new short-lived certificate CA and public key.", "parameters": [{"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "responses": {"200": {"description": "Create a short-lived certificate CA response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_single_response-19"}}}}, "4XX": {"description": "Create a short-lived certificate CA response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Zone-Level Access short-lived certificate CAs"], "x-api-token-group": ["Access: Apps and Policies Write", "Access: Apps and Policies Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.apps.ca", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
