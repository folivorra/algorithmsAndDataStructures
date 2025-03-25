/*
Package BTreeImplementation вся реализация структуры btree написана согласно книге
"Алгоритмы: построение и анализ" за авторством Томаса Кормена и его коллег
*/
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
	// если дерево неинициализировано, создаем корень
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
		// 0 потому что старый корень будет единственным потомком для нового корня
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

количество обращений к диску =  O(h), где h - высота дерева;
время вычислений = O(t*logₜn), где n - кол-во узлов
*/
func (t *BTree) Search(key int) (string, error) {
	if t.root == nil {
		return "", errors.New("root is nil")
	}
	// если корень инициализирован вызываем рекурсивную функцию для поиска
	return searchRecursively(t.root, key)
}

func searchRecursively(node *BTreeNode, key int) (string, error) {
	index := 0
	// перебираем ключи в узле
	for index < len(node.pairs) && key > node.pairs[index].key {
		index++
	}
	// если ключ совпадает возвращаем найденное значение
	// если узел листовой - некуда спускаться - ключа не существует в дереве
	// иначе спускаемся дальше по дереву
	if index < len(node.pairs) && key == node.pairs[index].key {
		return node.pairs[index].value, nil
	} else if node.leaf {
		return "", errors.New("key not found")
	} else {
		return searchRecursively(node.children[index], key)
	}
}

/*
Update обновляет значение по ключу,
во всем работает аналогично методу Search
*/
func (t *BTree) Update(key int, value string) error {
	if t.root == nil {
		return errors.New("root is nil")
	}
	return updateRecursively(t.root, key, value)
}

func updateRecursively(node *BTreeNode, key int, value string) error {
	index := 0
	// все также как в поиске
	for index < len(node.pairs) && key > node.pairs[index].key {
		index++
	}
	if index < len(node.pairs) && key == node.pairs[index].key {
		node.pairs[index].value = value
		return nil
	} else if node.leaf {
		return errors.New("key not found")
	} else {
		return updateRecursively(node.children[index], key, value)
	}
}

func (t *BTree) Delete(key int) {
	// если корень неинициализирован - завершаем функцию
	// иначе вызываем рекурсивную функцию для удаления
	if t.root == nil {
		return
	} else {
		deleteRecursively(t.root, key)
	}
}

func deleteRecursively(node *BTreeNode, key int) {
	index := 0
	for index < len(node.pairs) && key > node.pairs[index].key {
		index++
	}
	/*
		если нашли ключ и узел листовой - просто удаляем ключ и завершаем функцию
		если ключ найден и узел нелистовой:
		1)	если существует такой потомок Y предшевствующий ключу key, у которого минимум T ключей,
			то находим key` - предшественника key в поддереве, корнем которого ялвяется Y. Рекурсивно
			удаляем key` и заменяем key ключом key`;
		2)	иначе если существует такой потомок Z следующий за ключом key, у которого минимум T ключей,
			то находим key` - следующего за key в поддереве, корнем которого является Z. Рекурсивно
			удаляем key` и заменяем key ключом key`;
		3)	иначе вносим key и все ключи из Z в Y (при этом из изначального узла удаляются и key, и указатель Z),
			затем рекурсивно удаляем key из Y.
	*/
	if node.leaf && index < len(node.pairs) && key == node.pairs[index].key {
		node.pairs = append(node.pairs[:index], node.pairs[index+1:]...)
		return
	} else if node.leaf && !(index < len(node.pairs) && key == node.pairs[index].key) {
		return
	} else if !node.leaf && index < len(node.pairs) && key == node.pairs[index].key {
		if len(node.children[index].pairs) >= T { // 1)
			k := node.children[index].pairs[len(node.children[index].pairs)-1]
			deleteRecursively(node.children[index], k.key)
			node.pairs = append(node.pairs[:index], append([]KVPair{k}, node.pairs[index+1:]...)...)
		} else if len(node.children[index+1].pairs) >= T { // 2)
			k := node.children[index+1].pairs[0]
			deleteRecursively(node.children[index+1], k.key)
			node.pairs = append(node.pairs[:index], append([]KVPair{k}, node.pairs[index+1:]...)...)
		} else { // 3)
			zKeys := node.children[index+1].pairs
			node.children[index].pairs = append(node.children[index].pairs, append([]KVPair{node.pairs[index]}, zKeys...)...)
			node.pairs = append(node.pairs[:index], node.pairs[index+1:]...)
			node.children = append(node.children[:index+1], node.children[index+2:]...)
			deleteRecursively(node.children[index], key)
		}

		/*
			если ключа key нет во внутреннем листе, находим корень node.children[index] поддерева,
			которое должно содержать k.
			если node.children[index] содержит T-1 ключей, выполняем (чтобы гарантировать, что далее мы
			переходим в узел минимум с t ключами):
			1)	если один из непосредственных соседей-узлов содержит как минимум T ключей, передадим
				в node.children[index] ключ-разделитель между данным узлом и соседним, на его место поместим
				крайний ключ из соседнего узло и (если соседний узел нелистовой) переместим
				соответсвующий указатель из соседнего узла в node.children[index];
			2)
			далее мы рекурсивно удаляем key из соответствуюшего узла.
		*/

	} else if !node.leaf && !(index < len(node.pairs) && key == node.pairs[index].key) {
		next := (index != len(node.children)-1) && (len(node.children[index+1].pairs) >= T)
		prev := (index != 0) && (len(node.children[index-1].pairs) >= T)
		if (len(node.children[index].pairs) == T-1) && prev {
			leftNode := node.children[index-1]
			child := node.children[index]

			child.pairs = append([]KVPair{node.pairs[index-1]}, child.pairs...)
			node.pairs[index-1] = leftNode.pairs[len(leftNode.pairs)-1]
			leftNode.pairs = leftNode.pairs[:len(leftNode.pairs)-1]

			if !leftNode.leaf {
				child.children = append([]*BTreeNode{leftNode.children[len(leftNode.children)-1]}, child.children...)
				leftNode.children = leftNode.children[:len(leftNode.children)-1]
			}
		} else if (len(node.children[index].pairs) == T-1) && next {
			rightNode := node.children[index+1]
			child := node.children[index]

			child.pairs = append(child.pairs, node.pairs[index])
			node.pairs[index] = rightNode.pairs[0]
			rightNode.pairs = rightNode.pairs[1:]

			if !rightNode.leaf {
				child.children = append(child.children, rightNode.children[0])
				rightNode.children = rightNode.children[1:]
			}
		} else {
			// TODO: доделать последнее условие со слиянием узлов!!
		}
		deleteRecursively(node.children[index], key)
	}
}
