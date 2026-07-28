---
title: List CF1 Site Ramps
page_id: operation-get-accounts-account-id-magic-cf1-sites-cf1-site-id-ramps-aa219f72
path: operations/magic-cf1-site-ramps
description: Lists ramps (network connections) associated with a CF1 Site. Ramps represent GRE tunnels, IPsec tunnels, interconnects, or MCONN links.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/magic/cf1_sites/{cf1_site_id}/ramps
operation_ids:
    - magic-cf1-sites-list-cf1-site-ramps
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List CF1 Site Ramps

`GET /accounts/{account_id}/magic/cf1_sites/{cf1_site_id}/ramps`

Operation ID: `magic-cf1-sites-list-cf1-site-ramps`

Lists ramps (network connections) associated with a CF1 Site. Ramps represent GRE tunnels, IPsec tunnels, interconnects, or MCONN links.

## Definition

```yaml
{"operationId": "magic-cf1-sites-list-cf1-site-ramps", "summary": "List CF1 Site Ramps", "description": "Lists ramps (network connections) associated with a CF1 Site. Ramps represent GRE tunnels, IPsec tunnels, interconnects, or MCONN links.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "cf1_site_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "responses": {"200": {"description": "List CF1 Site Ramps response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_cf1_site_ramps_collection_response"}}}}, "4XX": {"description": "List CF1 Site Ramps response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic CF1 Site Ramps"], "x-api-token-group": ["Magic WAN Write", "Magic WAN Read", "Magic Transit Read", "Magic Transit Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}}
```
