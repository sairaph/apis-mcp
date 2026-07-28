---
title: dns-records_LOCRecord
page_id: schema-dns-records-locrecord-c7f18d77
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dns-records_LOCRecord

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-records_dns-record-shared-fields"}, {"properties": {"content": {"description": "Formatted LOC content. See 'data' to set LOC properties.", "example": "IN LOC 37 46 46 N 122 23 35 W 0m 100m 0m 0m", "readOnly": true, "type": "string", "x-auditable": true}, "data": {"description": "Components of a LOC record.", "type": "object", "properties": {"altitude": {"description": "Altitude of location in meters.", "type": "number", "example": 0, "maximum": 42849672.95, "minimum": -100000, "x-auditable": true}, "lat_degrees": {"description": "Degrees of latitude.", "type": "number", "example": 37, "maximum": 90, "minimum": 0, "x-auditable": true}, "lat_direction": {"description": "Latitude direction.", "type": "string", "example": "N", "enum": ["N", "S"], "x-auditable": true}, "lat_minutes": {"description": "Minutes of latitude.", "type": "number", "example": 46, "maximum": 59, "minimum": 0, "x-auditable": true}, "lat_seconds": {"description": "Seconds of latitude.", "type": "number", "example": 46, "maximum": 59.999, "minimum": 0, "x-auditable": true}, "long_degrees": {"description": "Degrees of longitude.", "type": "number", "example": 122, "maximum": 180, "minimum": 0, "x-auditable": true}, "long_direction": {"description": "Longitude direction.", "type": "string", "example": "W", "enum": ["E", "W"], "x-auditable": true}, "long_minutes": {"description": "Minutes of longitude.", "type": "number", "example": 23, "maximum": 59, "minimum": 0, "x-auditable": true}, "long_seconds": {"description": "Seconds of longitude.", "type": "number", "example": 35, "maximum": 59.999, "minimum": 0, "x-auditable": true}, "precision_horz": {"description": "Horizontal precision of location.", "type": "number", "example": 0, "maximum": 90000000, "minimum": 0, "x-auditable": true}, "precision_vert": {"description": "Vertical precision of location.", "type": "number", "example": 0, "maximum": 90000000, "minimum": 0, "x-auditable": true}, "size": {"description": "Size of location in meters.", "type": "number", "example": 100, "maximum": 90000000, "minimum": 0, "x-auditable": true}}}, "type": {"description": "Record type.", "type": "string", "example": "LOC", "enum": ["LOC"], "x-auditable": true}}, "type": "object"}], "title": "LOC Record"}
```
