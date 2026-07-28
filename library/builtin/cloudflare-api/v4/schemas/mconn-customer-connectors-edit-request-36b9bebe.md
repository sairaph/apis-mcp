---
title: mconn_customer_connectors_edit_request
page_id: schema-mconn-customer-connectors-edit-request-36b9bebe
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mconn_customer_connectors_edit_request

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/mconn_customer_connector_fields"}, {"properties": {"provision_license": {"description": "When true, regenerate license key for the connector.", "type": "boolean"}}, "type": "object"}]}
```
