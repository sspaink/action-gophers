# action-gophers

A Github action written in Go that posts a Go Gopher gif on your PRs!

## Usage

```yaml
name: Example
on:
  pull_request:
    branches: [ master ]

jobs:
  my_job:
    runs-on: ubuntu-latest

    steps:
    - name: Comment with gopher
      uses: BattleBas/action-gophers@master
      with:
        GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```