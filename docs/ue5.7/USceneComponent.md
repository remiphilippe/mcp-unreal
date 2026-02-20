# USceneComponent

**Parent**: UActorComponent
**Module**: Engine

USceneComponent adds a transform (location, rotation, scale) to UActorComponent, making it the base class for all components that have a position in the world. It supports parent-child attachment hierarchies. All visible components (meshes, lights, cameras) inherit from USceneComponent.

## Key Properties

- `RelativeLocation` — Position relative to parent (or world if root)
- `RelativeRotation` — Rotation relative to parent
- `RelativeScale3D` — Scale relative to parent
- `bAbsoluteLocation` — If true, location is in world space regardless of parent
- `bAbsoluteRotation` — If true, rotation is in world space regardless of parent
- `bAbsoluteScale` — If true, scale is in world space regardless of parent
- `bVisible` — Whether this component and children are visible
- `Mobility` — EComponentMobility: Static, Stationary, or Movable

## Key Functions

- `SetRelativeLocation(FVector NewLocation)` — Sets position relative to parent.
- `SetWorldLocation(FVector NewLocation)` — Sets position in world space.
- `SetRelativeRotation(FRotator NewRotation)` — Sets rotation relative to parent.
- `SetWorldRotation(FRotator NewRotation)` — Sets rotation in world space.
- `SetRelativeScale3D(FVector NewScale)` — Sets scale relative to parent.
- `AddLocalOffset(FVector DeltaLocation)` — Moves component in local space.
- `AddWorldOffset(FVector DeltaLocation)` — Moves component in world space.
- `GetComponentLocation()` — Returns world-space location.
- `GetComponentRotation()` — Returns world-space rotation.
- `GetForwardVector()` — Returns the forward direction vector.
- `GetRightVector()` — Returns the right direction vector.
- `GetUpVector()` — Returns the up direction vector.
- `AttachToComponent(USceneComponent* Parent, FAttachmentTransformRules Rules)` — Attaches to another component.
- `DetachFromComponent(FDetachmentTransformRules Rules)` — Detaches from parent.
- `GetAttachParent()` — Returns the parent component.
- `GetChildrenComponents(bool bIncludeAllDescendants, TArray<USceneComponent*>& Children)` — Gets attached child components.
