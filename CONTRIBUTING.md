# Contributing `GoDo`

Thank you for contributing `GoDo` :tada:

## Found a Bug?

If you find a bug in the source code, you can help us by [submitting an issue]
to our [GitHub Repository]. Even better, you can submit a Pull Request with a
fix.

## Commit Message Format

Commit messages should follow the [conventional commit style].

```text
type(scope): subject
BLANK LINE
body
BLANK LINE
footer
```

### Type

Must be one of the following:

- **ci:** Changes to our CI configuration files and scripts
- **chore:** Updating documentation etc, no production code changes
- **feat:** A new feature
- **fix:** A bug fix
- **perf:** A code change that improves performance
- **refactor:** A code change that neither fixes a bug nor adds a feature
- **style:** Changes that do not affect the meaning of the code
- **test:** Adding missing tests or correcting existing tests

### Footer

If appropriate the footer should contain a [closing reference to an issue].

The **footer** can contain any information about **Breaking Changes** and is
also the place to reference GitHub issues that this commit **Closes**.

**Breaking Changes** should start with the word `BREAKING CHANGE:` with a space
or two newlines. The rest of the commit message is then used for this.

[submitting an issue]: https://github.com/rsHalford/godo/issues
[GitHub Repository]: https://github.com/rsHalford/godo
[conventional commit style]: https://www.conventionalcommits.org/en/v1.0.0/
[closing reference to an issue]: https://help.github.com/articles/closing-issues-via-commit-messages/
