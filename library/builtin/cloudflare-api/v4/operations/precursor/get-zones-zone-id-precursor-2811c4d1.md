---
title: Get Zone Precursor Config
page_id: operation-get-zones-zone-id-precursor-a68465c5
path: operations/precursor
description: |-
    Retrieve a zone's Precursor configuration: the zone-level
    `default_mode` and the ordered list of `enforcement_rules`.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/precursor
operation_ids:
    - precursor-for-a-zone-get-config
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Zone Precursor Config

`GET /zones/{zone_id}/precursor`

Operation ID: `precursor-for-a-zone-get-config`

Retrieve a zone's Precursor configuration: the zone-level
`default_mode` and the ordered list of `enforcement_rules`.

## Definition

```yaml
{"operationId": "precursor-for-a-zone-get-config", "summary": "Get Zone Precursor Config", "description": "Retrieve a zone's Precursor configuration: the zone-level\n`default_mode` and the ordered list of `enforcement_rules`.\n", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/precursor_identifier"}}], "responses": {"200": {"description": "Precursor config response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/precursor_precursor_config_response_body"}}}}, "4XX": {"description": "Precursor config response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/precursor_precursor_config_response_body"}, {"$ref": "#/components/schemas/precursor_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Precursor"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
