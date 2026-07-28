---
title: Rerun the Activation Check
page_id: operation-put-zones-zone-id-activation-check-f7de9f85
path: operations/zone
description: |-
    Triggeres a new activation check for a PENDING Zone. This can be
    triggered every 5 min for paygo/ent customers, every hour for FREE
    Zones.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/activation_check
operation_ids:
    - put-zones-zone_id-activation_check
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Rerun the Activation Check

`PUT /zones/{zone_id}/activation_check`

Operation ID: `put-zones-zone_id-activation_check`

Triggeres a new activation check for a PENDING Zone. This can be
triggered every 5 min for paygo/ent customers, every hour for FREE
Zones.

## Definition

```yaml
{"operationId": "put-zones-zone_id-activation_check", "summary": "Rerun the Activation Check", "description": "Triggeres a new activation check for a PENDING Zone. This can be\ntriggered every 5 min for paygo/ent customers, every hour for FREE\nZones.", "parameters": [{"name": "zone_id", "in": "path", "description": "Zone ID", "required": true, "schema": {"$ref": "#/components/schemas/zone-activation_identifier"}}], "responses": {"200": {"description": "Successful Response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/zone-activation_api-response-single"}, {"properties": {"result": {"type": "object", "properties": {"id": {"$ref": "#/components/schemas/zone-activation_identifier"}}}}, "type": "object"}]}}}}, "4XX": {"description": "Client Error", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/zone-activation_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Zone"], "x-api-token-group": ["Zone Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zones.activation-check", "x-fern-sdk-method-name": "trigger"}
```
