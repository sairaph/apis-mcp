---
title: ContainerReferenceEnvironment
page_id: schema-containerreferenceenvironment-547618c8
path: schemas
description: Reference to a previously created container to reuse.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ContainerReferenceEnvironment

Reference to a previously created container to reuse.

```yaml
{"description": "Reference to a previously created container to reuse.", "example": {"container_id": "cntr_abc123", "type": "container_reference"}, "properties": {"container_id": {"description": "Identifier of an existing container to reuse (max 20 characters).", "example": "cntr_abc123", "maxLength": 20, "minLength": 1, "pattern": "^[\\w-]+$", "type": "string"}, "type": {"enum": ["container_reference"], "type": "string"}}, "required": ["type", "container_id"], "type": "object"}
```
