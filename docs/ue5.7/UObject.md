# UObject

**Parent**: None
**Module**: CoreUObject

UObject is the base class for almost all Unreal Engine objects. It provides the foundation for UE's reflection system (UCLASS, UPROPERTY, UFUNCTION macros), garbage collection via reference tracking, serialization for save/load and asset pipelines, the Class Default Object (CDO) pattern for instancing, and stable object naming and path addressing. Every Blueprint class, Actor, Component, and Asset inherits from UObject.

## Key Properties

- `Outer` — The UObject that owns this object in the object hierarchy (e.g., an actor's component has the actor as Outer)
- `Class` — The UClass descriptor for this object's type, used by the reflection system
- `Name` — The FName identifying this object within its Outer's namespace
- `ObjectFlags` — Bitmask of EObjectFlags controlling GC, serialization, and editor behavior (e.g., RF_ClassDefaultObject, RF_NeedLoad, RF_Transient)

## Key Functions

- `GetName()` — Returns the FString name of this object within its Outer.
- `GetPathName()` — Returns the full dot-separated path (e.g., `/Game/Maps/Level.Level:Actor.Component`).
- `GetClass()` — Returns the UClass descriptor for this object's runtime type.
- `GetOuter()` — Returns the UObject that owns this object in the hierarchy.
- `GetWorld()` — Returns the UWorld associated with this object (may return nullptr for non-world objects).
- `IsA(UClass* SomeBase)` — Returns true if this object is an instance of the given class or a subclass.
- `GetDefaultObject()` — Returns the Class Default Object (CDO) for this object's class.
- `StaticClass()` — Static method returning the UClass for this C++ type; used with SpawnActor, IsA, etc.
- `CreateDefaultSubobject<T>(FName SubobjectName)` — Creates a component or subobject during constructor; registers it with the CDO.
- `Rename(const TCHAR* NewName, UObject* NewOuter)` — Renames this object and optionally moves it to a new Outer.
- `MarkPendingKill()` — Marks object for garbage collection; use `IsValid(Obj)` to check validity after.
- `IsValid(const UObject* Obj)` — Free function; safe null and pending-kill check (prefer over `!= nullptr`).
- `PostInitProperties()` — Virtual; called after the constructor and property initialization. Override for post-CDO setup.
- `PostLoad()` — Virtual; called after an object is loaded from disk. Override for fixup logic.
- `Serialize(FArchive& Ar)` — Virtual; implement for custom serialization beyond UPROPERTY auto-serialization.
