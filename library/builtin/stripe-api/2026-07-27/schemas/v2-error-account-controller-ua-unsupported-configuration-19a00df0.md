---
title: v2.error.account_controller_ua_unsupported_configuration
page_id: schema-v2-error-account-controller-ua-unsupported-configuration-19a00df0
path: schemas
description: Connect integration combination is not supported when UA beta is enabled.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.account_controller_ua_unsupported_configuration

Connect integration combination is not supported when UA beta is enabled.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "doc_url", "invalid_permutation", "message", "user_message"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["account_controller_ua_unsupported_configuration"]}, "doc_url": {"type": "string", "description": "A URL to more information about the error reported"}, "invalid_permutation": {"required": ["dashboard", "fees_collector", "losses_collector", "requirements_collector"], "type": "object", "properties": {"dashboard": {"type": "string", "description": "The value provided for the dashboard field."}, "fees_collector": {"type": "string", "description": "The value provided for the fees_collector field."}, "losses_collector": {"type": "string", "description": "The value provided for the losses_collector field."}, "requirements_collector": {"type": "string", "description": "The value provided for the requirements_collector field."}}, "description": "The invalid permutation provided."}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}, "user_message": {"type": "string", "description": "A user-friendly message that can be shown to end-users"}}, "description": "Information about the error that occurred"}}, "description": "Connect integration combination is not supported when UA beta is enabled."}
```
