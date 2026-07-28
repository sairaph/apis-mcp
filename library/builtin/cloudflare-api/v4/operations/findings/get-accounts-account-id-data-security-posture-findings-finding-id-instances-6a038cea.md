---
title: List instances of a finding
page_id: operation-get-accounts-account-id-data-security-posture-findings-finding-id-instan-68f890ae
path: operations/findings
description: Lists all security finding instances for a given security finding.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/data-security/posture/findings/{finding_id}/instances
operation_ids:
    - ListFindingInstances
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List instances of a finding

`GET /accounts/{account_id}/data-security/posture/findings/{finding_id}/instances`

Operation ID: `ListFindingInstances`

Lists all security finding instances for a given security finding.

## Definition

```yaml
{"operationId": "ListFindingInstances", "summary": "List instances of a finding", "description": "Lists all security finding instances for a given security finding.", "parameters": [{"$ref": "#/components/parameters/posture-api_AccountTag"}, {"$ref": "#/components/parameters/posture-api_Archived"}, {"$ref": "#/components/parameters/posture-api_FindingIdByte"}, {"name": "cursor", "in": "query", "description": "A cursor for pagination. Obtained from the `result_info.cursor` field of a previous response.", "schema": {"type": "string"}}, {"$ref": "#/components/parameters/posture-api_Direction"}, {"$ref": "#/components/parameters/posture-api_MaxAfflictionDate"}, {"$ref": "#/components/parameters/posture-api_MinAfflictionDate"}, {"$ref": "#/components/parameters/posture-api_FindingInstanceOrder"}, {"$ref": "#/components/parameters/posture-api_Page"}, {"$ref": "#/components/parameters/posture-api_PerPage"}, {"$ref": "#/components/parameters/posture-api_Search"}, {"$ref": "#/components/parameters/posture-api_RemediationStatuses"}, {"$ref": "#/components/parameters/posture-api_FindingInstanceIds"}, {"$ref": "#/components/parameters/posture-api_AssetIds"}], "responses": {"200": {"description": "OK: Successful HTTP request", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_paginated-finding-instance-list"}}}}, "400": {"description": "Bad Request: Invalid request parameters"}, "404": {"description": "Not Found: Finding not found"}}, "security": [{"api_token": []}], "tags": ["findings"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "data-security.posture", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
