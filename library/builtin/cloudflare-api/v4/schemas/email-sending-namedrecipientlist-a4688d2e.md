---
title: email-sending_NamedRecipientList
page_id: schema-email-sending-namedrecipientlist-a4688d2e
path: schemas
description: Recipient(s). Optional if cc or bcc is provided. A single email string, a named address object, or an array of either.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email-sending_NamedRecipientList

Recipient(s). Optional if cc or bcc is provided. A single email string, a named address object, or an array of either.

```yaml
{"description": "Recipient(s). Optional if cc or bcc is provided. A single email string, a named address object, or an array of either.", "example": ["recipient-a@example.com", {"address": "recipient-b@example.com", "name": "Recipient B"}], "oneOf": [{"$ref": "#/components/schemas/email-sending_EmailAddressString"}, {"$ref": "#/components/schemas/email-sending_EmailAddressObject"}, {"items": {"oneOf": [{"$ref": "#/components/schemas/email-sending_EmailAddressString"}, {"$ref": "#/components/schemas/email-sending_EmailAddressObject"}]}, "type": "array"}]}
```
