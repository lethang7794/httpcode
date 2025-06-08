# Contributing to HTTP Code CLI Tool

Thank you for your interest in contributing to the HTTP Code CLI Tool! 🎉

## Development Setup

1. **Fork and clone the repository**
   ```bash
   git clone https://github.com/lethang7794/httpcode.git
   cd httpcode
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Run tests**
   ```bash
   go test ./...
   # or use the test script
   ./run_tests.sh
   ```

4. **Build and test locally**
   ```bash
   go build -o httpcode .
   ./httpcode --help
   ```

## Commit Message Convention

We use conventional commits for automatic semantic versioning:

### Commit Types

- **feat:** New feature (minor version bump)
- **fix:** Bug fix (patch version bump)  
- **breaking:** Breaking change (major version bump)
- **docs:** Documentation changes (patch version bump)
- **style:** Code style changes (patch version bump)
- **refactor:** Code refactoring (patch version bump)
- **test:** Adding or updating tests (patch version bump)
- **chore:** Maintenance tasks (patch version bump)

### Examples

```bash
feat: add fuzzy search functionality
fix: resolve status code lookup issue
breaking: change API structure
docs: update README with new examples
test: add tests for list command
```

## Pull Request Process

1. **Create a feature branch**
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. **Make your changes**
   - Write tests for new functionality
   - Update documentation if needed
   - Follow Go best practices

3. **Test your changes**
   ```bash
   go test ./...
   go build .
   ```

4. **Commit with conventional messages**
   ```bash
   git commit -m "feat: add new search functionality"
   ```

5. **Push and create PR**
   ```bash
   git push origin feature/your-feature-name
   ```

## Code Style

- Follow standard Go formatting (`go fmt`)
- Use meaningful variable and function names
- Add comments for exported functions
- Keep functions focused and small

## Testing

- Write tests for new functionality
- Maintain or improve test coverage
- Test on multiple Go versions (1.19, 1.20, 1.21)

## Release Process

Releases are automated based on commit messages:

- Push to `main` branch triggers the release workflow
- Version is automatically determined from commit messages
- GitHub release is created with binaries for multiple platforms

## Questions?

Feel free to open an issue for any questions or discussions!
