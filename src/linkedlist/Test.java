package linkedlist;

public class Test {
    public static void main(String[] args) {
        DoublyLinkedList doublyLinkedList = new DefaultDoublyLinkedList();
        doublyLinkedList.add("Ahihi");
        doublyLinkedList.add("ahuhu");
        System.out.println(doublyLinkedList.toString());
        doublyLinkedList.addFirst("first");
        doublyLinkedList.addLast("last");
        System.out.println(doublyLinkedList.toString());
        System.out.println(doublyLinkedList.contains("Ahihi"));
        System.out.println(doublyLinkedList.indexOf("ahuhu"));
        doublyLinkedList.removeAt(2);
        System.out.println(doublyLinkedList.toString());
    }
}
