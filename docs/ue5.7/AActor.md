# AActor

**Parent**: UObject
**Module**: Engine

AActor is the base class for all actors that can be placed or spawned in a level. Actors support 3D transformations (location, rotation, scale), component attachment hierarchies, replication for networking, and lifecycle events. They are the primary building block for gameplay objects in Unreal Engine.

## Key Properties

- `RootComponent` — The root USceneComponent that defines the transform for this actor
- `bReplicates` — Whether this actor replicates to network clients
- `bCanBeDamaged` — Whether this actor can take damage (ApplyDamage)
- `bHidden` — Whether this actor is hidden in game
- `Tags` — Array of FName tags for gameplay identification
- `Owner` — The AActor that owns this actor (set via SetOwner)
- `Instigator` — The APawn responsible for damage caused by this actor
- `PrimaryActorTick` — Tick function configuration (bCanEverTick, TickInterval)

## Key Functions

- `BeginPlay()` — Called when the game starts or when the actor is spawned. Override for initialization.
- `Tick(float DeltaTime)` — Called every frame. Must enable with PrimaryActorTick.bCanEverTick = true.
- `EndPlay(EEndPlayReason::Type Reason)` — Called when actor is destroyed or level unloaded.
- `GetActorLocation()` — Returns the FVector world position of the root component.
- `SetActorLocation(FVector NewLocation)` — Teleports actor to new location.
- `SetActorRotation(FRotator NewRotation)` — Sets actor rotation.
- `SetActorScale3D(FVector NewScale)` — Sets actor scale.
- `AddActorWorldOffset(FVector DeltaLocation)` — Moves actor by delta with optional sweep.
- `GetComponentByClass(TSubclassOf<UActorComponent> ComponentClass)` — Finds first component of given class.
- `Destroy()` — Marks the actor for destruction at end of frame.
- `SetLifeSpan(float InLifespan)` — Sets timer to auto-destroy after N seconds.
- `AttachToActor(AActor* ParentActor, FAttachmentTransformRules Rules)` — Attaches this actor to another.
- `GetWorld()` — Returns the UWorld this actor belongs to.
