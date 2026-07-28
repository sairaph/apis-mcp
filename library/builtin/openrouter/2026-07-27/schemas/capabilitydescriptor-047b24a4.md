---
title: CapabilityDescriptor
page_id: schema-capabilitydescriptor-047b24a4
path: schemas
description: A typed descriptor for one supported request parameter.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# CapabilityDescriptor

A typed descriptor for one supported request parameter.

```yaml
{"description": "A typed descriptor for one supported request parameter.", "discriminator": {"mapping": {"boolean": "#/components/schemas/BooleanCapability", "enum": "#/components/schemas/EnumCapability", "range": "#/components/schemas/RangeCapability", "x-speakeasy-unknown-values": "allow"}, "propertyName": "type"}, "example": {"type": "enum", "values": ["1K", "2K", "4K"]}, "oneOf": [{"$ref": "#/components/schemas/EnumCapability"}, {"$ref": "#/components/schemas/RangeCapability"}, {"$ref": "#/components/schemas/BooleanCapability"}]}
```
