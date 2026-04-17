# 30-minute and 60-minute interview variants below, aligned with your scorecard and this repository’s risk areas.

## **30-Minute Variant (Signal Check)**
Goal: quickly identify clear no-hire vs potential mid/senior.

1. Setup and context (3 min)
- Ask candidate to summarize the problem and key constraints.
- Expected signal: mentions parsing, movement rules, collisions, ordering, and concurrency ambiguity.

2. Fast code reading (5 min)
- Candidate inspects current implementation and lists top 3 risks.
- Expected signal: identifies incomplete parser/simulator, rectangle axis bug risk, and missing tests.

3. Live implementation slice (15 min)
- Task: implement minimal parser plus command execution for one mower with correct boundaries.
- Must produce deterministic output for one provided input.
- Expected signal: clean function boundaries, simple data model, basic validation.

4. Testing slice (5 min)
- Add 2 to 3 focused tests (table-driven preferred).
- Expected signal: edge cases, not only happy path.

5. Debrief (2 min)
- Ask: how would you extend to multi-mower + deterministic concurrency?

Scoring emphasis:
- Core Delivery 50%
- Testing 25%
- Architecture/SoC 20%
- Performance/Resource 5%

Pass guidance:
- Junior+: parses + executes basic commands correctly with at least one edge test.
- Mid: structured design + solid tests + clear extension strategy.
- Senior signal in 30 min: proactively calls out spec loopholes and deterministic conflict model.

---

## **60-Minute Variant (Level Differentiator)**
Goal: distinguish mid vs senior through architecture and correctness under constraints.

1. Problem framing and assumptions (5 min)
- Candidate states assumptions and open spec questions.
- Expected signal: identifies at least 2 loopholes.

2. Design first (10 min)
- Whiteboard/verbally define:
- parser contract
- domain types
- simulator contract
- output contract
- Expected signal: explicit separation of concerns.

3. Implementation phase 1 (20 min)
- Build parser + sequential simulator for multiple mowers.
- Enforce boundaries and output ordering.
- Expected signal: readable flow, explicit error handling, no overengineering.

4. Implementation phase 2 (10 min)
- Add collision rule (no shared cell).
- Candidate explains deterministic behavior for conflicting moves.

5. Testing and validation (10 min)
- Add unit tests + one acceptance test from sample.
- Optional: race-aware plan for concurrent extension.

6. Debrief and scaling discussion (5 min)
- Ask for concurrency approach that preserves determinism and testability.
- Ask what they would benchmark and why.

Scoring emphasis:
- Core Delivery 40%
- Architecture/SoC 25%
- Testing 20%
- Performance/Resource 15%

Pass guidance:
- Mid: complete sequential multi-mower simulation with robust tests.
- Senior-: plus clear deterministic concurrency design and strong spec-gap handling.
- Strong Senior: plus explicit tradeoffs, clean abstractions, and high signal test strategy.

---

## **Interviewer Prompts To Reuse**
1. Which assumptions did you make because the spec is ambiguous?
2. How do you guarantee deterministic output with concurrent execution?
3. What would you test first if this code goes to CI today?
4. What is the smallest refactor that improves maintainability without slowing delivery?

Ready-To-Print Interviewer Script (60 minutes, Go Backend)

Candidate: ____________________  
Interviewer: ____________________  
Date: ____________________  

## 1. Opening (2 min)
Prompt:
1. “You’ll work on a mower simulation in Go. Please think aloud and explain assumptions.”
2. “Focus on correctness first, then structure, then performance.”

Strong signals:
1. Clarifies requirements before coding.
2. Mentions deterministic behavior and testability early.

Weak signals:
1. Starts coding immediately with no requirements check.
2. Treats ambiguity as irrelevant.

---

## 2. Problem Framing (8 min)
Ask exactly:
1. “Summarize the core rules in your own words.”
2. “What are the top 3 risks in the current implementation?”
3. “What specification ambiguities do you want to resolve before coding?”

Expected strong answers:
1. Rules: parsing, movement, boundaries, collisions, ordered output.
2. Risks: missing parser/simulator split, weak tests, correctness bug risk in coordinate semantics.
3. Ambiguities: simultaneous conflict tie-break, invalid initial states, malformed input behavior.

Expected weak answers:
1. Only describes happy path.
2. Misses collisions or ordered output requirement.
3. Says “spec is clear enough” without naming gaps.

Score hints:
1. This section maps mostly to Architecture/SoC and Specification Loopholes gate.

---

## 3. Design Discussion (10 min)
Ask exactly:
1. “Describe your minimal architecture before implementation.”
2. “Where do parsing, domain logic, simulation, and output formatting live?”
3. “How will you keep concurrency deterministic if we add it?”

Expected strong answers:
1. Small components with clear boundaries.
2. `main` as orchestration only.
3. Deterministic tick/step model or clear conflict-resolution contract.

Expected weak answers:
1. Everything stays in one large function.
2. I/O mixed into domain entities.
3. “Use goroutines” with no deterministic conflict plan.

Score hints:
1. This section maps to Separation of Concerns, Programming Patterns, and Performance strategy.

---

## 4. Implementation Task (20 min)
Give this task:
1. “Implement sequential, spec-compliant simulation for multiple mowers.”
2. “Support `L/R/F`, boundaries, collision discard, and ordered final output.”
3. “Do not optimize prematurely. Keep code simple.”

While coding, ask:
1. “Why did you choose this data structure for occupancy?”
2. “What error handling policy are you applying for malformed input?”
3. “What part would you isolate first for tests?”

Strong signals:
1. Small pure functions and explicit data types.
2. Clear error paths and input validation.
3. Avoids overengineering while maintaining extensibility.

Weak signals:
1. Hardcoded flows or hidden global state.
2. Silent invalid input behavior without rationale.
3. Adds concurrency before core correctness.

---

## 5. Testing Task (12 min)
Ask exactly:
1. “Add tests that prove boundary behavior.”
2. “Add tests for collision discard and output order.”
3. “Add one acceptance-style test from sample input/output.”

Expected strong answers:
1. Table-driven tests.
2. Includes edge and failure cases, not only happy path.
3. Clear assertions on deterministic output.

Expected weak answers:
1. One superficial happy-path test.
2. No collision or malformed input coverage.
3. No explanation of what behavior each test protects.

Score hints:
1. This section maps directly to Unit, Integration, and Robustness criteria.

---

## 6. Concurrency and Scale Discussion (6 min)
Ask exactly:
1. “How would you parallelize safely on multi-CPU machines?”
2. “How do you preserve deterministic results?”
3. “What would you benchmark first and why?”

Expected strong answers:
1. Defines a deterministic execution model under concurrency.
2. Explicitly handles contention and conflicts.
3. Benchmarks parser throughput and simulation step cost before micro-optimizing.

Expected weak answers:
1. “Add goroutines everywhere.”
2. No tie-break strategy for simultaneous moves.
3. Optimization claims without measurement plan.

---

## 7. Final Debrief (2 min)
Ask exactly:
1. “What tradeoff did you make intentionally?”
2. “What would be your next 2 improvements if this goes to production?”

Strong signals:
1. Can defend tradeoffs between speed, correctness, and maintainability.
2. Prioritizes tests and architecture before fancy optimization.

Weak signals:
1. No clear tradeoff awareness.
2. Random roadmap with no risk prioritization.

---

## Quick Scoring Overlay
Use your existing weighted scorecard and apply this interpretation during debrief:

1. `Strong Senior`: deterministic design, clean separation, high-signal tests, clear loophole handling.
2. `Mid`: correct implementation and decent structure/tests, but weaker on ambiguity handling or concurrency strategy.
3. `Junior`: partial correctness, limited tests, heavy guidance needed on architecture and edge cases.

## Red Flags To Record
1. Ignores ambiguous requirements.
2. Confuses correctness with optimization.
3. Cannot explain own design decisions.
4. Writes code that works only for the sample input.
