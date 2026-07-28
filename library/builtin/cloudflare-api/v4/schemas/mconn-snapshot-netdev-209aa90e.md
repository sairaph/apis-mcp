---
title: mconn_snapshot_netdev
page_id: schema-mconn-snapshot-netdev-209aa90e
path: schemas
description: Snapshot Netdev
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mconn_snapshot_netdev

Snapshot Netdev

```yaml
{"description": "Snapshot Netdev", "type": "object", "properties": {"name": {"description": "Name of the network device", "type": "string"}, "recv_bytes": {"description": "Total bytes received", "type": "number"}, "recv_compressed": {"description": "Compressed packets received", "type": "number"}, "recv_drop": {"description": "Packets dropped", "type": "number"}, "recv_errs": {"description": "Bad packets received", "type": "number"}, "recv_fifo": {"description": "FIFO overruns", "type": "number"}, "recv_frame": {"description": "Frame alignment errors", "type": "number"}, "recv_multicast": {"description": "Multicast packets received", "type": "number"}, "recv_packets": {"description": "Total packets received", "type": "number"}, "sent_bytes": {"description": "Total bytes transmitted", "type": "number"}, "sent_carrier": {"description": "Number of packets not sent due to carrier errors", "type": "number"}, "sent_colls": {"description": "Number of collisions", "type": "number"}, "sent_compressed": {"description": "Number of compressed packets transmitted", "type": "number"}, "sent_drop": {"description": "Number of packets dropped during transmission", "type": "number"}, "sent_errs": {"description": "Number of transmission errors", "type": "number"}, "sent_fifo": {"description": "FIFO overruns", "type": "number"}, "sent_packets": {"description": "Total packets transmitted", "type": "number"}}, "required": ["name", "recv_bytes", "recv_packets", "recv_errs", "recv_drop", "recv_fifo", "recv_frame", "recv_compressed", "recv_multicast", "sent_bytes", "sent_packets", "sent_errs", "sent_drop", "sent_fifo", "sent_colls", "sent_carrier", "sent_compressed"]}
```
