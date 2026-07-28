---
title: Create CF1 Site Ramps
page_id: operation-post-accounts-account-id-magic-cf1-sites-cf1-site-id-ramps-1f357fea
path: operations/magic-cf1-site-ramps
description: Creates ramps (network connections) for a CF1 Site.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/magic/cf1_sites/{cf1_site_id}/ramps
operation_ids:
    - magic-cf1-sites-create-cf1-site-ramps
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create CF1 Site Ramps

`POST /accounts/{account_id}/magic/cf1_sites/{cf1_site_id}/ramps`

Operation ID: `magic-cf1-sites-create-cf1-site-ramps`

Creates ramps (network connections) for a CF1 Site.

## Definition

```yaml
{"operationId": "magic-cf1-sites-create-cf1-site-ramps", "summary": "Create CF1 Site Ramps", "description": "Creates ramps (network connections) for a CF1 Site.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "cf1_site_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "array", "items": {"$ref": "#/components/schemas/magic_cf1_site_ramp_body"}, "minItems": 1}}}}, "responses": {"200": {"description": "Create CF1 Site Ramps response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_cf1_site_ramps_collection_response"}}}}, "4XX": {"description": "Create CF1 Site Ramps response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic CF1 Site Ramps"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}}
```
