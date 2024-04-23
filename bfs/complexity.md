The time complexity of this algorithm can be analyzed as follows:

In the worst case, each person in the graph needs to be visited once, as we're performing a breadth-first search (BFS)
traversal.
At each step of the traversal, we add the neighbors of the current person to the queue, which takes O(k) time, where k
is the number of neighbors.
Since each person is processed once, and adding neighbors to the queue takes O(k) time for each person, the overall time
complexity is O(V + E), where V is the number of vertices (people) and E is the number of edges (relationships between
people).

The space complexity of this algorithm can be analyzed as follows:

We use a queue to keep track of the people to check. In the worst case, the queue may contain all the vertices in the
graph before finding the mango seller.
Additionally, we use a map to keep track of checked people to avoid revisiting them, which may also contain all the
vertices in the graph.
Therefore, the space complexity is O(V), where V is the number of vertices (people) in the graph.

Overall:
Time Complexity: O(V + E)
Space Complexity: O(V)