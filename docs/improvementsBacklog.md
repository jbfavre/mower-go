# Improvement Action Plan — mower-go

Backlog of improvement spotted by an analysis of the repo by Claude Sonnet 4.6 model

## **1. Overall Improvements**

| Focus area | Priority | Status | Recommendation | Evidence |
|---|---|---:|---|---|
| Code base and execution successful | High | Done | Buildability is now fixed with Go modules and a resolvable import path. Keep this as the baseline for any candidate exercise. | go.mod, Main.go |
| Code base and execution successful | High | Open | Make execution spec-compliant. The current program still runs hardcoded moves and only prints raw file lines instead of parsing and simulating the input format. | Main.go, specification.md |
| Code quality | High | Open | Fix the rectangle coordinate logic. `IsInPosition` currently compares `x` with `Height` and `y` with `Weight`, which is semantically wrong and can hide movement bugs. | Rectangle.go |
| Code quality | Medium | Open | Replace the current orientation output with a proper string representation. Printing raw enum integers makes the result unusable for final output validation. | Mower.go |
| Code quality | Medium | Open | Remove demo/debug behavior from the executable path. `OK`, repeated `Print()`, and hardcoded `Advance()` calls obscure actual program behavior. | Main.go |
| Programming patterns | High | Open | Introduce typed commands and an execution pipeline instead of imperative calls in `main`. `L`, `R`, and `F` should be modeled explicitly. | Main.go, specification.md |
| Programming patterns | Medium | Open | Remove identifier shadowing and rely on explicit naming when refactoring starts. `var mower = mower.NewMower(...)` is legal but brittle and harder to read. | Main.go |
| Programming patterns | Medium | Open | Replace positional struct construction with named fields where possible to avoid silent breakage if field order changes. | Mower.go |
| Code and repository architecture | High | Open | Move from a demo layout to an application layout: parser, domain, simulator, output. `main` should only wire components together. | Main.go |
| Code and repository architecture | Medium | Open | Revisit package layout once behavior exists. `src/mower` works now, but it is still a legacy-style structure under modules. | mower, go.mod |
| Separation of concerns | High | Open | Create a parser responsible only for input decoding and validation, separate from movement rules and output formatting. | Main.go, specification.md |
| Separation of concerns | High | Open | Add a simulator/service layer responsible for boundaries, occupancy, and execution ordering. Those rules do not belong in `main`. | Area.go, specification.md |
| Separation of concerns | Medium | Open | Remove I/O from domain objects. `Print()` mixes domain state with console output and makes tests less direct. | Mower.go |

## **2. Code Changes**

| Focus area | Priority | Status | Recommendation | Evidence |
|---|---|---:|---|---|
| Resources usage | Medium | Open | Define scanner limits explicitly or document them. The current scanner uses default token limits and reads input without any validation strategy. | Main.go |
| Resources usage | Medium | Open | Make input handling deterministic and fail-fast. Right now the file is opened and scanned, but malformed content is neither rejected nor surfaced. | Main.go, file.txt |
| Resources usage | Low | Open | Consider avoiding direct `log.Fatal` in the long term so command execution and error reporting stay testable. | Main.go |
| Performances | High | Open | Implement concurrent mower processing only with deterministic conflict resolution. The spec requires multi-CPU execution, but correctness must stay stable. | specification.md |
| Performances | High | Open | Use step-based simulation for all mowers rather than letting each mower consume its entire command stream independently. That is the cleanest way to enforce collision rules under concurrency. | specification.md |
| Performances | Medium | Open | Introduce an occupancy structure for coordinates and a clear concurrency strategy before parallel execution. The current `Area` abstraction does not model mower presence at all. | Area.go, specification.md |
| Performances | Medium | Open | Add benchmarks only after the parser and simulator exist. Performance work is premature until the actual execution path is implemented. | mower |
| Testing | High | Open | Replace the empty rectangle test with table-driven boundary tests covering negative coordinates, edges, corners, and the current axis bug. | Rectangle_test.go, Rectangle.go |
| Testing | High | Open | Add unit tests for mower rotations and forward movement, especially boundary-blocked movement. | Mower.go |
| Testing | High | Open | Add parser tests for invalid coordinates, invalid commands, incomplete mower definitions, and empty instruction lines. | specification.md |
| Testing | High | Open | Add simulator tests for out-of-bounds discard, collision discard, and ordered final output. These are core business rules. | specification.md |
| Testing | Medium | Open | Add an end-to-end test for the official example and use it as the baseline acceptance test. | specification.md |
| Testing | Medium | Open | Add race-detector test runs once concurrent simulation exists. | specification.md |

## **3. Specifications**

| Focus area | Priority | Status | Recommendation | Evidence |
|---|---|---:|---|---|
| Loop holes | High | Open | Clarify how simultaneous move conflicts are resolved when two mowers target the same empty cell in the same step. The spec requires concurrency but does not define tie-breaking. | specification.md |
| Loop holes | High | Open | Clarify whether starting positions are guaranteed unique and valid. The spec forbids overlapping mowers during execution, but does not explicitly define invalid initial states. | specification.md |
| Loop holes | High | Open | Clarify whether concurrency is a functional requirement or an implementation constraint. This affects whether deterministic sequential simulation is acceptable for correctness-first solutions. | specification.md |
| Loop holes | Medium | Open | Clarify expected behavior for malformed input: unknown orientation, unknown command, missing lines, extra whitespace, empty file. | specification.md |
| Loop holes | Medium | Open | Align repository fixtures with the current spec. The spec uses `L/R/F`, while the sample input file still contains `G/A` commands from a different command set. | specification.md, file.txt |
| Loop holes | Medium | Open | Clarify output formatting exactly: spacing, one mower per line, and whether output must omit debug text entirely. | specification.md, Main.go |

The highest-value candidate discriminator now is not buildability anymore, since that is fixed. It is whether the candidate can turn this prototype into a spec-driven design with a clean parser/simulator split, correct movement semantics, and a deterministic concurrency story.

# **Hiring Rubric**

Here is a candidate evaluation rubric derived from the updated action plan. It is structured to help you assess level through observable outcomes rather than style preferences.

| Area | Junior | Mid | Senior |
|---|---|---|---|
| Build and execution | Makes the project build successfully and keeps the executable runnable from the repository root. Usually stops at module/import fixes and basic command-line behavior. | Makes the project build and run predictably, removes hardcoded demo flow, and wires a real execution path from input to output. | Makes build and execution reproducible and defensible: clear entrypoint, reliable failure behavior, spec-aligned runtime behavior, and a credible path to CI. |
| Code quality | Fixes obvious defects when pointed out, such as broken imports or simple logic bugs. | Identifies and fixes correctness bugs proactively, including coordinate semantics and invalid output behavior. | Systematically improves correctness boundaries, reduces ambiguity in APIs, and prioritizes fixes based on risk and maintainability. |
| Programming patterns | Uses basic Go syntax correctly but may leave imperative or ad hoc flows in place. | Introduces typed commands, cleaner constructors, and more idiomatic flow instead of embedding business behavior in main. | Chooses patterns deliberately, keeps APIs small and stable, avoids overengineering, and explains tradeoffs clearly. |
| Code and repository architecture | Keeps most code in place and makes local fixes. | Splits parser, domain, and execution responsibilities into separate units once complexity justifies it. | Designs a small but scalable architecture with clear package boundaries and a coherent repository layout that supports future changes cleanly. |
| Separation of concerns | Can move a little logic out of main when asked. | Separates parsing, simulation, and rendering/output into distinct responsibilities with testable boundaries. | Enforces clean boundaries consistently and removes hidden coupling between I/O, domain logic, and orchestration. |
| Resource usage | Usually does not consider scanner limits, malformed input handling, or operational failure modes unless prompted. | Handles input validation and failure paths reasonably, and avoids obviously wasteful processing. | Thinks explicitly about bounded input, operational behavior, and failure modes without adding unnecessary complexity. |
| Performance | May mention concurrency because the spec says so, but typically does not implement it safely. | Can discuss or implement a practical execution strategy and avoid premature optimization. | Designs deterministic concurrent processing, understands contention tradeoffs, and can justify whether concurrency helps for the workload. |
| Testing | Adds a few happy-path tests after implementation. | Adds unit tests around geometry, movement, and parsing, usually table-driven. | Builds a layered test strategy: unit, integration, example-based acceptance, and race-aware validation for concurrent paths. |
| Specification handling | Implements the visible example but may miss ambiguous or underspecified behavior. | Identifies some gaps in the spec and makes reasonable assumptions. | Surfaces loopholes early, asks the right clarifying questions, and designs behavior that stays deterministic under ambiguity. |

## **What Good Evidence Looks Like**
| Level signal | Strong indicators |
|---|---|
| Junior | Gets go.mod and Main.go into a buildable state, fixes a few local bugs, adds a small number of basic tests. |
| Mid | Replaces the hardcoded flow in Main.go with parser-driven execution, adds meaningful tests around Mower.go and Rectangle.go, and produces correct output for the example in specification.md. |
| Senior | Produces a coherent design around the gaps in specification.md, explicitly addresses repository/spec mismatch with file.txt, and can defend deterministic concurrency and testing choices. |

## **Suggested Scoring**
Use a simple 0 to 2 score per area.
1. `0`: Missing or incorrect.
2. `1`: Adequate but partial.
3. `2`: Strong and well justified.

A practical interpretation:
1. Mostly `0-1` with few `2`: Junior.
2. Mostly `1` with some `2`: Mid.
3. Mostly `2`, especially on architecture, testing, and specification handling: Senior.

## **Critical differentiators for this repository**
1. Can the candidate move from prototype code in Main.go to a spec-driven design?
2. Do they catch the correctness issue in Rectangle.go without being told?
3. Do they notice the command mismatch between specification.md and file.txt?
4. Can they explain how to satisfy the concurrency requirement without introducing nondeterministic behavior?
