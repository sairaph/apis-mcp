---
title: cc_DurableObjectsConfigurationScriptAndClass
page_id: schema-cc-durableobjectsconfigurationscriptandclass-aa85cd08
path: schemas
description: Durable object configuration using script and class names
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cc_DurableObjectsConfigurationScriptAndClass

Durable object configuration using script and class names

```yaml
{"description": "Durable object configuration using script and class names", "type": "object", "properties": {"class_name": {"description": "The class name of the durable object.", "type": "string"}, "script_name": {"description": "The script name where the durable object class is defined.", "type": "string"}}, "required": ["script_name", "class_name"]}
```
