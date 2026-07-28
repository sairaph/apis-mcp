---
title: Validate ownership challenge
page_id: operation-post-zones-zone-id-logpush-ownership-validate-a826b9f7
path: operations/logpush-jobs-for-a-zone
description: Validates ownership challenge of the destination.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/logpush/ownership/validate
operation_ids:
    - post-zones-zone_id-logpush-ownership-validate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Validate ownership challenge

`POST /zones/{zone_id}/logpush/ownership/validate`

Operation ID: `post-zones-zone_id-logpush-ownership-validate`

Validates ownership challenge of the destination.

## Definition

```yaml
{"operationId": "post-zones-zone_id-logpush-ownership-validate", "summary": "Validate ownership challenge", "description": "Validates ownership challenge of the destination.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/logpush_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"destination_conf": {"$ref": "#/components/schemas/logpush_destination_conf"}, "ownership_challenge": {"$ref": "#/components/schemas/logpush_ownership_challenge"}}, "required": ["destination_conf", "ownership_challenge"]}}}}, "responses": {"200": {"description": "Validate ownership challenge response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/logpush_validate_ownership_response"}}}}, "4XX": {"description": "Validate ownership challenge response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/logpush_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Logpush jobs for a zone"], "x-api-token-group": ["Logs Write"], "x-cfPermissionsRequired": {"enum": ["#logs:edit"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "logpush.ownership", "x-fern-sdk-method-name": "validate", "x-forge-hidden": true}
```
