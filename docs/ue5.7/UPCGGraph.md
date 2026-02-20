# UPCGGraph

**Parent**: UObject
**Module**: PCG

PCG graph asset containing nodes that define procedural generation logic. A `UPCGGraph` is the authored representation of a procedural pipeline — a directed acyclic graph of `UPCGNode` objects whose edges carry typed `FPCGData` between them. It is referenced by `UPCGComponent` instances and executed by the PCG subsystem.

## Key Properties

- `InputNode` — The special entry node of type `UPCGNode` that provides spatial input data (actor bounds, landscape, point clouds) from the owning `UPCGComponent`
- `OutputNode` — The special exit node that collects the final data produced by the graph and hands it back to the component for spawning or property application
- `bLandscapeUsesMetadata` — When true, landscape sampling nodes include per-layer metadata in their output point data for use in downstream filter and transform nodes
- `HiGenGridSize` — Grid cell size used for hierarchical generation; controls how the PCG subsystem partitions large worlds for streaming and incremental updates
- `bUseHierarchicalGeneration` — Enables world-partition-aware generation that splits execution into grid cells matching `HiGenGridSize`; required for large open worlds

## Key Functions

- `AddNode(UPCGSettings* Settings)` — Creates a new `UPCGNode` wrapping the given settings object, adds it to the graph, and returns a pointer to the new node
- `RemoveNode(UPCGNode* Node)` — Detaches all edges connected to the node and removes it from the graph; does not delete the node's settings object
- `GetNodes()` — Returns the full array of `UPCGNode*` in this graph excluding the input and output nodes
- `GetInputNode()` — Returns the graph's dedicated input node; shorthand for accessing `InputNode`
- `GetOutputNode()` — Returns the graph's dedicated output node; shorthand for accessing `OutputNode`
- `AddEdge(UPCGNode* From, FName FromLabel, UPCGNode* To, FName ToLabel)` — Connects an output pin named `FromLabel` on `From` to an input pin named `ToLabel` on `To`; returns true on success, false if the pin types are incompatible
- `RemoveEdge(UPCGNode* From, FName FromLabel, UPCGNode* To, FName ToLabel)` — Disconnects a specific edge by source node, source pin label, destination node, and destination pin label
- `AddLabeledEdge(UPCGNode* From, FName FromLabel, UPCGNode* To, FName ToLabel)` — Variant of `AddEdge` that creates a labeled data channel edge, used for named data routing between nodes
- `GetEdges()` — Returns the array of `FPCGEdge` structs describing all connections in the graph
- `ForceNotificationOfChange()` — Marks the graph dirty and broadcasts change notifications to all registered listeners including dependent component instances
- `Compile()` — Validates node connectivity and pin types, computes execution order, and bakes the result into an optimized representation used by the runtime executor
