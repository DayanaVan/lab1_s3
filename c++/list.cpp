#include "list.h"

// Конструктор узла списка
ListNode::ListNode(std::string key)
{
    this->key = key;      
    this->next = nullptr; 
    this->prev = nullptr; 
}

// Конструктор списка
List::List()
{
    this->tail = nullptr; 
    this->head = nullptr; 
}

// Вставка элемента по индексу
void List::insert(std::string key, int index)
{
    // Если индекс 0 - вставляем в начало
    if(index == 0)
    {
        addHead(key);
        return;
    }
    
    // Находим узел, после которого нужно вставить новый элемент
    ListNode *node = this->head;
    for(int i = 0; i < index - 1; i++)
    {
        if(node == nullptr) // Если список закончился раньше - выходим
            break;
        node = node->next;
    }
    
    if(node == nullptr) // Если узел не найден - выходим
        return;
        
    // Если следующий узел - хвост, добавляем в конец
    if(node->next == this->tail)
    {
        addTail(key);
        return;
    }
    
    // Создаем новый узел и связываем его с соседними
    ListNode *newNode = new ListNode(key);
    newNode->next = node->next;
    newNode->prev = node;
    node->next = newNode;
    newNode->next->prev = newNode; 
}

// Удаление элемента по индексу
void List::remove(int index)
{
    // Если индекс 0 - удаляем из начала
    if(index == 0)
    {
        removeHead();
        return;
    }

    // Находим узел, предшествующий удаляемому
    ListNode *node = this->head;
    for(int i = 0; i < index - 1; i++)
    {
        if(node == nullptr) // Если список закончился раньше - выходим
            break;
        node = node->next;
    }
    
    if(node == nullptr) // Если узел не найден - выходим
        return;
        
    // Если следующий узел - хвост, удаляем хвост
    if(node->next == this->tail)
    {
        removeTail();
        return;
    }
    
    // Удаляем узел и перестраиваем связи
    ListNode *toDelete = node->next;
    if(toDelete != nullptr)
    {
        node->next = toDelete->next;
        toDelete->next->prev = node;
        delete toDelete;
    }
}

// Добавление элемента в конец списка
void List::addTail(std::string key)
{
    ListNode *node = new ListNode(key);
    
    // Если список пустой, новый узел становится и головой и хвостом
    if(this->tail == nullptr)
    {
        this->tail = node;
        this->head = node;
        return;
    }
    
    // Добавляем узел в конец и обновляем хвост
    node->prev = this->tail;
    this->tail->next = node;
    this->tail = node;
}

// Добавление элемента в начало списка
void List::addHead(std::string key)
{
    ListNode *node = new ListNode(key);
    
    // Если список пустой, новый узел становится и головой и хвостом
    if(this->tail == nullptr)
    {
        this->tail = node;
        this->head = node;
        return;
    }
    
    // Добавляем узел в начало и обновляем голову
    node->next = this->head;
    this->head->prev = node;
    this->head = node;
}

// Удаление последнего элемента
void List::removeTail()
{
    if(this->tail == nullptr) // Если список пуст - ничего не делаем
        return;
        
    // Если в списке только один элемент
    if(this->tail == this->head)
    {
        delete this->head;
        this->tail = nullptr;
        this->head = nullptr;
        return;
    }
    
    // Удаляем хвост и обновляем указатели
    this->tail = this->tail->prev;
    delete this->tail->next;
    this->tail->next = nullptr;
}

// Удаление первого элемента
void List::removeHead()
{
    if(this->head == nullptr) // Если список пуст - ничего не делаем
        return;
        
    // Если в списке только один элемент
    if(this->tail == this->head)
    {
        delete this->tail;
        this->tail = nullptr;
        this->head = nullptr;
        return;
    }
    
    // Удаляем голову и обновляем указатели
    this->head = this->head->next;
    delete this->head->prev;
    this->head->prev = nullptr;
}

// Печать списка с начала до конца
std::string List::printFromHead()
{
    ListNode *node = this->head;
    if(node == nullptr) // Если список пуст - возвращаем пустую строку
        return "";
        
    std::string a = "";
    // Проходим по всем узлам, кроме последнего
    while(node->next != nullptr)
    {
        a += node->key + ' '; 
        node = node->next;
    }
    a += node->key; 
    return a;
}

// Печать списка с конца до начала
std::string List::printFromTail()
{
    ListNode *node = this->tail;
    if(node == nullptr) // Если список пуст - возвращаем пустую строку
        return "";
        
    std::string a = "";
    // Проходим по всем узлам, кроме первого
    while(node->prev != nullptr)
    {
        a += node->key + ' ';
        node = node->prev;
    }
    a += node->key;
    return a;
}

// Удаление элемента по ключу и номеру вхождения
void List::removeKey(std::string key, int num)
{
    if(num < 0) 
        return;
        
    int n = 0;
    ListNode *node = this->head;
    
    // Поиск нужного вхождения ключа
    while(node != nullptr)
    {
        if(node->key == key)
        {
            if(n == num) // Нашли нужное вхождение
            {
                // В зависимости от положения узла используем соответствующий метод
                if(node == this->tail)
                {
                    removeTail();
                }
                else if(node == this->head)
                {
                    removeHead();
                }
                else
                {
                    // Удаляем узел из середины списка
                    node->prev->next = node->next;
                    node->next->prev = node->prev;
                    delete node;
                }
                return;
            }
            n++;
        }
        node = node->next;
    }
}

// Поиск элемента по ключу и номеру вхождения
ListNode *List::find(std::string key, int num)
{
    if(num < 0) // Проверка корректности номера вхождения
        return nullptr;
        
    int n = 0;
    ListNode *node = this->head;
    
    // Поиск нужного вхождения ключа
    while(node != nullptr)
    {
        if(node->key == key)
        {
            if(n == num) // Нашли нужное вхождение
                return node;
            n++;
        }
        node = node->next;
    }
    return nullptr; // Не нашли
}

// Добавление элемента перед заданным значением
void List::addBefore(std::string target, std::string key, int occurrence)
{
    if (occurrence < 0) // Проверка корректности номера вхождения
        return;

    int n = 0;
    ListNode* current = this->head;

    // Поиск нужного вхождения target
    while (current != nullptr)
    {
        if (current->key == target)
        {
            if (n == occurrence) // Нашли нужное вхождение
            {
                if (current == this->head) // Если target в начале списка
                {
                    addHead(key);
                }
                else // Если target в середине списка
                {
                    ListNode* newNode = new ListNode(key);
                    newNode->next = current;
                    newNode->prev = current->prev;
                    current->prev->next = newNode;
                    current->prev = newNode;
                }
                return;
            }
            n++;
        }
        current = current->next;
    }
}

// Удаление элемента перед заданным значением
void List::removeBefore(std::string target, int occurrence)
{
    // Проверка корректности ввода и наличия хотя бы 2 элементов
    if (occurrence < 0 || this->head == nullptr || this->head->next == nullptr)
        return;

    int n = 0;
    ListNode* current = this->head->next; // Начинаем со второго элемента

    // Поиск нужного вхождения target
    while (current != nullptr)
    {
        if (current->key == target)
        {
            if (n == occurrence) // Нашли нужное вхождение
            {
                ListNode* toDelete = current->prev; // Элемент для удаления
                if (toDelete == this->head) // Если удаляем голову
                {
                    removeHead();
                }
                else // Если удаляем элемент из середины
                {
                    toDelete->prev->next = current;
                    current->prev = toDelete->prev;
                    delete toDelete;
                }
                return;
            }
            n++;
        }
        current = current->next;
    }
}

// Добавление элемента после заданного значения
void List::addAfter(std::string target, std::string key, int occurrence)
{
    if (occurrence < 0) // Проверка корректности номера вхождения
        return;

    int n = 0;
    ListNode* current = this->head;

    // Поиск нужного вхождения target
    while (current != nullptr)
    {
        if (current->key == target)
        {
            if (n == occurrence) // Нашли нужное вхождение
            {
                if (current == this->tail) // Если target в конце списка
                {
                    addTail(key);
                }
                else // Если target в середине списка
                {
                    ListNode* newNode = new ListNode(key);
                    newNode->next = current->next;
                    newNode->prev = current;
                    current->next->prev = newNode;
                    current->next = newNode;
                }
                return;
            }
            n++;
        }
        current = current->next;
    }
}

// Удаление элемента после заданного значения
void List::removeAfter(std::string target, int occurrence)
{
    // Проверка корректности ввода и наличия хотя бы 2 элементов
    if (occurrence < 0 || this->head == nullptr || this->tail == nullptr)
        return;

    int n = 0;
    ListNode* current = this->head;

    // Поиск нужного вхождения target
    while (current != nullptr && current->next != nullptr) // Проверяем, что есть следующий элемент
    {
        if (current->key == target)
        {
            if (n == occurrence) // Нашли нужное вхождение
            {
                ListNode* toDelete = current->next; // Элемент для удаления

                if (toDelete == this->tail) // Если удаляем хвост
                {
                    removeTail();
                }
                else if (toDelete == this->head) 
                {
                    removeHead();
                }
                else // Если удаляем элемент из середины
                {
                    current->next = toDelete->next;
                    toDelete->next->prev = current;
                    delete toDelete;
                }
                return;
            }
            n++;
        }
        current = current->next;
    }
}