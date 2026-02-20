<!-- Source: https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraSystem -->

A Niagara System contains multiple Niagara Emitters to create various effects. Niagara Systems can be placed in the world, unlike Emitters, and expose User Parameters to configure an effect at runtime.

|  |  |
| --- | --- |
| _Name_ | UNiagaraSystem |
| _Type_ | class |
| _Header File_ | /Engine/Plugins/FX/Niagara/Source/Niagara/Classes/NiagaraSystem.h |
| _Include Path_ | #include "NiagaraSystem.h" |

## Syntax

```cpp

UCLASS (BlueprintType, MinimalAPI, Meta=(LoadBehavior="LazyOnDemand"))

class UNiagaraSystem :

    public UFXSystemAsset ,

    public INiagaraParameterDefinitionsSubscriber
```

## Inheritance Hierarchy

- [UObjectBase](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObjectBase) → [UObjectBaseUtility](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObjectBaseUtility) → [UObject](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObject) → [UFXSystemAsset](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/UFXSystemAsset) → **UNiagaraSystem**

## Implements Interfaces

- [INiagaraParameterDefinitionsSubscriber](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/INiagaraParameterDefinitionsSubs-)

## Constructors

| Name | Remarks | Include Path | [Unreal Specifiers](https://dev.epicgames.com/documentation/unreal-engine/ufunctions-in-unreal-engine#functionspecifiers) |
| --- | --- | --- | --- |
| [UNiagaraSystem](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraSystem/__ctor)<br>(<br>[FVTableHelper](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/FVTableHelper)& Helper<br>) |  | NiagaraSystem.h |  |
| [UNiagaraSystem](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraSystem/__ctor)<br>(<br>const [FObjectInitializer](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/FObjectInitializer)& ObjectInitializer<br>) |  | NiagaraSystem.h |  |

## Destructors

| Name | Remarks | Include Path | [Unreal Specifiers](https://dev.epicgames.com/documentation/unreal-engine/ufunctions-in-unreal-engine#functionspecifiers) |
| --- | --- | --- | --- |
| ~UNiagaraSystem() |  | NiagaraSystem.h |  |

## Structs

| Name | Remarks |
| --- | --- |
| [FStaticBuffersDeletor](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraSystem/FStaticBuffersDeletor) |  |

## Enums

### Protected

| Name | Remarks |
| --- | --- |
| [ERequestCompileStatus](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraSystem/ERequestCompileStatus) | Used to track the status of compilation request, when using on demand they will initially be in RequestPendingOnDemand. |

## Typedefs

| Name | Type | Remarks | Include Path |
| --- | --- | --- | --- |
| FOnScalabilityChanged | TMulticastDelegate\_NoParams< void > |  | NiagaraSystem.h |
| FOnSystemCompiled | TMulticastDelegate\_OneParam< void, [UNiagaraSystem](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraSystem) \\* > |  | NiagaraSystem.h |
| FOnSystemPostEditChange | TMulticastDelegate\_OneParam< void, [UNiagaraSystem](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraSystem) \\* > |  | NiagaraSystem.h |

## Constants

| Name | Type | Remarks | Include Path |
| --- | --- | --- | --- |
| ComputeEmitterExecutionOrderMessageId | const [FGuid](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FGuid) |  | NiagaraSystem.h |
| kStartNewOverlapGroupBit | int32 | When an index inside the EmitterExecutionOrder array has this bit set, it means the corresponding emitter cannot execute in parallel with the previous emitters due to a data dependency. | NiagaraSystem.h |
| ResolveDIsMessageId | const [FGuid](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FGuid) |  | NiagaraSystem.h |

## Variables

### Public

| Name | Type | Remarks | Include Path | [Unreal Specifiers](https://dev.epicgames.com/documentation/unreal-engine/unreal-engine-uproperties#propertyspecifiers) |
| --- | --- | --- | --- | --- |
| bCastShadow | uint8 | When enabled this is the default value set on the component. | NiagaraSystem.h | - EditAnywhere<br>- BlueprintReadOnly<br>- Category="Rendering"<br>- Meta=(DisplayName="Default Cast Shadows", EditCondition="bOverrideCastShadow") |
| bDumpDebugEmitterInfo | uint8 |  | NiagaraSystem.h | - EditAnywhere<br>- Category="Debug"<br>- Transient<br>- AdvancedDisplay |
| bDumpDebugSystemInfo | uint8 |  | NiagaraSystem.h | - EditAnywhere<br>- Category="Debug"<br>- Transient<br>- AdvancedDisplay |
| bFixedBounds | uint32 | Whether or not fixed bounds are enabled. | NiagaraSystem.h | - EditAnywhere<br>- Category="System"<br>- Meta=(SkipSystemResetOnChange="true", InlineEditConditionToggle) |
| bFullyLoaded | uint8 |  | NiagaraSystem.h |  |
| bOverrideCastShadow | uint8 | Various optional overrides for component properties when spawning a system. | NiagaraSystem.h | - EditAnywhere<br>- Category="Rendering"<br>- Meta=(InlineEditConditionToggle="bCastShadow") |
| bOverrideCustomDepthStencilValue | uint8 |  | NiagaraSystem.h | - EditAnywhere<br>- Category="Rendering"<br>- Meta=(InlineEditConditionToggle="CustomDepthStencilValue") |
| bOverrideCustomDepthStencilWriteMask | uint8 |  | NiagaraSystem.h | - EditAnywhere<br>- Category="Rendering"<br>- Meta=(InlineEditConditionToggle="CustomDepthStencilWriteMask") |
| bOverrideReceivesDecals | uint8 |  | NiagaraSystem.h | - EditAnywhere<br>- Category="Rendering"<br>- Meta=(InlineEditConditionToggle="bReceivesDecals") |
| bOverrideRenderCustomDepth | uint8 |  | NiagaraSystem.h | - EditAnywhere<br>- Category="Rendering"<br>- Meta=(InlineEditConditionToggle="bRenderCustomDepth") |
| bOverrideTranslucencySortDistanceOffset | uint8 |  | NiagaraSystem.h | - EditAnywhere<br>- Category="Rendering"<br>- Meta=(InlineEditConditionToggle="TranslucencySortDistanceOffset") |
| bOverrideTranslucencySortPriority | uint8 |  | NiagaraSystem.h | - EditAnywhere<br>- Category="Rendering"<br>- Meta=(InlineEditConditionToggle="TranslucencySortPriority") |
| bReceivesDecals | uint8 | When enabled this is the default value set on the component. Whether the primitive receives decals. | NiagaraSystem.h | - EditAnywhere<br>- BlueprintReadOnly<br>- Category="Rendering"<br>- Meta=(DisplayName="Default Receives Decals", EditCondition="bOverrideReceivesDecals") |
| bRenderCustomDepth | uint8 | When enabled this is the default value set on the component. | NiagaraSystem.h | - EditAnywhere<br>- BlueprintReadOnly<br>- Category="Rendering"<br>- Meta=(DisplayName="Default Render CustomDepth Pass", EditCondition="bOverrideRenderCustomDepth") |
| bRequireCurrentFrameData | uint8 | When enabled, we follow the settings on the [UNiagaraComponent](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraComponent) for tick order. | NiagaraSystem.h | - EditAnywhere<br>- Category="Performance"<br>- AdvancedDisplay |
| bSupportLargeWorldCoordinates | uint8 | If true then position type values will be rebased on system activation to fit into a float precision vector. | NiagaraSystem.h | - EditAnywhere<br>- AdvancedDisplay<br>- Category="Rendering" |
| Category | [FText](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FText) | Category of this system. | NiagaraSystem.h | - AssetRegistrySearchable<br>- Meta=(SkipSystemResetOnChange="true") |
| CustomDepthStencilValue | int32 | When enabled this is the default value set on the component. | NiagaraSystem.h | - EditAnywhere<br>- BlueprintReadOnly<br>- AdvancedDisplay<br>- Category="Rendering"<br>- Meta=(DisplayName="Default CustomDepthStencil Value", editcondition="bOverrideCustomDepthStencilWriteMask", UIMin="0", UIMax="255") |
| CustomDepthStencilWriteMask | [ERendererStencilMask](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/ERendererStencilMask) | When enabled this is the default value set on the component. Mask used for stencil buffer writes. | NiagaraSystem.h | - EditAnywhere<br>- BlueprintReadOnly<br>- AdvancedDisplay<br>- Category="Rendering"<br>- Meta=(DisplayName="Default CustomDepthStencil Write Mask", editcondition="bOverrideCustomDepthStencilValue") |
| EditorOnlyAddedParameters | [FNiagaraParameterStore](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraParameterStore) |  | NiagaraSystem.h | - Transient |
| LibraryVisibility | [ENiagaraScriptLibraryVisibility](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/ENiagaraScriptLibraryVisibility) | If this system is exposed to the library, or should be explicitly hidden. | NiagaraSystem.h | - AssetRegistrySearchable<br>- Meta=(SkipSystemResetOnChange="true") |
| ParameterDefinitionsSubscriptions | [TArray](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArray) < [FParameterDefinitionsSubscription](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FParameterDefinitionsSubscriptio-) > | Subscriptions to definitions of parameters. | NiagaraSystem.h |  |
| PreviewMoviePath | [FSoftObjectPath](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/FSoftObjectPath) |  | NiagaraSystem.h | - EditAnywhere<br>- AdvancedDisplay<br>- Category="Asset Options"<br>- AssetRegistrySearchable<br>- Meta=(SkipSystemResetOnChange="true", AllowedClasses="/Script/MediaAssets.FileMediaSource") |
| ScratchPadScripts | [TArray](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArray) < [TObjectPtr](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TObjectPtr) < [UNiagaraScript](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraScript) \> > |  | NiagaraSystem.h |  |
| TemplateAssetDescription | [FText](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FText) |  | NiagaraSystem.h | - EditAnywhere<br>- AdvancedDisplay<br>- Category="Asset Options"<br>- AssetRegistrySearchable<br>- DisplayName="Asset Description"<br>- Meta=(SkipSystemResetOnChange="true") |
| ThumbnailImage | [TObjectPtr](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TObjectPtr) < class [UTexture2D](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/UTexture2D) > | Internal: The thumbnail image. | NiagaraSystem.h |  |
| TranslucencySortDistanceOffset | float | When enabled this is the default value set on the component. | NiagaraSystem.h | - EditAnywhere<br>- BlueprintReadOnly<br>- AdvancedDisplay<br>- Category="Rendering"<br>- Meta=(editcondition="bOverrideTranslucencySortDistanceOffset") |
| TranslucencySortPriority | int32 | When enabled this is the default value set on the component. | NiagaraSystem.h | - EditAnywhere<br>- BlueprintReadOnly<br>- AdvancedDisplay<br>- Category="Rendering"<br>- Meta=(editcondition="bOverrideTranslucencySortPriority") |
| UpdateContext | [FNiagaraSystemUpdateContext](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraSystemUpdateContext) |  | NiagaraSystem.h | - Transient |

### Protected

| Name | Type | Remarks | Include Path | [Unreal Specifiers](https://dev.epicgames.com/documentation/unreal-engine/unreal-engine-uproperties#propertyspecifiers) |
| --- | --- | --- | --- | --- |
| bInitialOwnerVelocityFromActor | uint8 | When enabled we use the owner actor's velocity for the first frame. | NiagaraSystem.h | - EditAnywhere<br>- AdvancedDisplay<br>- Category="System" |
| LargeWorldCoordinateTileUpdateMode | [TOptional](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TOptional) < [ENiagaraLwcTileUpdateMode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/ENiagaraLwcTileUpdateMode) > |  | NiagaraSystem.h | - EditAnywhere<br>- Category=Rendering<br>- AdvancedDisplay<br>- Meta=(EditCondition="bSupportLargeWorldCoordinates", EditConditionHides, DisplayAfter=bSupportLargeWorldCoordinates) |

## Functions

### Public

| Name | Remarks | Include Path | [Unreal Specifiers](https://dev.epicgames.com/documentation/unreal-engine/ufunctions-in-unreal-engine#functionspecifiers) |
| --- | --- | --- | --- |
| [FNiagaraEmitterHandle](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraEmitterHandle) [AddEmitterHandle](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraSystem/AddEmitterHandle)<br>(<br>[UNiagaraEmitter](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraEmitter)& SourceEmitter,<br>[FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName) EmitterName,<br>[FGuid](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FGuid) EmitterVersion<br>) | Adds a new emitter handle to this System. | NiagaraSystem.h |  |
| void [AddEmitterHandleDirect](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraSystem/AddEmitterHandleDirect)<br>(<br>[FNiagaraEmitterHandle](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraEmitterHandle)& EmitterHandleToAdd<br>) | Adds a new emitter handle to this system without copying the original asset. | NiagaraSystem.h |  |
| void AddToInstanceCountStat<br>(<br>int32 NumInstances,<br>bool bSolo<br>) const |  | NiagaraSystem.h |  |
| bool AllDIsPostSimulateCanOverlapFrames() |  | NiagaraSystem.h |  |
| bool AllowCullingForLocalPlayers() |  | NiagaraSystem.h |  |
| bool AllowScalabilityForLocalPlayerFX() |  | NiagaraSystem.h |  |
| bool AllowValidation() |  | NiagaraSystem.h |  |
| bool AsyncWorkCanOverlapTickGroups() |  | NiagaraSystem.h |  |
| void CacheFromCompiledData() | Cache data & accessors from the compiled data, allows us to avoid per instance. | NiagaraSystem.h |  |
| bool CanObtainEmitterAttribute<br>(<br>const [FNiagaraVariableBase](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraVariableBase)& InVarWithUniqueNameNamespace,<br>[FNiagaraTypeDefinition](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraTypeDefinition)& OutBoundType<br>) const |  | NiagaraSystem.h |  |
| bool CanObtainSystemAttribute<br>(<br>const [FNiagaraVariableBase](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraVariableBase)& InVar,<br>[FNiagaraTypeDefinition](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraTypeDefinition)& OutBoundType<br>) const |  | NiagaraSystem.h |  |
| bool CanObtainUserVariable<br>(<br>const [FNiagaraVariableBase](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraVariableBase)& InVar<br>) const |  | NiagaraSystem.h |  |
| virtual bool ChangeEmitterVersion<br>(<br>const [FVersionedNiagaraEmitter](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FVersionedNiagaraEmitter)& Emitter,<br>const [FGuid](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FGuid)& NewVersion<br>) |  | NiagaraSystem.h |  |
| bool CompileRequestsShouldBlockGC() |  | NiagaraSystem.h |  |
| bool ComputeEmitterPriority<br>(<br>int32 EmitterIdx,<br>[TArray](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArray) < int32, TInlineAllocator< 32 > >& EmitterPriorities,<br>const [TBitArray](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TBitArray) < TInlineAllocator< 32 > >& EmitterDependencyGraph<br>) | Computes emitter priorities based on the dependency information. | NiagaraSystem.h |  |
| void ComputeEmittersExecutionOrder() | Computes the order in which the emitters in the Emitters array will be ticked and stores the results in EmitterExecutionOrder. | NiagaraSystem.h |  |
| void ComputeRenderersDrawOrder() | Computes the order in which renderers will render | NiagaraSystem.h |  |
| [FNiagaraEmitterHandle](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraEmitterHandle) [DuplicateEmitterHandle](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraSystem/DuplicateEmitterHandle)<br>(<br>const [FNiagaraEmitterHandle](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraEmitterHandle)& EmitterHandleToDuplicate,<br>[FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName) EmitterName<br>) | Duplicates an existing emitter handle and adds it to the System. | NiagaraSystem.h |  |
| void EnsureFullyLoaded() |  | NiagaraSystem.h |  |
| void FindDataInterfaceDependencies<br>(<br>[FVersionedNiagaraEmitterData](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FVersionedNiagaraEmitterData)\\* EmitterData,<br>[UNiagaraScript](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraScript)\\* Script,<br>[TArray](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArray) < [FVersionedNiagaraEmitter](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FVersionedNiagaraEmitter) >& Dependencies<br>) | Queries all the data interfaces in the array for emitter dependencies. | NiagaraSystem.h |  |
| void FindEventDependencies<br>(<br>[FVersionedNiagaraEmitterData](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FVersionedNiagaraEmitterData)\\* EmitterData,<br>[TArray](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArray) < [FVersionedNiagaraEmitter](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FVersionedNiagaraEmitter) >& Dependencies<br>) | Looks at all the event handlers in the emitter to determine which other emitters it depends on. | NiagaraSystem.h |  |
| void ForceGraphToRecompileOnNextCheck() |  | NiagaraSystem.h |  |
| void [ForEachPlatformSet](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraSystem/ForEachPlatformSet)<br>(<br>TAction Func<br>) | Performs the passed action for all FNiagaraPlatformSets used by this system. | NiagaraSystem.h |  |
| void [ForEachScript](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraSystem/ForEachScript)<br>(<br>TAction Func<br>) const | Performs the passed action for all scripts in this system. | NiagaraSystem.h |  |
| void GatherStaticVariables<br>(<br>[TArray](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArray) < [FNiagaraVariable](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraVariable) >& OutVars,<br>[TArray](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArray) < [FNiagaraVariable](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraVariable) >& OutEmitterVars<br>) const |  | NiagaraSystem.h |  |
| int32 & GetActiveInstancesCount() |  | NiagaraSystem.h |  |
| const [FGuid](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FGuid) & GetAssetGuid() |  | NiagaraSystem.h |  |
| const [UNiagaraBakerSettings](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraBakerSettings) \\* GetBakerGeneratedSettings() |  | NiagaraSystem.h |  |
| [UNiagaraBakerSettings](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraBakerSettings) \\* GetBakerSettings() |  | NiagaraSystem.h |  |
| const [TSharedPtr](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TSharedPtr) < [FNiagaraGraphCachedDataBase](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraGraphCachedDataBase), [ESPMode::ThreadSafe](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/ESPMode) \> & GetCachedTraversalData() | Get the cached parameter map traversal for this emitter. | NiagaraSystem.h |  |
| bool GetCompileForEdit() |  | NiagaraSystem.h |  |
| const [FString](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FString) & GetCrashReporterTag() |  | NiagaraSystem.h |  |
| [ENiagaraCullProxyMode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/ENiagaraCullProxyMode) GetCullProxyMode() |  | NiagaraSystem.h |  |
| const [FNiagaraSystemScalabilityOverride](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraSystemScalabilityOverrid-) & GetCurrentOverrideSettings() |  | NiagaraSystem.h |  |
| const [UNiagaraEditorDataBase](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraEditorDataBase) \\* [GetEditorData](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraSystem/GetEditorData) () | Gets editor specific data stored with this system. | NiagaraSystem.h |  |
| [UNiagaraEditorDataBase](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraEditorDataBase) \\* [GetEditorData](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraSystem/GetEditorData) () | Gets editor specific data stored with this system. | NiagaraSystem.h |  |
| [UNiagaraEditorParametersAdapterBase](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraEditorParametersAdapterB-) \\* GetEditorParameters() | Gets editor specific parameters stored with this system | NiagaraSystem.h |  |
| [UNiagaraEffectType](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraEffectType) \\* GetEffectType() |  | NiagaraSystem.h |  |
| const [TArray](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArray) < [TSharedRef](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TSharedRef) < const [FNiagaraEmitterCompiledData](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraEmitterCompiledData) > \> & GetEmitterCompiledData() |  | NiagaraSystem.h |  |
| TConstArrayView< [FNiagaraEmitterExecutionIndex](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraEmitterExecutionIndex) \> GetEmitterExecutionOrder() |  | NiagaraSystem.h |  |
| TConstArrayView< [FNiagaraDataSetAccessor](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraDataSetAccessor) < [ENiagaraExecutionState](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/ENiagaraExecutionState) > \> GetEmitterExecutionStateAccessors() |  | NiagaraSystem.h |  |
| [FNiagaraEmitterHandle](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraEmitterHandle) & [GetEmitterHandle](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraSystem/GetEmitterHandle)<br>(<br>int Idx<br>) |  | NiagaraSystem.h |  |
| const [FNiagaraEmitterHandle](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraEmitterHandle) & [GetEmitterHandle](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraSystem/GetEmitterHandle)<br>(<br>int Idx<br>) const |  | NiagaraSystem.h |  |
| [TArray](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArray) < [FNiagaraEmitterHandle](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraEmitterHandle) \> & [GetEmitterHandles](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraSystem/GetEmitterHandles) () | Gets an array of the emitter handles. | NiagaraSystem.h |  |
| const [TArray](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArray) < [FNiagaraEmitterHandle](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraEmitterHandle) \> & [GetEmitterHandles](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraSystem/GetEmitterHandles) () |  | NiagaraSystem.h |  |
| TConstArrayView< [FNiagaraDataSetAccessor](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraDataSetAccessor) < [FNiagaraSpawnInfo](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraSpawnInfo) > \> GetEmitterSpawnInfoAccessors<br>(<br>int32 EmitterIndex<br>) const |  | NiagaraSystem.h |  |
| [FNiagaraUserRedirectionParameterStore](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraUserRedirectionParameter-) & [GetExposedParameters](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraSystem/GetExposedParameters) () |  | NiagaraSystem.h |  |
| const [FNiagaraUserRedirectionParameterStore](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraUserRedirectionParameter-) & [GetExposedParameters](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraSystem/GetExposedParameters) () | From the last compile, what are the variables that were exported out of the system for external use? | NiagaraSystem.h |  |
| FBox GetFixedBounds() |  | NiagaraSystem.h |  |
| float GetFixedTickDeltaTime() |  | NiagaraSystem.h |  |
| FBox GetInitialStreamingBounds() |  | NiagaraSystem.h |  |
| bool GetIsolateEnabled() |  | NiagaraSystem.h |  |
| [ENiagaraLwcTileUpdateMode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/ENiagaraLwcTileUpdateMode) GetLargeWorldCoordinateTileUpdateMode() |  | NiagaraSystem.h |  |
| [TOptional](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TOptional) < float > GetMaxDeltaTime() |  | NiagaraSystem.h |  |
| void GetMaxInstanceCounts<br>(<br>int32& OutSystemInstanceMax,<br>int32& OutFXTypeInstanceMax,<br>bool bBudgetAdjusted<br>) const |  | NiagaraSystem.h |  |
| [FNiagaraMessageStore](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraMessageStore) & GetMessageStore() |  | NiagaraSystem.h |  |
| int GetNumEmitters() |  | NiagaraSystem.h |  |
| bool GetOverrideScalabilitySettings() |  | NiagaraSystem.h |  |
| [UNiagaraParameterCollectionInstance](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraParameterCollectionInsta-) \\* GetParameterCollectionOverride<br>(<br>[UNiagaraParameterCollection](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraParameterCollection)\\* Collection<br>) |  | NiagaraSystem.h |  |
| [ENiagaraPerformanceStateScopeMode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/ENiagaraPerformanceStateScopeMod-) GetPerformanceStateScopeMode() | Returns information on the availability of performance stat scopes | NiagaraSystem.h |  |
| int32 GetRandomSeed() |  | NiagaraSystem.h |  |
| TConstArrayView< [FNiagaraRendererExecutionIndex](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraRendererExecutionIndex) \> GetRendererCompletionOrder() |  | NiagaraSystem.h |  |
| TConstArrayView< int32 > GetRendererDrawOrder() |  | NiagaraSystem.h |  |
| TConstArrayView< [FNiagaraRendererExecutionIndex](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraRendererExecutionIndex) \> GetRendererPostTickOrder() |  | NiagaraSystem.h |  |
| [FNiagaraSystemScalabilityOverrides](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraSystemScalabilityOverrid-_1) & GetScalabilityOverrides() |  | NiagaraSystem.h |  |
| [FNiagaraPlatformSet](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraPlatformSet) & [GetScalabilityPlatformSet](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraSystem/GetScalabilityPlatformSet) () |  | NiagaraSystem.h |  |
| const [FNiagaraPlatformSet](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraPlatformSet) & [GetScalabilityPlatformSet](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraSystem/GetScalabilityPlatformSet) () |  | NiagaraSystem.h |  |
| const [FNiagaraSystemScalabilitySettings](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraSystemScalabilitySetting-) & GetScalabilitySettings() |  | NiagaraSystem.h |  |
| const [FNiagaraSystemStaticBuffers](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraSystemStaticBuffers) \\* GetStaticBuffers() |  | NiagaraSystem.h |  |
| [TStatId](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TStatId) GetStatID<br>(<br>bool bGameThread,<br>bool bConcurrent<br>) const |  | NiagaraSystem.h |  |
| const [FNiagaraSystemCompiledData](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraSystemCompiledData) & GetSystemCompiledData() |  | NiagaraSystem.h |  |
| const [FNiagaraDataSetAccessor](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraDataSetAccessor) < [ENiagaraExecutionState](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/ENiagaraExecutionState) \> & GetSystemExecutionStateAccessor() |  | NiagaraSystem.h |  |
| [FNiagaraSystemScalabilityOverrides](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraSystemScalabilityOverrid-_1) & GetSystemScalabilityOverrides() |  | NiagaraSystem.h |  |
| const [UNiagaraScript](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraScript) \\* [GetSystemSpawnScript](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraSystem/GetSystemSpawnScript) () |  | NiagaraSystem.h |  |
| [UNiagaraScript](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraScript) \\* [GetSystemSpawnScript](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraSystem/GetSystemSpawnScript) () | Gets the System script which is used to populate the System parameters and parameter bindings. | NiagaraSystem.h |  |
| const FNiagaraSystemStateData & GetSystemStateData() | Access the code system state data. | NiagaraSystem.h |  |
| const TCHAR \* GetSystemStateModeString() | Used for debug HUD / viewport to convey what mode we are running in | NiagaraSystem.h |  |
| const [UNiagaraScript](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraScript) \\* [GetSystemUpdateScript](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraSystem/GetSystemUpdateScript) () |  | NiagaraSystem.h |  |
| [UNiagaraScript](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraScript) \\* [GetSystemUpdateScript](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraSystem/GetSystemUpdateScript) () |  | NiagaraSystem.h |  |
| int32 GetWarmupTickCount() |  | NiagaraSystem.h |  |
| float GetWarmupTickDelta() |  | NiagaraSystem.h |  |
| float GetWarmupTime() |  | NiagaraSystem.h |  |
| void GraphSourceChanged() |  | NiagaraSystem.h |  |
| void [HandleVariableRemoved](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraSystem/HandleVariableRemoved)<br>(<br>const [FNiagaraVariable](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraVariable)& InOldVariable,<br>bool bUpdateContexts<br>) | Helper method to handle when an internal variable has been removed. | NiagaraSystem.h |  |
| void [HandleVariableRenamed](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraSystem/HandleVariableRenamed)<br>(<br>const [FNiagaraVariable](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraVariable)& InOldVariable,<br>const [FNiagaraVariable](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraVariable)& InNewVariable,<br>bool bUpdateContexts<br>) | Helper method to handle when an internal variable has been renamed. | NiagaraSystem.h |  |
| bool HasActiveCompilations() | Returns true if there are any active compilations. | NiagaraSystem.h |  |
| bool HasAnyGPUEmitters() |  | NiagaraSystem.h |  |
| bool HasDIsWithPostSimulateTick() |  | NiagaraSystem.h |  |
| bool HasFixedTickDelta() |  | NiagaraSystem.h |  |
| bool HasOutstandingCompilationRequests<br>(<br>bool bIncludingGPUShaders<br>) const | Are there any pending compile requests? | NiagaraSystem.h |  |
| void InvalidateActiveCompiles() | Invalidates any active compilation requests which will ignore their results. | NiagaraSystem.h |  |
| void InvalidateCachedData() |  | NiagaraSystem.h |  |
| bool IsAllowedByScalability() | Returns true if this emitter's platform filter allows it on this platform and quality level. | NiagaraSystem.h |  |
| bool IsInitialOwnerVelocityFromActor() |  | NiagaraSystem.h |  |
| bool IsLooping() |  | NiagaraSystem.h |  |
| bool IsReadyToRun() |  | NiagaraSystem.h |  |
| bool IsValid() | Returns true if this system is valid and can be instanced. False otherwise. | NiagaraSystem.h |  |
| void KillAllActiveCompilations() | Tries to abort all running shader compilations | NiagaraSystem.h |  |
| bool NeedsDeterminism() |  | NiagaraSystem.h |  |
| bool NeedsGPUContextInitForDataInterfaces() |  | NiagaraSystem.h |  |
| bool NeedsRequestCompile() | Do we have any pending compilation requests or not | NiagaraSystem.h |  |
| bool NeedsSortedSignificanceCull() |  | NiagaraSystem.h |  |
| bool NeedsWarmup() |  | NiagaraSystem.h |  |
| void OnCompiledDataInterfaceChanged() | Updates any post compile data based upon data interfaces. | NiagaraSystem.h |  |
| void OnCompiledUObjectChanged() | Updates the system post [UObject](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObject) change. | NiagaraSystem.h |  |
| FOnScalabilityChanged & OnScalabilityChanged() | Delegate called on effect type or effect type value change | NiagaraSystem.h |  |
| FOnSystemCompiled & OnSystemCompiled() | Delegate called when the system's dependencies have all been compiled. | NiagaraSystem.h |  |
| FOnSystemPostEditChange & OnSystemPostEditChange() | Delegate called on PostEditChange. | NiagaraSystem.h |  |
| bool PollForCompilationComplete<br>(<br>bool bFlushRequestCompile<br>) | If we have a pending compile request, is it done with yet? | NiagaraSystem.h |  |
| void PrecachePSOs() |  | NiagaraSystem.h |  |
| void PrepareRapidIterationParametersForCompilation() | Updates the rapid iteration parameters for all scripts referenced by the system. | NiagaraSystem.h |  |
| bool ReferencesInstanceEmitter<br>(<br>const [FVersionedNiagaraEmitter](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FVersionedNiagaraEmitter)& Emitter<br>) const | Determines if this system has the supplied emitter as an editable and simulating emitter instance. | NiagaraSystem.h |  |
| void RefreshSystemParametersFromEmitter<br>(<br>const [FNiagaraEmitterHandle](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraEmitterHandle)& EmitterHandle<br>) | Updates the system's rapid iteration parameters from a specific emitter. | NiagaraSystem.h |  |
| void RegisterActiveInstance() |  | NiagaraSystem.h |  |
| void RemoveEmitterHandle<br>(<br>const [FNiagaraEmitterHandle](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraEmitterHandle)& EmitterHandleToDelete<br>) | Removes the provided emitter handle. | NiagaraSystem.h |  |
| void RemoveEmitterHandlesById<br>(<br>const TSet< [FGuid](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FGuid) >& HandlesToRemove<br>) | Removes the emitter handles which have an Id in the supplied set. | NiagaraSystem.h |  |
| void RemoveSystemParametersForEmitter<br>(<br>const [FNiagaraEmitterHandle](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraEmitterHandle)& EmitterHandle<br>) | Removes the system's rapid iteration parameters for a specific emitter. | NiagaraSystem.h |  |
| void ReportAnalyticsData<br>(<br>bool bIsCooking<br>) |  | NiagaraSystem.h |  |
| bool RequestCompile<br>(<br>bool bForce,<br>[FNiagaraSystemUpdateContext](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraSystemUpdateContext)\\* OptionalUpdateContext,<br>const [ITargetPlatform](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Developer/TargetPlatform/ITargetPlatform)\\* TargetPlatform<br>) | Request that any dirty scripts referenced by this system be compiled. | NiagaraSystem.h |  |
| void ResetToEmptySystem() | Resets internal data leaving it in a state which would have minimal cost to exist in headless builds (servers) | NiagaraSystem.h |  |
| void ResolveWarmupTickCount() |  | NiagaraSystem.h |  |
| void SetBakeOutRapidIterationOnCook<br>(<br>bool bBakeOut<br>) |  | NiagaraSystem.h |  |
| void SetBakerGeneratedSettings<br>(<br>[UNiagaraBakerSettings](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraBakerSettings)\\* Settings<br>) |  | NiagaraSystem.h |  |
| void SetCompileForEdit<br>(<br>bool bNewCompileForEdit<br>) |  | NiagaraSystem.h |  |
| void SetEffectType<br>(<br>[UNiagaraEffectType](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraEffectType)\\* EffectType<br>) |  | NiagaraSystem.h |  |
| void SetFixedBounds<br>(<br>const FBox& Box<br>) |  | NiagaraSystem.h |  |
| void SetIsolateEnabled<br>(<br>bool bIsolate<br>) |  | NiagaraSystem.h |  |
| void SetOverrideScalabilitySettings<br>(<br>bool bOverride<br>) |  | NiagaraSystem.h |  |
| void SetTrimAttributesOnCook<br>(<br>bool bTrim<br>) |  | NiagaraSystem.h |  |
| void SetWarmupTickDelta<br>(<br>float InWarmupTickDelta<br>) |  | NiagaraSystem.h |  |
| void SetWarmupTime<br>(<br>float InWarmupTime<br>) |  | NiagaraSystem.h |  |
| bool ShouldCompressAttributes() |  | NiagaraSystem.h |  |
| bool ShouldDisableDebugSwitches() |  | NiagaraSystem.h |  |
| bool ShouldIgnoreParticleReadsForAttributeTrim() |  | NiagaraSystem.h |  |
| bool ShouldTrimAttributes() |  | NiagaraSystem.h |  |
| bool ShouldUseRapidIterationParameters() |  | NiagaraSystem.h |  |
| bool SupportsLargeWorldCoordinates() |  | NiagaraSystem.h |  |
| void UnregisterActiveInstance() |  | NiagaraSystem.h |  |
| void UpdateScalability() |  | NiagaraSystem.h |  |
| void UpdateSystemAfterLoad() |  | NiagaraSystem.h |  |
| bool UsesCollection<br>(<br>const [UNiagaraParameterCollection](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraParameterCollection)\\* Collection<br>) const |  | NiagaraSystem.h |  |
| bool [UsesEmitter](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraSystem/UsesEmitter)<br>(<br>[UNiagaraEmitterBase](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraEmitterBase)\\* Emitter<br>) const |  | NiagaraSystem.h |  |
| bool [UsesEmitter](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraSystem/UsesEmitter)<br>(<br>const [FVersionedNiagaraEmitterBase](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FVersionedNiagaraEmitterBase)& VersionedEmitter<br>) const |  | NiagaraSystem.h |  |
| bool [UsesEmitter](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraSystem/UsesEmitter)<br>(<br>const [FVersionedNiagaraEmitter](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FVersionedNiagaraEmitter)& VersionedEmitter<br>) const |  | NiagaraSystem.h |  |
| bool UsesScript<br>(<br>const [UNiagaraScript](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraScript)\\* Script<br>) const |  | NiagaraSystem.h |  |
| void WaitForCompilationComplete<br>(<br>bool bIncludingGPUShaders,<br>bool bShowProgress<br>) | Blocks until all active compile jobs have finished | NiagaraSystem.h |  |
| void WaitForCompilationComplete\_SkipPendingOnDemand<br>(<br>bool bIncludingGPUShaders,<br>bool bShowProgress<br>) | Blocks until all active compile jobs have finished but does not flush any requests that were deferred from loading. | NiagaraSystem.h |  |

#### Overridden from [UObject](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObject)

| Name | Remarks | Include Path | [Unreal Specifiers](https://dev.epicgames.com/documentation/unreal-engine/ufunctions-in-unreal-engine#functionspecifiers) |
| --- | --- | --- | --- |
| virtual void BeginCacheForCookedPlatformData<br>(<br>const [ITargetPlatform](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Developer/TargetPlatform/ITargetPlatform)\\* TargetPlatform<br>) |  | NiagaraSystem.h |  |
| virtual void BeginDestroy() |  | NiagaraSystem.h |  |
| virtual void GetAssetRegistryTagMetadata<br>(<br>[TMap](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TMap) < [FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName), FAssetRegistryTagMetadata >& OutMetadata<br>) const |  | NiagaraSystem.h |  |
| virtual void [GetAssetRegistryTags](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraSystem/GetAssetRegistryTags)<br>(<br>[TArray](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArray) < FAssetRegistryTag >& OutTags<br>) const |  | NiagaraSystem.h |  |
| virtual void [GetAssetRegistryTags](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraSystem/GetAssetRegistryTags)<br>(<br>[FAssetRegistryTagsContext](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/FAssetRegistryTagsContext) Context<br>) const |  | NiagaraSystem.h |  |
| virtual void GetResourceSizeEx<br>(<br>[FResourceSizeEx](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FResourceSizeEx)& CumulativeResourceSize<br>) |  | NiagaraSystem.h |  |
| virtual bool IsCachedCookedPlatformDataLoaded<br>(<br>const [ITargetPlatform](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Developer/TargetPlatform/ITargetPlatform)\\* TargetPlatform<br>) |  | NiagaraSystem.h |  |
| virtual bool IsReadyForFinishDestroy() |  | NiagaraSystem.h |  |
| virtual void OnCookEvent<br>(<br>UE::Cook::ECookEvent CookEvent,<br>[UE::Cook::FCookEventContext](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/FCookEventContext)& CookContext<br>) |  | NiagaraSystem.h |  |
| virtual void PostDuplicate<br>(<br>bool bDuplicateForPIE<br>) |  | NiagaraSystem.h |  |
| virtual void PostEditChangeProperty<br>(<br>[FPropertyChangedEvent](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/FPropertyChangedEvent)& PropertyChangedEvent<br>) |  | NiagaraSystem.h |  |
| virtual void PostInitProperties() |  | NiagaraSystem.h |  |
| virtual void PostLoad() |  | NiagaraSystem.h |  |
| virtual void PostRename<br>(<br>[UObject](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObject)\\* OldOuter,<br>const [FName](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FName) OldName<br>) |  | NiagaraSystem.h |  |
| virtual void PreEditChange<br>(<br>[FProperty](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/FProperty)\\* PropertyThatWillChange<br>) |  | NiagaraSystem.h |  |
| virtual void PreSave<br>(<br>[FObjectPreSaveContext](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/FObjectPreSaveContext) ObjectSaveContext<br>) |  | NiagaraSystem.h |  |
| virtual void Serialize<br>(<br>[FArchive](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FArchive)& Ar<br>) |  | NiagaraSystem.h |  |

#### Overridden from [UObjectBaseUtility](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObjectBaseUtility)

| Name | Remarks | Include Path | [Unreal Specifiers](https://dev.epicgames.com/documentation/unreal-engine/ufunctions-in-unreal-engine#functionspecifiers) |
| --- | --- | --- | --- |
| virtual bool CanBeClusterRoot() |  | NiagaraSystem.h |  |

#### Overridden from [INiagaraParameterDefinitionsSubscriber](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/INiagaraParameterDefinitionsSubs-)

| Name | Remarks | Include Path | [Unreal Specifiers](https://dev.epicgames.com/documentation/unreal-engine/ufunctions-in-unreal-engine#functionspecifiers) |
| --- | --- | --- | --- |
| virtual [TArray](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArray) < [UNiagaraScriptSourceBase](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraScriptSourceBase) \\* \> GetAllSourceScripts() | Get all [UNiagaraScriptSourceBase](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraScriptSourceBase) of this subscriber. | NiagaraSystem.h |  |
| virtual [TArray](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArray) < [UNiagaraEditorParametersAdapterBase](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraEditorParametersAdapterB-) \\* \> GetEditorOnlyParametersAdapters() | Get All adapters to editor only script vars owned directly by this subscriber. | NiagaraSystem.h |  |
| virtual [TArray](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArray) < [INiagaraParameterDefinitionsSubscriber](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/INiagaraParameterDefinitionsSubs-) \\* \> [GetOwnedParameterDefinitionsSubscribers](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraSystem/GetOwnedParameterDefinitionsSubs-) () | Get all subscribers that are owned by this subscriber. | NiagaraSystem.h |  |
| virtual [TArray](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArray) < [FParameterDefinitionsSubscription](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FParameterDefinitionsSubscriptio-) \> & [GetParameterDefinitionsSubscriptions](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraSystem/GetParameterDefinitionsSubscript-) () |  | NiagaraSystem.h |  |
| virtual const [TArray](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArray) < [FParameterDefinitionsSubscription](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FParameterDefinitionsSubscriptio-) \> & [GetParameterDefinitionsSubscriptions](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraSystem/GetParameterDefinitionsSubscript-) () |  | NiagaraSystem.h |  |
| virtual [FString](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FString) GetSourceObjectPathName() | Get the path to the [UObject](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObject) of this subscriber. | NiagaraSystem.h |  |

### Protected

| Name | Remarks | Include Path | [Unreal Specifiers](https://dev.epicgames.com/documentation/unreal-engine/ufunctions-in-unreal-engine#functionspecifiers) |
| --- | --- | --- | --- |
| void GenerateStatID() |  | NiagaraSystem.h |  |
| void UpdateStatID() |  | NiagaraSystem.h |  |

### Static

| Name | Remarks | Include Path | [Unreal Specifiers](https://dev.epicgames.com/documentation/unreal-engine/ufunctions-in-unreal-engine#functionspecifiers) |
| --- | --- | --- | --- |
| static void [AppendToClassSchema](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraSystem/AppendToClassSchema)<br>(<br>[FAppendToClassSchemaContext](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/FAppendToClassSchemaContext)& Context<br>) | Append config values or settings that can change how instances of the class are cooked, including especially values that determine how version upgraded are conducted. | NiagaraSystem.h |  |
| static void DeclareConstructClasses<br>(<br>[TArray](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArray) < [FTopLevelAssetPath](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/FTopLevelAssetPath) >& OutConstructClasses,<br>const [UClass](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UClass)\\* SpecificSubclass<br>) |  | NiagaraSystem.h |  |
| static void RecomputeExecutionOrderForDataInterface<br>(<br>[UNiagaraDataInterface](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraDataInterface)\\* DataInterface<br>) |  | NiagaraSystem.h |  |
| static void RecomputeExecutionOrderForEmitter<br>(<br>const [FVersionedNiagaraEmitter](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FVersionedNiagaraEmitter)& InEmitter<br>) |  | NiagaraSystem.h |  |
| static void RequestCompileForEmitter<br>(<br>const [FVersionedNiagaraEmitter](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FVersionedNiagaraEmitter)& InEmitter<br>) |  | NiagaraSystem.h |  |

## Deprecated Variables

| Name | Type | Remarks | Include Path | [Unreal Specifiers](https://dev.epicgames.com/documentation/unreal-engine/unreal-engine-uproperties#propertyspecifiers) |
| --- | --- | --- | --- | --- |
| AssetTags\_DEPRECATED | [TArray](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArray) < [FNiagaraAssetTagDefinitionReference](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraAssetTagDefinitionRefere-) > |  | NiagaraSystem.h | - Meta=(DeprecatedProperty) |
| bExposeToLibrary\_DEPRECATED | bool | Deprecated library exposure bool. | NiagaraSystem.h |  |
| bIsTemplateAsset\_DEPRECATED | bool | Deprecated template asset bool. Use the TemplateSpecification enum instead. | NiagaraSystem.h |  |
| MessageKeyToMessageMap\_DEPRECATED | [TMap](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TMap) < [FGuid](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FGuid), [TObjectPtr](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TObjectPtr) < [UNiagaraMessageDataBase](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/UNiagaraMessageDataBase) \> > | Messages associated with the System asset. | NiagaraSystem.h |  |
| ScalabilityOverrides\_DEPRECATED | [TArray](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArray) < [FNiagaraSystemScalabilityOverride](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/FNiagaraSystemScalabilityOverrid-) > |  | NiagaraSystem.h |  |
| TemplateSpecification\_DEPRECATED | [ENiagaraScriptTemplateSpecification](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/Niagara/ENiagaraScriptTemplateSpecificat-) | If this system is a regular system, a template or a behavior example. | NiagaraSystem.h |  |

* * *