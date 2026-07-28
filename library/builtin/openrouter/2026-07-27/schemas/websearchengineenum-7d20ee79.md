---
title: WebSearchEngineEnum
page_id: schema-websearchengineenum-7d20ee79
path: schemas
description: Which search engine to use. "auto" (default) uses native if the provider supports it, otherwise Exa. "native" forces the provider's built-in search. "exa" forces the Exa search API. "firecrawl" uses Firecrawl (requires BYOK). "parallel" uses the Parallel search API. "perplexity" uses the Perplexity Search API (raw ranked results).
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# WebSearchEngineEnum

Which search engine to use. "auto" (default) uses native if the provider supports it, otherwise Exa. "native" forces the provider's built-in search. "exa" forces the Exa search API. "firecrawl" uses Firecrawl (requires BYOK). "parallel" uses the Parallel search API. "perplexity" uses the Perplexity Search API (raw ranked results).

```yaml
{"description": "Which search engine to use. \"auto\" (default) uses native if the provider supports it, otherwise Exa. \"native\" forces the provider's built-in search. \"exa\" forces the Exa search API. \"firecrawl\" uses Firecrawl (requires BYOK). \"parallel\" uses the Parallel search API. \"perplexity\" uses the Perplexity Search API (raw ranked results).", "enum": ["native", "exa", "parallel", "firecrawl", "perplexity", "auto"], "example": "auto", "type": "string", "x-speakeasy-unknown-values": "allow"}
```
