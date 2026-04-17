# **Interview Scorecard: Mower-Go (Go Backend)**

Candidate: ____________________  
Date: ____________________  
Interviewer: ____________________  
Duration: ____________________

Scoring per criterion: `0 = Not demonstrated`, `1 = Partial`, `2 = Strong`  
Weighted points per row: `Weight x (Score / 2)`

---

## **A) Core Delivery (45 pts)**

| Criterion | Weight | Score (0-2) | Weighted points | Notes / Evidence |
|---|---:|---:|---:|---|
| Build and run reliability | 10 |  |  |  |
| Spec-compliant parsing | 10 |  |  |  |
| Movement correctness (L/R/F + boundaries) | 10 |  |  |  |
| Collision and occupancy handling | 10 |  |  |  |
| Output format and ordering | 5 |  |  |  |

Subtotal A: ______ / 45

---

## **B) Architecture and Design (25 pts)**

| Criterion | Weight | Score (0-2) | Weighted points | Notes / Evidence |
|---|---:|---:|---:|---|
| Separation of concerns (parser/domain/simulator/output) | 10 |  |  |  |
| Go programming patterns (clarity, idioms, tradeoffs) | 8 |  |  |  |
| Repository/package architecture quality | 7 |  |  |  |

Subtotal B: ______ / 25

---

## **C) Performance and Resource Usage (15 pts)**

| Criterion | Weight | Score (0-2) | Weighted points | Notes / Evidence |
|---|---:|---:|---:|---|
| Concurrency strategy + determinism | 8 |  |  |  |
| Resource handling (input limits, errors, runtime behavior) | 4 |  |  |  |
| Performance discipline (measurement before optimization) | 3 |  |  |  |

Subtotal C: ______ / 15

---

## **D) Testing and Validation (15 pts)**

| Criterion | Weight | Score (0-2) | Weighted points | Notes / Evidence |
|---|---:|---:|---:|---|
| Unit test depth and quality | 6 |  |  |  |
| Integration / acceptance coverage | 5 |  |  |  |
| Concurrency robustness (`-race`, conflict scenarios) | 4 |  |  |  |

Subtotal D: ______ / 15

---

## **Specification Loopholes Check (Gate)**  
Mark each item identified by candidate.

- [ ] Tie-break rule for simultaneous moves to same target cell is ambiguous  
- [ ] Initial mower positions validity/uniqueness not fully specified  
- [ ] Concurrency requirement vs deterministic output needs explicit design  
- [ ] Malformed input behavior not fully defined  
- [ ] Command set/fixture mismatch risks are recognized

Loopholes identified count: ______

Rule: if fewer than 2 identified, cap recommendation at Mid.

---

## **Total Score**
- A + B + C + D = ______ / 100

## **Recommendation Band**
- 85-100: Strong Senior
- 70-84: Senior-
- 55-69: Mid
- 40-54: Junior+
- <40: Junior / Not ready

Final recommendation: ____________________

---

## **Quick Debrief Prompts**
1. What tradeoff did the candidate make between correctness and speed?
2. Did they keep the design simple while still extensible?
3. What did they test first, and why?
4. Would you trust this person to own the next iteration independently?