# Interview scorecard for this repository, ready to use live.

## **How To Use**
1. Score each criterion from 0 to 2.
2. Add weighted points.
3. Use the level bands at the end for final recommendation.
4. Capture short evidence notes during the interview.

## **Score Scale**
- 0: Not demonstrated or incorrect
- 1: Partially correct, guidance needed
- 2: Correct, autonomous, and well justified

## **A. Core Delivery (45 points)**

| Criterion | Weight | 0 | 1 | 2 | Evidence notes |
|---|---:|---|---|---|---|
| Build and run reliability | 10 | Cannot make project build | Builds with help, fragile run path | Builds and runs cleanly, explains why | |
| Spec-compliant input parsing | 10 | No parser / ad hoc parsing | Partial parser, weak validation | Complete parser with validation and clear errors | |
| Correct movement semantics | 10 | Wrong L/R/F behavior | Mostly correct, misses edge cases | Fully correct including boundaries | |
| Collision and occupancy rules | 10 | Ignores or breaks rule | Partial handling | Correctly enforces no shared cell occupancy | |
| Output contract adherence | 5 | Wrong format/order | Mostly right | Exact expected output and order | |

## **B. Architecture and Design (25 points)**

| Criterion | Weight | 0 | 1 | 2 | Evidence notes |
|---|---:|---|---|---|---|
| Separation of concerns | 10 | Logic mixed in main | Some split, still coupled | Clear parser/domain/simulator/output boundaries | |
| Programming patterns in Go | 8 | Non-idiomatic or unsafe patterns | Mixed quality | Idiomatic, simple, maintainable choices | |
| Repository/package structure | 7 | Disorganized | Acceptable with inconsistencies | Coherent layout supporting growth and testing | |

## **C. Performance and Resource Usage (15 points)**

| Criterion | Weight | 0 | 1 | 2 | Evidence notes |
|---|---:|---|---|---|---|
| Concurrency strategy | 8 | No strategy or nondeterministic | Basic parallelism, unclear correctness | Deterministic concurrent model with clear conflict handling | |
| Resource handling | 4 | Ignores errors/limits | Handles basics | Thoughtful handling of input size, errors, and execution behavior | |
| Performance discipline | 3 | Premature optimization | Some measurement awareness | Benchmarks or reasoned tradeoff decisions | |

## **D. Testing and Validation (15 points)**

| Criterion | Weight | 0 | 1 | 2 | Evidence notes |
|---|---:|---|---|---|---|
| Unit tests quality | 6 | Minimal/empty tests | Basic happy-path tests | Table-driven tests covering boundaries and failure cases | |
| Integration or acceptance coverage | 5 | None | Partial scenario coverage | End-to-end coverage for spec example and key rules | |
| Concurrency and robustness tests | 4 | None | Some attempts | Includes race-aware checks and conflict scenarios | |

## **Specification Loopholes Check (Gate)**
If candidate does not identify at least 2 of these, cap final recommendation to Mid:
1. Simultaneous move tie-break rules are ambiguous.
2. Initial position validity and uniqueness are underspecified.
3. Concurrency requirement vs determinism tradeoff needs explicit design.
4. Input mismatch risk between command alphabets and fixtures.
5. Malformed input behavior is not fully specified.

## **Weighted Scoring Formula**
- For each row: row points = weight × (score / 2)
- Final score = sum of all row points, max 100

## **Recommendation Bands**
- 85 to 100: Strong Senior
- 70 to 84: Senior-
- 55 to 69: Mid
- 40 to 54: Junior+
- Below 40: Junior / Not ready for this scope

## **Decision Add-ons**
- Hire signal: candidate is correct, explains tradeoffs, and keeps design simple.
- Concern signal: candidate adds complexity before proving correctness.
- Strong senior signal: candidate makes concurrency deterministic and testable, not just parallel.
