# 24. Swap Nodes in Pairs

## Intuition
Linked list

## Approach
// head=1->2->3->4
// p1=1, p2=2, dummy=n->1->2->3->4, prevNode=n
// swap-1
// prevNode.Next=2 >> dummy=n->2->3->4
// p2Next=3
// swap(p1,p2) >> p1=2, p2=1
// link next >> dummy=n->2->1->3->4
// continue next round >> prevNode=1, p1=3, p2=4
// swap-2
// prevNode.Next=4 >> dummy=n->2->1->4
// p2Next=n
// swap(p1,p2) >> p1=4, p2=3
// link next >> dummy=n->2->1->4->3
// continue next round >> prevNode=3, p1=n, break