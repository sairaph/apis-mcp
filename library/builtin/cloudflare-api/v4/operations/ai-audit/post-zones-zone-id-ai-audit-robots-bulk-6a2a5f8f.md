---
title: Bulk get robots.txt rules
page_id: operation-post-zones-zone-id-ai-audit-robots-bulk-19e25bcb
path: operations/ai-audit
description: Fetches and parses robots.txt files for multiple domains within a zone in a single request. Each domain must belong to the specified zone. Results are keyed by hostname.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/ai-audit/robots/bulk
operation_ids:
    - ai-audit-bulk-get-robots
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Bulk get robots.txt rules

`POST /zones/{zone_id}/ai-audit/robots/bulk`

Operation ID: `ai-audit-bulk-get-robots`

Fetches and parses robots.txt files for multiple domains within a zone in a single request. Each domain must belong to the specified zone. Results are keyed by hostname.

## Definition

```yaml
{"operationId": "ai-audit-bulk-get-robots", "summary": "Bulk get robots.txt rules", "description": "Fetches and parses robots.txt files for multiple domains within a zone in a single request. Each domain must belong to the specified zone. Results are keyed by hostname.", "parameters": [{"name": "zone_id", "in": "path", "description": "Identifier of the zone.", "required": true, "schema": {"type": "string"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"description": "Array of domain hostnames to fetch robots.txt for. Each domain must end with the zone name. Maximum 25 domains per request.", "type": "array", "items": {"type": "string"}, "example": ["example.com", "blog.example.com"], "maxItems": 25}}}}, "responses": {"200": {"description": "Successful response with parsed robots.txt rules keyed by hostname.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/ai-audit_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/ai-audit_bulk_robots_rules"}}, "type": "object"}]}}}}, "400": {"description": "Bad request (invalid domains, exceeds 25-domain limit, or domain not in zone).", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/ai-audit_api-response-common-failure"}}}}, "401": {"description": "Unauthorized (missing or invalid authentication).", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/ai-audit_api-response-common-failure"}}}}, "403": {"description": "Forbidden (insufficient permissions or entitlement).", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/ai-audit_api-response-common-failure"}}}}, "404": {"description": "Domain not found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/ai-audit_api-response-common-failure"}}}}, "408": {"description": "Request timeout (bulk fetch exceeded global timeout).", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/ai-audit_api-response-common-failure"}}}}, "500": {"description": "Internal server error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/ai-audit_api-response-common-failure"}}}}, "503": {"description": "Upstream authentication service unavailable.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/ai-audit_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Audit"], "x-api-token-group": ["Zone Settings Write", "Zone Settings Read"]}
```
