---
title: vuln_scanner_bola-variable-capture
page_id: schema-vuln-scanner-bola-variable-capture-ab560cbe
path: schemas
description: A variable to capture from the response body.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# vuln_scanner_bola-variable-capture

A variable to capture from the response body.

```yaml
{"description": "A variable to capture from the response body.", "type": "object", "properties": {"json_path": {"description": "JSONPath expression used for capture, e.g. `\"$.id\"`.", "type": "string"}, "name": {"description": "Variable name, e.g. `\"resource_id\"`.", "type": "string"}}, "required": ["name", "json_path"]}
```
