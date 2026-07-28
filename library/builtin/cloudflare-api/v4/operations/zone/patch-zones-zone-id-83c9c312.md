---
title: Edit Zone
page_id: operation-patch-zones-zone-id-5e889ebc
path: operations/zone
description: Edits a zone. Only one zone property can be changed at a time.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}
operation_ids:
    - zones-0-patch
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Edit Zone

`PATCH /zones/{zone_id}`

Operation ID: `zones-0-patch`

Edits a zone. Only one zone property can be changed at a time.

## Definition

```yaml
{"operationId": "zones-0-patch", "summary": "Edit Zone", "description": "Edits a zone. Only one zone property can be changed at a time.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zones_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"paused": {"$ref": "#/components/schemas/zones_paused"}, "plan": {"description": "(Deprecated) Please use the `/zones/{zone_id}/subscription` API\nto update a zone's plan. Changing this value will create/cancel\nassociated subscriptions. To view available plans for this zone,\nsee Zone Plans.\n", "type": "object", "properties": {"id": {"$ref": "#/components/schemas/zones_identifier"}}}, "type": {"description": "A full zone implies that DNS is hosted with Cloudflare. A partial\nzone is typically a partner-hosted zone or a CNAME setup. This\nparameter is only available to Enterprise customers or if it has\nbeen explicitly enabled on a zone.\n", "type": "string", "example": "full", "enum": ["full", "partial", "secondary", "internal"]}, "vanity_name_servers": {"$ref": "#/components/schemas/zones_vanity_name_servers"}}, "example": {"paused": true}}}}}, "responses": {"200": {"description": "Edit Zone response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/zones_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/zones_zone"}}, "type": "object"}]}}}}, "4XX": {"description": "Edit Zone response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zones_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Zone"], "x-api-token-group": ["Zone Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zones", "x-fern-sdk-method-name": "edit"}
```
