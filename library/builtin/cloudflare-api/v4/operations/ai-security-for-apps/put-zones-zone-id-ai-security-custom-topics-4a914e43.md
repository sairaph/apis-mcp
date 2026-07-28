---
title: Set AI Security for Apps Custom Topics
page_id: operation-put-zones-zone-id-ai-security-custom-topics-c35c60c2
path: operations/ai-security-for-apps
description: |-
    Set the AI Security for Apps custom topic categories for a zone.

    A maximum of 20 custom topics can be configured per zone.
    Each topic label must be 2–20 characters using only lowercase letters (a–z), digits (0–9), and hyphens.
    Each topic description must be 2–50 printable ASCII characters.

    Changes can take up to a minute to propagate to the zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/ai-security/custom-topics
operation_ids:
    - ai-security-custom-topics-put
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Set AI Security for Apps Custom Topics

`PUT /zones/{zone_id}/ai-security/custom-topics`

Operation ID: `ai-security-custom-topics-put`

Set the AI Security for Apps custom topic categories for a zone.

A maximum of 20 custom topics can be configured per zone.
Each topic label must be 2–20 characters using only lowercase letters (a–z), digits (0–9), and hyphens.
Each topic description must be 2–50 printable ASCII characters.

Changes can take up to a minute to propagate to the zone.

## Definition

```yaml
{"operationId": "ai-security-custom-topics-put", "summary": "Set AI Security for Apps Custom Topics", "description": "Set the AI Security for Apps custom topic categories for a zone.\n\nA maximum of 20 custom topics can be configured per zone.\nEach topic label must be 2–20 characters using only lowercase letters (a–z), digits (0–9), and hyphens.\nEach topic description must be 2–50 printable ASCII characters.\n\nChanges can take up to a minute to propagate to the zone.\n", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waf-product-api-bundle_zone_id"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/waf-product-api-bundle_custom-topics"}}}}, "responses": {"200": {"description": "Set AI Security for Apps custom topics response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/waf-product-api-bundle_response-custom-topics"}}}}, "4XX": {"description": "Set AI Security for Apps custom topics failure response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/waf-product-api-bundle_response-custom-topics"}, {"$ref": "#/components/schemas/waf-product-api-bundle_api-response-common-failure-3"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["AI Security for Apps"]}
```
