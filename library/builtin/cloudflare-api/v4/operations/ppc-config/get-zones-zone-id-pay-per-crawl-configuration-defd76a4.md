---
title: Get the pay-per-crawl config
page_id: operation-get-zones-zone-id-pay-per-crawl-configuration-f9400b89
path: operations/ppc-config
description: Gets the pay-per-crawl config for a zone including the bot configuration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/pay-per-crawl/configuration
operation_ids:
    - pay-per-crawl.getConfig
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get the pay-per-crawl config

`GET /zones/{zone_id}/pay-per-crawl/configuration`

Operation ID: `pay-per-crawl.getConfig`

Gets the pay-per-crawl config for a zone including the bot configuration.

## Definition

```yaml
{"operationId": "pay-per-crawl.getConfig", "summary": "Get the pay-per-crawl config", "description": "Gets the pay-per-crawl config for a zone including the bot configuration.", "parameters": [{"name": "zone_id", "in": "path", "description": "zone id", "required": true, "schema": {"type": "string"}, "x-auditable": true}], "responses": {"200": {"description": "OK", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/pay-per-crawl_getConfigResponse"}}}}, "400": {"description": "Bad Request", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/pay-per-crawl_apiErrorResponse"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["ppc_config"], "x-api-token-group": ["Zone Settings Write", "Zone Settings Read"]}
```
