<!-- Source: https://dev.epicgames.com/documentation/en-us/unreal-engine/input-overview-in-unreal-engine -->

The **PlayerInput** Object is responsible for converting input from the player into data that Actors (like PlayerControllers or Pawns) can understand and make use of. It is part of an input processing flow that translates hardware input from players into game events and movement with PlayerInput mappings and InputComponents.

For an example of setting up Input, refer to the [Setup Input](https://dev.epicgames.com/documentation/en-us/unreal-engine/setting-up-user-inputs-in-unreal-engine) documentation.

## Hardware Input

Hardware input from a player is straightforward. It commonly includes key presses, mouse clicks or mouse movement, and controller button presses or joystick movement. Specialized input devices that don't conform to standard axis or button indices, or that have unusual input ranges, can be configured manually by using the [Raw Input Plugin](https://dev.epicgames.com/documentation/en-us/unreal-engine/rawinput-plugin-in-unreal-engine).

## PlayerInput

PlayerInput is a UObject within the PlayerController class that manages player input. It is only spawned on the client. Two structs are defined within PlayerInput. The first, **FInputActionKeyMapping**, defines an ActionMapping. The other, **FInputAxisKeyMapping**, defines an AxisMapping. The hardware input definitions used in both ActionMappings and AxisMappings are established in InputCoreTypes.

- **ActionMappings**: Map a discrete button or key press to a "friendly name" that will later be bound to event-driven behavior. The end effect is that pressing (and/or releasing) a key, mouse button, or keypad button directly triggers some game behavior.

- **AxisMappings**: Map keyboard, controller, or mouse inputs to a "friendly name" that will later be bound to continuous game behavior, such as movement. The inputs mapped in AxisMappings are continuously polled, even if they are just reporting that their input value is currently zero. This allows for smooth transitions in movement or other game behavior, rather than the discrete game events triggered by inputs in ActionMappings.

Hardware axes, such as controller joysticks, provide degrees of input, rather than discrete 1 (pressed) or 0 (not pressed) input. That is, they can be moved to a small degree or a large degree, and your character's movement can vary accordingly. While these input methods are ideal for providing scalable amounts of movement input, AxisMappings can also map common movement keys, like WASD or Up, Down, Left, Right, to continuously-polled game behavior.

### Setting Input Mappings

Input mappings are stored in configuration files, and can be edited in the Input section of Project Settings.

1. In the Level Editor, select **Edit > Project Settings**.
2. Click **Input** in the **Project Settings** tab that appears.

In this window, you can:
- Change the properties of (hardware) axis inputs
- Add or edit ActionMappings
- Add or edit AxisMappings

## InputComponent

**InputComponents** are most commonly present in Pawns and Controllers, although they can be set in other Actors and Level Scripts if desired. The InputComponent links the AxisMappings and ActionMappings in your project to game actions, usually functions, set up either in C++ code or Blueprint graphs.

The priority stack for input handling by InputComponents is as follows (highest priority first):

1. Actors with "Accepts input" enabled, from most-recently enabled to least-recently enabled. If you want a particular Actor to always be the first one considered for input handling, you can re-enable its "Accepts input" and it will be moved to the top of the stack.
2. Controllers.
3. Level Script.
4. Pawns.

If one InputComponent takes the input, it is not available further down the stack.

## Input Processing Procedure

### Example - Moving Forward

This example is taken from the First Person template provided with Unreal Engine.

1. **Hardware Input from Player:** The player presses W.
2. **PlayerInput Mapping:** The AxisMapping translates W to "MoveForward" with a scale of 1.
3. **InputComponent Priority Stack:** Proceeding through the InputComponent priority stack, the first binding of the "MoveForward" input is in the AFirstPersonBaseCodeCharacter class. This class is the current player's Pawn, so its InputComponent is checked last.

```cpp
void AFirstPersonBaseCodeCharacter::SetupPlayerInputComponent(class UInputComponent* InputComponent)
{
    // set up gameplay key bindings
    check(InputComponent);
    ...
    InputComponent->BindAxis("MoveForward", this, &AFirstPersonBaseCodeCharacter::MoveForward);
    ...
}
```

This step could also be accomplished in Blueprints by having an InputAxis MoveForward node in the Character's EventGraph. Whatever this node is connected to is what will execute when W is pressed.

4. **Game Logic:** AFirstPersonBaseCodeCharacter's MoveForward function executes.

```cpp
void AFirstPersonBaseCodeCharacter::MoveForward(float Value)
{
    if ( (Controller != NULL) && (Value != 0.0f) )
    {
        // find out which way is forward
        FRotator Rotation = Controller->GetControlRotation();
        // Limit pitch when walking or falling
        if ( CharacterMovement->IsMovingOnGround() || CharacterMovement->IsFalling() )
        {
            Rotation.Pitch = 0.0f;
        }
        // add movement in that direction
        const FVector Direction = FRotationMatrix(Rotation).GetScaledAxis(EAxis::X);
        AddMovementInput(Direction, Value);
    }
}
```

## Touch Interface

By default, games running on touch devices will have two virtual joysticks (like a console controller). You can change this in your **Project Settings**, in the **Input** section, with the **Default Touch Interface** property. This points to a Touch Interface Setup asset. The default one, **DefaultVirtualJoysticks** is located in shared engine content (`/Engine/MobileResources/HUD/DefaultVirtualJoysticks.DefaultVirtualJoysticks`). There is also a Left Stick only version, **LeftVirtualJoystickOnly**, for games that do not need to turn the camera.

If you do not want any virtual joysticks, just clear the Default Touch Interface property. Additionally, you can force the touch interface for your game independent of the platform it is running by checking Always Show Touch Interface (or by running the PC game with -faketouches).

## Enhanced Input Plugin

For projects that require more advanced input features, like complex input handling or runtime control remapping, the [Enhanced Input Plugin](https://dev.epicgames.com/documentation/en-us/unreal-engine/enhanced-input-in-unreal-engine) gives developers an easy upgrade path and backward compatibility with the engine's default input system. This plugin implements features like radial dead zones, chorded actions, contextual input and prioritization, and the ability to extend your own filtering and processing of raw input data in an Asset-based environment.

## Getting Started

To configure your project to use Enhanced Input, enable the Enhanced Input Plugin. You can do this by opening the **Edit** dropdown menu in the editor and selecting **Plugins**. Under the **Input** section of the Plugin List, find and enable the Enhanced Input Plugin, then restart the editor.

Once the editor has restarted, you can set your project to use Enhanced Input Plugin classes instead of the default Unreal Engine input handlers. Go to the **Edit** dropdown menu and choose **Project Settings**. From there, locate the **Input** section (under the **Engine** heading) and find the **Default Classes** settings. To use Enhanced Input, change these settings to EnhancedPlayerInput and EnhancedInputComponent, respectively.

## Core Concepts

The Enhanced Input system has four main concepts:

- **Input Actions** are the communication link between the Enhanced Input system and your project's code. An Input Action can be anything that an interactive character might do, like jumping or opening a door, but could also be used to indicate user input states, like holding a button that changes walking movement to running.

- **Input Mapping Contexts** map user inputs to Actions and can be dynamically added, removed, or prioritized for each user.

- **Modifiers** adjust the value of raw input coming from the user's devices. An Input Mapping Context can have any number of modifiers associated with each raw input for an Input Action. Common Modifiers include dead zones, input smoothing over multiple frames, conversion of input vectors from local to world space, and several others.

- **Triggers** use post-Modifier input values, or the output magnitudes of other Input Actions, to determine whether or not an Input Action should activate.

## Input Actions

Input Actions are the connection between the system and your project's code. You can create an Input Action by right-clicking in the **Content Browser**, expanding the **Input** option, and choosing **Input Action**. To trigger an Input Action, you must include it in an Input Mapping Context, and add that Input Mapping Context to the local player's **Enhanced Input Local Player Subsystem**.

To make your Pawn class respond to a triggered Input Action, you must bind it to the appropriate type of **Trigger Event** in SetupPlayerInputComponent:

```cpp
if (UEnhancedInputComponent* PlayerEnhancedInputComponent = Cast<UEnhancedInputComponent>(PlayerInputComponent))
{
    // This calls the handler function on the tick when MyInputAction starts
    if (MyInputAction)
    {
        PlayerEnhancedInputComponent->BindAction(MyInputAction, ETriggerEvent::Started, this, &AMyPawn::MyInputHandlerFunction);
    }

    // This calls the handler function by name on every tick while the input conditions are met
    if (MyOtherInputAction)
    {
        PlayerEnhancedInputComponent->BindAction(MyOtherInputAction, ETriggerEvent::Triggered, this, TEXT("MyOtherInputHandlerFunction"));
    }
}
```

When binding Input Actions, you can choose between four different handler function signatures:

| Return Type | Parameters | Usage Notes |
| --- | --- | --- |
| void | `()` | Suitable for simple cases where you don't need any extra information from the Enhanced Input system. |
| void | `(const FInputActionValue& ActionValue)` | Provides access to the current value of the Input Action. |
| void | `(const FInputActionInstance& ActionInstance)` | Provides access to the current value of the Input Action, the type of trigger event, and relevant timers. |
| void | `(FInputActionValue ActionValue, float ElapsedTime, float TriggeredTime)` | Signature used when dynamically binding to a UFunction by its name; parameters are optional. |

## Input Mapping Contexts

Input Mapping Contexts describe the rules for triggering one or more Input Actions. Its basic structure is a hierarchy with a list of Input Actions at the top level. Under the Input Action level is a list of user inputs that can trigger each Input Action, such as keys, buttons, and movement axes. The bottom level contains a list of Input Triggers and Input Modifiers for each user input, which you can use to determine how an input's raw value is filtered or processed, and what restrictions it must meet in order to drive the Input Action at the top of its hierarchy. Any input can have multiple Input Modifiers and Input Triggers. These will evaluate in order, using the output of each step as the input for the next.

Once you have populated an Input Mapping Context, you can add it to the Local Player associated with the Pawn's Player Controller:

```cpp
if (APlayerController* PC = Cast<APlayerController>(GetController()))
{
    if (UEnhancedInputLocalPlayerSubsystem* Subsystem = ULocalPlayer::GetSubsystem<UEnhancedInputLocalPlayerSubsystem>(PC->GetLocalPlayer()))
    {
        Subsystem->ClearAllMappings();
        Subsystem->AddMappingContext(MyInputMappingContext, MyInt32Priority);
    }
}
```

## Input Modifiers

Input Modifiers are pre-processors that alter the raw input values that Unreal Engine receives before sending them on to Input Triggers. The Enhanced Input Plugin ships with a variety of Input Modifiers to perform tasks like changing the order of axes, implementing "dead zones", converting axial input to world space, and several others.

## Input Triggers

Input Triggers determine whether or not a user input, after passing through an optional list of Input Modifiers, should activate the corresponding Input Action within its Input Mapping Context. Most Input Triggers analyze the input itself, checking for minimum actuation values and validating patterns like short taps, prolonged holds, or the typical "press" or "release" events. The one exception to this rule is the "Chorded Action" Input Trigger, which requires another Input Action to be triggered. By default, any user activity on an input will trigger on every tick.
