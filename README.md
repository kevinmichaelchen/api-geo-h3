# api-geo-h3
This is a proof-of-concept component that wraps [H3](https://h3geo.org/),
a hexagonal hierarchical geospatial indexing system.

This service needs to have the following capabilities:
* [`geotoh3`](https://h3geo.org/docs/api/indexing/#geotoh3) — convert lat/lng to hex cell (step 2)
* [`kring`](https://h3geo.org/docs/api/traversal/#kring) — get a cell's neighboring cells (step 4)

<img src="./dispatch.png" />
