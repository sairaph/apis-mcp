---
title: Start On-Demand Zone Scan
page_id: operation-post-zones-zone-id-security-center-insights-scans-00117bdc
path: operations/security-center-scans
description: Initiates an on-demand security scan for a specific zone. Rate limited to 5 scans per account per 24-hour window (shared with account-level scans).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/security-center/insights/scans
operation_ids:
    - start-security-center-zone-scan
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Start On-Demand Zone Scan

`POST /zones/{zone_id}/security-center/insights/scans`

Operation ID: `start-security-center-zone-scan`

Initiates an on-demand security scan for a specific zone. Rate limited to 5 scans per account per 24-hour window (shared with account-level scans).

## Definition

```yaml
{"operationId": "start-security-center-zone-scan", "summary": "Start On-Demand Zone Scan", "description": "Initiates an on-demand security scan for a specific zone. Rate limited to 5 scans per account per 24-hour window (shared with account-level scans).", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/security-center_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/security-center_newScanRequest"}}}}, "responses": {"200": {"description": "The request was successful.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/security-center_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/security-center_newScanResponse"}}, "type": "object"}]}}}}, "4XX": {"description": "A client error occurred.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/security-center_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}], "tags": ["Security Center Scans"], "x-api-token-group": ["Zone Settings Write"]}
```
