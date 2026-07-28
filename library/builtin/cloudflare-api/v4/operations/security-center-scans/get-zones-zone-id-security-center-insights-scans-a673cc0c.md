---
title: Get Recent Zone Scans
page_id: operation-get-zones-zone-id-security-center-insights-scans-17aed742
path: operations/security-center-scans
description: Returns the most recent on-demand scans for a specific zone, up to a maximum of 5. Each scan includes its ID, start time, and current status. Results include both zone-specific scans and account-wide scans (which cover all zones). Also returns quota information showing how many scans have been used and how many remain in the current 24-hour window.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/security-center/insights/scans
operation_ids:
    - get-security-center-zone-scans
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Recent Zone Scans

`GET /zones/{zone_id}/security-center/insights/scans`

Operation ID: `get-security-center-zone-scans`

Returns the most recent on-demand scans for a specific zone, up to a maximum of 5. Each scan includes its ID, start time, and current status. Results include both zone-specific scans and account-wide scans (which cover all zones). Also returns quota information showing how many scans have been used and how many remain in the current 24-hour window.

## Definition

```yaml
{"operationId": "get-security-center-zone-scans", "summary": "Get Recent Zone Scans", "description": "Returns the most recent on-demand scans for a specific zone, up to a maximum of 5. Each scan includes its ID, start time, and current status. Results include both zone-specific scans and account-wide scans (which cover all zones). Also returns quota information showing how many scans have been used and how many remain in the current 24-hour window.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/security-center_identifier"}}], "responses": {"200": {"description": "The request was successful.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/security-center_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/security-center_scansListResponse"}}, "type": "object"}]}}}}, "4XX": {"description": "A client error occurred.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/security-center_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}], "tags": ["Security Center Scans"], "x-api-token-group": ["Zone Settings Write", "Zone Settings Read"]}
```
