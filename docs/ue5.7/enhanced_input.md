# Enhanced Input System

**Module**: EnhancedInput

The Enhanced Input system replaces the legacy input system in UE 5.x. It separates input into three layers: Input Actions (what happened), Input Mapping Contexts (which keys trigger it), and Input Modifiers/Triggers (how to process the raw input). This is the recommended input system for all new UE 5.7 projects.

## Core Classes

- `UInputAction` — Defines a gameplay action (e.g., "Jump", "Move", "Look"). Has a ValueType: bool, float, FVector2D, or FVector3D.
- `UInputMappingContext` — Maps physical keys to Input Actions. Multiple contexts can be active with priorities.
- `UEnhancedInputComponent` — Component for binding Input Actions to functions. Replaces UInputComponent.
- `UEnhancedInputLocalPlayerSubsystem` — Manages active mapping contexts per player.
- `UInputModifier` — Modifies raw input values (deadzone, sensitivity, negate, swizzle).
- `UInputTrigger` — Defines when an action fires (pressed, released, held, tap, chord).

## Setup Pattern

1. Create UInputAction assets for each gameplay action
2. Create UInputMappingContext and map keys to actions
3. In your character/controller BeginPlay, add the mapping context:

```cpp
UEnhancedInputLocalPlayerSubsystem* Subsystem = ULocalPlayer::GetSubsystem<UEnhancedInputLocalPlayerSubsystem>(GetLocalPlayer());
Subsystem->AddMappingContext(DefaultMappingContext, 0);
```

4. Bind actions in SetupPlayerInputComponent:

```cpp
UEnhancedInputComponent* EnhancedInput = Cast<UEnhancedInputComponent>(InputComponent);
EnhancedInput->BindAction(MoveAction, ETriggerEvent::Triggered, this, &AMyCharacter::Move);
EnhancedInput->BindAction(JumpAction, ETriggerEvent::Started, this, &AMyCharacter::Jump);
```

## Key Functions

- `UEnhancedInputLocalPlayerSubsystem::AddMappingContext(UInputMappingContext* Context, int32 Priority)` — Activates a mapping context.
- `UEnhancedInputLocalPlayerSubsystem::RemoveMappingContext(UInputMappingContext* Context)` — Deactivates a mapping context.
- `UEnhancedInputComponent::BindAction(UInputAction* Action, ETriggerEvent Event, UObject* Object, FunctionPtr)` — Binds a function to an input action.
- `UInputAction::GetValueType()` — Returns the value type (bool, Axis1D, Axis2D, Axis3D).
