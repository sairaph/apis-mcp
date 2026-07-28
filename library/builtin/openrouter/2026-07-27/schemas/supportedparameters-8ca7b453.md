---
title: SupportedParameters
page_id: schema-supportedparameters-8ca7b453
path: schemas
description: Union of supported parameters across every endpoint of this model. Coarse discovery aid; the definitive per-endpoint set is behind the endpoints URL.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# SupportedParameters

Union of supported parameters across every endpoint of this model. Coarse discovery aid; the definitive per-endpoint set is behind the endpoints URL.

```yaml
{"additionalProperties": {"$ref": "#/components/schemas/CapabilityDescriptor"}, "description": "Union of supported parameters across every endpoint of this model. Coarse discovery aid; the definitive per-endpoint set is behind the endpoints URL.", "example": {"output_compression": {"max": 100, "min": 0, "type": "range"}, "resolution": {"type": "enum", "values": ["1K", "2K", "4K"]}, "seed": {"type": "boolean"}}, "type": "object"}
```
