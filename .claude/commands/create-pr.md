# Open Pull Request Command

Create and open a comprehensive, well-structured pull request from the current branch with intelligent analysis of changes and proper categorization.

## Usage Pattern

When invoked, this command will:
1. **Analyze all changes** on the current branch compared to the base branch
2. **Generate a comprehensive PR description** with proper formatting
3. **Apply appropriate labels** based on change analysis
4. **Open the PR** with all metadata configured

## Change Analysis Process

1. **Scope Detection**: Analyze modified files to determine:
    - Type of changes (feature, fix, refactor, docs, etc.)
    - Components/modules affected
    - Breaking changes or backwards compatibility concerns
    - Required reviewers based on code ownership

2. **Commit Analysis**: Review all commits to:
    - Extract key changes from commit messages
    - Group related changes together
    - Identify the primary purpose of the PR

3. **Impact Assessment**: Evaluate:
    - Test coverage for changed code
    - Documentation updates needed or included
    - Potential risks or side effects
    - Performance implications

## PR Description Template

Generate PR descriptions with the following structure:

```markdown
## Summary
[Brief 1-2 sentence overview of what this PR accomplishes]

## Motivation
[Why these changes are needed - reference issue numbers if applicable]

## Changes Made

### Features/Fixes
- [Bullet list of key changes, grouped logically]
- [Use checkboxes for multi-part features: - [x] Completed part]

### Technical Details
[Explanation of implementation approach, architectural decisions, or complex logic]

## Testing
- [ ] Unit tests added/updated
- [ ] Integration tests added/updated
- [ ] Manual testing completed
- [ ] Performance impact assessed

### Test Coverage
[Report on test coverage for modified files]

## Documentation
- [ ] API documentation updated
- [ ] README updated if needed
- [ ] Inline comments added for complex logic
- [ ] CHANGELOG entry added

## Screenshots/Examples
[If UI changes or visible output changes, include before/after]

## Breaking Changes
[List any breaking changes with migration instructions]

## Checklist
- [ ] Code follows project style guidelines
- [ ] Self-review completed
- [ ] Tests pass locally
- [ ] No console errors or warnings
- [ ] Accessibility considerations addressed (if applicable)

## Related Issues
Closes #[issue number]
Related to #[issue number]

## Additional Notes
[Any additional context, deployment notes, or considerations for reviewers]
```

## Labeling Strategy

Automatically apply labels based on analysis:

### Type Labels (Primary - choose one):
- `feat`: New features or functionality
- `fix`: Bug fixes
- `docs`: Documentation only changes
- `refactor`: Code refactoring without behavior changes
- `test`: Test additions or modifications
- `perf`: Performance improvements
- `style`: Code style/formatting changes
- `build`: Build system or dependency changes
- `ci`: CI/CD configuration changes

### Scope Labels (Multiple allowed):
- Component-specific labels based on affected files
- `breaking-change`: For backwards-incompatible changes
- `security`: Security-related changes
- `accessibility`: Accessibility improvements
- `dependencies`: Dependency updates

### Status Labels:
- `ready-for-review`: PR is complete and ready
- `work-in-progress`: Still being developed
- `needs-tests`: Tests need to be added
- `needs-docs`: Documentation needed

### Size Labels:
- `size/XS`: < 10 lines changed
- `size/S`: 10-100 lines changed
- `size/M`: 100-500 lines changed
- `size/L`: 500-1000 lines changed
- `size/XL`: > 1000 lines changed

## Implementation Steps

1. **Gather Information**:
   ```bash
   # Get current branch name
   current_branch=$(git branch --show-current)
   
   # Get base branch (usually main or master)
   base_branch=$(git symbolic-ref refs/remotes/origin/HEAD | sed 's@^refs/remotes/origin/@@')
   
   # Get list of changed files
   git diff --name-only $base_branch...$current_branch
   
   # Get commit messages
   git log --oneline $base_branch...$current_branch
   ```

2. **Analyze Changes**:
    - Parse changed files by type and component
    - Extract conventional commit types from commit messages
    - Check for test file changes
    - Identify documentation updates

3. **Generate PR Body**:
    - Use the template above
    - Fill in sections based on analysis
    - Include relevant issue references
    - Add code examples if helpful

4. **Create Pull Request**:
   ```bash
   # Open PR with generated content
   gh pr create \
     --title "[Type](scope): Brief description" \
     --body "$(cat generated_pr_body.md)" \
     --base $base_branch \
     --head $current_branch \
     --label "type,scope,size" \
     --assignee "@me" \
     --reviewer "team-lead,code-owner"
   ```

## Title Format Guidelines

Generate PR titles following conventional format:
- `feat(component): add new feature description`
- `fix(api): resolve issue with endpoint validation`
- `docs: update README with new installation steps`
- `refactor(auth): simplify token validation logic`
- `test(user): add integration tests for profile updates`

Keep titles:
- Under 72 characters
- Descriptive but concise
- Starting with lowercase (except for proper nouns)
- Without periods at the end

## Review Assignment Logic

Automatically suggest reviewers based on:
1. **Code ownership**: Check CODEOWNERS file
2. **File patterns**: Assign domain experts for specific areas
3. **Recent contributors**: People who recently modified affected files
4. **Team structure**: Respect team review requirements

## Draft vs Ready PR Logic

Create as draft if:
- Tests are failing
- TODO comments exist in changed files
- Documentation is incomplete
- Commit messages contain "WIP" or "draft"
- Large number of changes (>1000 lines) needing careful review

Otherwise, create as ready for review.

## Example Workflow

```bash
# Analyze current branch changes
git diff --stat main...$(git branch --show-current)

# Generate PR description based on changes
# [Automated analysis and generation happens here]

# Open PR with all metadata
gh pr create \
  --title "feat(http-gen): add support for multipart form data" \
  --body "$(cat <<'EOF'
## Summary
Adds comprehensive support for multipart/form-data in HTTP request generation.

## Motivation
Closes #234 - Users need to upload files through generated HTTP clients.

## Changes Made
### Features
- Added MultipartFormData builder class
- Implemented file upload handling
- Added streaming support for large files

### Technical Details
Uses streams to handle large file uploads efficiently without loading entire files into memory.

## Testing
- [x] Unit tests added
- [x] Integration tests added
- [x] Manual testing completed

## Documentation
- [x] API documentation updated
- [x] README examples added

## Related Issues
Closes #234
EOF
)" \
  --label "feat,http-gen,size/L" \
  --assignee "@me"
```

## Quality Checks

Before opening the PR, verify:
- **Build passes**: Run build/compile commands
- **Tests pass**: Execute test suite
- **Linting clean**: No style violations
- **Documentation complete**: All public APIs documented
- **Commits organized**: Following commit guidelines from your smart commit strategy

## Integration with Other Commands

This command integrates with:
- **Smart Git Commit**: Ensures commits are well-organized before PR
- **Documentation Update**: Verifies docs are current
- **Issue Creation**: Links to relevant issues

## Benefits

- **Consistent PR quality**: All PRs follow the same high standard
- **Faster reviews**: Reviewers have all needed context
- **Better tracking**: Proper labels and links for project management
- **Reduced back-and-forth**: Comprehensive information upfront
- **Automated best practices**: Enforces team standards automatically