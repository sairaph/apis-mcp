---
title: account
page_id: schema-account-9af21132
path: schemas
description: |-
    For new integrations, we recommend using the [Accounts v2 API](/api/v2/core/accounts), in place of /v1/accounts and /v1/customers to represent a user.

    This is an object representing a Stripe account. You can retrieve it to see
    properties on the account like its current requirements or if the account is
    enabled to make live charges or receive payouts.

    For accounts where [controller.requirement_collection](/api/accounts/object#account_object-controller-requirement_collection)
    is `application`, which includes Custom accounts, the properties below are always
    returned.

    For accounts where [controller.requirement_collection](/api/accounts/object#account_object-controller-requirement_collection)
    is `stripe`, which includes Standard and Express accounts, some properties are only returned
    until you create an [Account Link](/api/account_links) or [Account Session](/api/account_sessions)
    to start Connect Onboarding. Learn about the [differences between accounts](/connect/accounts).
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# account

For new integrations, we recommend using the [Accounts v2 API](/api/v2/core/accounts), in place of /v1/accounts and /v1/customers to represent a user.

This is an object representing a Stripe account. You can retrieve it to see
properties on the account like its current requirements or if the account is
enabled to make live charges or receive payouts.

For accounts where [controller.requirement_collection](/api/accounts/object#account_object-controller-requirement_collection)
is `application`, which includes Custom accounts, the properties below are always
returned.

For accounts where [controller.requirement_collection](/api/accounts/object#account_object-controller-requirement_collection)
is `stripe`, which includes Standard and Express accounts, some properties are only returned
until you create an [Account Link](/api/account_links) or [Account Session](/api/account_sessions)
to start Connect Onboarding. Learn about the [differences between accounts](/connect/accounts).

```yaml
{"title": "Account", "required": ["id", "object"], "type": "object", "properties": {"business_profile": {"description": "Business information about the account.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/account_business_profile"}]}, "business_type": {"type": "string", "description": "The business type.", "nullable": true, "enum": ["company", "government_entity", "individual", "non_profit"], "x-stripeBypassValidation": true}, "capabilities": {"$ref": "#/components/schemas/account_capabilities"}, "charges_enabled": {"type": "boolean", "description": "Whether the account can process charges."}, "company": {"$ref": "#/components/schemas/legal_entity_company"}, "controller": {"$ref": "#/components/schemas/account_unification_account_controller"}, "country": {"maxLength": 5000, "type": "string", "description": "The account's country."}, "created": {"type": "integer", "description": "Time at which the account was connected. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "default_currency": {"maxLength": 5000, "type": "string", "description": "Three-letter ISO currency code representing the default currency for the account. This must be a currency that [Stripe supports in the account's country](https://stripe.com/docs/payouts)."}, "details_submitted": {"type": "boolean", "description": "Whether account details have been submitted. Accounts with Stripe Dashboard access, which includes Standard accounts, cannot receive payouts before this is true. Accounts where this is false should be directed to [an onboarding flow](/connect/onboarding) to finish submitting account details."}, "email": {"maxLength": 5000, "type": "string", "description": "An email address associated with the account. It's not used for authentication and Stripe doesn't market to this field without explicit approval from the platform.", "nullable": true}, "external_accounts": {"title": "ExternalAccountList", "required": ["data", "has_more", "object", "url"], "type": "object", "properties": {"data": {"type": "array", "description": "The list contains all external accounts that have been attached to the Stripe account. These may be bank accounts or cards.", "items": {"title": "Polymorphic", "anyOf": [{"$ref": "#/components/schemas/bank_account"}, {"$ref": "#/components/schemas/card"}], "x-stripeBypassValidation": true}}, "has_more": {"type": "boolean", "description": "True if this list has another page of items after this one that can be fetched."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value. Always has the value `list`.", "enum": ["list"]}, "url": {"maxLength": 5000, "type": "string", "description": "The URL where this list can be accessed."}}, "description": "External accounts (bank accounts and debit cards) currently attached to this account. External accounts are only returned for requests where `controller[is_controller]` is true.", "x-expandableFields": ["data"]}, "future_requirements": {"$ref": "#/components/schemas/account_future_requirements"}, "groups": {"description": "The groups associated with the account.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/account_group_membership"}]}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "individual": {"$ref": "#/components/schemas/person"}, "metadata": {"type": "object", "additionalProperties": {"maxLength": 500, "type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["account"]}, "payouts_enabled": {"type": "boolean", "description": "Whether the funds in this account can be paid out."}, "requirements": {"$ref": "#/components/schemas/account_requirements"}, "settings": {"description": "Options for customizing how the account functions within Stripe.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/account_settings"}]}, "tos_acceptance": {"$ref": "#/components/schemas/account_tos_acceptance"}, "type": {"type": "string", "description": "The Stripe account type. Can be `standard`, `express`, `custom`, or `none`.", "enum": ["custom", "express", "none", "standard"]}}, "description": "For new integrations, we recommend using the [Accounts v2 API](/api/v2/core/accounts), in place of /v1/accounts and /v1/customers to represent a user.\n\nThis is an object representing a Stripe account. You can retrieve it to see\nproperties on the account like its current requirements or if the account is\nenabled to make live charges or receive payouts.\n\nFor accounts where [controller.requirement_collection](/api/accounts/object#account_object-controller-requirement_collection)\nis `application`, which includes Custom accounts, the properties below are always\nreturned.\n\nFor accounts where [controller.requirement_collection](/api/accounts/object#account_object-controller-requirement_collection)\nis `stripe`, which includes Standard and Express accounts, some properties are only returned\nuntil you create an [Account Link](/api/account_links) or [Account Session](/api/account_sessions)\nto start Connect Onboarding. Learn about the [differences between accounts](/connect/accounts).", "x-expandableFields": ["business_profile", "capabilities", "company", "controller", "external_accounts", "future_requirements", "groups", "individual", "requirements", "settings", "tos_acceptance"], "x-resourceId": "account"}
```
