# UPCGSettings

**Parent**: UPCGSettingsInterface
**Module**: PCG

Base class for PCG node settings that defines node behavior. Every `UPCGNode` in a `UPCGGraph` owns a `UPCGSettings` subclass instance that describes what the node does, what data it consumes and produces, and how it is displayed in the editor. Subclass this to implement custom PCG nodes.

## Key Properties

- `bEnabled` — When false the node is bypassed during execution: its input data passes directly to output without processing, equivalent to a no-op pass-through
- `bDebug` — When true the PCG debugger visualizes this node's output in the editor viewport as colored points or meshes, useful during graph authoring
- `CachedOverridableParams` — Array of `FPCGSettingsOverridableParam` describing which properties on this settings object can be overridden at runtime via PCG attribute sets or Blueprint
- `FilterOnTags` — Set of `FName` tags; the node only processes input data elements that carry all listed tags, enabling selective filtering without explicit filter nodes
- `PassThroughOnTags` — Set of `FName` tags; input elements carrying these tags skip this node's processing and flow directly to output unchanged
- `DeterminismSettings` — `FPCGDeterminismSettings` struct controlling how the node handles non-deterministic operations (ordering, floating point, randomness) for repeatable generation
- `bExposeToLibrary` — When true this settings class appears in the PCG node palette under its category, making it available for graph authors to place

## Key Functions

- `GetDefaultNodeTitle()` — Returns the `FText` display name shown on the node tile in the graph editor; override to provide a custom name
- `GetNodeTooltipText()` — Returns the `FText` tooltip shown on hover in the graph editor; should describe what the node does and its expected inputs
- `GetInputPinLabel(uint32 Index)` — Returns the `FName` label for the input pin at `Index`; used by the graph editor and edge routing logic
- `GetOutputPinLabel(uint32 Index)` — Returns the `FName` label for the output pin at `Index`
- `GetInputPinTypes()` — Returns an `EPCGDataType` bitmask describing the data types accepted on each input pin; used for connection validation
- `GetOutputPinTypes()` — Returns an `EPCGDataType` bitmask describing the data types produced on each output pin
- `UseSeed()` — Returns true if this node uses the PCG seed for randomness; when true the node is re-executed when the component seed changes
- `GetSeed(const UPCGComponent* Component)` — Computes the effective integer seed for this node by combining the component seed with a node-specific offset, ensuring unique random streams per node
- `ExecuteInternal(FPCGContext* Context)` — Core execution entry point called by the PCG runtime; override this in subclasses to implement node logic; reads input data from `Context`, processes it, and writes results back to `Context`
