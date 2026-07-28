---
title: AnthropicRefusalStopDetails
page_id: schema-anthropicrefusalstopdetails-f2ed71c3
path: schemas
description: Structured information about a refusal
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AnthropicRefusalStopDetails

Structured information about a refusal

```yaml
{"description": "Structured information about a refusal", "example": {"category": "cyber", "explanation": "The request was refused due to policy.", "type": "refusal"}, "properties": {"category": {"enum": ["cyber", "bio", null], "type": ["string", "null"], "x-speakeasy-unknown-values": "allow"}, "explanation": {"type": ["string", "null"]}, "type": {"enum": ["refusal"], "type": "string"}}, "required": ["type", "category", "explanation"], "type": ["object", "null"]}
```
