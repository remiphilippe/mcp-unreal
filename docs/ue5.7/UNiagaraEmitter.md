# UNiagaraEmitter

**Parent**: UObject
**Module**: Niagara

Niagara emitter asset that encapsulates all spawn, update, event, and render logic for one particle group. Emitters are assembled into a `UNiagaraSystem` via `FNiagaraEmitterHandle`. An emitter owns a set of scripts (spawn, update, event handlers) and renderer properties that describe how particles look. Emitters can run on the CPU or GPU and optionally inherit from a parent emitter using Niagara's emitter inheritance system.

## Key Properties

- `SimTarget` — Determines execution backend: `CPUSim` (flexible, Blueprint accessible) or `GPUComputeSim` (high particle count, GPU only)
- `bFixedBounds` — When true, uses the static `FixedBounds` box instead of computing dynamic bounds each frame
- `FixedBounds` — `FBox` defining the static local-space bounds when `bFixedBounds` is enabled
- `AllocationMode` — Controls how particle memory is allocated: automatic or manual (`PreAllocationCount`)
- `PreAllocationCount` — Number of particles to pre-allocate when `AllocationMode` is manual; reduces runtime allocation spikes
- `bDeterministic` — When true, the emitter uses `RandomSeed` for reproducible random sequences; useful for testing and cinematics
- `RandomSeed` — Seed value used for deterministic random number generation when `bDeterministic` is enabled
- `bLocalSpace` — When true, particles are simulated in the emitter's local coordinate space and move with the component transform
- `bRequiresPersistentIDs` — When true, each particle gets a stable integer ID that persists across its lifetime; required for event-driven effects and ribbons
- `SpawnScriptProps` — `FNiagaraEmitterScriptProperties` containing the particle spawn script and its variable bindings
- `UpdateScriptProps` — `FNiagaraEmitterScriptProperties` containing the particle update script run every tick
- `EventHandlerScriptProps` — Array of `FNiagaraEventScriptProperties` for scripts that respond to Niagara events from other emitters
- `RendererProperties` — Array of `UNiagaraRendererProperties*` defining how particles are rendered: `UNiagaraSpriteRendererProperties` (billboards), `UNiagaraMeshRendererProperties` (static meshes), `UNiagaraRibbonRendererProperties` (trails), `UNiagaraLightRendererProperties` (point lights)

## Key Functions

- `GetRenderers()` — Returns a const reference to the array of `UNiagaraRendererProperties*` attached to this emitter
- `GetNumRenderers()` — Returns the renderer count as an `int32`
- `GetRendererAt(int32 RendererIndex)` — Returns the `UNiagaraRendererProperties*` at the given index
- `AddRenderer(UNiagaraRendererProperties* Renderer)` — Appends a renderer to the emitter's renderer list
- `RemoveRenderer(UNiagaraRendererProperties* Renderer)` — Removes and destroys the specified renderer from the list
- `GetEmitterSpawnScriptProperties()` — Returns a reference to the `FNiagaraEmitterScriptProperties` for the emitter-level spawn script
- `GetEmitterUpdateScriptProperties()` — Returns a reference to the `FNiagaraEmitterScriptProperties` for the emitter-level update script
- `GetParticleSpawnScriptProperties()` — Returns a reference to the `FNiagaraEmitterScriptProperties` for the per-particle spawn script
- `GetParticleUpdateScriptProperties()` — Returns a reference to the `FNiagaraEmitterScriptProperties` for the per-particle update script
- `SetSimTarget(ENiagaraSimTarget InSimTarget)` — Changes the simulation target between CPU and GPU; requires recompile
- `GetSimTarget()` — Returns the current `ENiagaraSimTarget` enum value
- `MergeChangesFromParent()` — Re-applies changes from a parent emitter asset through Niagara's emitter inheritance system, merging parent overrides with local modifications
