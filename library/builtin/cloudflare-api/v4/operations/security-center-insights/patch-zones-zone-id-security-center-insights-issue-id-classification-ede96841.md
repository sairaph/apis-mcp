---
title: Updates Zone Security Center Insight Classification
page_id: operation-patch-zones-zone-id-security-center-insights-issue-id-classification-6a716986
path: operations/security-center-insights
description: Updates the user classification for a zone-specific Security Center insight. Valid values are 'false_positive' or 'accept_risk'. To reset, set classification to null. Cannot change directly between classification values - must reset to null first.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/security-center/insights/{issue_id}/classification
operation_ids:
    - update-zone-security-center-insight-classification
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Updates Zone Security Center Insight Classification

`PATCH /zones/{zone_id}/security-center/insights/{issue_id}/classification`

Operation ID: `update-zone-security-center-insight-classification`

Updates the user classification for a zone-specific Security Center insight. Valid values are 'false_positive' or 'accept_risk'. To reset, set classification to null. Cannot change directly between classification values - must reset to null first.

## Definition

```yaml
{"operationId": "update-zone-security-center-insight-classification", "summary": "Updates Zone Security Center Insight Classification", "description": "Updates the user classification for a zone-specific Security Center insight. Valid values are 'false_positive' or 'accept_risk'. To reset, set classification to null. Cannot change directly between classification values - must reset to null first.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/security-center_identifier"}}, {"name": "issue_id", "in": "path", "required": true, "schema": {"type": "string"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/security-center_userClassificationUpdate"}}}}, "responses": {"200": {"description": "The request was successful.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/security-center_api-response-single"}]}}}}, "4XX": {"description": "A client error occurred.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/security-center_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}], "tags": ["Security Center Insights"]}
```
