package main

import (
	"sync/atomic"
	"testing"
)

// TestNewLockFreeStack 测试 NewLockFreeStack 函数
func TestNewLockFreeStack(t *testing.T) {
	stack := NewLockFreeStack()
	if stack == nil {
		t.Error("Expected a non-nil stack, got nil")
	}
}

// TestPush 测试 Push 函数
func TestPush(t *testing.T) {
	stack := NewLockFreeStack()

	// 测试用例1：向空栈中压入一个元素
	stack.Push(1)
	top := (*Node)(atomic.LoadPointer(&stack.top))
	if top.v != 1 {
		t.Errorf("Expected top element to be 1, got %v", top.v)
	}

	// 测试用例2：向非空栈中压入多个元素
	stack.Push(2)
	top = (*Node)(atomic.LoadPointer(&stack.top))
	if top.v != 2 {
		t.Errorf("Expected top element to be 2, got %v", top.v)
	}
}

// TestPop 测试 Pop 函数
func TestPop(t *testing.T) {
	stack := NewLockFreeStack()

	// 测试用例1：从空栈中弹出元素
	value, ok := stack.Pop()
	if ok {
		t.Error("Expected Pop to return false for empty stack, got true")
	}
	if value != nil {
		t.Errorf("Expected nil value from Pop on empty stack, got %v", value)
	}

	// 测试用例2：从非空栈中弹出元素
	stack.Push(1)
	stack.Push(2)
	value, ok = stack.Pop()
	if !ok {
		t.Error("Expected Pop to return true for non-empty stack, got false")
	}
	if value != 2 {
		t.Errorf("Expected popped value to be 2, got %v", value)
	}

	// 检查栈顶元素是否被正确移除
	top := (*Node)(atomic.LoadPointer(&stack.top))
	if top.v != 1 {
		t.Errorf("Expected top element to be 1 after Pop, got %v", top.v)
	}
}
