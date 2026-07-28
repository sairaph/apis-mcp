---
title: PublicPricing
page_id: schema-publicpricing-745b6dda
path: schemas
description: Pricing information for the model
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# PublicPricing

Pricing information for the model

```yaml
{"description": "Pricing information for the model", "example": {"completion": "0.00006", "image": "0", "prompt": "0.00003", "request": "0"}, "properties": {"audio": {"description": "Price in USD per audio input token", "type": "string"}, "audio_output": {"description": "Price in USD per audio output token", "type": "string"}, "completion": {"description": "Price in USD per token for completion (output) generation", "type": "string"}, "discount": {"description": "Fractional discount applied to this endpoint's pricing; the price is multiplied by (1 - discount) (0 = no discount, 1 = free)", "format": "double", "type": "number"}, "image": {"description": "Price in USD per input image", "type": "string"}, "image_output": {"description": "Price in USD per output image", "type": "string"}, "image_token": {"description": "Price in USD per image token", "type": "string"}, "input_audio_cache": {"description": "Price in USD per cached audio input token", "type": "string"}, "input_cache_read": {"description": "Price in USD per cached input token (read)", "type": "string"}, "input_cache_write": {"description": "Price per cache-write token, in USD per token. For providers with multiple cache TTLs (e.g. Anthropic), this is the default (5-minute) cache-write rate.", "type": "string"}, "input_cache_write_1h": {"description": "Price per 1-hour cache-write token, in USD per token. Only present for providers that price an extended (1-hour) cache TTL separately, such as Anthropic.", "type": "string"}, "internal_reasoning": {"description": "Price in USD per internal reasoning token", "type": "string"}, "overrides": {"description": "Conditional overrides of the base pricing (e.g. long-context or time-based pricing). An entry applies when all of its condition fields (e.g. min_prompt_tokens, or the utc_start/utc_end time window) match the request; among applicable entries, later entries win per key; price keys absent from an entry inherit the base price. The top-level pricing keys always reflect the price that applies under default conditions.", "items": {"$ref": "#/components/schemas/PricingOverride"}, "type": "array"}, "prompt": {"description": "Price in USD per token for prompt (input) processing", "type": "string"}, "request": {"description": "Price in USD per request", "type": "string"}, "web_search": {"description": "Price in USD per web search", "type": "string"}}, "required": ["prompt", "completion"], "type": "object"}
```
