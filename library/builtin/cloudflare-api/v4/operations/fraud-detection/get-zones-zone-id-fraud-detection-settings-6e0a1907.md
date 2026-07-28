---
title: Get Fraud Detection Settings
page_id: operation-get-zones-zone-id-fraud-detection-settings-c2ec7d09
path: operations/fraud-detection
description: Retrieve Fraud Detection settings for a zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/fraud_detection/settings
operation_ids:
    - fraud-detection-zone-get-settings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Fraud Detection Settings

`GET /zones/{zone_id}/fraud_detection/settings`

Operation ID: `fraud-detection-zone-get-settings`

Retrieve Fraud Detection settings for a zone.

## Definition

```yaml
{"operationId": "fraud-detection-zone-get-settings", "summary": "Get Fraud Detection Settings", "description": "Retrieve Fraud Detection settings for a zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/fraud_identifier"}}], "responses": {"200": {"description": "Fraud Detection settings response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/fraud_fraud_settings_response_body"}}}}, "4XX": {"description": "Fraud Detection settings response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/fraud_fraud_settings_response_body"}, {"$ref": "#/components/schemas/fraud_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Fraud Detection"], "x-api-token-group": ["Fraud Detection Read", "Fraud Detection Write"]}
```
