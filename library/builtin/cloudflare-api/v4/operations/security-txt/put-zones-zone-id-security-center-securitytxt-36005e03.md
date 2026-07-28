---
title: Updates security.txt
page_id: operation-put-zones-zone-id-security-center-securitytxt-3a7234f2
path: operations/security-txt
description: Updates the security.txt file configuration for a zone, which provides security researchers with vulnerability reporting information.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/security-center/securitytxt
operation_ids:
    - update-security-txt
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Updates security.txt

`PUT /zones/{zone_id}/security-center/securitytxt`

Operation ID: `update-security-txt`

Updates the security.txt file configuration for a zone, which provides security researchers with vulnerability reporting information.

## Definition

```yaml
{"operationId": "update-security-txt", "summary": "Updates security.txt", "description": "Updates the security.txt file configuration for a zone, which provides security researchers with vulnerability reporting information.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/security-center_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/security-center_securityTxt"}}}}, "responses": {"200": {"description": "The request was successful.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/security-center_api-response-single"}]}}}}, "4XX": {"description": "A client error occurred.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/security-center_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}], "tags": ["security.txt"], "x-api-token-group": ["Zone Settings Write"]}
```
