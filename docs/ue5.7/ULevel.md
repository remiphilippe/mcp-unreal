# ULevel

**Parent**: UObject
**Module**: Engine

Represents a single persistent or streaming sublevel within a `UWorld`. A level owns an array of all actors placed or spawned within it, a BSP model, precomputed lighting data, and a reference to its `ALevelScriptActor`. The persistent level is always loaded; streaming levels are managed by `ULevelStreamingVolume` or `ULevelStreaming` and loaded/unloaded at runtime. Levels are stored as `.umap` assets on disk.

## Key Properties

- `Actors` — `TArray<AActor*>` of every actor belonging to this level; includes the `AWorldSettings` actor at index 0
- `OwningWorld` — The `UWorld*` this level is part of; may differ from `GWorld` for streaming levels
- `Model` — The `UModel*` representing BSP brush geometry for this level; legacy geometry from the CSG workflow
- `LevelBuildDataId` — `FGuid` identifying the lighting build data set; changes when the level is re-lit
- `bIsLightingScenario` — When true, this level acts as a lighting scenario and its static lighting is applied to the entire world when streamed in
- `bIsVisible` — Runtime visibility flag; reflects whether the level's actors are currently rendered and ticking
- `LevelScriptActor` — The `ALevelScriptActor*` containing Blueprint logic authored in the Level Blueprint for this level
- `NavListStart` — First navigation-relevant actor in the level's navigation linked list (internal navigation system use)
- `NavListEnd` — Last navigation-relevant actor in the level's navigation linked list
- `WorldSettings` — Cached pointer to the `AWorldSettings` actor for fast access; always the first actor in `Actors`
- `PrecomputedLightVolume` — Volumetric precomputed lighting data (`FPrecomputedLightVolume`) used for dynamic object indirect lighting
- `LevelColor` — `FLinearColor` used to tint this level's actors in the editor viewport for visual identification of streaming levels

## Key Functions

- `GetLevelScriptActor()` — Returns the `ALevelScriptActor*` for this level's Level Blueprint; may be null for levels with no Blueprint logic
- `GetWorldSettings()` — Returns the `AWorldSettings*` actor; this actor stores per-level game mode, gravity, and global settings
- `GetActors()` — Returns a const reference to the `TArray<AActor*>` of all actors in this level
- `GetWorld()` — Returns the `UWorld*` that owns this level; override of `UObject::GetWorld()`
- `SortActorList()` — Reorders the `Actors` array to ensure `AWorldSettings` is first and `APlayerController`/`APawn` actors are in a stable order; called internally after level load
- `IncrementalUpdateComponents(int32 NumComponentsToUpdate, bool bRerunConstructionScripts)` — Registers a batch of actor components with the world in slices to spread the cost over multiple frames during streaming
- `InitializeActors()` — Calls `AActor::PreInitializeComponents` on all uninitialized actors in this level; part of the level load sequence
- `RouteActorInitialize()` — Calls `AActor::PostInitializeComponents` and `AActor::BeginPlay` on actors that have been initialised but not yet begun play
- `HasVisibilityChangeRequestPending()` — Returns `true` if a streaming request to change this level's visibility is in flight
- `SetVisibility(bool bIsVisible)` — Shows or hides all actors in this level and toggles their tick/render state; used by the streaming system
- `IsCurrentLevel()` — Returns `true` if this is the world's current level (the target for `SpawnActor` calls with no explicit level argument)
- `IsPersistentLevel()` — Returns `true` if this is the persistent (always-loaded) level of its owning world
- `GetOutermost()` — Returns the `UPackage*` (`UObject::GetOutermost`) corresponding to the `.umap` package on disk
