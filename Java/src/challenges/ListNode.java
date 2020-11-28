package challenges;

import java.util.ArrayList;
import java.util.List;

public class ListNode {
    int val;
    ListNode next;

    public ListNode() {}
    public ListNode(int val) {
        this.val = val;
    }
    public ListNode(int val, ListNode next) {
        this.val = val;
        this.next = next;
    }

    /**
     * Let's try making a linked list based on a List of ints
     * @param list
     */
    public static ListNode fromList(List<Integer> list) {
        ListNode thisNode = new ListNode(list.get(0));
        ListNode tempNode = thisNode;
        for (int i = 1; i < list.size(); i++) {
            ListNode currNode = new ListNode(list.get(i));
            tempNode.next = currNode;
            tempNode = tempNode.next;
        }
        return thisNode;
    }

    public static String listToString(ListNode theNode) {
        List<Integer> nodeVals = new ArrayList<>();
        while (theNode != null) {
            nodeVals.add(theNode.val);
            theNode = theNode.next;
        }
        return nodeVals.toString();
    }
}
