# UEdGraph

**Parent**: UObject
**Module**: Engine

Blueprint graph structure containing nodes and connections. Each Blueprint can own multiple `UEdGraph` instances representing event graphs, function graphs, macro graphs, and animation state machines. The graph stores an ordered list of `UEdGraphNode` objects whose pins are connected to form the visual script.

## Key Properties

- `Schema` — The `UEdGraphSchema` subclass governing this graph's rules: what nodes are allowed, how pins are typed, and how connections are validated (e.g. `UEdGraphSchema_K2` for Blueprint, `UAnimationGraphSchema` for anim graphs)
- `Nodes` — Array of `UEdGraphNode*` representing every node currently placed in this graph
- `bEditable` — When false the graph is read-only in the editor; used for interface function stubs and auto-generated graphs
- `bAllowDeletion` — When false the graph cannot be deleted by the user; typically set on the default event graph
- `bAllowRenaming` — When false the user cannot rename the graph; used on the event graph and auto-generated function implementations
- `GraphGuid` — Stable `FGuid` identifier for this graph, used for cross-graph node references and diffing
- `SubGraphs` — Array of child `UEdGraph*` embedded within this graph, used by composite nodes and collapsed graphs

## Key Functions

- `AddNode(UEdGraphNode* Node, bool bFromUI, bool bSelectNewNode)` — Adds a node to the graph, fires change notifications, and optionally selects it in the editor
- `RemoveNode(UEdGraphNode* Node)` — Removes a node and its pin connections from the graph and fires change notifications
- `GetAllNodes(TArray<UEdGraphNode*>& Nodes)` — Populates an array with all nodes in this graph and all sub-graphs
- `GetNodesOfClass(TArray<T*>& Nodes)` — Template helper that returns only nodes matching a specific `UEdGraphNode` subclass
- `FindNode(const FGuid& NodeGuid)` — Looks up a node by its stable GUID, returning `nullptr` if not found
- `MoveNodesToAnotherGraph(UEdGraph* DestinationGraph, bool bIsLoading, bool bInIsCompiling)` — Transfers a set of nodes to another graph, updating outer pointers and connection references
- `NotifyGraphChanged()` — Broadcasts the `OnGraphChanged` delegate so listeners (compiler, editor panels) react to structural changes
- `GetSchema()` — Returns the `UEdGraphSchema` for this graph; prefer this over casting `Schema` directly
- `CreateDefaultNodesForGraph(UEdGraph& Graph)` — Asks the schema to insert the default entry and result nodes appropriate for this graph type
- `GetFName()` — Returns the `FName` of this graph, used as the function or event name when compiling
- `GetGoodPlaceForNewNode()` — Computes a `FVector2D` position in graph space that avoids overlapping existing nodes, useful when programmatically adding nodes
