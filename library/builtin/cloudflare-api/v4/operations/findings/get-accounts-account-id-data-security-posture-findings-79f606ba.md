---
title: List posture findings
page_id: operation-get-accounts-account-id-data-security-posture-findings-0c674173
path: operations/findings
description: "List all security findings that have been identified as being problematic. \nThis will return a list of findings regardless if they have been ignored or not."
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/data-security/posture/findings
operation_ids:
    - ListFindings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List posture findings

`GET /accounts/{account_id}/data-security/posture/findings`

Operation ID: `ListFindings`

List all security findings that have been identified as being problematic.
This will return a list of findings regardless if they have been ignored or not.

## Definition

```yaml
{"operationId": "ListFindings", "summary": "List posture findings", "description": "List all security findings that have been identified as being problematic. \nThis will return a list of findings regardless if they have been ignored or not.", "parameters": [{"$ref": "#/components/parameters/posture-api_AccountTag"}, {"name": "cursor", "in": "query", "description": "A cursor for pagination. Obtained from the `result_info.cursor` field of a previous response.", "schema": {"type": "string"}}, {"$ref": "#/components/parameters/posture-api_Direction"}, {"$ref": "#/components/parameters/posture-api_Ignored"}, {"$ref": "#/components/parameters/posture-api_IntegrationId"}, {"$ref": "#/components/parameters/posture-api_MaxAfflictionDate"}, {"$ref": "#/components/parameters/posture-api_MinAfflictionDate"}, {"$ref": "#/components/parameters/posture-api_Observation"}, {"$ref": "#/components/parameters/posture-api_FindingOrder"}, {"$ref": "#/components/parameters/posture-api_Page"}, {"$ref": "#/components/parameters/posture-api_PerPage"}, {"$ref": "#/components/parameters/posture-api_Product"}, {"$ref": "#/components/parameters/posture-api_Search"}, {"$ref": "#/components/parameters/posture-api_Severity"}, {"$ref": "#/components/parameters/posture-api_Type"}, {"$ref": "#/components/parameters/posture-api_Vendor"}, {"$ref": "#/components/parameters/posture-api_FindingTypeIds"}], "responses": {"200": {"description": "OK: Successful HTTP request", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_paginated-finding-list"}}}}, "400": {"description": "Bad Request: Invalid parameters", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_error-response"}}}}, "401": {"description": "Unauthorized: Authentication required", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_error-response"}}}}, "403": {"description": "Forbidden: Insufficient permissions", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_error-response"}}}}}, "security": [{"api_token": []}], "tags": ["findings"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"]}
```
