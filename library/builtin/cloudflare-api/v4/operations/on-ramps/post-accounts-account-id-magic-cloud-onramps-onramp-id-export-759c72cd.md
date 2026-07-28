---
title: Export as Terraform
page_id: operation-post-accounts-account-id-magic-cloud-onramps-onramp-id-export-0c1e6b5d
path: operations/on-ramps
description: Export an On-ramp to terraform ready file(s) (Closed Beta).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/magic/cloud/onramps/{onramp_id}/export
operation_ids:
    - onramps-export
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Export as Terraform

`POST /accounts/{account_id}/magic/cloud/onramps/{onramp_id}/export`

Operation ID: `onramps-export`

Export an On-ramp to terraform ready file(s) (Closed Beta).

## Definition

```yaml
{"operationId": "onramps-export", "summary": "Export as Terraform", "description": "Export an On-ramp to terraform ready file(s) (Closed Beta).", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mcn_account_id"}}, {"name": "onramp_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mcn_onramp_id"}}], "responses": {"201": {"description": "Exported file.", "headers": {"Content-Disposition": {"schema": {"type": "string", "example": "attachment; filename=\"exported-file.zip\""}}}, "content": {"application/zip": {"schema": {"type": "string", "format": "binary"}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "401": {"description": "Invalid Credentials.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "404": {"description": "Not Found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "409": {"description": "Conflict.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "500": {"description": "Internal Server Error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["On-ramps"], "x-api-token-group": ["Magic WAN Write"]}
```
