---
title: PricingOverride
page_id: schema-pricingoverride-c3dfd78d
path: schemas
description: A conditional override of the base pricing. An entry applies only when all of its condition fields (e.g. min_prompt_tokens, or the utc_start/utc_end time window) match the request; among applicable entries, later entries win per price key; price keys absent from an entry inherit the base price.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# PricingOverride

A conditional override of the base pricing. An entry applies only when all of its condition fields (e.g. min_prompt_tokens, or the utc_start/utc_end time window) match the request; among applicable entries, later entries win per price key; price keys absent from an entry inherit the base price.

```yaml
{"description": "A conditional override of the base pricing. An entry applies only when all of its condition fields (e.g. min_prompt_tokens, or the utc_start/utc_end time window) match the request; among applicable entries, later entries win per price key; price keys absent from an entry inherit the base price.", "example": {"completion": "0.00002", "min_prompt_tokens": 200000, "prompt": "0.000005"}, "properties": {"audio": {"description": "Overridden price in USD per audio input token", "type": "string"}, "completion": {"description": "Overridden price in USD per token for completion (output) generation", "type": "string"}, "input_audio_cache": {"description": "Overridden price in USD per cached audio input token", "type": "string"}, "input_cache_read": {"description": "Overridden price in USD per cached input token (read)", "type": "string"}, "input_cache_write": {"description": "Overridden price in USD per cache-write token", "type": "string"}, "input_cache_write_1h": {"description": "Overridden price in USD per 1-hour cache-write token", "type": "string"}, "min_prompt_tokens": {"description": "Condition: the entry applies when the total prompt tokens of a request are strictly greater than this threshold", "format": "double", "type": "number"}, "prompt": {"description": "Overridden price in USD per token for prompt (input) processing", "type": "string"}, "utc_end": {"description": "Condition: exclusive end of a daily UTC time window as an HHMM clock number (e.g. 400 = 04:00)", "format": "double", "type": "number"}, "utc_start": {"description": "Condition: inclusive start of a daily UTC time window as an HHMM clock number (e.g. 100 = 01:00, 1030 = 10:30). The entry applies while the current UTC time is inside the half-open window [utc_start, utc_end), which may wrap past midnight (utc_start > utc_end).", "format": "double", "type": "number"}}, "type": "object"}
```
