---
title: List Access Applications
page_id: operation-get-zones-zone-id-access-apps-09670252
path: operations/zone-level-access-applications
description: List all Access Applications in a zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/access/apps
operation_ids:
    - zone-level-access-applications-list-access-applications
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Access Applications

`GET /zones/{zone_id}/access/apps`

Operation ID: `zone-level-access-applications-list-access-applications`

List all Access Applications in a zone.

## Definition

```yaml
{"operationId": "zone-level-access-applications-list-access-applications", "summary": "List Access Applications", "description": "List all Access Applications in a zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "responses": {"200": {"description": "List Access Applications response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_response_collection-22"}}}}, "4XX": {"description": "List Access Applications response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Zone-Level Access applications"], "x-api-token-group": ["Access: Apps and Policies Revoke", "Access: Apps and Policies Write", "Access: Apps and Policies Read", "Access: Apps and Policies Revoke", "Access: Apps and Policies Write", "Access: Apps and Policies Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.apps", "x-fern-sdk-method-name": "list", "x-forge-hidden": true, "x-stability": "beta"}
```
