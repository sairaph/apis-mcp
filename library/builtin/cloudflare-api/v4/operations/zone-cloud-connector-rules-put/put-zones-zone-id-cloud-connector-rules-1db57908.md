---
title: Put Rules
page_id: operation-put-zones-zone-id-cloud-connector-rules-35a0a6dd
path: operations/zone-cloud-connector-rules-put
description: Updates Cloud Connector rules for a zone, replacing the existing rule configuration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/cloud_connector/rules
operation_ids:
    - zone-cloud-conenctor-rules-put
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Put Rules

`PUT /zones/{zone_id}/cloud_connector/rules`

Operation ID: `zone-cloud-conenctor-rules-put`

Updates Cloud Connector rules for a zone, replacing the existing rule configuration.

## Definition

```yaml
{"operationId": "zone-cloud-conenctor-rules-put", "summary": "Put Rules", "description": "Updates Cloud Connector rules for a zone, replacing the existing rule configuration.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloud-connector_identifier"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "array", "items": {"$ref": "#/components/schemas/cloud-connector_rule"}}}}}, "responses": {"200": {"description": "Cloud Connector rules response", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/cloud-connector_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/cloud-connector_rules"}}, "type": "object"}]}}}}, "4XX": {"description": "Cloud Connector response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/cloud-connector_api-response-common-failure"}]}}}}, "5XX": {"description": "Cloud Connector response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/cloud-connector_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Zone Cloud Connector Rules PUT"], "x-api-token-group": ["Cloud Connector Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
