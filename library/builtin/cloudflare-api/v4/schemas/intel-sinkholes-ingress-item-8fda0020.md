---
title: intel-sinkholes_ingress_item
page_id: schema-intel-sinkholes-ingress-item-8fda0020
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# intel-sinkholes_ingress_item

```yaml
{"type": "object", "properties": {"cidr": {"description": "The CIDR block for the ingress rule.", "type": "string"}, "created_on": {"description": "The date and time when the ingress rule was created.", "type": "string", "format": "date-time"}, "id": {"description": "The unique identifier for the ingress rule.", "type": "string"}, "modified_on": {"description": "The date and time when the ingress rule was last modified.", "type": "string", "format": "date-time"}, "sinkhole_id": {"description": "The sinkhole this ingress rule belongs to.", "type": "string"}, "zone_tag": {"description": "The zone tag associated with this ingress rule.", "type": "string"}}, "example": {"cidr": "192.0.2.0/24", "created_on": "2023-06-01T10:00:00Z", "id": "de32ae5203724ed08dcc26e971a4d22f", "modified_on": "2023-06-15T14:30:00Z", "sinkhole_id": "93defa6e909e464e8c89a85859f36d3c", "zone_tag": "4c961e9d94f40aa922775483b9ee18cf"}}
```
