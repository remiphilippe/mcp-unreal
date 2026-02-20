# UAnimInstance

**Parent**: UObject
**Module**: Engine

Animation Blueprint runtime instance that drives skeletal mesh animation. Each `USkeletalMeshComponent` that uses Animation Blueprint mode owns one `UAnimInstance`. Contains state machines, blend spaces, and montage playback logic. Override `NativeInitializeAnimation` and `NativeUpdateAnimation` in C++ subclasses to implement code-driven animation logic alongside Blueprint-driven graphs.

## Key Properties

- `CurrentSkeleton` — `USkeleton*` asset this instance is bound to; must match the skeleton of the owning `USkeletalMeshComponent`
- `RootMotionMode` — `ERootMotionMode` controlling how root motion extracted from animations is applied: `NoRootMotionExtraction`, `IgnoreRootMotion`, `RootMotionFromMontagesOnly`, `RootMotionFromEverything`
- `bUsingCopyPoseFromMesh` — `bool` enabling this instance to copy its pose from a linked mesh component rather than evaluating its own graph
- `bReceiveNotifiesFromLinkedInstances` — `bool` allowing animation notifies fired by linked animation instances to propagate to this instance

## Key Functions

- `NativeInitializeAnimation()` — Called once when the instance is created and bound to a skeletal mesh; override in C++ to cache component and owner references
- `NativeUpdateAnimation(float DeltaSeconds)` — Called every frame before the animation graph is evaluated; override to update variables that drive state machines and blend spaces
- `Montage_Play(UAnimMontage* MontageToPlay, float InPlayRate, EMontagePlayReturnType ReturnValueType, float InTimeToStartMontageAt, bool bStopAllMontages)` — Starts playback of a montage; returns play duration or position depending on `ReturnValueType`
- `Montage_Stop(float InBlendOutTime, const UAnimMontage* Montage)` — Stops the specified montage with a blend-out over `InBlendOutTime` seconds; pass `nullptr` to stop all active montages
- `Montage_Pause(const UAnimMontage* Montage)` — Pauses playback of the specified montage at its current position
- `Montage_Resume(const UAnimMontage* Montage)` — Resumes a previously paused montage
- `Montage_IsPlaying(const UAnimMontage* Montage)` — Returns `bool` indicating whether the specified montage is currently active and not paused
- `Montage_GetCurrentSection(const UAnimMontage* Montage)` — Returns `FName` of the montage section currently playing
- `GetCurrentActiveMontage()` — Returns `UAnimMontage*` for the highest-weighted currently active montage, or `nullptr` if none
- `GetCurrentStateName(int32 MachineIndex)` — Returns `FName` of the active state in the state machine at `MachineIndex`
- `GetInstanceTransitionTimeElapsed(int32 MachineIndex, int32 TransitionIndex)` — Returns `float` seconds elapsed since the specified transition began in the given state machine
- `GetSkelMeshComponent()` — Returns `USkeletalMeshComponent*` that owns this animation instance
- `TryGetPawnOwner()` — Returns `APawn*` if the owning actor is a pawn, `nullptr` otherwise; convenience accessor for common gameplay use
- `GetOwningActor()` — Returns `AActor*` that owns the skeletal mesh component driving this instance
- `SavePoseSnapshot(FName SnapshotName)` — Captures the current evaluated pose and stores it under `SnapshotName` for later retrieval by a `PoseSnapshot` node in the animation graph
