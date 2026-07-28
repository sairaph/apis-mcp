---
title: Delete an Access policy
page_id: operation-delete-zones-zone-id-access-apps-app-id-policies-policy-id-f2229c2a
path: operations/zone-level-access-policies
description: Delete an Access policy.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/access/apps/{app_id}/policies/{policy_id}
operation_ids:
    - zone-level-access-policies-delete-an-access-policy
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete an Access policy

`DELETE /zones/{zone_id}/access/apps/{app_id}/policies/{policy_id}`

Operation ID: `zone-level-access-policies-delete-an-access-policy`

Delete an Access policy.

## Definition

```yaml
{"operationId": "zone-level-access-policies-delete-an-access-policy", "summary": "Delete an Access policy", "description": "Delete an Access policy.", "parameters": [{"name": "policy_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}, {"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "responses": {"202": {"description": "Delete an Access policy response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_id_response"}}}}, "4XX": {"description": "Delete an Access policy response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Zone-Level Access policies"], "x-api-token-group": ["Access: Apps and Policies Write", "Access: Apps and Policies Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.apps.policies", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
