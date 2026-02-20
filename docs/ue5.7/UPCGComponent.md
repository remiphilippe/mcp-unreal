# UPCGComponent

**Parent**: UActorComponent
**Module**: PCG

Procedural Content Generation component that runs PCG graphs on actors. Attach this component to any actor to execute a `UPCGGraph` in the context of that actor's transform and bounds. The component manages the full generation lifecycle: scheduling, execution, result tracking, and cleanup.

## Key Properties

- `Graph` — Reference to the `UPCGGraph` asset to execute; swapping this triggers a regeneration if the component is active
- `GenerationTrigger` — Controls when generation runs: `GenerateOnLoad` (auto on level load), `GenerateOnDemand` (only via explicit `Generate()` call), `GenerateAtRuntime` (ticks at runtime for dynamic content)
- `bGenerated` — True after a successful generation pass; false after `Cleanup()` or on a dirty graph
- `bRegenerateInEditor` — When true the component regenerates in the editor viewport when properties change, enabling live preview
- `Seed` — Integer seed passed to the graph for deterministic randomness; changing it produces a different but reproducible result
- `InputType` — Specifies what spatial data the component provides to the graph as its primary input (actor bounds, landscape, primitives, etc.)
- `bParseActorComponents` — When true the component introspects sibling components on the owner actor and exposes them as additional PCG inputs
- `bActivated` — Master on/off switch; setting to false suppresses generation without destroying existing results

## Key Functions

- `Generate()` — Schedules and executes the assigned `PCGGraph` using the current seed and input data, replacing any existing generated output
- `CleanupLocal(bool bRemoveComponents, bool bSave)` — Removes generated actors and components that are local to this component without affecting partitioned or remote results
- `Cleanup()` — Full teardown of all generated output including partitioned grid cells; restores the actor to its pre-generation state
- `NotifyPropertiesChangedFromBlueprint()` — Called from Blueprint when a property that affects generation has changed; triggers a dirty + regenerate cycle
- `SetGraph(UPCGGraph* InGraph)` — Replaces the current graph reference and re-runs generation if the component is active and auto-generation is configured
- `GetGeneratedActors()` — Returns the array of `AActor*` instances spawned by the last generation pass
- `GetGeneratedComponents()` — Returns the array of `UActorComponent*` instances added to the owner or child actors during the last generation pass
- `ForceNotificationOfGeneration()` — Fires the post-generation delegate without re-running the graph, used to synchronize listeners after a manual data update
- `ClearPCGLink()` — Severs the link between this component and a PCG partition actor, used during level unload or component destruction
- `GetSeed()` — Returns the effective seed value, combining the component's `Seed` property with any global seed offset configured in project settings
