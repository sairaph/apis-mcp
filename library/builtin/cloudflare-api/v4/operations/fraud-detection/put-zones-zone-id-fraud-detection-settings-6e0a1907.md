---
title: Update Fraud Detection Settings
page_id: operation-put-zones-zone-id-fraud-detection-settings-fad0a2be
path: operations/fraud-detection
description: |-
    Update Fraud Detection settings for a zone.

    Notes on `username_expressions` behavior:
    - If omitted or set to null, expressions are not modified.
    - If provided as an empty array `[]`, all expressions will be cleared.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/fraud_detection/settings
operation_ids:
    - fraud-detection-zone-update-settings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Fraud Detection Settings

`PUT /zones/{zone_id}/fraud_detection/settings`

Operation ID: `fraud-detection-zone-update-settings`

Update Fraud Detection settings for a zone.

Notes on `username_expressions` behavior:
- If omitted or set to null, expressions are not modified.
- If provided as an empty array `[]`, all expressions will be cleared.

## Definition

```yaml
{"operationId": "fraud-detection-zone-update-settings", "summary": "Update Fraud Detection Settings", "description": "Update Fraud Detection settings for a zone.\n\nNotes on `username_expressions` behavior:\n- If omitted or set to null, expressions are not modified.\n- If provided as an empty array `[]`, all expressions will be cleared.\n", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/fraud_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"examples": {"clear_expressions": {"summary": "Leave user_profiles unchanged, clear expressions only", "value": {"username_expressions": []}}, "configure_auth_status": {"summary": "Configure authentication status codes", "value": {"authentication_settings": {"failure_criteria": {"kind": "status_code", "status_codes": [401, 403]}, "success_criteria": {"kind": "status_code", "status_codes": [200, 201]}}}}, "enable_user_profiles": {"summary": "Enable user profiles and set two username extraction expressions", "value": {"user_profiles": "enabled", "username_expressions": ["http.request.body.form[\"username\"][0]", "lookup_json_string(http.request.body.raw, \"username\")"]}}, "partial_update_auth_status": {"summary": "Update only success codes, leave failure codes unchanged", "value": {"authentication_settings": {"success_criteria": {"kind": "status_code", "status_codes": [200]}}}}}, "schema": {"$ref": "#/components/schemas/fraud_fraud_settings"}}}}, "responses": {"200": {"description": "Updated Fraud Detection settings response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/fraud_fraud_settings_response_body"}}}}, "4XX": {"description": "Update Fraud Detection settings failure", "content": {"application/json": {"examples": {"invalid_username_expression": {"summary": "Invalid username expression type", "value": {"errors": [{"code": 10400, "message": "Bad Request"}], "messages": [{"code": 10400, "message": "'http.request.body.form[\"username\"]' is not a valid value for username because the expression has the wrong type: got array[bytes], want bytes"}], "result": null, "success": false}}}, "schema": {"allOf": [{"$ref": "#/components/schemas/fraud_fraud_settings_response_body"}, {"$ref": "#/components/schemas/fraud_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Fraud Detection"], "x-api-token-group": ["Fraud Detection Write"]}
```
