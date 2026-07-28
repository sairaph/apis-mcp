---
title: cloudforce-one_MetaInputEntry
page_id: schema-cloudforce-one-metainputentry-495e70a6
path: schemas
description: 'A YARA meta entry. Value type is resolved server-side. Constrained keys: ''detection'' must be one of MALICIOUS, SUSPICIOUS, SPAM, SPOOF.'
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cloudforce-one_MetaInputEntry

A YARA meta entry. Value type is resolved server-side. Constrained keys: 'detection' must be one of MALICIOUS, SUSPICIOUS, SPAM, SPOOF.

```yaml
{"description": "A YARA meta entry. Value type is resolved server-side. Constrained keys: 'detection' must be one of MALICIOUS, SUSPICIOUS, SPAM, SPOOF.", "type": "object", "properties": {"key": {"type": "string", "maxLength": 128, "minLength": 1, "pattern": "^[a-zA-Z_][a-zA-Z0-9_]*$"}, "value": {"anyOf": [{"maxLength": 10000, "type": "string"}, {"type": "number"}, {"type": "boolean"}]}}, "required": ["key", "value"]}
```
