---
title: BashServerToolEnvironment
page_id: schema-bashservertoolenvironment-8c981ce8
path: schemas
description: Execution environment for the bash server tool.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# BashServerToolEnvironment

Execution environment for the bash server tool.

```yaml
{"description": "Execution environment for the bash server tool.", "discriminator": {"mapping": {"container_auto": "#/components/schemas/ContainerAutoEnvironment", "container_reference": "#/components/schemas/ContainerReferenceEnvironment"}, "propertyName": "type"}, "example": {"type": "container_auto"}, "oneOf": [{"$ref": "#/components/schemas/ContainerAutoEnvironment"}, {"$ref": "#/components/schemas/ContainerReferenceEnvironment"}]}
```
