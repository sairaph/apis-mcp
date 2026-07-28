---
title: Get robots.txt rules
page_id: operation-get-zones-zone-id-ai-audit-robots-6f553ea9
path: operations/ai-audit
description: Fetches and parses the robots.txt file for a zone or a specific subdomain within the zone. Returns parsed user-agent rules, content signals, and sitemaps.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/ai-audit/robots
operation_ids:
    - ai-audit-get-robots
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get robots.txt rules

`GET /zones/{zone_id}/ai-audit/robots`

Operation ID: `ai-audit-get-robots`

Fetches and parses the robots.txt file for a zone or a specific subdomain within the zone. Returns parsed user-agent rules, content signals, and sitemaps.

## Definition

```yaml
{"operationId": "ai-audit-get-robots", "summary": "Get robots.txt rules", "description": "Fetches and parses the robots.txt file for a zone or a specific subdomain within the zone. Returns parsed user-agent rules, content signals, and sitemaps.", "parameters": [{"name": "zone_id", "in": "path", "description": "Identifier of the zone.", "required": true, "schema": {"type": "string"}}, {"name": "subdomain", "in": "query", "description": "Optional subdomain to fetch robots.txt for. If omitted, fetches robots.txt for the zone apex domain.", "schema": {"type": "string"}}], "responses": {"200": {"description": "Successful response with parsed robots.txt rules.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/ai-audit_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/ai-audit_robots_rules"}}, "type": "object"}]}}}}, "400": {"description": "Bad request (invalid subdomain or missing parameters).", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/ai-audit_api-response-common-failure"}}}}, "401": {"description": "Unauthorized (missing or invalid authentication).", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/ai-audit_api-response-common-failure"}}}}, "403": {"description": "Forbidden (insufficient permissions or entitlement).", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/ai-audit_api-response-common-failure"}}}}, "404": {"description": "Domain not found or robots.txt not found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/ai-audit_api-response-common-failure"}}}}, "500": {"description": "Internal server error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/ai-audit_api-response-common-failure"}}}}, "503": {"description": "Upstream authentication service unavailable.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/ai-audit_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Audit"], "x-api-token-group": ["Zone Settings Write", "Zone Settings Read"]}
```
