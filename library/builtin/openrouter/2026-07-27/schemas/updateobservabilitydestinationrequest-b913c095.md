---
title: UpdateObservabilityDestinationRequest
page_id: schema-updateobservabilitydestinationrequest-b913c095
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# UpdateObservabilityDestinationRequest

```yaml
{"example": {"enabled": false, "name": "Updated Langfuse"}, "properties": {"api_key_hashes": {"description": "Optional allowlist of OpenRouter API key hashes. `null` clears the filter (all keys). Omitting leaves the current value. Must contain at least one hash if provided.", "example": null, "items": {"type": "string"}, "minItems": 1, "type": ["array", "null"]}, "config": {"additionalProperties": {}, "description": "Provider-specific configuration fields to update. Masked values are ignored; unset fields keep their current value.", "example": {"baseUrl": "https://us.cloud.langfuse.com", "publicKey": "pk-l...EfGh", "secretKey": "sk-l...AbCd"}, "type": "object"}, "enabled": {"description": "Whether the destination is enabled.", "example": true, "type": "boolean"}, "filter_rules": {"allOf": [{"$ref": "#/components/schemas/ObservabilityFilterRulesConfigNullable"}, {"description": "Optional structured filter rules. `null` clears the rules. Omitting keeps the current value."}]}, "name": {"description": "Human-readable name for the destination.", "example": "Production Langfuse", "type": "string"}, "privacy_mode": {"description": "When true, request/response bodies are not forwarded — only metadata.", "example": false, "type": "boolean"}, "sampling_rate": {"description": "Sampling rate between 0.0001 and 1 (1 = 100%).", "example": 1, "format": "double", "type": "number"}}, "type": "object"}
```
