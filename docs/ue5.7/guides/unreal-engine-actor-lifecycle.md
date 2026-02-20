<!-- Source: https://dev.epicgames.com/documentation/en-us/unreal-engine/unreal-engine-actor-lifecycle -->

This document is a high-level overview of the lifecycle of an [Actor](https://dev.epicgames.com/documentation/en-us/unreal-engine/actors-in-unreal-engine), which includes:

- How an Actor is instantiated or spawned into the level, including how the Actor is initialized.
- How an Actor is marked PendingKill and then removed or destroyed through Garbage Collection.
- The flow chart below shows the primary paths for how an Actor is instanced. No matter how an Actor is created, they all follow the same path to their destruction.

## Load from Disk

The Load From Disk path occurs for any Actor that is already in the level, like when [UEngine::LoadMap](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/Engine/UEngine/LoadMap?application_version=5.5) occurs, or when [level streaming](https://dev.epicgames.com/documentation/en-us/unreal-engine/level-streaming-in-unreal-engine) calls [UWorld::AddToWorld](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/Engine/UWorld/AddToWorld?application_version=5.5).

1. Actors in a package/level are loaded from disk.
2. The serialized Actor calls [PostLoad](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObject/UObject/PostLoad?application_version=5.5) after it has finished loading from disk. Any custom versioning and fixup behavior should be implemented here. PostLoad is mutually exclusive with [AActor::PostActorCreated](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/GameFramework/AActor/PostActorCreated?application_version=5.5).
3. The world calls [UAISystemBase::InitializeActorsForPlay](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/AI/UAISystemBase/InitializeActorsForPlay?application_version=5.5) to prepare Actors to start gameplay.
4. The level calls [ULevel::RouteActorInitialize](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/Engine/ULevel/RouteActorInitialize?application_version=5.5) for any non-initialized Actors and any seamless travel carry-over.
   1. [AActor::PreInitializeComponents](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/GameFramework/AActor/PreInitializeComponents?application_version=5.5) is called before InitializeComponent is called on the Actor's components.
   2. [UActorComponent::InitializeComponent](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/Components/UActorComponent/InitializeComponent?application_version=5.5) is a helper function for the creation of each component defined on the Actor.
   3. [AActor::PostInitializeComponents](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/GameFramework/AActor/PostInitializeComponents?application_version=5.5) is called after the Actor's components have been initialized.
5. [AActor::BeginPlay](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/GameFramework/AActor/BeginPlay?application_version=5.5) is called when the level is started.

## Play in Editor

In the Play in Editor path, Actors are copied from the Editor instead of being loaded from the disk. Then, the copied Actors initialize similarly to the flow described in the Load From Disk path.

1. Actors in the Editor are duplicated into a new World.
2. [UObject::PostDuplicate](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObject/UObject/PostDuplicate?application_version=5.5) is called.
3. [UAISystemBase::InitializeActorsForPlay](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/AI/UAISystemBase/InitializeActorsForPlay?application_version=5.5)
4. [ULevel::RouteActorInitialize](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/Engine/ULevel/RouteActorInitialize?application_version=5.5) for any non-initialized Actors and covers any seamless travel carry over.
   1. [AActor::PreInitializeComponents](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/GameFramework/AActor/PreInitializeComponents?application_version=5.5) is called before InitializeComponent is called on the Actor's components.
   2. [UActorComponent::InitializeComponent](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/Components/UActorComponent/InitializeComponent?application_version=5.5) is a helper function for the creation of each component defined on the Actor.
   3. [AActor::PostInitializeComponents](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/GameFramework/AActor/PostInitializeComponents?application_version=5.5) is called after the Actor's components have been initialized.
5. [AActor::BeginPlay](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/GameFramework/AActor/BeginPlay?application_version=5.5) is called when the level is started.

## Spawning

When you spawn an instance of an Actor, this is the path that will be followed:

1. [UWorld::SpawnActor](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/Engine/UWorld/SpawnActor?application_version=5.5) is called.
2. [AActor::PostSpawnInitialize](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/GameFramework/AActor/PostSpawnInitialize?application_version=5.5) is called after the Actor is spawned in the world.
3. [AActor::PostActorCreated](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/GameFramework/AActor/PostActorCreated?application_version=5.5) is called for spawned Actors after its creation, any constructor implementation behavior should go here. PostActorCreated is mutually exclusive with PostLoad.
4. [AActor::ExecuteConstruction](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/GameFramework/AActor/ExecuteConstruction?application_version=5.5):
5. [AActor::OnConstruction](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/GameFramework/AActor/OnConstruction?application_version=5.5) - The construction of the Actor, this is where Blueprint Actors have their components created and Blueprint variables are initialized.
6. [AActor::PostActorConstruction](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/GameFramework/AActor/PostActorConstruction?application_version=5.5):
   1. [AActor::PreInitializeComponents](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/GameFramework/AActor/PreInitializeComponents?application_version=5.5) Called before InitializeComponent is called on the Actor's components.
   2. [UActorComponent::InitializeComponent](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/Components/UActorComponent/InitializeComponent?application_version=5.5) is a Helper function for the creation of each component defined on the Actor.
   3. [AActor::PostInitializeComponents](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/GameFramework/AActor/PostInitializeComponents?application_version=5.5) is called after the Actor's components have been initialized.
7. [UWorld::OnActorSpawned](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/Engine/UWorld/AddOnActorSpawnedHandler?application_version=5.5) is broadcast on UWorld.
8. [AActor::BeginPlay](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/GameFramework/AActor/BeginPlay?application_version=5.5) is called.

## Deferred Spawn

An Actor can be Deferred Spawned by having any properties set to "Expose on Spawn."

1. [UWorld::SpawnActorDeferred](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/Engine/UWorld/SpawnActorDeferred?application_version=5.5) is meant to spawn procedural Actors, allows additional setup before Blueprint construction script.
2. Everything in SpawnActor occurs, but after [AActor::PostActorCreated](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/GameFramework/AActor/PostActorCreated?application_version=5.5) the following occurs:
   1. Do setup and call various "initialization functions" with a valid but incomplete Actor instance.
   2. [AActor::FinishSpawning](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/GameFramework/AActor/FinishSpawning?application_version=5.5) is called to Finalize the Actor, picks up at ExecuteConstruction in the Spawn Actor line.

## End of Actor Lifecycle

You can destroy Actors in several ways, but the way they are removed from the world follows the same method:

- [AActor::Destroy](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/GameFramework/AActor/Destroy?application_version=5.5) is called manually by a game any time an Actor is meant to be removed, but gameplay is still occurring. The Actor is marked pending kill and removed from Level's array of Actors.

- [AActor::EndPlay](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/GameFramework/AActor/EndPlay?application_version=5.5) is called in several places to guarantee the life of the Actor is coming to an end. EndPlay is called from:
  - An explicit call to Destroy.
  - When Play in Editor has Ended.
  - Level Transitions (seamless travel or load map).
  - A streaming level containing the Actor is unloaded.
  - The lifetime of the Actor has expired.
  - Application shut down (All Actors are Destroyed).

Regardless of how this happens, the Actor will be marked `RF_PendingKill` so that UE de-allocates it from memory during the next garbage collection cycle. Also, rather than checking for pending kill manually, consider using an `FWeakObjectPtr<AActor>` as it is cleaner.

Actors may not necessarily be destroyed when their EndPlay is called. For example, if `s.ForceGCAfterLevelStreamedOut` is `false` and a sublevel is quickly reloaded then an Actor's EndPlay would be called, but the actor may be "resurrected" and would be the exact same actor that previously existed along with its local variables that were not re-initialized to its defaults.

- [AActor::OnDestroyed](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/GameFramework/AActor?application_version=5.5) - This is a legacy response to Destroy. We recommend you move any logic here to EndPlay as it is called by level transition and other game cleanup functions.

## Garbage Collection

Some time after an object has been marked for destruction, Garbage Collection will remove it from memory, freeing any resources it was using.

The following functions are called on the object during its destruction:

1. [UObject::BeginDestroy](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObject/UObject/BeginDestroy?application_version=5.5) - This is the object's chance to free up memory and handle other multithreaded resources (ie: graphics thread proxy objects). Most gameplay functionality related to being destroyed should have been handled earlier, in EndPlay.

2. [UObject::IsReadyForFinishDestroy](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObject/UObject/IsReadyForFinishDestroy?application_version=5.5) - The garbage collection process will call this function to determine whether or not the object is ready to be deallocated permanently. By returning false, this function can defer actual destruction of the object until the next garbage collection pass.

3. [UObject::FinishDestroy](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObject/UObject/FinishDestroy?application_version=5.5) - Finally, the object is really going to be destroyed, and this is another chance to free up internal data structures. This is the last call before memory is freed.

### Advanced Garbage Collection

The [Garbage Collection](https://dev.epicgames.com/documentation/en-us/unreal-engine/unreal-object-handling-in-unreal-engine) process in Unreal Engine builds clusters of objects that are all destroyed together. **Clustering** reduces the total time and overall memory churn associated with garbage collection compared to deleting objects individually.

As an object loads, it may create subobjects. By combining the object and its subobjects into a single cluster for the garbage collector, the engine can delay freeing the resources used by the cluster until the entire object is ready to be freed, and can then free all of the resources at once.

Garbage collection does not need to be configured or modified at all for most projects, but there are some specific cases in which the "clustering" behavior of the garbage collector can be altered to improve efficiency:

- **Clustering** - Turn clustering off. In **Project Settings**, under the **Garbage Collection** section, the **Create Garbage Collector UObject Clusters** option can be set to false. For most projects, this will result in less efficient garbage collection, so it is recommended in cases where performance testing reveals that it is beneficial.
