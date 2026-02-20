# USkeletalMeshComponent

**Parent**: USkinnedMeshComponent
**Module**: Engine

Renders skeletal meshes and drives their animation. Manages the active `UAnimInstance`, handles physics asset simulation, exposes bone and socket transforms, and supports morph targets. Used on all animated characters, creatures, and mechanical objects that require bone-based deformation.

## Key Properties

- `AnimationMode` — `EAnimationMode` enum controlling the source of animation: `AnimationBlueprint` (driven by an `AnimClass` instance), `AnimationSingleNode` (plays a single `UAnimationAsset`), `AnimationCustomMode` (externally driven)
- `AnimClass` — `TSubclassOf<UAnimInstance>` specifying the Animation Blueprint class instantiated at runtime when `AnimationMode` is `AnimationBlueprint`
- `AnimScriptInstance` — `UAnimInstance*` the currently active animation instance; valid after `InitializeAnimScriptInstance` has been called
- `SkeletalMesh` / `SkeletalMeshAsset` — `USkeletalMesh*` asset providing the geometry, skeleton hierarchy, and default materials; property name varies by engine version (`SkeletalMeshAsset` in UE 5.1+)
- `bPauseAnims` — `bool` freezing animation evaluation each frame while leaving the current pose in place; physics simulation continues
- `GlobalAnimRateScale` — `float` multiplier applied to the animation update rate for all animations on this component; `0.0` effectively pauses animation
- `KinematicBonesUpdateType` — `EKinematicBonesUpdateToPhysics` enum controlling which bones are updated when mixing animation with physics: `SkipAllBones`, `SkipSimulatingBones`, `SkipFixedBones`
- `PhysicsTransformUpdateMode` — `EPhysicsTransformUpdateMode` enum determining how simulated physics bodies write back to bone transforms: `SimulationUpatesComponentTransform`, `ComponentTransformIsKinematic`

## Key Functions

- `SetAnimInstanceClass(TSubclassOf<UAnimInstance> NewClass)` — Replaces the current animation Blueprint class and re-initializes the animation instance; safe to call at runtime
- `GetAnimInstance()` — Returns `UAnimInstance*` for the currently active animation instance, or `nullptr` if none is running
- `PlayAnimation(UAnimationAsset* NewAnimToPlay, bool bLooping)` — Sets `AnimationMode` to `AnimationSingleNode` and begins playing the given asset; replaces any current single-node animation
- `SetAnimation(UAnimationAsset* NewAnimToPlay)` — Assigns an animation asset for single-node playback without starting it; call `Play` separately
- `Play(bool bLooping)` — Starts or resumes single-node animation playback from the current position
- `Stop()` — Stops single-node animation playback and resets to the first frame
- `SetSkeletalMesh(USkeletalMesh* NewMesh, bool bReinitPose)` — Swaps the skeletal mesh asset at runtime; `bReinitPose` resets the current pose to the reference pose when `true`
- `GetBoneTransform(int32 BoneIndex, EBoneSpaces::Type BoneSpace)` — Returns `FTransform` of the bone at `BoneIndex` in the specified space (`BoneSpace`); use `EBoneSpaces::WorldSpace` or `EBoneSpaces::ComponentSpace`
- `GetSocketLocation(FName InSocketName)` — Returns `FVector` world-space position of the named socket or bone
- `GetBoneName(int32 BoneIndex)` — Returns `FName` of the bone at the given skeleton index
- `GetNumBones()` — Returns `int32` total number of bones in the referenced skeleton
- `SetMorphTarget(FName MorphTargetName, float Value, bool bRemoveZeroWeight)` — Sets the blend weight of a named morph target; `Value` is clamped to `[0.0, 1.0]`; `bRemoveZeroWeight` removes the entry from the active list when `Value` reaches zero for efficiency
- `ClearMorphTargets()` — Resets all morph target weights to zero and clears the active morph target list
