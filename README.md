# Scaffold Example Registry

This repository serves as a template registry for the [Scaffold CLI tool](https://github.com/kjuulh/scaffold). It provides a foundation for creating and managing your own collection of templates.

## Getting Started

### Setup

1. Fork this repository to your GitHub account
2. Clone your forked repository:
   ```bash
   git clone git@github.com:yourusername/scaffold-registry.git
   cd scaffold-registry
   ```
3. Set the environment variable to point to your registry:
   ```bash
   # Add to your .bashrc, .zshrc, or equivalent
   export SCAFFOLD_REGISTRY=git@github.com:yourusername/scaffold-registry.git
   ```

### Creating Templates

Generate a new template skeleton:
```bash
./scaffold scaffold --name my-new-template
```

This will create a new template in the `registry/my-new-template` directory with the following structure:
- `scaffold.yaml`: Configuration file for your template
- `scaffold_test.go`: Test file for verifying template functionality
- `files/`: Directory containing template files with `.gotmpl` extension
- `testdata/`: Directory for test output

### Template Development

Modify the generated files to build your template:
1. Edit `scaffold.yaml` to define inputs, default paths, and other template behavior
2. Add or modify template files in the `files/` directory using Go template syntax
3. Update tests in `scaffold_test.go` to validate your template's functionality

### Testing Templates

Run tests to ensure your templates work correctly:
```bash
go test ./...
```

For individual template testing:
```bash
cd registry/my-new-template
go test
```

The testing system is snapshot-based:
- The test output will be stored in `testdata/your_test_name/actual`
- Expected output should be in `testdata/your_test_name/expected`
- To accept changes: remove the `expected` directory and rename `actual` to `expected`

## Usage with Scaffold

Once your registry is set up, you can use your templates with the Scaffold CLI:

```bash
scaffold
# Or directly specify a template
scaffold my-new-template
```

## Best Practices

1. **Keep templates focused**: Each template should serve a specific purpose
2. **Document inputs**: Include clear descriptions in `scaffold.yaml` for all required inputs
3. **Test thoroughly**: Create tests for different input combinations
4. **Provide sensible defaults**: Make templates easy to use with minimal required input
5. **Version control**: Tag releases of your registry for stable template versions
