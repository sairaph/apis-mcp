---
title: ORAnthropicNullableCaller
page_id: schema-oranthropicnullablecaller-3b86aaa3
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ORAnthropicNullableCaller

```yaml
{"discriminator": {"mapping": {"code_execution_20250825": "#/components/schemas/AnthropicCodeExecution20250825Caller", "code_execution_20260120": "#/components/schemas/AnthropicCodeExecution20260120Caller", "direct": "#/components/schemas/AnthropicDirectCaller"}, "propertyName": "type"}, "example": {"type": "direct"}, "oneOf": [{"$ref": "#/components/schemas/AnthropicDirectCaller"}, {"$ref": "#/components/schemas/AnthropicCodeExecution20250825Caller"}, {"$ref": "#/components/schemas/AnthropicCodeExecution20260120Caller"}, {"type": "null"}]}
```
