using UnrealBuildTool;

public class MCPTestProjectEditorTarget : TargetRules
{
	public MCPTestProjectEditorTarget(TargetInfo Target) : base(Target)
	{
		Type = TargetType.Editor;
		DefaultBuildSettings = BuildSettingsVersion.V6;
		IncludeOrderVersion = EngineIncludeOrderVersion.Latest;
		ExtraModuleNames.Add("MCPTestProject");
	}
}
