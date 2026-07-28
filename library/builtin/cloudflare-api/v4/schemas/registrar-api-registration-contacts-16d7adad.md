---
title: registrar-api_registration_contacts
page_id: schema-registrar-api-registration-contacts-16d7adad
path: schemas
description: |-
    Contact data for the registration request.

    The per-extension schema returned by
    `GET /accounts/{account_id}/registrar/extensions/{extension}` is the
    authoritative contract for which contact roles are accepted. Every
    currently supported extension requires only `contacts.registrant` from
    API callers. Additional roles such as `technical`, `administrator`, and
    `billing` may be provided when the extension schema includes them. If a
    registry requires one of those roles and the caller omits it, Cloudflare
    may derive that contact from `contacts.registrant`.

    If the `contacts` object is omitted entirely from the request, or if
    `contacts.registrant` is not provided, the system will use the account's
    default address book entry as the registrant contact. This default must be
    pre-configured by the account owner at
    `https://dash.cloudflare.com/{account_id}/domains/registrations`, where
    they can create or update the address book entry and accept the required
    agreement. No API exists for managing address book entries at this time.

    If no default address book entry exists and no registrant contact is
    provided, the registration request will fail with a validation error.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# registrar-api_registration_contacts

Contact data for the registration request.

The per-extension schema returned by
`GET /accounts/{account_id}/registrar/extensions/{extension}` is the
authoritative contract for which contact roles are accepted. Every
currently supported extension requires only `contacts.registrant` from
API callers. Additional roles such as `technical`, `administrator`, and
`billing` may be provided when the extension schema includes them. If a
registry requires one of those roles and the caller omits it, Cloudflare
may derive that contact from `contacts.registrant`.

If the `contacts` object is omitted entirely from the request, or if
`contacts.registrant` is not provided, the system will use the account's
default address book entry as the registrant contact. This default must be
pre-configured by the account owner at
`https://dash.cloudflare.com/{account_id}/domains/registrations`, where
they can create or update the address book entry and accept the required
agreement. No API exists for managing address book entries at this time.

If no default address book entry exists and no registrant contact is
provided, the registration request will fail with a validation error.

```yaml
{"description": "Contact data for the registration request.\n\nThe per-extension schema returned by\n`GET /accounts/{account_id}/registrar/extensions/{extension}` is the\nauthoritative contract for which contact roles are accepted. Every\ncurrently supported extension requires only `contacts.registrant` from\nAPI callers. Additional roles such as `technical`, `administrator`, and\n`billing` may be provided when the extension schema includes them. If a\nregistry requires one of those roles and the caller omits it, Cloudflare\nmay derive that contact from `contacts.registrant`.\n\nIf the `contacts` object is omitted entirely from the request, or if\n`contacts.registrant` is not provided, the system will use the account's\ndefault address book entry as the registrant contact. This default must be\npre-configured by the account owner at\n`https://dash.cloudflare.com/{account_id}/domains/registrations`, where\nthey can create or update the address book entry and accept the required\nagreement. No API exists for managing address book entries at this time.\n\nIf no default address book entry exists and no registrant contact is\nprovided, the registration request will fail with a validation error.\n", "type": "object", "properties": {"administrator": {"$ref": "#/components/schemas/registrar-api_registration_contact"}, "billing": {"$ref": "#/components/schemas/registrar-api_registration_contact"}, "registrant": {"$ref": "#/components/schemas/registrar-api_registration_contact"}, "technical": {"$ref": "#/components/schemas/registrar-api_registration_contact"}}}
```
