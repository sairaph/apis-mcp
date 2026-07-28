---
title: Archives Zone Security Center Insight
page_id: operation-put-zones-zone-id-security-center-insights-issue-id-dismiss-94b02df4
path: operations/security-center-insights
description: Archives a zone-specific Security Center insight, removing it from the active zone insights while preserving historical data.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/security-center/insights/{issue_id}/dismiss
operation_ids:
    - archive-zone-security-center-insight
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Archives Zone Security Center Insight

`PUT /zones/{zone_id}/security-center/insights/{issue_id}/dismiss`

Operation ID: `archive-zone-security-center-insight`

Archives a zone-specific Security Center insight, removing it from the active zone insights while preserving historical data.

## Definition

```yaml
{"operationId": "archive-zone-security-center-insight", "summary": "Archives Zone Security Center Insight", "description": "Archives a zone-specific Security Center insight, removing it from the active zone insights while preserving historical data.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/security-center_identifier"}}, {"name": "issue_id", "in": "path", "required": true, "schema": {"type": "string"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"dismiss": {"type": "boolean", "default": true, "x-auditable": true}}}}}}, "responses": {"200": {"description": "The request was successful.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/security-center_api-response-single"}]}}}}, "4XX": {"description": "A client error occurred.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/security-center_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}], "tags": ["Security Center Insights"], "x-api-token-group": ["Zone Settings Write"]}
```
