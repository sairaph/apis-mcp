---
title: Delete CF1 Site Ramp
page_id: operation-delete-accounts-account-id-magic-cf1-sites-cf1-site-id-ramps-ramp-id-bacd4fe2
path: operations/magic-cf1-site-ramps
description: Deletes a specific ramp from a CF1 Site.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/magic/cf1_sites/{cf1_site_id}/ramps/{ramp_id}
operation_ids:
    - magic-cf1-sites-delete-cf1-site-ramp
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete CF1 Site Ramp

`DELETE /accounts/{account_id}/magic/cf1_sites/{cf1_site_id}/ramps/{ramp_id}`

Operation ID: `magic-cf1-sites-delete-cf1-site-ramp`

Deletes a specific ramp from a CF1 Site.

## Definition

```yaml
{"operationId": "magic-cf1-sites-delete-cf1-site-ramp", "summary": "Delete CF1 Site Ramp", "description": "Deletes a specific ramp from a CF1 Site.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "cf1_site_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "ramp_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "responses": {"200": {"description": "Delete CF1 Site Ramp response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_cf1_site_ramp_single_response"}}}}, "4XX": {"description": "Delete CF1 Site Ramp response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic CF1 Site Ramps"], "x-api-token-group": ["Magic WAN Write", "Magic Transit Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}}
```
