---
title: v2.error.capability_cannot_be_unrequested_due_to_other_capability_requirement
page_id: schema-v2-error-capability-cannot-be-unrequested-due-to-other-capability-requir-28791919
path: schemas
description: Feature cannot be unrequested due to being a requirement for another feature.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.capability_cannot_be_unrequested_due_to_other_capability_requirement

Feature cannot be unrequested due to being a requirement for another feature.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "message", "user_message"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["capability_cannot_be_unrequested_due_to_other_capability_requirement"]}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}, "user_message": {"type": "string", "description": "A user-friendly message that can be shown to end-users"}}, "description": "Information about the error that occurred"}}, "description": "Feature cannot be unrequested due to being a requirement for another feature."}
```
