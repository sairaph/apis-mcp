---
title: List Device DEX tests
page_id: operation-get-accounts-account-id-dex-devices-dex-tests-ea899c45
path: operations/dex-synthetic-application-monitoring
description: Fetch all DEX tests.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dex/devices/dex_tests
operation_ids:
    - device-dex-test-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Device DEX tests

`GET /accounts/{account_id}/dex/devices/dex_tests`

Operation ID: `device-dex-test-details`

Fetch all DEX tests.

## Definition

```yaml
{"operationId": "device-dex-test-details", "summary": "List Device DEX tests", "description": "Fetch all DEX tests.", "parameters": [{"name": "account_id", "in": "path", "description": "Unique Cloudflare account ID.", "required": true, "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_account_identifier"}}, {"name": "page", "in": "query", "description": "Page number of paginated results.", "schema": {"type": "number", "example": 1, "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "description": "Number of results per page.", "schema": {"type": "number", "example": 10, "default": 10, "maximum": 50, "minimum": 1}}, {"name": "testName", "in": "query", "description": "Filter by test name.", "schema": {"type": "string"}}, {"name": "kind", "in": "query", "description": "Filter by test type.", "schema": {"type": "string", "enum": ["http", "traceroute"]}}], "responses": {"200": {"description": "Device DEX test details response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/digital-experience-monitoring_dex-response_collection"}}}}, "4XX": {"description": "Device DEX test response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/digital-experience-monitoring_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["DEX Synthetic Application Monitoring"], "x-api-token-group": ["Cloudflare DEX Write", "Cloudflare DEX Read", "Zero Trust Report", "Zero Trust Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.dex.tests", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
