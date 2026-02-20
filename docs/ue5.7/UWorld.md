# UWorld

**Parent**: UObject
**Module**: Engine

UWorld represents a game world containing actors, components, and level geometry. It manages the simulation state including physics, rendering, and actor lifecycle. There is typically one persistent world with optional streaming sub-levels loaded at runtime.

## Key Properties

- `PersistentLevel` — The main ULevel that is always loaded
- `bIsWorldInitialized` — Whether the world has completed initialization

## Key Functions

- `SpawnActor(UClass* Class, FTransform const* Transform, FActorSpawnParameters Params)` — Spawns a new actor into the world. Primary method for runtime actor creation.
- `SpawnActorDeferred<T>(UClass* Class, FTransform Transform)` — Spawns actor but delays BeginPlay until FinishSpawning is called. Use when you need to set properties before initialization.
- `DestroyActor(AActor* Actor)` — Removes an actor from the world.
- `GetTimerManager()` — Returns FTimerManager for setting timers.
- `LineTraceSingleByChannel(FHitResult& OutHit, FVector Start, FVector End, ECollisionChannel Channel)` — Performs a line trace (raycast) for collision detection.
- `SweepSingleByChannel(FHitResult& OutHit, FVector Start, FVector End, FQuat Rot, ECollisionChannel Channel, FCollisionShape Shape)` — Performs a shape sweep for collision.
- `OverlapMultiByChannel(TArray<FOverlapResult>& OutOverlaps, FVector Pos, FQuat Rot, ECollisionChannel Channel, FCollisionShape Shape)` — Tests for overlapping geometry.
- `GetFirstPlayerController()` — Returns the first local player controller.
- `GetGameInstance()` — Returns the UGameInstance for persistent state across levels.
- `GetAuthGameMode()` — Returns the AGameModeBase for server-side game rules.
- `ServerTravel(const FString& URL)` — Triggers a server-initiated level change.
