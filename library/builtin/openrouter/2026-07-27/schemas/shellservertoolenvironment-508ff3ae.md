---
title: ShellServerToolEnvironment
page_id: schema-shellservertoolenvironment-508ff3ae
path: schemas
description: Server-side execution environment for the shell tool. Only container-backed environments are supported; "local" shells are not.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ShellServerToolEnvironment

Server-side execution environment for the shell tool. Only container-backed environments are supported; "local" shells are not.

```yaml
{"description": "Server-side execution environment for the shell tool. Only container-backed environments are supported; \"local\" shells are not.", "discriminator": {"mapping": {"container_auto": "#/components/schemas/ContainerAutoEnvironment", "container_reference": "#/components/schemas/ContainerReferenceEnvironment"}, "propertyName": "type"}, "example": {"type": "container_auto"}, "oneOf": [{"$ref": "#/components/schemas/ContainerAutoEnvironment"}, {"$ref": "#/components/schemas/ContainerReferenceEnvironment"}]}
```
