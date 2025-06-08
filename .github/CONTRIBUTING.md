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
   ./scripts/run_tests.sh
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

## CI/CD Pipeline

The project uses a unified GitHub Actions workflow:

### For Pull Requests
- **Tests only**: Runs on Go 1.23.0
- **Coverage**: Reports uploaded to Codecov
- **Must pass**: Before merging

### For Main Branch Pushes
- **Full pipeline**: Test → Tag → Release
- **Automatic versioning**: Based on commit messages
- **GoReleaser**: Builds cross-platform binaries
- **GitHub Release**: Created automatically

### Pipeline Flow
```
1. Tests run on Go 1.23.0
2. If tests pass and on main branch:
   - Analyze commits for version bump
   - Create semantic version tag
   - Run GoReleaser for release
```

## Release Process

Releases are fully automated:

- **Push to main** with conventional commits
- **Pipeline determines** version bump automatically
- **Tag created** and pushed to repository
- **GoReleaser builds** cross-platform binaries
- **GitHub release** created with assets

### Example Release Flow
```bash
# Your changes
git commit -m "feat: add interactive search"
git push origin main

# Automated result:
# 1. Tests pass ✅
# 2. Version bumped to v1.1.0 (feat = minor)
# 3. Tag v1.1.0 created and pushed
# 4. GoReleaser builds binaries
# 5. GitHub release created
```

## Questions?

Feel free to open an issue for any questions or discussions!
