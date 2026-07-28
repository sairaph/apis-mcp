---
title: cc_ApplicationVersion
page_id: schema-cc-applicationversion-86915088
path: schemas
description: An application with the configuration of its version
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cc_ApplicationVersion

An application with the configuration of its version

```yaml
{"description": "An application with the configuration of its version", "type": "object", "properties": {"configuration": {"$ref": "#/components/schemas/cc_ModifyUserDeploymentConfiguration"}, "percentage": {"type": "integer"}, "version": {"type": "integer"}}, "required": ["configuration", "percentage", "version"]}
```
