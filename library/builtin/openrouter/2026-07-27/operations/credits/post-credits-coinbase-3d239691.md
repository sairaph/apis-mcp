---
title: Deprecated Coinbase Commerce charge endpoint
page_id: operation-post-credits-coinbase-7e373772
path: operations/credits
description: Deprecated. The Coinbase APIs used by this endpoint have been deprecated, so Coinbase Commerce charges have been removed. Use the web credits purchase flow instead.
source: https://openrouter.ai/openapi.json
http_methods:
    - POST
api_endpoints:
    - /credits/coinbase
operation_ids:
    - createCoinbaseCharge
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Deprecated Coinbase Commerce charge endpoint

`POST /credits/coinbase`

Operation ID: `createCoinbaseCharge`

Deprecated. The Coinbase APIs used by this endpoint have been deprecated, so Coinbase Commerce charges have been removed. Use the web credits purchase flow instead.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"deprecated": true, "description": "Deprecated. The Coinbase APIs used by this endpoint have been deprecated, so Coinbase Commerce charges have been removed. Use the web credits purchase flow instead.", "operationId": "createCoinbaseCharge", "responses": {"200": {"description": "This endpoint is deprecated and will never return a 200 response."}, "410": {"content": {"application/json": {"example": {"error": {"code": 410, "message": "The Coinbase APIs used by this endpoint have been deprecated, so the Coinbase Commerce credits API has been removed. Use the web credits purchase flow instead."}}, "schema": {"$ref": "#/components/schemas/GoneResponse"}}}, "description": "Gone - Endpoint has been permanently removed or deprecated"}}, "security": [], "summary": "Deprecated Coinbase Commerce charge endpoint", "tags": ["Credits"], "x-fern-ignore": true, "x-speakeasy-ignore": true, "x-speakeasy-name-override": "createCoinbaseCharge"}
```
