---
title: Delete Leaked Credential Checks Custom Detection
page_id: operation-delete-zones-zone-id-leaked-credential-checks-detections-detection-id-b2c0c0ab
path: operations/leaked-credential-checks
description: Remove user-defined detection pattern for Leaked Credential Checks.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/leaked-credential-checks/detections/{detection_id}
operation_ids:
    - waf-product-api-leaked-credentials-delete-detection
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Leaked Credential Checks Custom Detection

`DELETE /zones/{zone_id}/leaked-credential-checks/detections/{detection_id}`

Operation ID: `waf-product-api-leaked-credentials-delete-detection`

Remove user-defined detection pattern for Leaked Credential Checks.

## Definition

```yaml
{"operationId": "waf-product-api-leaked-credentials-delete-detection", "summary": "Delete Leaked Credential Checks Custom Detection", "description": "Remove user-defined detection pattern for Leaked Credential Checks.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waf-product-api-bundle_identifier"}}, {"name": "detection_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waf-product-api-bundle_detection-id"}}], "responses": {"200": {"description": "Delete Leaked Credential Checks custom detection response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/waf-product-api-bundle_api-response-common"}}}}, "4XX": {"description": "Delete Leaked Credential Checks custom detection failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/waf-product-api-bundle_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Leaked Credential Checks"], "x-api-token-group": ["Zone WAF Write", "Account WAF Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "leaked-credential-checks.detections", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
