---
title: Retrieves security.txt
page_id: operation-get-zones-zone-id-security-center-securitytxt-d4bb7b52
path: operations/security-txt
description: Retrieves the current security.txt file configuration for a zone, used for security vulnerability reporting.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/security-center/securitytxt
operation_ids:
    - get-security-txt
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Retrieves security.txt

`GET /zones/{zone_id}/security-center/securitytxt`

Operation ID: `get-security-txt`

Retrieves the current security.txt file configuration for a zone, used for security vulnerability reporting.

## Definition

```yaml
{"operationId": "get-security-txt", "summary": "Retrieves security.txt", "description": "Retrieves the current security.txt file configuration for a zone, used for security vulnerability reporting.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/security-center_identifier"}}], "responses": {"200": {"description": "The request was successful.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/security-center_api-response-common"}, {"properties": {"result": {"anyOf": [{"$ref": "#/components/schemas/security-center_securityTxt"}]}}, "type": "object"}]}}}}, "4XX": {"description": "A client error occurred.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/security-center_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}], "tags": ["security.txt"], "x-api-token-group": ["Zone Settings Write", "Zone Settings Read"]}
```
