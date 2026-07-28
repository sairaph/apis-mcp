---
title: zone-analytics-api_user_dashboard_response
page_id: schema-zone-analytics-api-user-dashboard-response-d8b7a745
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zone-analytics-api_user_dashboard_response

```yaml
{"allOf": [{"$ref": "#/components/schemas/zone-analytics-api_api-response-common"}, {"properties": {"query": {"$ref": "#/components/schemas/zone-analytics-api_query_response"}, "result": {"description": "Per-zone analytics data for all zones accessible to the user.", "type": "array", "items": {"$ref": "#/components/schemas/zone-analytics-api_user_zone"}}}}]}
```
