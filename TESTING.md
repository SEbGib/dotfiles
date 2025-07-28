# Testing Guide

This document describes the comprehensive testing strategy for the Dotfiles TUI project.

## 🧪 Test Structure

### Test Types

1. **Unit Tests** - Test individual functions and methods
2. **Integration Tests** - Test component interactions
3. **Benchmark Tests** - Performance testing
4. **End-to-End Tests** - Full workflow testing

### Test Organization

```
internal/
├── tui/
│   ├── main_test.go          # Main menu tests
│   ├── editor_test.go        # Editor functionality tests
│   └── integration_test.go   # Cross-component tests
├── scripts/
│   └── integration_test.go   # Script runner tests
└── testutil/
    └── helpers.go            # Test utilities and helpers
```

## 🚀 Running Tests

### Quick Test Commands

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Run benchmarks
make bench

# Run comprehensive test suite
./run-tests.sh
```

### Detailed Test Commands

```bash
# Unit tests with race detection
go test -v -race ./...

# Tests with coverage report
go test -v -race -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html

# Specific package tests
go test -v ./internal/tui
go test -v ./internal/scripts

# Run specific test
go test -v ./internal/tui -run TestNewMainModel

# Benchmarks with memory stats
go test -bench=. -benchmem ./...

# Integration tests only
go test -v -tags=integration ./...
```

## 📊 Test Coverage

### Coverage Goals

- **Minimum Coverage**: 70%
- **Target Coverage**: 85%
- **Critical Components**: 90%+

### Coverage Reports

```bash
# Generate coverage report
go test -coverprofile=coverage.out ./...

# View coverage in terminal
go tool cover -func=coverage.out

# Generate HTML report
go tool cover -html=coverage.out -o coverage.html
```

## 🔧 Test Utilities

### ModelTester

Helper for testing Bubble Tea models:

```go
tester := testutil.NewModelTester(t)
model := NewMainModel()

// Test view content
tester.AssertViewContains(model, "expected text")

// Send key events
updatedModel := tester.SendKey(model, tea.KeyDown)
```

### TempFileHelper

Helper for creating temporary files in tests:

```go
fileHelper := testutil.NewTempFileHelper(t)
testFile := fileHelper.CreateFile("test.conf", "content")
```

### MockEnvironment

Helper for mocking environment variables:

```go
mockEnv := testutil.NewMockEnvironment(t)
defer mockEnv.Restore()
mockEnv.SetEnv("HOME", "/tmp/test")
```

## 🧪 Test Categories

### Unit Tests

Test individual components in isolation:

- **Model Creation**: Test model initialization
- **State Management**: Test state transitions
- **View Rendering**: Test UI output
- **Key Handling**: Test user input processing

### Integration Tests

Test component interactions:

- **Model Transitions**: Test navigation between screens
- **File Operations**: Test editor with real files
- **Script Integration**: Test shell script execution
- **Error Handling**: Test error scenarios

### Benchmark Tests

Performance testing:

- **View Rendering**: Measure rendering performance
- **Model Updates**: Measure update performance
- **Memory Usage**: Track memory allocations

## 🔍 Testing Best Practices

### Test Structure

```go
func TestFeatureName(t *testing.T) {
    // Arrange
    model := NewTestModel()
    
    // Act
    result := model.DoSomething()
    
    // Assert
    if result != expected {
        t.Errorf("Expected %v, got %v", expected, result)
    }
}
```

### Test Naming

- Use descriptive test names
- Follow pattern: `Test<Component><Behavior>`
- Example: `TestMainModelNavigation`

### Test Data

- Use table-driven tests for multiple scenarios
- Create test fixtures for complex data
- Use temporary files for file operations

### Error Testing

```go
func TestErrorHandling(t *testing.T) {
    // Test error conditions
    _, err := functionThatCanFail()
    if err == nil {
        t.Error("Expected error but got none")
    }
}
```

## 🚦 Continuous Integration

### GitHub Actions

The project uses GitHub Actions for automated testing:

- **Multiple OS**: Ubuntu and macOS
- **Multiple Go versions**: 1.21, 1.22
- **Test Coverage**: Automatic coverage reporting
- **Security Scanning**: Vulnerability detection
- **Linting**: Code quality checks

### Local CI Simulation

```bash
# Run the same tests as CI
./run-tests.sh

# Check linting
golangci-lint run

# Security scan
gosec ./...

# Vulnerability check
govulncheck ./...
```

## 📈 Test Metrics

### Key Metrics

- **Test Coverage**: Percentage of code covered by tests
- **Test Duration**: Time taken to run all tests
- **Benchmark Results**: Performance measurements
- **Failure Rate**: Percentage of failing tests

### Monitoring

```bash
# Coverage threshold check
COVERAGE=$(go tool cover -func=coverage.out | grep total | awk '{print $3}' | sed 's/%//')
if (( $(echo "$COVERAGE < 70" | bc -l) )); then
    echo "Coverage below threshold"
    exit 1
fi
```

## 🐛 Debugging Tests

### Common Issues

1. **Race Conditions**: Use `-race` flag to detect
2. **Flaky Tests**: Add timeouts and retries
3. **Environment Dependencies**: Mock external dependencies
4. **File System Issues**: Use temporary directories

### Debug Commands

```bash
# Run with verbose output
go test -v ./...

# Run specific test with debug info
go test -v ./internal/tui -run TestSpecificTest

# Run with race detection
go test -race ./...

# Run with CPU profiling
go test -cpuprofile=cpu.prof ./...
```

## 📚 Test Documentation

### Writing Test Documentation

- Document complex test scenarios
- Explain test data and fixtures
- Document performance expectations
- Include troubleshooting guides

### Test Comments

```go
// TestComplexScenario tests the interaction between multiple components
// when handling edge cases like network failures and file system errors.
// This test uses mocked dependencies to ensure consistent behavior.
func TestComplexScenario(t *testing.T) {
    // Test implementation
}
```

## 🎯 Testing Checklist

Before submitting code:

- [ ] All tests pass locally
- [ ] Coverage meets minimum threshold
- [ ] New features have tests
- [ ] Edge cases are covered
- [ ] Performance tests added for critical paths
- [ ] Integration tests updated
- [ ] Documentation updated

## 🔄 Test Maintenance

### Regular Tasks

- Review and update test coverage
- Remove obsolete tests
- Update test data and fixtures
- Optimize slow tests
- Update CI configuration

### Test Refactoring

- Extract common test utilities
- Reduce test duplication
- Improve test readability
- Update test documentation