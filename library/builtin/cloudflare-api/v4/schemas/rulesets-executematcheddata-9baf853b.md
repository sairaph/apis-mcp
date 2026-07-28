---
title: rulesets_ExecuteMatchedData
page_id: schema-rulesets-executematcheddata-9baf853b
path: schemas
description: The configuration to use for matched data logging.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_ExecuteMatchedData

The configuration to use for matched data logging.

```yaml
{"description": "The configuration to use for matched data logging.", "type": "object", "properties": {"public_key": {"description": "The public key to encrypt matched data logs with.", "type": "string", "example": "iGqBmyIUxuWt1rvxoAharN9FUXneUBxA/Y19PyyrEG0=", "minLength": 1, "title": "Public Key"}}, "required": ["public_key"], "title": "Matched Data"}
```
