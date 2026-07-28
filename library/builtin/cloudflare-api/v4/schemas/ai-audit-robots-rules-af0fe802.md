---
title: ai-audit_robots_rules
page_id: schema-ai-audit-robots-rules-af0fe802
path: schemas
description: Parsed robots.txt rules for a single domain.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# ai-audit_robots_rules

Parsed robots.txt rules for a single domain.

```yaml
{"description": "Parsed robots.txt rules for a single domain.", "type": "object", "properties": {"sitemaps": {"description": "List of sitemap URLs found in robots.txt.", "type": "array", "items": {"type": "string"}}, "status": {"description": "HTTP status code from fetching the robots.txt file.", "type": "integer"}, "userAgents": {"description": "Map of user-agent string to its parsed rules.", "type": "object", "additionalProperties": {"$ref": "#/components/schemas/ai-audit_user_agent_rules"}}}, "required": ["userAgents"]}
```
