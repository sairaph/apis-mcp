---
title: Rules
page_id: operation-get-zones-zone-id-cloud-connector-rules-6c78589d
path: operations/zone-cloud-connector-rules-get
description: Retrieves the Cloud Connector rules configured for a zone. Rules define how traffic is routed to cloud services.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/cloud_connector/rules
operation_ids:
    - zone-cloud-connector-rules
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Rules

`GET /zones/{zone_id}/cloud_connector/rules`

Operation ID: `zone-cloud-connector-rules`

Retrieves the Cloud Connector rules configured for a zone. Rules define how traffic is routed to cloud services.

## Definition

```yaml
{"operationId": "zone-cloud-connector-rules", "summary": "Rules", "description": "Retrieves the Cloud Connector rules configured for a zone. Rules define how traffic is routed to cloud services.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloud-connector_identifier"}}], "responses": {"200": {"description": "Cloud Connector rules response", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/cloud-connector_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/cloud-connector_rules"}}, "type": "object"}]}}}}, "4XX": {"description": "Cloud Connector response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/cloud-connector_api-response-common-failure"}]}}}}, "5XX": {"description": "Cloud Connector response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/cloud-connector_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Zone Cloud Connector Rules GET"], "x-api-token-group": ["Cloud Connector Read", "Cloud Connector Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
