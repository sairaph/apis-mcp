---
title: List DEX test analytics
page_id: operation-get-accounts-account-id-dex-tests-overview-e3950bf9
path: operations/dex-synthetic-application-monitoring
description: List DEX tests with overview metrics.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dex/tests/overview
operation_ids:
    - dex-endpoints-list-tests-overview
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List DEX test analytics

`GET /accounts/{account_id}/dex/tests/overview`

Operation ID: `dex-endpoints-list-tests-overview`

List DEX tests with overview metrics.

## Definition

```yaml
{"operationId": "dex-endpoints-list-tests-overview", "summary": "List DEX test analytics", "description": "List DEX tests with overview metrics.", "parameters": [{"name": "account_id", "in": "path", "description": "Unique Cloudflare account ID.", "required": true, "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_account_identifier"}}, {"name": "colo", "in": "query", "description": "Optionally filter result stats to a Cloudflare colo. Cannot be used in combination with deviceId param.", "schema": {"description": "Cloudflare colo airport code.", "type": "string", "example": "SJC"}}, {"name": "testName", "in": "query", "description": "Optionally filter results by test name.", "schema": {"type": "string"}}, {"name": "deviceId", "in": "query", "description": "Optionally filter result stats to a specific device(s). Cannot be used in combination with colo param.", "schema": {"type": "array", "items": {"description": "Unique identifier for the physical device (UUID).", "example": "cb49c27f-7f97-49c5-b6f3-f7c01ead0fd7", "type": "string"}}}, {"name": "registration_id", "in": "query", "description": "Optionally filter results to a specific device registration. Must be used in combination with a single deviceId.", "schema": {"description": "Unique identifier for the device registration (UUID). On multi-user devices, this uniquely identifies a user's registration on the device.", "type": "string", "example": "a1b2c3d4-e5f6-7890-abcd-ef1234567890"}}, {"name": "page", "in": "query", "description": "Page number of paginated results", "schema": {"type": "number", "example": 1, "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "description": "Number of items per page", "schema": {"type": "number", "example": 10, "default": 10, "maximum": 50, "minimum": 1}}, {"name": "kind", "in": "query", "description": "Filter by test type.", "schema": {"type": "string", "enum": ["http", "traceroute"]}}], "responses": {"200": {"description": "Success response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/digital-experience-monitoring_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/digital-experience-monitoring_tests_response"}}}]}}}}, "4XX": {"description": "Failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/digital-experience-monitoring_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["DEX Synthetic Application Monitoring"], "x-api-token-group": ["Cloudflare DEX Write", "Cloudflare DEX Read", "Zero Trust Report", "Zero Trust Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.dex.overview.tests", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
