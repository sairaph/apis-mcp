---
title: issuing_authorization_network_data
page_id: schema-issuing-authorization-network-data-8f90d423
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# issuing_authorization_network_data

```yaml
{"title": "IssuingAuthorizationNetworkData", "type": "object", "properties": {"acquiring_institution_id": {"maxLength": 5000, "type": "string", "description": "Identifier assigned to the acquirer by the card network. Sometimes this value is not provided by the network; in this case, the value will be `null`.", "nullable": true}, "system_trace_audit_number": {"maxLength": 5000, "type": "string", "description": "The System Trace Audit Number (STAN) is a 6-digit identifier assigned by the acquirer. Prefer `network_data.transaction_id` if present, unless you have special requirements.", "nullable": true}, "transaction_id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the authorization assigned by the card network used to match subsequent messages, disputes, and transactions.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
