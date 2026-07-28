---
title: Traceroute
page_id: operation-post-accounts-account-id-diagnostics-traceroute-a603a496
path: operations/diagnostics
description: Run traceroutes from Cloudflare colos.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/diagnostics/traceroute
operation_ids:
    - diagnostics-traceroute
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Traceroute

`POST /accounts/{account_id}/diagnostics/traceroute`

Operation ID: `diagnostics-traceroute`

Run traceroutes from Cloudflare colos.

## Definition

```yaml
{"operationId": "diagnostics-traceroute", "summary": "Traceroute", "description": "Run traceroutes from Cloudflare colos.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic-transit_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"properties": {"colos": {"$ref": "#/components/schemas/magic-transit_colos"}, "options": {"$ref": "#/components/schemas/magic-transit_options"}, "targets": {"$ref": "#/components/schemas/magic-transit_targets"}}, "required": ["targets"]}}}}, "responses": {"200": {"description": "Traceroute response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic-transit_traceroute_response_collection"}}}}, "4XX": {"description": "Traceroute response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/magic-transit_traceroute_response_collection"}, {"$ref": "#/components/schemas/magic-transit_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Diagnostics"], "x-api-token-group": ["Magic Transit Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}}
```
