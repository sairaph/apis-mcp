---
title: terminal_onboarding_link_link_options
page_id: schema-terminal-onboarding-link-link-options-db38dc8e
path: schemas
description: Link type options associated with the current onboarding link object.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# terminal_onboarding_link_link_options

Link type options associated with the current onboarding link object.

```yaml
{"title": "TerminalOnboardingLinkLinkOptions", "type": "object", "properties": {"apple_terms_and_conditions": {"description": "The options associated with the Apple Terms and Conditions link type.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/terminal_onboarding_link_apple_terms_and_conditions"}]}}, "description": "Link type options associated with the current onboarding link object.", "x-expandableFields": ["apple_terms_and_conditions"]}
```
