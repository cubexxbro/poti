# Engineering Architecture: Data Ingestion & Performance

This document outlines the performance characteristics of `poti` and the mechanisms governing its operational speed.

## 1. The Ingestion Pipeline
`poti` operates on a three-stage asynchronous pipeline to convert raw CIDR registries into actionable network reconnaissance targets:

1. **Network Sync Stage**: Fetches remote subnet registries (e.g., `country-ip-blocks`). 
    * *Bottleneck*: Governed by upstream bandwidth and global routing latency.
    * *Optimization*: Uses `io.TeeReader` to provide real-time stream feedback, ensuring operators are informed during packet loss or low-throughput scenarios.

2. **Lattice Formulation (The "Heavy" Stage)**:
    * *Process*: Parsing CIDR blocks (e.g., `192.168.0.0/16`) into individual IP nodes within memory.
    * *Why it can be slow*: This is a high-CPU, O(n) operation. Parsing millions of nodes requires extensive bitwise masking and memory allocation.
    * *User Feedback*: In `v0.7.3`, we implemented a real-time terminal progress counter (`[i/total]`) to prevent perceived application hanging.

3. **Concurrency Engine**: 
    * *Execution*: Distributes tasks across 30+ asynchronous worker routines.
    * *Advantage*: Effectively hides I/O wait times, allowing parallel TCP handshake evaluation.

## 2. Performance Tuning & Best Practices

To maximize throughput, operators should consider the following environment variables:

* **Caching (Persistent Layer)**: 
    * Once a country registry is fetched, it is stored in `~/.poti_cache`. Subsequent runs will bypass the network sync stage entirely, reducing startup latency from "seconds" to "milliseconds".
    * *Maintenance*: Use the `clear` command to purge stale cache data and force a fresh registry sync.

* **Network Environment**:
    * Execution speed is correlated with the node's proximity to the primary registry hosts (typically GitHub-based raw endpoints). Using a proxy or high-speed data path can significantly reduce Sync Stage duration.

## 3. Future Scaling Strategies
As the global asset registry continues to expand, our roadmap includes:
* **Binary Indexing**: Transitioning from text-based CIDR files to compressed binary maps for faster memory-mapped I/O.
* **Predictive Prefetching**: Background caching of commonly used regions.
