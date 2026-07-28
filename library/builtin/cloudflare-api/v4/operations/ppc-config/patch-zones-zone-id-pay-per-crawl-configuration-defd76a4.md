---
title: Changes pay-per-crawl config for a zone
page_id: operation-patch-zones-zone-id-pay-per-crawl-configuration-de817d6e
path: operations/ppc-config
description: Changes the pay-per-crawl config for a zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/pay-per-crawl/configuration
operation_ids:
    - pay-per-crawl.patchConfig
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Changes pay-per-crawl config for a zone

`PATCH /zones/{zone_id}/pay-per-crawl/configuration`

Operation ID: `pay-per-crawl.patchConfig`

Changes the pay-per-crawl config for a zone.

## Definition

```yaml
{"operationId": "pay-per-crawl.patchConfig", "summary": "Changes pay-per-crawl config for a zone", "description": "Changes the pay-per-crawl config for a zone.", "parameters": [{"name": "zone_id", "in": "path", "description": "zone id", "required": true, "schema": {"type": "string"}, "x-auditable": true}], "requestBody": {"$ref": "#/components/requestBodies/pay-per-crawl_DaricConfig"}, "responses": {"200": {"description": "OK", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/pay-per-crawl_getConfigResponse"}}}}, "400": {"description": "Bad Request", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/pay-per-crawl_apiErrorResponse"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["ppc_config"], "x-api-token-group": ["Zone Settings Write"]}
```
