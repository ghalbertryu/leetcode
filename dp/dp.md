## DP 就是 Dynamic Programming（動態規劃）

簡單來說，它是一種 把大問題拆解成小問題，並把小問題的解存起來重複利用 的演算法思維。

### 特點
1. 重疊子問題 (Overlapping Subproblems)許多子問題會被重複計算，
例如費氏數列 F(n) = F(n-1) + F(n-2)。
DP 會把計算過的結果存起來 (通常用陣列或 hash map)，避免重算。
2. 最優子結構 (Optimal Substructure)
大問題的解可以由小問題的解組合而成。
例如找最短路徑，從 A 到 C 的最短路徑可能經過 B，那就只需要「A→B」和「B→C」的最短路徑。