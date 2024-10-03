class ListNode {
    int val;
    ListNode next;
    ListNode() {}
    ListNode(int val) { this.val = val; }
    ListNode(int val, ListNode next) { this.val = val; this.next = next; }

    public static ListNode deserialize(String data) {
        if (data.equals("[]") || data.isEmpty()) return null;
        String[] nodes = data.replaceAll("\\[|\\]", "").split(",");
        ListNode head = new ListNode(Integer.parseInt(nodes[0].trim()));
        ListNode current = head;
        for (int i = 1; i < nodes.length; i++) {
            current.next = new ListNode(Integer.parseInt(nodes[i].trim()));
            current = current.next;
        }
        return head;
    }

    public static String serialize(ListNode head) {
        if (head == null) return "[]";
        StringBuilder sb = new StringBuilder("[");
        while (head != null) {
            sb.append(head.val);
            if (head.next != null) sb.append(",");
            head = head.next;
        }
        return sb.append("]").toString();
    }
}

public class Solution {
    public ListNode mergeTwoLists(ListNode list1, ListNode list2) {
        if (list1 == null) return list2;
        if (list2 == null) return list1;
        if (list1.val < list2.val) {
            list1.next = mergeTwoLists(list1.next, list2);
            return list1;
        } else {
            list2.next = mergeTwoLists(list1, list2.next);
            return list2;
        }
    }

    public static void main(String[] args) {
        ListNode list1 = ListNode.deserialize("[1,2,4]");
        ListNode list2 = ListNode.deserialize("[1,3,4]");
        Solution sol = new Solution();
        System.out.println(ListNode.serialize(sol.mergeTwoLists(list1, list2)));
    }
}
