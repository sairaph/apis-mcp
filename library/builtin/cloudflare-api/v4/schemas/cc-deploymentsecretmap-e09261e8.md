---
title: cc_DeploymentSecretMap
page_id: schema-cc-deploymentsecretmap-e09261e8
path: schemas
description: Specifies how secrets are accessed in containers, defining the name of the secret within the container and the corresponding account secret name.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cc_DeploymentSecretMap

Specifies how secrets are accessed in containers, defining the name of the secret within the container and the corresponding account secret name.

```yaml
{"description": "Specifies how secrets are accessed in containers, defining the name of the secret within the container and the corresponding account secret name.", "type": "object", "properties": {"name": {"description": "The name of the secret within the container", "type": "string"}, "secret": {"description": "Corresponding secret name from the account", "type": "string"}, "type": {"$ref": "#/components/schemas/cc_SecretAccessType"}}, "required": ["name", "type", "secret"]}
```
