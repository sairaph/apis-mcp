---
title: zones_search_for_agents
page_id: schema-zones-search-for-agents-d93974f7
path: schemas
description: |-
    When enabled, Cloudflare provisions an AI Search instance for the zone
    and exposes a /.well-known/ai-search endpoint that AI agents can query.
    Markdown responses also receive an agent: YAML capability block advertising
    the search endpoint.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_search_for_agents

When enabled, Cloudflare provisions an AI Search instance for the zone
and exposes a /.well-known/ai-search endpoint that AI agents can query.
Markdown responses also receive an agent: YAML capability block advertising
the search endpoint.

```yaml
{"description": "When enabled, Cloudflare provisions an AI Search instance for the zone\nand exposes a /.well-known/ai-search endpoint that AI agents can query.\nMarkdown responses also receive an agent: YAML capability block advertising\nthe search endpoint.", "allOf": [{"$ref": "#/components/schemas/zones_base"}, {"properties": {"id": {"description": "ID of the zone setting.", "example": "search_for_agents", "enum": ["search_for_agents"]}, "value": {"$ref": "#/components/schemas/zones_search_for_agents_value"}}}], "title": "Search for Agents"}
```
