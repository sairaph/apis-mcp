---
title: terminal_onboarding_link_apple_terms_and_conditions
page_id: schema-terminal-onboarding-link-apple-terms-and-conditions-a17cc085
path: schemas
description: Options associated with the Apple Terms and Conditions link type.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# terminal_onboarding_link_apple_terms_and_conditions

Options associated with the Apple Terms and Conditions link type.

```yaml
{"title": "TerminalOnboardingLinkAppleTermsAndConditions", "required": ["merchant_display_name"], "type": "object", "properties": {"allow_relinking": {"type": "boolean", "description": "Whether the link should also support users relinking their Apple account.", "nullable": true}, "merchant_display_name": {"maxLength": 5000, "type": "string", "description": "The business name of the merchant accepting Apple's Terms and Conditions."}}, "description": "Options associated with the Apple Terms and Conditions link type.", "x-expandableFields": []}
```
