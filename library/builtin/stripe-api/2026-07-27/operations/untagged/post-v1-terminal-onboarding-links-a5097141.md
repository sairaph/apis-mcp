---
title: Create an Onboarding Link
page_id: operation-post-v1-terminal-onboarding-links-699beb8f
path: operations/untagged
description: <p>Creates a new <code>OnboardingLink</code> object that contains a redirect_url used for onboarding onto Tap to Pay on iPhone.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/terminal/onboarding_links
operation_ids:
    - PostTerminalOnboardingLinks
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Create an Onboarding Link

`POST /v1/terminal/onboarding_links`

Operation ID: `PostTerminalOnboardingLinks`

<p>Creates a new <code>OnboardingLink</code> object that contains a redirect_url used for onboarding onto Tap to Pay on iPhone.</p>

## Definition

```yaml
{"summary": "Create an Onboarding Link", "description": "<p>Creates a new <code>OnboardingLink</code> object that contains a redirect_url used for onboarding onto Tap to Pay on iPhone.</p>", "operationId": "PostTerminalOnboardingLinks", "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["link_options", "link_type"], "type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "link_options": {"title": "link_options_params", "type": "object", "properties": {"apple_terms_and_conditions": {"title": "apple_terms_and_conditions_params", "required": ["merchant_display_name"], "type": "object", "properties": {"allow_relinking": {"type": "boolean"}, "merchant_display_name": {"maxLength": 5000, "type": "string"}}}}, "description": "Specific fields needed to generate the desired link type."}, "link_type": {"type": "string", "description": "The type of link being generated.", "enum": ["apple_terms_and_conditions"]}, "on_behalf_of": {"maxLength": 5000, "type": "string", "description": "Stripe account ID to generate the link for."}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "link_options": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/terminal.onboarding_link"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
