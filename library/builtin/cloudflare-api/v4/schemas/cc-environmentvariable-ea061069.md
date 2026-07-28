---
title: cc_EnvironmentVariable
page_id: schema-cc-environmentvariable-ea061069
path: schemas
description: An environment variable with a value set
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cc_EnvironmentVariable

An environment variable with a value set

```yaml
{"description": "An environment variable with a value set", "type": "object", "properties": {"name": {"$ref": "#/components/schemas/cc_EnvironmentVariableName"}, "value": {"$ref": "#/components/schemas/cc_EnvironmentVariableValue"}}, "required": ["name", "value"]}
```
