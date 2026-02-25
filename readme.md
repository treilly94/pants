# TODO:
- envs
- promote image

# Build dependants
```
pants --changed-since=main --changed-dependents=transitive list
```
[link](https://www.pantsbuild.org/stable/docs/using-pants/advanced-target-selection#running-over-changed-files-with---changed-since)

or

```
pants dependents pkg/package1 | xargs pants list 
```
[link](https://www.pantsbuild.org/stable/docs/using-pants/advanced-target-selection#piping-to-other-pants-runs)