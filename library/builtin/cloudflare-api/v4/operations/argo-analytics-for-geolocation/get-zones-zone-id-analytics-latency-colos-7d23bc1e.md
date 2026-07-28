---
title: Argo Analytics for a zone at different PoPs
page_id: operation-get-zones-zone-id-analytics-latency-colos-bccfb1ef
path: operations/argo-analytics-for-geolocation
description: Retrieves Argo Smart Routing analytics broken down by geographic points of presence (PoPs). Shows latency improvements and routing efficiency per location.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/analytics/latency/colos
operation_ids:
    - argo-analytics-for-geolocation-argo-analytics-for-a-zone-at-different-po-ps
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Argo Analytics for a zone at different PoPs

`GET /zones/{zone_id}/analytics/latency/colos`

Operation ID: `argo-analytics-for-geolocation-argo-analytics-for-a-zone-at-different-po-ps`

Retrieves Argo Smart Routing analytics broken down by geographic points of presence (PoPs). Shows latency improvements and routing efficiency per location.

## Definition

```yaml
{"operationId": "argo-analytics-for-geolocation-argo-analytics-for-a-zone-at-different-po-ps", "summary": "Argo Analytics for a zone at different PoPs", "description": "Retrieves Argo Smart Routing analytics broken down by geographic points of presence (PoPs). Shows latency improvements and routing efficiency per location.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/argo-analytics_identifier"}}], "responses": {"200": {"description": "Argo Analytics for a zone at different PoPs response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/argo-analytics_response_single"}}}}, "4XX": {"description": "Argo Analytics for a zone at different PoPs response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/argo-analytics_response_single"}, {"$ref": "#/components/schemas/argo-analytics_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Argo Analytics for Geolocation"], "x-api-token-group": ["Analytics Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
