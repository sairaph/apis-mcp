---
title: terminal.onboarding_link
page_id: schema-terminal-onboarding-link-1d62166e
path: schemas
description: Returns redirect links used for onboarding onto Tap to Pay on iPhone.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# terminal.onboarding_link

Returns redirect links used for onboarding onto Tap to Pay on iPhone.

```yaml
{"title": "TerminalOnboardingLinkOnboardingLink", "required": ["link_options", "link_type", "object", "redirect_url"], "type": "object", "properties": {"link_options": {"$ref": "#/components/schemas/terminal_onboarding_link_link_options"}, "link_type": {"type": "string", "description": "The type of link being generated.", "enum": ["apple_terms_and_conditions"]}, "object": {"type": "string", "enum": ["terminal.onboarding_link"]}, "on_behalf_of": {"maxLength": 5000, "type": "string", "description": "Stripe account ID to generate the link for.", "nullable": true}, "redirect_url": {"maxLength": 5000, "type": "string", "description": "The link passed back to the user for their onboarding."}}, "description": "Returns redirect links used for onboarding onto Tap to Pay on iPhone.", "x-expandableFields": ["link_options"], "x-resourceId": "terminal.onboarding_link"}
```
