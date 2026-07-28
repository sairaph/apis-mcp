---
title: List short-lived certificate CAs
page_id: operation-get-zones-zone-id-access-apps-ca-89493c75
path: operations/zone-level-access-short-lived-certificate-cas
description: Lists short-lived certificate CAs and their public keys.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/access/apps/ca
operation_ids:
    - zone-level-access-short-lived-certificate-c-as-list-short-lived-certificate-c-as
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List short-lived certificate CAs

`GET /zones/{zone_id}/access/apps/ca`

Operation ID: `zone-level-access-short-lived-certificate-c-as-list-short-lived-certificate-c-as`

Lists short-lived certificate CAs and their public keys.

## Definition

```yaml
{"operationId": "zone-level-access-short-lived-certificate-c-as-list-short-lived-certificate-c-as", "summary": "List short-lived certificate CAs", "description": "Lists short-lived certificate CAs and their public keys.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "responses": {"200": {"description": "List short-lived certificate CAs response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_response_collection-20"}}}}, "4XX": {"description": "List short-lived certificate CAs response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Zone-Level Access short-lived certificate CAs"], "x-api-token-group": ["Access: Apps and Policies Write", "Access: Apps and Policies Read", "Access: Apps and Policies Write", "Access: Apps and Policies Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.apps.ca", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
