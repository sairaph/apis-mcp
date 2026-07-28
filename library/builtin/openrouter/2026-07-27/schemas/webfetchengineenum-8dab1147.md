---
title: WebFetchEngineEnum
page_id: schema-webfetchengineenum-8dab1147
path: schemas
description: Which fetch engine to use. "auto" (default) uses native if the provider supports it, otherwise Exa. "native" forces the provider's built-in fetch. "exa" uses Exa Contents API. "openrouter" uses direct HTTP fetch. "firecrawl" uses Firecrawl scrape (requires BYOK). "parallel" uses the Parallel extract API.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# WebFetchEngineEnum

Which fetch engine to use. "auto" (default) uses native if the provider supports it, otherwise Exa. "native" forces the provider's built-in fetch. "exa" uses Exa Contents API. "openrouter" uses direct HTTP fetch. "firecrawl" uses Firecrawl scrape (requires BYOK). "parallel" uses the Parallel extract API.

```yaml
{"description": "Which fetch engine to use. \"auto\" (default) uses native if the provider supports it, otherwise Exa. \"native\" forces the provider's built-in fetch. \"exa\" uses Exa Contents API. \"openrouter\" uses direct HTTP fetch. \"firecrawl\" uses Firecrawl scrape (requires BYOK). \"parallel\" uses the Parallel extract API.", "enum": ["auto", "native", "openrouter", "exa", "parallel", "firecrawl"], "example": "auto", "type": "string", "x-speakeasy-unknown-values": "allow"}
```
