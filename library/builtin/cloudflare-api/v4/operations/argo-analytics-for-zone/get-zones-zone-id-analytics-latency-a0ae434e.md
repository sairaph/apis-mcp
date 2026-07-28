---
title: Argo Analytics for a zone
page_id: operation-get-zones-zone-id-analytics-latency-d96f4aba
path: operations/argo-analytics-for-zone
description: Retrieves aggregate Argo Smart Routing analytics for a zone, including latency improvements, bandwidth savings, and routing statistics.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/analytics/latency
operation_ids:
    - argo-analytics-for-zone-argo-analytics-for-a-zone
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Argo Analytics for a zone

`GET /zones/{zone_id}/analytics/latency`

Operation ID: `argo-analytics-for-zone-argo-analytics-for-a-zone`

Retrieves aggregate Argo Smart Routing analytics for a zone, including latency improvements, bandwidth savings, and routing statistics.

## Definition

```yaml
{"operationId": "argo-analytics-for-zone-argo-analytics-for-a-zone", "summary": "Argo Analytics for a zone", "description": "Retrieves aggregate Argo Smart Routing analytics for a zone, including latency improvements, bandwidth savings, and routing statistics.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/argo-analytics_identifier"}}, {"name": "bins", "in": "query", "schema": {"type": "string"}}], "responses": {"200": {"description": "Argo Analytics for a zone response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/argo-analytics_response_single"}}}}, "4XX": {"description": "Argo Analytics for a zone response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/argo-analytics_response_single"}, {"$ref": "#/components/schemas/argo-analytics_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Argo Analytics for Zone"], "x-api-token-group": ["Analytics Read"], "x-cfPermissionsRequired": {"enum": ["#analytics:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
