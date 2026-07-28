---
title: v2.error.individual_additional_person_not_allowed
page_id: schema-v2-error-individual-additional-person-not-allowed-2bce8bf9
path: schemas
description: Additional person is added for an individual business type.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.individual_additional_person_not_allowed

Additional person is added for an individual business type.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "message"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["individual_additional_person_not_allowed"]}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}}, "description": "Information about the error that occurred"}}, "description": "Additional person is added for an individual business type."}
```
