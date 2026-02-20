# UNiagaraComponent

**Parent**: UFXSystemComponent
**Module**: Niagara

Scene component that spawns and controls a `UNiagaraSystem` particle effect in the world. Attach it to any actor to place an effect at a specific transform, or use the `UNiagaraFunctionLibrary` helpers to spawn standalone instances. Provides a runtime API for activating, deactivating, pausing, and setting user-exposed Niagara variables without reloading the asset.

## Key Properties

- `Asset` — The `UNiagaraSystem` asset this component instantiates and simulates
- `bAutoActivate` — When true, the system begins simulating immediately when the component is registered with the world
- `bAutoDestroy` — When true, the component destroys itself after the system completes all emitter loops
- `SeekDelta` — Time step used when seeking the simulation forward during warmup or timeline scrubbing
- `MaxTimeBeforeForceUpdateTransform` — Maximum seconds the component can go without updating its world transform before a forced update is triggered
- `bForceSolo` — Forces this component to run in its own isolated execution group instead of being batched with other instances; useful for debugging but costly at scale
- `bAutoManageAttachment` — When true, the component automatically attaches/detaches from a parent component based on activation state
- `OverrideParameters` — `FNiagaraUserRedirectionParameterStore` holding per-instance parameter overrides that shadow the asset's default exposed parameters

## Key Functions

- `SetAsset(UNiagaraSystem* InAsset)` — Swaps the system asset at runtime and reinitialises the component; triggers a full reset
- `GetAsset()` — Returns the currently assigned `UNiagaraSystem*`
- `Activate(bool bReset)` — Starts or restarts simulation; pass `true` to reset all emitters to their initial state before activating
- `Deactivate()` — Signals all emitters to complete their current loops then stop; particles in flight finish naturally
- `DeactivateImmediate()` — Stops simulation immediately and clears all live particles without waiting for emitters to complete
- `ResetSystem()` — Resets the simulation to t=0 and re-activates; equivalent to `Activate(true)`
- `ReinitializeSystem()` — Fully reinitialises the system from the asset, recompiling parameter bindings; use after changing `Asset` at runtime
- `SetVariableFloat(FName VarName, float Value)` — Sets a user-exposed float parameter by name on this component instance
- `SetVariableVec3(FName VarName, FVector Value)` — Sets a user-exposed `FVector` (Vector3) parameter by name
- `SetVariableLinearColor(FName VarName, FLinearColor Value)` — Sets a user-exposed `FLinearColor` parameter by name
- `SetVariableBool(FName VarName, bool Value)` — Sets a user-exposed boolean parameter by name
- `SetVariableObject(FName VarName, UObject* Value)` — Sets a user-exposed object reference parameter by name
- `SetNiagaraVariableFloat(const FString& InVariableName, float InValue)` — Blueprint-callable variant of `SetVariableFloat` using `FString` name
- `GetNiagaraParticleCount()` — Returns the current total live particle count across all emitters as an `int32`
- `IsComplete()` — Returns `true` if all emitters have finished and the system is no longer simulating
- `IsPaused()` — Returns `true` if the system simulation is currently paused
- `SetPaused(bool bPaused)` — Pauses or unpauses the simulation without destroying live particles
