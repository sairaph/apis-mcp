---
title: mconn_customer_connector
page_id: schema-mconn-customer-connector-05b3a092
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mconn_customer_connector

```yaml
{"type": "object", "properties": {"activated": {"type": "boolean", "x-auditable": true, "x-stainless-terraform-configurability": "computed_optional"}, "device": {"$ref": "#/components/schemas/mconn_customer_device"}, "id": {"$ref": "#/components/schemas/mconn_uuid"}, "interrupt_window_days_of_week": {"description": "Allowed days of the week for upgrades. Default is all days.", "type": "array", "items": {"$ref": "#/components/schemas/mconn_day_of_week"}, "x-auditable": true, "x-stainless-terraform-configurability": "computed_optional"}, "interrupt_window_duration_hours": {"type": "number", "maximum": 24, "minimum": 1, "x-auditable": true, "x-stainless-terraform-configurability": "computed_optional"}, "interrupt_window_embargo_dates": {"description": "List of dates (YYYY-MM-DD) when upgrades are blocked.", "type": "array", "items": {"$ref": "#/components/schemas/mconn_embargo_date"}, "maxItems": 100, "x-auditable": true, "x-stainless-terraform-configurability": "computed_optional"}, "interrupt_window_hour_of_day": {"type": "number", "x-auditable": true, "x-stainless-terraform-configurability": "computed_optional"}, "last_heartbeat": {"type": "string", "x-auditable": true}, "last_seen_version": {"type": "string", "x-auditable": true}, "last_updated": {"type": "string", "x-auditable": true}, "license_key": {"type": "string", "x-stainless-terraform-configurability": "computed_optional"}, "notes": {"type": "string", "x-auditable": true, "x-stainless-terraform-configurability": "computed_optional"}, "primary": {"type": "boolean", "default": true, "x-auditable": true, "x-stainless-terraform-configurability": "computed_optional"}, "site_id": {"type": "string", "x-auditable": true, "x-stainless-terraform-configurability": "computed_optional"}, "timezone": {"type": "string", "x-auditable": true, "x-stainless-terraform-configurability": "computed_optional"}}, "required": ["id", "notes", "activated", "last_updated", "timezone", "interrupt_window_hour_of_day", "interrupt_window_duration_hours", "interrupt_window_days_of_week", "interrupt_window_embargo_dates", "primary"]}
```
