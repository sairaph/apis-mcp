---
title: load-balancing_health_details
page_id: schema-load-balancing-health-details-d6cbfe8e
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# load-balancing_health_details

```yaml
{"allOf": [{"$ref": "#/components/schemas/load-balancing_api-response-single"}, {"properties": {"result": {"description": "A list of regions from which to run health checks. Null means every Cloudflare data center.", "type": "object", "example": {"pool_id": "17b5962d775c646f3f9725cbc7a53df4", "pop_health": {"Amsterdam, NL": {"healthy": true, "origins": [{"2001:DB8::5": {"failure_reason": "No failures", "healthy": true, "response_code": 401, "rtt": "12.1ms"}}]}}}, "properties": {"pool_id": {"description": "Pool ID.", "type": "string", "example": "17b5962d775c646f3f9725cbc7a53df4", "x-auditable": true}, "pop_health": {"description": "List of regions and associated health status.", "type": "object", "properties": {"healthy": {"description": "Whether health check in region is healthy.", "type": "boolean", "example": true, "x-auditable": true}, "origins": {"type": "array", "items": {"$ref": "#/components/schemas/load-balancing_origin-health"}}}}}}}, "type": "object"}]}
```
