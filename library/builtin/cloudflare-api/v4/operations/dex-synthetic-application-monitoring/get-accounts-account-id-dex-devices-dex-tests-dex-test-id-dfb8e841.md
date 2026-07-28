---
title: Get Device DEX test
page_id: operation-get-accounts-account-id-dex-devices-dex-tests-dex-test-id-7ded049c
path: operations/dex-synthetic-application-monitoring
description: Fetch a single DEX test.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dex/devices/dex_tests/{dex_test_id}
operation_ids:
    - device-dex-test-get-device-dex-test
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Device DEX test

`GET /accounts/{account_id}/dex/devices/dex_tests/{dex_test_id}`

Operation ID: `device-dex-test-get-device-dex-test`

Fetch a single DEX test.

## Definition

```yaml
{"operationId": "device-dex-test-get-device-dex-test", "summary": "Get Device DEX test", "description": "Fetch a single DEX test.", "parameters": [{"name": "account_id", "in": "path", "description": "Unique Cloudflare account ID.", "required": true, "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_account_identifier"}}, {"name": "dex_test_id", "in": "path", "description": "Unique identifier for a DEX test.", "required": true, "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_schemas-test-id"}}], "responses": {"200": {"description": "Device DEX test details response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/digital-experience-monitoring_dex-single_response"}}}}, "4XX": {"description": "Device DEX test response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/digital-experience-monitoring_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["DEX Synthetic Application Monitoring"], "x-api-token-group": ["Cloudflare DEX Write", "Cloudflare DEX Read", "Zero Trust Report", "Zero Trust Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.dex.tests", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
