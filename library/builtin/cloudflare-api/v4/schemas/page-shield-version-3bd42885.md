---
title: page-shield_version
page_id: schema-page-shield-version-3bd42885
path: schemas
description: The version of the analyzed script.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# page-shield_version

The version of the analyzed script.

```yaml
{"description": "The version of the analyzed script.", "type": "object", "properties": {"cryptomining_score": {"$ref": "#/components/schemas/page-shield_cryptomining_score"}, "dataflow_score": {"allOf": [{"$ref": "#/components/schemas/page-shield_dataflow_score"}], "deprecated": true}, "fetched_at": {"$ref": "#/components/schemas/page-shield_fetched_at"}, "hash": {"$ref": "#/components/schemas/page-shield_hash"}, "js_integrity_score": {"$ref": "#/components/schemas/page-shield_js_integrity_score"}, "magecart_score": {"$ref": "#/components/schemas/page-shield_magecart_score"}, "malware_score": {"$ref": "#/components/schemas/page-shield_malware_score"}, "obfuscation_score": {"allOf": [{"$ref": "#/components/schemas/page-shield_obfuscation_score"}], "deprecated": true}}}
```
