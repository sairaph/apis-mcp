---
title: Deletes security.txt
page_id: operation-delete-zones-zone-id-security-center-securitytxt-6c9ab1a5
path: operations/security-txt
description: Removes the security.txt file configuration for a zone. The /.well-known/security.txt endpoint will no longer be served.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/security-center/securitytxt
operation_ids:
    - delete-security-txt
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Deletes security.txt

`DELETE /zones/{zone_id}/security-center/securitytxt`

Operation ID: `delete-security-txt`

Removes the security.txt file configuration for a zone. The /.well-known/security.txt endpoint will no longer be served.

## Definition

```yaml
{"operationId": "delete-security-txt", "summary": "Deletes security.txt", "description": "Removes the security.txt file configuration for a zone. The /.well-known/security.txt endpoint will no longer be served.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/security-center_identifier"}}], "responses": {"200": {"description": "The request was successful.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/security-center_api-response-single"}]}}}}, "4XX": {"description": "A client error occurred.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/security-center_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}], "tags": ["security.txt"], "x-api-token-group": ["Zone Settings Write"]}
```
