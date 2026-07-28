---
title: Create Zone
page_id: operation-post-zones-5fc832c2
path: operations/zone
description: |-
    Creates a new zone (domain) in your Cloudflare account.

    The zone is created in a pending state and must be activated by updating your domain's
    nameservers to point to Cloudflare, or by completing the verification process for partial
    (CNAME) setups.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones
operation_ids:
    - zones-post
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Zone

`POST /zones`

Operation ID: `zones-post`

Creates a new zone (domain) in your Cloudflare account.

The zone is created in a pending state and must be activated by updating your domain's
nameservers to point to Cloudflare, or by completing the verification process for partial
(CNAME) setups.

## Definition

```yaml
{"operationId": "zones-post", "summary": "Create Zone", "description": "Creates a new zone (domain) in your Cloudflare account.\n\nThe zone is created in a pending state and must be activated by updating your domain's\nnameservers to point to Cloudflare, or by completing the verification process for partial\n(CNAME) setups.\n", "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"account": {"type": "object", "properties": {"id": {"$ref": "#/components/schemas/zones_identifier"}}}, "name": {"$ref": "#/components/schemas/zones_name"}, "type": {"$ref": "#/components/schemas/zones_type"}}, "required": ["name", "account"]}}}}, "responses": {"200": {"description": "Create Zone response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/zones_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/zones_zone"}}, "type": "object"}]}}}}, "4XX": {"description": "Create Zone response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zones_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Zone"], "x-api-token-group": ["Zone Zone Edit", "Zone DNS Edit"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zones", "x-fern-sdk-method-name": "create"}
```
