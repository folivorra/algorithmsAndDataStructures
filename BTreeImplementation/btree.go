package BTreeImplementation

import "errors"

/*
T - это минимальная степень дерева, в узле может быть T-1 <= X <= 2T-1 ключей
*/
const T = 2

/*
KVPair - это структура для представления ассоциативного элемента ключ-значение
*/
type KVPair struct {
	key   int
	value string
}

/*
BTreeNode - это структура для представления узла дерева
*/
type BTreeNode struct {
	leaf     bool
	pairs    []KVPair
	children []*BTreeNode
}

/*
BTree - это структура для представления дерева, в себе имеет только указатель на корень,
остальные узлы образуются рекурсивно
*/
type BTree struct {
	root *BTreeNode
}

/*
newNode инициализирует с дефолтными значениями узел,
leaf - флаг, указывающий на то является ли узел листовым
*/
func newNode(leaf bool) *BTreeNode {
	return &BTreeNode{
		leaf:     leaf,
		pairs:    []KVPair{},
		children: []*BTreeNode{},
	}
}

/*
splitChild поднимает средний ключ к родителю и разделяет узел,
parent - указатель на родительский узел,
i - индекс переполненного потомка
*/
func (t *BTree) splitChild(parent *BTreeNode, i int) {
	// T это середина потомка по которому будет идти разрез
	T := len(parent.children[i].pairs) / 2
	child := parent.children[i]
	// узел который будет отделяться "вправо"
	rightNode := newNode(child.leaf)

	// перенос половины ключей в новый узел
	rightNode.pairs = append(rightNode.pairs, child.pairs[T+1:]...)
	// ключ для вставки в родительский узел
	KVForInsert := KVPair{child.pairs[T].key, child.pairs[T].value}
	// старый узел очищается на перенесенную половину
	child.pairs = child.pairs[:T]

	// перенос потомков если потомок не был листовым
	if !child.leaf {
		rightNode.children = append(rightNode.children, child.children[T+1:]...)
		child.children = child.children[:T+1]
	}

	// вставка нового узла в родителя
	// сначала поднимаем ключ в родительский узел на позицию i
	parent.pairs = append(parent.pairs[:i], append([]KVPair{KVForInsert}, parent.pairs[i:]...)...)
	// затем новосозданный узел с ключами большими чем child.pairs[T] добавляется в список детей родителя
	parent.children = append(parent.children[:i+1], append([]*BTreeNode{rightNode}, parent.children[i+1:]...)...)
}

/*
insertNonFull вставляет ключ в неполный узел,
node - указатель на узел, в который будет производится вставка,
key и value - ключ и значение для вставки
*/
func (t *BTree) insertNonFull(node *BTreeNode, key int, value string) {
	i := len(node.pairs) - 1
	// если узел листовой, то вставляем в нужное место
	if node.leaf {
		// добавление пустой структуры чтобы избежать паники выхода за границы
		node.pairs = append(node.pairs, KVPair{})
		// в цикле ищем место для ключа
		for i >= 0 && key < node.pairs[i].key {
			node.pairs[i+1] = node.pairs[i]
			i--
		}
		// вставка ключа
		node.pairs[i+1] = KVPair{key, value}
	} else {
		// если узел не листовой то сначала ищем нужного потомка
		for i >= 0 && key < node.pairs[i].key {
			i--
		}
		i++
		// если дочерний узел переполнен, разделяем его
		if len(node.children[i].pairs) == 2*T-1 {
			t.splitChild(node, i)
			// выбираем в какой из двух узлов пойдем
			if key > node.pairs[i].key {
				i++
			}
		}

		// рекурсивно вставляем в дочерний узел
		t.insertNonFull(node.children[i], key, value)
	}
}

/*
Insert вставляет ключ и значение в дерево,
key и value - ключ и значение для вставки
*/
func (t *BTree) Insert(key int, value string) {
	// если дерево пустое, создаем корень
	if t.root == nil {
		t.root = newNode(true)
		t.root.pairs = append(t.root.pairs, KVPair{key, value})
		return
	}

	// если корень переполнен, создаем новый корень и разделяем старый
	if len(t.root.pairs) == 2*T-1 {
		oldRoot := t.root
		t.root = newNode(false)
		t.root.children = append(t.root.children, oldRoot)
		t.splitChild(t.root, 0)
	}

	// вставляем ключ в неполный узел
	t.insertNonFull(t.root, key, value)
}

/*
Search ищет значение по ключу,
key - ключ для поиск,
функция возвращает указатель на узел и индекс,
по которым можно обратиться к искомому значению (node.pairs[i].value)

количество обращений к диску =  O(h), где h - высота дерева
время вычислений = O(t*logₜn), где n - кол-во узлов
*/
func (t *BTree) Search(key int) (string, error) {
	if t.root == nil {
		return "", errors.New("root is nil")
	}
	return searchRecursively(t.root, key)
}

func searchRecursively(node *BTreeNode, key int) (string, error) {
	index := 0
	for index < len(node.pairs) && key > node.pairs[index].key {
		index++
	}
	if index < len(node.pairs) && key == node.pairs[index].key {
		return node.pairs[index].value, nil
	} else if node.leaf {
		return "", errors.New("value not found")
	} else {
		return searchRecursively(node.children[index], key)
	}
}
