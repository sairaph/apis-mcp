---
title: Create an Access policy
page_id: operation-post-zones-zone-id-access-apps-app-id-policies-be3b8724
path: operations/zone-level-access-policies
description: Create a new Access policy for an application.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/access/apps/{app_id}/policies
operation_ids:
    - zone-level-access-policies-create-an-access-policy
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create an Access policy

`POST /zones/{zone_id}/access/apps/{app_id}/policies`

Operation ID: `zone-level-access-policies-create-an-access-policy`

Create a new Access policy for an application.

## Definition

```yaml
{"operationId": "zone-level-access-policies-create-an-access-policy", "summary": "Create an Access policy", "description": "Create a new Access policy for an application.", "parameters": [{"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"properties": {"approval_groups": {"$ref": "#/components/schemas/access_approval_groups-2"}, "approval_required": {"$ref": "#/components/schemas/access_approval_required-2"}, "decision": {"$ref": "#/components/schemas/access_decision-2"}, "exclude": {"$ref": "#/components/schemas/access_exclude-3"}, "include": {"$ref": "#/components/schemas/access_include"}, "isolation_required": {"$ref": "#/components/schemas/access_isolation_required-2"}, "name": {"$ref": "#/components/schemas/access_name-9"}, "precedence": {"$ref": "#/components/schemas/access_precedence-2"}, "purpose_justification_prompt": {"$ref": "#/components/schemas/access_purpose_justification_prompt"}, "purpose_justification_required": {"$ref": "#/components/schemas/access_purpose_justification_required-2"}, "require": {"$ref": "#/components/schemas/access_require-3"}}, "required": ["name", "decision", "include"]}}}}, "responses": {"201": {"description": "Create an Access policy response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_single_response-22"}}}}, "4XX": {"description": "Create an Access policy response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Zone-Level Access policies"], "x-api-token-group": ["Access: Apps and Policies Write", "Access: Apps and Policies Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.apps.policies", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
