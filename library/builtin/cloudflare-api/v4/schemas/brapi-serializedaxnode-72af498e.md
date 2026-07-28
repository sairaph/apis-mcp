---
title: brapi_SerializedAXNode
page_id: schema-brapi-serializedaxnode-72af498e
path: schemas
description: Accessibility tree node
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# brapi_SerializedAXNode

Accessibility tree node

```yaml
{"description": "Accessibility tree node", "type": "object", "properties": {"autocomplete": {"type": "string"}, "checked": {"oneOf": [{"type": "boolean"}, {"enum": ["mixed"], "type": "string"}]}, "children": {"type": "array", "items": {"$ref": "#/components/schemas/brapi_SerializedAXNode"}}, "description": {"type": "string"}, "disabled": {"type": "boolean"}, "expanded": {"type": "boolean"}, "focused": {"type": "boolean"}, "haspopup": {"type": "string"}, "invalid": {"type": "string"}, "keyshortcuts": {"type": "string"}, "level": {"type": "number"}, "modal": {"type": "boolean"}, "multiline": {"type": "boolean"}, "multiselectable": {"type": "boolean"}, "name": {"type": "string"}, "orientation": {"type": "string"}, "pressed": {"oneOf": [{"type": "boolean"}, {"enum": ["mixed"], "type": "string"}]}, "readonly": {"type": "boolean"}, "required": {"type": "boolean"}, "role": {"type": "string"}, "roledescription": {"type": "string"}, "selected": {"type": "boolean"}, "value": {"oneOf": [{"type": "string"}, {"type": "number"}]}, "valuemax": {"type": "number"}, "valuemin": {"type": "number"}, "valuetext": {"type": "string"}}, "required": ["role"]}
```
