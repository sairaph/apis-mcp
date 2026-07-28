---
title: Create Device DEX test
page_id: operation-post-accounts-account-id-dex-devices-dex-tests-c6dee57a
path: operations/dex-synthetic-application-monitoring
description: Create a DEX test.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/dex/devices/dex_tests
operation_ids:
    - device-dex-test-create-device-dex-test
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Device DEX test

`POST /accounts/{account_id}/dex/devices/dex_tests`

Operation ID: `device-dex-test-create-device-dex-test`

Create a DEX test.

## Definition

```yaml
{"operationId": "device-dex-test-create-device-dex-test", "summary": "Create Device DEX test", "description": "Create a DEX test.", "parameters": [{"name": "account_id", "in": "path", "description": "Unique Cloudflare account ID.", "required": true, "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/digital-experience-monitoring_device-dex-test-schemas-http"}}}}, "responses": {"200": {"description": "Update Dex test response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/digital-experience-monitoring_dex-single_response"}}}}, "4XX": {"description": "Update Dex test response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/digital-experience-monitoring_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["DEX Synthetic Application Monitoring"], "x-api-token-group": ["Cloudflare DEX Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.dex.tests", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
