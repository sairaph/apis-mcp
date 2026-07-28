---
title: List all finding types
page_id: operation-get-accounts-account-id-data-security-posture-finding-types-fc0c689c
path: operations/finding-types
description: List all available finding types with pagination support.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/data-security/posture/finding_types
operation_ids:
    - ListFindingTypes
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List all finding types

`GET /accounts/{account_id}/data-security/posture/finding_types`

Operation ID: `ListFindingTypes`

List all available finding types with pagination support.

## Definition

```yaml
{"operationId": "ListFindingTypes", "summary": "List all finding types", "description": "List all available finding types with pagination support.", "parameters": [{"$ref": "#/components/parameters/posture-api_AccountTag"}, {"$ref": "#/components/parameters/posture-api_Page"}, {"$ref": "#/components/parameters/posture-api_PerPage"}], "responses": {"200": {"description": "OK: Successful HTTP request", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_paginated-finding-type-list"}}}}, "400": {"description": "Bad Request: Invalid request parameters"}}, "security": [{"api_token": []}], "tags": ["finding_types"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "data-security.posture", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
