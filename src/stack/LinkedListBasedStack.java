package stack;

import linkedlist.DefaultDoublyLinkedList;
import linkedlist.DoublyLinkedList;

import java.util.EmptyStackException;
import java.util.Iterator;

public class LinkedListBasedStack<T> implements StackADT<T> {
    private final DoublyLinkedList<T> list = new DefaultDoublyLinkedList<>();
    public LinkedListBasedStack() {}
    public LinkedListBasedStack(T element) {
        push(element);
    }
    @Override
    public void push(T element) {
        list.addLast(element);
    }

    @Override
    public T pop() {
        if (isEmpty()) throw new EmptyStackException();
        return list.removeLast();
    }

    @Override
    public T top() {
        if(isEmpty()) throw new EmptyStackException();
        return list.peekLast();
    }

    @Override
    public int size() {
        return list.size();
    }

    @Override
    public boolean isEmpty() {
        return list.isEmpty();
    }

    @Override
    public Iterator iterator() {
        return list.iterator();
    }
}
