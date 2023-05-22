// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

//用solidity实现List的双向链表的结构
//运用数组实现双向链表
//数据结构，使用2个结构体，1表示节点，2表示链表
//提供四个方法，增 删 改 查
struct Node {
    uint prev;
    uint next;
    uint data;
}

struct List {
    uint first;
    uint last;
    Node[] nodes;
}

contract List2 {
    uint constant public Null = 0;

    function _add(List storage list, uint data) internal returns(bool) {
        if(list.nodes.length == Null) {
            Node memory node = Node(Null, Null, data);
            list.nodes.push(node);
        }
        Node memory node = Node(0, 0, data);
        return true;
    }


}