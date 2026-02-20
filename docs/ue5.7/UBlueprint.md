# UBlueprint

**Parent**: UBlueprintCore
**Module**: Engine

Blueprint asset containing visual scripting graphs. Represents a Blueprint class in the content browser, storing the graph data, generated class reference, and metadata needed to compile and use the Blueprint at runtime.

## Key Properties

- `ParentClass` — The parent class this Blueprint extends, set at creation and driving which UClass is generated
- `GeneratedClass` — The UClass compiled from this Blueprint's graphs; used at runtime to instantiate Blueprint actors and objects
- `BlueprintType` — Enum controlling how the Blueprint is used: `Normal` (standard actor/object Blueprint), `Const` (all functions are const), `MacroLibrary` (reusable macro collection), `Interface` (defines a contract), `FunctionLibrary` (static utility functions)
- `bRecompileOnLoad` — When true the Blueprint is recompiled during editor load, used to recover from stale bytecode
- `Status` — Compilation state: `BS_Unknown`, `BS_Dirty` (needs recompile), `BS_Error` (last compile failed), `BS_UpToDate` (clean and valid)
- `BlueprintSystemVersion` — Integer version stamp incremented on structural changes; used to detect stale derived data
- `NewVariables` — Array of `FBPVariableDescription` entries declaring variables visible in the My Blueprint panel
- `ComponentTemplates` — Array of `UActorComponent` archetypes used as default subobject templates for actor Blueprints

## Key Functions

- `GetBlueprintClass()` — Returns the generated `UClass` for this Blueprint, equivalent to reading `GeneratedClass`
- `GetAllGraphs(TArray<UEdGraph*>& Graphs)` — Populates an array with every graph in the Blueprint including event graphs, function graphs, and macro graphs
- `GetFunctionGraphs(TArray<UEdGraph*>& Graphs)` — Returns only the function graphs (non-event, non-macro user-defined functions)
- `GetMacroGraphs(TArray<UEdGraph*>& Graphs)` — Returns only the macro library graphs defined in this Blueprint
- `GetEventGraphs(TArray<UEdGraph*>& Graphs)` — Returns only the event graphs (typically one named `EventGraph`)
- `MarkBlueprintAsStructurallyModified()` — Marks the Blueprint dirty and signals that variable layout or class hierarchy has changed, triggering a full recompile
- `MarkBlueprintAsModified()` — Marks the Blueprint dirty for a lighter recompile pass (logic changes without structural impact)
- `GeneratedClassUpdate()` — Propagates changes from the Blueprint data to the generated class without a full compile cycle
- `Rename(const TCHAR* NewName, UObject* NewOuter, ERenameFlags Flags)` — Renames the Blueprint asset and updates the generated class name to match
- `IsNormalBlueprintType()` — Returns true when `BlueprintType == BPTYPE_Normal`, used to guard operations that only apply to standard Blueprints
- `SupportsInputEvents()` — Returns true if this Blueprint type can contain input-related event nodes such as `InputAction` or `InputAxis`
- `SupportsEventGraphs()` — Returns true if the Blueprint type allows event graphs; false for function libraries and interfaces
