---
title: mconn_snapshot_dhcp_lease
page_id: schema-mconn-snapshot-dhcp-lease-66113fa1
path: schemas
description: Snapshot DHCP lease
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mconn_snapshot_dhcp_lease

Snapshot DHCP lease

```yaml
{"description": "Snapshot DHCP lease", "type": "object", "properties": {"client_id": {"description": "Client ID of the device the IP Address was leased to", "type": "string"}, "expiry_time": {"description": "Expiry time of the DHCP lease (seconds since the Unix epoch)", "type": "number"}, "hostname": {"description": "Hostname of the device the IP Address was leased to", "type": "string"}, "interface_name": {"description": "Name of the network interface", "type": "string"}, "ip_address": {"description": "IP Address that was leased", "type": "string"}, "mac_address": {"description": "MAC Address of the device the IP Address was leased to", "type": "string"}}, "required": ["interface_name", "expiry_time", "mac_address", "ip_address", "hostname", "client_id"]}
```
