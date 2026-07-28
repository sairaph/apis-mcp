---
title: teams-devices_ip_profile_match
page_id: schema-teams-devices-ip-profile-match-f5fac2ce
path: schemas
description: 'The wirefilter expression to match registrations. Available values: "identity.name", "identity.email", "identity.groups.id", "identity.groups.name", "identity.groups.email", "identity.saml_attributes".'
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_ip_profile_match

The wirefilter expression to match registrations. Available values: "identity.name", "identity.email", "identity.groups.id", "identity.groups.name", "identity.groups.email", "identity.saml_attributes".

```yaml
{"description": "The wirefilter expression to match registrations. Available values: \"identity.name\", \"identity.email\", \"identity.groups.id\", \"identity.groups.name\", \"identity.groups.email\", \"identity.saml_attributes\".", "type": "string", "example": "identity.email == \"test@cloudflare.com\"", "maxLength": 10000, "x-auditable": true}
```
