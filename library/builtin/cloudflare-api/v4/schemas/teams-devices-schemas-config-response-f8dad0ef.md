---
title: teams-devices_schemas-config_response
page_id: schema-teams-devices-schemas-config-response-f8dad0ef
path: schemas
description: The configuration object containing information for the WARP client to detect the managed network.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_schemas-config_response

The configuration object containing information for the WARP client to detect the managed network.

```yaml
{"description": "The configuration object containing information for the WARP client to detect the managed network.", "type": "object", "example": {"sha256": "b5bb9d8014a0f9b1d61e21e796d78dccdf1352f23cd32812f4850b878ae4944c", "tls_sockaddr": "foo.bar:1234"}, "oneOf": [{"$ref": "#/components/schemas/teams-devices_tls_config_response"}]}
```
