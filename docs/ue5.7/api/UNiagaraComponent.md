<!-- Source: https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraComponent -->

[UNiagaraComponent](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraComponent) is the primitive component for a Niagara System.

|  |  |
| --- | --- |
| _Name_ | UNiagaraComponent |
| _Type_ | class |
| _Header File_ | /Engine/Plugins/FX/Niagara/Source/Niagara/Public/NiagaraComponent.h |
| _Include Path_ | #include "NiagaraComponent.h" |

## Syntax

```cpp

UCLASS (ClassGroup=(Rendering, Common), Blueprintable, HideCategories=Object,

       HideCategories=Physics, HideCategories=Collision, ShowCategories=Trigger, EditInlineNew,

       Meta=(BlueprintSpawnableComponent, DisplayName="Niagara Particle System Component"),

       MinimalAPI)

class UNiagaraComponent : public UFXSystemComponent
```

## Inheritance Hierarchy

- [UObjectBase](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObjectBase) → [UObjectBaseUtility](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObjectBaseUtility) → [UObject](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObject) → [UActorComponent](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/UActorComponent) → [USceneComponent](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/USceneComponent) → [UPrimitiveComponent](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/UPrimitiveComponent) → [UFXSystemComponent](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/UFXSystemComponent) → **UNiagaraComponent**

## Implements Interfaces

- [IAsyncPhysicsStateProcessor](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/IAsyncPhysicsStateProcessor)
- [IInterface\_AssetUserData](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/IInterface_AssetUserData)
- [IInterface\_AsyncCompilation](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/IInterface_AsyncCompilation)
- [INavRelevantInterface](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/INavRelevantInterface)
- [IPhysicsComponent](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/IPhysicsComponent)

## Derived Classes

- [UNiagaraCullProxyComponent](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraCullProxyComponent)

## Constructors

| Name | Remarks | Include Path | [Unreal Specifiers](https://dev.epicgames.com/documentation/unreal-engine/ufunctions-in-unreal-engine#functionspecifiers) |
| --- | --- | --- | --- |
| UNiagaraComponent<br>(<br>const [FObjectInitializer](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/FObjectInitializer)& ObjectInitializer<br>) |  | NiagaraComponent.h |  |

## Structs

| Name | Remarks |
| --- | --- |
| [FEmitterOverrideInfo](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraComponent/FEmitterOverrideInfo) |  |

## Typedefs

| Name | Type | Remarks | Include Path |
| --- | --- | --- | --- |
| FOnSynchronizedWithAssetParameters | TMulticastDelegate\_NoParams< void > |  | NiagaraComponent.h |
| FOnSystemInstanceChanged | TMulticastDelegate\_NoParams< void > |  | NiagaraComponent.h |

## Variables

### Public

| Name | Type | Remarks | Include Path | [Unreal Specifiers](https://dev.epicgames.com/documentation/unreal-engine/unreal-engine-uproperties#propertyspecifiers) |
| --- | --- | --- | --- | --- |
| AutoAttachLocationRule | [EAttachmentRule](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/EAttachmentRule) | Options for how we handle our location when we attach to the AutoAttachParent, if bAutoManageAttachment is true. | NiagaraComponent.h | - EditAnywhere<br>- BlueprintReadWrite<br>- Category=Attachment<br>- Meta=(EditCondition="bAutoManageAttachment") |
| AutoAttachParent | [TWeakObjectPtr](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TWeakObjectPtr) < [USceneComponent](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/USceneComponent) > | Component we automatically attach to when activated, if bAutoManageAttachment is true. | NiagaraComponent.h | - VisibleInstanceOnly<br>- BlueprintReadWrite<br>- Category=Attachment<br>- Meta=(EditCondition="bAutoManageAttachment") |
| AutoAttachRotationRule | [EAttachmentRule](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/EAttachmentRule) | Options for how we handle our rotation when we attach to the AutoAttachParent, if bAutoManageAttachment is true. | NiagaraComponent.h | - EditAnywhere<br>- BlueprintReadWrite<br>- Category=Attachment<br>- Meta=(EditCondition="bAutoManageAttachment") |
| AutoAttachScaleRule | [EAttachmentRule](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/EAttachmentRule) | Options for how we handle our scale when we attach to the AutoAttachParent, if bAutoManageAttachment is true. | NiagaraComponent.h | - EditAnywhere<br>- BlueprintReadWrite<br>- Category=Attachment<br>- Meta=(EditCondition="bAutoManageAttachment") |
| AutoAttachSocketName | [FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName) | Socket we automatically attach to on the AutoAttachParent, if bAutoManageAttachment is true. | NiagaraComponent.h | - EditAnywhere<br>- BlueprintReadWrite<br>- Category=Attachment<br>- Meta=(EditCondition="bAutoManageAttachment") |
| bAutoAttachWeldSimulatedBodies | uint32 | Option for how we handle bWeldSimulatedBodies when we attach to the AutoAttachParent, if bAutoManageAttachment is true. | NiagaraComponent.h | - EditAnywhere<br>- BlueprintReadWrite<br>- Category=Attachment<br>- Meta=(EditCondition="bAutoManageAttachment") |
| bAutoManageAttachment | uint32 | True if we should automatically attach to AutoAttachParent when activated, and detach from our parent when completed. | NiagaraComponent.h | - EditAnywhere<br>- BlueprintReadOnly<br>- Category=Attachment |
| bEnablePreviewLODDistance | uint32 |  | NiagaraComponent.h |  |
| bWaitForCompilationOnActivate | uint32 |  | NiagaraComponent.h | - EditAnywhere<br>- Category=Compilation |
| MaxTimeBeforeForceUpdateTransform | float | Time between forced UpdateTransforms for systems that use dynamically calculated bounds, Which is effectively how often the bounds are shrunk. | NiagaraComponent.h |  |
| OnSystemFinished | [FOnNiagaraSystemFinished](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FOnNiagaraSystemFinished) | Called when the particle system is done. | NiagaraComponent.h | - BlueprintAssignable<br>- DuplicateTransient |
| PoolingMethod | [ENCPoolMethod](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/ENCPoolMethod) | How to handle pooling for this component instance. | NiagaraComponent.h |  |
| PreviewLODDistance | float |  | NiagaraComponent.h |  |
| PreviewMaxDistance | float |  | NiagaraComponent.h |  |
| SimCacheDebugNumFramesToCapture | [TOptional](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TOptional) < int32 > |  | NiagaraComponent.h | - EditAnywhere<br>- Category="Niagara Utilities"<br>- Meta=(UIMin="1", ClampMin="1", DisplayName="NumFramesToCapture") |

### Protected

| Name | Type | Remarks | Include Path | [Unreal Specifiers](https://dev.epicgames.com/documentation/unreal-engine/unreal-engine-uproperties#propertyspecifiers) |
| --- | --- | --- | --- | --- |
| AgeUpdateMode | [ENiagaraAgeUpdateMode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/ENiagaraAgeUpdateMode) | Defines the mode use when updating the System age. | NiagaraComponent.h |  |
| Asset | [TObjectPtr](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TObjectPtr) < [UNiagaraSystem](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraSystem) > |  | NiagaraComponent.h | - EditAnywhere<br>- Category="Niagara"<br>- Meta=(DisplayName="Niagara System Asset") |
| AssetExposedParametersChangedHandle | [FDelegateHandle](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FDelegateHandle) |  | NiagaraComponent.h |  |
| bActivateShouldResetWhenReady | uint32 | Should we try and reset when ready? | NiagaraComponent.h |  |
| bAllowScalability | uint32 | Controls whether we allow scalability culling for this component. | NiagaraComponent.h | - EditAnywhere<br>- Category=Niagara<br>- BlueprintGetter=GetAllowScalability<br>- BlueprintSetter=SetAllowScalability |
| bAutoDestroy | uint32 |  | NiagaraComponent.h |  |
| bAwaitingActivationDueToNotReady | uint32 | Did we try and activate but fail due to the asset being not yet ready. Keep looping. | NiagaraComponent.h |  |
| bCanRenderWhileSeeking | bool | Whether or not the component can render while seeking to the desired age. | NiagaraComponent.h |  |
| bDesiredPauseState | uint32 | Stores the current state for pause/unpause desired by the use. | NiagaraComponent.h |  |
| bDidAutoAttach | uint32 | Did we auto attach during activation? Used to determine if we should restore the relative transform during detachment. | NiagaraComponent.h |  |
| bDuringUpdateContextReset | uint32 | Flag to mark us as currently changing auto attachment as part of Activate/Deactivate so we don't reset in the OnAttachmentChanged() callback. | NiagaraComponent.h |  |
| bEnableGpuComputeDebug | uint32 | When true the GPU simulation debug display will enabled, allowing information used during simulation to be visualized. | NiagaraComponent.h | - EditAnywhere<br>- Category=Parameters |
| bForceLocalPlayerEffect | uint32 | Flag allowing us to force this Effect to be considered a LocalPlayer Effect. | NiagaraComponent.h |  |
| bForceSolo | uint32 | When true, this component's system will be force to update via a slower "solo" path rather than the more optimal batched path with other instances of the same system. | NiagaraComponent.h |  |
| bIsCulledByScalability | uint32 | True if this component has been culled by the scalability manager. | NiagaraComponent.h |  |
| bIsFullyComplete | uint32 | True if this component has been fully completed. | NiagaraComponent.h |  |
| bIsSeeking | bool | Whether or not the component is currently seeking to the desired time. | NiagaraComponent.h |  |
| bLockDesiredAgeDeltaTimeToSeekDelta | bool |  | NiagaraComponent.h |  |
| bOverrideWarmupSettings | uint32 | When true then this instance will override the system's warmup settings. | NiagaraComponent.h | - EditAnywhere<br>- Category=Warmup |
| bOwnerAllowsScalabiltiy | uint32 | Whether the owner of this component allows it to be scalability culled. | NiagaraComponent.h |  |
| bRecachePSOs | uint32 | Request recache the PSOs | NiagaraComponent.h |  |
| bRenderingEnabled | uint32 |  | NiagaraComponent.h |  |
| CreateSceneComponentUtilsFunction | [TFunction](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TFunction) < [INiagaraSceneComponentUtils](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/INiagaraSceneComponentUtils) \*()> |  | NiagaraComponent.h |  |
| CullProxy | [TObjectPtr](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TObjectPtr) < class [UNiagaraCullProxyComponent](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraCullProxyComponent) > |  | NiagaraComponent.h | - Transient |
| CurrLocalBounds | FBox |  | NiagaraComponent.h |  |
| CustomTimeDilation | float |  | NiagaraComponent.h |  |
| DesiredAge | float | The desired age of the System instance. | NiagaraComponent.h |  |
| EmitterOverrideInfos | [TArray](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArray) < FEmitterOverrideInfo > |  | NiagaraComponent.h |  |
| ForceUpdateTransformTime | float |  | NiagaraComponent.h |  |
| InstanceParameterOverrides | [TMap](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TMap) < [FNiagaraVariableBase](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraVariableBase), [FNiagaraVariant](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraVariant) > |  | NiagaraComponent.h | - EditAnywhere<br>- Category="Niagara" |
| InstanceParameterOverridesCache | [TMap](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TMap) < [FNiagaraVariableBase](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraVariableBase), [FNiagaraVariant](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraVariant) > |  | NiagaraComponent.h | - Transient |
| LastHandledDesiredAge | float | The last desired age value that was handled by the tick function. | NiagaraComponent.h |  |
| MaxSimTime | float | The maximum amount of time in seconds to spend seeking to the desired age in a single frame. | NiagaraComponent.h |  |
| OnSynchronizedWithAssetParametersDelegate | FOnSynchronizedWithAssetParameters |  | NiagaraComponent.h |  |
| OnSystemInstanceChangedDelegate | FOnSystemInstanceChanged |  | NiagaraComponent.h |  |
| OverrideParameters | [FNiagaraUserRedirectionParameterStore](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraUserRedirectionParameter-) |  | NiagaraComponent.h |  |
| RandomSeedOffset | int32 | Offsets the deterministic random seed of all emitters. | NiagaraComponent.h | - EditAnywhere<br>- Category="Randomness" |
| SavedAutoAttachRelativeLocation | FVector | Saved relative transform before auto attachment. | NiagaraComponent.h |  |
| SavedAutoAttachRelativeRotation | FRotator |  | NiagaraComponent.h |  |
| SavedAutoAttachRelativeScale3D | FVector |  | NiagaraComponent.h |  |
| ScalabilityEffectType | [UNiagaraEffectType](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraEffectType) \* |  | NiagaraComponent.h |  |
| ScalabilityManagerHandle | int32 |  | NiagaraComponent.h |  |
| SeekDelta | float | The delta time used when seeking to the desired age. | NiagaraComponent.h |  |
| SimCache | [TObjectPtr](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TObjectPtr) < class [UNiagaraSimCache](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraSimCache) > |  | NiagaraComponent.h | - Transient |
| SystemFixedBounds | FBox |  | NiagaraComponent.h |  |
| SystemInstanceController | FNiagaraSystemInstanceControllerPtr |  | NiagaraComponent.h |  |
| TemplateParameterOverrides | [TMap](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TMap) < [FNiagaraVariableBase](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraVariableBase), [FNiagaraVariant](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraVariant) > |  | NiagaraComponent.h | - EditAnywhere<br>- Category="Niagara" |
| TemplateParameterOverridesCache | [TMap](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TMap) < [FNiagaraVariableBase](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraVariableBase), [FNiagaraVariant](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraVariant) > |  | NiagaraComponent.h | - Transient |
| TickBehavior | [ENiagaraTickBehavior](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/ENiagaraTickBehavior) | Allows you to control how Niagara selects the tick group, changing this while an instance is active will result in not change as it is cached. | NiagaraComponent.h | - EditAnywhere<br>- Category="Niagara"<br>- Meta=(DisplayName="Niagara Tick Behavior") |
| WarmupTickCount | int32 | Number of ticks to process for warmup of the system. | NiagaraComponent.h | - EditAnywhere<br>- Category=Warmup<br>- Meta=(EditCondition="bOverrideWarmupSettings", ClampMin="0") |
| WarmupTickDelta | float | Delta time used when ticking the system in warmup mode. | NiagaraComponent.h | - EditAnywhere<br>- Category=Warmup<br>- Meta=(EditCondition="bOverrideWarmupSettings", ForceUnits=s, UIMin="0.01", UIMax="1") |

## Functions

### Public

| Name | Remarks | Include Path | [Unreal Specifiers](https://dev.epicgames.com/documentation/unreal-engine/ufunctions-in-unreal-engine#functionspecifiers) |
| --- | --- | --- | --- |
| void AdvanceSimulation<br>(<br>int32 TickCount,<br>float TickDeltaSeconds<br>) | Advances this system's simulation by the specified number of ticks and delta time. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara |
| void [AdvanceSimulationByTime](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraComponent/AdvanceSimulationByTime)<br>(<br>float SimulateTime,<br>float TickDeltaSeconds<br>) | Advances this system's simulation by the specified time in seconds and delta time. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara |
| void BeginUpdateContextReset() |  | NiagaraComponent.h |  |
| void ClearEmitterFixedBounds<br>(<br>[FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName) EmitterName<br>) | Clear any previously set fixed bounds for the emitter instance. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara |
| void [ClearSimCache](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraComponent/ClearSimCache)<br>(<br>bool bResetSystem<br>) | Clear any active simulation cache. | NiagaraComponent.h | - BlueprintCallable<br>- Category=SimCache |
| void ClearSystemFixedBounds() | Clear any previously set fixed bounds for the system instance. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara |
| [INiagaraSceneComponentUtils](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/INiagaraSceneComponentUtils) \\* CreateSceneComponentUtils() | DO NOT USE THIS IN EXTERNAL CODE as it is subject to change Methods to handle interop between Actor/Entity/Desc | NiagaraComponent.h |  |
| void DestroyInstance() |  | NiagaraComponent.h |  |
| void DestroyInstanceNotComponent() |  | NiagaraComponent.h |  |
| void EndUpdateContextReset() |  | NiagaraComponent.h |  |
| void EnsureOverrideParametersConsistent() |  | NiagaraComponent.h |  |
| [FNiagaraVariant](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraVariant) [FindParameterOverride](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraComponent/FindParameterOverride)<br>(<br>const [FNiagaraVariableBase](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraVariableBase)& InKey<br>) const | Find the value of an overridden parameter. | NiagaraComponent.h |  |
| [ENiagaraAgeUpdateMode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/ENiagaraAgeUpdateMode) GetAgeUpdateMode() |  | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Get Age Update Mode") |
| bool GetAllowScalability() |  | NiagaraComponent.h | - BlueprintGetter |
| [UNiagaraSystem](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraSystem) \\* GetAsset() |  | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Get Niagara System Asset") |
| [FNiagaraVariant](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraVariant) [GetCurrentParameterValue](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraComponent/GetCurrentParameterValue)<br>(<br>const [FNiagaraVariableBase](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraVariableBase)& InKey<br>) const | Gets the current value of a parameter which is being used by the simulation. | NiagaraComponent.h |  |
| float GetCustomTimeDilation() |  | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara |
| [UNiagaraDataInterface](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraDataInterface) \\* GetDataInterface<br>(<br>const [FString](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FString)& Name<br>) |  | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara |
| float [GetDesiredAge](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraComponent/GetDesiredAge) () | Gets the desired age of the System instance. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Get Desired Age") |
| FBox [GetEmitterFixedBounds](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraComponent/GetEmitterFixedBounds)<br>(<br>[FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName) EmitterName<br>) const | Gets the fixed bounds for an emitter instance. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara |
| [ENiagaraExecutionState](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/ENiagaraExecutionState) GetExecutionState() |  | NiagaraComponent.h |  |
| bool GetForceLocalPlayerEffect() |  | NiagaraComponent.h | - BlueprintGetter |
| bool GetForceSolo() |  | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Is In Forced Solo Mode") |
| bool [GetLockDesiredAgeDeltaTimeToSeekDelta](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraComponent/GetLockDesiredAgeDeltaTimeToSeek-) () | Gets whether or not the delta time used to tick the system instance when using desired age is locked to the seek delta. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara |
| float [GetMaxSimTime](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraComponent/GetMaxSimTime) () | Get the maximum CPU time in seconds we will simulate to the desired age, when we go beyond this limit ticks will be processed in the next frame. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Get Max Desired Age Tick Delta") |
| [ENiagaraOcclusionQueryMode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/ENiagaraOcclusionQueryMode) GetOcclusionQueryMode() |  | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara |
| const [FNiagaraParameterStore](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraParameterStore) & [GetOverrideParameters](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraComponent/GetOverrideParameters) () |  | NiagaraComponent.h |  |
| [FNiagaraUserRedirectionParameterStore](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraUserRedirectionParameter-) & [GetOverrideParameters](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraComponent/GetOverrideParameters) () |  | NiagaraComponent.h |  |
| [FParticlePerfStatsContext](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/FParticlePerfStatsContext) GetPerfStatsContext() |  | NiagaraComponent.h |  |
| float GetPreviewLODDistance() |  | NiagaraComponent.h | - BlueprintCallable<br>- Category=Preview<br>- Meta=(Keywords="preview LOD Distance scalability") |
| bool GetPreviewLODDistanceEnabled() |  | NiagaraComponent.h | - BlueprintCallable<br>- Category=Preview<br>- Meta=(Keywords="preview LOD Distance scalability") |
| int32 GetRandomSeedOffset() |  | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Get Random Seed Offset") |
| bool GetRenderingEnabled() | Gets whether or not rendering is enabled for this component. | NiagaraComponent.h |  |
| [ENiagaraExecutionState](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/ENiagaraExecutionState) GetRequestedExecutionState() |  | NiagaraComponent.h |  |
| int32 GetScalabilityManagerHandle() |  | NiagaraComponent.h |  |
| float [GetSeekDelta](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraComponent/GetSeekDelta) () | Gets the delta value which is used when seeking from the current age, to the desired age. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Get Desired Age Seek Delta") |
| [UNiagaraSimCache](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraSimCache) \\* GetSimCache() | Get the active simulation cache, will return null if we do not have an active one. | NiagaraComponent.h | - BlueprintCallable<br>- Category=SimCache |
| FBox [GetSystemFixedBounds](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraComponent/GetSystemFixedBounds) () | Gets the fixed bounds for the system instance. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara |
| [FNiagaraSystemInstance](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraSystemInstance) \\* GetSystemInstance() |  | NiagaraComponent.h |  |
| FNiagaraSystemInstanceControllerPtr [GetSystemInstanceController](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraComponent/GetSystemInstanceController) () |  | NiagaraComponent.h |  |
| FNiagaraSystemInstanceControllerConstPtr [GetSystemInstanceController](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraComponent/GetSystemInstanceController) () |  | NiagaraComponent.h |  |
| [TSharedPtr](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TSharedPtr) < [FNiagaraSystemSimulation](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraSystemSimulation), [ESPMode::ThreadSafe](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/ESPMode) \> GetSystemSimulation() |  | NiagaraComponent.h |  |
| [ENiagaraTickBehavior](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/ENiagaraTickBehavior) GetTickBehavior() |  | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Get Tick Behavior") |
| bool GetVariableBool<br>(<br>[FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName) InVariableName,<br>bool& bIsValid<br>) const | Gets a Niagara bool parameter by name. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Get Niagara Variable (Bool)", Keywords="user parameter variable bool") |
| [FLinearColor](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FLinearColor) GetVariableColor<br>(<br>[FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName) InVariableName,<br>bool& bIsValid<br>) const | Gets a Niagara Linear Color parameter by name. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Get Niagara Variable (LinearColor)", Keywords="user parameter variable color") |
| float GetVariableFloat<br>(<br>[FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName) InVariableName,<br>bool& bIsValid<br>) const | Gets a Niagara float parameter by name. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Get Niagara Variable (Float)", Keywords="user parameter variable float") |
| int32 GetVariableInt<br>(<br>[FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName) InVariableName,<br>bool& bIsValid<br>) const | Gets a Niagara int parameter by name. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Get Niagara Variable (Int32)", Keywords="user parameter variable int") |
| FMatrix GetVariableMatrix<br>(<br>[FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName) InVariableName,<br>bool& bIsValid<br>) const | Gets a Niagara matrix parameter by name. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Get Niagara Variable (Matrix)", Keywords="user parameter variable matrix") |
| FVector GetVariablePosition<br>(<br>[FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName) InVariableName,<br>bool& bIsValid<br>) const | Gets a Niagara Position parameter by name. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Get Niagara Variable (Position)", Keywords="user parameter variable vector position lwc") |
| FQuat GetVariableQuat<br>(<br>[FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName) InVariableName,<br>bool& bIsValid<br>) const | Gets a Niagara quaternion parameter by name. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Get Niagara Variable (Quaternion)", Keywords="user parameter variable quaternion rotation") |
| FVector2D GetVariableVec2<br>(<br>[FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName) InVariableName,<br>bool& bIsValid<br>) const | Gets a Niagara Vector2 parameter by name. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Get Niagara Variable (Vector2)", Keywords="user parameter variable vector") |
| FVector GetVariableVec3<br>(<br>[FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName) InVariableName,<br>bool& bIsValid<br>) const | Gets a Niagara Vector3 parameter by name. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Get Niagara Variable (Vector3)", Keywords="user parameter variable vector") |
| FVector4 GetVariableVec4<br>(<br>[FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName) InVariableName,<br>bool& bIsValid<br>) const | Gets a Niagara Vector4 parameter by name. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Get Niagara Variable (Vector4)", Keywords="user parameter variable vector") |
| bool HasParameterOverride<br>(<br>const [FNiagaraVariableBase](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraVariableBase)& InKey<br>) const |  | NiagaraComponent.h |  |
| void [InitForPerformanceBaseline](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraComponent/InitForPerformanceBaseline) () | Initializes this component for capturing a performance baseline. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Performance<br>- Meta=(Keywords="Niagara Performance") |
| bool InitializeSystem() |  | NiagaraComponent.h |  |
| bool IsActiveForUpdateContext() |  | NiagaraComponent.h |  |
| bool IsComplete() |  | NiagaraComponent.h |  |
| bool IsLocalPlayerEffect() | Is this an effect on or linked to the local player. | NiagaraComponent.h |  |
| bool IsPaused() |  | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara |
| bool IsRegisteredWithScalabilityManager() |  | NiagaraComponent.h |  |
| bool IsUsingCullProxy() |  | NiagaraComponent.h |  |
| bool IsWorldReadyToRun() |  | NiagaraComponent.h |  |
| void OnPooledReuse<br>(<br>[UWorld](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/UWorld)\\* NewWorld<br>) |  | NiagaraComponent.h |  |
| FOnSynchronizedWithAssetParameters & OnSynchronizedWithAssetParameters() |  | NiagaraComponent.h |  |
| FOnSystemInstanceChanged & OnSystemInstanceChanged() |  | NiagaraComponent.h |  |
| void PostLoadNormalizeOverrideNames() |  | NiagaraComponent.h |  |
| void ReinitializeSystem() | Called on when an external object wishes to force this System to reinitialize itself from the System data. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Reinitialize System") |
| void RemoveParameterOverride<br>(<br>const [FNiagaraVariableBase](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraVariableBase)& InKey<br>) | Remove an override for a given parameter if one exists. | NiagaraComponent.h |  |
| void ResetSystem() | Resets the System to it's initial pre-simulated state. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Reset System") |
| bool ResolveOwnerAllowsScalability<br>(<br>bool bRegister<br>) |  | NiagaraComponent.h |  |
| void [SeekToDesiredAge](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraComponent/SeekToDesiredAge)<br>(<br>float InDesiredAge<br>) | Sets the desired age of the System instance and designates that this change is a seek. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Seek to Desired Age") |
| void SetAgeUpdateMode<br>(<br>[ENiagaraAgeUpdateMode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/ENiagaraAgeUpdateMode) InAgeUpdateMode<br>) | Sets the age update mode for the System instance. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Set Age Update Mode") |
| void [SetAllowScalability](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraComponent/SetAllowScalability)<br>(<br>bool bAllow<br>) | Set whether this component is allowed to perform scalability checks and potentially be culled etc. | NiagaraComponent.h | - BlueprintSetter<br>- Category=Scalability<br>- Meta=(Keywords="LOD scalability") |
| void [SetAsset](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraComponent/SetAsset)<br>(<br>[UNiagaraSystem](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraSystem)\\* InAsset,<br>bool bResetExistingOverrideParameters<br>) | Switch which asset the component is using. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Set Niagara System Asset") |
| void SetAutoDestroy<br>(<br>bool bInAutoDestroy<br>) |  | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Set Auto Destroy") |
| void SetCanRenderWhileSeeking<br>(<br>bool bInCanRenderWhileSeeking<br>) | Sets whether or not the system can render while seeking. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Set Can Render While Seeking") |
| void SetCreateSceneComponentUtilsFunction<br>(<br>[TFunction](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TFunction) < [INiagaraSceneComponentUtils](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/INiagaraSceneComponentUtils)\*()\> Function<br>) |  | NiagaraComponent.h |  |
| void [SetCustomTimeDilation](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraComponent/SetCustomTimeDilation)<br>(<br>float Dilation<br>) | Sets the custom time dilation value for the component. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara |
| void [SetDesiredAge](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraComponent/SetDesiredAge)<br>(<br>float InDesiredAge<br>) | Sets the desired age of the System instance. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Set Desired Age") |
| void [SetEmitterFixedBounds](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraComponent/SetEmitterFixedBounds)<br>(<br>[FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName) EmitterName,<br>FBox LocalBounds<br>) | Sets the fixed bounds for an emitter instance, this overrides all other bounds. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara |
| void SetForceLocalPlayerEffect<br>(<br>bool bIsPlayerEffect<br>) |  | NiagaraComponent.h | - BlueprintSetter<br>- Category=Scalability<br>- Meta=(Keywords="LOD scalability") |
| void SetForceSolo<br>(<br>bool bInForceSolo<br>) |  | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Set Forced Solo Mode") |
| void SetGpuComputeDebug<br>(<br>bool bEnableDebug<br>) |  | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara |
| void [SetLockDesiredAgeDeltaTimeToSeekDelta](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraComponent/SetLockDesiredAgeDeltaTimeToSeek-)<br>(<br>bool bLock<br>) | Sets whether or not the delta time used to tick the system instance when using desired age is locked to the seek delta. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara |
| void SetLODDistance<br>(<br>float InLODDistance,<br>float InMaxLODDistance<br>) |  | NiagaraComponent.h |  |
| void [SetMaxSimTime](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraComponent/SetMaxSimTime)<br>(<br>float InMaxTime<br>) | Sets the maximum CPU time in seconds we will simulate to the desired age, when we go beyond this limit ticks will be processed in the next frame. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Set Max Desired Age Tick Delta") |
| void SetNiagaraVariableActor<br>(<br>const [FString](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FString)& InVariableName,<br>[AActor](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/AActor)\\* Actor<br>) |  | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Set Niagara Variable By String (Actor)", Keywords="user parameter variable actor") |
| void SetNiagaraVariableBool<br>(<br>const [FString](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FString)& InVariableName,<br>bool InValue<br>) | Sets a Niagara bool parameter by name, overriding locally if necessary. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Set Niagara Variable By String (Bool)", Keywords="user parameter variable bool") |
| void SetNiagaraVariableFloat<br>(<br>const [FString](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FString)& InVariableName,<br>float InValue<br>) | Sets a Niagara float parameter by name, overriding locally if necessary. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Set Niagara Variable By String (Float)", Keywords="user parameter variable float") |
| void SetNiagaraVariableInt<br>(<br>const [FString](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FString)& InVariableName,<br>int32 InValue<br>) | Sets a Niagara int parameter by name, overriding locally if necessary. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Set Niagara Variable By String (Int32)", Keywords="user parameter variable int") |
| void SetNiagaraVariableLinearColor<br>(<br>const [FString](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FString)& InVariableName,<br>const [FLinearColor](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FLinearColor)& InValue<br>) | Sets a Niagara [FLinearColor](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FLinearColor) parameter by name, overriding locally if necessary. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Set Niagara Variable By String (LinearColor)", Keywords="user parameter variable color") |
| void SetNiagaraVariableMatrix<br>(<br>const [FString](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FString)& InVariableName,<br>const FMatrix& InValue<br>) | Sets a Niagara matrix parameter by name, overriding locally if necessary. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Set Niagara Variable By String (Matrix)", Keywords="user parameter variable matrix") |
| void SetNiagaraVariableObject<br>(<br>const [FString](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FString)& InVariableName,<br>[UObject](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObject)\\* Object<br>) |  | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Set Niagara Variable By String (Object)", Keywords="user parameter variable object") |
| void SetNiagaraVariablePosition<br>(<br>const [FString](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FString)& InVariableName,<br>FVector InValue<br>) | Sets a Niagara Position parameter by name, overriding locally if necessary. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Set Niagara Variable By String (Position)", Keywords="user parameter variable vector position lwc") |
| void SetNiagaraVariableQuat<br>(<br>const [FString](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FString)& InVariableName,<br>const FQuat& InValue<br>) | Sets a Niagara quaternion parameter by name, overriding locally if necessary. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Set Niagara Variable By String (Quaternion)", Keywords="user parameter variable quaternion rotation") |
| void SetNiagaraVariableVec2<br>(<br>const [FString](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FString)& InVariableName,<br>FVector2D InValue<br>) | Sets a Niagara Vector2 parameter by name, overriding locally if necessary. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Set Niagara Variable By String (Vector2)", Keywords="user parameter variable vector") |
| void SetNiagaraVariableVec3<br>(<br>const [FString](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FString)& InVariableName,<br>FVector InValue<br>) | Sets a Niagara Vector3 parameter by name, overriding locally if necessary. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Set Niagara Variable By String (Vector3)", Keywords="user parameter variable vector") |
| void SetNiagaraVariableVec4<br>(<br>const [FString](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FString)& InVariableName,<br>const FVector4& InValue<br>) | Sets a Niagara Vector4 parameter by name, overriding locally if necessary. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Set Niagara Variable By String (Vector4)", Keywords="user parameter variable vector") |
| void SetOcclusionQueryMode<br>(<br>[ENiagaraOcclusionQueryMode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/ENiagaraOcclusionQueryMode) Mode<br>) |  | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara |
| void SetParameterOverride<br>(<br>const [FNiagaraVariableBase](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraVariableBase)& InKey,<br>const [FNiagaraVariant](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraVariant)& InValue<br>) |  | NiagaraComponent.h |  |
| void SetPaused<br>(<br>bool bInPaused<br>) |  | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara |
| void SetPreviewLODDistance<br>(<br>bool bEnablePreviewLODDistance,<br>float PreviewLODDistance,<br>float PreviewMaxDistance<br>) |  | NiagaraComponent.h | - BlueprintCallable<br>- Category=Preview<br>- Meta=(Keywords="preview LOD Distance scalability") |
| void SetRandomSeedOffset<br>(<br>int32 NewRandomSeedOffset<br>) |  | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Set Random Seed Offset") |
| void SetRenderingEnabled<br>(<br>bool bInRenderingEnabled<br>) | Sets whether or not rendering is enabled for this component. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Set Rendering Enabled") |
| void [SetSeekDelta](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraComponent/SetSeekDelta)<br>(<br>float InSeekDelta<br>) | Sets the delta value which is used when seeking from the current age, to the desired age. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Set Desired Age Seek Delta") |
| void [SetSimCache](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraComponent/SetSimCache)<br>(<br>[UNiagaraSimCache](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraSimCache)\\* SimCache,<br>bool bResetSystem<br>) | Sets the simulation cache to use for the component. | NiagaraComponent.h | - BlueprintCallable<br>- Category=SimCache |
| void [SetSystemFixedBounds](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraComponent/SetSystemFixedBounds)<br>(<br>FBox LocalBounds<br>) | Sets the fixed bounds for the system instance, this overrides all other bounds. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara |
| void [SetSystemSignificanceIndex](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraComponent/SetSystemSignificanceIndex)<br>(<br>int32 InIndex<br>) | The significant index for this component. | NiagaraComponent.h |  |
| void SetTickBehavior<br>(<br>[ENiagaraTickBehavior](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/ENiagaraTickBehavior) NewTickBehavior<br>) |  | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Set Tick Behavior") |
| void SetUserParametersToDefaultValues() | Removes all local overrides and replaces them with the values from the source System - note: this also removes the editor overrides from the component as it is used by the pooling mechanism to prevent values leaking between different instances. | NiagaraComponent.h |  |
| void SetVariableActor<br>(<br>[FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName) InVariableName,<br>[AActor](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/AActor)\\* Actor<br>) |  | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Set Niagara Variable (Actor)", Keywords="user parameter variable actor") |
| void SetVariableBool<br>(<br>[FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName) InVariableName,<br>bool InValue<br>) | Sets a Niagara bool parameter by name, overriding locally if necessary. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Set Niagara Variable (Bool)", Keywords="user parameter variable bool") |
| void SetVariableFloat<br>(<br>[FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName) InVariableName,<br>float InValue<br>) | Sets a Niagara float parameter by name, overriding locally if necessary. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Set Niagara Variable (Float)", Keywords="user parameter variable float") |
| void SetVariableInt<br>(<br>[FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName) InVariableName,<br>int32 InValue<br>) | Sets a Niagara int parameter by name, overriding locally if necessary. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Set Niagara Variable (Int32)", Keywords="user parameter variable int") |
| void SetVariableLinearColor<br>(<br>[FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName) InVariableName,<br>const [FLinearColor](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FLinearColor)& InValue<br>) | Sets a Niagara [FLinearColor](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FLinearColor) parameter by name, overriding locally if necessary. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Set Niagara Variable (LinearColor)", Keywords="user parameter variable color") |
| void SetVariableMaterial<br>(<br>[FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName) InVariableName,<br>[UMaterialInterface](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/UMaterialInterface)\\* Object<br>) |  | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Set Niagara Variable (Material)", Keywords="user parameter variable material") |
| void SetVariableMatrix<br>(<br>[FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName) InVariableName,<br>const FMatrix& InValue<br>) | Sets a Niagara matrix parameter by name, overriding locally if necessary. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Set Niagara Variable (Matrix)", Keywords="user parameter variable matrix") |
| void SetVariableObject<br>(<br>[FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName) InVariableName,<br>[UObject](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObject)\\* Object<br>) |  | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Set Niagara Variable (Object)", Keywords="user parameter variable object") |
| void SetVariablePosition<br>(<br>[FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName) InVariableName,<br>FVector InValue<br>) | Sets a Niagara Position parameter by name, overriding locally if necessary. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Set Niagara Variable (Position)", Keywords="user parameter variable vector position lwc") |
| void SetVariableQuat<br>(<br>[FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName) InVariableName,<br>const FQuat& InValue<br>) | Sets a Niagara quaternion parameter by name, overriding locally if necessary. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Set Niagara Variable (Quaternion)", Keywords="user parameter variable quaternion rotation") |
| void SetVariableStaticMesh<br>(<br>[FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName) InVariableName,<br>[UStaticMesh](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/UStaticMesh)\\* InValue<br>) |  | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Set Niagara Variable (Static Mesh)", Keywords="user parameter variable mesh") |
| void SetVariableTexture<br>(<br>[FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName) InVariableName,<br>[UTexture](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/UTexture)\\* Texture<br>) |  | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Set Niagara Variable (Texture)", Keywords="user parameter variable texture") |
| void SetVariableTextureRenderTarget<br>(<br>[FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName) InVariableName,<br>[UTextureRenderTarget](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/UTextureRenderTarget)\\* TextureRenderTarget<br>) |  | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Set Niagara Variable (TextureRenderTarget)") |
| void SetVariableVec2<br>(<br>[FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName) InVariableName,<br>FVector2D InValue<br>) | Sets a Niagara Vector2 parameter by name, overriding locally if necessary. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Set Niagara Variable (Vector2)", Keywords="user parameter variable vector") |
| void SetVariableVec3<br>(<br>[FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName) InVariableName,<br>FVector InValue<br>) | Sets a Niagara Vector3 parameter by name, overriding locally if necessary. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Set Niagara Variable (Vector3)", Keywords="user parameter variable vector") |
| void SetVariableVec4<br>(<br>[FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName) InVariableName,<br>const FVector4& InValue<br>) | Sets a Niagara Vector4 parameter by name, overriding locally if necessary. | NiagaraComponent.h | - BlueprintCallable<br>- Category=Niagara<br>- Meta=(DisplayName="Set Niagara Variable (Vector4)", Keywords="user parameter variable vector") |
| void UpgradeDeprecatedParameterOverrides() |  | NiagaraComponent.h |  |

#### Overridden from [UFXSystemComponent](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/UFXSystemComponent)

| Name | Remarks | Include Path | [Unreal Specifiers](https://dev.epicgames.com/documentation/unreal-engine/ufunctions-in-unreal-engine#functionspecifiers) |
| --- | --- | --- | --- |
| virtual void ActivateSystem<br>(<br>bool bFlagAsJustAttached<br>) |  | NiagaraComponent.h |  |
| virtual void DeactivateImmediate() |  | NiagaraComponent.h |  |
| virtual uint32 GetApproxMemoryUsage() |  | NiagaraComponent.h |  |
| virtual [UFXSystemAsset](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/UFXSystemAsset) \\* GetFXSystemAsset() |  | NiagaraComponent.h |  |
| virtual void ReleaseToPool() |  | NiagaraComponent.h |  |
| virtual void SetActorParameter<br>(<br>[FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName) ParameterName,<br>[AActor](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/AActor)\\* Param<br>) |  | NiagaraComponent.h |  |
| virtual void [SetAutoAttachmentParameters](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraComponent/SetAutoAttachmentParameters)<br>(<br>[USceneComponent](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/USceneComponent)\\* Parent,<br>[FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName) SocketName,<br>[EAttachmentRule](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/EAttachmentRule) LocationRule,<br>[EAttachmentRule](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/EAttachmentRule) RotationRule,<br>[EAttachmentRule](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/EAttachmentRule) ScaleRule<br>) | Set AutoAttachParent, AutoAttachSocketName, AutoAttachLocationRule, AutoAttachRotationRule, AutoAttachScaleRule to the specified parameters. | NiagaraComponent.h |  |
| virtual void SetBoolParameter<br>(<br>[FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName) ParameterName,<br>bool Param<br>) |  | NiagaraComponent.h |  |
| virtual void SetColorParameter<br>(<br>[FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName) ParameterName,<br>[FLinearColor](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FLinearColor) Param<br>) |  | NiagaraComponent.h |  |
| virtual void SetEmitterEnable<br>(<br>[FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName) EmitterName,<br>bool bNewEnableState<br>) |  | NiagaraComponent.h |  |
| virtual void SetFloatParameter<br>(<br>[FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName) ParameterName,<br>float Param<br>) |  | NiagaraComponent.h |  |
| virtual void SetIntParameter<br>(<br>[FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName) ParameterName,<br>int Param<br>) |  | NiagaraComponent.h |  |
| virtual void SetUseAutoManageAttachment<br>(<br>bool bAutoManage<br>) |  | NiagaraComponent.h |  |
| virtual void SetVectorParameter<br>(<br>[FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName) ParameterName,<br>FVector Param<br>) |  | NiagaraComponent.h |  |

#### Overridden from [UPrimitiveComponent](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/UPrimitiveComponent)

| Name | Remarks | Include Path | [Unreal Specifiers](https://dev.epicgames.com/documentation/unreal-engine/ufunctions-in-unreal-engine#functionspecifiers) |
| --- | --- | --- | --- |
| virtual void CollectPSOPrecacheData<br>(<br>const [FPSOPrecacheParams](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/FPSOPrecacheParams)& BasePrecachePSOParams,<br>FMaterialInterfacePSOPrecacheParamsList& OutParams<br>) |  | NiagaraComponent.h |  |
| virtual [FPrimitiveSceneProxy](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/FPrimitiveSceneProxy) \\* CreateSceneProxy() |  | NiagaraComponent.h |  |
| virtual int32 GetNumMaterials() |  | NiagaraComponent.h |  |
| virtual void GetStreamingRenderAssetInfo<br>(<br>[FStreamingTextureLevelContext](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/FStreamingTextureLevelContext)& LevelContext,<br>[TArray](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArray) < [FStreamingRenderAssetPrimitiveInfo](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/FStreamingRenderAssetPrimitiveIn-) >& OutStreamingRenderAssets<br>) const |  | NiagaraComponent.h |  |
| virtual void GetUsedMaterials<br>(<br>[TArray](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArray) < [UMaterialInterface](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/UMaterialInterface)\\* >& OutMaterials,<br>bool bGetDebugMaterials<br>) const |  | NiagaraComponent.h |  |

#### Overridden from [USceneComponent](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/USceneComponent)

| Name | Remarks | Include Path | [Unreal Specifiers](https://dev.epicgames.com/documentation/unreal-engine/ufunctions-in-unreal-engine#functionspecifiers) |
| --- | --- | --- | --- |
| virtual FBoxSphereBounds CalcBounds<br>(<br>const FTransform& LocalToWorld<br>) const |  | NiagaraComponent.h |  |
| virtual bool IsVisible() |  | NiagaraComponent.h |  |
| virtual void OnAttachmentChanged() |  | NiagaraComponent.h |  |
| virtual void OnChildAttached<br>(<br>[USceneComponent](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/USceneComponent)\\* ChildComponent<br>) |  | NiagaraComponent.h |  |
| virtual void OnChildDetached<br>(<br>[USceneComponent](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/USceneComponent)\\* ChildComponent<br>) |  | NiagaraComponent.h |  |

#### Overridden from [UActorComponent](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/UActorComponent)

| Name | Remarks | Include Path | [Unreal Specifiers](https://dev.epicgames.com/documentation/unreal-engine/ufunctions-in-unreal-engine#functionspecifiers) |
| --- | --- | --- | --- |
| virtual void Activate<br>(<br>bool bReset<br>) |  | NiagaraComponent.h |  |
| virtual const [UObject](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObject) \\* AdditionalStatObject() |  | NiagaraComponent.h |  |
| virtual void Deactivate() |  | NiagaraComponent.h |  |
| virtual bool IsReadyForOwnerToAutoDestroy() |  | NiagaraComponent.h |  |
| virtual void OnComponentCreated() |  | NiagaraComponent.h |  |
| virtual void OnComponentDestroyed<br>(<br>bool bDestroyingHierarchy<br>) |  | NiagaraComponent.h |  |
| virtual void PostApplyToComponent() |  | NiagaraComponent.h |  |
| virtual bool RequiresGameThreadEndOfFrameRecreate() |  | NiagaraComponent.h |  |
| virtual void TickComponent<br>(<br>float DeltaTime,<br>enum [ELevelTick](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/ELevelTick) TickType,<br>[FActorComponentTickFunction](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/FActorComponentTickFunction)\\* ThisTickFunction<br>) |  | NiagaraComponent.h |  |

#### Overridden from [UObject](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObject)

| Name | Remarks | Include Path | [Unreal Specifiers](https://dev.epicgames.com/documentation/unreal-engine/ufunctions-in-unreal-engine#functionspecifiers) |
| --- | --- | --- | --- |
| virtual bool CanEditChange<br>(<br>const [FProperty](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/FProperty)\\* InProperty<br>) const |  | NiagaraComponent.h |  |
| virtual void GetResourceSizeEx<br>(<br>[FResourceSizeEx](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FResourceSizeEx)& CumulativeResourceSize<br>) |  | NiagaraComponent.h |  |
| virtual void PostEditChangeProperty<br>(<br>[FPropertyChangedEvent](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/FPropertyChangedEvent)& PropertyChangedEvent<br>) |  | NiagaraComponent.h |  |
| virtual void PostLoad() |  | NiagaraComponent.h |  |
| virtual void PreEditChange<br>(<br>[FProperty](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/FProperty)\\* PropertyAboutToChange<br>) |  | NiagaraComponent.h |  |
| virtual void Serialize<br>(<br>FStructuredArchive::FRecord Record<br>) |  | NiagaraComponent.h |  |

#### Overridden from [UObjectBase](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObjectBase)

| Name | Remarks | Include Path | [Unreal Specifiers](https://dev.epicgames.com/documentation/unreal-engine/ufunctions-in-unreal-engine#functionspecifiers) |
| --- | --- | --- | --- |
| virtual [FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName) GetFNameForStatID() |  | NiagaraComponent.h |  |

### Protected

#### Overridden from [UActorComponent](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/UActorComponent)

| Name | Remarks | Include Path | [Unreal Specifiers](https://dev.epicgames.com/documentation/unreal-engine/ufunctions-in-unreal-engine#functionspecifiers) |
| --- | --- | --- | --- |
| virtual void ApplyWorldOffset<br>(<br>const FVector& InOffset,<br>bool bWorldShift<br>) |  | NiagaraComponent.h |  |
| virtual void CreateRenderState\_Concurrent<br>(<br>[FRegisterComponentContext](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/FRegisterComponentContext)\\* Context<br>) |  | NiagaraComponent.h |  |
| virtual void DestroyRenderState\_Concurrent() |  | NiagaraComponent.h |  |
| virtual void OnEndOfFrameUpdateDuringTick() |  | NiagaraComponent.h |  |
| virtual void OnRegister() |  | NiagaraComponent.h |  |
| virtual void OnUnregister() |  | NiagaraComponent.h |  |
| virtual void SendRenderDynamicData\_Concurrent() |  | NiagaraComponent.h |  |

#### Overridden from [UObject](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObject)

| Name | Remarks | Include Path | [Unreal Specifiers](https://dev.epicgames.com/documentation/unreal-engine/ufunctions-in-unreal-engine#functionspecifiers) |
| --- | --- | --- | --- |
| virtual void BeginDestroy() |  | NiagaraComponent.h |  |

### Static

| Name | Remarks | Include Path | [Unreal Specifiers](https://dev.epicgames.com/documentation/unreal-engine/ufunctions-in-unreal-engine#functionspecifiers) |
| --- | --- | --- | --- |
| static void DeclareConstructClasses<br>(<br>[TArray](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArray) < [FTopLevelAssetPath](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/FTopLevelAssetPath) >& OutConstructClasses,<br>const [UClass](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UClass)\\* SpecificSubclass<br>) |  | NiagaraComponent.h |  |

## See Also

- [ANiagaraActor](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/ANiagaraActor)

- [UNiagaraSystem](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraSystem)

* * *