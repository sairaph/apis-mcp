---
title: Set can_be_enabled setting on zones
page_id: operation-patch-accounts-account-id-pay-per-crawl-zones-can-be-enabled-139776d9
path: operations/ppc-config
description: Allows an account admin to set the can_be_enabled setting on a list of zones.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/pay-per-crawl/zones_can_be_enabled
operation_ids:
    - pay-per-crawl.setZonesCanBeEnabled
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Set can_be_enabled setting on zones

`PATCH /accounts/{account_id}/pay-per-crawl/zones_can_be_enabled`

Operation ID: `pay-per-crawl.setZonesCanBeEnabled`

Allows an account admin to set the can_be_enabled setting on a list of zones.

## Definition

```yaml
{"operationId": "pay-per-crawl.setZonesCanBeEnabled", "summary": "Set can_be_enabled setting on zones", "description": "Allows an account admin to set the can_be_enabled setting on a list of zones.", "parameters": [{"name": "account_id", "in": "path", "description": "account id", "required": true, "schema": {"type": "string"}, "x-auditable": true}], "requestBody": {"$ref": "#/components/requestBodies/pay-per-crawl_ZonesCanBeEnabledPayload"}, "responses": {"200": {"description": "OK", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/pay-per-crawl_apiNoResultResponse"}}}}, "400": {"description": "Bad Request", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/pay-per-crawl_apiErrorResponse"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["ppc_config"], "x-api-token-group": ["Account Settings Write"]}
```
