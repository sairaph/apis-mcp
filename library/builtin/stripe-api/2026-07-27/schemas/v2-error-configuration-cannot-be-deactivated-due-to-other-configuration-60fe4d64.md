---
title: v2.error.configuration_cannot_be_deactivated_due_to_other_configuration
page_id: schema-v2-error-configuration-cannot-be-deactivated-due-to-other-configuration-60fe4d64
path: schemas
description: Cannot deactivate a configuration due to another configuration depending on it.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.configuration_cannot_be_deactivated_due_to_other_configuration

Cannot deactivate a configuration due to another configuration depending on it.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "message"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["configuration_cannot_be_deactivated_due_to_other_configuration"]}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}}, "description": "Information about the error that occurred"}}, "description": "Cannot deactivate a configuration due to another configuration depending on it."}
```
