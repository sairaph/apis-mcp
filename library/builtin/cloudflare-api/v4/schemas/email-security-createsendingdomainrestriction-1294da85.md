---
title: email-security_CreateSendingDomainRestriction
page_id: schema-email-security-createsendingdomainrestriction-1294da85
path: schemas
description: Create a sending domain restriction.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email-security_CreateSendingDomainRestriction

Create a sending domain restriction.

```yaml
{"description": "Create a sending domain restriction.", "allOf": [{"$ref": "#/components/schemas/email-security_SendingDomainRestriction"}], "required": ["domain", "exclude"]}
```
