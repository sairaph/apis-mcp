---
title: teams-devices_schemas-match
page_id: schema-teams-devices-schemas-match-523764b6
path: schemas
description: 'The wirefilter expression to match devices. Available values: "identity.email", "identity.groups.id", "identity.groups.name", "identity.groups.email", "identity.service_token_uuid", "identity.saml_attributes", "network", "os.name", "os.version".'
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_schemas-match

The wirefilter expression to match devices. Available values: "identity.email", "identity.groups.id", "identity.groups.name", "identity.groups.email", "identity.service_token_uuid", "identity.saml_attributes", "network", "os.name", "os.version".

```yaml
{"description": "The wirefilter expression to match devices. Available values: \"identity.email\", \"identity.groups.id\", \"identity.groups.name\", \"identity.groups.email\", \"identity.service_token_uuid\", \"identity.saml_attributes\", \"network\", \"os.name\", \"os.version\".", "type": "string", "example": "identity.email == \"test@cloudflare.com\"", "maxLength": 500}
```
