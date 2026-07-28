---
title: v2.error.configuration_cannot_be_deactivated_due_to_other_capability_requirement
page_id: schema-v2-error-configuration-cannot-be-deactivated-due-to-other-capability-req-b161501c
path: schemas
description: Configuration cannot be deactivated due to a dependency with another capability.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.configuration_cannot_be_deactivated_due_to_other_capability_requirement

Configuration cannot be deactivated due to a dependency with another capability.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "message"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["configuration_cannot_be_deactivated_due_to_other_capability_requirement"]}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}}, "description": "Information about the error that occurred"}}, "description": "Configuration cannot be deactivated due to a dependency with another capability."}
```
