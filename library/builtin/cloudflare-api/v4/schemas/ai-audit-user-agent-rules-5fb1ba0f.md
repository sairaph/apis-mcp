---
title: ai-audit_user_agent_rules
page_id: schema-ai-audit-user-agent-rules-5fb1ba0f
path: schemas
description: Parsed rules for a specific user-agent.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# ai-audit_user_agent_rules

Parsed rules for a specific user-agent.

```yaml
{"description": "Parsed rules for a specific user-agent.", "type": "object", "properties": {"allow": {"description": "List of allowed path patterns.", "type": "array", "items": {"type": "string"}}, "contentSignals": {"$ref": "#/components/schemas/ai-audit_content_signal"}, "crawlDelay": {"description": "Crawl delay in seconds.", "type": "number"}, "disallow": {"description": "List of disallowed path patterns.", "type": "array", "items": {"type": "string"}}}, "required": ["allow", "disallow"]}
```
