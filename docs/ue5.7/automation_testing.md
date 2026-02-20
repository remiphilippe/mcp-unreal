# Automation Testing in UE 5.7

**Module**: AutomationTest (Core)

UE's automation framework provides unit, functional, and integration tests that run headless (-nullrhi) or in-editor. Tests are registered via macros and discovered at runtime. The mcp-unreal server uses this framework via the run_tests and list_tests tools.

## Test Types

- **Simple Tests** — Single function, runs once. Use `IMPLEMENT_SIMPLE_AUTOMATION_TEST`.
- **Complex Tests** — Parameterized, runs multiple times with different parameters. Use `IMPLEMENT_COMPLEX_AUTOMATION_TEST`.
- **Latent Tests** — Async tests that span multiple frames. Use `ADD_LATENT_AUTOMATION_COMMAND`.
- **Functional Tests** — Level-based tests using AFunctionalTest actors placed in test maps.

## Writing a Simple Test

```cpp
IMPLEMENT_SIMPLE_AUTOMATION_TEST(FMyTest, "Project.Category.TestName",
    EAutomationTestFlags::EditorContext | EAutomationTestFlags::ProductContext |
    EAutomationTestFlags::HighPriority)

bool FMyTest::RunTest(const FString& Parameters)
{
    // Test code
    TestEqual(TEXT("Value should be 42"), MyValue, 42);
    TestTrue(TEXT("Should be valid"), bIsValid);
    return true;
}
```

## Running Tests Headless

```bash
UnrealEditor-Cmd /path/to/Project.uproject \
    -ExecCmds="Automation RunTests Project.Category" \
    -nullrhi -unattended -nopause -nosplash -nosound \
    -log=TestLog.txt
```

## Key Assertion Functions

- `TestEqual(Description, Actual, Expected)` — Asserts equality
- `TestTrue(Description, Value)` — Asserts true
- `TestFalse(Description, Value)` — Asserts false
- `TestNull(Description, Pointer)` — Asserts null
- `TestNotNull(Description, Pointer)` — Asserts not null
- `TestValid(Description, SharedPtr)` — Asserts valid shared pointer
- `AddError(Message)` — Records a test error
- `AddWarning(Message)` — Records a test warning
- `AddInfo(Message)` — Records test info

## Log Output Format

Test results appear in the UE log with these patterns:
- `LogAutomationTest: Test Completed. Result={Success|Fail} Test={TestName}`
- `LogAutomationTest: BeginEvents: ...`
- `LogAutomationTest: Error: ...` (failure details)
- `LogAutomationController: ... tests completed in ... seconds`
