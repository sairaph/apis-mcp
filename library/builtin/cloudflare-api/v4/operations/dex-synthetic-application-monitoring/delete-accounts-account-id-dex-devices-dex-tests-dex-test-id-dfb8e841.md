---
title: Delete Device DEX test
page_id: operation-delete-accounts-account-id-dex-devices-dex-tests-dex-test-id-f6176f25
path: operations/dex-synthetic-application-monitoring
description: Delete a Device DEX test. Returns the remaining device dex tests for the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/dex/devices/dex_tests/{dex_test_id}
operation_ids:
    - device-dex-test-delete-device-dex-test
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Device DEX test

`DELETE /accounts/{account_id}/dex/devices/dex_tests/{dex_test_id}`

Operation ID: `device-dex-test-delete-device-dex-test`

Delete a Device DEX test. Returns the remaining device dex tests for the account.

## Definition

```yaml
{"operationId": "device-dex-test-delete-device-dex-test", "summary": "Delete Device DEX test", "description": "Delete a Device DEX test. Returns the remaining device dex tests for the account.", "parameters": [{"name": "account_id", "in": "path", "description": "Unique Cloudflare account ID.", "required": true, "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_account_identifier"}}, {"name": "dex_test_id", "in": "path", "description": "Unique identifier for a DEX test.", "required": true, "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_uuid"}}], "responses": {"200": {"description": "Delete Device DEX test response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/digital-experience-monitoring_dex-delete-response-collection"}}}}, "4XX": {"description": "Delete DEX test response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/digital-experience-monitoring_dex-response_collection"}, {"$ref": "#/components/schemas/digital-experience-monitoring_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["DEX Synthetic Application Monitoring"], "x-api-token-group": ["Cloudflare DEX Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.dex.tests", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
