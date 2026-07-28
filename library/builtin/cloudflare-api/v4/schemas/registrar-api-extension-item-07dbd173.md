---
title: registrar-api_extension-item
page_id: schema-registrar-api-extension-item-07dbd173
path: schemas
description: Extension entry with metadata and JSON Schema documents for the registration operation.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# registrar-api_extension-item

Extension entry with metadata and JSON Schema documents for the registration operation.

```yaml
{"description": "Extension entry with metadata and JSON Schema documents for the registration operation.", "type": "object", "properties": {"metadata": {"description": "Extension metadata", "type": "object", "properties": {"name": {"description": "The full name of the extension. For example, \"co.uk\", or \"uk\"", "type": "string"}, "tld": {"description": "The tld of the extension. For example, for \"co.uk\", it's \"uk\". For \"uk\", it's \"uk\"", "type": "string"}}, "required": ["name", "tld"]}, "registration_schema": {"description": "JSON Schema describing the expected input structure for registration operations on this extension.", "type": "object"}}, "required": ["metadata", "registration_schema"]}
```
